package world

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	js "persons_generator/core/storage/json_storage"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/location"
	"persons_generator/engine/entities/religion"

	"github.com/google/uuid"
)

type World struct {
	ID                        uuid.UUID
	Size                      int
	MaxDistanceValue          float64
	Year                      int
	Locations                 [][]*location.Location
	Cultures                  []*culture.Culture
	Religions                 []*religion.Religion
	CultureReligionReferences []*religion.CultureReference

	storageFolderName       string
	defaultHumanAmount      int
	defaultMalePercentage   float64
	defaultFemalePercentage float64

	populationNumber int
}

func New(
	cfg Config,
	s int,
	cultures []*culture.Culture,
	religions []*religion.Religion,
	refs []*religion.CultureReference,
) (*World, error) {
	w := &World{
		ID:                        uuid.New(),
		Size:                      s,
		MaxDistanceValue:          GetMaxDistanceValue(s),
		Year:                      0,
		Cultures:                  cultures,
		Religions:                 religions,
		CultureReligionReferences: refs,

		storageFolderName:       cfg.StorageFolderName,
		defaultHumanAmount:      cfg.DefaultHumanAmount,
		defaultMalePercentage:   cfg.DefaultMalePercentage,
		defaultFemalePercentage: cfg.DefaultFemalePercentage,
	}
	if err := w.seed(); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not seed for (*World).New")
	}
	w.populationNumber = len(w.GetPersons())

	return w, nil
}

func NewByPreferred(cfg Config, preferred *Preference) (*World, error) {
	w := &World{
		ID: uuid.New(),

		storageFolderName:       cfg.StorageFolderName,
		defaultHumanAmount:      cfg.DefaultHumanAmount,
		defaultMalePercentage:   cfg.DefaultMalePercentage,
		defaultFemalePercentage: cfg.DefaultFemalePercentage,
	}
	p, err := w.preparePreference(preferred)
	if err != nil {
		return nil, err
	}

	size, err := getSizeByPreference(p)
	if err != nil {
		return nil, err
	}
	w.Size = size
	w.CultureReligionReferences = p.ReligionCultures
	w.Cultures = religion.ExtractCulturesFromCultureReferences(p.ReligionCultures)
	w.Religions = religion.ExtractReligionsFromCultureReferences(p.ReligionCultures)

	if err := w.seed(); err != nil {
		return nil, err
	}

	return w, nil
}

func (w *World) Save() error {
	if w == nil {
		return wrapped_error.NewInternalServerError(nil, "can not save nil world")
	}

	storage := js.New(js.Config{StorageFolderName: w.storageFolderName})
	dirname := fmt.Sprintf("world_%s", w.ID.String())
	if err := storage.MkDir(dirname); err != nil {
		return wrapped_error.NewInternalServerError(err, "can not create dir for the world")
	}

	m, err := newMetadata(w).Marshal()
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not marshal metadata for the world")
	}
	metadataFilename := fmt.Sprintf("%s/metadata_%s.json", dirname, w.ID.String())
	if err := storage.Store(metadataFilename, m); err != nil {
		return err
	}

	concurrentJobs := make(chan struct{}, 5)
	var wg sync.WaitGroup
	wg.Add(w.Size * w.Size)
	errCh := make(chan error)
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] == nil {
				return wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not save loc for <nil> location (x: %d, y: %d)", x, y))
			}

			concurrentJobs <- struct{}{}
			go func(x, y int) {
				defer wg.Done()

				l, err := w.Locations[y][x].Marshal()
				if err != nil {
					errCh <- wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not save loc (x: %d, y: %d)", x, y))
					return
				}
				locFilename := fmt.Sprintf("%s/loc_%s_y%d_x%d.json", dirname, w.ID.String(), y, x)
				if err := storage.Store(locFilename, l); err != nil {
					errCh <- err
					return
				}
				<-concurrentJobs
			}(x, y)
			select {
			case err := <-errCh:
				return err
			default:
			}
		}
	}
	select {
	case err := <-errCh:
		return err
	default:
	}
	wg.Wait()
	close(concurrentJobs)
	close(errCh)

	return nil
}

func ReadByID(storageFolderName string, id uuid.UUID) (*World, error) {
	filename := strings.Join([]string{"world", id.String()}, "_") + ".json"
	b, err := js.
		New(js.Config{StorageFolderName: storageFolderName}).
		Get(filename)
	if err != nil {
		return nil, err
	}

	var out World
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

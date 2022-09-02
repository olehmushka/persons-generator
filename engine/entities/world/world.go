package world

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"

	js "persons_generator/core/storage/json_storage"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/coordinate"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/human/human"
	"persons_generator/engine/entities/location"
	"persons_generator/engine/entities/religion"

	"github.com/google/uuid"
)

type World struct {
	ID                        uuid.UUID
	Size                      int
	Locations                 [][]*location.Location
	Cultures                  []*culture.Culture
	Religions                 []*religion.Religion
	CultureReligionReferences []*religion.CultureReference

	storageFolderName       string
	defaultHumanAmount      int
	defaultMalePercentage   float64
	defaultFemalePercentage float64
}

func New(cfg Config, s int) *World {
	return &World{
		ID:   uuid.New(),
		Size: s,

		storageFolderName: cfg.StorageFolderName,
	}
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
	w.seedLocations()
	if err := w.seedCultures(); err != nil {
		return nil, err
	}
	if err := w.seedReligions(); err != nil {
		return nil, err
	}

	return w, nil
}

func (w *World) preparePreference(in *Preference) (*Preference, error) {
	if in == nil {
		hp := HumanPreference{
			Amount:       w.defaultHumanAmount,
			MaleAmount:   int(float64(w.defaultHumanAmount) * w.defaultMalePercentage),
			FemaleAmount: int(float64(w.defaultHumanAmount) * w.defaultFemalePercentage),
		}
		var culturesAmount, religionsAmount int
		switch {
		case hp.Amount < 10:
			culturesAmount = 1
			religionsAmount = 1
		case hp.Amount >= 10 && hp.Amount < 25:
			culturesAmount = 2
			religionsAmount = 2
		case hp.Amount >= 25 && hp.Amount < 50:
			culturesAmount = 3
			religionsAmount = 3
		case hp.Amount >= 50 && hp.Amount < 75:
			culturesAmount = 5
			religionsAmount = 4
		case hp.Amount >= 75 && hp.Amount < 100:
			culturesAmount = 8
			religionsAmount = 5
		case hp.Amount >= 100 && hp.Amount < 250:
			culturesAmount = 13
			religionsAmount = 6
		case hp.Amount >= 250 && hp.Amount < 500:
			culturesAmount = 21
			religionsAmount = 7
		case hp.Amount >= 500 && hp.Amount < 750:
			culturesAmount = 34
			religionsAmount = 8
		case hp.Amount >= 750 && hp.Amount < 1000:
			culturesAmount = 55
			religionsAmount = 9
		case hp.Amount > 1000:
			culturesAmount = 89
			religionsAmount = 10
		}
		cultures, err := culture.NewMany(culture.Config{StorageFolderName: w.storageFolderName}, culturesAmount, nil)
		if err != nil {
			return nil, err
		}
		references, err := religion.NewReferences(
			religion.Config{StorageFolderName: w.storageFolderName},
			religionsAmount,
			cultures,
		)
		if err != nil {
			return nil, err
		}

		return &Preference{
			Human:            hp,
			ReligionCultures: references,
		}, nil
	}

	return in, nil
}

func getSizeByPreference(pref *Preference) (int, error) {
	if pref == nil {
		return 0, wrapped_error.New(http.StatusInternalServerError, nil, "can not get size of world without preference")
	}
	squareSide := math.Sqrt(float64(pref.Human.Amount)) * 1.5

	return int(math.Ceil(squareSide)), nil
}

func (w *World) seedLocations() {
	size := w.Size
	w.Locations = make([][]*location.Location, 0, size)
	for y := 0; y < size; y++ {
		w.Locations = append(w.Locations, make([]*location.Location, 0, size))
		for x := 0; x < size; x++ {
			w.Locations[y] = append(w.Locations[y], &location.Location{
				Coordinate: &coordinate.Coordinate{X: x, Y: y},
				Population: make([]*human.Human, 0),
			})
		}
	}
}

func (w *World) seedCultures() error {
	totalLocs := w.Size * w.Size
	var (
		locsPerCulture = totalLocs / len(w.Cultures)
		reminder       = totalLocs - (locsPerCulture * len(w.Cultures))
		toFillCultures = make(map[string]int, len(w.Cultures))
	)
	for i, c := range w.Cultures {
		toFillCultures[c.Name] = locsPerCulture
		if i == 0 {
			toFillCultures[c.Name] += reminder
		}
	}

	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			var cultureName string
			for {
				randCultureName, err := culture.GetRandomCultureName(w.Cultures)
				if err != nil {
					return wrapped_error.New(http.StatusInternalServerError, err, "can not generate random culture_name")
				}
				if rem, ok := toFillCultures[randCultureName]; ok && rem > 0 {
					cultureName = randCultureName
					break
				}
			}
			w.Locations[y][x].InitCulture = culture.GetCultureByName(cultureName, w.Cultures)
			toFillCultures[cultureName]--
		}
	}

	return nil
}

func (w *World) seedReligions() error {
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] == nil || w.Locations[y][x].InitCulture == nil {
				return wrapped_error.New(http.StatusInternalServerError, nil, fmt.Sprintf("location is nil or its culture is nil (x: %d, y: %d)", x, y))
			}

			cultureName := w.Locations[y][x].InitCulture.Name
			w.Locations[y][x].InitReligion = religion.GetReligionByCultureNameFromCultureReferences(w.CultureReligionReferences, cultureName)
		}
	}

	return nil
}

func (w *World) PrintLocationCultures() {
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if x != 0 {
				fmt.Printf("\t")
			}
			fmt.Printf("%s", w.Locations[y][x].InitCulture.Name)
		}
		fmt.Printf("\n")
	}
}

func (w *World) PrintLocationReligions() {
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if x != 0 {
				fmt.Printf("\t")
			}
			fmt.Printf("%s", w.Locations[y][x].InitReligion.Name)
		}
		fmt.Printf("\n")
	}
}

func (w *World) GetHumans(params GetHumansParams) ([]*human.Human, error) {
	out := make([]*human.Human, 0, w.Size*w.Size)
	for y := 0; y < w.Size; y++ {
		for x := 0; x < w.Size; x++ {
			if w.Locations[y][x] != nil || len(w.Locations[y][x].Population) > 0 {
				out = append(out, w.Locations[y][x].Population...)
			}
		}
	}

	return out, nil
}

func (w *World) Save() error {
	if w == nil {
		return wrapped_error.New(http.StatusInternalServerError, nil, "can not save nil world")
	}

	b, err := json.MarshalIndent(w, "", " ")
	if err != nil {
		return err
	}

	return js.
		New(js.Config{StorageFolderName: w.storageFolderName}).
		Store(strings.Join([]string{"world", w.ID.String()}, "_")+".json", b)
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

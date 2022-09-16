package world

import (
	"encoding/json"
	"strings"

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
}

func New(
	cfg Config,
	s int,
	cultures []*culture.Culture,
	religions []*religion.Religion,
) (*World, error) {
	w := &World{
		ID:               uuid.New(),
		Size:             s,
		MaxDistanceValue: GetMaxDistanceValue(s),
		Year:             0,
		Cultures:         cultures,
		Religions:        religions,

		storageFolderName:       cfg.StorageFolderName,
		defaultHumanAmount:      cfg.DefaultHumanAmount,
		defaultMalePercentage:   cfg.DefaultMalePercentage,
		defaultFemalePercentage: cfg.DefaultFemalePercentage,
	}
	if err := w.seed(); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not seed for (*World).New")
	}

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

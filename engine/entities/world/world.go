package world

import (
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/location"
	"persons_generator/engine/entities/religion"

	"github.com/google/uuid"
)

type World struct {
	ID                        uuid.UUID
	Size                      int
	MaxPersonsNumberPerLoc    int
	MaxDistanceValue          float64
	Year                      int
	Locations                 [][]*location.Location
	DeathWorldLocations       [][]*location.Location
	Cultures                  []*culture.Culture
	Religions                 []*religion.Religion
	CultureReligionReferences []*religion.CultureReference

	defaultHumanAmount      int
	defaultMalePercentage   float64
	defaultFemalePercentage float64

	PopulationNumber     int
	DeadPopulationNumber int
	religionsSimilarity  map[string]float64
	culturesSimilarity   map[string]float64
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
		MaxPersonsNumberPerLoc:    750,

		defaultHumanAmount:      cfg.DefaultHumanAmount,
		defaultMalePercentage:   cfg.DefaultMalePercentage,
		defaultFemalePercentage: cfg.DefaultFemalePercentage,
	}
	if err := w.seed(); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not seed for (*World).New")
	}
	w.PopulationNumber = w.CalculatePersonsNumber()

	return w, nil
}

func NewByPreferred(cfg Config, preferred *Preference) (*World, error) {
	w := &World{
		ID: uuid.New(),

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

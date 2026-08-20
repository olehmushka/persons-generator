// Package world models the year-by-year population simulation: a square
// grid of Locations, each seeded with a starting population and an initial
// culture/religion. RunYear advances every location's Population by one
// year -- aging, marriage, births, and deaths -- and RunWorld drives that
// loop until a target stop year, reporting progress on a channel.
package world

import (
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/location"
	"persons_generator/engine/entities/religion"

	"github.com/google/uuid"
)

// World is a Size x Size grid of Locations. Locations holds the living
// population; DeathWorldLocations mirrors it 1:1 and accumulates anyone who
// has died, keyed by the same coordinates they died at.
type World struct {
	ID                        string
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
	id string,
	s int,
	cultures []*culture.Culture,
	religions []*religion.Religion,
	refs []*religion.CultureReference,
) (*World, error) {
	w := &World{
		ID:                        id,
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
		ID: uuid.New().String(),

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

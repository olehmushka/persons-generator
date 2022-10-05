package entities

import (
	"github.com/google/uuid"
)

type World struct {
	ID                        uuid.UUID           `json:"id"`
	Size                      int                 `json:"size"`
	MaxPersonsNumberPerLoc    int                 `json:"max_persons_number_per_loc"`
	MaxDistanceValue          float64             `json:"max_distance_value"`
	Year                      int                 `json:"year"`
	CulturesIDs               []uuid.UUID         `json:"cultures_ids"`
	ReligionsIDs              []uuid.UUID         `json:"religions_ids"`
	CultureReligionReferences []*CultureReference `json:"culture_religion_references"`
	PopulationNumber          int                 `json:"population_number"`
	DeadPopulationNumber      int                 `json:"dead_population_number"`
	Duration                  string              `json:"duration"`
}

type CultureReference struct {
	ReligionID uuid.UUID `json:"religion_id"`
	CultureID  uuid.UUID `json:"culture_id"`
}

package entities

type World struct {
	ID                        string              `json:"id"`
	Size                      int                 `json:"size"`
	MaxPersonsNumberPerLoc    int                 `json:"max_persons_number_per_loc"`
	MaxDistanceValue          float64             `json:"max_distance_value"`
	Year                      int                 `json:"year"`
	CulturesIDs               []string            `json:"cultures_ids"`
	ReligionsIDs              []string            `json:"religions_ids"`
	CultureReligionReferences []*CultureReference `json:"culture_religion_references"`
	PopulationNumber          int                 `json:"population_number"`
	DeadPopulationNumber      int                 `json:"dead_population_number"`
	Duration                  string              `json:"duration"`
}

type CultureReference struct {
	ReligionID string `json:"religion_id"`
	CultureID  string `json:"culture_id"`
}

package person

type Chronology struct {
	BirthYear int     `json:"birth_year"`
	DeathYear int     `json:"death_year"`
	Events    []Event `json:"events"`
}

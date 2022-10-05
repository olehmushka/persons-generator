package mq

type RunAndSaveWorldPayload struct {
	WorldID             string            `json:"world_id"`
	StopYear            int               `json:"stop_year"`
	Amount              int               `json:"amount"`
	MaleAmount          int               `json:"male_amount"`
	FemaleAmount        int               `json:"female_amount"`
	ReligionCultureRels map[string]string `json:"religion_culture_rels"`
}

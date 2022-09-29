package mq

import "github.com/google/uuid"

type RunAndSaveWorldPayload struct {
	WorldID             uuid.UUID               `json:"world_id"`
	StopYear            int                     `json:"stop_year"`
	Amount              int                     `json:"amount"`
	MaleAmount          int                     `json:"male_amount"`
	FemaleAmount        int                     `json:"female_amount"`
	ReligionCultureRels map[uuid.UUID]uuid.UUID `json:"religion_culture_rels"`
}

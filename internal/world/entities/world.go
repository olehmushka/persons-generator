package entities

import "github.com/google/uuid"

type World struct {
	ID                  uuid.UUID   `json:"id"`
	PersonsAmount       int         `json:"persons_amount"`
	MalePersonsAmount   int         `json:"male_persons_amount"`
	FemalePersonsAmount int         `json:"female_persons_amount"`
	ReligionIDs         []uuid.UUID `json:"religion_ids"`
	CultureIDs          []uuid.UUID `json:"culture_ids"`
}

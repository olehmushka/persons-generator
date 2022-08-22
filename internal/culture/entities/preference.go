package entities

import "github.com/google/uuid"

type CulturePreference struct {
	Names  []string `json:"names"`
	Amount int      `json:"amount"`
	Kind   string   `json:"kind"`
}

type ReligionPreference struct {
	CultureIDs []uuid.UUID
}

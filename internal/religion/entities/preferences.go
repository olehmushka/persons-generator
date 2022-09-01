package entities

import "github.com/google/uuid"

type Preference struct {
	CultureIDs []uuid.UUID
	Amount     int
}

package services

import (
	engineWorld "persons_generator/engine/entities/world"
	"persons_generator/internal/world/entities"
)

func serializeWorld(in *engineWorld.World) *entities.World {
	if in == nil {
		return nil
	}

	return &entities.World{
		ID: in.ID,
		// PersonsAmount       int         `json:"persons_amount"`
		// MalePersonsAmount   int         `json:"male_persons_amount"`
		// FemalePersonsAmount int         `json:"female_persons_amount"`
		// ReligionIDs         []uuid.UUID `json:"religion_ids"`
		// CultureIDs          []uuid.UUID `json:"culture_ids"`
	}
}

package services

import (
	engineWorld "persons_generator/engine/entities/world"
	"persons_generator/internal/world/entities"

	"github.com/google/uuid"
)

func serializeWorld(in *engineWorld.World) *entities.World {
	if in == nil {
		return nil
	}

	religions := make([]uuid.UUID, 0, len(in.Religions))
	for _, r := range in.Religions {
		if r == nil {
			continue
		}
		religions = append(religions, r.ID)
	}

	cultures := make([]uuid.UUID, 0, len(in.Cultures))
	for _, c := range in.Cultures {
		if c == nil {
			continue
		}
		cultures = append(cultures, c.ID)
	}

	return &entities.World{
		ID: in.ID,
		// PersonsAmount       int         `json:"persons_amount"`
		// MalePersonsAmount   int         `json:"male_persons_amount"`
		// FemalePersonsAmount int         `json:"female_persons_amount"`
		ReligionIDs: religions,
		CultureIDs:  cultures,
	}
}

package engine

import (
	"persons_generator/engine/entities/religion"
	"persons_generator/internal/religion/entities"
)

func serializeReligions(in []*religion.Religion) []*entities.Religion {
	out := make([]*entities.Religion, len(in))
	for i := range out {
		out[i] = serializeReligion(in[i])
	}

	return out
}

func serializeReligion(in *religion.Religion) *entities.Religion {
	if in == nil {
		return nil
	}

	return &entities.Religion{}
}

package services

import (
	"encoding/json"
	"persons_generator/core/wrapped_error"
	engineWorld "persons_generator/engine/entities/world"
	"persons_generator/internal/world/entities"
)

func serializeSerializedWorld(in *engineWorld.SerializedWorld) (*entities.World, error) {
	if in == nil {
		return nil, nil
	}

	b, err := json.Marshal(in)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not marshal in serialize world")
	}

	var out entities.World
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not unmarshal in serialize world")
	}

	return &out, nil
}

func serializeSerializedWorlds(in []*engineWorld.SerializedWorld) ([]*entities.World, error) {
	out := make([]*entities.World, len(in))
	for i := range out {
		var err error
		out[i], err = serializeSerializedWorld(in[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can no serialize worlds")
		}
	}

	return out, nil
}

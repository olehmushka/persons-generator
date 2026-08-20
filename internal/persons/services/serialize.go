package services

import (
	"encoding/json"
	"persons_generator/core/wrapped_error"
	enginePerson "persons_generator/engine/entities/person"
	"persons_generator/internal/persons/entities"
)

func serializeSerializedPerson(in *enginePerson.SerializedPerson) (*entities.Person, error) {
	if in == nil {
		return nil, nil
	}

	b, err := json.Marshal(in)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not marshal in serialize person")
	}

	var out entities.Person
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not unmarshal in serialize person")
	}

	return &out, nil
}

func serializeSerializedPeople(in []*enginePerson.SerializedPerson) ([]*entities.Person, error) {
	out := make([]*entities.Person, len(in))
	for i := range out {
		var err error
		out[i], err = serializeSerializedPerson(in[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can no serialize persons")
		}
	}

	return out, nil
}

package http

import (
	"encoding/json"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
	personsEntities "persons_generator/internal/persons/entities"
)

func serializeCultures(cultures []*culture.SerializedCulture) ([]*SerializedCulture, error) {
	out := make([]*SerializedCulture, len(cultures))
	for i := range out {
		var err error
		out[i], err = serializeCulture(cultures[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can no serialize cultures")
		}
	}

	return out, nil
}

func serializeCulture(in *culture.SerializedCulture) (*SerializedCulture, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not marshal in serialize culture")
	}

	var out SerializedCulture
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not unmarshal in serialize culture")
	}

	return &out, nil
}

func serializeReligions(religions []*religion.SerializedReligion) ([]*SerializedReligion, error) {
	out := make([]*SerializedReligion, len(religions))
	for i := range out {
		var err error
		out[i], err = serializeReligion(religions[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can no serialize religions")
		}
	}

	return out, nil
}

func serializeReligion(in *religion.SerializedReligion) (*SerializedReligion, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not marshal in serialize religion")
	}

	var out SerializedReligion
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not unmarshal in serialize religion")
	}

	return &out, nil
}

func serializePersons(in []*personsEntities.Person) []*Person {
	out := make([]*Person, len(in))
	for i := range out {
		out[i] = serializePerson(in[i])
	}

	return out
}

func serializePerson(in *personsEntities.Person) *Person {
	return &Person{}
}

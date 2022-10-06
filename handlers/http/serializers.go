package http

import (
	"encoding/json"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
	languageEntities "persons_generator/internal/language/entities"
	personsEntities "persons_generator/internal/persons/entities"
	worldEntities "persons_generator/internal/world/entities"
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

func serializeLanguages(langs []*languageEntities.Language) ([]*Language, error) {
	out := make([]*Language, len(langs))
	for i := range out {
		var err error
		out[i], err = serializeLanguage(langs[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can no serialize languages")
		}
	}

	return out, nil
}

func serializeLanguage(in *languageEntities.Language) (*Language, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not marshal in serialize language")
	}

	var out Language
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not unmarshal in serialize language")
	}

	return &out, nil
}

func serializeLanguageSubfamilies(langs []*languageEntities.Subfamily) ([]*LanguageSubfamily, error) {
	out := make([]*LanguageSubfamily, len(langs))
	for i := range out {
		var err error
		out[i], err = serializeLanguageSubfamily(langs[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can no serialize language  subfamilies")
		}
	}

	return out, nil
}

func serializeLanguageSubfamily(in *languageEntities.Subfamily) (*LanguageSubfamily, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not marshal in serialize language subfamily")
	}

	var out LanguageSubfamily
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not unmarshal in serialize language subfamily")
	}

	return &out, nil
}

func serializeWorlds(langs []*worldEntities.World) ([]*World, error) {
	out := make([]*World, len(langs))
	for i := range out {
		var err error
		out[i], err = serializeWorld(langs[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can no serialize worlds")
		}
	}

	return out, nil
}

func serializeWorld(in *worldEntities.World) (*World, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not marshal in serialize world")
	}

	var out World
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not unmarshal in serialize world")
	}

	return &out, nil
}

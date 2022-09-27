package location

import (
	"encoding/json"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/coordinate"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/person"
	"persons_generator/engine/entities/religion"

	"github.com/google/uuid"
)

type Location struct {
	Coordinate *coordinate.Coordinate
	Population []*person.Person

	InitCulture  *culture.Culture
	InitReligion *religion.Religion
}

type loc struct {
	Coordinate *coordinate.Coordinate     `json:"coordinate"`
	Population []*person.SerializedPerson `json:"population"`
	CultureID  uuid.UUID                  `json:"culture_id"`
	ReligionID uuid.UUID                  `json:"religion_id"`
}

func newLoc(in *Location) *loc {
	if in == nil {
		return nil
	}

	population := make([]*person.SerializedPerson, len(in.Population))
	for i := range population {
		population[i] = in.Population[i].Serialize()
	}

	return &loc{
		Coordinate: in.Coordinate,
		Population: population,
		CultureID:  in.InitCulture.ID,
		ReligionID: in.InitReligion.ID,
	}
}

func (l *loc) Marshal() ([]byte, error) {
	if l == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not marshal <nil> loc")
	}

	b, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not marshal loc")
	}

	return b, nil
}

func (l *Location) Marshal() ([]byte, error) {
	if l == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not marshal <nil> location")
	}

	return newLoc(l).Marshal()
}

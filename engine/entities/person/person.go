package person

import (
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/person/human"
	"persons_generator/engine/entities/person/traits"
	"persons_generator/engine/entities/religion"

	"github.com/google/uuid"
)

type Person struct {
	ID       uuid.UUID          `json:"id"`
	OwnName  string             `json:"own_name"`
	Culture  *culture.Culture   `json:"culture,omitempty"`
	Religion *religion.Religion `json:"religion,omitempty"`
	Human    *human.Human       `json:"human"`

	Traits     []*traits.Trait `json:"traits"`
	Spouces    []*Person       `json:"spouces"`
	Chronology Chronology      `json:"chronology"`
}

func New(h *human.Human, c *culture.Culture, r *religion.Religion, year int) (*Person, error) {
	if h == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "person can not be generated without human inside")
	}
	if c == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "person can not be generated without culture inside")
	}
	if r == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "person can not be generated without religion inside")
	}
	p := &Person{
		ID:       uuid.New(),
		Human:    h,
		Culture:  c,
		Religion: r,
		Chronology: Chronology{
			BirthYear: year,
			DeathYear: -1,
			Events:    []Event{},
		},
	}
	ownName, err := c.Language.GetOwnName(h.Sex)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not get own_name for person by sex (sex=%s)", h.Sex))
	}
	p.OwnName = ownName

	return p, nil
}

func UniquePersons(persons []*Person) []*Person {
	if len(persons) <= 1 {
		return persons
	}

	preOut := make(map[uuid.UUID]*Person)
	for _, c := range persons {
		preOut[c.ID] = c
	}

	out := make([]*Person, 0, len(preOut))
	for _, c := range preOut {
		out = append(out, c)
	}

	return out
}

func RemovePersonFromSliceByID(persons []*Person, id uuid.UUID) []*Person {
	if len(persons) == 0 {
		return persons
	}
	out := make([]*Person, 0, len(persons))
	for _, p := range persons {
		if p.ID == id {
			continue
		}
		out = append(out, p)
	}

	return out
}

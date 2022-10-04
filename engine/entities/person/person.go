package person

import (
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/coordinate"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/human"
	"persons_generator/engine/entities/person/human/presets"
	"persons_generator/engine/entities/person/traits"
	"persons_generator/engine/entities/religion"

	"github.com/google/uuid"
)

type Person struct {
	ID         uuid.UUID              `json:"id"`
	OwnName    string                 `json:"own_name"`
	Culture    *culture.Culture       `json:"culture,omitempty"`
	Religion   *religion.Religion     `json:"religion,omitempty"`
	Human      *human.Human           `json:"human"`
	Father     *Person                `json:"father"`
	Mother     *Person                `json:"mother"`
	Coordinate *coordinate.Coordinate `json:"coordinate"`

	Traits        []*traits.Trait `json:"traits"`
	Spouces       []*Person       `json:"spouces"`
	Chronology    Chronology      `json:"chronology"`
	Metadata      Metadata        `json:"metadata"`
	PrebirthStats *PrebirthStats  `json:"prebirth_stats"`
}

func New(
	h *human.Human,
	c *culture.Culture,
	r *religion.Religion,
	year int,
	father,
	mother *Person,
	coor *coordinate.Coordinate,
) (*Person, error) {
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
		Metadata:   Metadata{},
		Father:     father,
		Mother:     mother,
		Coordinate: coor,
	}
	ownName, err := c.Language.GetOwnName(h.Sex)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not get own_name for person by sex (sex=%s)", h.Sex))
	}
	p.OwnName = ownName

	return p, nil
}

func NewBase(c *culture.Culture, r *religion.Religion, coor *coordinate.Coordinate) (*Person, error) {
	if c == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "person can not be generated without culture inside")
	}
	if r == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "person can not be generated without religion inside")
	}

	sex, err := gender.GetRandomSex()
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate sex for new base person")
	}

	gene, err := presets.GetPresetByCulture(c)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate gene by culture for new base person")
	}

	h, err := human.New(sex, gene, nil, nil)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate human from gene for new base person")
	}

	return New(h, c, r, 0, nil, nil, coor)
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

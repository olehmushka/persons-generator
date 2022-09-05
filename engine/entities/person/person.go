package person

import (
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/person/human"
	"persons_generator/engine/entities/religion"
)

type Person struct {
	OwnName  string             `json:"own_name"`
	Culture  *culture.Culture   `json:"culture,omitempty"`
	Religion *religion.Religion `json:"religion,omitempty"`
	Human    *human.Human       `json:"human"`
}

func New(h *human.Human, c *culture.Culture, r *religion.Religion) (*Person, error) {
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
		Human:    h,
		Culture:  c,
		Religion: r,
	}
	ownName, err := c.Language.GetOwnName(h.Sex)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not get own_name for person by sex (sex=%s)", h.Sex))
	}
	p.OwnName = ownName

	return p, nil
}

func (p *Person) HaveSex(other *Person) error {
	return nil
}

package human

import (
	"encoding/json"
	"fmt"

	"persons_generator/entities/culture"
	g "persons_generator/entities/gender"
	"persons_generator/entities/human/body"
	"persons_generator/entities/human/gene"
	"persons_generator/entities/religion"
)

type Human struct {
	OwnName  string
	Culture  *culture.Culture
	Religion *religion.Religion
	Sex      g.Sex
	Gene     gene.Gene

	Body *body.Body
}

func (h Human) Bytes() []byte {
	zero := Human{}
	if h == zero {
		return nil
	}
	if out, err := json.Marshal(h); err == nil {
		return out
	}

	return nil
}

func (h Human) Print() {
	zero := Human{}
	if h == zero {
		fmt.Printf("\nHuman is nil\n")
	}

	if out, err := json.Marshal(h); err == nil {
		fmt.Printf("\n%s\n", string(out))
	}
}

func NewHuman(
	c *culture.Culture,
	r *religion.Religion,
	sex g.Sex,
	g gene.Gene,
) (*Human, error) {
	b, err := g.Produce()
	if err != nil {
		return nil, err
	}
	var h Human
	if err := json.Unmarshal(b.Bytes(), &h); err != nil {
		return nil, err
	}
	h.OwnName = c.Language.GetOwnName(sex)
	h.Culture = c
	h.Religion = r
	h.Gene = g
	h.Sex = sex

	return &h, nil
}

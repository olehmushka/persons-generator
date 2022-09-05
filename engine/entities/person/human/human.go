package human

import (
	"encoding/json"
	"fmt"

	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/body"
	"persons_generator/engine/entities/person/gene"
	"persons_generator/engine/entities/person/psycho"
)

type Human struct {
	Sex    g.Sex     `json:"sex"`
	Gene   gene.Gene `json:"gene"`
	Age    int       `json:"age"`
	Father *Human    `json:"father,omitempty"`
	Mother *Human    `json:"mather,omitempty"`

	Body   *body.Body     `json:"body"`
	Psycho *psycho.Psycho `json:"psycho"`
}

func New(sex g.Sex, g gene.Gene, f, m *Human) (*Human, error) {
	b, err := g.Produce(sex)
	if err != nil {
		return nil, err
	}
	var h Human
	if err := json.Unmarshal(b.Bytes(), &h); err != nil {
		return nil, err
	}
	h.Gene = g
	h.Sex = sex
	h.Father = f
	h.Mother = m

	return &h, nil
}

func (h Human) Zero() bool {
	return h == Human{}
}

func (h Human) Bytes() []byte {
	if h.Zero() {
		return nil
	}
	if out, err := json.Marshal(h); err == nil {
		return out
	}

	return nil
}

func (h Human) Print() {
	if h.Zero() {
		fmt.Printf("\nHuman is nil\n")
	}

	if out, err := json.Marshal(h); err == nil {
		fmt.Printf("\n%s\n", string(out))
	}
}

func (h *Human) ProduceChildren(father *Human) ([]*Human, error) {
	if father == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not produce children without father")
	}
	if h == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not produce children without mother")
	}

	return nil, nil
}

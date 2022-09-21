package human

import (
	"encoding/json"

	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/body"
	"persons_generator/engine/entities/person/gene"
	"persons_generator/engine/entities/person/psycho"
)

type Human struct {
	Sex      g.Sex     `json:"sex"`
	Gene     gene.Gene `json:"gene"`
	Age      int       `json:"age"`
	Metadata Metadata  `json:"metadata"`

	Father           *Human   `json:"father,omitempty"`
	Mother           *Human   `json:"mother,omitempty"`
	Children         []*Human `json:"children,omitempty"`
	PreBirthChildren []*Human `json:"pre_birth_children,omitempty"`

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
	h.Metadata = Metadata{}

	return &h, nil
}

func (h Human) Zero() bool {
	if h.Sex != "" {
		return false
	}
	if h.Gene != nil {
		return false
	}
	if h.Age != 0 {
		return false
	}
	if h.Father != nil {
		return false
	}
	if h.Mother != nil {
		return false
	}
	if len(h.Children) != 0 {
		return false
	}
	if len(h.PreBirthChildren) != 0 {
		return false
	}
	if h.Body != nil {
		return false
	}
	if h.Psycho != nil {
		return false
	}

	return true
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

func GetSimilarityCoef(h1, h2 *Human) float64 {
	return 1
}

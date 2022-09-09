package human

import (
	"encoding/json"
	"fmt"

	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/body"
	"persons_generator/engine/entities/person/psycho"

	"persons_generator/engine/entities/person/gene"
)

type HumanGene struct {
	T string `json:"t"`

	BodyGene   gene.Gene `json:"body_gene"`
	PsychoGene gene.Gene `json:"psycho_gene"`
}

func NewGene(bodyGene gene.Gene, psychoGene gene.Gene) gene.Gene {
	return &HumanGene{
		T: "human_gene",

		BodyGene:   bodyGene,
		PsychoGene: psychoGene,
	}
}

func (g *HumanGene) Type() string {
	return g.T
}

func (g *HumanGene) Produce(sex g.Sex) (gene.Byteble, error) {
	b, err := g.BodyGene.Produce(sex)
	if err != nil {
		return Human{}, wrapped_error.NewInternalServerError(err, "can not produce body from gene")
	}
	var bodyVal body.Body
	if err := json.Unmarshal(b.Bytes(), &bodyVal); err != nil {
		return Human{}, wrapped_error.NewInternalServerError(err, "can not unmarshal body produced from gene")
	}

	p, err := g.PsychoGene.Produce(sex)
	if err != nil {
		return Human{}, wrapped_error.NewInternalServerError(err, "can not produce psycho from gene")
	}
	var psychoVal psycho.Psycho
	if err := json.Unmarshal(p.Bytes(), &psychoVal); err != nil {
		return Human{}, wrapped_error.NewInternalServerError(err, "can not unmarshal psycho produced from gene")
	}

	return Human{
		Sex:    sex,
		Body:   &bodyVal,
		Psycho: &psychoVal,
	}, nil
}

func (g *HumanGene) Children() []gene.Gene {
	return []gene.Gene{
		g.BodyGene,
		g.PsychoGene,
	}
}

func (g *HumanGene) Bytes() []byte {
	return nil
}

func (g *HumanGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> human genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}
	var (
		hostChildren  = g.Children()
		guestChildren = in.Children()
	)
	if len(g.Children()) != len(in.Children()) {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same number of children for human gene (host_children_number=%d, guest_children_number=%d)", len(hostChildren), len(guestChildren)))
	}

	args := make([]gene.Gene, 0, len(hostChildren))
	for i, child := range hostChildren {
		arg, err := child.Pair(guestChildren[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not pair children for psycho gene")
		}
		args = append(args, arg)
	}

	return NewGene(args[0], args[1]), nil
}

package human

import (
	"encoding/json"

	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/body"

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

	return Human{
		Sex:  sex,
		Body: &bodyVal,
	}, nil
}

func (g *HumanGene) Children() []gene.Gene {
	return []gene.Gene{
		g.BodyGene,
	}
}

func (g *HumanGene) Bytes() []byte {
	return nil
}

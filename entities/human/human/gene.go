package human

import (
	"encoding/json"

	"persons_generator/entities/human/body"

	"persons_generator/entities/human/gene"
)

type HumanGene struct {
	t string

	BodyGene gene.Gene
}

func NewHumanGene(bodyGene gene.Gene) gene.Gene {
	return &HumanGene{
		t:        "human_gene",
		BodyGene: bodyGene,
	}
}

func (g *HumanGene) Type() string {
	return g.t
}

func (g *HumanGene) Produce() (gene.Byteble, error) {
	b, err := g.BodyGene.Produce()
	if err != nil {
		return Human{}, err
	}
	var bodyVal body.Body
	if err := json.Unmarshal(b.Bytes(), &bodyVal); err != nil {
		return Human{}, err
	}

	return Human{
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

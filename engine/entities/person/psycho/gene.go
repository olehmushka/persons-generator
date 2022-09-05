package psycho

import (
	"persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type PsychoGene struct {
	T string `json:"t"`
}

func NewPsychoGene() gene.Gene {
	return &PsychoGene{
		T: "psycho_gene",
	}
}

func (g *PsychoGene) Type() string {
	return g.T
}

func (g *PsychoGene) Produce(sex gender.Sex) (gene.Byteble, error) {
	return Psycho{}, nil
}

func (g *PsychoGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *PsychoGene) Bytes() []byte {
	return nil
}

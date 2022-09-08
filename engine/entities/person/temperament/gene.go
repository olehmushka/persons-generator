package temperament

import (
	"persons_generator/core/tools"
	"persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
	pm "persons_generator/engine/probability_machine"
)

type TemperamentGene struct {
	T string `json:"t"`

	Stats map[string]float64 `json:"stats"`
}

func NewTemperamentGene(stats map[string]float64) gene.Gene {
	return &TemperamentGene{
		T: "temperament_gene",

		Stats: stats,
	}
}

func (g *TemperamentGene) Type() string {
	return g.T
}

func (g *TemperamentGene) Produce(sex gender.Sex) (gene.Byteble, error) {
	return Temperament{
		Name: tools.Search(AllTemperamentTypes, func(e string) string { return e }, pm.GetRandomFromSeveral(g.Stats)),
	}, nil
}

func (g *TemperamentGene) Children() []gene.Gene {
	return nil
}

func (g *TemperamentGene) Bytes() []byte {
	return nil
}

func (g *TemperamentGene) Pair(in gene.Gene) (gene.Gene, error) {
	return in, nil
}

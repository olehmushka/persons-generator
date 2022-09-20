package temperament

import (
	"fmt"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
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
	t, err := pm.GetRandomFromSeveral(g.Stats)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate temperament")
	}

	return Temperament{
		Name: tools.Search(AllTemperamentTypes, func(e string) string { return e }, t),
	}, nil
}

func (g *TemperamentGene) Children() []gene.Gene {
	return nil
}

func (g *TemperamentGene) Bytes() []byte {
	return nil
}

func (g *TemperamentGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> temperament genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}

	inStr, ok := in.(*TemperamentGene)
	if !ok {
		return nil, wrapped_error.NewInternalServerError(nil, "can not case input gene to TemperamentGene")
	}

	return NewTemperamentGene(tools.MergeMaps(g.Stats, inStr.Stats)), nil
}

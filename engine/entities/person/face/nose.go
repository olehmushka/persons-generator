package face

import (
	"encoding/json"
	"fmt"

	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
	pm "persons_generator/engine/probability_machine"
)

type Nose struct {
	Type NoseType `json:"type"`
}

func (n Nose) Zero() bool {
	return n == Nose{}
}

func (n Nose) Bytes() []byte {
	if n.Zero() {
		return nil
	}

	if out, err := json.Marshal(n); err == nil {
		return out
	}

	return nil
}

type NoseGene struct {
	T string `json:"t"`

	NoseTypeStats map[string]float64 `json:"nose_type_stats"`
}

func NewNoseGene(stats map[string]float64) gene.Gene {
	return &NoseGene{
		T: "nose_gene",

		NoseTypeStats: stats,
	}
}

func (g *NoseGene) Type() string {
	return g.T
}

func (g *NoseGene) Produce(sex g.Sex) (gene.Byteble, error) {
	return Nose{Type: NoseType(pm.GetRandomFromSeveral(g.NoseTypeStats))}, nil
}

func (g *NoseGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *NoseGene) Bytes() []byte {
	return nil
}

func (g *NoseGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> nose genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}

	inStr, ok := in.(*NoseGene)
	if !ok {
		return nil, wrapped_error.NewInternalServerError(nil, "can not case input gene to NoseGene")
	}

	return NewNoseGene(tools.MergeMaps(g.NoseTypeStats, inStr.NoseTypeStats)), nil
}

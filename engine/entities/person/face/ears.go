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

type Ears struct {
	Type EarsType `json:"type"`
}

func (n Ears) Zero() bool {
	return n == Ears{}
}

func (n Ears) Bytes() []byte {
	if n.Zero() {
		return nil
	}

	if out, err := json.Marshal(n); err == nil {
		return out
	}

	return nil
}

type EarsGene struct {
	T string `json:"t"`

	EarsTypeStats map[string]float64 `json:"ears_type_stats"`
}

func NewEarsGene(stats map[string]float64) gene.Gene {
	return &EarsGene{
		T: "ears_gene",

		EarsTypeStats: stats,
	}
}

func (g *EarsGene) Type() string {
	return g.T
}

func (g *EarsGene) Produce(sex g.Sex) (gene.Byteble, error) {
	return Ears{Type: EarsType(pm.GetRandomFromSeveral(g.EarsTypeStats))}, nil
}

func (g *EarsGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *EarsGene) Bytes() []byte {
	return nil
}

func (g *EarsGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> ears genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}

	inStr, ok := in.(*EarsGene)
	if !ok {
		return nil, wrapped_error.NewInternalServerError(nil, "can not case input gene to EarsGene")
	}

	return NewEarsGene(tools.MergeMaps(g.EarsTypeStats, inStr.EarsTypeStats)), nil
}

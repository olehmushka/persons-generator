package face

import (
	"encoding/json"

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

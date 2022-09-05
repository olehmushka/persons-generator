package face

import (
	"encoding/json"

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

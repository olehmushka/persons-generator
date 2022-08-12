package face

import (
	"encoding/json"

	"persons_generator/engine/entities/human/gene"
	pm "persons_generator/engine/probability_machine"
)

type Nose struct {
	Type NoseType
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
	t             string
	NoseTypeStats map[string]float64
}

func NewNoseGene(stats map[string]float64) gene.Gene {
	return &NoseGene{
		t:             "nose_gene",
		NoseTypeStats: stats,
	}
}

func (g *NoseGene) Type() string {
	return g.t
}

func (g *NoseGene) Produce() (gene.Byteble, error) {
	return Nose{Type: NoseType(pm.GetRandomFromSeveral(g.NoseTypeStats))}, nil
}

func (g *NoseGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *NoseGene) Bytes() []byte {
	return nil
}

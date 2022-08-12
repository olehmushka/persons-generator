package face

import (
	"encoding/json"

	"persons_generator/entities/human/gene"
	pm "persons_generator/probability_machine"
)

type Ears struct {
	Type EarsType
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
	t             string
	EarsTypeStats map[string]float64
}

func NewEarsGene(stats map[string]float64) gene.Gene {
	return &EarsGene{
		t:             "ears_gene",
		EarsTypeStats: stats,
	}
}

func (g *EarsGene) Type() string {
	return g.t
}

func (g *EarsGene) Produce() (gene.Byteble, error) {
	return Ears{Type: EarsType(pm.GetRandomFromSeveral(g.EarsTypeStats))}, nil
}

func (g *EarsGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *EarsGene) Bytes() []byte {
	return nil
}

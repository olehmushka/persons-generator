package face

import (
	"encoding/json"

	"persons_generator/engine/entities/human/gene"
	pm "persons_generator/engine/probability_machine"
)

type Mouth struct {
	LipsType LipsType
}

func (n Mouth) Zero() bool {
	return n == Mouth{}
}

func (n Mouth) Bytes() []byte {
	if n.Zero() {
		return nil
	}

	if out, err := json.Marshal(n); err == nil {
		return out
	}

	return nil
}

type MouthGene struct {
	t             string
	LipsTypeStats map[string]float64
}

func NewMouthGene(lipsStats map[string]float64) gene.Gene {
	return &MouthGene{
		t:             "mouth_gene",
		LipsTypeStats: lipsStats,
	}
}

func (g *MouthGene) Type() string {
	return g.t
}

func (g *MouthGene) Produce() (gene.Byteble, error) {
	return Mouth{LipsType: LipsType(pm.GetRandomFromSeveral(g.LipsTypeStats))}, nil
}

func (g *MouthGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *MouthGene) Bytes() []byte {
	return nil
}

package face

import (
	"encoding/json"

	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
	pm "persons_generator/engine/probability_machine"
)

type Mouth struct {
	LipsType LipsType `json:"lips_type"`
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
	T string `json:"t"`

	LipsTypeStats map[string]float64 `json:"lips_type_stats"`
}

func NewMouthGene(lipsStats map[string]float64) gene.Gene {
	return &MouthGene{
		T: "mouth_gene",

		LipsTypeStats: lipsStats,
	}
}

func (g *MouthGene) Type() string {
	return g.T
}

func (g *MouthGene) Produce(sex g.Sex) (gene.Byteble, error) {
	return Mouth{LipsType: LipsType(pm.GetRandomFromSeveral(g.LipsTypeStats))}, nil
}

func (g *MouthGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *MouthGene) Bytes() []byte {
	return nil
}

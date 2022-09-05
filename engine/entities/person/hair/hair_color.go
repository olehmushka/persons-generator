package hair

import (
	"encoding/json"
	"persons_generator/core/tools"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/gene"
	pm "persons_generator/engine/probability_machine"
)

type HairColor color.Color

func (c HairColor) Zero() bool {
	return c == HairColor{}
}

func (n HairColor) Bytes() []byte {
	if n.Zero() {
		return nil
	}

	if out, err := json.Marshal(n); err == nil {
		return out
	}

	return nil
}

func (n HairColor) Print() {
}

type HairColorGene struct {
	T string `json:"t"`

	Stats map[string]float64 `json:"stats"`
}

func NewHairColorGene(stats map[string]float64) gene.Gene {
	return &HairColorGene{
		T: "hair_color_gene",

		Stats: stats,
	}
}

func (g *HairColorGene) Type() string {
	return g.T
}

func (g *HairColorGene) Produce(sex g.Sex) (gene.Byteble, error) {
	return tools.Search(
		AllHairColors,
		func(e HairColor) string { return e.Name },
		pm.GetRandomFromSeveral(g.Stats),
	), nil
}

func (g *HairColorGene) Children() []gene.Gene {
	return nil
}

func (g *HairColorGene) Bytes() []byte {
	return nil
}

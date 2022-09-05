package face

import (
	"encoding/json"

	"persons_generator/core/tools"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/gene"
	pm "persons_generator/engine/probability_machine"
)

type EyesColor color.Color

func (n EyesColor) Zero() bool {
	return n == EyesColor{}
}

func (n EyesColor) Bytes() []byte {
	if n.Zero() {
		return nil
	}

	if out, err := json.Marshal(n); err == nil {
		return out
	}

	return nil
}

func (n EyesColor) Print() {
}

type EyesColorGene struct {
	T string `json:"t"`

	Stats map[string]float64 `json:"stats"`
}

func NewEyesColorGene(stats map[string]float64) gene.Gene {
	return &EyesColorGene{
		T: "eyes_color_gene",

		Stats: stats,
	}
}

func (g *EyesColorGene) Type() string {
	return g.T
}

func (g *EyesColorGene) Produce(sex g.Sex) (gene.Byteble, error) {
	return tools.Search(
		AllEyesColors,
		func(e EyesColor) string { return e.Name },
		pm.GetRandomFromSeveral(g.Stats),
	), nil
}

func (g *EyesColorGene) Children() []gene.Gene {
	return nil
}

func (g *EyesColorGene) Bytes() []byte {
	return nil
}

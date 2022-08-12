package face

import (
	"encoding/json"

	"persons_generator/entities/human/color"
	"persons_generator/entities/human/gene"
	pm "persons_generator/probability_machine"
	"persons_generator/tools"
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
	t     string
	Stats map[string]float64
}

func NewEyesColorGene(stats map[string]float64) gene.Gene {
	return &EyesColorGene{
		t:     "eyes_color_gene",
		Stats: stats,
	}
}

func (g *EyesColorGene) Type() string {
	return g.t
}

func (g *EyesColorGene) Produce() (gene.Byteble, error) {
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

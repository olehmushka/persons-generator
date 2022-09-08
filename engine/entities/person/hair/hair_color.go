package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
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
	palette := tools.Search(
		color.AllHairColorPalettes,
		func(e string) string { return e },
		pm.GetRandomFromSeveral(g.Stats),
	)
	colors := color.GetHairColorsByPalette(palette)
	if len(colors) == 0 {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not get hair_colors by palette (palette=%s)", palette))
	}
	out, err := tools.RandomValueOfSlice(pm.RandFloat64, colors)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not get random hair_color")
	}

	return HairColor(out), nil
}

func (g *HairColorGene) Children() []gene.Gene {
	return nil
}

func (g *HairColorGene) Bytes() []byte {
	return nil
}

func (g *HairColorGene) Pair(in gene.Gene) (gene.Gene, error) {
	return in, nil
}

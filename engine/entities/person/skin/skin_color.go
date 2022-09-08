package skin

import (
	"encoding/json"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/gene"
	pm "persons_generator/engine/probability_machine"
)

type SkinColor color.Color

func (sc SkinColor) Zero() bool {
	return sc == SkinColor{}
}

func (sc SkinColor) Bytes() []byte {
	if sc.Zero() {
		return nil
	}

	if out, err := json.Marshal(sc); err == nil {
		return out
	}

	return nil
}

func (n SkinColor) Print() {
}

type SkinColorGene struct {
	T string `json:"t"`

	Stats     map[string]float64 `json:"stats"`
	MeanIndex int                `json:"mean_index"`
}

func NewSkinColorGene(stats map[string]float64, meanIndex int) gene.Gene {
	return &SkinColorGene{
		T: "skin_color_gene",

		Stats:     stats,
		MeanIndex: meanIndex,
	}
}

func (g *SkinColorGene) Type() string {
	return g.T
}

func (g *SkinColorGene) Produce(sex gender.Sex) (gene.Byteble, error) {
	palette := tools.Search(
		color.AllSkinColorPalettes,
		func(e string) string { return e },
		pm.GetRandomFromSeveral(g.Stats),
	)
	out, err := color.GetRandomColorByPaletteNorm(palette, g.MeanIndex)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not get random skin_color")
	}

	return SkinColor(out), nil
}

func (g *SkinColorGene) Children() []gene.Gene {
	return nil
}

func (g *SkinColorGene) Bytes() []byte {
	return nil
}

func (g *SkinColorGene) Pair(in gene.Gene) (gene.Gene, error) {
	return in, nil
}

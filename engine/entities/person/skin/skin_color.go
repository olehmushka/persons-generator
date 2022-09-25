package skin

import (
	"encoding/json"
	"fmt"
	"math"
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

func (sc SkinColor) Value() color.Color {
	return color.Color(sc)
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
	MeanIndex float64            `json:"mean_index"`
}

func NewSkinColorGene(stats map[string]float64, meanIndex float64) gene.Gene {
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
	p, err := pm.GetRandomFromSeveral(g.Stats)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate skin color")
	}

	palette := tools.Search(
		color.AllSkinColorPalettes,
		func(e string) string { return e },
		p,
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
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> skin_color genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}

	inStr, ok := in.(*SkinColorGene)
	if !ok {
		return nil, wrapped_error.NewInternalServerError(nil, "can not case input gene to SkinColorGene")
	}
	meanIndex := math.Floor((inStr.MeanIndex + g.MeanIndex) / 2)
	stats := tools.MergeMaps(g.Stats, inStr.Stats)

	return NewSkinColorGene(stats, meanIndex), nil
}

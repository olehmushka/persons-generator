package face

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

type EyesColor color.Color

func (n EyesColor) Zero() bool {
	return n == EyesColor{}
}

func (n EyesColor) Value() color.Color {
	return color.Color(n)
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

	Stats map[string]float64 `json:"stats"` // map[palette_name]probability
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
	cp, err := pm.GetRandomFromSeveral(g.Stats)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate eyes color")
	}
	palette := tools.Search(
		color.AllEyesColorPalettes,
		func(e string) string { return e },
		cp,
	)
	colors := color.GetEyesColorsByPalette(palette)
	if len(colors) == 0 {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not get colors by palette (palette=%s)", palette))
	}
	out, err := tools.RandomValueOfSlice(pm.RandFloat64, colors)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not get random color")
	}

	return EyesColor(out), nil
}

func (g *EyesColorGene) Children() []gene.Gene {
	return nil
}

func (g *EyesColorGene) Bytes() []byte {
	return nil
}

func (g *EyesColorGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> eyes_color genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}

	inStr, ok := in.(*EyesColorGene)
	if !ok {
		return nil, wrapped_error.NewInternalServerError(nil, "can not case input gene to EyesColorGene")
	}

	return NewEyesColorGene(tools.MergeMaps(g.Stats, inStr.Stats)), nil
}

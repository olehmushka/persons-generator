package size

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
	pm "persons_generator/engine/probability_machine"
)

type Height struct {
	HeightInCm float64 `json:"height_in_cm"`
}

func (h Height) Zero() bool {
	return h == Height{}
}

func (h Height) Bytes() []byte {
	if h.Zero() {
		return nil
	}
	if out, err := json.Marshal(h); err == nil {
		return out
	}

	return nil
}

type HeightGene struct {
	T string `json:"t"`

	MinHeight    float64 `json:"min_height"`
	MediumHeight float64 `json:"medium_height"`
	MaxHeight    float64 `json:"max_height"`
}

func NewHeightGene(min, medium, max float64) gene.Gene {
	minHeight, mediumHeight, maxHeight := normalizeHeightStats(min, medium, max)

	return &HeightGene{
		T: "height_gene",

		MinHeight:    minHeight,
		MediumHeight: mediumHeight,
		MaxHeight:    maxHeight,
	}
}

func (g *HeightGene) Type() string {
	return g.T
}

func (g *HeightGene) Produce(sex gender.Sex) (gene.Byteble, error) {
	h, err := pm.RandFloat64NormInRange(g.MinHeight, g.MaxHeight, HeightStandartDeviation, g.MediumHeight)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce height (min=%f, max=%f, standart_deviation=%f, medium(mean)=%f)", g.MinHeight, g.MaxHeight, HeightStandartDeviation, g.MediumHeight))
	}
	switch {
	case sex.IsMale():
		h = h * 1
	case sex.IsFemale():
		h = h * FemaleMaleRelation
	}

	return Height{
		HeightInCm: tools.RoundFloat64(h, 2),
	}, nil
}

func (g *HeightGene) Children() []gene.Gene {
	return nil
}

func (g *HeightGene) Bytes() []byte {
	return nil
}

func (g *HeightGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> height genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}

	inStr, ok := in.(*HeightGene)
	if !ok {
		return nil, wrapped_error.NewInternalServerError(nil, "can not case input gene to HeightGene")
	}
	var (
		min    = (inStr.MinHeight + g.MinHeight) / 2
		medium = (inStr.MediumHeight + g.MediumHeight) / 2
		max    = (inStr.MaxHeight + g.MaxHeight) / 2
	)
	min, medium, max = normalizeHeightStats(min, medium, max)

	return NewHeightGene(min, medium, max), nil
}

func normalizeHeightStats(min, medium, max float64) (float64, float64, float64) {
	var (
		outMin    = min
		outMedium = medium
		outMax    = max
	)
	if outMin < MinPossibleHeight {
		outMin = MinPossibleHeight
	}
	if outMax > MaxPossibleHeight {
		outMax = MaxPossibleHeight
	}
	if MinPossibleHeight < outMedium || outMedium > MaxPossibleHeight {
		outMedium = (outMin + outMax) / 2
	}

	return outMin, outMedium, outMax
}

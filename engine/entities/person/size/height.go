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
	if min < MinPossibleHeight {
		min = MinPossibleHeight
	}
	if max > MaxPossibleHeight {
		max = MaxPossibleHeight
	}
	if MinPossibleHeight < medium && medium > MaxPossibleHeight {
		medium = (min + max) / 2
	}

	return &HeightGene{
		T: "height_gene",

		MinHeight:    min,
		MediumHeight: medium,
		MaxHeight:    max,
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
	return []gene.Gene{}
}

func (g *HeightGene) Bytes() []byte {
	return nil
}

package size

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
	pm "persons_generator/engine/probability_machine"
)

type ShoeSize struct {
	USSize         int        `json:"us_size"`
	EUSize         int        `json:"eu_size"`
	UKSize         float64    `json:"uk_size"`
	FeetLengthInch float64    `json:"feet_length_inch"`
	FeetLengthCm   float64    `json:"feet_length_cm"`
	Sex            gender.Sex `json:"sex"`
}

func (ss ShoeSize) Zero() bool {
	return ss == ShoeSize{}
}

func (ss ShoeSize) Bytes() []byte {
	if ss.Zero() {
		return nil
	}
	if out, err := json.Marshal(ss); err == nil {
		return out
	}

	return nil
}

type ShoeSizeGene struct {
	T string `json:"t"`

	MinSizeInCm    float64 `json:"min_size_in_cm"`
	MediumSizeInCm float64 `json:"medium_size_in_cm"`
	MaxSizeInCm    float64 `json:"max_size_in_cm"`
}

func NewShoeSizeGene(min, medium, max float64) gene.Gene {
	return &ShoeSizeGene{
		T: "shoe_size_gene",

		MinSizeInCm:    min,
		MediumSizeInCm: medium,
		MaxSizeInCm:    max,
	}
}

func (g *ShoeSizeGene) Type() string {
	return g.T
}

func (g *ShoeSizeGene) Produce(sex gender.Sex) (gene.Byteble, error) {
	min, medium, max := prepareShoeSizeParams(sex, g.MinSizeInCm, g.MediumSizeInCm, g.MaxSizeInCm)
	ss, err := pm.RandFloat64NormInRange(min, max, HeightStandartDeviation, medium)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce shoe_size (min=%f, max=%f, standart_deviation=%f, medium(mean)=%f)", min, max, HeightStandartDeviation, medium))
	}

	return GetSizeByCm(sex, ss), nil
}

func (g *ShoeSizeGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *ShoeSizeGene) Bytes() []byte {
	return nil
}

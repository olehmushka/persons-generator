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
	minSizeInCm, mediumSizeInCm, maxSizeInCm := normalizeShoeSizeStats(min, medium, max)
	return &ShoeSizeGene{
		T: "shoe_size_gene",

		MinSizeInCm:    minSizeInCm,
		MediumSizeInCm: mediumSizeInCm,
		MaxSizeInCm:    maxSizeInCm,
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

func (g *ShoeSizeGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> shoe_size genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}

	inStr, ok := in.(*ShoeSizeGene)
	if !ok {
		return nil, wrapped_error.NewInternalServerError(nil, "can not case input gene to ShoeSizeGene")
	}
	var (
		min    = (inStr.MinSizeInCm + g.MinSizeInCm) / 2
		medium = (inStr.MediumSizeInCm + g.MediumSizeInCm) / 2
		max    = (inStr.MaxSizeInCm + g.MaxSizeInCm) / 2
	)
	min, medium, max = normalizeShoeSizeStats(min, medium, max)

	return NewShoeSizeGene(min, medium, max), nil
}

func normalizeShoeSizeStats(min, medium, max float64) (float64, float64, float64) {
	var (
		outMin    = min
		outMedium = medium
		outMax    = max
		absMin    = GetMinShoeSizeInCm(gender.FemaleSex)
		absMax    = GetMaxShoeSizeInCm(gender.MaleSex)
	)
	if outMin > absMin {
		outMin = absMin
	}
	if outMax > absMax {
		outMax = absMax
	}
	if absMin < outMedium || outMedium > absMax {
		outMedium = (outMin + outMax) / 2
	}

	return outMin, outMedium, outMax
}

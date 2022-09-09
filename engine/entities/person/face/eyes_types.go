package face

import (
	"fmt"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
	pm "persons_generator/engine/probability_machine"
)

type EyesType string

const (
	RoundEyesType              EyesType = "round_eyes"
	RoundishAlmondEyesType     EyesType = "roundish_almond_eyes"
	AlmondEyesType             EyesType = "almond_eyes"
	DroopyEyesType             EyesType = "droopy_eyes"
	DroopyHoodedEyesType       EyesType = "droopy_hooded_eyes"
	HoodedEyesType             EyesType = "hooded_eyes"
	AsianEyesType              EyesType = "asian_eyes"
	ChildishRoundAsianEyesType EyesType = "childish_round_asian_eyes"
)

var AllEyesTypes = []EyesType{
	RoundEyesType,
	RoundishAlmondEyesType,
	AlmondEyesType,
	DroopyEyesType,
	DroopyHoodedEyesType,
	HoodedEyesType,
	AsianEyesType,
	ChildishRoundAsianEyesType,
}

func (n EyesType) Zero() bool {
	return n == ""
}

func (n EyesType) Bytes() []byte {
	if n.Zero() {
		return []byte("")
	}
	return []byte(n)
}

type EyesTypeGene struct {
	T string `json:"t"`

	Stats map[string]float64 `json:"stats"`
}

func NewEyesTypeGene(stats map[string]float64) gene.Gene {
	return &EyesTypeGene{
		T: "eyes_type_gene",

		Stats: stats,
	}
}

func (g *EyesTypeGene) Type() string {
	return g.T
}

func (g *EyesTypeGene) Produce(sex g.Sex) (gene.Byteble, error) {
	return EyesType(pm.GetRandomFromSeveral(g.Stats)), nil
}

func (g *EyesTypeGene) Children() []gene.Gene {
	return nil
}

func (g *EyesTypeGene) Bytes() []byte {
	return nil
}

func (g *EyesTypeGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> eyes_type genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}

	inStr, ok := in.(*EyesTypeGene)
	if !ok {
		return nil, wrapped_error.NewInternalServerError(nil, "can not case input gene to EyesTypeGene")
	}

	return NewEyesTypeGene(tools.MergeMaps(g.Stats, inStr.Stats)), nil
}

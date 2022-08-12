package face

import (
	"persons_generator/entities/human/gene"
	pm "persons_generator/probability_machine"
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
	t     string
	Stats map[string]float64
}

func NewEyesTypeGene(stats map[string]float64) gene.Gene {
	return &EyesTypeGene{
		t:     "eyes_type_gene",
		Stats: stats,
	}
}

func (g *EyesTypeGene) Type() string {
	return g.t
}

func (g *EyesTypeGene) Produce() (gene.Byteble, error) {
	return EyesType(pm.GetRandomFromSeveral(g.Stats)), nil
}

func (g *EyesTypeGene) Children() []gene.Gene {
	return nil
}

func (g *EyesTypeGene) Bytes() []byte {
	return nil
}

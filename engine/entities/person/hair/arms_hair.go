package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type ArmsHair struct {
	Density *HairDensity `json:"density"`
}

func (ah ArmsHair) Zero() bool {
	return ah == ArmsHair{}
}

func (ah ArmsHair) Bytes() []byte {
	if ah.Zero() {
		return nil
	}

	if out, err := json.Marshal(ah); err == nil {
		return out
	}

	return nil
}

type ArmsHairGene struct {
	T string `json:"t"`

	HairDensityGene gene.Gene `json:"hair_density_gene"`
}

func NewArmsHairGene(
	armsHairDensityGene gene.Gene,
) gene.Gene {
	return &ArmsHairGene{
		T: "arms_hair_gene",

		HairDensityGene: armsHairDensityGene,
	}
}

func (g *ArmsHairGene) Type() string {
	return g.T
}

func (g *ArmsHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	density, err := g.HairDensityGene.Produce(sex)
	if err != nil {
		return ArmsHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce arms_hair_density from gene by sex (sex=%s)", sex))
	}
	var densityVal HairDensity
	if err := json.Unmarshal(density.Bytes(), &densityVal); err != nil {
		return ArmsHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal arms_hair_density produced from gene")
	}

	return ArmsHair{
		Density: &densityVal,
	}, nil
}

func (g *ArmsHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairDensityGene,
	}
}

func (g *ArmsHairGene) Bytes() []byte {
	return nil
}

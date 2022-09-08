package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type AxillaryHair struct {
	Density *HairDensity `json:"density"`
}

func (ah AxillaryHair) Zero() bool {
	return ah == AxillaryHair{}
}

func (ah AxillaryHair) Bytes() []byte {
	if ah.Zero() {
		return nil
	}

	if out, err := json.Marshal(ah); err == nil {
		return out
	}

	return nil
}

type AxillaryHairGene struct {
	T string `json:"t"`

	HairDensityGene gene.Gene `json:"hair_density_gene"`
}

func NewAxillaryHairGene(
	axillaryHairDensityGene gene.Gene,
) gene.Gene {
	return &AxillaryHairGene{
		T: "axillary_hair_gene",

		HairDensityGene: axillaryHairDensityGene,
	}
}

func (g *AxillaryHairGene) Type() string {
	return g.T
}

func (g *AxillaryHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	density, err := g.HairDensityGene.Produce(sex)
	if err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce axillary_hair_density from gene by sex (sex=%s)", sex))
	}
	var densityVal HairDensity
	if err := json.Unmarshal(density.Bytes(), &densityVal); err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal axillary_hair_density produced from gene")
	}

	return AxillaryHair{
		Density: &densityVal,
	}, nil
}

func (g *AxillaryHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairDensityGene,
	}
}

func (g *AxillaryHairGene) Bytes() []byte {
	return nil
}

func (g *AxillaryHairGene) Pair(in gene.Gene) (gene.Gene, error) {
	return in, nil
}

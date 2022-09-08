package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type LegsHair struct {
	Density *HairDensity `json:"density"`
}

func (lh LegsHair) Zero() bool {
	return lh == LegsHair{}
}

func (lh LegsHair) Bytes() []byte {
	if lh.Zero() {
		return nil
	}

	if out, err := json.Marshal(lh); err == nil {
		return out
	}

	return nil
}

type LegsHairGene struct {
	T string `json:"t"`

	HairDensityGene gene.Gene `json:"hair_density_gene"`
}

func NewLegsHairGene(
	legsHairDensityGene gene.Gene,
) gene.Gene {
	return &LegsHairGene{
		T: "legs_hair_gene",

		HairDensityGene: legsHairDensityGene,
	}
}

func (g *LegsHairGene) Type() string {
	return g.T
}

func (g *LegsHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	density, err := g.HairDensityGene.Produce(sex)
	if err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce legs_hair_density from gene by sex (sex=%s)", sex))
	}
	var densityVal HairDensity
	if err := json.Unmarshal(density.Bytes(), &densityVal); err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal legs_hair_density produced from gene")
	}

	return LegsHair{
		Density: &densityVal,
	}, nil
}

func (g *LegsHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairDensityGene,
	}
}

func (g *LegsHairGene) Bytes() []byte {
	return nil
}

func (g *LegsHairGene) Pair(in gene.Gene) (gene.Gene, error) {
	return in, nil
}

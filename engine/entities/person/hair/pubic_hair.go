package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type PubicHair struct {
	Density *HairDensity `json:"density"`
}

func (ph PubicHair) Zero() bool {
	return ph == PubicHair{}
}

func (ph PubicHair) Bytes() []byte {
	if ph.Zero() {
		return nil
	}

	if out, err := json.Marshal(ph); err == nil {
		return out
	}

	return nil
}

type PubicHairGene struct {
	T string `json:"t"`

	HairDensityGene gene.Gene `json:"hair_density_gene"`
}

func NewPubicHairGene(
	pubicHairDensityGene gene.Gene,
) gene.Gene {
	return &PubicHairGene{
		T: "pubic_hair_gene",

		HairDensityGene: pubicHairDensityGene,
	}
}

func (g *PubicHairGene) Type() string {
	return g.T
}

func (g *PubicHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	density, err := g.HairDensityGene.Produce(sex)
	if err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce pubic_hair_density from gene by sex (sex=%s)", sex))
	}
	var densityVal HairDensity
	if err := json.Unmarshal(density.Bytes(), &densityVal); err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal pubic_hair_density produced from gene")
	}

	return PubicHair{
		Density: &densityVal,
	}, nil
}

func (g *PubicHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairDensityGene,
	}
}

func (g *PubicHairGene) Bytes() []byte {
	return nil
}

func (g *PubicHairGene) Pair(in gene.Gene) (gene.Gene, error) {
	return in, nil
}

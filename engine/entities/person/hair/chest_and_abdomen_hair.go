package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type ChestAndAbdomenHair struct {
	Density *HairDensity `json:"density"`
}

func (h ChestAndAbdomenHair) Zero() bool {
	return h == ChestAndAbdomenHair{}
}

func (h ChestAndAbdomenHair) Bytes() []byte {
	if h.Zero() {
		return nil
	}

	if out, err := json.Marshal(h); err == nil {
		return out
	}

	return nil
}

type ChestAndAbdomenHairGene struct {
	T string `json:"t"`

	HairDensityGene gene.Gene `json:"hair_density_gene"`
}

func NewChestAndAbdomenHairGene(
	chestAndAbdomenHairDensityGene gene.Gene,
) gene.Gene {
	return &ChestAndAbdomenHairGene{
		T: "chest_and_abdomen_hair_gene",

		HairDensityGene: chestAndAbdomenHairDensityGene,
	}
}

func (g *ChestAndAbdomenHairGene) Type() string {
	return g.T
}

func (g *ChestAndAbdomenHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	density, err := g.HairDensityGene.Produce(sex)
	if err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce chest_and_abdomen_hair_density from gene by sex (sex=%s)", sex))
	}
	var densityVal HairDensity
	if err := json.Unmarshal(density.Bytes(), &densityVal); err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal chest_and_abdomen_hair_density produced from gene")
	}

	return ChestAndAbdomenHair{
		Density: &densityVal,
	}, nil
}

func (g *ChestAndAbdomenHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairDensityGene,
	}
}

func (g *ChestAndAbdomenHairGene) Bytes() []byte {
	return nil
}

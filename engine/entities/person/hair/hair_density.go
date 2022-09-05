package hair

import (
	"encoding/json"
	"persons_generator/core/tools"
	"persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
	pm "persons_generator/engine/probability_machine"
)

type HairDensity struct {
	Type string `json:"type"`
}

func (hd HairDensity) Zero() bool {
	return hd == HairDensity{}
}

func (hd HairDensity) Bytes() []byte {
	if hd.Zero() {
		return nil
	}

	if out, err := json.Marshal(hd); err == nil {
		return out
	}

	return nil
}

func (hd HairDensity) Print() {
}

type HairDensityGene struct {
	T string `json:"t"`

	Stats map[string]float64 `json:"stats"`
}

func NewHairDensityGene(stats map[string]float64) gene.Gene {
	return &HairDensityGene{
		T: "hair_density_gene",

		Stats: stats,
	}
}

func (g *HairDensityGene) Type() string {
	return g.T
}

func (g *HairDensityGene) Produce(sex gender.Sex) (gene.Byteble, error) {
	return tools.Search(
		AllHairDensities,
		func(e HairDensity) string { return e.Type },
		pm.GetRandomFromSeveral(g.Stats),
	), nil
}

func (g *HairDensityGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *HairDensityGene) Bytes() []byte {
	return nil
}

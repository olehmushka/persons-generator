package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
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
	hd, err := pm.GetRandomFromSeveral(g.Stats)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate hair density")
	}

	return tools.Search(
		AllHairDensities,
		func(e HairDensity) string { return e.Type },
		hd,
	), nil
}

func (g *HairDensityGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *HairDensityGene) Bytes() []byte {
	return nil
}

func (g *HairDensityGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> hair_density genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}

	inStr, ok := in.(*HairDensityGene)
	if !ok {
		return nil, wrapped_error.NewInternalServerError(nil, "can not case input gene to HairDensityGene")
	}

	return NewHairDensityGene(tools.MergeMaps(g.Stats, inStr.Stats)), nil
}

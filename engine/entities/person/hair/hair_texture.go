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

type HairTexture struct {
	Type string `json:"type"`
}

func (ht HairTexture) Zero() bool {
	return ht == HairTexture{}
}

func (ht HairTexture) Bytes() []byte {
	if ht.Zero() {
		return nil
	}

	if out, err := json.Marshal(ht); err == nil {
		return out
	}

	return nil
}

func (ht HairTexture) Print() {
}

type HairTextureGene struct {
	T string `json:"t"`

	Stats map[string]float64 `json:"stats"`
}

func NewHairTextureGene(stats map[string]float64) gene.Gene {
	return &HairTextureGene{
		T: "hair_texture_gene",

		Stats: stats,
	}
}

func (g *HairTextureGene) Type() string {
	return g.T
}

func (g *HairTextureGene) Produce(sex gender.Sex) (gene.Byteble, error) {
	ht, err := pm.GetRandomFromSeveral(g.Stats)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate hair texture")
	}

	return tools.Search(
		AllHairTextures,
		func(e HairTexture) string { return e.Type },
		ht,
	), nil
}

func (g *HairTextureGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *HairTextureGene) Bytes() []byte {
	return nil
}

func (g *HairTextureGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> hair_texture genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}

	inStr, ok := in.(*HairTextureGene)
	if !ok {
		return nil, wrapped_error.NewInternalServerError(nil, "can not case input gene to HairTextureGene")
	}

	return NewHairTextureGene(tools.MergeMaps(g.Stats, inStr.Stats)), nil
}

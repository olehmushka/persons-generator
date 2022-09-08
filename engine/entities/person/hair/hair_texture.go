package hair

import (
	"encoding/json"
	"persons_generator/core/tools"
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
	return tools.Search(
		AllHairTextures,
		func(e HairTexture) string { return e.Type },
		pm.GetRandomFromSeveral(g.Stats),
	), nil
}

func (g *HairTextureGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *HairTextureGene) Bytes() []byte {
	return nil
}

func (g *HairTextureGene) Pair(in gene.Gene) (gene.Gene, error) {
	return in, nil
}

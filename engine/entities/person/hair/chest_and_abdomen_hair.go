package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type ChestAndAbdomenHair struct {
	Color   *HairColor   `json:"color"`
	Texture *HairTexture `json:"texture"`
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

	HairColorGene   gene.Gene `json:"hair_color_gene"`
	HairTextureGene gene.Gene `json:"hair_texture_gene"`
}

func NewChestAndAbdomenHairGene(
	chestAndAbdomenHairColorGene gene.Gene,
	chestAndAbdomenHairTextureGene gene.Gene,
) gene.Gene {
	return &ChestAndAbdomenHairGene{
		T: "chest_and_abdomen_hair_gene",

		HairColorGene:   chestAndAbdomenHairColorGene,
		HairTextureGene: chestAndAbdomenHairTextureGene,
	}
}

func (g *ChestAndAbdomenHairGene) Type() string {
	return g.T
}

func (g *ChestAndAbdomenHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	color, err := g.HairColorGene.Produce(sex)
	if err != nil {
		return ChestAndAbdomenHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce chest_and_abdomen_hair_color from gene by sex (sex=%s)", sex))
	}
	var colorVal HairColor
	if err := json.Unmarshal(color.Bytes(), &colorVal); err != nil {
		return ChestAndAbdomenHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal chest_and_abdomen_hair_color produced from gene")
	}

	texture, err := g.HairTextureGene.Produce(sex)
	if err != nil {
		return ChestAndAbdomenHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce chest_and_abdomen_hair_texture from gene by sex (sex=%s)", sex))
	}
	var textureVal HairTexture
	if err := json.Unmarshal(texture.Bytes(), &textureVal); err != nil {
		return ChestAndAbdomenHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal chest_and_abdomen_hair_texture produced from gene")
	}

	return ChestAndAbdomenHair{
		Texture: &textureVal,
		Color:   &colorVal,
	}, nil
}

func (g *ChestAndAbdomenHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairColorGene,
		g.HairTextureGene,
	}
}

func (g *ChestAndAbdomenHairGene) Bytes() []byte {
	return nil
}

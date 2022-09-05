package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type ArmsHair struct {
	Color   *HairColor   `json:"color"`
	Texture *HairTexture `json:"texture"`
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

	HairColorGene   gene.Gene `json:"hair_color_gene"`
	HairTextureGene gene.Gene `json:"hair_texture_gene"`
}

func NewArmsHairGene(
	armsHairColorGene gene.Gene,
	armsHairTextureGene gene.Gene,
) gene.Gene {
	return &ArmsHairGene{
		T: "arms_hair_gene",

		HairColorGene:   armsHairColorGene,
		HairTextureGene: armsHairTextureGene,
	}
}

func (g *ArmsHairGene) Type() string {
	return g.T
}

func (g *ArmsHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	color, err := g.HairColorGene.Produce(sex)
	if err != nil {
		return ArmsHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce arms_hair_color from gene by sex (sex=%s)", sex))
	}
	var colorVal HairColor
	if err := json.Unmarshal(color.Bytes(), &colorVal); err != nil {
		return ArmsHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal arms_hair_color produced from gene")
	}

	texture, err := g.HairTextureGene.Produce(sex)
	if err != nil {
		return ArmsHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce arms_hair_texture from gene by sex (sex=%s)", sex))
	}
	var textureVal HairTexture
	if err := json.Unmarshal(texture.Bytes(), &textureVal); err != nil {
		return ArmsHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal arms_hair_texture produced from gene")
	}

	return ArmsHair{
		Texture: &textureVal,
		Color:   &colorVal,
	}, nil
}

func (g *ArmsHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairColorGene,
		g.HairTextureGene,
	}
}

func (g *ArmsHairGene) Bytes() []byte {
	return nil
}

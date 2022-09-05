package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type LegsHair struct {
	Color   *HairColor   `json:"color"`
	Texture *HairTexture `json:"texture"`
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

	HairColorGene   gene.Gene `json:"hair_color_gene"`
	HairTextureGene gene.Gene `json:"hair_texture_gene"`
}

func NewLegsHairGene(
	legsHairColorGene gene.Gene,
	legsHairTextureGene gene.Gene,
) gene.Gene {
	return &LegsHairGene{
		T: "legs_hair_gene",

		HairColorGene:   legsHairColorGene,
		HairTextureGene: legsHairTextureGene,
	}
}

func (g *LegsHairGene) Type() string {
	return g.T
}

func (g *LegsHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	color, err := g.HairColorGene.Produce(sex)
	if err != nil {
		return LegsHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce legs_hair_color from gene by sex (sex=%s)", sex))
	}
	var colorVal HairColor
	if err := json.Unmarshal(color.Bytes(), &colorVal); err != nil {
		return LegsHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal legs_hair_color produced from gene")
	}

	texture, err := g.HairTextureGene.Produce(sex)
	if err != nil {
		return LegsHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce legs_hair_color from gene by sex (sex=%s)", sex))
	}
	var textureVal HairTexture
	if err := json.Unmarshal(texture.Bytes(), &textureVal); err != nil {
		return LegsHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal legs_hair_color produced from gene")
	}

	return LegsHair{
		Texture: &textureVal,
		Color:   &colorVal,
	}, nil
}

func (g *LegsHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairColorGene,
		g.HairTextureGene,
	}
}

func (g *LegsHairGene) Bytes() []byte {
	return nil
}

package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type FeetHair struct {
	Color   *HairColor   `json:"color"`
	Texture *HairTexture `json:"texture"`
}

func (fh FeetHair) Zero() bool {
	return fh == FeetHair{}
}

func (fh FeetHair) Bytes() []byte {
	if fh.Zero() {
		return nil
	}

	if out, err := json.Marshal(fh); err == nil {
		return out
	}

	return nil
}

type FeetHairGene struct {
	T string `json:"t"`

	HairColorGene   gene.Gene `json:"hair_color_gene"`
	HairTextureGene gene.Gene `json:"hair_texture_gene"`
}

func NewFeetHairGene(
	feetHairColorGene gene.Gene,
	feetHairTextureGene gene.Gene,
) gene.Gene {
	return &FeetHairGene{
		T: "feet_hair_gene",

		HairColorGene:   feetHairColorGene,
		HairTextureGene: feetHairTextureGene,
	}
}

func (g *FeetHairGene) Type() string {
	return g.T
}

func (g *FeetHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	color, err := g.HairColorGene.Produce(sex)
	if err != nil {
		return ScaplHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce feet_hair_color from gene by sex (sex=%s)", sex))
	}
	var colorVal HairColor
	if err := json.Unmarshal(color.Bytes(), &colorVal); err != nil {
		return FeetHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal feet_hair_color produced from gene")
	}

	texture, err := g.HairTextureGene.Produce(sex)
	if err != nil {
		return FeetHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce feet_hair_color from gene by sex (sex=%s)", sex))
	}
	var textureVal HairTexture
	if err := json.Unmarshal(texture.Bytes(), &textureVal); err != nil {
		return FeetHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal feet_hair_color produced from gene")
	}

	return FeetHair{
		Texture: &textureVal,
		Color:   &colorVal,
	}, nil
}

func (g *FeetHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairColorGene,
		g.HairTextureGene,
	}
}

func (g *FeetHairGene) Bytes() []byte {
	return nil
}

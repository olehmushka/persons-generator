package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type PubicHair struct {
	Color   *HairColor   `json:"color"`
	Texture *HairTexture `json:"texture"`
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

	HairColorGene   gene.Gene `json:"hair_color_gene"`
	HairTextureGene gene.Gene `json:"hair_texture_gene"`
}

func NewPubicHairGene(
	pubicHairColorGene gene.Gene,
	pubicHairTextureGene gene.Gene,
) gene.Gene {
	return &PubicHairGene{
		T: "pubic_hair_gene",

		HairColorGene:   pubicHairColorGene,
		HairTextureGene: pubicHairTextureGene,
	}
}

func (g *PubicHairGene) Type() string {
	return g.T
}

func (g *PubicHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	color, err := g.HairColorGene.Produce(sex)
	if err != nil {
		return PubicHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce pubic_hair_color from gene by sex (sex=%s)", sex))
	}
	var colorVal HairColor
	if err := json.Unmarshal(color.Bytes(), &colorVal); err != nil {
		return PubicHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal pubic_hair_color produced from gene")
	}

	texture, err := g.HairTextureGene.Produce(sex)
	if err != nil {
		return PubicHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce pubic_hair_color from gene by sex (sex=%s)", sex))
	}
	var textureVal HairTexture
	if err := json.Unmarshal(texture.Bytes(), &textureVal); err != nil {
		return PubicHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal pubic_hair_color produced from gene")
	}

	return PubicHair{
		Texture: &textureVal,
		Color:   &colorVal,
	}, nil
}

func (g *PubicHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairColorGene,
		g.HairTextureGene,
	}
}

func (g *PubicHairGene) Bytes() []byte {
	return nil
}

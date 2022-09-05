package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type ScaplHair struct {
	Color   *HairColor   `json:"color"`
	Texture *HairTexture `json:"texture"`
}

func (sh ScaplHair) Zero() bool {
	return sh == ScaplHair{}
}

func (sh ScaplHair) Bytes() []byte {
	if sh.Zero() {
		return nil
	}

	if out, err := json.Marshal(sh); err == nil {
		return out
	}

	return nil
}

type ScaplHairGene struct {
	T string `json:"t"`

	HairColorGene   gene.Gene `json:"hair_color_gene"`
	HairTextureGene gene.Gene `json:"hair_texture_gene"`
}

func NewScalpHairGene(
	scalpHairColorGene gene.Gene,
	scalpHairTextureGene gene.Gene,
) gene.Gene {
	return &ScaplHairGene{
		T:               "scapl_hair_gene",
		HairColorGene:   scalpHairColorGene,
		HairTextureGene: scalpHairTextureGene,
	}
}

func (g *ScaplHairGene) Type() string {
	return g.T
}

func (g *ScaplHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	color, err := g.HairColorGene.Produce(sex)
	if err != nil {
		return ScaplHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce scapl_hair_color from gene by sex (sex=%s)", sex))
	}
	var colorVal HairColor
	if err := json.Unmarshal(color.Bytes(), &colorVal); err != nil {
		return ScaplHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal scapl_hair_color produced from gene")
	}

	texture, err := g.HairTextureGene.Produce(sex)
	if err != nil {
		return ScaplHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce head_hair_texture from gene by sex (sex=%s)", sex))
	}
	var textureVal HairTexture
	if err := json.Unmarshal(texture.Bytes(), &textureVal); err != nil {
		return ScaplHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal head_hair_texture produced from gene")
	}

	return ScaplHair{
		Texture: &textureVal,
		Color:   &colorVal,
	}, nil
}

func (g *ScaplHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairColorGene,
		g.HairTextureGene,
	}
}

func (g *ScaplHairGene) Bytes() []byte {
	return nil
}

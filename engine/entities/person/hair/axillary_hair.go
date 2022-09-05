package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type AxillaryHair struct {
	Color   *HairColor   `json:"color"`
	Texture *HairTexture `json:"texture"`
}

func (ah AxillaryHair) Zero() bool {
	return ah == AxillaryHair{}
}

func (ah AxillaryHair) Bytes() []byte {
	if ah.Zero() {
		return nil
	}

	if out, err := json.Marshal(ah); err == nil {
		return out
	}

	return nil
}

type AxillaryHairGene struct {
	T string `json:"t"`

	HairColorGene   gene.Gene `json:"hair_color_gene"`
	HairTextureGene gene.Gene `json:"hair_texture_gene"`
}

func NewAxillaryHairGene(
	axillaryHairColorGene gene.Gene,
	axillaryHairTextureGene gene.Gene,
) gene.Gene {
	return &AxillaryHairGene{
		T: "axillary_hair_gene",

		HairColorGene:   axillaryHairColorGene,
		HairTextureGene: axillaryHairTextureGene,
	}
}

func (g *AxillaryHairGene) Type() string {
	return g.T
}

func (g *AxillaryHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	color, err := g.HairColorGene.Produce(sex)
	if err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce axillary_hair_color from gene by sex (sex=%s)", sex))
	}
	var colorVal HairColor
	if err := json.Unmarshal(color.Bytes(), &colorVal); err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal axillary_hair_color produced from gene")
	}

	texture, err := g.HairTextureGene.Produce(sex)
	if err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce axillary_hair_texture from gene by sex (sex=%s)", sex))
	}
	var textureVal HairTexture
	if err := json.Unmarshal(texture.Bytes(), &textureVal); err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal axillary_hair_texture produced from gene")
	}

	return AxillaryHair{
		Texture: &textureVal,
		Color:   &colorVal,
	}, nil
}

func (g *AxillaryHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairColorGene,
		g.HairTextureGene,
	}
}

func (g *AxillaryHairGene) Bytes() []byte {
	return nil
}

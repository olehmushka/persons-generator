package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type FaceHair struct {
	Color   *HairColor   `json:"color"`
	Texture *HairTexture `json:"texture"`
}

func (fh FaceHair) Zero() bool {
	return fh == FaceHair{}
}

func (fh FaceHair) Bytes() []byte {
	if fh.Zero() {
		return nil
	}

	if out, err := json.Marshal(fh); err == nil {
		return out
	}

	return nil
}

type FaceHairGene struct {
	T string `json:"t"`

	HairColorGene   gene.Gene `json:"hair_color_gene"`
	HairTextureGene gene.Gene `json:"hair_texture_gene"`
}

func NewFaceHairGene(
	feetHairColorGene gene.Gene,
	feetHairTextureGene gene.Gene,
) gene.Gene {
	return &FeetHairGene{
		T: "face_hair_gene",

		HairColorGene:   feetHairColorGene,
		HairTextureGene: feetHairTextureGene,
	}
}

func (g *FaceHairGene) Type() string {
	return g.T
}

func (g *FaceHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	color, err := g.HairColorGene.Produce(sex)
	if err != nil {
		return FaceHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce face_hair_color from gene by sex (sex=%s)", sex))
	}
	var colorVal HairColor
	if err := json.Unmarshal(color.Bytes(), &colorVal); err != nil {
		return FaceHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal face_hair_color produced from gene")
	}

	texture, err := g.HairTextureGene.Produce(sex)
	if err != nil {
		return FaceHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce face_hair_color from gene by sex (sex=%s)", sex))
	}
	var textureVal HairTexture
	if err := json.Unmarshal(texture.Bytes(), &textureVal); err != nil {
		return FaceHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal face_hair_color produced from gene")
	}

	return FaceHair{
		Texture: &textureVal,
		Color:   &colorVal,
	}, nil
}

func (g *FaceHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairColorGene,
		g.HairTextureGene,
	}
}

func (g *FaceHairGene) Bytes() []byte {
	return nil
}

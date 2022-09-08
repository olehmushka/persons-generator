package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type Hair struct {
	Scapl           *ScaplHair           `json:"scapl"`
	Face            *FaceHair            `json:"face"`
	Axillary        *AxillaryHair        `json:"axillary"`
	ChestAndAbdomen *ChestAndAbdomenHair `json:"chest_and_abdomen"`
	Arms            *ArmsHair            `json:"arms"`
	Pubic           *PubicHair           `json:"pubic"`
	Legs            *LegsHair            `json:"legs"`
	Feet            *FeetHair            `json:"feet"`
	Color           *HairColor           `json:"color"`
}

func (h Hair) Zero() bool {
	return h == Hair{}
}

func (h Hair) Bytes() []byte {
	if h.Zero() {
		return nil
	}
	if out, err := json.Marshal(h); err == nil {
		return out
	}

	return nil
}

type HairGene struct {
	T string `json:"t"`

	ScaplHairGene           gene.Gene `json:"scalp_hair_gene"`
	FaceHairGene            gene.Gene `json:"face_hair_gene"`
	AxillaryHairGene        gene.Gene `json:"axillary_hair_gene"`
	ChestAndAbdomenHairGene gene.Gene `json:"chest_and_abdomen_hair_gene"`
	ArmsHairGene            gene.Gene `json:"arms_hair_gene"`
	PubicHairGene           gene.Gene `json:"pubic_hair_gene"`
	LegsHairGene            gene.Gene `json:"legs_hair_gene"`
	FeetHairGene            gene.Gene `json:"feet_hair_gene"`
	HairColorGene           gene.Gene `json:"hair_color_gene"`
}

func NewHairGene(
	scaplHairGene gene.Gene,
	faceHairGene gene.Gene,
	axillaryHairGene gene.Gene,
	chestAndAbdomenHairGene gene.Gene,
	armsHairGene gene.Gene,
	pubicHairGene gene.Gene,
	legsHairGene gene.Gene,
	feetHairGene gene.Gene,
	hairColorGene gene.Gene,
) gene.Gene {
	return &HairGene{
		T: "hair_gene",

		ScaplHairGene:           scaplHairGene,
		FaceHairGene:            faceHairGene,
		AxillaryHairGene:        axillaryHairGene,
		ChestAndAbdomenHairGene: chestAndAbdomenHairGene,
		ArmsHairGene:            armsHairGene,
		PubicHairGene:           pubicHairGene,
		LegsHairGene:            legsHairGene,
		FeetHairGene:            feetHairGene,
		HairColorGene:           hairColorGene,
	}
}

func (g *HairGene) Type() string {
	return g.T
}

func (g *HairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	scapl, err := g.ScaplHairGene.Produce(sex)
	if err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce scapl_hair from gene by sex (sex=%s)", sex))
	}

	var scaplVal ScaplHair
	if err := json.Unmarshal(scapl.Bytes(), &scaplVal); err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal scapl_hair produced from gene")
	}

	face, err := g.FaceHairGene.Produce(sex)
	if err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce face_hair from gene by sex (sex=%s)", sex))
	}

	var faceVal FaceHair
	if err := json.Unmarshal(face.Bytes(), &faceVal); err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal face_hair produced from gene")
	}

	axillary, err := g.AxillaryHairGene.Produce(sex)
	if err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce axillary_hair from gene by sex (sex=%s)", sex))
	}

	var axillaryVal AxillaryHair
	if err := json.Unmarshal(axillary.Bytes(), &axillaryVal); err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal axillary_hair produced from gene")
	}

	chestAndAbdomen, err := g.ChestAndAbdomenHairGene.Produce(sex)
	if err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce chest_and_abdomen_hair from gene by sex (sex=%s)", sex))
	}

	var chestAndAbdomenVal ChestAndAbdomenHair
	if err := json.Unmarshal(chestAndAbdomen.Bytes(), &chestAndAbdomenVal); err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal chest_and_abdomen_hair produced from gene")
	}

	arms, err := g.ArmsHairGene.Produce(sex)
	if err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce arms_hair from gene by sex (sex=%s)", sex))
	}

	var armsVal ArmsHair
	if err := json.Unmarshal(arms.Bytes(), &armsVal); err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal arms_hair produced from gene")
	}

	pubic, err := g.PubicHairGene.Produce(sex)
	if err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce pubic_hair from gene by sex (sex=%s)", sex))
	}

	var pubicVal PubicHair
	if err := json.Unmarshal(pubic.Bytes(), &pubicVal); err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal pubic_hair produced from gene")
	}

	legs, err := g.LegsHairGene.Produce(sex)
	if err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce legs_hair from gene by sex (sex=%s)", sex))
	}

	var legsVal LegsHair
	if err := json.Unmarshal(legs.Bytes(), &legsVal); err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal legs_hair produced from gene")
	}

	feet, err := g.FeetHairGene.Produce(sex)
	if err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce feet_hair from gene by sex (sex=%s)", sex))
	}

	var feetVal FeetHair
	if err := json.Unmarshal(feet.Bytes(), &feetVal); err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal feet_hair produced from gene")
	}

	color, err := g.HairColorGene.Produce(sex)
	if err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce hair_color from gene by sex (sex=%s)", sex))
	}
	var colorVal HairColor
	if err := json.Unmarshal(color.Bytes(), &colorVal); err != nil {
		return Hair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal hair_color produced from gene")
	}

	return Hair{
		Scapl:           &scaplVal,
		Face:            &faceVal,
		Axillary:        &axillaryVal,
		ChestAndAbdomen: &chestAndAbdomenVal,
		Arms:            &armsVal,
		Pubic:           &pubicVal,
		Legs:            &legsVal,
		Feet:            &feetVal,
		Color:           &colorVal,
	}, nil
}

func (g *HairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.ScaplHairGene,
		g.FaceHairGene,
		g.AxillaryHairGene,
		g.ChestAndAbdomenHairGene,
		g.ArmsHairGene,
		g.PubicHairGene,
		g.LegsHairGene,
		g.FeetHairGene,
		g.HairColorGene,
	}
}

func (g *HairGene) Bytes() []byte {
	return nil
}

func (g *HairGene) Pair(in gene.Gene) (gene.Gene, error) {
	return in, nil
}

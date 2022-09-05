package body

import (
	"encoding/json"
	"fmt"

	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/face"
	"persons_generator/engine/entities/person/gene"
	"persons_generator/engine/entities/person/hair"
)

type BodyGene struct {
	T string `json:"t"`

	FaceGene gene.Gene `json:"face_gene"`
	HairGene gene.Gene `json:"hair_gene"`
}

func NewBodyGene(
	faceGene gene.Gene,
	hairGene gene.Gene,
) gene.Gene {
	return &BodyGene{
		T: "body_gene",

		FaceGene: faceGene,
		HairGene: hairGene,
	}
}

func (g *BodyGene) Type() string {
	return g.T
}

func (g *BodyGene) Produce(sex g.Sex) (gene.Byteble, error) {
	f, err := g.FaceGene.Produce(sex)
	if err != nil {
		return Body{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce face from gene by sex (sex=%s)", sex))
	}
	var faceVal face.Face
	if err := json.Unmarshal(f.Bytes(), &faceVal); err != nil {
		return Body{}, wrapped_error.NewInternalServerError(err, "can not unmarshal face produced from gene")
	}

	h, err := g.HairGene.Produce(sex)
	if err != nil {
		return Body{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce head_hair from gene by sex (sex=%s)", sex))
	}
	var hairVal hair.Hair
	if err := json.Unmarshal(h.Bytes(), &hairVal); err != nil {
		return Body{}, wrapped_error.NewInternalServerError(err, "can not unmarshal head_hair produced from gene")
	}

	return Body{
		Face: faceVal,
		Hair: hairVal,
	}, nil
}

func (g *BodyGene) Children() []gene.Gene {
	return []gene.Gene{
		g.FaceGene,
		g.HairGene,
	}
}

func (g *BodyGene) Bytes() []byte {
	return nil
}

package body

import (
	"encoding/json"

	"persons_generator/engine/entities/human/face"
	"persons_generator/engine/entities/human/gene"
)

type BodyGene struct {
	t string

	FaceGene gene.Gene
}

func NewBodyGene(faceGene gene.Gene) gene.Gene {
	return &BodyGene{
		t:        "body_gene",
		FaceGene: faceGene,
	}
}

func (g *BodyGene) Type() string {
	return g.t
}

func (g *BodyGene) Produce() (gene.Byteble, error) {
	f, err := g.FaceGene.Produce()
	if err != nil {
		return Body{}, err
	}
	var faceVal face.Face
	if err := json.Unmarshal(f.Bytes(), &faceVal); err != nil {
		return Body{}, err
	}

	return Body{
		Face: faceVal,
	}, nil
}

func (g *BodyGene) Children() []gene.Gene {
	return []gene.Gene{
		g.FaceGene,
	}
}

func (g *BodyGene) Bytes() []byte {
	return nil
}

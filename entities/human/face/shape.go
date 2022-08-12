package face

import (
	"persons_generator/entities/human/gene"
	pm "persons_generator/probability_machine"
)

type Shape string

const (
	Oval   Shape = "oval"
	Circle Shape = "circle"
	Oblong Shape = "oblong"

	VTriangle Shape = "v_triangle"
	Square    Shape = "Square"
	Diamond   Shape = "diamond"

	Rectangle Shape = "rectangle"
	Heart     Shape = "heart"
	ATriangle Shape = "a_triangle"
)

func (s Shape) Zero() bool {
	return s == ""
}

func (s Shape) Bytes() []byte {
	if s.Zero() {
		return []byte("")
	}
	return []byte(s)
}

type FaceShapeGene struct {
	t     string
	Stats map[string]float64
}

func NewFaceShapeGene(stats map[string]float64) gene.Gene {
	return &FaceShapeGene{
		t:     "face_shape_gene",
		Stats: stats,
	}
}

func (g *FaceShapeGene) Type() string {
	return g.t
}

func (g *FaceShapeGene) Produce() (gene.Byteble, error) {
	return Shape(pm.GetRandomFromSeveral(g.Stats)), nil
}

func (g *FaceShapeGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *FaceShapeGene) Bytes() []byte {
	return nil
}

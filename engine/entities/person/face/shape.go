package face

import (
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
	pm "persons_generator/engine/probability_machine"
)

type Shape string

const (
	Oval   Shape = "oval"
	Circle Shape = "circle"
	Oblong Shape = "oblong"

	VTriangle Shape = "v_triangle"
	Square    Shape = "square"
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
	T string `json:"t"`

	Stats map[string]float64 `json:"stats"`
}

func NewFaceShapeGene(stats map[string]float64) gene.Gene {
	return &FaceShapeGene{
		T: "face_shape_gene",

		Stats: stats,
	}
}

func (g *FaceShapeGene) Type() string {
	return g.T
}

func (g *FaceShapeGene) Produce(sex g.Sex) (gene.Byteble, error) {
	return Shape(pm.GetRandomFromSeveral(g.Stats)), nil
}

func (g *FaceShapeGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *FaceShapeGene) Bytes() []byte {
	return nil
}

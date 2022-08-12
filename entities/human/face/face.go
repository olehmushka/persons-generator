package face

import (
	"encoding/json"

	"persons_generator/entities/human/gene"
)

type Face struct {
	Shape Shape
	Eyes  *Eyes
	Ears  *Ears
	Nose  *Nose
	Mouth *Mouth
}

func (n Face) Zero() bool {
	return n == Face{}
}

func (f Face) Bytes() []byte {
	if f.Zero() {
		return nil
	}
	if out, err := json.Marshal(f); err == nil {
		return out
	}

	return nil
}

type FaceGene struct {
	t string

	ShapeGene gene.Gene
	EyesGene  gene.Gene
	EarsGene  gene.Gene
	NoseGene  gene.Gene
	MouthGene gene.Gene
}

func NewFaceGene(
	shapeGene gene.Gene,
	eyesGene gene.Gene,
	earsGene gene.Gene,
	noseGene gene.Gene,
	mouthGene gene.Gene,
) gene.Gene {
	return &FaceGene{
		t:         "face_gene",
		ShapeGene: shapeGene,
		EyesGene:  eyesGene,
		EarsGene:  earsGene,
		NoseGene:  noseGene,
		MouthGene: mouthGene,
	}
}

func (g *FaceGene) Type() string {
	return g.t
}

func (g *FaceGene) Produce() (gene.Byteble, error) {
	shape, err := g.ShapeGene.Produce()
	if err != nil {
		return Face{}, err
	}

	eyes, err := g.EyesGene.Produce()
	if err != nil {
		return Face{}, err
	}
	var eyesVal Eyes
	if err := json.Unmarshal(eyes.Bytes(), &eyesVal); err != nil {
		return Face{}, err
	}

	ears, err := g.EarsGene.Produce()
	if err != nil {
		return Face{}, err
	}
	var earsVal Ears
	if err := json.Unmarshal(ears.Bytes(), &earsVal); err != nil {
		return Face{}, err
	}

	nose, err := g.NoseGene.Produce()
	if err != nil {
		return Face{}, err
	}
	var noseVal Nose
	if err := json.Unmarshal(nose.Bytes(), &noseVal); err != nil {
		return Face{}, err
	}

	mouth, err := g.MouthGene.Produce()
	if err != nil {
		return Face{}, err
	}
	var mouthVal Mouth
	if err := json.Unmarshal(mouth.Bytes(), &mouthVal); err != nil {
		return Face{}, err
	}

	return Face{
		Shape: Shape(string(shape.Bytes())),
		Eyes:  &eyesVal,
		Ears:  &earsVal,
		Nose:  &noseVal,
		Mouth: &mouthVal,
	}, nil
}

func (g *FaceGene) Children() []gene.Gene {
	return []gene.Gene{
		g.ShapeGene,
		g.EyesGene,
		g.EarsGene,
		g.NoseGene,
		g.MouthGene,
	}
}

func (g *FaceGene) Bytes() []byte {
	return nil
}

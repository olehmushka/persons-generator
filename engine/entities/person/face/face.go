package face

import (
	"encoding/json"
	"fmt"

	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type Face struct {
	Shape Shape  `json:"shape"`
	Eyes  *Eyes  `json:"eyes"`
	Ears  *Ears  `json:"ears"`
	Nose  *Nose  `json:"nose"`
	Mouth *Mouth `json:"mouth"`
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
	T string `json:"t"`

	ShapeGene gene.Gene `json:"shape_gene"`
	EyesGene  gene.Gene `json:"eyes_gene"`
	EarsGene  gene.Gene `json:"ears_gene"`
	NoseGene  gene.Gene `json:"nose_gene"`
	MouthGene gene.Gene `json:"mouth_gene"`
}

func NewFaceGene(
	shapeGene gene.Gene,
	eyesGene gene.Gene,
	earsGene gene.Gene,
	noseGene gene.Gene,
	mouthGene gene.Gene,
) gene.Gene {
	return &FaceGene{
		T: "face_gene",

		ShapeGene: shapeGene,
		EyesGene:  eyesGene,
		EarsGene:  earsGene,
		NoseGene:  noseGene,
		MouthGene: mouthGene,
	}
}

func (g *FaceGene) Type() string {
	return g.T
}

func (g *FaceGene) Produce(sex g.Sex) (gene.Byteble, error) {
	shape, err := g.ShapeGene.Produce(sex)
	if err != nil {
		return Face{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce shape from gene by sex (sex=%s)", sex))
	}

	eyes, err := g.EyesGene.Produce(sex)
	if err != nil {
		return Face{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce eyes from gene by sex (sex=%s)", sex))
	}
	var eyesVal Eyes
	if err := json.Unmarshal(eyes.Bytes(), &eyesVal); err != nil {
		return Face{}, wrapped_error.NewInternalServerError(err, "can not unmarshal eyes produced from gene")
	}

	ears, err := g.EarsGene.Produce(sex)
	if err != nil {
		return Face{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce ears from gene by sex (sex=%s)", sex))
	}
	var earsVal Ears
	if err := json.Unmarshal(ears.Bytes(), &earsVal); err != nil {
		return Face{}, wrapped_error.NewInternalServerError(err, "can not unmarshal ears produced from gene")
	}

	nose, err := g.NoseGene.Produce(sex)
	if err != nil {
		return Face{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce nose from gene by sex (sex=%s)", sex))
	}
	var noseVal Nose
	if err := json.Unmarshal(nose.Bytes(), &noseVal); err != nil {
		return Face{}, wrapped_error.NewInternalServerError(err, "can not unmarshal nose produced from gene")
	}

	mouth, err := g.MouthGene.Produce(sex)
	if err != nil {
		return Face{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce mouth from gene by sex (sex=%s)", sex))
	}
	var mouthVal Mouth
	if err := json.Unmarshal(mouth.Bytes(), &mouthVal); err != nil {
		return Face{}, wrapped_error.NewInternalServerError(err, "can not unmarshal mouth produced from gene")
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

func (g *FaceGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> face genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}
	var (
		hostChildren  = g.Children()
		guestChildren = in.Children()
	)
	if len(g.Children()) != len(in.Children()) {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same number of children for face gene (host_children_number=%d, guest_children_number=%d)", len(hostChildren), len(guestChildren)))
	}

	args := make([]gene.Gene, 0, len(hostChildren))
	for i, child := range hostChildren {
		arg, err := child.Pair(guestChildren[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not pair children for face gene")
		}
		args = append(args, arg)
	}

	return NewFaceGene(args[0], args[1], args[2], args[3], args[4]), nil
}

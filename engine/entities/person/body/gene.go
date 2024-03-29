package body

import (
	"encoding/json"
	"fmt"

	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/face"
	"persons_generator/engine/entities/person/gene"
	"persons_generator/engine/entities/person/hair"
	"persons_generator/engine/entities/person/size"
	"persons_generator/engine/entities/person/skin"
)

type BodyGene struct {
	T string `json:"t"`

	FaceGene gene.Gene `json:"face_gene"`
	HairGene gene.Gene `json:"hair_gene"`
	SizeGene gene.Gene `json:"size_gene"`
	SkinGene gene.Gene `json:"skin_gene"`
}

func NewBodyGene(
	faceGene gene.Gene,
	hairGene gene.Gene,
	sizeGene gene.Gene,
	skinGene gene.Gene,
) gene.Gene {
	return &BodyGene{
		T: "body_gene",

		FaceGene: faceGene,
		HairGene: hairGene,
		SizeGene: sizeGene,
		SkinGene: skinGene,
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
		return Body{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce body_hair from gene by sex (sex=%s)", sex))
	}
	var hairVal hair.Hair
	if err := json.Unmarshal(h.Bytes(), &hairVal); err != nil {
		return Body{}, wrapped_error.NewInternalServerError(err, "can not unmarshal body_hair produced from gene")
	}

	s, err := g.SizeGene.Produce(sex)
	if err != nil {
		return Body{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce body_size from gene by sex (sex=%s)", sex))
	}
	var sizeVal size.Size
	if err := json.Unmarshal(s.Bytes(), &sizeVal); err != nil {
		return Body{}, wrapped_error.NewInternalServerError(err, "can not unmarshal body_size produced from gene")
	}

	sn, err := g.SkinGene.Produce(sex)
	if err != nil {
		return Body{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce body_skin from gene by sex (sex=%s)", sex))
	}
	var skinVal skin.Skin
	if err := json.Unmarshal(sn.Bytes(), &skinVal); err != nil {
		return Body{}, wrapped_error.NewInternalServerError(err, "can not unmarshal body_skin produced from gene")
	}

	return Body{
		Face: faceVal,
		Hair: hairVal,
		Size: sizeVal,
		Skin: skinVal,
	}, nil
}

func (g *BodyGene) Children() []gene.Gene {
	return []gene.Gene{
		g.FaceGene,
		g.HairGene,
		g.SizeGene,
		g.SkinGene,
	}
}

func (g *BodyGene) Bytes() []byte {
	return nil
}

func (g *BodyGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> body genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}
	var (
		hostChildren  = g.Children()
		guestChildren = in.Children()
	)
	if len(g.Children()) != len(in.Children()) {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same number of children for body gene (host_children_number=%d, guest_children_number=%d)", len(hostChildren), len(guestChildren)))
	}

	args := make([]gene.Gene, 0, len(hostChildren))
	for i, child := range hostChildren {
		arg, err := child.Pair(guestChildren[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not pair children for body gene")
		}
		args = append(args, arg)
	}

	return NewBodyGene(args[0], args[1], args[2], args[3]), nil
}

package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type FaceHair struct {
	Density *HairDensity `json:"density"`
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

	HairDensityGene gene.Gene `json:"hair_density_gene"`
}

func NewFaceHairGene(
	feetHairDensityGene gene.Gene,
) gene.Gene {
	return &FaceHairGene{
		T: "face_hair_gene",

		HairDensityGene: feetHairDensityGene,
	}
}

func (g *FaceHairGene) Type() string {
	return g.T
}

func (g *FaceHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	density, err := g.HairDensityGene.Produce(sex)
	if err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce face_hair_density from gene by sex (sex=%s)", sex))
	}
	var densityVal HairDensity
	if err := json.Unmarshal(density.Bytes(), &densityVal); err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal face_hair_density produced from gene")
	}

	return FaceHair{
		Density: &densityVal,
	}, nil
}

func (g *FaceHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairDensityGene,
	}
}

func (g *FaceHairGene) Bytes() []byte {
	return nil
}

func (g *FaceHairGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> face_hair genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}
	var (
		hostChildren  = g.Children()
		guestChildren = in.Children()
	)
	if len(g.Children()) != len(in.Children()) {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same number of children for face_hair gene (host_children_number=%d, guest_children_number=%d)", len(hostChildren), len(guestChildren)))
	}

	args := make([]gene.Gene, 0, len(hostChildren))
	for i, child := range hostChildren {
		arg, err := child.Pair(guestChildren[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not pair children for face_hair gene")
		}
		args = append(args, arg)
	}

	return NewFaceHairGene(args[0]), nil
}

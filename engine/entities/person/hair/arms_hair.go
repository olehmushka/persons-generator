package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type ArmsHair struct {
	Density *HairDensity `json:"density"`
}

func (ah ArmsHair) Zero() bool {
	return ah == ArmsHair{}
}

func (ah ArmsHair) Bytes() []byte {
	if ah.Zero() {
		return nil
	}

	if out, err := json.Marshal(ah); err == nil {
		return out
	}

	return nil
}

type ArmsHairGene struct {
	T string `json:"t"`

	HairDensityGene gene.Gene `json:"hair_density_gene"`
}

func NewArmsHairGene(
	armsHairDensityGene gene.Gene,
) gene.Gene {
	return &ArmsHairGene{
		T: "arms_hair_gene",

		HairDensityGene: armsHairDensityGene,
	}
}

func (g *ArmsHairGene) Type() string {
	return g.T
}

func (g *ArmsHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	density, err := g.HairDensityGene.Produce(sex)
	if err != nil {
		return ArmsHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce arms_hair_density from gene by sex (sex=%s)", sex))
	}
	var densityVal HairDensity
	if err := json.Unmarshal(density.Bytes(), &densityVal); err != nil {
		return ArmsHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal arms_hair_density produced from gene")
	}

	return ArmsHair{
		Density: &densityVal,
	}, nil
}

func (g *ArmsHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairDensityGene,
	}
}

func (g *ArmsHairGene) Bytes() []byte {
	return nil
}

func (g *ArmsHairGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> arms_hair genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}
	var (
		hostChildren  = g.Children()
		guestChildren = in.Children()
	)
	if len(g.Children()) != len(in.Children()) {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same number of children for arms_hair gene (host_children_number=%d, guest_children_number=%d)", len(hostChildren), len(guestChildren)))
	}

	args := make([]gene.Gene, 0, len(hostChildren))
	for i, child := range hostChildren {
		arg, err := child.Pair(guestChildren[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not pair children for arms_hair gene")
		}
		args = append(args, arg)
	}

	return NewArmsHairGene(args[0]), nil
}

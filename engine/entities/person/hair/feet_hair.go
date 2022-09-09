package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type FeetHair struct {
	Density *HairDensity `json:"density"`
}

func (fh FeetHair) Zero() bool {
	return fh == FeetHair{}
}

func (fh FeetHair) Bytes() []byte {
	if fh.Zero() {
		return nil
	}

	if out, err := json.Marshal(fh); err == nil {
		return out
	}

	return nil
}

type FeetHairGene struct {
	T string `json:"t"`

	HairDensityGene gene.Gene `json:"hair_density_gene"`
}

func NewFeetHairGene(
	feetHairDensityGene gene.Gene,
) gene.Gene {
	return &FeetHairGene{
		T: "feet_hair_gene",

		HairDensityGene: feetHairDensityGene,
	}
}

func (g *FeetHairGene) Type() string {
	return g.T
}

func (g *FeetHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	density, err := g.HairDensityGene.Produce(sex)
	if err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce feet_hair_density from gene by sex (sex=%s)", sex))
	}
	var densityVal HairDensity
	if err := json.Unmarshal(density.Bytes(), &densityVal); err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal feet_hair_density produced from gene")
	}

	return FeetHair{
		Density: &densityVal,
	}, nil
}

func (g *FeetHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairDensityGene,
	}
}

func (g *FeetHairGene) Bytes() []byte {
	return nil
}

func (g *FeetHairGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> feet_hair genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}
	var (
		hostChildren  = g.Children()
		guestChildren = in.Children()
	)
	if len(g.Children()) != len(in.Children()) {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same number of children for feet_hair gene (host_children_number=%d, guest_children_number=%d)", len(hostChildren), len(guestChildren)))
	}

	args := make([]gene.Gene, 0, len(hostChildren))
	for i, child := range hostChildren {
		arg, err := child.Pair(guestChildren[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not pair children for feet_hair gene")
		}
		args = append(args, arg)
	}

	return NewFeetHairGene(args[0]), nil
}

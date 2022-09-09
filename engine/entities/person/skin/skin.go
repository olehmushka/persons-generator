package skin

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type Skin struct {
	Color *SkinColor `json:"color"`
}

func (s Skin) Zero() bool {
	return s == Skin{}
}

func (s Skin) Bytes() []byte {
	if s.Zero() {
		return nil
	}
	if out, err := json.Marshal(s); err == nil {
		return out
	}

	return nil
}

type SkinGene struct {
	T string `json:"t"`

	SkinColorGene gene.Gene `json:"skin_color_gene"`
}

func NewSkinGene(
	skinColorGene gene.Gene,
) gene.Gene {
	return &SkinGene{
		T: "skin_gene",

		SkinColorGene: skinColorGene,
	}
}

func (g *SkinGene) Type() string {
	return g.T
}

func (g *SkinGene) Produce(sex gender.Sex) (gene.Byteble, error) {
	color, err := g.SkinColorGene.Produce(sex)
	if err != nil {
		return Skin{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce skin_color from gene by sex (sex=%s)", sex))
	}
	var colorVal SkinColor
	if err := json.Unmarshal(color.Bytes(), &colorVal); err != nil {
		return Skin{}, wrapped_error.NewInternalServerError(err, "can not unmarshal skin_color produced from gene")
	}

	return Skin{
		Color: &colorVal,
	}, nil
}

func (g *SkinGene) Children() []gene.Gene {
	return []gene.Gene{
		g.SkinColorGene,
	}
}

func (g *SkinGene) Bytes() []byte {
	return nil
}

func (g *SkinGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> skin genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}
	var (
		hostChildren  = g.Children()
		guestChildren = in.Children()
	)
	if len(g.Children()) != len(in.Children()) {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same number of children for skin gene (host_children_number=%d, guest_children_number=%d)", len(hostChildren), len(guestChildren)))
	}

	args := make([]gene.Gene, 0, len(hostChildren))
	for i, child := range hostChildren {
		arg, err := child.Pair(guestChildren[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not pair children for skin gene")
		}
		args = append(args, arg)
	}

	return NewSkinGene(args[0]), nil
}

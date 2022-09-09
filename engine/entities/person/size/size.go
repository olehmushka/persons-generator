package size

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type Size struct {
	Height   *Height   `json:"height"`
	ShoeSize *ShoeSize `json:"shoe_size"`
}

func (s Size) Zero() bool {
	return s == Size{}
}

func (s Size) Bytes() []byte {
	if s.Zero() {
		return nil
	}
	if out, err := json.Marshal(s); err == nil {
		return out
	}

	return nil
}

type SizeGene struct {
	T string `json:"t"`

	HeightGene   gene.Gene
	ShoeSizeGene gene.Gene
}

func NewSizeGene(
	heightGene gene.Gene,
	shoeSizeGene gene.Gene,
) gene.Gene {
	return &SizeGene{
		T: "size_gene",

		HeightGene:   heightGene,
		ShoeSizeGene: shoeSizeGene,
	}
}

func (g *SizeGene) Type() string {
	return g.T
}

func (g *SizeGene) Produce(sex gender.Sex) (gene.Byteble, error) {
	height, err := g.HeightGene.Produce(sex)
	if err != nil {
		return Size{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce height_size from gene by sex (sex=%s)", sex))
	}
	var heightVal Height
	if err := json.Unmarshal(height.Bytes(), &heightVal); err != nil {
		return Size{}, wrapped_error.NewInternalServerError(err, "can not unmarshal height_size produced from gene")
	}

	shoeSize, err := g.ShoeSizeGene.Produce(sex)
	if err != nil {
		return Size{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce shoe_size from gene by sex (sex=%s)", sex))
	}
	var shoeSizeVal ShoeSize
	if err := json.Unmarshal(shoeSize.Bytes(), &shoeSizeVal); err != nil {
		return Size{}, wrapped_error.NewInternalServerError(err, "can not unmarshal shoe_size produced from gene")
	}

	return Size{
		Height:   &heightVal,
		ShoeSize: &shoeSizeVal,
	}, nil
}

func (g *SizeGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HeightGene,
		g.ShoeSizeGene,
	}
}

func (g *SizeGene) Bytes() []byte {
	return nil
}

func (g *SizeGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> size genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}
	var (
		hostChildren  = g.Children()
		guestChildren = in.Children()
	)
	if len(g.Children()) != len(in.Children()) {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same number of children for size gene (host_children_number=%d, guest_children_number=%d)", len(hostChildren), len(guestChildren)))
	}

	args := make([]gene.Gene, 0, len(hostChildren))
	for i, child := range hostChildren {
		arg, err := child.Pair(guestChildren[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not pair children for size gene")
		}
		args = append(args, arg)
	}

	return NewSizeGene(args[0], args[1]), nil
}

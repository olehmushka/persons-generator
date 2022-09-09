package hair

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type ScaplHair struct {
	Texture *HairTexture `json:"texture"`
	Density *HairDensity `json:"density"`
}

func (sh ScaplHair) Zero() bool {
	return sh == ScaplHair{}
}

func (sh ScaplHair) Bytes() []byte {
	if sh.Zero() {
		return nil
	}

	if out, err := json.Marshal(sh); err == nil {
		return out
	}

	return nil
}

type ScaplHairGene struct {
	T string `json:"t"`

	HairTextureGene gene.Gene `json:"hair_texture_gene"`
	HairDensityGene gene.Gene `json:"hair_density_gene"`
}

func NewScalpHairGene(
	scalpHairTextureGene gene.Gene,
	scalpHairDensityGene gene.Gene,
) gene.Gene {
	return &ScaplHairGene{
		T: "scapl_hair_gene",

		HairTextureGene: scalpHairTextureGene,
		HairDensityGene: scalpHairDensityGene,
	}
}

func (g *ScaplHairGene) Type() string {
	return g.T
}

func (g *ScaplHairGene) Produce(sex g.Sex) (gene.Byteble, error) {
	texture, err := g.HairTextureGene.Produce(sex)
	if err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce scapl_hair_texture from gene by sex (sex=%s)", sex))
	}
	var textureVal HairTexture
	if err := json.Unmarshal(texture.Bytes(), &textureVal); err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal scapl_hair_texture produced from gene")
	}

	density, err := g.HairDensityGene.Produce(sex)
	if err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce scapl_hair_density from gene by sex (sex=%s)", sex))
	}
	var densityVal HairDensity
	if err := json.Unmarshal(density.Bytes(), &densityVal); err != nil {
		return AxillaryHair{}, wrapped_error.NewInternalServerError(err, "can not unmarshal scapl_hair_density produced from gene")
	}

	return ScaplHair{
		Texture: &textureVal,
		Density: &densityVal,
	}, nil
}

func (g *ScaplHairGene) Children() []gene.Gene {
	return []gene.Gene{
		g.HairTextureGene,
		g.HairDensityGene,
	}
}

func (g *ScaplHairGene) Bytes() []byte {
	return nil
}

func (g *ScaplHairGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> scalp_hair genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}
	var (
		hostChildren  = g.Children()
		guestChildren = in.Children()
	)
	if len(g.Children()) != len(in.Children()) {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same number of children for scalp_hair gene (host_children_number=%d, guest_children_number=%d)", len(hostChildren), len(guestChildren)))
	}

	args := make([]gene.Gene, 0, len(hostChildren))
	for i, child := range hostChildren {
		arg, err := child.Pair(guestChildren[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not pair children for scalp_hair gene")
		}
		args = append(args, arg)
	}

	return NewScalpHairGene(args[0], args[1]), nil
}

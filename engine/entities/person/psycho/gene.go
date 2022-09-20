package psycho

import (
	"encoding/json"
	"fmt"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
	"persons_generator/engine/entities/person/temperament"
)

type PsychoGene struct {
	T string `json:"t"`

	TemperamentGene gene.Gene `json:"temperament_gene"`
}

func NewPsychoGene(
	temperamentGene gene.Gene,
) gene.Gene {
	return &PsychoGene{
		T: "psycho_gene",

		TemperamentGene: temperamentGene,
	}
}

func (g *PsychoGene) Type() string {
	return g.T
}

func (g *PsychoGene) Produce(sex gender.Sex) (gene.Byteble, error) {
	t, err := g.TemperamentGene.Produce(sex)
	if err != nil {
		return Psycho{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce temperament from gene by sex (sex=%s)", sex))
	}
	var temperamentVal temperament.Temperament
	if err := json.Unmarshal(t.Bytes(), &temperamentVal); err != nil {
		return Psycho{}, wrapped_error.NewInternalServerError(err, "can not unmarshal temperament produced from gene")
	}

	so, err := ProduceSexualOrientation()
	if err != nil {
		return Psycho{}, err
	}

	return Psycho{
		Temperament:       temperamentVal,
		SexualOrientation: so,
	}, nil
}

func (g *PsychoGene) Children() []gene.Gene {
	return []gene.Gene{
		g.TemperamentGene,
	}
}

func (g *PsychoGene) Bytes() []byte {
	return nil
}

func (g *PsychoGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> psycho genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}
	var (
		hostChildren  = g.Children()
		guestChildren = in.Children()
	)
	if len(g.Children()) != len(in.Children()) {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same number of children for psycho gene (host_children_number=%d, guest_children_number=%d)", len(hostChildren), len(guestChildren)))
	}

	args := make([]gene.Gene, 0, len(hostChildren))
	for i, child := range hostChildren {
		arg, err := child.Pair(guestChildren[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not pair children for psycho gene")
		}
		args = append(args, arg)
	}

	return NewPsychoGene(args[0]), nil
}

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

	return Psycho{
		Temperament: temperamentVal,
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
	return in, nil
}

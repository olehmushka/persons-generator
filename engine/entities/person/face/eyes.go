package face

import (
	"encoding/json"
	"fmt"

	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
)

type Eyes struct {
	Type  EyesType  `json:"type"`
	Color EyesColor `json:"color"`
}

func (n Eyes) Zero() bool {
	return n == Eyes{}
}

func (n Eyes) Bytes() []byte {
	if n.Zero() {
		return nil
	}

	if out, err := json.Marshal(n); err == nil {
		return out
	}

	return nil
}

type EyesGene struct {
	T string `json:"t"`

	EyesTypeGene  gene.Gene `json:"eyes_type_gene"`
	EyesColorGene gene.Gene `json:"eyes_color_gene"`
}

func NewEyesGene(eyesTypeGene gene.Gene, eyesColorGene gene.Gene) gene.Gene {
	return &EyesGene{
		T: "eyes_gene",

		EyesTypeGene:  eyesTypeGene,
		EyesColorGene: eyesColorGene,
	}
}

func (g *EyesGene) Type() string {
	return g.T
}

func (g *EyesGene) Produce(sex g.Sex) (gene.Byteble, error) {
	eyesColor, err := g.EyesColorGene.Produce(sex)
	if err != nil {
		return Eyes{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce eye_color from gene by sex (sex=%s)", sex))
	}
	var colorVal EyesColor
	if err := json.Unmarshal(eyesColor.Bytes(), &colorVal); err != nil {
		return Eyes{}, wrapped_error.NewInternalServerError(err, "can not unmarshal eye_color produced from gene")
	}

	eyesType, err := g.EyesTypeGene.Produce(sex)
	if err != nil {
		return Eyes{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not produce eye_type from gene by sex (sex=%s)", sex))
	}

	return Eyes{
		Type:  EyesType(string(eyesType.Bytes())),
		Color: colorVal,
	}, nil
}

func (g *EyesGene) Children() []gene.Gene {
	return []gene.Gene{
		g.EyesColorGene,
		g.EyesTypeGene,
	}
}

func (g *EyesGene) Bytes() []byte {
	return nil
}

func (g *EyesGene) Pair(in gene.Gene) (gene.Gene, error) {
	return in, nil
}

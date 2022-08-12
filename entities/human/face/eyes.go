package face

import (
	"encoding/json"

	"persons_generator/entities/human/gene"
)

type Eyes struct {
	Type  EyesType
	Color EyesColor
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
	t string

	EyesTypeGene  gene.Gene
	EyesColorGene gene.Gene
}

func NewEyesGene(eyesTypeGene gene.Gene, eyesColorGene gene.Gene) gene.Gene {
	return &EyesGene{
		t:             "eyes_gene",
		EyesTypeGene:  eyesTypeGene,
		EyesColorGene: eyesColorGene,
	}
}

func (g *EyesGene) Type() string {
	return g.t
}

func (g *EyesGene) Produce() (gene.Byteble, error) {
	eyesColor, err := g.EyesColorGene.Produce()
	if err != nil {
		return Eyes{}, err
	}
	var colorVal EyesColor
	if err := json.Unmarshal(eyesColor.Bytes(), &colorVal); err != nil {
		return Eyes{}, err
	}

	eyesType, err := g.EyesTypeGene.Produce()
	if err != nil {
		return Eyes{}, err
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

package face

import (
	"encoding/json"
	"fmt"

	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/person/gene"
	pm "persons_generator/engine/probability_machine"
)

type Mouth struct {
	LipsType LipsType `json:"lips_type"`
}

func (n Mouth) Zero() bool {
	return n == Mouth{}
}

func (n Mouth) Bytes() []byte {
	if n.Zero() {
		return nil
	}

	if out, err := json.Marshal(n); err == nil {
		return out
	}

	return nil
}

type MouthGene struct {
	T string `json:"t"`

	LipsTypeStats map[string]float64 `json:"lips_type_stats"`
}

func NewMouthGene(lipsStats map[string]float64) gene.Gene {
	return &MouthGene{
		T: "mouth_gene",

		LipsTypeStats: lipsStats,
	}
}

func (g *MouthGene) Type() string {
	return g.T
}

func (g *MouthGene) Produce(sex g.Sex) (gene.Byteble, error) {
	lt, err := pm.GetRandomFromSeveral(g.LipsTypeStats)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate lips type")
	}

	return Mouth{LipsType: LipsType(lt)}, nil
}

func (g *MouthGene) Children() []gene.Gene {
	return []gene.Gene{}
}

func (g *MouthGene) Bytes() []byte {
	return nil
}

func (g *MouthGene) Pair(in gene.Gene) (gene.Gene, error) {
	if in == nil || g == nil {
		return nil, wrapped_error.NewInternalServerError(nil, "can not pair <nil> mouth genes")
	}
	if g.Type() != in.Type() {
		return nil, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not pair genes with not the same types (first_type=%s, second_type=%s)", g.Type(), in.Type()))
	}

	inStr, ok := in.(*MouthGene)
	if !ok {
		return nil, wrapped_error.NewInternalServerError(nil, "can not case input gene to MouthGene")
	}

	return NewMouthGene(tools.MergeMaps(g.LipsTypeStats, inStr.LipsTypeStats)), nil
}

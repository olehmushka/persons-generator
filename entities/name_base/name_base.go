package culture

import (
	bc "persons_generator/entities/base_culture"
	pm "persons_generator/probability_machine"
)

type NameBase struct {
	Culture *bc.BaseCulture
	Words   []string
	Index   int
	Min     int
	Max     int
	Dupl    string
	M       float64
}

func NameBasesToStrSlices(nbs []*NameBase) [][]string {
	out := make([][]string, len(nbs))
	for i := range out {
		out[i] = nbs[i].Words
	}

	return out
}

func GetRandomNmeBase() *NameBase {
	return AllNameBases[pm.RandInt(len(AllNameBases)-1)]
}

package religion

import "math"

type GoodnessNature struct {
	Level    Level    `json:"level"`
	Goodness Goodness `json:"goodness"`
}

func GetGoodnessNatureSimilarityCoef(gn1, gn2 GoodnessNature) float64 {
	return math.Abs(gn1.Goodness.Value()*gn1.Level.GetCoef() - gn2.Goodness.Value()*gn2.Level.GetCoef())
}

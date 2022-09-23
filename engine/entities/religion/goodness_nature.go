package religion

import (
	"fmt"
	"math"
)

type GoodnessNature struct {
	Level    Level    `json:"level"`
	Goodness Goodness `json:"goodness"`
}

func (gn *GoodnessNature) String() string {
	if gn == nil {
		return ""
	}

	return fmt.Sprintf("%s (%s)", gn.Goodness.String(), gn.Level.String())
}

func GetGoodnessNatureSimilarityCoef(gn1, gn2 GoodnessNature) float64 {
	return math.Abs(gn1.Goodness.Value()*gn1.Level.GetCoef() - gn2.Goodness.Value()*gn2.Level.GetCoef())
}

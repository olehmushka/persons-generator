package influence

import (
	"persons_generator/core/wrapped_error"
	pm "persons_generator/engine/probability_machine"
)

type Influence string

const (
	StrongInfluence   Influence = "strong"
	ModerateInfluence Influence = "moderate"
	WeakInfluence     Influence = "weak"
)

func GetInfluenceByProbability(strong, moderate, weak float64) (Influence, error) {
	i, err := pm.GetRandomFromSeveral(map[string]float64{
		string(StrongInfluence):   pm.PrepareProbability(strong),
		string(ModerateInfluence): pm.PrepareProbability(moderate),
		string(WeakInfluence):     pm.PrepareProbability(weak),
	})
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generate infuence")
	}

	return Influence(i), nil
}

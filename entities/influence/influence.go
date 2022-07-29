package influence

import pm "persons_generator/probability_machine"

type Influence string

const (
	StrongInfluence   Influence = "strong"
	ModerateInfluence Influence = "moderate"
	WeakInfluence     Influence = "weak"
)

func GetInfluenceByProbability(strong, moderate, weak float64) Influence {
	return Influence(pm.GetRandomFromSeveral(map[string]float64{
		string(StrongInfluence):   pm.PrepareProbability(strong),
		string(ModerateInfluence): pm.PrepareProbability(moderate),
		string(WeakInfluence):     pm.PrepareProbability(weak),
	}))
}

package religion

import pm "persons_generator/probability-machine"

type Influence string

const (
	StrongInfluence   Influence = "Strong"
	ModerateInfluence Influence = "Moderate"
	WeakInfluence     Influence = "Weak"
)

func geInfluenceByProbability(strong, moderate, weak float64) Influence {
	m := map[string]float64{
		string(StrongInfluence):   pm.PrepareProbability(strong),
		string(ModerateInfluence): pm.PrepareProbability(moderate),
		string(WeakInfluence):     pm.PrepareProbability(weak),
	}

	return Influence(pm.GetRandomFromSeveral(m))
}

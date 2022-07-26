package religion

import pm "persons_generator/probability_machine"

type Goodness string

const (
	Good    Goodness = "good"
	Neutral Goodness = "neutral"
	Evil    Goodness = "evil"
)

func getGoodnessByProbability(good, neutral, evil float64) Goodness {
	return Goodness(pm.GetRandomFromSeveral(map[string]float64{
		string(Good):    pm.PrepareProbability(good),
		string(Neutral): pm.PrepareProbability(neutral),
		string(Evil):    pm.PrepareProbability(evil),
	}))
}

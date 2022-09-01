package religion

import pm "persons_generator/engine/probability_machine"

type Goodness string

func (g Goodness) String() string {
	return string(g)
}

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

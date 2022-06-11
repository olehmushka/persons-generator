package religion

import pm "persons_generator/probability-machine"

type Goodness string

const (
	Good    Goodness = "Good"
	Neutral Goodness = "Neutral"
	Evil    Goodness = "Evil"
)

func getGoodnessByProbability(good, neutral, evil float64) Goodness {
	m := map[string]float64{
		string(Good):    pm.PrepareProbability(good),
		string(Neutral): pm.PrepareProbability(neutral),
		string(Evil):    pm.PrepareProbability(evil),
	}
	return Goodness(pm.GetRandomFromSeveral(m))
}

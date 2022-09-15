package religion

import pm "persons_generator/engine/probability_machine"

type Goodness string

func (g Goodness) String() string {
	return string(g)
}

func (g Goodness) Value() float64 {
	if g == Good {
		return 1
	}
	if g == Neutral {
		return 0.5
	}
	if g == Evil {
		return 0
	}

	return 0
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

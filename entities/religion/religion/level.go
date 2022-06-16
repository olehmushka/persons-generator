package religion

import pm "persons_generator/probability-machine"

type Level string

const (
	Major  Level = "Major"
	Middle Level = "Middle"
	Minor  Level = "Minor"
)

func getLevelByProbability(major, middle, minor float64) Level {
	m := map[string]float64{
		string(Major):  pm.PrepareProbability(major),
		string(Middle): pm.PrepareProbability(middle),
		string(Minor):  pm.PrepareProbability(minor),
	}
	return Level(pm.GetRandomFromSeveral(m))
}

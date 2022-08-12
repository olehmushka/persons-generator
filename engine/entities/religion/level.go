package religion

import pm "persons_generator/engine/probability_machine"

type Level string

const (
	Major  Level = "major"
	Middle Level = "middle"
	Minor  Level = "minor"
)

func getLevelByProbability(major, middle, minor float64) Level {
	return Level(pm.GetRandomFromSeveral(map[string]float64{
		string(Major):  pm.PrepareProbability(major),
		string(Middle): pm.PrepareProbability(middle),
		string(Minor):  pm.PrepareProbability(minor),
	}))
}

func (l Level) GetLevelCoef() float64 {
	switch l {
	case Major:
		return 1.25
	case Middle:
		return 1
	case Minor:
		return 0.75
	}

	return 1
}

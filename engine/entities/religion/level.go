package religion

import (
	"persons_generator/core/wrapped_error"
	pm "persons_generator/engine/probability_machine"
)

type Level string

func (l Level) String() string {
	return string(l)
}

const (
	Major  Level = "major"
	Middle Level = "middle"
	Minor  Level = "minor"
)

func getLevelByProbability(major, middle, minor float64) (Level, error) {
	l, err := pm.GetRandomFromSeveral(map[string]float64{
		string(Major):  pm.PrepareProbability(major),
		string(Middle): pm.PrepareProbability(middle),
		string(Minor):  pm.PrepareProbability(minor),
	})
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generate level")
	}
	return Level(l), nil
}

func (l Level) GetCoef() float64 {
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

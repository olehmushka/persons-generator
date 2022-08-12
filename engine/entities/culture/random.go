package culture

import (
	"errors"

	"persons_generator/core/tools"
	pm "persons_generator/engine/probability_machine"
)

func GetRandomCultureName(cultures []*Culture) (string, error) {
	if len(cultures) == 0 {
		return "", errors.New("can not get random culture_name from empty cultures slice")
	}

	return tools.RandomValueOfSlice(pm.RandFloat64, cultures).Name, nil
}

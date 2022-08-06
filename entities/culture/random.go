package culture

import (
	"errors"

	pm "persons_generator/probability_machine"
	"persons_generator/tools"
)

func GetRandomCultureName(cultures []*Culture) (string, error) {
	if len(cultures) == 0 {
		return "", errors.New("can not get random culture_name from empty cultures slice")
	}

	return tools.RandomValueOfSlice(pm.RandFloat64, cultures).Name, nil
}

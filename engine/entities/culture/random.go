package culture

import (
	"net/http"
	"persons_generator/core/tools"
	we "persons_generator/core/wrapped_error"
	pm "persons_generator/engine/probability_machine"
)

func GetRandomCultureName(cultures []*Culture) (string, error) {
	if len(cultures) == 0 {
		return "", we.New(http.StatusInternalServerError, nil, "can not get random culture_name from empty cultures slice")
	}
	c, err := tools.RandomValueOfSlice(pm.RandFloat64, cultures)
	if err != nil {
		return "", err
	}

	return c.Name, nil
}

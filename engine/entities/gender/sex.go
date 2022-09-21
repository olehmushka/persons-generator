package gender

import (
	"persons_generator/core/wrapped_error"
	pm "persons_generator/engine/probability_machine"
)

type Sex string

const (
	MaleSex   Sex = "male_sex"
	FemaleSex Sex = "female_sex"
)

func (s Sex) String() string {
	return string(s)
}

func (s Sex) IsMale() bool {
	return s == MaleSex
}

func (s Sex) IsFemale() bool {
	return s == FemaleSex
}

func GetRandomSex() (Sex, error) {
	isMale, err := pm.GetRandomBool(0.5)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generate bool for generation random sex")
	}
	if isMale {
		return MaleSex, nil
	}

	return FemaleSex, nil
}

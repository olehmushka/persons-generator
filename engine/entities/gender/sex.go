package gender

import pm "persons_generator/engine/probability_machine"

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
	list := []Sex{MaleSex, FemaleSex}
	i, err := pm.RandInt(len(list) - 1)
	if err != nil {
		return "", err
	}

	return list[i], nil
}

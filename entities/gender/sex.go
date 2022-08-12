package gender

import pm "persons_generator/probability_machine"

type Sex string

const (
	MaleSex   Sex = "male_sex"
	FemaleSex Sex = "female_sex"
)

func GetRandomSex() Sex {
	list := []Sex{MaleSex, FemaleSex}

	return list[pm.RandInt(len(list)-1)]
}

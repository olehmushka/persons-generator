package human

import (
	"fmt"

	"persons_generator/entities/culture"
	pm "persons_generator/probability_machine"
)

type Sex string

const (
	Male   Sex = "male"
	Female Sex = "female"
)

func (s Sex) ToCultureSex() culture.Sex {
	switch s {
	case Male:
		return culture.Male
	case Female:
		return culture.Female
	}
	panic(fmt.Sprintf("unexpected sex(sex=%s)", s))
}

var AllSexes = []Sex{Male, Female}

func GetRandomSex() Sex {
	return AllSexes[pm.RandInt(len(AllSexes)-1)]
}

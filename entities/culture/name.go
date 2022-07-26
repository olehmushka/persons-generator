package culture

import (
	bc "persons_generator/entities/base_culture"
	pm "persons_generator/probability_machine"
)

type Name struct {
	Culture       bc.Culture
	Name          string
	Sex           Sex
	Probability   float64
	IsInheritable bool
}

func GetRandomFirstName(sex Sex, c bc.Culture) *Name {
	names := GetAvailableFirstNamesByCulture(sex, c)
	n := pm.GetRandomFromSeveral(PrepareNameToChoose(names))

	return FindNameByName(n, names)
}

func GetRandomLastName(c bc.Culture) *Name {
	names := GetAvailableLastNamesByCulture(c)
	n := pm.GetRandomFromSeveral(PrepareNameToChoose(names))

	return FindNameByName(n, names)
}

func GetAvailableFirstNamesByCulture(sex Sex, c bc.Culture) []*Name {
	return []*Name{}
}

func GetAvailableLastNamesByCulture(c bc.Culture) []*Name {
	return []*Name{}
}

func FindAvailableNames(c *Culture, allNames []*Name) []*Name {
	names := make([]*Name, 0, len(allNames))
	for _, n := range allNames {
		if bc.IsEqualCulture(n.Culture, c) {
			names = append(names, n)
		}
	}

	return names
}

func FindNameByName(name string, allNames []*Name) *Name {
	for _, n := range allNames {
		if n.Name == name {
			return n
		}
	}

	return nil
}

func PrepareNameToChoose(names []*Name) map[string]float64 {
	m := make(map[string]float64, len(names))
	for _, name := range names {
		m[name.Name] = name.Probability
	}

	return m
}

func MergeNames(names ...[]*Name) []*Name {
	length := 0
	for _, ns := range names {
		length += len(ns)
	}
	merged := make([]*Name, length)
	var count int
	for _, ns := range names {
		for _, n := range ns {
			merged[count] = n
			count++
		}
	}

	return merged
}

func IsNamesEqual(n1, n2 *Name) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil {
		return false
	}
	if n2 == nil {
		return false
	}

	return n1.Name == n2.Name
}

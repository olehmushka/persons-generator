package proto_culture

import (
	"strings"

	pm "persons_generator/probability_machine"
)

type ProtoCulture struct {
	Name string
}

func (c *ProtoCulture) GetCultureName() string {
	if c == nil {
		return ""
	}

	return c.Name
}

func (c *ProtoCulture) HasLastName() bool {
	return false
}

func GetRandomProtoCultureFromList(list []*ProtoCulture) *ProtoCulture {
	return list[pm.RandInt(len(list)-1)]
}

func GetRandomProtoCulture() *ProtoCulture {
	return GetRandomProtoCultureFromList(AllProtoCultures)
}

func IsEqualProtoCulture(pc1, pc2 *ProtoCulture) bool {
	if pc1 == nil && pc2 == nil {
		return true
	}
	if pc1 == nil {
		return false
	}
	if pc2 == nil {
		return false
	}

	return pc1.Name == pc2.Name
}

func GetProtoCultureByPref(pref string) *ProtoCulture {
	for _, c := range AllProtoCultures {
		if strings.Contains(strings.ToLower(c.Name), strings.ToLower(pref)) {
			return c
		}
	}

	return nil
}

func GetProtoCulturesByPref(prefs []string) []*ProtoCulture {
	pcs := make([]*ProtoCulture, 0, len(prefs))
	for _, pref := range prefs {
		if pc := GetProtoCultureByPref(pref); pc != nil {
			pcs = append(pcs, pc)
		}
	}

	return Unique(pcs)
}

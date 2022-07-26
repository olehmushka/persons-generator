package proto_culture

import pm "persons_generator/probability_machine"

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

func GetRandomProtoCulture() *ProtoCulture {
	return AllProtoCultures[pm.RandInt(len(AllProtoCultures)-1)]
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

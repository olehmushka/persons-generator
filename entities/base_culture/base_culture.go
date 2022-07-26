package culture

import (
	pc "persons_generator/entities/proto_culture"
	pm "persons_generator/probability_machine"
)

type BaseCulture struct {
	Name string

	Proto             *pc.ProtoCulture
	InheritedCultures []*BaseCulture
}

func (c *BaseCulture) GetCultureName() string {
	if c == nil {
		return ""
	}

	return c.Name
}

func (c *BaseCulture) HasLastName() bool {
	if c == nil {
		return false
	}
	for _, cl := range CulturesWithoutLastName {
		if IsEqualCulture(cl, c) {
			return false
		}
	}

	return true
}

func GetBaseCultureByName(name string) *BaseCulture {
	for _, c := range AllBaseCultures {
		if c.Name == name {
			return c
		}
	}

	return nil
}

func GetBaseCulturesByProto(p *pc.ProtoCulture) []*BaseCulture {
	if p == nil {
		return nil
	}

	cultures := make([]*BaseCulture, 0, len(AllBaseCultures))
	for _, c := range AllBaseCultures {
		if c == nil {
			continue
		}
		if pc.IsEqualProtoCulture(c.Proto, p) {
			cultures = append(cultures, c)
		}
	}

	return cultures
}

func GetRandomBaseCultures(cultures []*BaseCulture, min, max int) []*BaseCulture {
	var (
		bcs    = PrepareBaseCultureToChoose(cultures, 0.1)
		l      = pm.RandIntInRange(min, max)
		cs     = pm.GetRandomSeveralFromSeveral(bcs, l)
		result = make([]*BaseCulture, l)
	)
	for i := range result {
		result[i] = GetBaseCultureByName(cs[i])
	}

	return result
}

func PrepareBaseCultureToChoose(bcs []*BaseCulture, prob float64) map[string]float64 {
	m := make(map[string]float64, len(bcs))
	for _, name := range bcs {
		m[name.Name] = prob
	}

	return m
}

func IsEqualCulture(pc1, pc2 *BaseCulture) bool {
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

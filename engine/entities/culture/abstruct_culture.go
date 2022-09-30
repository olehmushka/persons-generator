package culture

import (
	"persons_generator/core/tools"
	pm "persons_generator/engine/probability_machine"
)

type AbstructCulture struct {
	Name string `json:"name" bson:"name"`
	Root *Root  `json:"root" bson:"root"`
}

func (c *AbstructCulture) IsZero() bool {
	return c == nil
}

func (c *AbstructCulture) SerializeName() string {
	if c == nil {
		return ""
	}

	return c.Name
}

func getAbstractCulture(r *Root, proto []*Culture) (*AbstructCulture, error) {
	acs := make([]*AbstructCulture, 0, len(proto))
	for _, p := range proto {
		acs = append(acs, p.Abstuct)
	}

	out := make([]*AbstructCulture, 0, len(acs))
	for _, c := range UniqueAbstractCulture(acs) {
		if c.Root == r {
			out = append(out, c)
		}
	}

	return tools.RandomValueOfSlice(pm.RandFloat64, out)
}

func UniqueAbstractCulture(acs []*AbstructCulture) []*AbstructCulture {
	if len(acs) <= 1 {
		return acs
	}

	preOut := make(map[string]*AbstructCulture)
	for _, t := range acs {
		preOut[t.Name] = t
	}

	out := make([]*AbstructCulture, 0, len(preOut))
	for _, t := range preOut {
		out = append(out, t)
	}

	return out
}

func GetAbstructCultureSimilarityCoef(c1, c2 *AbstructCulture) float64 {
	if c1.IsZero() && c2.IsZero() {
		return 1
	}
	if c1.IsZero() || c2.IsZero() {
		return 0
	}

	if c1.Name == c2.Name {
		return 1
	}
	if IsRootEqual(c1.Root, c2.Root) {
		return 0.5
	}

	return 0
}

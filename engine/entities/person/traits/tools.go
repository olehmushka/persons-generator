package traits

func Unique(traits []*Trait) []*Trait {
	if len(traits) <= 1 {
		return traits
	}

	preOut := make(map[string]*Trait)
	for _, c := range traits {
		preOut[c.Name] = c
	}

	out := make([]*Trait, 0, len(preOut))
	for _, c := range preOut {
		out = append(out, c)
	}

	return out
}

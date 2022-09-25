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

func ExtractTraitNames(ts []*Trait) []string {
	if len(ts) == 0 {
		return []string{}
	}

	out := make([]string, len(ts))
	for i := range out {
		out[i] = ts[i].Name
	}

	return out
}

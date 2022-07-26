package proto_culture

import "sort"

func Unique(s []*ProtoCulture) []*ProtoCulture {
	if len(s) < 1 {
		return s
	}

	// sort
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].Name < s[j].Name
	})

	prev := 1
	for curr := 1; curr < len(s); curr++ {
		if s[curr-1].Name != s[curr].Name {
			s[prev] = s[curr]
			prev++
		}
	}

	return s[:prev]
}

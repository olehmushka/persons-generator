package tools

import "math/rand"

func Shuffle[T interface{}](in []T) []T {
	out := make([]T, len(in))
	for i, v := range rand.Perm(len(in)) {
		out[v] = in[i]
	}

	return out
}

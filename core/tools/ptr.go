package tools

func PointersToValues[T any](in []*T) []T {
	out := make([]T, len(in))
	for i := range out {
		out[i] = PointerToValue(in[i])
	}

	return out
}

func PointerToValue[T any](in *T) T {
	return *in
}

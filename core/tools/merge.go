package tools

func Merge[T interface{}](slOfEls ...[]T) []T {
	out := make([]T, 0)
	for _, els := range slOfEls {
		out = append(out, els...)
	}

	return out
}

func MergeUnique[T int | float64 | string](slOfEls ...[]T) []T {
	return Unique(Merge(slOfEls...))
}

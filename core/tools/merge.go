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

func MergeMaps(args ...map[string]float64) map[string]float64 {
	outSum := make(map[string]float64)
	outCount := make(map[string]float64)
	for _, arg := range args {
		for key, value := range arg {
			val, ok := outSum[key]
			if ok {
				outSum[key] = val + value
				outCount[key]++
			} else {
				outSum[key] = value
				outCount[key] = 1
			}
		}
	}
	out := make(map[string]float64, len(outSum))
	for key, sum := range outSum {
		out[key] = sum / outCount[key]
	}

	return out
}

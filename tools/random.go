package tools

import (
	"math"
)

func RandomValueOfSlice[T interface{}](randSrc func(float64) float64, sl []T) T {
	if len(sl) == 0 {
		var zero T
		return zero
	}
	return sl[int(math.Floor(randSrc(1)*float64(len(sl))))]
}

func RandomValuesOfSlice[T interface{}](randSrc func(float64) float64, sl []T, amount int) []T {
	if amount == 0 {
		return []T{}
	}
	if len(sl) <= amount {
		return sl
	}

	preOut := make(map[int]T)
	for {
		index := int(math.Floor(randSrc(1) * float64(len(sl))))
		preOut[index] = sl[index]
		if len(preOut) == amount {
			break
		}
	}

	out := make([]T, 0, amount)
	for _, v := range preOut {
		out = append(out, v)
	}

	return out
}

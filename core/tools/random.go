package tools

import (
	"math"
)

func RandomValueOfSlice[T interface{}](randSrc func(float64) (float64, error), sl []T) (T, error) {
	var zero T
	if len(sl) == 0 {
		return zero, nil
	}
	r, err := randSrc(1)
	if err != nil {
		return zero, err
	}

	return sl[int(math.Floor(r*float64(len(sl))))], nil
}

func RandomValuesOfSlice[T interface{}](randSrc func(float64) (float64, error), sl []T, amount int) ([]T, error) {
	if amount == 0 {
		return []T{}, nil
	}
	if len(sl) <= amount {
		return sl, nil
	}

	preOut := make(map[int]T)
	for {
		r, err := randSrc(1)
		if err != nil {
			return nil, err
		}
		index := int(math.Floor(r * float64(len(sl))))
		preOut[index] = sl[index]
		if len(preOut) == amount {
			break
		}
	}

	out := make([]T, 0, amount)
	for _, v := range preOut {
		out = append(out, v)
	}

	return out, nil
}

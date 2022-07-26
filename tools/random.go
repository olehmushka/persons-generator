package tools

import (
	"math"
)

func RandomValueOfSlice[T interface{}](randSrc func(float64) float64, sl []T) T {
	return sl[int(math.Floor(randSrc(1)*float64(len(sl))))]
}

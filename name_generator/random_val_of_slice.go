package name_generator

import (
	"math"

	pm "persons_generator/probability_machine"
)

// return random value from the slice
func RandomValOfSlice[T interface{}](sl []T) T {
	return sl[int(math.Floor(pm.RandFloat64(1)*float64(len(sl))))]
}

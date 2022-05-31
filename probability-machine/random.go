package probabilitymachine

import "math/rand"

func GetRandomBool(trueProb float64) bool {
	return GetRandomFloatInRange(0, 100) < trueProb
}

func GetRandomFloatInRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

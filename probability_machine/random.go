package probability_machine

import (
	"crypto/rand"
	"math/big"

	expRand "golang.org/x/exp/rand"
)

func GetRandomBool(trueProb float64) bool {
	n := RandFloat64(1)
	return n < trueProb
}

func RandIntInRange(min int, max int) int {
	if min >= max {
		panic("[RandIntInRange] min >= max")
	}

	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	return int(int64(min) + n.Int64())
}

func RandInt(max int) int {
	return RandIntInRange(0, max)
}

func RandFloat64InRange(min, max float64) float64 {
	if min >= max {
		panic("[RandFloat64InRange] min >= max")
	}

	s := expRand.NewSource(uint64(RandIntInRange(0, 100)))
	return min + expRand.New(s).Float64()*(max-min)
}

func RandFloat64(max float64) float64 {
	return RandFloat64InRange(0, max)
}

package probabilitymachine

import (
	"crypto/rand"
	"math/big"

	expRand "golang.org/x/exp/rand"
)

func GetRandomBool(trueProb float64) bool {
	n := randFloat64(0, 1)
	return n < trueProb
}

func randInt(min int, max int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	return int(int64(min) + n.Int64())
}

func randFloat64(min, max float64) float64 {
	s := expRand.NewSource(uint64(randInt(0, 100)))
	return min + expRand.New(s).Float64()*(max-min)
}

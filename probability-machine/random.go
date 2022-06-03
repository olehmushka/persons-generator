package probabilitymachine

import (
	"crypto/rand"
	"math/big"
)

// func init() {
// 	rand.Seed(time.Now().UTC().UnixNano())
// }

func GetRandomBool(trueProb int) bool {
	n := randInt(0, 100)
	return n < trueProb
}

func randInt(min int, max int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	return int(int64(min) + n.Int64())
}

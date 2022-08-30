package probability_machine

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"persons_generator/core/wrapped_error"

	expRand "golang.org/x/exp/rand"
)

func GetRandomBool(trueProb float64) (bool, error) {
	n, err := RandFloat64(1)
	if err != nil {
		return false, err
	}
	return n < trueProb, nil
}

func RandIntInRange(min int, max int) (int, error) {
	if min >= max {
		return 0, wrapped_error.New(http.StatusInternalServerError, nil, "[RandIntInRange] min >= max")
	}

	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	return int(int64(min) + n.Int64()), nil
}

func RandInt(max int) (int, error) {
	return RandIntInRange(0, max)
}

func RandFloat64InRange(min, max float64) (float64, error) {
	if min >= max {
		return 0, wrapped_error.New(http.StatusInternalServerError, nil, "[RandFloat64InRange] min >= max")
	}

	rInt, err := RandIntInRange(0, 100)
	if err != nil {
		return 0, err
	}
	s := expRand.NewSource(uint64(rInt))
	return min + expRand.New(s).Float64()*(max-min), nil
}

func RandFloat64(max float64) (float64, error) {
	return RandFloat64InRange(0, max)
}

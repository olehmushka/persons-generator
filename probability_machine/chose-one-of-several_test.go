package probability_machine

import "testing"

func TestGetRandomFromSeveral(t *testing.T) {
	m1 := map[string]float64{
		"one":   0.7,
		"two":   0.2,
		"three": 0.9,
		"four":  0.8,
		"five":  0.1,
		"six":   1,
	}
	t.Logf("map: %+v | chosen: %s\n", m1, GetRandomFromSeveral(m1))
}

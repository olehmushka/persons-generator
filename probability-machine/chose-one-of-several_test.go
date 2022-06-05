package probabilitymachine

import "testing"

func TestGetRandomFromSeveral(t *testing.T) {
	m1 := map[string]int{
		"one":   70,
		"two":   20,
		"three": 90,
		"four":  80,
		"five":  10,
		"six":   100,
	}
	t.Logf("map: %+v | chosen: %s\n", m1, GetRandomFromSeveral(m1))
}

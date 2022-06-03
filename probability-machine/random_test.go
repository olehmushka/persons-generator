package probabilitymachine

import "testing"

func TestGetRandomBool(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(GetRandomBool(i * 10))
	}
}

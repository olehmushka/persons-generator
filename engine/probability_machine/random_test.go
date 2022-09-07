package probability_machine

import (
	"fmt"
	"testing"
)

func TestGetRandomBool(t *testing.T) {
	tCases := []struct {
		name           string
		input          float64
		expectedOutput bool
		times          int
		skip           bool
	}{
		{
			name:           "should always get false",
			input:          0,
			expectedOutput: false,
			times:          10_000,
			skip:           false,
		},
		{
			name:           "should always get true",
			input:          1,
			expectedOutput: true,
			times:          10_000,
			skip:           false,
		},
		{
			name:           "should be skipped",
			input:          0.5,
			expectedOutput: true,
			times:          10,
			skip:           true,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			for i := 0; i < tc.times; i++ {
				output, _ := GetRandomBool(tc.input)
				if output != tc.expectedOutput && !tc.skip {
					tt.Errorf("unexpected output (expected=%t but actual=%t)", tc.expectedOutput, output)
				}
				if tc.skip {
					tt.Log(output)
				}
			}
		})
	}
}

func TestRandFloat64InRange(t *testing.T) {
	var (
		min = 0.1
		max = 0.4
	)

	t.Log(RandFloat64InRange(min, max))
}

func TestRandFloat64Norm(t *testing.T) {
	for i := 0; i < 10; i++ {
		val, err := RandFloat64Norm(1, 0)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%f\n", val)
	}
}

func TestRandFloat64NormInRange(t *testing.T) {
	for i := 0; i < 100; i++ {
		val, err := RandFloat64NormInRange(0, 1, 1, 0.1)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%f\n", val)
	}
}

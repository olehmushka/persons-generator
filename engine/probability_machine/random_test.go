package probability_machine

import "testing"

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

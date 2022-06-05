package probabilitymachine

import "testing"

func TestGetRandomBool(t *testing.T) {
	tCases := []struct {
		name           string
		input          int
		expectedOutput bool
		times          int
	}{
		{
			name:           "should always get false",
			input:          0,
			expectedOutput: false,
			times:          10_000,
		},
		{
			name:           "should always get true",
			input:          100,
			expectedOutput: true,
			times:          10_000,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			for i := 0; i < tc.times; i++ {
				if output := GetRandomBool(tc.input); output != tc.expectedOutput {
					t.Errorf("unexpected output (expected=%t but actual=%t)", tc.expectedOutput, output)
				}
			}
		})
	}
}

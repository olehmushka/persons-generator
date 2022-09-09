package tools

import "testing"

func TestMax(t *testing.T) {
	tCases := []struct {
		name     string
		args     []float64
		expected float64
	}{
		{
			name:     "should return 5",
			args:     []float64{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "should return 5",
			args:     []float64{0.1, 0.2, 0.3, 0.4, 0.5},
			expected: 0.5,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			if out := MaxNumeric(tc.args...); out != tc.expected {
				tt.Errorf("expected = %f, actual = %f", tc.expected, out)
			}
		})
	}
}

package language

import "testing"

func TestGetLanguageKinship(t *testing.T) {
	tCases := []struct {
		name     string
		l1, l2   *Language
		expected int
	}{
		{
			name:     "should return for belarusian & ukrainian",
			l1:       Belarusian,
			l2:       Ukrainian,
			expected: 1,
		},
		{
			name:     "should return for belarusian & russian",
			l1:       Russian,
			l2:       Ukrainian,
			expected: 2,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			if out := GetLanguageKinship(tc.l1, tc.l2); out != tc.expected {
				tt.Errorf("expected = %d, actual = %d", tc.expected, out)
			}
		})
	}
}

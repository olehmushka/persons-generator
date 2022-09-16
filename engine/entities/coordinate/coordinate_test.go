package coordinate

import "testing"

func TestCalcDistanceValue(t *testing.T) {
	tCases := []struct {
		name             string
		p1, p2           *Coordinate
		expectedDistance float64
	}{
		{
			p1: &Coordinate{
				X: 0,
				Y: 0,
			},
			p2: &Coordinate{
				X: 10,
				Y: 0,
			},
			expectedDistance: 10,
		},
		{
			p1: &Coordinate{
				X: 0,
				Y: 0,
			},
			p2: &Coordinate{
				X: 1,
				Y: 1,
			},
			expectedDistance: 1,
		},
		{
			p1: &Coordinate{
				X: 0,
				Y: 0,
			},
			p2: &Coordinate{
				X: 10,
				Y: 10,
			},
			expectedDistance: 15,
		},
		{
			p1: &Coordinate{
				X: 10,
				Y: 20,
			},
			p2: &Coordinate{
				X: 110,
				Y: 40,
			},
			expectedDistance: 102,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			distance := CalcDistanceValue(tc.p1, tc.p2)
			if distance != tc.expectedDistance {
				t.Logf("expected distance = %f but actual distance = %f", tc.expectedDistance, distance)
				t.Failed()
			}
		})
	}
}

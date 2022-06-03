package coordinate

import (
	"testing"

	"persons_generator/entities"
)

func TestCalcDistance(t *testing.T) {
	tCases := []struct {
		name             string
		p1, p2           *entities.Coordinate
		expectedDistance int
	}{
		{
			p1: &entities.Coordinate{
				X: 0,
				Y: 0,
			},
			p2: &entities.Coordinate{
				X: 10,
				Y: 0,
			},
			expectedDistance: 10,
		},
		{
			p1: &entities.Coordinate{
				X: 10,
				Y: 20,
			},
			p2: &entities.Coordinate{
				X: 110,
				Y: 40,
			},
			expectedDistance: 102,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			distance := CalcDistance(tc.p1, tc.p2)
			if distance != tc.expectedDistance {
				t.Logf("expected distance = %d but actual distance = %d", tc.expectedDistance, distance)
				t.Failed()
			}
		})
	}
}

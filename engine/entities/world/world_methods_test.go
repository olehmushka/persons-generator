package world

import (
	"testing"

	"persons_generator/engine/entities/coordinate"
	"persons_generator/engine/entities/location"
	"persons_generator/engine/entities/person"
)

func newTestWorld(grid [][]*location.Location) *World {
	return &World{
		Size:      len(grid),
		Locations: grid,
	}
}

func locWithPopulation(n int) *location.Location {
	pop := make([]*person.Person, n)
	for i := range pop {
		pop[i] = &person.Person{}
	}

	return &location.Location{
		Coordinate: &coordinate.Coordinate{},
		Population: pop,
	}
}

func TestGetPersons(t *testing.T) {
	tCases := []struct {
		name string
		grid [][]*location.Location
		want int
	}{
		{
			name: "all-nil locations",
			grid: [][]*location.Location{
				{nil, nil},
				{nil, nil},
			},
			want: 0,
		},
		{
			name: "nil mixed with populated",
			grid: [][]*location.Location{
				{locWithPopulation(3), nil},
				{nil, locWithPopulation(2)},
			},
			want: 5,
		},
		{
			name: "all populated, including an empty-population location",
			grid: [][]*location.Location{
				{locWithPopulation(4), locWithPopulation(0)},
				{locWithPopulation(1), locWithPopulation(2)},
			},
			want: 7,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			w := newTestWorld(tc.grid)
			if got := len(w.GetPersons()); got != tc.want {
				tt.Errorf("unexpected persons count (expected = %d, actual = %d)", tc.want, got)
			}
			if got := w.CalculatePersonsNumber(); got != tc.want {
				tt.Errorf("unexpected CalculatePersonsNumber (expected = %d, actual = %d)", tc.want, got)
			}
		})
	}
}

package culture

import (
	"testing"

	bc "persons_generator/entities/base_culture"
)

func TestMergeNames(t *testing.T) {
	tCases := []struct {
		name           string
		inputNames     [][]*Name
		expectedOutput []*Name
	}{
		{
			name: "5 length for 4 and 1 inputs",
			inputNames: [][]*Name{
				{
					{
						Culture:       bc.Mesopotamian,
						Name:          "Baal",
						Sex:           Male,
						IsInheritable: true,
					},
					{
						Culture:       bc.Mesopotamian,
						Name:          "Anatu",
						Sex:           Female,
						IsInheritable: true,
					},
					{
						Culture:       bc.Mesopotamian,
						Name:          "Lilith",
						Sex:           Female,
						IsInheritable: true,
					},
					{
						Culture:       bc.Mesopotamian,
						Name:          "Mot",
						Sex:           Female,
						IsInheritable: true,
					},
				},
				{
					{
						Culture:       bc.Mesopotamian,
						Name:          "Tiamat",
						Sex:           Female,
						IsInheritable: true,
					},
				},
			},
			expectedOutput: []*Name{
				{
					Culture:       bc.Mesopotamian,
					Name:          "Baal",
					Sex:           Male,
					IsInheritable: true,
				},
				{
					Culture:       bc.Mesopotamian,
					Name:          "Anatu",
					Sex:           Female,
					IsInheritable: true,
				},
				{
					Culture:       bc.Mesopotamian,
					Name:          "Lilith",
					Sex:           Female,
					IsInheritable: true,
				},
				{
					Culture:       bc.Mesopotamian,
					Name:          "Mot",
					Sex:           Female,
					IsInheritable: true,
				},
				{
					Culture:       bc.Mesopotamian,
					Name:          "Tiamat",
					Sex:           Female,
					IsInheritable: true,
				},
			},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			output := MergeNames(tc.inputNames...)
			if len(output) != len(tc.expectedOutput) {
				tt.Logf("unexpected output length (expected length = %d, actual length = %d)", len(tc.expectedOutput), len(output))
				return
			}
			for i, n := range output {
				if IsNamesEqual(n, tc.expectedOutput[i]) {
					tt.Logf("unexpected name (expected name = %+v, actual name = %+v)", tc.expectedOutput[i], n)
				}
			}
		})
	}
}

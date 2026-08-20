package http

import (
	"testing"
)

func intPtr(v int) *int {
	return &v
}

func TestGetAmounts(t *testing.T) {
	tCases := []struct {
		name                string
		personsAmount       *int
		malePersonsAmount   *int
		femalePersonsAmount *int
		wantAmount          int
		wantMaleAmount      int
		wantFemaleAmount    int
		wantErr             bool
	}{
		{
			name:             "all nil -> defaults to 100 (50/50)",
			wantAmount:       100,
			wantMaleAmount:   50,
			wantFemaleAmount: 50,
		},
		{
			name:             "only persons_amount, even",
			personsAmount:    intPtr(10),
			wantAmount:       10,
			wantMaleAmount:   5,
			wantFemaleAmount: 5,
		},
		{
			name:             "only persons_amount, odd (female gets the extra)",
			personsAmount:    intPtr(9),
			wantAmount:       9,
			wantMaleAmount:   4,
			wantFemaleAmount: 5,
		},
		{
			name:             "only persons_amount == 1",
			personsAmount:    intPtr(1),
			wantAmount:       1,
			wantMaleAmount:   1,
			wantFemaleAmount: 0,
		},
		{
			name:          "only persons_amount < 1 -> error",
			personsAmount: intPtr(0),
			wantErr:       true,
		},
		{
			name:              "persons_amount + male_persons_amount",
			personsAmount:     intPtr(10),
			malePersonsAmount: intPtr(4),
			wantAmount:        10,
			wantMaleAmount:    4,
			wantFemaleAmount:  6,
		},
		{
			name:              "persons_amount < male_persons_amount -> error",
			personsAmount:     intPtr(3),
			malePersonsAmount: intPtr(4),
			wantErr:           true,
		},
		{
			name:              "only male_persons_amount",
			malePersonsAmount: intPtr(7),
			wantAmount:        7,
			wantMaleAmount:    7,
			wantFemaleAmount:  0,
		},
		{
			name:              "only male_persons_amount < 1 -> error",
			malePersonsAmount: intPtr(0),
			wantErr:           true,
		},
		{
			name:                "male_persons_amount + female_persons_amount",
			malePersonsAmount:   intPtr(3),
			femalePersonsAmount: intPtr(5),
			wantAmount:          8,
			wantMaleAmount:      3,
			wantFemaleAmount:    5,
		},
		{
			name:                "male_persons_amount < 1 (with female set) -> error",
			malePersonsAmount:   intPtr(0),
			femalePersonsAmount: intPtr(5),
			wantErr:             true,
		},
		{
			name:                "female_persons_amount < 1 (with male set) -> error",
			malePersonsAmount:   intPtr(5),
			femalePersonsAmount: intPtr(0),
			wantErr:             true,
		},
		{
			name:                "persons_amount + female_persons_amount",
			personsAmount:       intPtr(10),
			femalePersonsAmount: intPtr(4),
			wantAmount:          10,
			wantMaleAmount:      6,
			wantFemaleAmount:    4,
		},
		{
			name:                "persons_amount < female_persons_amount -> error",
			personsAmount:       intPtr(3),
			femalePersonsAmount: intPtr(4),
			wantErr:             true,
		},
		{
			name:                "only female_persons_amount",
			femalePersonsAmount: intPtr(6),
			wantAmount:          6,
			wantMaleAmount:      0,
			wantFemaleAmount:    6,
		},
		{
			name:                "only female_persons_amount < 1 -> error",
			femalePersonsAmount: intPtr(0),
			wantErr:             true,
		},
		{
			name:                "all three, consistent",
			personsAmount:       intPtr(10),
			malePersonsAmount:   intPtr(4),
			femalePersonsAmount: intPtr(6),
			wantAmount:          10,
			wantMaleAmount:      4,
			wantFemaleAmount:    6,
		},
		{
			name:                "all three, inconsistent sum -> error",
			personsAmount:       intPtr(10),
			malePersonsAmount:   intPtr(4),
			femalePersonsAmount: intPtr(5),
			wantErr:             true,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			amount, maleAmount, femaleAmount, err := getAmounts(tc.personsAmount, tc.malePersonsAmount, tc.femalePersonsAmount)
			if tc.wantErr {
				if err == nil {
					tt.Fatal("expected an error, got <nil>")
				}
				return
			}
			if err != nil {
				tt.Fatalf("unexpected error: %+v", err)
			}
			if amount != tc.wantAmount {
				tt.Errorf("unexpected amount (expected = %d, actual = %d)", tc.wantAmount, amount)
			}
			if maleAmount != tc.wantMaleAmount {
				tt.Errorf("unexpected maleAmount (expected = %d, actual = %d)", tc.wantMaleAmount, maleAmount)
			}
			if femaleAmount != tc.wantFemaleAmount {
				tt.Errorf("unexpected femaleAmount (expected = %d, actual = %d)", tc.wantFemaleAmount, femaleAmount)
			}
		})
	}
}

func TestParseReligionCultureRelations(t *testing.T) {
	tCases := []struct {
		name    string
		in      []string
		want    map[string]string
		wantErr bool
	}{
		{
			name: "empty",
			in:   []string{},
			want: map[string]string{},
		},
		{
			name: "one valid pair",
			in:   []string{"culture-1:religion-1"},
			want: map[string]string{"religion-1": "culture-1"},
		},
		{
			name: "several valid pairs",
			in:   []string{"culture-1:religion-1", "culture-2:religion-2"},
			want: map[string]string{"religion-1": "culture-1", "religion-2": "culture-2"},
		},
		{
			name:    "malformed pair, no colon",
			in:      []string{"culture-1_religion-1"},
			wantErr: true,
		},
		{
			name:    "malformed pair, too many colons",
			in:      []string{"culture-1:religion-1:extra"},
			wantErr: true,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			got, err := parseReligionCultureRelations(tc.in)
			if tc.wantErr {
				if err == nil {
					tt.Fatal("expected an error, got <nil>")
				}
				return
			}
			if err != nil {
				tt.Fatalf("unexpected error: %+v", err)
			}
			if len(got) != len(tc.want) {
				tt.Fatalf("unexpected length (expected = %d, actual = %d)", len(tc.want), len(got))
			}
			for k, v := range tc.want {
				if got[k] != v {
					tt.Errorf("unexpected value for key %q (expected = %s, actual = %s)", k, v, got[k])
				}
			}
		})
	}
}

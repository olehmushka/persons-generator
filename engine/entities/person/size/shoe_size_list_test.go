package size

import (
	"persons_generator/engine/entities/gender"
	"testing"
)

func TestGetSizeByCm(t *testing.T) {
	tCases := []struct {
		name string
		sex  gender.Sex
		size float64
		out  *ShoeSize
	}{
		{
			name: "should found 39 size by 24.5 cm",
			sex:  gender.MaleSex,
			size: 24.5,
			out:  GetSizeByEUSize(39, AllMaleShoeSizes),
		},
		{
			name: "should found 40 size by 25 cm",
			sex:  gender.MaleSex,
			size: 25,
			out:  GetSizeByEUSize(40, AllMaleShoeSizes),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			out := GetSizeByCm(tc.sex, tc.size)
			if out == nil && tc.out == nil {
				return
			}
			if out == nil || tc.out == nil {
				tt.Errorf("unextected out (expected out = %+v, actual out = %+v)", tc.out, out)
				return
			}

			if out.EUSize != tc.out.EUSize {
				tt.Errorf("unextected out (expected out = %+v, actual out = %+v)", tc.out, out)
			}
		})
	}
}

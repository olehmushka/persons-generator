package size

import (
	"math"
	"persons_generator/core/tools"
	"persons_generator/engine/entities/gender"
)

var AllFemaleShoeSizes = []ShoeSize{
	{
		USSize:         5,
		EUSize:         35,
		UKSize:         3,
		FeetLengthInch: 8.6,
		FeetLengthCm:   21.8,
		Sex:            gender.FemaleSex,
	},
	{
		USSize:         6,
		EUSize:         36,
		UKSize:         4,
		FeetLengthInch: 8.9,
		FeetLengthCm:   22.5,
		Sex:            gender.FemaleSex,
	},
	{
		USSize:         7,
		EUSize:         37,
		UKSize:         5,
		FeetLengthInch: 9.1,
		FeetLengthCm:   23.1,
		Sex:            gender.FemaleSex,
	},
	{
		USSize:         8,
		EUSize:         38,
		UKSize:         6,
		FeetLengthInch: 9.4,
		FeetLengthCm:   23.8,
		Sex:            gender.FemaleSex,
	},
	{
		USSize:         9,
		EUSize:         39,
		UKSize:         7,
		FeetLengthInch: 9.6,
		FeetLengthCm:   24.4,
		Sex:            gender.FemaleSex,
	},
	{
		USSize:         10,
		EUSize:         40,
		UKSize:         8,
		FeetLengthInch: 9.9,
		FeetLengthCm:   25.1,
		Sex:            gender.FemaleSex,
	},
	{
		USSize:         11,
		EUSize:         41,
		UKSize:         9,
		FeetLengthInch: 10.1,
		FeetLengthCm:   25.7,
		Sex:            gender.FemaleSex,
	},
	{
		USSize:         12,
		EUSize:         42,
		UKSize:         10,
		FeetLengthInch: 10.4,
		FeetLengthCm:   26.4,
		Sex:            gender.FemaleSex,
	},
}

var AllMaleShoeSizes = []ShoeSize{
	{
		USSize:         6,
		EUSize:         39,
		UKSize:         6,
		FeetLengthInch: 9.6,
		FeetLengthCm:   24.4,
		Sex:            gender.MaleSex,
	},
	{
		USSize:         7,
		EUSize:         40,
		UKSize:         6.5,
		FeetLengthInch: 9.9,
		FeetLengthCm:   25.1,
		Sex:            gender.MaleSex,
	},
	{
		USSize:         8,
		EUSize:         41,
		UKSize:         7.5,
		FeetLengthInch: 10.1,
		FeetLengthCm:   25.7,
		Sex:            gender.MaleSex,
	},
	{
		USSize:         9,
		EUSize:         42,
		UKSize:         8.5,
		FeetLengthInch: 10.4,
		FeetLengthCm:   26.4,
		Sex:            gender.MaleSex,
	},
	{
		USSize:         10,
		EUSize:         43,
		UKSize:         9.5,
		FeetLengthInch: 10.6,
		FeetLengthCm:   27,
		Sex:            gender.MaleSex,
	},
	{
		USSize:         11,
		EUSize:         44,
		UKSize:         10.5,
		FeetLengthInch: 10.9,
		FeetLengthCm:   27.7,
		Sex:            gender.MaleSex,
	},
	{
		USSize:         12,
		EUSize:         45,
		UKSize:         11.5,
		FeetLengthInch: 11.2,
		FeetLengthCm:   28.3,
		Sex:            gender.MaleSex,
	},
	{
		USSize:         13,
		EUSize:         46,
		UKSize:         12.5,
		FeetLengthInch: 11.4,
		FeetLengthCm:   29,
		Sex:            gender.MaleSex,
	},
	{
		USSize:         14,
		EUSize:         47,
		UKSize:         13.5,
		FeetLengthInch: 11.7,
		FeetLengthCm:   29.6,
		Sex:            gender.MaleSex,
	},
	{
		USSize:         15,
		EUSize:         48,
		UKSize:         14.5,
		FeetLengthInch: 11.9,
		FeetLengthCm:   30.3,
		Sex:            gender.MaleSex,
	},
}

func getMinShoeSizeInCm(sizes []ShoeSize) float64 {
	var min float64
	for i, size := range sizes {
		if i == 0 {
			min = size.FeetLengthCm
		}

		if min > size.FeetLengthCm {
			min = size.FeetLengthCm
		}
	}

	return min
}

func getMaxShoeSizeInCm(sizes []ShoeSize) float64 {
	var max float64
	for i, size := range sizes {
		if i == 0 {
			max = size.FeetLengthCm
		}

		if max < size.FeetLengthCm {
			max = size.FeetLengthCm
		}
	}

	return max
}

func GetMinShoeSizeInCm(sex gender.Sex) float64 {
	switch {
	case sex.IsMale():
		return getMinShoeSizeInCm(AllFemaleShoeSizes)
	case sex.IsFemale():
		return getMinShoeSizeInCm(AllMaleShoeSizes)
	}
	return 0
}

func GetMaxShoeSizeInCm(sex gender.Sex) float64 {
	switch {
	case sex.IsMale():
		return getMaxShoeSizeInCm(AllMaleShoeSizes)
	case sex.IsFemale():
		return getMaxShoeSizeInCm(AllFemaleShoeSizes)
	}
	return 0
}

func prepareShoeSizeParams(sex gender.Sex, min, medium, max float64) (float64, float64, float64) {
	var (
		absMin = GetMinShoeSizeInCm(sex)
		absMax = GetMaxShoeSizeInCm(sex)
	)
	if min > absMin {
		min = absMin
	}
	if max > absMax {
		max = absMax
	}
	if absMin < medium && medium > absMax {
		medium = (min + max) / 2
	}

	return min, medium, max
}

func GetSizeByCm(sex gender.Sex, size float64) *ShoeSize {
	switch {
	case sex.IsMale():
		out := getSizeByCmFromSizes(size, AllMaleShoeSizes)
		return &out
	case sex.IsFemale():
		out := getSizeByCmFromSizes(size, AllFemaleShoeSizes)
		return &out
	}
	return nil
}

func getSizeByCmFromSizes(size float64, sizes []ShoeSize) ShoeSize {
	m := make(map[int]float64)
	for _, s := range sizes {
		m[s.EUSize] = math.Abs(s.FeetLengthCm - size)
	}
	var (
		minDiff   float64 = math.MaxFloat64
		minEUSize int
	)
	for euSize, diff := range m {
		if diff < minDiff {
			minDiff = diff
			minEUSize = euSize
		}
	}

	return tools.Search(
		sizes,
		func(e ShoeSize) int { return e.EUSize },
		minEUSize,
	)
}

func GetSizeByEUSize(euSize int, sizes []ShoeSize) *ShoeSize {
	for _, s := range sizes {
		if s.EUSize == euSize {
			return &s
		}
	}

	return nil
}

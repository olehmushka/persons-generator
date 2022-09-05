package hair

var AllHairDensities = []HairDensity{
	LowHairDensity,
	MediumHairDensity,
	HighHairDensity,
}

var (
	LowHairDensity = HairDensity{
		Type: "low",
	}
	MediumHairDensity = HairDensity{
		Type: "medium",
	}
	HighHairDensity = HairDensity{
		Type: "high",
	}
)

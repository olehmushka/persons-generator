package hair

var AllHairTextures = []HairTexture{
	BlackHairTexture,
	CurlyHairTexture,
	FineHairTexture,
	StraightHairTexture,
	ThickHairTexture,
	ThinHairTexture,
	WavyHairTexture,
}

var (
	BlackHairTexture = HairTexture{
		Type: "black",
	}
	CurlyHairTexture = HairTexture{
		Type: "curly",
	}
	FineHairTexture = HairTexture{
		Type: "fine",
	}
	StraightHairTexture = HairTexture{
		Type: "straight",
	}
	ThickHairTexture = HairTexture{
		Type: "thick",
	}
	ThinHairTexture = HairTexture{
		Type: "thin",
	}
	WavyHairTexture = HairTexture{
		Type: "wavy",
	}
)

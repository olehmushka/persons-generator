package color

var AllSkinColors = []Color{
	// Olive
	RustyNailSkinColor,
	CopperSkinColor,
	RawSiennaSkinColor,
	PorscheSkinColor,
	CherokeeSkinColor,
	FleshSkinColor,
	// Brown
	MetallicBronzeSkinColor,
	OldCopperSkinColor,
	SpicyMixSkinColor,
	PottersClaySkinColor,
	MuddyWatersSkinColor,
	AntiqueBrassSkinColor,
	WhiskeySkinColor,
	// Red
	CopperCanyonSkinColor,
	PaarlSkinColor,
	RedDamaskSkinColor,
	ApricotSkinColor,
	TumbleweedSkinColor,
	SandyBrownSkinColor,
	HitPinkSkinColor,
	// Pink
	ManhattanSkinColor,
	RomanticSkinColor,
	WheatSkinColor,
	// Caucasian
	TacaoSkinColor,
	ApricotPeachSkinColor,
}

var (
	// Olive
	RustyNailSkinColor = Color{
		Name:    "rusty_nail",
		Hex:     "#89530c",
		Palette: OliveSkinColorPalette,
	}
	CopperSkinColor = Color{
		Name:    "copper",
		Hex:     "#a76432",
		Palette: OliveSkinColorPalette,
	}
	RawSiennaSkinColor = Color{
		Name:    "raw_sienna",
		Hex:     "#d0914e",
		Palette: OliveSkinColorPalette,
	}
	PorscheSkinColor = Color{
		Name:    "porsche",
		Hex:     "#e9b36d",
		Palette: OliveSkinColorPalette,
	}
	CherokeeSkinColor = Color{
		Name:    "cherokee",
		Hex:     "#f9c286",
		Palette: OliveSkinColorPalette,
	}
	FleshSkinColor = Color{
		Name:    "flesh",
		Hex:     "#ffd1a2",
		Palette: OliveSkinColorPalette,
	}
	// Brown
	MetallicBronzeSkinColor = Color{
		Name:    "metallic_bronze",
		Hex:     "#56351f",
		Palette: BrownSkinColorPalette,
	}
	OldCopperSkinColor = Color{
		Name:    "old_copper",
		Hex:     "#774f35",
		Palette: BrownSkinColorPalette,
	}
	SpicyMixSkinColor = Color{
		Name:    "spicy_mix",
		Hex:     "#8f5e46",
		Palette: BrownSkinColorPalette,
	}
	PottersClaySkinColor = Color{
		Name:    "potters_clay",
		Hex:     "#925534",
		Palette: BrownSkinColorPalette,
	}
	MuddyWatersSkinColor = Color{
		Name:    "muddy_waters",
		Hex:     "#bb7c53",
		Palette: BrownSkinColorPalette,
	}
	AntiqueBrassSkinColor = Color{
		Name:    "antique_brass",
		Hex:     "#c68a63",
		Palette: BrownSkinColorPalette,
	}
	WhiskeySkinColor = Color{
		Name:    "whiskey",
		Hex:     "#d3965f",
		Palette: BrownSkinColorPalette,
	}
	// Red
	CopperCanyonSkinColor = Color{
		Name:    "copper_canyon",
		Hex:     "#823a11",
		Palette: RedSkinColorPalette,
	}
	PaarlSkinColor = Color{
		Name:    "paarl",
		Hex:     "#9f5227",
		Palette: RedSkinColorPalette,
	}
	RedDamaskSkinColor = Color{
		Name:    "red_damask",
		Hex:     "#db764c",
		Palette: RedSkinColorPalette,
	}
	ApricotSkinColor = Color{
		Name:    "apricot",
		Hex:     "#e98961",
		Palette: RedSkinColorPalette,
	}
	TumbleweedSkinColor = Color{
		Name:    "tumbleweed",
		Hex:     "#daa07a",
		Palette: RedSkinColorPalette,
	}
	SandyBrownSkinColor = Color{
		Name:    "sandy_brown",
		Hex:     "#f49f68",
		Palette: RedSkinColorPalette,
	}
	HitPinkSkinColor = Color{
		Name:    "hit_pink",
		Hex:     "#ffb07c",
		Palette: RedSkinColorPalette,
	}
	// Pink
	ManhattanSkinColor = Color{
		Name:    "manhattan",
		Hex:     "#f5b99d",
		Palette: PinkSkinColorPalette,
	}
	RomanticSkinColor = Color{
		Name:    "romantic",
		Hex:     "#fecab2",
		Palette: PinkSkinColorPalette,
	}
	WheatSkinColor = Color{
		Name:    "wheat",
		Hex:     "#f6d7c5",
		Palette: PinkSkinColorPalette,
	}
	// Caucasian
	TacaoSkinColor = Color{
		Name:    "tacao",
		Hex:     "#f0bc97",
		Palette: CaucasianSkinColorPalette,
	}
	ApricotPeachSkinColor = Color{
		Name:    "apricot_peach",
		Hex:     "#f9cfb7",
		Palette: CaucasianSkinColorPalette,
	}
)

func GetSkinColorsByPalette(palette string) []Color {
	return getColorsByPalette(palette, AllSkinColors)
}

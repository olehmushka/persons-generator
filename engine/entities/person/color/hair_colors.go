package color

var AllHairColors = []Color{
	// Black
	BlackRussianHairColor,
	CocoaBrownHairColor,
	SepiaBlackHairColor,
	// MediumBrown
	EclipseHairColor,
	TamarindHairColor,
	CedarHairColor,
	DeepOakHairColor,
	VanCleefHairColor,
	CioccolatoHairColor,
	HairyHeathHairColor,
	// LightBrown
	WalnutHairColor,
	RussetHairColor,
	BullShotHairColor,
	PaarlHairColor,
	CopperHairColor,

	// Blonde
	SepiaSkinHairColor,
	CapePalliserHairColor,
	DriftwoodHairColor,
	TwineHairColor,
	TanHairColor,
	TumbleweedHairColor,
	GoldSandHairColor,
	LightApricotHairColor,
	// DarkRed
	ChocolateHairColor,
	MahoganyHairColor,
	LonestarHairColor,
	RedwoodHairColor,
	PuebloHairColor,
	TamarilloHairColor,
	RoofTerracottaHairColor,
	TabascoHairColor,
	// StrawberryBlonde
	TuscanyHairColor,
	RedDamaskHairColor,
	BurntSiennaHairColor,
	SandyBrownHairColor,
	HitPinkHairColor,
	WaxFlowerHairColor,
}

var (
	// Black
	BlackRussianHairColor = Color{
		Name:    "black_russian",
		Hex:     "#00041b",
		Palette: BlackHairColorPalette,
	}
	CocoaBrownHairColor = Color{
		Name:    "cocoa_brown",
		Hex:     "#1e1614",
		Palette: BlackHairColorPalette,
	}
	SepiaBlackHairColor = Color{
		Name:    "sepia_black",
		Hex:     "#2d0406",
		Palette: BlackHairColorPalette,
	}
	// MediumBrown
	EclipseHairColor = Color{
		Name:    "eclipse",
		Hex:     "#321f18",
		Palette: MediumBrownHairColorPalette,
	}
	TamarindHairColor = Color{
		Name:    "tamarind",
		Hex:     "#362016",
		Palette: MediumBrownHairColorPalette,
	}
	CedarHairColor = Color{
		Name:    "cedar",
		Hex:     "#3c2013",
		Palette: MediumBrownHairColorPalette,
	}
	DeepOakHairColor = Color{
		Name:    "deep_oak",
		Hex:     "#432210",
		Palette: MediumBrownHairColorPalette,
	}
	VanCleefHairColor = Color{
		Name:    "van_cleef",
		Hex:     "#4f220e",
		Palette: MediumBrownHairColorPalette,
	}
	CioccolatoHairColor = Color{
		Name:    "cioccolato",
		Hex:     "#58260c",
		Palette: MediumBrownHairColorPalette,
	}
	HairyHeathHairColor = Color{
		Name:    "hairy_heath",
		Hex:     "#622d11",
		Palette: MediumBrownHairColorPalette,
	}
	// LightBrown
	WalnutHairColor = Color{
		Name:    "walnut",
		Hex:     "#713f17",
		Palette: LightBrownHairColorPalette,
	}
	RussetHairColor = Color{
		Name:    "russet",
		Hex:     "#7e461a",
		Palette: LightBrownHairColorPalette,
	}
	BullShotHairColor = Color{
		Name:    "bull_shot",
		Hex:     "#8e4f23",
		Palette: LightBrownHairColorPalette,
	}
	PaarlHairColor = Color{
		Name:    "paarl",
		Hex:     "#9e5a25",
		Palette: LightBrownHairColorPalette,
	}
	CopperHairColor = Color{
		Name:    "copper",
		Hex:     "#b36d37",
		Palette: LightBrownHairColorPalette,
	}
	// Blonde
	SepiaSkinHairColor = Color{
		Name:    "sepia_skin",
		Hex:     "#9a683d",
		Palette: BlondeHairColorPalette,
	}
	CapePalliserHairColor = Color{
		Name:    "cape_palliser",
		Hex:     "#a6754b",
		Palette: BlondeHairColorPalette,
	}
	DriftwoodHairColor = Color{
		Name:    "driftwood",
		Hex:     "#a87d49",
		Palette: BlondeHairColorPalette,
	}
	TwineHairColor = Color{
		Name:    "twine",
		Hex:     "#c19569",
		Palette: BlondeHairColorPalette,
	}
	TanHairColor = Color{
		Name:    "tan",
		Hex:     "#d6b082",
		Palette: BlondeHairColorPalette,
	}
	TumbleweedHairColor = Color{
		Name:    "tumbleweed",
		Hex:     "#d9af85",
		Palette: BlondeHairColorPalette,
	}
	GoldSandHairColor = Color{
		Name:    "gold_sand",
		Hex:     "#e7c397",
		Palette: BlondeHairColorPalette,
	}
	LightApricotHairColor = Color{
		Name:    "light_apricot",
		Hex:     "#fcd9ad",
		Palette: BlondeHairColorPalette,
	}
	// DarkRed
	ChocolateHairColor = Color{
		Name:    "chocolate",
		Hex:     "#430102",
		Palette: DarkRedHairColorPalette,
	}
	MahoganyHairColor = Color{
		Name:    "mahogany",
		Hex:     "#631b09",
		Palette: DarkRedHairColorPalette,
	}
	LonestarHairColor = Color{
		Name:    "lonestar",
		Hex:     "#690101",
		Palette: DarkRedHairColorPalette,
	}
	RedwoodHairColor = Color{
		Name:    "redwood",
		Hex:     "#6c280d",
		Palette: DarkRedHairColorPalette,
	}
	PuebloHairColor = Color{
		Name:    "pueblo",
		Hex:     "#792412",
		Palette: DarkRedHairColorPalette,
	}
	TamarilloHairColor = Color{
		Name:    "tamarillo",
		Hex:     "#810e0e",
		Palette: DarkRedHairColorPalette,
	}
	RoofTerracottaHairColor = Color{
		Name:    "roof_terracotta",
		Hex:     "#a61b1b",
		Palette: DarkRedHairColorPalette,
	}
	TabascoHairColor = Color{
		Name:    "tabasco",
		Hex:     "#923612",
		Palette: DarkRedHairColorPalette,
	}
	// StrawberryBlonde
	TuscanyHairColor = Color{
		Name:    "tuscany",
		Hex:     "#ba542b",
		Palette: StrawberryBlondeHairColorPalette,
	}
	RedDamaskHairColor = Color{
		Name:    "red_damask",
		Hex:     "#de764c",
		Palette: StrawberryBlondeHairColorPalette,
	}
	BurntSiennaHairColor = Color{
		Name:    "burnt_sienna",
		Hex:     "#e8855e",
		Palette: StrawberryBlondeHairColorPalette,
	}
	SandyBrownHairColor = Color{
		Name:    "sandy_brown",
		Hex:     "#f29570",
		Palette: StrawberryBlondeHairColorPalette,
	}
	HitPinkHairColor = Color{
		Name:    "hit_pink",
		Hex:     "#ffa580",
		Palette: StrawberryBlondeHairColorPalette,
	}
	WaxFlowerHairColor = Color{
		Name:    "wax_flower",
		Hex:     "#ffba9e",
		Palette: StrawberryBlondeHairColorPalette,
	}
)

func GetHairColorsByPalette(palette string) []Color {
	return getColorsByPalette(palette, AllHairColors)
}

package color

var AllEyesColors = []Color{
	// Brown
	ColaEyesColor,
	IndianTanEyesColor,
	BrownBrambleEyesColor,
	CafeRoyaleEyesColor,
	RawUmberEyesColor,
	// Hazel
	PungaEyesColor,
	DallasEyesColor,
	PottersClayEyesColor,
	CapePalliserEyesColor,
	// Amber
	KormaEyesColor,
	LightKormaEyesColor,
	DesertEyesColor,
	KumeraEyesColor,
	LuxorGoldEyesColor,
	MarigoldEyesColor,
	// Green
	MadrasEyesColor,
	VerdunGreenEyesColor,
	AntiqueBronzeEyesColor,
	SaratogaEyesColor,
	HimalayaEyesColor,
	// Grey
	FinchEyesColor,
	GoBenEyesColor,
	GoBenLightEyesColor,
	FlaxSmokeEyesColor,
	WillowGroveEyesColor,
	CamouflageGreenEyesColor,
	SpanishGreenEyesColor,
	PewterEyesColor,
	NorwayEyesColor,
	GumLeafEyesColor,
	// TealBlue
	GreenPeaDarkEyesColor,
	GreenPeaEyesColor,
	GreenPeaLightEyesColor,
	AmazonEyesColor,
	ViridianEyesColor,
	PatinaEyesColor,
	SeaNymphEyesColor,
	// Cyan
	MountainMeadowEyesColor,
	JungleGreenEyesColor,
	KeppelEyesColor,
	FountainBlueEyesColor,
	DownyEyesColor,
	MonteCarloEyesColor,
	SinbadEyesColor,
	ZigguratEyesColor,
	// Indigo
	PortGoreEyesColor,
	RhinoEyesColor,
	ChambrayEyesColor,
	VictoriaEyesColor,
	BlueVioletEyesColor,
	DanubeEyesColor,
	ChetwodeBlueEyesColor,
	// Violet
	StudioEyesColor,
	// Red
	CardinalEyesColor,
	RazzmatazzEyesColor,
}

var (
	// Brown
	ColaEyesColor = Color{
		Name:    "cola",
		Hex:     "#3a1e00",
		Palette: BrownEyesColorPalette,
	}
	IndianTanEyesColor = Color{
		Name:    "indian_tan",
		Hex:     "#4e2a02",
		Palette: BrownEyesColorPalette,
	}
	BrownBrambleEyesColor = Color{
		Name:    "brown_bramble",
		Hex:     "#5e3407",
		Palette: BrownEyesColorPalette,
	}
	CafeRoyaleEyesColor = Color{
		Name:    "cafe_royale",
		Hex:     "#6b3e0d",
		Palette: BrownEyesColorPalette,
	}
	RawUmberEyesColor = Color{
		Name:    "raw_umber",
		Hex:     "#7a4b13",
		Palette: BrownEyesColorPalette,
	}

	// Hazel
	PungaEyesColor = Color{
		Name:    "punga",
		Hex:     "#523913",
		Palette: HazelEyesColorPalette,
	}
	DallasEyesColor = Color{
		Name:    "dallas",
		Hex:     "#6e5226",
		Palette: HazelEyesColorPalette,
	}
	PottersClayEyesColor = Color{
		Name:    "potters_clay",
		Hex:     "#836332",
		Palette: HazelEyesColorPalette,
	}
	CapePalliserEyesColor = Color{
		Name:    "cape_palliser",
		Hex:     "#957645",
		Palette: HazelEyesColorPalette,
	}

	// Amber
	KormaEyesColor = Color{
		Name:    "korma",
		Hex:     "#81460E",
		Palette: AmberEyesColorPalette,
	}
	LightKormaEyesColor = Color{
		Name:    "light_korma",
		Hex:     "#945211",
		Palette: AmberEyesColorPalette,
	}
	DesertEyesColor = Color{
		Name:    "desert",
		Hex:     "#aa631d",
		Palette: AmberEyesColorPalette,
	}
	KumeraEyesColor = Color{
		Name:    "kumera",
		Hex:     "#8a6121",
		Palette: AmberEyesColorPalette,
	}
	LuxorGoldEyesColor = Color{
		Name:    "luxor_gold",
		Hex:     "#a07126",
		Palette: AmberEyesColorPalette,
	}
	MarigoldEyesColor = Color{
		Name:    "marigold",
		Hex:     "#b97f2c",
		Palette: AmberEyesColorPalette,
	}

	// Green
	MadrasEyesColor = Color{
		Name:    "madras",
		Hex:     "#373400",
		Palette: GreenEyesColorPalette,
	}
	VerdunGreenEyesColor = Color{
		Name:    "verdun_green",
		Hex:     "#464200",
		Palette: GreenEyesColorPalette,
	}
	AntiqueBronzeEyesColor = Color{
		Name:    "antique_bronze",
		Hex:     "#565107",
		Palette: GreenEyesColorPalette,
	}
	SaratogaEyesColor = Color{
		Name:    "saratoga",
		Hex:     "#62590f",
		Palette: GreenEyesColorPalette,
	}
	HimalayaEyesColor = Color{
		Name:    "himalaya",
		Hex:     "#746a19",
		Palette: GreenEyesColorPalette,
	}

	// Grey
	FinchEyesColor = Color{
		Name:    "finch",
		Hex:     "#5c5942",
		Palette: GreyEyesColorPalette,
	}
	GoBenEyesColor = Color{
		Name:    "go_ben",
		Hex:     "#706c4c",
		Palette: GreyEyesColorPalette,
	}
	GoBenLightEyesColor = Color{
		Name:    "go_ben_light",
		Hex:     "#6e724a",
		Palette: GreyEyesColorPalette,
	}
	FlaxSmokeEyesColor = Color{
		Name:    "flax_smoke",
		Hex:     "#76815b",
		Palette: GreyEyesColorPalette,
	}
	WillowGroveEyesColor = Color{
		Name:    "willow_grove",
		Hex:     "#677d5f",
		Palette: GreyEyesColorPalette,
	}
	CamouflageGreenEyesColor = Color{
		Name:    "camouflage_green",
		Hex:     "#7b8676",
		Palette: GreyEyesColorPalette,
	}
	SpanishGreenEyesColor = Color{
		Name:    "spanish_green",
		Hex:     "#869787",
		Palette: GreyEyesColorPalette,
	}
	PewterEyesColor = Color{
		Name:    "pewter",
		Hex:     "#90a79c",
		Palette: GreyEyesColorPalette,
	}
	NorwayEyesColor = Color{
		Name:    "norway",
		Hex:     "#9fb29f",
		Palette: GreyEyesColorPalette,
	}
	GumLeafEyesColor = Color{
		Name:    "gum_leaf",
		Hex:     "#b3ccc0",
		Palette: GreyEyesColorPalette,
	}

	// TealBlue
	GreenPeaDarkEyesColor = Color{
		Name:    "green_pea_dark",
		Hex:     "#124738",
		Palette: TealBlueEyesColorPalette,
	}
	GreenPeaEyesColor = Color{
		Name:    "green_pea",
		Hex:     "#195644",
		Palette: TealBlueEyesColorPalette,
	}
	GreenPeaLightEyesColor = Color{
		Name:    "green_pea_light",
		Hex:     "#236350",
		Palette: TealBlueEyesColorPalette,
	}
	AmazonEyesColor = Color{
		Name:    "amazon",
		Hex:     "#327461",
		Palette: TealBlueEyesColorPalette,
	}
	ViridianEyesColor = Color{
		Name:    "viridian",
		Hex:     "#437e6d",
		Palette: TealBlueEyesColorPalette,
	}
	PatinaEyesColor = Color{
		Name:    "patina",
		Hex:     "#5f9384",
		Palette: TealBlueEyesColorPalette,
	}
	SeaNymphEyesColor = Color{
		Name:    "sea_nymph",
		Hex:     "#799e93",
		Palette: TealBlueEyesColorPalette,
	}

	// Cyan
	MountainMeadowEyesColor = Color{
		Name:    "mountain_meadow",
		Hex:     "#1b978d",
		Palette: CyanEyesColorPalette,
	}
	JungleGreenEyesColor = Color{
		Name:    "jungle_green",
		Hex:     "#27a79d",
		Palette: CyanEyesColorPalette,
	}
	KeppelEyesColor = Color{
		Name:    "keppel",
		Hex:     "#39b2a8",
		Palette: CyanEyesColorPalette,
	}
	FountainBlueEyesColor = Color{
		Name:    "fountainblue",
		Hex:     "#50c2b9",
		Palette: CyanEyesColorPalette,
	}
	DownyEyesColor = Color{
		Name:    "downy",
		Hex:     "#67c9c1",
		Palette: CyanEyesColorPalette,
	}
	MonteCarloEyesColor = Color{
		Name:    "monte_carlo",
		Hex:     "#7ed1ca",
		Palette: CyanEyesColorPalette,
	}
	SinbadEyesColor = Color{
		Name:    "sinbad",
		Hex:     "#99d5d0",
		Palette: CyanEyesColorPalette,
	}
	ZigguratEyesColor = Color{
		Name:    "ziggurat",
		Hex:     "#b7dcd9",
		Palette: CyanEyesColorPalette,
	}

	// Indigo
	PortGoreEyesColor = Color{
		Name:    "port_gore",
		Hex:     "#282a5a",
		Palette: IndigoEyesColorPalette,
	}
	RhinoEyesColor = Color{
		Name:    "rhino",
		Hex:     "#2d3067",
		Palette: IndigoEyesColorPalette,
	}
	ChambrayEyesColor = Color{
		Name:    "chambray",
		Hex:     "#364385",
		Palette: IndigoEyesColorPalette,
	}
	VictoriaEyesColor = Color{
		Name:    "cictoria",
		Hex:     "#475496",
		Palette: IndigoEyesColorPalette,
	}
	BlueVioletEyesColor = Color{
		Name:    "blue_violet",
		Hex:     "#5666b6",
		Palette: IndigoEyesColorPalette,
	}
	DanubeEyesColor = Color{
		Name:    "danube",
		Hex:     "#6376d3",
		Palette: IndigoEyesColorPalette,
	}
	ChetwodeBlueEyesColor = Color{
		Name:    "chetwode_blue",
		Hex:     "#7182d9",
		Palette: IndigoEyesColorPalette,
	}

	// Violet
	StudioEyesColor = Color{
		Name:    "studio",
		Hex:     "#8c48b9",
		Palette: VioletEyesColorPalette,
	}

	// Red
	CardinalEyesColor = Color{
		Name:    "cardinal",
		Hex:     "#ad1446",
		Palette: RedEyesColorPalette,
	}
	RazzmatazzEyesColor = Color{
		Name:    "razzmatazz",
		Hex:     "#d60366",
		Palette: RedEyesColorPalette,
	}
)

func GetEyesColorsByPalette(palette string) []Color {
	return getColorsByPalette(palette, AllEyesColors)
}

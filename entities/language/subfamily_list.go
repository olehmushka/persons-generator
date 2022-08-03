package language

import "persons_generator/tools"

var AllSubfamilies = tools.Merge(
	AllIndoEuropeanSubfamilies,
	[]*Subfamily{BasqueSubfamily, JaponicSubfamily, KoreanicSubfamily, KartoZanSubfamily, SvanSubfamily},
	AllUralicSubfamilies,
	AllTurkicSubfamilies,
	AllMongolicSubfamilies,
	AllTungusicSubfamilies,
	AllNigerCongoSubfamilies,
	AllAustronesianSubfamilies,
	AllSinoTibetanSubfamilies,
	AllAfroAsiaticSubfamilies,
	AllNiloSaharanSubfamilies,
	AllOtoMangueanSubfamilies,
	AllUtoAztecanSubfamilies,
	AllAustroAsiaticSubfamilies,
	AllVieticSubfamilies,
	AllTaiKadaiSubfamilies,
	AllDravidianSubfamilies,
	AllTupianSubfamilies,
	AllEskimoAleutSubfamilies,
	AllQuechuaSubfamilies,
	AllFantasySubfamilies,
)

// IndoEuropeanFamily
var (
	AlbanianSubfamily = &Subfamily{
		Name:     "albanian_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	AnatolianSubfamily = &Subfamily{
		Name:     "anatolian_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	ArmenianSubfamily = &Subfamily{
		Name:     "armenian_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	BaltoSlavicSubfamily = &Subfamily{
		Name:     "balto_slavic_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	CelticSubfamily = &Subfamily{
		Name:     "celtic_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	DacianSubfamily = &Subfamily{
		Name:     "dacian_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	GermanicSubfamily = &Subfamily{
		Name:     "germanic_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	GreekSubfamily = &Subfamily{
		Name:     "greek_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	IllyrianSubfamily = &Subfamily{
		Name:     "illyrian_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	IndoIranianSubfamily = &Subfamily{
		Name:     "indo_iranian_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	ItalicSubfamily = &Subfamily{
		Name:     "italic_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	LiburnianSubfamily = &Subfamily{
		Name:     "liburnian_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	LusitanianSubfamily = &Subfamily{
		Name:     "lusitanian_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	MessapicSubfamily = &Subfamily{
		Name:     "messapic_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	PhrygianSubfamily = &Subfamily{
		Name:     "phrygian_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	ThracianSubfamily = &Subfamily{
		Name:     "thracian_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	TocharianSubfamily = &Subfamily{
		Name:     "tocharian_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	TysenianSubfamily = &Subfamily{
		Name:     "tysenian_lang_subfamily",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
)

// IndoEuropeanFamilyBranches
var (
	// Slavic
	SlavicSubfamily = &Subfamily{
		Name:              "slavic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: BaltoSlavicSubfamily,
		IsLiving:          true,
	}
	RuthenianSubfamily = &Subfamily{
		Name:              "ruthenian_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: SlavicSubfamily,
		IsLiving:          true,
	}
	MoscovianSubfamily = &Subfamily{
		Name:              "moscovian_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: SlavicSubfamily,
		IsLiving:          true,
	}
	SouthSlavicSubfamily = &Subfamily{
		Name:              "south_slavic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: SlavicSubfamily,
		IsLiving:          true,
	}
	EasternSouthSlavicSubfamily = &Subfamily{
		Name:              "eastern_south_slavic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: SouthSlavicSubfamily,
		IsLiving:          true,
	}
	WesternSouthSlavicSubfamily = &Subfamily{
		Name:              "western_south_slavic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: SouthSlavicSubfamily,
		IsLiving:          true,
	}
	WestSlavicSubfamily = &Subfamily{
		Name:              "west_slavic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: SlavicSubfamily,
		IsLiving:          true,
	}
	LechiticSubfamily = &Subfamily{
		Name:              "lechitic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: WestSlavicSubfamily,
		IsLiving:          true,
	}
	SorbianSubfamily = &Subfamily{
		Name:              "sorbian_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: WestSlavicSubfamily,
		IsLiving:          true,
	}
	CzechSlovakSubfamily = &Subfamily{
		Name:              "czech_slovak_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: WestSlavicSubfamily,
		IsLiving:          true,
	}
	BalticSubfamily = &Subfamily{
		Name:              "baltic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: BaltoSlavicSubfamily,
		IsLiving:          true,
	}
	EasternBalticSubfamily = &Subfamily{
		Name:              "eastern_baltic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: BalticSubfamily,
		IsLiving:          true,
	}
	// Celtic
	BrittonicSubfamily = &Subfamily{
		Name:              "brittonic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: CelticSubfamily,
		IsLiving:          true,
	}
	WesternBrittonicSubfamily = &Subfamily{
		Name:              "western_brittonic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: BrittonicSubfamily,
		IsLiving:          true,
	}
	SouthWesternBrittonicSubfamily = &Subfamily{
		Name:              "south_western_brittonic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: BrittonicSubfamily,
		IsLiving:          true,
	}
	GoidelicSubfamily = &Subfamily{
		Name:              "goidelic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: CelticSubfamily,
		IsLiving:          true,
	}
	// Germanic
	NorthGermanicSubfamily = &Subfamily{
		Name:              "north_germanic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: GermanicSubfamily,
		IsLiving:          true,
	}
	WestGermanicSubfamily = &Subfamily{
		Name:              "west_germanic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: GermanicSubfamily,
		IsLiving:          true,
	}
	EastGermanicSubfamily = &Subfamily{
		Name:              "east_germanic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: GermanicSubfamily,
		IsLiving:          false,
	}
	ElbeGermanicSubfamily = &Subfamily{
		Name:              "elbe_germanic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: GermanicSubfamily,
		IsLiving:          true,
	}
	NorthSeaGermanicSubfamily = &Subfamily{
		Name:              "north_sea_germanic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: GermanicSubfamily,
		IsLiving:          true,
	}
	WeserRhineGermanicSubfamily = &Subfamily{
		Name:              "weser_rhine_germanic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: GermanicSubfamily,
		IsLiving:          true,
	}
	// IndoIranian
	IndoAryanSubfamily = &Subfamily{
		Name:              "indo_aryan_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: IndoIranianSubfamily,
		IsLiving:          true,
	}
	IranianSubfamily = &Subfamily{
		Name:              "iranian_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: IndoIranianSubfamily,
		IsLiving:          true,
	}
	NuriastaniSubfamily = &Subfamily{
		Name:              "nuriastani_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: IndoIranianSubfamily,
		IsLiving:          true,
	}
	// Italic
	LatinoFaliscanSubfamily = &Subfamily{
		Name:              "latino_faliscan_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: ItalicSubfamily,
		IsLiving:          true,
	}
	OscoUmbrianSubfamily = &Subfamily{
		Name:              "osco_umbrian_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: ItalicSubfamily,
		IsLiving:          false,
	}
	RomanceSubfamily = &Subfamily{
		Name:              "romance_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: LatinoFaliscanSubfamily,
		IsLiving:          true,
	}
	IberoRomanceSubfamily = &Subfamily{
		Name:              "ibero_romance_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: RomanceSubfamily,
		IsLiving:          true,
	}
	OccitanoRomanceSubfamily = &Subfamily{
		Name:              "occitano_romance_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: RomanceSubfamily,
		IsLiving:          true,
	}
	GalloRomanceSubfamily = &Subfamily{
		Name:              "gallo_romance_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: RomanceSubfamily,
		IsLiving:          true,
	}
	RhaetoRomanceSubfamily = &Subfamily{
		Name:              "rhaeto_romance_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: RomanceSubfamily,
		IsLiving:          true,
	}
	GalloItalicSubfamily = &Subfamily{
		Name:              "gallo_italic_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: RomanceSubfamily,
		IsLiving:          true,
	}
	ItaloDalmatianSubfamily = &Subfamily{
		Name:              "italo_dalmatian_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: RomanceSubfamily,
		IsLiving:          true,
	}
	EasternRomanceSubfamily = &Subfamily{
		Name:              "eastern_romance_lang_subfamily",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: RomanceSubfamily,
		IsLiving:          true,
	}
)

var AllIndoEuropeanSubfamilies = []*Subfamily{
	AlbanianSubfamily,
	AnatolianSubfamily,
	ArmenianSubfamily,
	BaltoSlavicSubfamily,
	CelticSubfamily,
	DacianSubfamily,
	GermanicSubfamily,
	GreekSubfamily,
	IllyrianSubfamily,
	IndoIranianSubfamily,
	ItalicSubfamily,
	LiburnianSubfamily,
	LusitanianSubfamily,
	MessapicSubfamily,
	PhrygianSubfamily,
	ThracianSubfamily,
	TocharianSubfamily,
	TysenianSubfamily,

	SlavicSubfamily,
	RuthenianSubfamily,
	MoscovianSubfamily,
	SouthSlavicSubfamily,
	EasternSouthSlavicSubfamily,
	WesternSouthSlavicSubfamily,
	WestSlavicSubfamily,
	LechiticSubfamily,
	SorbianSubfamily,
	CzechSlovakSubfamily,
	BalticSubfamily,
	EasternBalticSubfamily,
	// Celtic
	BrittonicSubfamily,
	WesternBrittonicSubfamily,
	SouthWesternBrittonicSubfamily,
	GoidelicSubfamily,
	// Germanic
	NorthGermanicSubfamily,
	WestGermanicSubfamily,
	EastGermanicSubfamily,
	ElbeGermanicSubfamily,
	NorthSeaGermanicSubfamily,
	WeserRhineGermanicSubfamily,
	// IndoIranian
	IndoAryanSubfamily,
	IranianSubfamily,
	NuriastaniSubfamily,
	// Italic
	LatinoFaliscanSubfamily,
	OscoUmbrianSubfamily,
	RomanceSubfamily,
	IberoRomanceSubfamily,
	OccitanoRomanceSubfamily,
	GalloRomanceSubfamily,
	RhaetoRomanceSubfamily,
	GalloItalicSubfamily,
	ItaloDalmatianSubfamily,
	EasternRomanceSubfamily,
}

// Basque
var (
	BasqueSubfamily = &Subfamily{
		Name:     "basque_lang_subfamily",
		Family:   BasqueFamily,
		IsLiving: true,
	}
)

// Japonic
var (
	JaponicSubfamily = &Subfamily{
		Name:     "japonic_lang_subfamily",
		Family:   JaponicFamily,
		IsLiving: true,
	}
)

// Koreanic
var (
	KoreanicSubfamily = &Subfamily{
		Name:     "koreanic_lang_subfamily",
		Family:   KoreanicFamily,
		IsLiving: true,
	}
)

// Kartvelian
var (
	KartoZanSubfamily = &Subfamily{
		Name:     "karto_zan_lang_subfamily",
		Family:   KartvelianFamily,
		IsLiving: true,
	}
	SvanSubfamily = &Subfamily{
		Name:     "svan_lang_subfamily",
		Family:   KartvelianFamily,
		IsLiving: true,
	}
)

// Uralic
var (
	FinnoUgricSubfamily = &Subfamily{
		Name:     "finno_ugric_lang_subfamily",
		Family:   UralicFamily,
		IsLiving: true,
	}
	SamoyedicSubfamily = &Subfamily{
		Name:     "samoyedic_lang_subfamily",
		Family:   UralicFamily,
		IsLiving: true,
	}
	FinnoPermicSubfamily = &Subfamily{
		Name:              "finno_permic_lang_subfamily",
		Family:            UralicFamily,
		ExtendedSubfamily: FinnoUgricSubfamily,
		IsLiving:          true,
	}
	UgricSubfamily = &Subfamily{
		Name:              "ugric_lang_subfamily",
		Family:            UralicFamily,
		ExtendedSubfamily: FinnoUgricSubfamily,
		IsLiving:          true,
	}
	PermicSubfamily = &Subfamily{
		Name:              "Permic_lang_subfamily",
		Family:            UralicFamily,
		ExtendedSubfamily: FinnoPermicSubfamily,
		IsLiving:          true,
	}
	BaltoFinnicSubfamily = &Subfamily{
		Name:              "BaltoFinnic_lang_subfamily",
		Family:            UralicFamily,
		ExtendedSubfamily: FinnoPermicSubfamily,
		IsLiving:          true,
	}
	SamiSubfamily = &Subfamily{
		Name:              "Sami_lang_subfamily",
		Family:            UralicFamily,
		ExtendedSubfamily: FinnoPermicSubfamily,
		IsLiving:          true,
	}
	MordvinSubfamily = &Subfamily{
		Name:              "Mordvin_lang_subfamily",
		Family:            UralicFamily,
		ExtendedSubfamily: FinnoPermicSubfamily,
		IsLiving:          true,
	}
)

var AllUralicSubfamilies = []*Subfamily{
	FinnoUgricSubfamily,
	SamoyedicSubfamily,
	FinnoPermicSubfamily,
	UgricSubfamily,
	PermicSubfamily,
	BaltoFinnicSubfamily,
	SamiSubfamily,
	MordvinSubfamily,
}

// Turkic
var (
	ShazTurkicSubfamily = &Subfamily{
		Name:     "shaz_turkic_lang_subfamily",
		Family:   TurkicFamily,
		IsLiving: true,
	}
	LirTurkicSubfamily = &Subfamily{
		Name:     "lir_turkic_lang_subfamily",
		Family:   TurkicFamily,
		IsLiving: true,
	}
	OghuzSubfamily = &Subfamily{
		Name:              "oghuz_lang_subfamily",
		Family:            TurkicFamily,
		ExtendedSubfamily: ShazTurkicSubfamily,
		IsLiving:          true,
	}
	KipchakSubfamily = &Subfamily{
		Name:              "kipchak_lang_subfamily",
		Family:            TurkicFamily,
		ExtendedSubfamily: ShazTurkicSubfamily,
		IsLiving:          true,
	}
	KarlukSubfamily = &Subfamily{
		Name:              "karluk_lang_subfamily",
		Family:            TurkicFamily,
		ExtendedSubfamily: ShazTurkicSubfamily,
		IsLiving:          true,
	}
	SiberianTurkicSubfamily = &Subfamily{
		Name:              "siberian_turkic_lang_subfamily",
		Family:            TurkicFamily,
		ExtendedSubfamily: ShazTurkicSubfamily,
		IsLiving:          true,
	}
)

var AllTurkicSubfamilies = []*Subfamily{
	ShazTurkicSubfamily,
	LirTurkicSubfamily,
	OghuzSubfamily,
	KipchakSubfamily,
	KarlukSubfamily,
	SiberianTurkicSubfamily,
}

// Mongolic
var (
	DagurSubfamily = &Subfamily{
		Name:     "dagur_lang_subfamily",
		Family:   MongolicFamily,
		IsLiving: true,
	}
	CentralMongolicSubfamily = &Subfamily{
		Name:     "central_mongolic_lang_subfamily",
		Family:   MongolicFamily,
		IsLiving: true,
	}
	SouthernMongolicSubfamily = &Subfamily{
		Name:     "southern_mongolic_lang_subfamily",
		Family:   MongolicFamily,
		IsLiving: true,
	}
	MogholSubfamily = &Subfamily{
		Name:     "moghol_lang_subfamily",
		Family:   MongolicFamily,
		IsLiving: true,
	}
	KhitanSubfamily = &Subfamily{
		Name:     "khitan_lang_subfamily",
		Family:   MongolicFamily,
		IsLiving: true,
	}
)

var AllMongolicSubfamilies = []*Subfamily{DagurSubfamily, CentralMongolicSubfamily, SouthernMongolicSubfamily, MogholSubfamily, KhitanSubfamily}

// Tungusic

var (
	JurchenicNanaicSubfamily = &Subfamily{
		Name:     "jurchenic_nanaic_lang_subfamily",
		Family:   TungusicFamily,
		IsLiving: true,
	}
	UdegheicSubfamily = &Subfamily{
		Name:     "udegheic_lang_subfamily",
		Family:   TungusicFamily,
		IsLiving: true,
	}

	JurchenicSubfamily = &Subfamily{
		Name:              "jurchenic_lang_subfamily",
		Family:            TungusicFamily,
		ExtendedSubfamily: JurchenicNanaicSubfamily,
		IsLiving:          true,
	}
	NanaicSubfamily = &Subfamily{
		Name:              "nanaic_lang_subfamily",
		Family:            TungusicFamily,
		ExtendedSubfamily: JurchenicNanaicSubfamily,
		IsLiving:          true,
	}
)

var AllTungusicSubfamilies = []*Subfamily{JurchenicNanaicSubfamily, UdegheicSubfamily, JurchenicSubfamily, NanaicSubfamily}

// NigerCongoFamily
var (
	DogonSubfamily = &Subfamily{
		Name:     "dogon_lang_subfamily",
		Family:   NigerCongoFamily,
		IsLiving: false,
	}
	MandeSubfamily = &Subfamily{
		Name:     "mande_lang_subfamily",
		Family:   NigerCongoFamily,
		IsLiving: false,
	}
	IjoidSubfamily = &Subfamily{
		Name:     "ijoid_lang_subfamily",
		Family:   NigerCongoFamily,
		IsLiving: false,
	}
	LafofaSubfamily = &Subfamily{
		Name:     "lafofa (kordofanian)_lang_subfamily",
		Family:   NigerCongoFamily,
		IsLiving: false,
	}
	KruSubfamily = &Subfamily{
		Name:     "kru_lang_subfamily",
		Family:   NigerCongoFamily,
		IsLiving: false,
	}
	SiamouSubfamily = &Subfamily{
		Name:     "siamou_lang_subfamily",
		Family:   NigerCongoFamily,
		IsLiving: false,
	}
	AtlanticCongoSubfamily = &Subfamily{
		Name:     "atlantic_congo_lang_subfamily",
		Family:   NigerCongoFamily,
		IsLiving: true,
	}
	VoltaCongoSubfamily = &Subfamily{
		Name:              "volta_congo_lang_subfamily",
		Family:            NigerCongoFamily,
		ExtendedSubfamily: AtlanticCongoSubfamily,
		IsLiving:          true,
	}
	BantuSubfamily = &Subfamily{
		Name:              "bantu_lang_subfamily",
		Family:            NigerCongoFamily,
		ExtendedSubfamily: AtlanticCongoSubfamily,
		IsLiving:          true,
	}
)

var AllNigerCongoSubfamilies = []*Subfamily{
	DogonSubfamily,
	MandeSubfamily,
	IjoidSubfamily,
	LafofaSubfamily,
	KruSubfamily,
	SiamouSubfamily,
	AtlanticCongoSubfamily,
	VoltaCongoSubfamily,
	BantuSubfamily,
}

// AustronesianFamily
var (
	RukaiSubfamily = &Subfamily{
		Name:     "rukai_lang_subfamily",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	TsouicSubfamily = &Subfamily{
		Name:     "tsouic_lang_subfamily",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	PuyumaSubfamily = &Subfamily{
		Name:     "puyuma_lang_subfamily",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	NorthwestFormosanSubfamily = &Subfamily{
		Name:     "northwest_formosan_lang_subfamily",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	EastFormosanSubfamily = &Subfamily{
		Name:     "east_formosan_lang_subfamily",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	WesternPlainsSubfamily = &Subfamily{
		Name:     "western_plains_lang_subfamily",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	AtayalicSubfamily = &Subfamily{
		Name:     "atayalic_lang_subfamily",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	BununSubfamily = &Subfamily{
		Name:     "bunun_lang_subfamily",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	PaiwanSubfamily = &Subfamily{
		Name:     "paiwan_lang_subfamily",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	MalayoPolynesianSubfamily = &Subfamily{
		Name:     "malayo_polynesian_lang_subfamily",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	OceanicSubfamily = &Subfamily{
		Name:              "oceanic_lang_subfamily",
		Family:            AustronesianFamily,
		ExtendedSubfamily: MalayoPolynesianSubfamily,
		IsLiving:          true,
	}
	PolynesianSubfamily = &Subfamily{
		Name:              "polynesian_lang_subfamily",
		Family:            AustronesianFamily,
		ExtendedSubfamily: OceanicSubfamily,
		IsLiving:          true,
	}
	MarquesicSubfamily = &Subfamily{
		Name:              "marquesic_lang_subfamily",
		Family:            AustronesianFamily,
		ExtendedSubfamily: PolynesianSubfamily,
		IsLiving:          true,
	}
)

var AllAustronesianSubfamilies = []*Subfamily{
	RukaiSubfamily,
	TsouicSubfamily,
	PuyumaSubfamily,
	NorthwestFormosanSubfamily,
	EastFormosanSubfamily,
	WesternPlainsSubfamily,
	AtayalicSubfamily,
	BununSubfamily,
	PaiwanSubfamily,
	MalayoPolynesianSubfamily,
	OceanicSubfamily,
	PolynesianSubfamily,
	MarquesicSubfamily,
}

// SinoTibetanFamily
var (
	SiniticSubfamily = &Subfamily{
		Name:     "sinitic_lang_subfamily",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	LoloBurmeseSubfamily = &Subfamily{
		Name:     "lolo_burmese_lang_subfamily",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	TibeticSubfamily = &Subfamily{
		Name:     "tibetic_lang_subfamily",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	KarenicSubfamily = &Subfamily{
		Name:     "karenic_lang_subfamily",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	BodoGaroSubfamily = &Subfamily{
		Name:     "bodo_garo_lang_subfamily",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	KukiChinSubfamily = &Subfamily{
		Name:     "kuki_chin_lang_subfamily",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	MeiteiSubfamily = &Subfamily{
		Name:     "meitei_lang_subfamily",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	TamangicSubfamily = &Subfamily{
		Name:     "tamangic_lang_subfamily",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	BaiSubfamily = &Subfamily{
		Name:     "bai_lang_subfamily",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	JingphoLuishSubfamily = &Subfamily{
		Name:     "jingpho_luish_lang_subfamily",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
)

// SinoTibetanFamilyBranches
var (
	// Chinese
	ChineseSubfamily = &Subfamily{
		Name:              "chinese_languages_lang_subfamily",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: SiniticSubfamily,
		IsLiving:          true,
	}
	GreaterBaiSubfamily = &Subfamily{
		Name:              "greater_bai_lang_subfamily",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: SiniticSubfamily,
		IsLiving:          true,
	}
	// LoloBurmese
	MondzishSubfamily = &Subfamily{
		Name:              "mondzish_lang_subfamily",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	BurmishSubfamily = &Subfamily{
		Name:              "burmish_lang_subfamily",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	HanoishSubfamily = &Subfamily{
		Name:              "hanoish_lang_subfamily",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	LahoishSubfamily = &Subfamily{
		Name:              "lahoish_lang_subfamily",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	NaxishSubfamily = &Subfamily{
		Name:              "naxish_lang_subfamily",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	NusoishSubfamily = &Subfamily{
		Name:              "nusoish_lang_subfamily",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	KazhuoishSubfamily = &Subfamily{
		Name:              "kazhuoish_lang_subfamily",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	LisoishSubfamily = &Subfamily{
		Name:              "lisoish_lang_subfamily",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	NisoishSubfamily = &Subfamily{
		Name:              "nisoish_lang_subfamily",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	QiangicSubfamily = &Subfamily{
		Name:              "qiangic_lang_subfamily",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	// Tibetic
	TibetoBurmanSubfamily = &Subfamily{
		Name:              "tibeto_burman_lang_subfamily",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: TibeticSubfamily,
		IsLiving:          true,
	}
	BodishSubfamily = &Subfamily{
		Name:              "bodish_lang_subfamily",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: TibetoBurmanSubfamily,
		IsLiving:          true,
	}
)

var AllSinoTibetanSubfamilies = []*Subfamily{
	SiniticSubfamily,
	LoloBurmeseSubfamily,
	TibeticSubfamily,
	KarenicSubfamily,
	BodoGaroSubfamily,
	KukiChinSubfamily,
	MeiteiSubfamily,
	TamangicSubfamily,
	BaiSubfamily,
	JingphoLuishSubfamily,
	// Chinese
	ChineseSubfamily,
	GreaterBaiSubfamily,
	// LoloBurmese
	MondzishSubfamily,
	BurmishSubfamily,
	HanoishSubfamily,
	LahoishSubfamily,
	NaxishSubfamily,
	NusoishSubfamily,
	KazhuoishSubfamily,
	LisoishSubfamily,
	NisoishSubfamily,
	QiangicSubfamily,
	// Tibetic
	TibetoBurmanSubfamily,
	BodishSubfamily,
}

// AfroAsiaticFamily
var (
	BerberSubfamily = &Subfamily{
		Name:     "berber_lang_subfamily",
		Family:   AfroAsiaticFamily,
		IsLiving: true,
	}
	ChadicSubfamily = &Subfamily{
		Name:     "chadic_lang_subfamily",
		Family:   AfroAsiaticFamily,
		IsLiving: true,
	}
	CushiticSubfamily = &Subfamily{
		Name:     "cushitic_lang_subfamily",
		Family:   AfroAsiaticFamily,
		IsLiving: true,
	}
	EgyptianSubfamily = &Subfamily{
		Name:     "egyptian_lang_subfamily",
		Family:   AfroAsiaticFamily,
		IsLiving: true,
	}
	SemiticSubfamily = &Subfamily{
		Name:     "semitic_lang_subfamily",
		Family:   AfroAsiaticFamily,
		IsLiving: true,
	}
	OmoticSubfamily = &Subfamily{
		Name:     "omotic_lang_subfamily",
		Family:   AfroAsiaticFamily,
		IsLiving: true,
	}

	AgawSubfamily = &Subfamily{
		Name:              "agaw_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: CushiticSubfamily,
		IsLiving:          true,
	}
	DullaySubfamily = &Subfamily{
		Name:              "dullay_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: CushiticSubfamily,
		IsLiving:          true,
	}
	HighlandEastCushiticSubfamily = &Subfamily{
		Name:              "highland_east_cushitic_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: CushiticSubfamily,
		IsLiving:          true,
	}
	LowlandEastCushiticSubfamily = &Subfamily{
		Name:              "lowland_east_cushitic_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: CushiticSubfamily,
		IsLiving:          true,
	}
	SouthCushiticSubfamily = &Subfamily{
		Name:              "south_cushitic_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: CushiticSubfamily,
		IsLiving:          true,
	}
	EastSemiticSubfamily = &Subfamily{
		Name:              "east_semitic_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: SemiticSubfamily,
		IsLiving:          true,
	}
	WestSemiticSubfamily = &Subfamily{
		Name:              "west_semitic_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: SemiticSubfamily,
		IsLiving:          true,
	}
	CanaaniteSubfamily = &Subfamily{
		Name:              "canaanite_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: WestSemiticSubfamily,
		IsLiving:          true,
	}
	AramaicSubfamily = &Subfamily{
		Name:              "aramaic_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: WestSemiticSubfamily,
		IsLiving:          true,
	}
	ArabicSubfamily = &Subfamily{
		Name:              "arabic_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: WestSemiticSubfamily,
		IsLiving:          true,
	}
	EthiopicSubfamily = &Subfamily{
		Name:              "ethiopic_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: WestSemiticSubfamily,
		IsLiving:          true,
	}
	SouthArabianSubfamily = &Subfamily{
		Name:              "modern_south_arabian_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: WestSemiticSubfamily,
		IsLiving:          true,
	}
	HebrewSubfamily = &Subfamily{
		Name:              "hebrew_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: CanaaniteSubfamily,
		IsLiving:          true,
	}
	WestChadicSubfamily = &Subfamily{
		Name:              "west_chadic_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: ChadicSubfamily,
		IsLiving:          true,
	}
	SomalicSubfamily = &Subfamily{
		Name:              "somalic_lang_subfamily",
		Family:            AfroAsiaticFamily,
		ExtendedSubfamily: LowlandEastCushiticSubfamily,
		IsLiving:          true,
	}
)

var AllAfroAsiaticSubfamilies = []*Subfamily{
	BerberSubfamily,
	ChadicSubfamily,
	CushiticSubfamily,
	EgyptianSubfamily,
	SemiticSubfamily,
	OmoticSubfamily,

	AgawSubfamily,
	DullaySubfamily,
	HighlandEastCushiticSubfamily,
	LowlandEastCushiticSubfamily,
	SouthCushiticSubfamily,
	EastSemiticSubfamily,
	WestSemiticSubfamily,
	CanaaniteSubfamily,
	AramaicSubfamily,
	ArabicSubfamily,
	EthiopicSubfamily,
	SouthArabianSubfamily,
	HebrewSubfamily,
	WestChadicSubfamily,
	SomalicSubfamily,
}

// NiloSaharanFamily
var (
	BertaSubfamily = &Subfamily{
		Name:     "berta_lang_subfamily",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	Baga = &Subfamily{
		Name:     "b'aga_lang_subfamily",
		Family:   NiloSaharanFamily,
		IsLiving: false,
	}
	FurSubfamily = &Subfamily{
		Name:     "fur_lang_subfamily",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	KaduSubfamily = &Subfamily{
		Name:     "kadu_lang_subfamily",
		Family:   NiloSaharanFamily,
		IsLiving: false,
	}
	KomanSubfamily = &Subfamily{
		Name:     "koman_lang_subfamily",
		Family:   NiloSaharanFamily,
		IsLiving: false,
	}
	KuliakSubfamily = &Subfamily{
		Name:     "kuliak_lang_subfamily",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	KunamaSubfamily = &Subfamily{
		Name:     "kunama_lang_subfamily",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	MabanSubfamily = &Subfamily{
		Name:     "maban_lang_subfamily",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	SaharanSubfamily = &Subfamily{
		Name:     "saharan_lang_subfamily",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	SonghaySubfamily = &Subfamily{
		Name:     "songhay_lang_subfamily",
		Family:   NiloSaharanFamily,
		IsLiving: false,
	}
	CentralSudanicSubfamily = &Subfamily{
		Name:     "central_sudanic_lang_subfamily",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	EasternSudanicSubfamily = &Subfamily{
		Name:     "eastern_sudanic_lang_subfamily",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	MimiDSubfamily = &Subfamily{
		Name:     "mimi_d_lang_subfamily",
		Family:   NiloSaharanFamily,
		IsLiving: false,
	}

	NubianSubfamily = &Subfamily{
		Name:              "nubian_lang_subfamily",
		Family:            NiloSaharanFamily,
		ExtendedSubfamily: EasternSudanicSubfamily,
		IsLiving:          true,
	}
	DajuSubfamily = &Subfamily{
		Name:              "daju_lang_subfamily",
		Family:            NiloSaharanFamily,
		ExtendedSubfamily: EasternSudanicSubfamily,
		IsLiving:          true,
	}
)

var AllNiloSaharanSubfamilies = []*Subfamily{
	BertaSubfamily,
	Baga,
	FurSubfamily,
	KaduSubfamily,
	KomanSubfamily,
	KuliakSubfamily,
	KunamaSubfamily,
	MabanSubfamily,
	SaharanSubfamily,
	SonghaySubfamily,
	CentralSudanicSubfamily,
	EasternSudanicSubfamily,
	MimiDSubfamily,

	NubianSubfamily,
	DajuSubfamily,
}

// OtoMangueanFamily
var (
	OtoPameanSubfamily = &Subfamily{
		Name:     "oto_pamean_lang_subfamily",
		Family:   OtoMangueanFamily,
		IsLiving: true,
	}
	ChinantecanSubfamily = &Subfamily{
		Name:     "chinantecan_lang_subfamily",
		Family:   OtoMangueanFamily,
		IsLiving: true,
	}
	TlapanecanSubfamily = &Subfamily{
		Name:     "tlapanecan_lang_subfamily",
		Family:   OtoMangueanFamily,
		IsLiving: true,
	}
	PopolocanSubfamily = &Subfamily{
		Name:     "popolocan_lang_subfamily",
		Family:   OtoMangueanFamily,
		IsLiving: true,
	}
	ZapotecanSubfamily = &Subfamily{
		Name:     "zapotecan_lang_subfamily",
		Family:   OtoMangueanFamily,
		IsLiving: true,
	}
	AmuzgoSubfamily = &Subfamily{
		Name:     "amuzgo_lang_subfamily",
		Family:   OtoMangueanFamily,
		IsLiving: true,
	}
	MixtecanSubfamily = &Subfamily{
		Name:     "mixtecan_lang_subfamily",
		Family:   OtoMangueanFamily,
		IsLiving: true,
	}
	MangueanSubfamily = &Subfamily{
		Name:     "manguean_lang_subfamily",
		Family:   OtoMangueanFamily,
		IsLiving: false,
	}
)

var AllOtoMangueanSubfamilies = []*Subfamily{
	OtoPameanSubfamily,
	ChinantecanSubfamily,
	TlapanecanSubfamily,
	PopolocanSubfamily,
	ZapotecanSubfamily,
	AmuzgoSubfamily,
	MixtecanSubfamily,
	MangueanSubfamily,
}

// UtoAztecanFamily
var (
	NorthernUtoAztecanSubfamily = &Subfamily{
		Name:     "northern_uto_aztecan_lang_subfamily",
		Family:   UtoAztecanFamily,
		IsLiving: true,
	}
	SouthernUtoAztecanSubfamily = &Subfamily{
		Name:     "northern_uto_aztecan_lang_subfamily",
		Family:   UtoAztecanFamily,
		IsLiving: true,
	}
	TepimanSubfamily = &Subfamily{
		Name:              "tepiman_lang_subfamily",
		Family:            UtoAztecanFamily,
		ExtendedSubfamily: SouthernUtoAztecanSubfamily,
		IsLiving:          true,
	}
	TarahumaranSubfamily = &Subfamily{
		Name:              "tarahumaran_lang_subfamily",
		Family:            UtoAztecanFamily,
		ExtendedSubfamily: SouthernUtoAztecanSubfamily,
		IsLiving:          true,
	}
	CahitaSubfamily = &Subfamily{
		Name:              "cahita_lang_subfamily",
		Family:            UtoAztecanFamily,
		ExtendedSubfamily: SouthernUtoAztecanSubfamily,
		IsLiving:          true,
	}
	OpatanSubfamily = &Subfamily{
		Name:              "opatan_lang_subfamily",
		Family:            UtoAztecanFamily,
		ExtendedSubfamily: SouthernUtoAztecanSubfamily,
		IsLiving:          true,
	}
	CoracholSubfamily = &Subfamily{
		Name:              "corachol_lang_subfamily",
		Family:            UtoAztecanFamily,
		ExtendedSubfamily: SouthernUtoAztecanSubfamily,
		IsLiving:          true,
	}
	AztecanSubfamily = &Subfamily{
		Name:              "aztecan_lang_subfamily",
		Family:            UtoAztecanFamily,
		ExtendedSubfamily: SouthernUtoAztecanSubfamily,
		IsLiving:          true,
	}
)

var AllUtoAztecanSubfamilies = []*Subfamily{
	NorthernUtoAztecanSubfamily,
	SouthernUtoAztecanSubfamily,
	TepimanSubfamily,
	TarahumaranSubfamily,
	CahitaSubfamily,
	OpatanSubfamily,
	CoracholSubfamily,
	AztecanSubfamily,
}

// AustroasiaticFamily
var (
	MundaSubfamily = &Subfamily{
		Name:     "munda_lang_subfamily",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	KhasiPalaungicSubfamily = &Subfamily{
		Name:     "khasi_palaungic_lang_subfamily",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	KhmuicSubfamily = &Subfamily{
		Name:     "khmuic_lang_subfamily",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	MangSubfamily = &Subfamily{
		Name:     "mang_lang_subfamily",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	PakanicSubfamily = &Subfamily{
		Name:     "pakanic_lang_subfamily",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	VieticSubfamily = &Subfamily{
		Name:     "vietic_lang_subfamily",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	KatuicSubfamily = &Subfamily{
		Name:     "katuic_lang_subfamily",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	BahnaricSubfamily = &Subfamily{
		Name:     "bahnaric_lang_subfamily",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	KhmericSubfamily = &Subfamily{
		Name:     "khmer_lang_subfamily",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	PearicSubfamily = &Subfamily{
		Name:     "pearic_lang_subfamily",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	MonicSubfamily = &Subfamily{
		Name:     "monic_lang_subfamily",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	AslianSubfamily = &Subfamily{
		Name:     "aslian_lang_subfamily",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	NicobareseSubfamily = &Subfamily{
		Name:     "nicobarese_lang_subfamily",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
)

var AllAustroAsiaticSubfamilies = []*Subfamily{
	MundaSubfamily,
	KhasiPalaungicSubfamily,
	KhmuicSubfamily,
	MangSubfamily,
	PakanicSubfamily,
	VieticSubfamily,
	KatuicSubfamily,
	BahnaricSubfamily,
	KhmericSubfamily,
	PearicSubfamily,
	MonicSubfamily,
	AslianSubfamily,
	NicobareseSubfamily,
}

// Vietic
var (
	VietMuongSubfamily = &Subfamily{
		Name:              "viet_muong_lang_subfamily",
		Family:            AustroasiaticFamily,
		ExtendedSubfamily: VieticSubfamily,
		IsLiving:          true,
	}
	CuoiSubfamily = &Subfamily{
		Name:              "cuoi_lang_subfamily",
		Family:            AustroasiaticFamily,
		ExtendedSubfamily: VieticSubfamily,
		IsLiving:          true,
	}
	ThavungSubfamily = &Subfamily{
		Name:              "thavung_lang_subfamily",
		Family:            AustroasiaticFamily,
		ExtendedSubfamily: VieticSubfamily,
		IsLiving:          true,
	}
	ChutSubfamily = &Subfamily{
		Name:              "chut_lang_subfamily",
		Family:            AustroasiaticFamily,
		ExtendedSubfamily: VieticSubfamily,
		IsLiving:          true,
	}
	MalengSubfamily = &Subfamily{
		Name:              "maleng_lang_subfamily",
		Family:            AustroasiaticFamily,
		ExtendedSubfamily: VieticSubfamily,
		IsLiving:          true,
	}
)

var AllVieticSubfamilies = []*Subfamily{
	VietMuongSubfamily,
	CuoiSubfamily,
	ThavungSubfamily,
	ThavungSubfamily,
	ChutSubfamily,
	MalengSubfamily,
}

// TaiKadaiFamily
var (
	KraSubfamily = &Subfamily{
		Name:     "kra_lang_subfamily",
		Family:   TaiKadaiFamily,
		IsLiving: true,
	}
	KamSuiSubfamily = &Subfamily{
		Name:     "kam_sui_lang_subfamily",
		Family:   TaiKadaiFamily,
		IsLiving: true,
	}
	LakkiaSubfamily = &Subfamily{
		Name:     "lakkia_lang_subfamily",
		Family:   TaiKadaiFamily,
		IsLiving: true,
	}
	BiaoSubfamily = &Subfamily{
		Name:     "biao_lang_subfamily",
		Family:   TaiKadaiFamily,
		IsLiving: true,
	}
	BeSubfamily = &Subfamily{
		Name:     "be_lang_subfamily",
		Family:   TaiKadaiFamily,
		IsLiving: true,
	}
	TaiSubfamily = &Subfamily{
		Name:     "tai_lang_subfamily",
		Family:   TaiKadaiFamily,
		IsLiving: true,
	}
	HlaiSubfamily = &Subfamily{
		Name:     "hlai_lang_subfamily",
		Family:   TaiKadaiFamily,
		IsLiving: true,
	}
)

var AllTaiKadaiSubfamilies = []*Subfamily{
	KraSubfamily,
	KamSuiSubfamily,
	LakkiaSubfamily,
	BiaoSubfamily,
	BeSubfamily,
	TaiSubfamily,
	HlaiSubfamily,
}

// DravidianFamily
var (
	NorthernDravidianSubfamily = &Subfamily{
		Name:     "northern_dravidian_lang_subfamily",
		Family:   DravidianFamily,
		IsLiving: true,
	}
	CentralDravidianSubfamily = &Subfamily{
		Name:     "central_dravidian_lang_subfamily",
		Family:   DravidianFamily,
		IsLiving: true,
	}
	SouthCentralDravidianSubfamily = &Subfamily{
		Name:     "south_central_dravidian_lang_subfamily",
		Family:   DravidianFamily,
		IsLiving: true,
	}
	SouthDravidianSubfamily = &Subfamily{
		Name:     "south_dravidian_lang_subfamily",
		Family:   DravidianFamily,
		IsLiving: true,
	}
)

var AllDravidianSubfamilies = []*Subfamily{
	NorthernDravidianSubfamily,
	CentralDravidianSubfamily,
	SouthCentralDravidianSubfamily,
	SouthDravidianSubfamily,
}

// TupianFamily
var (
	TupiGuaraniSubfamily = &Subfamily{
		Name:     "tupi_guarani_lang_subfamily",
		Family:   TupianFamily,
		IsLiving: true,
	}
	ArikemSubfamily = &Subfamily{
		Name:     "arikem_lang_subfamily",
		Family:   TupianFamily,
		IsLiving: true,
	}
	AwetiSubfamily = &Subfamily{
		Name:     "aweti_lang_subfamily",
		Family:   TupianFamily,
		IsLiving: true,
	}
	MaweSubfamily = &Subfamily{
		Name:     "mawe_lang_subfamily",
		Family:   TupianFamily,
		IsLiving: true,
	}
	MondeSubfamily = &Subfamily{
		Name:     "monde_lang_subfamily",
		Family:   TupianFamily,
		IsLiving: true,
	}
	MundurukuSubfamily = &Subfamily{
		Name:     "munduruku_lang_subfamily",
		Family:   TupianFamily,
		IsLiving: true,
	}
	PuruboraRamaramaSubfamily = &Subfamily{
		Name:     "purubora_ramarama_lang_subfamily",
		Family:   TupianFamily,
		IsLiving: true,
	}
	TupariSubfamily = &Subfamily{
		Name:     "tupari_lang_subfamily",
		Family:   TupianFamily,
		IsLiving: true,
	}
	YurunaSubfamily = &Subfamily{
		Name:     "yuruna_lang_subfamily",
		Family:   TupianFamily,
		IsLiving: true,
	}
)

var AllTupianSubfamilies = []*Subfamily{
	TupiGuaraniSubfamily,
	ArikemSubfamily,
	AwetiSubfamily,
	MaweSubfamily,
	MondeSubfamily,
	MundurukuSubfamily,
	PuruboraRamaramaSubfamily,
	TupariSubfamily,
	YurunaSubfamily,
}

// EskimoAleut
var (
	EskimoSubfamily = &Subfamily{
		Name:     "eskimo_lang_subfamily",
		Family:   EskimoAleutFamily,
		IsLiving: true,
	}
	AleutSubfamily = &Subfamily{
		Name:     "aleut_lang_subfamily",
		Family:   EskimoAleutFamily,
		IsLiving: true,
	}
)

var AllEskimoAleutSubfamilies = []*Subfamily{EskimoSubfamily, AleutSubfamily}

// Quechua
var (
	QuechuaISubfamily = &Subfamily{
		Name:     "quechua_i_lang_subfamily",
		Family:   QuechuaFamily,
		IsLiving: true,
	}
	QuechuaIISubfamily = &Subfamily{
		Name:     "quechua_ii_lang_subfamily",
		Family:   QuechuaFamily,
		IsLiving: true,
	}
)

var AllQuechuaSubfamilies = []*Subfamily{QuechuaISubfamily, QuechuaIISubfamily}

// Fantasy
var (
	QuendianSubfamily = &Subfamily{
		Name:     "quendian_lang_subfamily",
		Family:   ElvenFamily,
		IsLiving: true,
	}
	AldmerisSubfamily = &Subfamily{
		Name:     "aldmeris_lang_subfamily",
		Family:   ElvenFamily,
		IsLiving: true,
	}
	DwermerisSubfamily = &Subfamily{
		Name:     "dwermeris_lang_subfamily",
		Family:   ElvenFamily,
		IsLiving: true,
	}
	OrcishSubfamily = &Subfamily{
		Name:     "orcish_lang_subfamily",
		Family:   OrcishFamily,
		IsLiving: true,
	}
	GiantishSubfamily = &Subfamily{
		Name:     "giantish_lang_subfamily",
		Family:   GiantishFamily,
		IsLiving: true,
	}
	DraconicSubfamily = &Subfamily{
		Name:     "draconic_lang_subfamily",
		Family:   DraconicFamily,
		IsLiving: true,
	}
	ArachnidSubfamily = &Subfamily{
		Name:     "arachnid_lang_subfamily",
		Family:   ArachnidFamily,
		IsLiving: true,
	}
	SerpentsSubfamily = &Subfamily{
		Name:     "serpents_lang_subfamily",
		Family:   SerpentsFamily,
		IsLiving: true,
	}
)

var AllFantasySubfamilies = []*Subfamily{
	QuendianSubfamily,
	AldmerisSubfamily,
	DwermerisSubfamily,
	OrcishSubfamily,
	GiantishSubfamily,
	DraconicSubfamily,
	ArachnidSubfamily,
	SerpentsSubfamily,
}

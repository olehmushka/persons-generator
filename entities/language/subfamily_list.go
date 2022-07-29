package language

// IndoEuropeanFamily
var (
	AlbanianSubfamily = &Subfamily{
		Name:     "albanian",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	AnatolianSubfamily = &Subfamily{
		Name:     "anatolian",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	ArmenianSubfamily = &Subfamily{
		Name:     "armenian",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	BaltoSlavicSubfamily = &Subfamily{
		Name:     "balto_slavic",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	CelticSubfamily = &Subfamily{
		Name:     "celtic",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	DacianSubfamily = &Subfamily{
		Name:     "dacian",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	GermanicSubfamily = &Subfamily{
		Name:     "germanic",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	GreekSubfamily = &Subfamily{
		Name:     "greek",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	IllyrianSubfamily = &Subfamily{
		Name:     "illyrian",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	IndoIranianSubfamily = &Subfamily{
		Name:     "indo_iranian",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	ItalicSubfamily = &Subfamily{
		Name:     "italic",
		Family:   IndoEuropeanFamily,
		IsLiving: true,
	}
	LiburnianSubfamily = &Subfamily{
		Name:     "liburnian",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	LusitanianSubfamily = &Subfamily{
		Name:     "lusitanian",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	MessapicSubfamily = &Subfamily{
		Name:     "messapic",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	PhrygianSubfamily = &Subfamily{
		Name:     "phrygian",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	ThracianSubfamily = &Subfamily{
		Name:     "thracian",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
	TocharianSubfamily = &Subfamily{
		Name:     "tocharian",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
)

// IndoEuropeanFamilyBranches
var (
	// Slavic
	SlavicSubfamily = &Subfamily{
		Name:              "slavic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: BaltoSlavicSubfamily,
		IsLiving:          true,
	}
	RuthenianSubfamily = &Subfamily{
		Name:              "ruthenian",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: SlavicSubfamily,
		IsLiving:          true,
	}
	MoscovianSubfamily = &Subfamily{
		Name:              "moscovian",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: SlavicSubfamily,
		IsLiving:          true,
	}
	SouthSlavicSubfamily = &Subfamily{
		Name:              "south_slavic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: SlavicSubfamily,
		IsLiving:          true,
	}
	EasternSouthSlavicSubfamily = &Subfamily{
		Name:              "eastern_south_slavic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: SouthSlavicSubfamily,
		IsLiving:          true,
	}
	WesternSouthSlavicSubfamily = &Subfamily{
		Name:              "western_south_slavic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: SouthSlavicSubfamily,
		IsLiving:          true,
	}
	WestSlavicSubfamily = &Subfamily{
		Name:              "west_slavic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: SlavicSubfamily,
		IsLiving:          true,
	}
	LechiticSubfamily = &Subfamily{
		Name:              "lechitic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: WestSlavicSubfamily,
		IsLiving:          true,
	}
	SorbianSubfamily = &Subfamily{
		Name:              "sorbian",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: WestSlavicSubfamily,
		IsLiving:          true,
	}
	CzechSlovakSubfamily = &Subfamily{
		Name:              "czech_slovak",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: WestSlavicSubfamily,
		IsLiving:          true,
	}
	BalticSubfamily = &Subfamily{
		Name:              "baltic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: BaltoSlavicSubfamily,
		IsLiving:          true,
	}
	EasternBalticSubfamily = &Subfamily{
		Name:              "eastern_baltic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: BalticSubfamily,
		IsLiving:          true,
	}
	// Celtic
	BrittonicSubfamily = &Subfamily{
		Name:              "brittonic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: CelticSubfamily,
		IsLiving:          true,
	}
	WesternBrittonicSubfamily = &Subfamily{
		Name:              "western_brittonic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: BrittonicSubfamily,
		IsLiving:          true,
	}
	SouthWesternBrittonicSubfamily = &Subfamily{
		Name:              "south_western_brittonic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: BrittonicSubfamily,
		IsLiving:          true,
	}
	GoidelicSubfamily = &Subfamily{
		Name:              "goidelic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: CelticSubfamily,
		IsLiving:          true,
	}
	// Germanic
	NorthGermanicSubfamily = &Subfamily{
		Name:              "north_germanic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: GermanicSubfamily,
		IsLiving:          true,
	}
	WestGermanicSubfamily = &Subfamily{
		Name:              "west_germanic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: GermanicSubfamily,
		IsLiving:          true,
	}
	EastGermanicSubfamily = &Subfamily{
		Name:              "east_germanic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: GermanicSubfamily,
		IsLiving:          false,
	}
	// IndoIranian
	IndoAryanSubfamily = &Subfamily{
		Name:              "indo_aryan",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: IndoIranianSubfamily,
		IsLiving:          true,
	}
	IranianSubfamily = &Subfamily{
		Name:              "iranian",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: IndoIranianSubfamily,
		IsLiving:          true,
	}
	NuriastaniSubfamily = &Subfamily{
		Name:              "nuriastani",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: IndoIranianSubfamily,
		IsLiving:          true,
	}
	// Italic
	LatinoFaliscanSubfamily = &Subfamily{
		Name:              "latino_faliscan",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: ItalicSubfamily,
		IsLiving:          true,
	}
	OscoUmbrianSubfamily = &Subfamily{
		Name:              "osco_umbrian",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: ItalicSubfamily,
		IsLiving:          false,
	}
	RomanceSubfamily = &Subfamily{
		Name:              "romance",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: LatinoFaliscanSubfamily,
		IsLiving:          true,
	}
	IberoRomanceSubfamily = &Subfamily{
		Name:              "ibero_romance",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: RomanceSubfamily,
		IsLiving:          true,
	}
	OccitanoRomanceSubfamily = &Subfamily{
		Name:              "occitano_romance",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: RomanceSubfamily,
		IsLiving:          true,
	}
	GalloRomanceSubfamily = &Subfamily{
		Name:              "gallo_romance",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: RomanceSubfamily,
		IsLiving:          true,
	}
	RhaetoRomanceSubfamily = &Subfamily{
		Name:              "rhaeto_romance",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: RomanceSubfamily,
		IsLiving:          true,
	}
	GalloItalicSubfamily = &Subfamily{
		Name:              "gallo_italic",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: RomanceSubfamily,
		IsLiving:          true,
	}
	ItaloDalmatianSubfamily = &Subfamily{
		Name:              "italo_dalmatian",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: RomanceSubfamily,
		IsLiving:          true,
	}
	EasternRomanceSubfamily = &Subfamily{
		Name:              "eastern_romance",
		Family:            IndoEuropeanFamily,
		ExtendedSubfamily: RomanceSubfamily,
		IsLiving:          true,
	}
	// Tysenian
	TysenianSubfamily = &Subfamily{
		Name:     "tysenian",
		Family:   IndoEuropeanFamily,
		IsLiving: false,
	}
)

// Basque
var (
	BasqueSubfamily = &Subfamily{
		Name:     "basque",
		Family:   BasqueFamily,
		IsLiving: true,
	}
)

// Japonic
var (
	JaponicSubfamily = &Subfamily{
		Name:     "japonic",
		Family:   JaponicFamily,
		IsLiving: true,
	}
)

// Koreanic
var (
	KoreanicSubfamily = &Subfamily{
		Name:     "koreanic",
		Family:   KoreanicFamily,
		IsLiving: true,
	}
)

// Kartvelian
var (
	KartoZanSubfamily = &Subfamily{
		Name:     "karto_zan",
		Family:   KartvelianFamily,
		IsLiving: true,
	}
	SvanSubfamily = &Subfamily{
		Name:     "svan",
		Family:   KartvelianFamily,
		IsLiving: true,
	}
)

// NigerCongoFamily
var (
	DogonSubfamily = &Subfamily{
		Name:     "dogon",
		Family:   NigerCongoFamily,
		IsLiving: false,
	}
	MandeSubfamily = &Subfamily{
		Name:     "mande",
		Family:   NigerCongoFamily,
		IsLiving: false,
	}
	IjoidSubfamily = &Subfamily{
		Name:     "ijoid",
		Family:   NigerCongoFamily,
		IsLiving: false,
	}
	LafofaSubfamily = &Subfamily{
		Name:     "lafofa (kordofanian)",
		Family:   NigerCongoFamily,
		IsLiving: false,
	}
	KruSubfamily = &Subfamily{
		Name:     "kru",
		Family:   NigerCongoFamily,
		IsLiving: false,
	}
	SiamouSubfamily = &Subfamily{
		Name:     "siamou",
		Family:   NigerCongoFamily,
		IsLiving: false,
	}
	AtlanticCongoSubfamily = &Subfamily{
		Name:     "atlantic_congo",
		Family:   NigerCongoFamily,
		IsLiving: true,
	}
)

// AustronesianFamily
var (
	RukaiSubfamily = &Subfamily{
		Name:     "rukai",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	TsouicSubfamily = &Subfamily{
		Name:     "tsouic",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	PuyumaSubfamily = &Subfamily{
		Name:     "puyuma",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	NorthwestFormosanSubfamily = &Subfamily{
		Name:     "northwest_formosan",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	EastFormosanSubfamily = &Subfamily{
		Name:     "east_formosan",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	WesternPlainsSubfamily = &Subfamily{
		Name:     "western_plains",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	AtayalicSubfamily = &Subfamily{
		Name:     "atayalic",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	BununSubfamily = &Subfamily{
		Name:     "bunun",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	PaiwanSubfamily = &Subfamily{
		Name:     "paiwan",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
	MalayoPolynesianSubfamily = &Subfamily{
		Name:     "malayo_polynesian",
		Family:   AustronesianFamily,
		IsLiving: true,
	}
)

// SinoTibetanFamily
var (
	SiniticSubfamily = &Subfamily{
		Name:     "sinitic",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	LoloBurmeseSubfamily = &Subfamily{
		Name:     "lolo_burmese",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	TibeticSubfamily = &Subfamily{
		Name:     "tibetic",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	KarenicSubfamily = &Subfamily{
		Name:     "karenic",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	BodoGaroSubfamily = &Subfamily{
		Name:     "bodo_garo",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	KukiChinSubfamily = &Subfamily{
		Name:     "kuki_chin",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	MeiteiSubfamily = &Subfamily{
		Name:     "meitei",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	TamangicSubfamily = &Subfamily{
		Name:     "tamangic",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	BaiSubfamily = &Subfamily{
		Name:     "bai",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
	JingphoLuishSubfamily = &Subfamily{
		Name:     "jingpho_luish",
		Family:   SinoTibetanFamily,
		IsLiving: true,
	}
)

// SinoTibetanFamilyBranches
var (
	// Chinese
	ChineseSubfamily = &Subfamily{
		Name:              "chinese_languages",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: SiniticSubfamily,
		IsLiving:          true,
	}
	GreaterBaiSubfamily = &Subfamily{
		Name:              "greater_bai",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: SiniticSubfamily,
		IsLiving:          true,
	}
	// LoloBurmese
	MondzishSubfamily = &Subfamily{
		Name:              "mondzish",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	BurmishSubfamily = &Subfamily{
		Name:              "burmish",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	HanoishSubfamily = &Subfamily{
		Name:              "hanoish",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	LahoishSubfamily = &Subfamily{
		Name:              "lahoish",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	NaxishSubfamily = &Subfamily{
		Name:              "naxish",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	NusoishSubfamily = &Subfamily{
		Name:              "nusoish",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	KazhuoishSubfamily = &Subfamily{
		Name:              "kazhuoish",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	LisoishSubfamily = &Subfamily{
		Name:              "lisoish",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	NisoishSubfamily = &Subfamily{
		Name:              "nisoish",
		Family:            SinoTibetanFamily,
		ExtendedSubfamily: LoloBurmeseSubfamily,
		IsLiving:          true,
	}
	// Tibetic
	// ...
)

// AfroAsiaticFamily
var (
	BerberSubfamily = &Subfamily{
		Name:     "berber",
		Family:   AfroAsiaticFamily,
		IsLiving: true,
	}
	ChadicSubfamily = &Subfamily{
		Name:     "chadic",
		Family:   AfroAsiaticFamily,
		IsLiving: true,
	}
	CushiticSubfamily = &Subfamily{
		Name:     "cushitic",
		Family:   AfroAsiaticFamily,
		IsLiving: true,
	}
	EgyptianSubfamily = &Subfamily{
		Name:     "egyptian",
		Family:   AfroAsiaticFamily,
		IsLiving: true,
	}
	SemiticSubfamily = &Subfamily{
		Name:     "semitic",
		Family:   AfroAsiaticFamily,
		IsLiving: true,
	}
	OmoticSubfamily = &Subfamily{
		Name:     "omotic",
		Family:   AfroAsiaticFamily,
		IsLiving: true,
	}
)

// NiloSaharanFamily
var (
	BertaSubfamily = &Subfamily{
		Name:     "berta",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	Baga = &Subfamily{
		Name:     "b'aga",
		Family:   NiloSaharanFamily,
		IsLiving: false,
	}
	FurSubfamily = &Subfamily{
		Name:     "fur",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	KaduSubfamily = &Subfamily{
		Name:     "kadu",
		Family:   NiloSaharanFamily,
		IsLiving: false,
	}
	KomanSubfamily = &Subfamily{
		Name:     "koman",
		Family:   NiloSaharanFamily,
		IsLiving: false,
	}
	KuliakSubfamily = &Subfamily{
		Name:     "kuliak",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	KunamaSubfamily = &Subfamily{
		Name:     "kunama",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	MabanSubfamily = &Subfamily{
		Name:     "maban",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	SaharanSubfamily = &Subfamily{
		Name:     "saharan",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	SonghaySubfamily = &Subfamily{
		Name:     "songhay",
		Family:   NiloSaharanFamily,
		IsLiving: false,
	}
	CentralSudanicSubfamily = &Subfamily{
		Name:     "central_sudanic",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	EasternSudanicSubfamily = &Subfamily{
		Name:     "eastern_sudanic",
		Family:   NiloSaharanFamily,
		IsLiving: true,
	}
	MimiDSubfamily = &Subfamily{
		Name:     "mimi_d",
		Family:   NiloSaharanFamily,
		IsLiving: false,
	}
)

// OtoMangueanFamily
var (
	OtoPameanSubfamily = &Subfamily{
		Name:     "oto_pamean",
		Family:   OtoMangueanFamily,
		IsLiving: true,
	}
	ChinantecanSubfamily = &Subfamily{
		Name:     "chinantecan",
		Family:   OtoMangueanFamily,
		IsLiving: true,
	}
	TlapanecanSubfamily = &Subfamily{
		Name:     "tlapanecan",
		Family:   OtoMangueanFamily,
		IsLiving: true,
	}
	PopolocanSubfamily = &Subfamily{
		Name:     "popolocan",
		Family:   OtoMangueanFamily,
		IsLiving: true,
	}
	ZapotecanSubfamily = &Subfamily{
		Name:     "zapotecan",
		Family:   OtoMangueanFamily,
		IsLiving: true,
	}
	AmuzgoSubfamily = &Subfamily{
		Name:     "amuzgo",
		Family:   OtoMangueanFamily,
		IsLiving: true,
	}
	MixtecanSubfamily = &Subfamily{
		Name:     "mixtecan",
		Family:   OtoMangueanFamily,
		IsLiving: true,
	}
	MangueanSubfamily = &Subfamily{
		Name:     "manguean",
		Family:   OtoMangueanFamily,
		IsLiving: false,
	}
)

// AustroasiaticFamily
var (
	MundaSubfamily = &Subfamily{
		Name:     "munda",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	KhasiPalaungicSubfamily = &Subfamily{
		Name:     "khasi_palaungic",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	KhmuicSubfamily = &Subfamily{
		Name:     "khmuic",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	MangSubfamily = &Subfamily{
		Name:     "mang",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	PakanicSubfamily = &Subfamily{
		Name:     "pakanic",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	VieticSubfamily = &Subfamily{
		Name:     "vietic",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	KatuicSubfamily = &Subfamily{
		Name:     "katuic",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	BahnaricSubfamily = &Subfamily{
		Name:     "bahnaric",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	KhmerSubfamily = &Subfamily{
		Name:     "khmer",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	PearicSubfamily = &Subfamily{
		Name:     "pearic",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	MonicSubfamily = &Subfamily{
		Name:     "monic",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	AslianSubfamily = &Subfamily{
		Name:     "aslian",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
	NicobareseSubfamily = &Subfamily{
		Name:     "nicobarese",
		Family:   AustroasiaticFamily,
		IsLiving: true,
	}
)

// TaiKadaiFamily
var (
	KraSubfamily = &Subfamily{
		Name:     "kra",
		Family:   TaiKadaiFamily,
		IsLiving: true,
	}
	KamSuiSubfamily = &Subfamily{
		Name:     "kam_sui",
		Family:   TaiKadaiFamily,
		IsLiving: true,
	}
	LakkiaSubfamily = &Subfamily{
		Name:     "lakkia",
		Family:   TaiKadaiFamily,
		IsLiving: true,
	}
	BiaoSubfamily = &Subfamily{
		Name:     "biao",
		Family:   TaiKadaiFamily,
		IsLiving: true,
	}
	BeSubfamily = &Subfamily{
		Name:     "be",
		Family:   TaiKadaiFamily,
		IsLiving: true,
	}
	TaiSubfamily = &Subfamily{
		Name:     "tai",
		Family:   TaiKadaiFamily,
		IsLiving: true,
	}
	HlaiSubfamily = &Subfamily{
		Name:     "hlai",
		Family:   TaiKadaiFamily,
		IsLiving: true,
	}
)

// DravidianFamily
var (
	NorthernDravidianSubfamily = &Subfamily{
		Name:     "northern_dravidian",
		Family:   DravidianFamily,
		IsLiving: true,
	}
	CentralDravidianSubfamily = &Subfamily{
		Name:     "central_dravidian",
		Family:   DravidianFamily,
		IsLiving: true,
	}
	SouthCentralDravidianSubfamily = &Subfamily{
		Name:     "south_central_dravidian",
		Family:   DravidianFamily,
		IsLiving: true,
	}
	SouthDravidianSubfamily = &Subfamily{
		Name:     "south_dravidian",
		Family:   DravidianFamily,
		IsLiving: true,
	}
)

// TupianFamily
var (
	TupiGuaraniSubfamily = &Subfamily{
		Name:     "tupi_guarani",
		Family:   TupianFamily,
		IsLiving: true,
	}
	ArikemSubfamily = &Subfamily{
		Name:     "arikem",
		Family:   TupianFamily,
		IsLiving: true,
	}
	AwetïSubfamily = &Subfamily{
		Name:     "awetï",
		Family:   TupianFamily,
		IsLiving: true,
	}
	MawéSubfamily = &Subfamily{
		Name:     "mawé",
		Family:   TupianFamily,
		IsLiving: true,
	}
	MondéSubfamily = &Subfamily{
		Name:     "mondé",
		Family:   TupianFamily,
		IsLiving: true,
	}
	MundurukúSubfamily = &Subfamily{
		Name:     "mundurukú",
		Family:   TupianFamily,
		IsLiving: true,
	}
	PuruboráRamaramaSubfamily = &Subfamily{
		Name:     "puruborá_ramarama",
		Family:   TupianFamily,
		IsLiving: true,
	}
	TuparíSubfamily = &Subfamily{
		Name:     "tuparí",
		Family:   TupianFamily,
		IsLiving: true,
	}
	YurúnaSubfamily = &Subfamily{
		Name:     "yurúna",
		Family:   TupianFamily,
		IsLiving: true,
	}
)

// Fantasy
var (
	QuendianSubfamily = &Subfamily{
		Name:     "quendian",
		Family:   ElvenFamily,
		IsLiving: true,
	}
	AldmerisSubfamily = &Subfamily{
		Name:     "aldmeris",
		Family:   ElvenFamily,
		IsLiving: true,
	}
	DwermerisSubfamily = &Subfamily{
		Name:     "dwermeris",
		Family:   ElvenFamily,
		IsLiving: true,
	}
	OrcishSubfamily = &Subfamily{
		Name:     "orcish",
		Family:   OrcishFamily,
		IsLiving: true,
	}
	GiantishSubfamily = &Subfamily{
		Name:     "giantish",
		Family:   GiantishFamily,
		IsLiving: true,
	}
	DraconicSubfamily = &Subfamily{
		Name:     "draconic",
		Family:   DraconicFamily,
		IsLiving: true,
	}
	ArachnidSubfamily = &Subfamily{
		Name:     "arachnid",
		Family:   ArachnidFamily,
		IsLiving: true,
	}
	SerpentsSubfamily = &Subfamily{
		Name:     "serpents",
		Family:   SerpentsFamily,
		IsLiving: true,
	}
)

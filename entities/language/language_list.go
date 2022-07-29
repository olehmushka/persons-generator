package language

import "persons_generator/tools"

var AllLanguages = tools.Merge(
	AllIndoEuropeanLanguages,
	AllBasqueLanguages,
	AllJaponicLanguages,
	AllKoreanicLanguages,
	AllKartvelianLanguages,
	AllSinoTibetanLanguages,
	AllFantasyLanguages,
)

// IndoEuropeanFamily
var (
	Albanian = &Language{
		Name:      "albanian",
		Subfamily: AlbanianSubfamily,
		IsLiving:  true,
		WordBase:  AlbanianWordBase,
	}
	Armenian = &Language{
		Name:      "armenian",
		Subfamily: ArmenianSubfamily,
		IsLiving:  true,
		WordBase:  ArmenianWordBase,
	}
	// Slavic
	Ruthenian = &Language{
		Name:      "ruthenian",
		Subfamily: RuthenianSubfamily,
		IsLiving:  false,
		WordBase:  RuthenianWordBase,
	}
	Belarusian = &Language{
		Name:      "belarusian",
		Subfamily: RuthenianSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding belarusian to word bases
	}
	Ukrainian = &Language{
		Name:      "ukrainian",
		Subfamily: RuthenianSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding ukrainian to word bases
	}
	Russian = &Language{
		Name:      "russian",
		Subfamily: MoscovianSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding russian to word bases
	}
	Slovenian = &Language{
		Name:      "slovenian",
		Subfamily: WesternSouthSlavicSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding slovenian to word bases
	}
	Croatian = &Language{
		Name:      "croation",
		Subfamily: WesternSouthSlavicSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding croation to word bases
	}
	Serbian = &Language{
		Name:      "serbian",
		Subfamily: WesternSouthSlavicSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding serbian to word bases
	}
	Bosnian = &Language{
		Name:      "bosnian",
		Subfamily: WesternSouthSlavicSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding bosnian to word bases
	}
	Bulgarian = &Language{
		Name:      "bulgarian",
		Subfamily: EasternSouthSlavicSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding bulgarian to word bases
	}
	Macedonian = &Language{
		Name:      "macedonian",
		Subfamily: EasternSouthSlavicSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding macedonian to word bases
	}
	Polish = &Language{
		Name:      "polish",
		Subfamily: LechiticSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding polish to word bases
	}
	Silesian = &Language{
		Name:      "silesian",
		Subfamily: LechiticSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding silesian to word bases
	}
	LowerSorbian = &Language{
		Name:      "lower_sorbian",
		Subfamily: SorbianSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding lower_sorbian to word bases
	}
	UpperSorbian = &Language{
		Name:      "upper_sorbian",
		Subfamily: SorbianSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding upper_sorbian to word bases
	}
	Czech = &Language{
		Name:      "czech",
		Subfamily: CzechSlovakSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding czech to word bases
	}
	Slovak = &Language{
		Name:      "slovak",
		Subfamily: CzechSlovakSubfamily,
		IsLiving:  true,
		WordBase:  RuthenianWordBase, // @TODO can be changed after adding slovak to word bases
	}
	// Baltic
	Latvian = &Language{
		Name:      "latvian",
		Subfamily: EasternBalticSubfamily,
		IsLiving:  true,
		WordBase:  LithuanianWordBase, // @TODO can be changed after adding latvian to word bases
	}
	Latgalian = &Language{
		Name:      "latgalian",
		Subfamily: EasternBalticSubfamily,
		IsLiving:  true,
		WordBase:  LithuanianWordBase, // @TODO can be changed after adding latgalian to word bases
	}
	Lithuanian = &Language{
		Name:      "lithuanian",
		Subfamily: EasternBalticSubfamily,
		IsLiving:  true,
		WordBase:  LithuanianWordBase,
	}
	Semigallian = &Language{
		Name:      "semigallian",
		Subfamily: EasternBalticSubfamily,
		IsLiving:  false,
		WordBase:  LithuanianWordBase, // @TODO can be changed after adding semigallian to word bases
	}
	// Celtic
	Celtiberian = &Language{
		Name:      "celtiberian",
		Subfamily: CelticSubfamily,
		IsLiving:  false,
		WordBase:  CelticWordBase,
	}
	Gallaecian = &Language{
		Name:      "gallaecian",
		Subfamily: CelticSubfamily,
		IsLiving:  false,
		WordBase:  CelticWordBase,
	}
	Noric = &Language{
		Name:      "noric",
		Subfamily: CelticSubfamily,
		IsLiving:  false,
		WordBase:  CelticWordBase,
	}
	Pictish = &Language{
		Name:      "pictish",
		Subfamily: BrittonicSubfamily,
		IsLiving:  false,
		WordBase:  CelticWordBase,
	}
	Cumbric = &Language{
		Name:      "cumbric",
		Subfamily: BrittonicSubfamily,
		IsLiving:  false,
		WordBase:  CelticWordBase,
	}
	Welsh = &Language{
		Name:      "welsh",
		Subfamily: WesternBrittonicSubfamily,
		IsLiving:  true,
		WordBase:  CelticWordBase,
	}
	Cornish = &Language{
		Name:      "cornish",
		Subfamily: SouthWesternBrittonicSubfamily,
		IsLiving:  true,
		WordBase:  CelticWordBase,
	}
	Breton = &Language{
		Name:      "breton",
		Subfamily: SouthWesternBrittonicSubfamily,
		IsLiving:  true,
		WordBase:  CelticWordBase,
	}
	Irish = &Language{
		Name:      "irish",
		Subfamily: GoidelicSubfamily,
		IsLiving:  true,
		WordBase:  CelticWordBase,
	}
	Manx = &Language{
		Name:      "manx",
		Subfamily: GoidelicSubfamily,
		IsLiving:  true,
		WordBase:  CelticWordBase,
	}
	ScottishGaelic = &Language{
		Name:      "scottish_gaelic",
		Subfamily: GoidelicSubfamily,
		IsLiving:  true,
		WordBase:  CelticWordBase,
	}
	// Dacian
	Dacian = &Language{
		Name:      "dacian",
		Subfamily: DacianSubfamily,
		IsLiving:  false,
	}
	// Germanic
	Danish = &Language{
		Name:      "danish",
		Subfamily: NorthGermanicSubfamily,
		IsLiving:  true,
		WordBase:  NordicWordBase, // @TODO can be changed after adding danish to word bases
	}
	Faroese = &Language{
		Name:      "faroese",
		Subfamily: NorthGermanicSubfamily,
		IsLiving:  true,
		WordBase:  NordicWordBase, // @TODO can be changed after adding faroese to word bases
	}
	Icelandic = &Language{
		Name:      "icelandic",
		Subfamily: NorthGermanicSubfamily,
		IsLiving:  true,
		WordBase:  NordicWordBase, // @TODO can be changed after adding icelandic to word bases
	}
	Norwegian = &Language{
		Name:      "norwegian",
		Subfamily: NorthGermanicSubfamily,
		IsLiving:  true,
		WordBase:  NordicWordBase, // @TODO can be changed after adding norwegian to word bases
	}
	Swedish = &Language{
		Name:      "swedish",
		Subfamily: NorthGermanicSubfamily,
		IsLiving:  true,
		WordBase:  NordicWordBase,
	}
	Dalecarlian = &Language{
		Name:      "dalecarlian",
		Subfamily: NorthGermanicSubfamily,
		IsLiving:  true,
		WordBase:  NordicWordBase, // @TODO can be changed after adding dalecarlian to word bases
	}
	Gutnish = &Language{
		Name:      "gutnish",
		Subfamily: NorthGermanicSubfamily,
		IsLiving:  true,
		WordBase:  NordicWordBase, // @TODO can be changed after adding gutnish to word bases
	}
	English = &Language{
		Name:      "english",
		Subfamily: WestGermanicSubfamily,
		IsLiving:  true,
		WordBase:  EnglishWordBase,
	}
	Scots = &Language{
		Name:      "scots",
		Subfamily: WestGermanicSubfamily,
		IsLiving:  true,
		WordBase:  EnglishWordBase, // @TODO can be changed after adding scots to word bases
	}
	Yola = &Language{
		Name:      "yola",
		Subfamily: WestGermanicSubfamily,
		IsLiving:  true,
		WordBase:  EnglishWordBase, // @TODO can be changed after adding yola to word bases
	}
	Frisian = &Language{
		Name:      "frisian",
		Subfamily: WestGermanicSubfamily,
		IsLiving:  true,
		WordBase:  EnglishWordBase, // @TODO can be changed after adding frisian to word bases
	}
	German = &Language{
		Name:      "german",
		Subfamily: WestGermanicSubfamily,
		IsLiving:  true,
		WordBase:  GermanWordBase,
	}
	Dutch = &Language{
		Name:      "dutch",
		Subfamily: WestGermanicSubfamily,
		IsLiving:  true,
		WordBase:  GermanWordBase, // @TODO can be changed after adding dutch to word bases
	}
	Luxembourgish = &Language{
		Name:      "luxembourgish",
		Subfamily: WestGermanicSubfamily,
		IsLiving:  true,
		WordBase:  GermanWordBase, // @TODO can be changed after adding luxembourgish to word bases
	}
	Yiddish = &Language{
		Name:      "yiddish",
		Subfamily: WestGermanicSubfamily,
		IsLiving:  true,
		WordBase:  GermanWordBase, // @TODO can be changed after adding yiddish to word bases
	}
	Gothic = &Language{
		Name:      "gothic",
		Subfamily: EastGermanicSubfamily,
		IsLiving:  false,
		WordBase:  GermanWordBase, // @TODO can be changed after adding gothic to word bases
	}
	Vandalic = &Language{
		Name:      "vandalic",
		Subfamily: EastGermanicSubfamily,
		IsLiving:  false,
		WordBase:  GermanWordBase, // @TODO can be changed after adding vandalic to word bases
	}
	Burgundian = &Language{
		Name:      "burgundian",
		Subfamily: EastGermanicSubfamily,
		IsLiving:  false,
		WordBase:  GermanWordBase, // @TODO can be changed after adding burgundian to word bases
	}
	// Greek
	Greek = &Language{
		Name:      "greek",
		Subfamily: GreekSubfamily,
		IsLiving:  true,
		WordBase:  GreekWordBase,
	}
	// Illyrian
	Illyrian = &Language{
		Name:      "illyrian",
		Subfamily: IllyrianSubfamily,
		IsLiving:  false,
		WordBase:  IllyrianWordBase,
	}
	// IndoIranian
	Pashai = &Language{
		Name:      "pashai",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Chitrali = &Language{
		Name:      "chitrali",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Shina = &Language{
		Name:      "shina",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Kohistani = &Language{
		Name:      "kohistani",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Kashmiri = &Language{
		Name:      "kashmiri",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Punjabi = &Language{
		Name:      "punjabi",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Sindhi = &Language{
		Name:      "sindhi",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Rajasthani = &Language{
		Name:      "rajasthani",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Gujarati = &Language{
		Name:      "gujarati",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Bhili = &Language{
		Name:      "bhili",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Khandeshi = &Language{
		Name:      "khandeshi",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	HimachaliDogri = &Language{
		Name:      "himachali_dogri",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	GirhwaliKumaoni = &Language{
		Name:      "girhwali_kumaoni",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Nepali = &Language{
		Name:      "nepali",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Hindi = &Language{
		Name:      "hindi",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Bihari = &Language{
		Name:      "bihari",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	BengaliAssamese = &Language{
		Name:      "bengali_assamese",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Odia = &Language{
		Name:      "odia",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	Halbi = &Language{
		Name:      "halbi",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	MarathiKonkani = &Language{
		Name:      "marathi_konkani",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	SinhalaMaldivian = &Language{
		Name:      "sinhala_maldivian",
		Subfamily: IndoAryanSubfamily,
		IsLiving:  true,
		WordBase:  HindiWordBase,
	}
	// Iranian
	Farsi = &Language{
		Name:      "farsi",
		Subfamily: IranianSubfamily,
		IsLiving:  true,
		WordBase:  IranianWordBase,
	}
	Avestan = &Language{
		Name:      "avestan",
		Subfamily: IranianSubfamily,
		IsLiving:  false,
		WordBase:  IranianWordBase,
	}
	Dari = &Language{
		Name:      "dari",
		Subfamily: IranianSubfamily,
		IsLiving:  true,
		WordBase:  IranianWordBase,
	}
	Tajik = &Language{
		Name:      "tajik",
		Subfamily: IranianSubfamily,
		IsLiving:  true,
		WordBase:  IranianWordBase,
	}
	Pashto = &Language{
		Name:      "pashto",
		Subfamily: IranianSubfamily,
		IsLiving:  true,
		WordBase:  IranianWordBase,
	}
	Kurdish = &Language{
		Name:      "kurdish",
		Subfamily: IranianSubfamily,
		IsLiving:  true,
		WordBase:  IranianWordBase,
	}
	Luri = &Language{
		Name:      "luri",
		Subfamily: IranianSubfamily,
		IsLiving:  true,
		WordBase:  IranianWordBase,
	}
	Balochi = &Language{
		Name:      "balochi",
		Subfamily: IranianSubfamily,
		IsLiving:  true,
		WordBase:  IranianWordBase,
	}
	// Nuriastani
	KamkataVari = &Language{
		Name:      "kamkata_vari",
		Subfamily: NuriastaniSubfamily,
		WordBase:  KarnatakaWordBase, // @TODO change it into KamkataVari when available
	}
	VasiVari = &Language{
		Name:      "vasi_vari",
		Subfamily: NuriastaniSubfamily,
		WordBase:  KarnatakaWordBase, // @TODO change it into VasiVari when available
		IsLiving:  true,
	}
	Askunu = &Language{
		Name:      "askunu",
		Subfamily: NuriastaniSubfamily,
		WordBase:  KarnatakaWordBase, // @TODO change it into Askunu when available
		IsLiving:  true,
	}
	Waigali = &Language{
		Name:      "waigali",
		Subfamily: NuriastaniSubfamily,
		WordBase:  KarnatakaWordBase, // @TODO change it into Waigali when available
		IsLiving:  true,
	}
	Tregami = &Language{
		Name:      "tregami",
		Subfamily: NuriastaniSubfamily,
		WordBase:  KarnatakaWordBase, // @TODO change it into Tregami when available
		IsLiving:  true,
	}
	Zemiaki = &Language{
		Name:      "zemiaki",
		Subfamily: NuriastaniSubfamily,
		WordBase:  KarnatakaWordBase, // @TODO change it into Zemiaki when available
		IsLiving:  true,
	}
	// Italic
	Venetic = &Language{
		Name:      "venetic",
		Subfamily: ItalicSubfamily,
		WordBase:  ItalianWordBase, // @TODO change it into venetic when available
		IsLiving:  false,
	}
	Sicel = &Language{
		Name:      "sicel",
		Subfamily: ItalicSubfamily,
		WordBase:  ItalianWordBase, // @TODO change it into sicel when available
		IsLiving:  false,
	}
	Lusitanian = &Language{
		Name:      "lusitanian",
		Subfamily: ItalicSubfamily,
		WordBase:  ItalianWordBase, // @TODO change it into lusitanian when available
		IsLiving:  false,
	}
	Latin = &Language{
		Name:      "latin",
		Subfamily: LatinoFaliscanSubfamily,
		WordBase:  LatinWordBase, // @TODO change it into latin when available
		IsLiving:  true,
	}
	Faliscan = &Language{
		Name:      "faliscan",
		Subfamily: LatinoFaliscanSubfamily,
		WordBase:  LatinWordBase, // @TODO change it into faliscan when available
		IsLiving:  false,
	}
	Lanuvain = &Language{
		Name:      "lanuvain",
		Subfamily: LatinoFaliscanSubfamily,
		WordBase:  LatinWordBase, // @TODO change it into lanuvain when available
		IsLiving:  false,
	}
	Venetian = &Language{
		Name:      "venetian",
		Subfamily: RomanceSubfamily,
		WordBase:  ItalianWordBase, // @TODO change it into venetian when available
		IsLiving:  true,
	}
	Sardinian = &Language{
		Name:      "sardinian",
		Subfamily: RomanceSubfamily,
		WordBase:  ItalianWordBase, // @TODO change it into sardinian when available
		IsLiving:  true,
	}
	Portuguese = &Language{
		Name:      "portuguese",
		Subfamily: IberoRomanceSubfamily,
		WordBase:  PortugueseWordBase,
		IsLiving:  true,
	}
	Galician = &Language{
		Name:      "galician",
		Subfamily: IberoRomanceSubfamily,
		WordBase:  PortugueseWordBase,
		IsLiving:  true,
	}
	Asturleonese = &Language{
		Name:      "asturleonese",
		Subfamily: IberoRomanceSubfamily,
		WordBase:  SpanishWordBase,
		IsLiving:  true,
	}
	Mirandese = &Language{
		Name:      "mirandese",
		Subfamily: IberoRomanceSubfamily,
		WordBase:  PortugueseWordBase,
		IsLiving:  true,
	}
	Spanish = &Language{
		Name:      "spanish",
		Subfamily: IberoRomanceSubfamily,
		WordBase:  SpanishWordBase,
		IsLiving:  true,
	}
	Aragonese = &Language{
		Name:      "aragonese",
		Subfamily: IberoRomanceSubfamily,
		WordBase:  SpanishWordBase,
		IsLiving:  true,
	}
	Ladino = &Language{
		Name:      "ladino",
		Subfamily: IberoRomanceSubfamily,
		WordBase:  SpanishWordBase,
		IsLiving:  true,
	}
	Catalan = &Language{
		Name:      "catalan",
		Subfamily: OccitanoRomanceSubfamily,
		WordBase:  SpanishWordBase,
		IsLiving:  true,
	}
	Occitan = &Language{
		Name:      "occitan",
		Subfamily: OccitanoRomanceSubfamily,
		WordBase:  FrenchWordBase,
		IsLiving:  true,
	}
	Gascon = &Language{
		Name:      "gascon",
		Subfamily: OccitanoRomanceSubfamily,
		WordBase:  FrenchWordBase,
		IsLiving:  true,
	}
	French = &Language{
		Name:      "french",
		Subfamily: GalloRomanceSubfamily,
		WordBase:  FrenchWordBase,
		IsLiving:  true,
	}
	FrancoProvencal = &Language{
		Name:      "franco_provencal",
		Subfamily: GalloRomanceSubfamily,
		WordBase:  FrenchWordBase,
		IsLiving:  true,
	}
	Romansh = &Language{
		Name:      "romansh",
		Subfamily: RhaetoRomanceSubfamily,
		WordBase:  FrenchWordBase,
		IsLiving:  true,
	}
	Ladin = &Language{
		Name:      "ladin",
		Subfamily: RhaetoRomanceSubfamily,
		WordBase:  FrenchWordBase,
		IsLiving:  true,
	}
	Friulian = &Language{
		Name:      "friulian",
		Subfamily: RhaetoRomanceSubfamily,
		WordBase:  FrenchWordBase,
		IsLiving:  true,
	}
	Piedmontense = &Language{
		Name:      "piedmontense",
		Subfamily: GalloItalicSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  true,
	}
	Ligurian = &Language{
		Name:      "ligurian",
		Subfamily: GalloItalicSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  true,
	}
	Lombard = &Language{
		Name:      "lombard",
		Subfamily: GalloItalicSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  true,
	}
	Emilian = &Language{
		Name:      "emilian",
		Subfamily: GalloItalicSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  true,
	}
	Romagnol = &Language{
		Name:      "romagnol",
		Subfamily: GalloItalicSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  true,
	}
	Italian = &Language{
		Name:      "italian",
		Subfamily: ItaloDalmatianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  true,
	}
	Sicilian = &Language{
		Name:      "sicilian",
		Subfamily: ItaloDalmatianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  true,
	}
	Neapolitan = &Language{
		Name:      "neapolitan",
		Subfamily: ItaloDalmatianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  true,
	}
	Dalmatian = &Language{
		Name:      "dalmatian",
		Subfamily: ItaloDalmatianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  true,
	}
	Istriot = &Language{
		Name:      "istriot",
		Subfamily: ItaloDalmatianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  true,
	}
	Romanian = &Language{
		Name:      "romanian",
		Subfamily: EasternRomanceSubfamily,
		WordBase:  LatinWordBase,
		IsLiving:  true,
	}
	Aromanian = &Language{
		Name:      "aromanian",
		Subfamily: EasternRomanceSubfamily,
		WordBase:  LatinWordBase,
		IsLiving:  true,
	}
	MaglenoRomanian = &Language{
		Name:      "magleno_romanian",
		Subfamily: EasternRomanceSubfamily,
		WordBase:  LatinWordBase,
		IsLiving:  true,
	}
	IstroRomanian = &Language{
		Name:      "istro_romanian",
		Subfamily: EasternRomanceSubfamily,
		WordBase:  LatinWordBase,
		IsLiving:  true,
	}
	Oscan = &Language{
		Name:      "oscan",
		Subfamily: OscoUmbrianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  false,
	}
	Umbrian = &Language{
		Name:      "umbrian",
		Subfamily: OscoUmbrianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  false,
	}
	Volscian = &Language{
		Name:      "volscian",
		Subfamily: OscoUmbrianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  false,
	}
	Sabine = &Language{
		Name:      "sabine",
		Subfamily: OscoUmbrianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  false,
	}
	SouthPicene = &Language{
		Name:      "south_picene",
		Subfamily: OscoUmbrianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  false,
	}
	Marsian = &Language{
		Name:      "marsian",
		Subfamily: OscoUmbrianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  false,
	}
	Paeligni = &Language{
		Name:      "paeligni",
		Subfamily: OscoUmbrianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  false,
	}
	Hernican = &Language{
		Name:      "hernican",
		Subfamily: OscoUmbrianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  false,
	}
	Marrucinian = &Language{
		Name:      "marrucinian",
		Subfamily: OscoUmbrianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  false,
	}
	PreSamnite = &Language{
		Name:      "pre_samnite",
		Subfamily: OscoUmbrianSubfamily,
		WordBase:  ItalianWordBase,
		IsLiving:  false,
	}
	// Tysenian
	Rhaetic = &Language{
		Name:      "rhaetic",
		Subfamily: TysenianSubfamily,
		WordBase:  EtruscanWordBase,
		IsLiving:  false,
	}
	Etruscan = &Language{
		Name:      "etruscan",
		Subfamily: TysenianSubfamily,
		WordBase:  EtruscanWordBase,
		IsLiving:  false,
	}
)

var AllIndoEuropeanLanguages = []*Language{
	Albanian,
	Armenian,
	Belarusian,
	Ukrainian,
	Russian,
	Slovenian,
	Croatian,
	Serbian,
	Bosnian,
	Bulgarian,
	Macedonian,
	Polish,
	Silesian,
	LowerSorbian,
	UpperSorbian,
	Czech,
	Slovak,
	Latvian,
	Latgalian,
	Lithuanian,
	Semigallian,
	Celtiberian,
	Gallaecian,
	Noric,
	Pictish,
	Cumbric,
	Welsh,
	Cornish,
	Breton,
	Irish,
	Manx,
	ScottishGaelic,
	Dacian,
	Danish,
	Faroese,
	Icelandic,
	Norwegian,
	Swedish,
	Dalecarlian,
	Gutnish,
	English,
	Scots,
	Yola,
	Frisian,
	German,
	Dutch,
	Luxembourgish,
	Yiddish,
	Gothic,
	Vandalic,
	Burgundian,
	Greek,
	Illyrian,
	Pashai,
	Chitrali,
	Shina,
	Kohistani,
	Kashmiri,
	Punjabi,
	Sindhi,
	Rajasthani,
	Gujarati,
	Bhili,
	Khandeshi,
	HimachaliDogri,
	GirhwaliKumaoni,
	Nepali,
	Hindi,
	Bihari,
	BengaliAssamese,
	Odia,
	Halbi,
	MarathiKonkani,
	SinhalaMaldivian,
	Farsi,
	Avestan,
	Dari,
	Tajik,
	Pashto,
	Kurdish,
	Luri,
	Balochi,
	KamkataVari,
	VasiVari,
	Askunu,
	Waigali,
	Tregami,
	Zemiaki,
	Venetic,
	Sicel,
	Lusitanian,
	Latin,
	Faliscan,
	Lanuvain,
	Venetian,
	Sardinian,
	Portuguese,
	Galician,
	Asturleonese,
	Mirandese,
	Spanish,
	Aragonese,
	Ladino,
	Catalan,
	Occitan,
	Gascon,
	French,
	FrancoProvencal,
	Romansh,
	Ladin,
	Friulian,
	Piedmontense,
	Ligurian,
	Lombard,
	Emilian,
	Romagnol,
	Italian,
	Sicilian,
	Neapolitan,
	Dalmatian,
	Istriot,
	Romanian,
	Aromanian,
	MaglenoRomanian,
	IstroRomanian,
	Oscan,
	Umbrian,
	Volscian,
	Sabine,
	SouthPicene,
	Marsian,
	Paeligni,
	Hernican,
	Marrucinian,
	PreSamnite,
	Rhaetic,
	Etruscan,
}

// Basque
var (
	Basque = &Language{
		Name:      "basque",
		Subfamily: BasqueSubfamily,
		IsLiving:  true,
		WordBase:  BasqueWordBase,
	}
)
var AllBasqueLanguages = []*Language{Basque}

// Japonic
var (
	Japanese = &Language{
		Name:      "japanese",
		Subfamily: JaponicSubfamily,
		IsLiving:  true,
		WordBase:  JapaneseWordBase,
	}
)
var AllJaponicLanguages = []*Language{Japanese}

// Koreanic
var (
	Korean = &Language{
		Name:      "korean",
		Subfamily: KoreanicSubfamily,
		IsLiving:  true,
		WordBase:  KoreanWordBase,
	}
)
var AllKoreanicLanguages = []*Language{Korean}

// Kartvelian
var (
	Georgian = &Language{
		Name:      "Georgian",
		Subfamily: KartoZanSubfamily,
		WordBase:  GeorgianWordBase,
		IsLiving:  true,
	}
	Mingrelian = &Language{
		Name:      "Mingrelian",
		Subfamily: KartoZanSubfamily,
		WordBase:  GeorgianWordBase,
		IsLiving:  true,
	}
	Laz = &Language{
		Name:      "Laz",
		Subfamily: KartoZanSubfamily,
		WordBase:  GeorgianWordBase,
		IsLiving:  true,
	}
	Svan = &Language{
		Name:      "Svan",
		Subfamily: SvanSubfamily,
		WordBase:  GeorgianWordBase,
		IsLiving:  true,
	}
)

var AllKartvelianLanguages = []*Language{
	Georgian,
	Mingrelian,
	Laz,
	Svan,
}

// SinoTibetanFamily
var (
	// Sinitic
	Chinese = &Language{
		Name:      "chinese",
		Subfamily: ChineseSubfamily,
		IsLiving:  true,
		WordBase:  ChineseWordBase,
	}
	Cantonese = &Language{
		Name:      "cantonese",
		Subfamily: ChineseSubfamily,
		IsLiving:  true,
		WordBase:  CantoneseWordBase,
	}
	Bai = &Language{
		Name:      "bai",
		Subfamily: GreaterBaiSubfamily,
		IsLiving:  true,
		WordBase:  ChineseWordBase,
	}
	Caijia = &Language{
		Name:      "caijia",
		Subfamily: GreaterBaiSubfamily,
		IsLiving:  true,
		WordBase:  ChineseWordBase,
	}
	Longjia = &Language{
		Name:      "longjia",
		Subfamily: GreaterBaiSubfamily,
		IsLiving:  true,
		WordBase:  ChineseWordBase,
	}
	Luren = &Language{
		Name:      "luren",
		Subfamily: GreaterBaiSubfamily,
		IsLiving:  true,
		WordBase:  ChineseWordBase,
	}
	// LoloBurmese
	Muangphe = &Language{
		Name:      "muangphe",
		Subfamily: MondzishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Mango = &Language{
		Name:      "mango",
		Subfamily: MondzishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Manga = &Language{
		Name:      "manga",
		Subfamily: MondzishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Maang = &Language{
		Name:      "maang",
		Subfamily: MondzishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Mondzi = &Language{
		Name:      "mondzi",
		Subfamily: MondzishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Maza = &Language{
		Name:      "maza",
		Subfamily: MondzishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Mauphu = &Language{
		Name:      "mauphu",
		Subfamily: MondzishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Motang = &Language{
		Name:      "motang",
		Subfamily: MondzishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Mongphu = &Language{
		Name:      "mongphu",
		Subfamily: MondzishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Burmish = &Language{
		Name:      "burmish",
		Subfamily: BurmishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Jinuo = &Language{
		Name:      "jinuo",
		Subfamily: HanoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Sangkong = &Language{
		Name:      "sangkong",
		Subfamily: HanoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Bisu = &Language{
		Name:      "bisu",
		Subfamily: HanoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Phunoi = &Language{
		Name:      "phunoi",
		Subfamily: HanoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Pyen = &Language{
		Name:      "pyen",
		Subfamily: HanoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Sila = &Language{
		Name:      "sila",
		Subfamily: HanoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Phana = &Language{
		Name:      "phana",
		Subfamily: HanoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Akeu = &Language{
		Name:      "akeu",
		Subfamily: HanoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Hani = &Language{
		Name:      "hani",
		Subfamily: HanoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Piyo = &Language{
		Name:      "piyo",
		Subfamily: HanoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Enu = &Language{
		Name:      "enu",
		Subfamily: HanoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Mpi = &Language{
		Name:      "mpi",
		Subfamily: HanoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Kaduo = &Language{
		Name:      "kaduo",
		Subfamily: HanoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Laha = &Language{
		Name:      "laha",
		Subfamily: LahoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Kucong = &Language{
		Name:      "kucong",
		Subfamily: LahoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Namuyi = &Language{
		Name:      "namuyi",
		Subfamily: NaxishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Shixing = &Language{
		Name:      "shixing",
		Subfamily: NaxishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Naish = &Language{
		Name:      "naish",
		Subfamily: NaxishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Nusu = &Language{
		Name:      "nusu",
		Subfamily: NusoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Zauzou = &Language{
		Name:      "zauzou",
		Subfamily: NusoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Katso = &Language{
		Name:      "katso",
		Subfamily: KazhuoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Samu = &Language{
		Name:      "samu",
		Subfamily: KazhuoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Sanie = &Language{
		Name:      "sanie",
		Subfamily: KazhuoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Sadu = &Language{
		Name:      "sadu",
		Subfamily: KazhuoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Meuma = &Language{
		Name:      "meuma",
		Subfamily: KazhuoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Lisu = &Language{
		Name:      "lisu",
		Subfamily: LisoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Lolopo = &Language{
		Name:      "lolopo",
		Subfamily: LisoishSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
)

var AllSinoTibetanLanguages = []*Language{
	Chinese,
	Cantonese,
	Bai,
	Caijia,
	Longjia,
	Luren,
	Muangphe,
	Mango,
	Manga,
	Maang,
	Mondzi,
	Maza,
	Mauphu,
	Motang,
	Mongphu,
	Burmish,
	Jinuo,
	Sangkong,
	Bisu,
	Phunoi,
	Pyen,
	Sila,
	Phana,
	Akeu,
	Hani,
	Piyo,
	Enu,
	Mpi,
	Kaduo,
	Laha,
	Kucong,
	Namuyi,
	Shixing,
	Naish,
	Nusu,
	Zauzou,
	Katso,
	Samu,
	Sanie,
	Sadu,
	Meuma,
}

// Fantasy Languages
var (
	CommonEldarin = &Language{
		Name:      "common_eldarin",
		Subfamily: QuendianSubfamily,
		IsLiving:  true,
		WordBase:  SindarinWordBase, // @TODO can be changed after adding common_eldarin to word bases
	}
	Quenya = &Language{
		Name:      "quenya",
		Subfamily: QuendianSubfamily,
		IsLiving:  true,
		WordBase:  SindarinWordBase, // @TODO can be changed after adding quenya to word bases
	}
	Quendya = &Language{
		Name:      "quendya",
		Subfamily: QuendianSubfamily,
		IsLiving:  true,
		WordBase:  SindarinWordBase, // @TODO can be changed after adding quendya to word bases
	}
	ExilicQuendya = &Language{
		Name:      "exilic_quendya",
		Subfamily: QuendianSubfamily,
		IsLiving:  true,
		WordBase:  SindarinWordBase, // @TODO can be changed after adding exilic_quendya to word bases
	}
	CommonTelerin = &Language{
		Name:      "common_telerin",
		Subfamily: QuendianSubfamily,
		IsLiving:  true,
		WordBase:  SindarinWordBase, // @TODO can be changed after adding common_telerin to word bases
	}
	Telerin = &Language{
		Name:      "telerin",
		Subfamily: QuendianSubfamily,
		IsLiving:  true,
		WordBase:  SindarinWordBase, // @TODO can be changed after adding telerin to word bases
	}
	Sindarin = &Language{
		Name:      "sindarin",
		Subfamily: QuendianSubfamily,
		IsLiving:  true,
		WordBase:  SindarinWordBase,
	}
	Nandorin = &Language{
		Name:      "nandorin",
		Subfamily: QuendianSubfamily,
		IsLiving:  true,
		WordBase:  SindarinWordBase, // @TODO can be changed after adding nandorin to word bases
	}
	Avarin = &Language{
		Name:      "avarin",
		Subfamily: QuendianSubfamily,
		IsLiving:  true,
		WordBase:  SindarinWordBase, // @TODO can be changed after adding avarin to word bases
	}
	Aldmeris = &Language{
		Name:      "aldmeris",
		Subfamily: AldmerisSubfamily,
		IsLiving:  true,
		WordBase:  SindarinWordBase, // @TODO can be changed after adding aldmeris to word bases
	}
	Dunmeris = &Language{
		Name:      "dunmeris",
		Subfamily: AldmerisSubfamily,
		IsLiving:  true,
		WordBase:  DunmerisWordBase,
	}
	Bosmeris = &Language{
		Name:      "bosmeris",
		Subfamily: AldmerisSubfamily,
		IsLiving:  true,
		WordBase:  DunmerisWordBase, // @TODO can be changed after adding bosmeris to word bases
	}
	Dwemeris = &Language{
		Name:      "dwemeris",
		Subfamily: DwermerisSubfamily,
		IsLiving:  true,
		WordBase:  DwemerisWordBase,
	}
	Goblins = &Language{
		Name:      "goblins",
		Subfamily: OrcishSubfamily,
		IsLiving:  true,
		WordBase:  GoblinsWordBase,
	}
	Orcish = &Language{
		Name:      "orcish",
		Subfamily: OrcishSubfamily,
		IsLiving:  true,
		WordBase:  OrcWordBase,
	}
	Giantish = &Language{
		Name:      "giantish",
		Subfamily: GiantishSubfamily,
		IsLiving:  true,
		WordBase:  GiantWordBase,
	}
	Draconic = &Language{
		Name:      "draconic",
		Subfamily: DraconicSubfamily,
		IsLiving:  true,
		WordBase:  DraconicWordBase,
	}
	Arachnid = &Language{
		Name:      "arachnid",
		Subfamily: ArachnidSubfamily,
		IsLiving:  true,
		WordBase:  ArachnidWordBase,
	}
	Serpents = &Language{
		Name:      "serpents",
		Subfamily: SerpentsSubfamily,
		IsLiving:  true,
		WordBase:  SerpentsWordBase,
	}
)

var AllFantasyLanguages = []*Language{
	CommonEldarin,
	Quenya,
	Quendya,
	ExilicQuendya,
	CommonTelerin,
	Telerin,
	Sindarin,
	Nandorin,
	Avarin,
	Aldmeris,
	Dunmeris,
	Bosmeris,
	Dwemeris,
	Goblins,
	Orcish,
	Giantish,
	Draconic,
	Arachnid,
	Serpents,
}

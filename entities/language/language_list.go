package language

import "persons_generator/tools"

var AllLanguages = tools.Merge(
	AllIndoEuropeanLanguages,
	AllBasqueLanguages,
	AllEastAsianLanguages,
	AllKartvelianLanguages,
	AllSinoTibetanLanguages,
	AllAustroasiaticLanguages,
	AllUralicLanguages,
	AllTurkicLanguages,
	AllMongolicLanguages,
	AllTungusicLanguages,
	AllDravidianLanguages,
	AllUtoAztecanLanguages,
	AllEskimoAleutLanguages,
	AllAustonesianLanguages,
	AllAfricanLanguages,
	AllQuechuaLanguages,
	AllAfroAsiaticLanguages,
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
	Alamannian = &Language{
		Name:      "alamannian",
		Subfamily: ElbeGermanicSubfamily,
		IsLiving:  false,
		WordBase:  GermanWordBase, // @TODO can be changed after adding alamannian to word bases
	}
	Bavarian = &Language{
		Name:      "bavarian",
		Subfamily: ElbeGermanicSubfamily,
		IsLiving:  false,
		WordBase:  GermanWordBase, // @TODO can be changed after adding bavarian to word bases
	}
	Langobardian = &Language{
		Name:      "lombardian",
		Subfamily: ElbeGermanicSubfamily,
		IsLiving:  false,
		WordBase:  GermanWordBase, // @TODO can be changed after adding lombardian to word bases
	}
	Frankish = &Language{
		Name:      "frankish",
		Subfamily: WeserRhineGermanicSubfamily,
		IsLiving:  false,
		WordBase:  GermanWordBase, // @TODO can be changed after adding frankish to word bases
	}
	Anglic = &Language{
		Name:      "anglic",
		Subfamily: NorthSeaGermanicSubfamily,
		IsLiving:  false,
		WordBase:  GermanWordBase, // @TODO can be changed after adding anglic to word bases
	}
	Saxon = &Language{
		Name:      "saxon",
		Subfamily: NorthSeaGermanicSubfamily,
		IsLiving:  false,
		WordBase:  GermanWordBase, // @TODO can be changed after adding saxon to word bases
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
	Scythian = &Language{
		Name:      "scythian",
		Subfamily: IranianSubfamily,
		IsLiving:  false,
		WordBase:  IranianWordBase,
	}
	// Nuriastani
	KamkataVari = &Language{
		Name:      "kamkata_vari",
		Subfamily: NuriastaniSubfamily,
		WordBase:  KannadaWordBase, // @TODO change it into KamkataVari when available
	}
	VasiVari = &Language{
		Name:      "vasi_vari",
		Subfamily: NuriastaniSubfamily,
		WordBase:  KannadaWordBase, // @TODO change it into VasiVari when available
		IsLiving:  true,
	}
	Askunu = &Language{
		Name:      "askunu",
		Subfamily: NuriastaniSubfamily,
		WordBase:  KannadaWordBase, // @TODO change it into Askunu when available
		IsLiving:  true,
	}
	Waigali = &Language{
		Name:      "waigali",
		Subfamily: NuriastaniSubfamily,
		WordBase:  KannadaWordBase, // @TODO change it into Waigali when available
		IsLiving:  true,
	}
	Tregami = &Language{
		Name:      "tregami",
		Subfamily: NuriastaniSubfamily,
		WordBase:  KannadaWordBase, // @TODO change it into Tregami when available
		IsLiving:  true,
	}
	Zemiaki = &Language{
		Name:      "zemiaki",
		Subfamily: NuriastaniSubfamily,
		WordBase:  KannadaWordBase, // @TODO change it into Zemiaki when available
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
	// Tocharian
	Tocharian = &Language{
		Name:      "tocharian",
		Subfamily: TocharianSubfamily,
		WordBase:  ChineseWordBase,
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
	Alamannian,
	Bavarian,
	Langobardian,
	Frankish,
	Anglic,
	Saxon,
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
	Scythian,
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
	Tocharian,
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

// EastAsian
var (
	Japanese = &Language{
		Name:      "japanese",
		Subfamily: JaponicSubfamily,
		IsLiving:  true,
		WordBase:  JapaneseWordBase,
	}
	Korean = &Language{
		Name:      "korean",
		Subfamily: KoreanicSubfamily,
		IsLiving:  true,
		WordBase:  KoreanWordBase,
	}
)

var AllEastAsianLanguages = []*Language{Japanese, Korean}

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
	Burmese = &Language{
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
	Tangut = &Language{
		Name:      "tangut",
		Subfamily: QiangicSubfamily,
		WordBase:  ChineseWordBase,
		IsLiving:  true,
	}
	Tibetic = &Language{
		Name:      "tibetic",
		Subfamily: BodishSubfamily,
		IsLiving:  true,
		WordBase:  NigerianWordBase,
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
	Burmese,
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
	Lisu,
	Lolopo,
	Tangut,
	Tibetic,
}

// AustroasiaticFamily
var (
	Vietnamese = &Language{
		Name:      "vietnamese",
		Subfamily: VietMuongSubfamily,
		IsLiving:  true,
		WordBase:  VietnameseWordBase,
	}
	Kri = &Language{
		Name:      "kri",
		Subfamily: VieticSubfamily,
		IsLiving:  true,
		WordBase:  VietnameseWordBase,
	}
	Khmer = &Language{
		Name:      "khmer",
		Subfamily: KhmericSubfamily,
		IsLiving:  true,
		WordBase:  VietnameseWordBase,
	}
)

var AllAustroasiaticLanguages = []*Language{
	Vietnamese,
	Kri,
	Khmer,
}

// Uralic
var (
	Mari = &Language{
		Name:      "mari",
		Subfamily: FinnoPermicSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Merya = &Language{
		Name:      "merya",
		Subfamily: FinnoPermicSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Komi = &Language{
		Name:      "komi",
		Subfamily: PermicSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Permyak = &Language{
		Name:      "permyak",
		Subfamily: PermicSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Udmurt = &Language{
		Name:      "udmurt",
		Subfamily: PermicSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Meshchera = &Language{
		Name:      "meshchera",
		Subfamily: PermicSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Estonian = &Language{
		Name:      "estonian",
		Subfamily: BaltoFinnicSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Livonian = &Language{
		Name:      "livonian",
		Subfamily: BaltoFinnicSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Votic = &Language{
		Name:      "votic",
		Subfamily: BaltoFinnicSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Finnish = &Language{
		Name:      "finnish",
		Subfamily: BaltoFinnicSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Ingrian = &Language{
		Name:      "ingrian",
		Subfamily: BaltoFinnicSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Karelian = &Language{
		Name:      "karelian",
		Subfamily: BaltoFinnicSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Ludic = &Language{
		Name:      "ludic",
		Subfamily: BaltoFinnicSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Veps = &Language{
		Name:      "veps",
		Subfamily: BaltoFinnicSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Sami = &Language{
		Name:      "sami",
		Subfamily: SamiSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Erzya = &Language{
		Name:      "erzya",
		Subfamily: MordvinSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Moksha = &Language{
		Name:      "moksha",
		Subfamily: MordvinSubfamily,
		IsLiving:  true,
		WordBase:  FinnicWordBase,
	}
	Hungarian = &Language{
		Name:      "hungarian",
		Subfamily: UgricSubfamily,
		IsLiving:  true,
		WordBase:  HungarianWordBase,
	}
	Khanty = &Language{
		Name:      "khanty",
		Subfamily: UgricSubfamily,
		IsLiving:  true,
		WordBase:  HungarianWordBase,
	}
	Mansi = &Language{
		Name:      "mansi",
		Subfamily: UgricSubfamily,
		IsLiving:  true,
		WordBase:  HungarianWordBase,
	}
)

var AllUralicLanguages = []*Language{
	Mari,
	Merya,
	Komi,
	Permyak,
	Udmurt,
	Meshchera,
	Estonian,
	Livonian,
	Votic,
	Finnish,
	Ingrian,
	Karelian,
	Ludic,
	Veps,
	Sami,
	Erzya,
	Moksha,
	Hungarian,
	Khanty,
	Mansi,
}

// Turkic
var (
	Gagauz = &Language{
		Name:      "gagauz",
		Subfamily: OghuzSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Turkish = &Language{
		Name:      "turkish",
		Subfamily: OghuzSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Azerbaijani = &Language{
		Name:      "azerbaijani",
		Subfamily: OghuzSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Turkmen = &Language{
		Name:      "turkmen",
		Subfamily: OghuzSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Qashqai = &Language{
		Name:      "qashqai",
		Subfamily: OghuzSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Chaharmahali = &Language{
		Name:      "chaharmahali",
		Subfamily: OghuzSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Khorasani = &Language{
		Name:      "khorasani",
		Subfamily: OghuzSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Salar = &Language{
		Name:      "salar",
		Subfamily: OghuzSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Bashkir = &Language{
		Name:      "bashkir",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Tatar = &Language{
		Name:      "tatar",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	CrimeanTatar = &Language{
		Name:      "crimean_tatar",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	KarachayBalkar = &Language{
		Name:      "karachay_balkar",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Kumyk = &Language{
		Name:      "kumyk",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Karaim = &Language{
		Name:      "karaim",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Krymchak = &Language{
		Name:      "krymchak",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Urum = &Language{
		Name:      "urum",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Kazakh = &Language{
		Name:      "kazakh",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Karakalpak = &Language{
		Name:      "karakalpak",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Nogai = &Language{
		Name:      "nogai",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	SiberianTatar = &Language{
		Name:      "siberian_tatar",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Baraba = &Language{
		Name:      "baraba",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	SouthernAltai = &Language{
		Name:      "southern_altai",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Teleut = &Language{
		Name:      "teleut",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Kyrgyz = &Language{
		Name:      "kyrgyz",
		Subfamily: KipchakSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Uzbek = &Language{
		Name:      "uzbek",
		Subfamily: KarlukSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Uyghur = &Language{
		Name:      "uyghur",
		Subfamily: KarlukSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Aynu = &Language{
		Name:      "aynu",
		Subfamily: KarlukSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Ili = &Language{
		Name:      "ili",
		Subfamily: KarlukSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Khalaj = &Language{
		Name:      "khalaj",
		Subfamily: ShazTurkicSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Chuvash = &Language{
		Name:      "chuvash",
		Subfamily: OghuzSubfamily,
		IsLiving:  true,
		WordBase:  TurkishWordBase,
	}
	Bulgar = &Language{
		Name:      "bulgar",
		Subfamily: OghuzSubfamily,
		IsLiving:  false,
		WordBase:  TurkishWordBase,
	}
	Sabir = &Language{
		Name:      "sabir",
		Subfamily: OghuzSubfamily,
		IsLiving:  false,
		WordBase:  TurkishWordBase,
	}
	Khazar = &Language{
		Name:      "khazar",
		Subfamily: OghuzSubfamily,
		IsLiving:  false,
		WordBase:  TurkishWordBase,
	}
	Hunnic = &Language{
		Name:      "hunnic",
		Subfamily: OghuzSubfamily,
		IsLiving:  false,
		WordBase:  TurkishWordBase,
	}
	Tuoba = &Language{
		Name:      "tuoba",
		Subfamily: OghuzSubfamily,
		IsLiving:  false,
		WordBase:  TurkishWordBase,
	}
	Avar = &Language{
		Name:      "avar",
		Subfamily: OghuzSubfamily,
		IsLiving:  false,
		WordBase:  TurkishWordBase,
	}
	Cuman = &Language{
		Name:      "cuman",
		Subfamily: KipchakSubfamily,
		IsLiving:  false,
		WordBase:  TurkishWordBase,
	}
)

var AllTurkicLanguages = []*Language{
	Gagauz,
	Turkish,
	Azerbaijani,
	Turkmen,
	Qashqai,
	Chaharmahali,
	Khorasani,
	Salar,
	Bashkir,
	Tatar,
	CrimeanTatar,
	KarachayBalkar,
	Kumyk,
	Karaim,
	Krymchak,
	Urum,
	Kazakh,
	Karakalpak,
	Nogai,
	SiberianTatar,
	Baraba,
	SouthernAltai,
	Teleut,
	Kyrgyz,
	Uzbek,
	Uyghur,
	Aynu,
	Ili,
	Khalaj,
	Chuvash,
	Bulgar,
	Sabir,
	Khazar,
	Hunnic,
	Tuoba,
	Avar,
	Cuman,
}

// Mongolic
var (
	Dagur = &Language{
		Name:      "dagur",
		Subfamily: DagurSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Moghol = &Language{
		Name:      "moghol",
		Subfamily: MogholSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Khamnigan = &Language{
		Name:      "khamnigan",
		Subfamily: CentralMongolicSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Buryat = &Language{
		Name:      "buryat",
		Subfamily: CentralMongolicSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Mongolian = &Language{
		Name:      "mongolian",
		Subfamily: CentralMongolicSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Khalkha = &Language{
		Name:      "khalkha",
		Subfamily: CentralMongolicSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Chakhar = &Language{
		Name:      "chakhar",
		Subfamily: CentralMongolicSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Khorchin = &Language{
		Name:      "khorchin",
		Subfamily: CentralMongolicSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Ordos = &Language{
		Name:      "ordos",
		Subfamily: CentralMongolicSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Oirat = &Language{
		Name:      "oirat",
		Subfamily: CentralMongolicSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Monguor = &Language{
		Name:      "monguor",
		Subfamily: SouthernMongolicSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Bonan = &Language{
		Name:      "bonan",
		Subfamily: SouthernMongolicSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Santa = &Language{
		Name:      "santa",
		Subfamily: SouthernMongolicSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Kangjia = &Language{
		Name:      "kangjia",
		Subfamily: SouthernMongolicSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Khitan = &Language{
		Name:      "khitan",
		Subfamily: KhitanSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
)

var AllMongolicLanguages = []*Language{
	Dagur,
	Moghol,
	Khamnigan,
	Buryat,
	Mongolian,
	Khalkha,
	Chakhar,
	Khorchin,
	Ordos,
	Oirat,
	Monguor,
	Bonan,
	Santa,
	Kangjia,
	Khitan,
}

// Tungusic

var (
	Manchu = &Language{
		Name:      "manchu",
		Subfamily: JurchenicSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
	Xibe = &Language{
		Name:      "xibe",
		Subfamily: JurchenicSubfamily,
		IsLiving:  true,
		WordBase:  MongolianWordBase,
	}
)

var AllTungusicLanguages = []*Language{Manchu, Xibe}

// Dravidian
var (
	Tamil = &Language{
		Name:      "tamil",
		Subfamily: SouthDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Malayalam = &Language{
		Name:      "malayalam",
		Subfamily: SouthDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Irula = &Language{
		Name:      "irula",
		Subfamily: SouthDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Kodava = &Language{
		Name:      "kodava",
		Subfamily: SouthDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Kurumba = &Language{
		Name:      "kurumba",
		Subfamily: SouthDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Toda = &Language{
		Name:      "toda",
		Subfamily: SouthDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Kota = &Language{
		Name:      "kota",
		Subfamily: SouthDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Kannada = &Language{
		Name:      "kannada",
		Subfamily: SouthDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Badaga = &Language{
		Name:      "badaga",
		Subfamily: SouthDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Telugu = &Language{
		Name:      "telugu",
		Subfamily: SouthCentralDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Chenchu = &Language{
		Name:      "chenchu",
		Subfamily: SouthCentralDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Konda = &Language{
		Name:      "konda",
		Subfamily: SouthCentralDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	MukhaDora = &Language{
		Name:      "mukha_dora",
		Subfamily: SouthDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Manda = &Language{
		Name:      "manda",
		Subfamily: SouthCentralDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Pengo = &Language{
		Name:      "pengo",
		Subfamily: SouthCentralDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Kuvi = &Language{
		Name:      "kuvi",
		Subfamily: SouthCentralDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Kui = &Language{
		Name:      "kui",
		Subfamily: SouthCentralDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Gondi = &Language{
		Name:      "gondi",
		Subfamily: SouthCentralDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Kolami = &Language{
		Name:      "kolami",
		Subfamily: CentralDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Naiki = &Language{
		Name:      "naiki",
		Subfamily: CentralDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Gadaba = &Language{
		Name:      "gadaba",
		Subfamily: CentralDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Ollari = &Language{
		Name:      "ollari",
		Subfamily: CentralDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Kondekor = &Language{
		Name:      "kondekor",
		Subfamily: CentralDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Duruwa = &Language{
		Name:      "duruwa",
		Subfamily: CentralDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Oraon = &Language{
		Name:      "oraon",
		Subfamily: NorthernDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Kisan = &Language{
		Name:      "kisan",
		Subfamily: NorthernDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	KumarbhagPaharia = &Language{
		Name:      "kumarbhag_paharia",
		Subfamily: NorthernDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	SauriaPaharia = &Language{
		Name:      "sauria_paharia",
		Subfamily: NorthernDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
	Brahui = &Language{
		Name:      "brahui",
		Subfamily: NorthernDravidianSubfamily,
		IsLiving:  true,
		WordBase:  KannadaWordBase,
	}
)

var AllDravidianLanguages = []*Language{
	Tamil,
	Malayalam,
	Irula,
	Kodava,
	Kurumba,
	Toda,
	Kota,
	Kannada,
	Badaga,
	Telugu,
	Chenchu,
	Konda,
	MukhaDora,
	Manda,
	Pengo,
	Kuvi,
	Kui,
	Gondi,
	Kolami,
	Naiki,
	Gadaba,
	Ollari,
	Kondekor,
	Duruwa,
	Oraon,
	Kisan,
	KumarbhagPaharia,
	SauriaPaharia,
	Brahui,
}

// African
var (
	Kru = &Language{
		Name:      "kru",
		Subfamily: KruSubfamily,
		IsLiving:  true,
		WordBase:  NigerianWordBase,
	}
	Kwa = &Language{
		Name:      "kwa",
		Subfamily: VoltaCongoSubfamily,
		IsLiving:  true,
		WordBase:  NigerianWordBase,
	}
	Gur = &Language{
		Name:      "gur",
		Subfamily: VoltaCongoSubfamily,
		IsLiving:  true,
		WordBase:  NigerianWordBase,
	}
	Soninke = &Language{
		Name:      "soninke",
		Subfamily: MandeSubfamily,
		IsLiving:  true,
		WordBase:  NigerianWordBase,
	}
	Manding = &Language{
		Name:      "manding",
		Subfamily: MandeSubfamily,
		IsLiving:  true,
		WordBase:  NigerianWordBase,
	}
	Senegambian = &Language{
		Name:      "senegambian",
		Subfamily: AtlanticCongoSubfamily,
		IsLiving:  true,
		WordBase:  NigerianWordBase,
	}
	Swahili = &Language{
		Name:      "swahili",
		Subfamily: AtlanticCongoSubfamily,
		IsLiving:  true,
		WordBase:  SwahiliWordBase,
	}
	Tebu = &Language{
		Name:      "tebu",
		Subfamily: SaharanSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Hausa = &Language{
		Name:      "hausa",
		Subfamily: WestChadicSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Nobiin = &Language{
		Name:      "nobiin",
		Subfamily: NubianSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Kenzi = &Language{
		Name:      "kenzi",
		Subfamily: NubianSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Midob = &Language{
		Name:      "midob",
		Subfamily: NubianSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Birgid = &Language{
		Name:      "birgid",
		Subfamily: NubianSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	HillNubian = &Language{
		Name:      "hill_nubian",
		Subfamily: NubianSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Daju = &Language{
		Name:      "daju",
		Subfamily: DajuSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Somali = &Language{
		Name:      "somali",
		Subfamily: SomalicSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	// Songhay
	Korandje = &Language{
		Name:      "korandje",
		Subfamily: SonghaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	KoyraChiini = &Language{
		Name:      "koyra_chiini",
		Subfamily: SonghaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Tadaksahak = &Language{
		Name:      "tadaksahak",
		Subfamily: SonghaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Tasawaq = &Language{
		Name:      "tasawaq",
		Subfamily: SonghaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Tagdal = &Language{
		Name:      "tagdal",
		Subfamily: SonghaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	TondiSongwayKiini = &Language{
		Name:      "tondi_songway_kiini",
		Subfamily: SonghaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	HumburiSenni = &Language{
		Name:      "humburi_senni",
		Subfamily: SonghaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	KoyraboroSenni = &Language{
		Name:      "koyraboro_senni",
		Subfamily: SonghaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Zarma = &Language{
		Name:      "zarma",
		Subfamily: SonghaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	SonghoyboroCiine = &Language{
		Name:      "songhoyboro_ciine",
		Subfamily: SonghaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Dendi = &Language{
		Name:      "dendi",
		Subfamily: SonghaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
)

var AllAfricanLanguages = []*Language{
	Kru,
	Kwa,
	Gur,
	Soninke,
	Manding,
	Senegambian,
	Swahili,
	Tebu,
	Hausa,
	Nobiin,
	Kenzi,
	Midob,
	Birgid,
	HillNubian,
	Daju,
	Somali,
	Korandje,
	KoyraChiini,
	Tadaksahak,
	Tasawaq,
	Tagdal,
	TondiSongwayKiini,
	HumburiSenni,
	KoyraboroSenni,
	Zarma,
	SonghoyboroCiine,
	Dendi,
}

// UtoAztecan
var (
	Oodham = &Language{
		Name:      "oodham",
		Subfamily: TepimanSubfamily,
		IsLiving:  true,
		WordBase:  NahuatlWordBase,
	}
	PimaBajo = &Language{
		Name:      "pima_bajo",
		Subfamily: TepimanSubfamily,
		IsLiving:  true,
		WordBase:  NahuatlWordBase,
	}
	Tepehuan = &Language{
		Name:      "tepehuan",
		Subfamily: TepimanSubfamily,
		IsLiving:  true,
		WordBase:  NahuatlWordBase,
	}
	Southern = &Language{
		Name:      "southern",
		Subfamily: TepimanSubfamily,
		IsLiving:  true,
		WordBase:  NahuatlWordBase,
	}
	Tepecano = &Language{
		Name:      "tepecano",
		Subfamily: TepimanSubfamily,
		IsLiving:  false,
		WordBase:  NahuatlWordBase,
	}
	Tarahumara = &Language{
		Name:      "tarahumara",
		Subfamily: TarahumaranSubfamily,
		IsLiving:  true,
		WordBase:  NahuatlWordBase,
	}
	UpriverGuarijio = &Language{
		Name:      "upriver_guarijio",
		Subfamily: TarahumaranSubfamily,
		IsLiving:  true,
		WordBase:  NahuatlWordBase,
	}
	DownriverGuarijio = &Language{
		Name:      "downriver_guarijio",
		Subfamily: TarahumaranSubfamily,
		IsLiving:  true,
		WordBase:  NahuatlWordBase,
	}
	Tubar = &Language{
		Name:      "tubar",
		Subfamily: TarahumaranSubfamily,
		IsLiving:  false,
		WordBase:  NahuatlWordBase,
	}
	Yaqui = &Language{
		Name:      "yaqui",
		Subfamily: CahitaSubfamily,
		IsLiving:  true,
		WordBase:  NahuatlWordBase,
	}
	Mayo = &Language{
		Name:      "mayo",
		Subfamily: CahitaSubfamily,
		IsLiving:  true,
		WordBase:  NahuatlWordBase,
	}
	Opata = &Language{
		Name:      "opata",
		Subfamily: OpatanSubfamily,
		IsLiving:  false,
		WordBase:  NahuatlWordBase,
	}
	Eudeve = &Language{
		Name:      "eudeve",
		Subfamily: OpatanSubfamily,
		IsLiving:  false,
		WordBase:  NahuatlWordBase,
	}
	Cora = &Language{
		Name:      "Cora",
		Subfamily: CoracholSubfamily,
		IsLiving:  true,
		WordBase:  NahuatlWordBase,
	}
	Huichol = &Language{
		Name:      "Huichol",
		Subfamily: CoracholSubfamily,
		IsLiving:  true,
		WordBase:  NahuatlWordBase,
	}
	Pochutec = &Language{
		Name:      "pochutec",
		Subfamily: AztecanSubfamily,
		IsLiving:  false,
		WordBase:  NahuatlWordBase,
	}
	Pipil = &Language{
		Name:      "pipil",
		Subfamily: AztecanSubfamily,
		IsLiving:  true,
		WordBase:  NahuatlWordBase,
	}
	Nahuatl = &Language{
		Name:      "nahuatl",
		Subfamily: AztecanSubfamily,
		IsLiving:  true,
		WordBase:  NahuatlWordBase,
	}
)

var AllUtoAztecanLanguages = []*Language{
	Oodham,
	PimaBajo,
	Tepehuan,
	Southern,
	Tepecano,
	Tarahumara,
	UpriverGuarijio,
	DownriverGuarijio,
	Tubar,
	Yaqui,
	Mayo,
	Opata,
	Eudeve,
	Cora,
	Huichol,
	Pochutec,
	Pipil,
	Nahuatl,
}

// EskimoAleut
var Inuit = &Language{
	Name:      "inuit",
	Subfamily: EskimoSubfamily,
	IsLiving:  true,
	WordBase:  InuitWordBase,
}

var AllEskimoAleutLanguages = []*Language{Inuit}

// Austonesian
var (
	Hawaiian = &Language{
		Name:      "hawaiian",
		Subfamily: MarquesicSubfamily,
		IsLiving:  true,
		WordBase:  HawaiianWordBase,
	}
)

var AllAustonesianLanguages = []*Language{Hawaiian}

// Quechua
var (
	Ancash = &Language{
		Name:      "Ancash",
		Subfamily: QuechuaIISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	AltoPativilcaAltoMaranonAltoHuallaga = &Language{
		Name:      "alto_pativilca_alto_maranon_alto_huallaga",
		Subfamily: QuechuaISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	Yaru = &Language{
		Name:      "yaru",
		Subfamily: QuechuaISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	Wanka = &Language{
		Name:      "wanka",
		Subfamily: QuechuaISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	YauyosChincha = &Language{
		Name:      "yauyos_chincha",
		Subfamily: QuechuaISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	Pacaraos = &Language{
		Name:      "pacaraos",
		Subfamily: QuechuaISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	Lambayeque = &Language{
		Name:      "lambayeque",
		Subfamily: QuechuaIISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	Cajamarca = &Language{
		Name:      "cajamarca",
		Subfamily: QuechuaIISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	Lincha = &Language{
		Name:      "lincha",
		Subfamily: QuechuaIISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	Laraos = &Language{
		Name:      "laraos",
		Subfamily: QuechuaIISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	Ayacucho = &Language{
		Name:      "ayacucho",
		Subfamily: QuechuaIISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	Cusco = &Language{
		Name:      "cusco",
		Subfamily: QuechuaIISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	Puno = &Language{
		Name:      "puno",
		Subfamily: QuechuaIISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	NorthernBolivian = &Language{
		Name:      "northern_bolivian",
		Subfamily: QuechuaIISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	SouthernBolivia = &Language{
		Name:      "southern_bolivia",
		Subfamily: QuechuaIISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
	SantiagoDelEstero = &Language{
		Name:      "santiago_del_estero",
		Subfamily: QuechuaIISubfamily,
		IsLiving:  true,
		WordBase:  QuechuaWordBase,
	}
)

var AllQuechuaLanguages = []*Language{
	Ancash,
	AltoPativilcaAltoMaranonAltoHuallaga,
	Yaru,
	Wanka,
	YauyosChincha,
	Pacaraos,
	Lambayeque,
	Cajamarca,
	Lincha,
	Laraos,
	Ayacucho,
	Cusco,
	Puno,
	NorthernBolivian,
	SouthernBolivia,
	SantiagoDelEstero,
}

// AfroAsiaticFamily
var (
	// Berber
	Kabyle = &Language{
		Name:      "kabyle",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Tamazight = &Language{
		Name:      "tamazight",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Shilha = &Language{
		Name:      "shilha",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	SenhajaDeSrair = &Language{
		Name:      "senhaja_de_srair",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Ghomara = &Language{
		Name:      "ghomara",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Riffian = &Language{
		Name:      "riffian",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	AytSeghrouchen = &Language{
		Name:      "ayt_seghrouchen",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	AytWarayn = &Language{
		Name:      "ayt_warayn",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Shenwa = &Language{
		Name:      "shenwa",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Shawiya = &Language{
		Name:      "shawiya",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Zenaga = &Language{
		Name:      "zenaga",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Siwi = &Language{
		Name:      "siwi",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Nafusi = &Language{
		Name:      "nafusi",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Sokna = &Language{
		Name:      "sokna",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Ghadames = &Language{
		Name:      "ghadames",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Awjila = &Language{
		Name:      "awjila",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Tuareg = &Language{
		Name:      "tuareg",
		Subfamily: BerberSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	// Cushitic
	Beja = &Language{
		Name:      "beja",
		Subfamily: CushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Awngi = &Language{
		Name:      "awngi",
		Subfamily: AgawSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Bilen = &Language{
		Name:      "bilen",
		Subfamily: AgawSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Qimant = &Language{
		Name:      "qimant",
		Subfamily: AgawSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Xamtanga = &Language{
		Name:      "xamtanga",
		Subfamily: AgawSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Gawwada = &Language{
		Name:      "gawwada",
		Subfamily: DullaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Tsamai = &Language{
		Name:      "tsamai",
		Subfamily: DullaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Dihina = &Language{
		Name:      "dihina",
		Subfamily: DullaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Dobase = &Language{
		Name:      "dobase",
		Subfamily: DullaySubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Burji = &Language{
		Name:      "burji",
		Subfamily: HighlandEastCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Sidamo = &Language{
		Name:      "sidamo",
		Subfamily: HighlandEastCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Gedeo = &Language{
		Name:      "gedeo",
		Subfamily: HighlandEastCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	HadiyyaLibido = &Language{
		Name:      "hadiyya_libido",
		Subfamily: HighlandEastCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	KambaataAlaba = &Language{
		Name:      "kambaata_alaba",
		Subfamily: HighlandEastCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	SahoAfar = &Language{
		Name:      "saho_afar",
		Subfamily: LowlandEastCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	OmoTana = &Language{
		Name:      "omo_tana",
		Subfamily: LowlandEastCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Oromoid = &Language{
		Name:      "oromoid",
		Subfamily: LowlandEastCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Dullay = &Language{
		Name:      "dullay",
		Subfamily: LowlandEastCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Yaaku = &Language{
		Name:      "yaaku",
		Subfamily: LowlandEastCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Gorowa = &Language{
		Name:      "gorowa",
		Subfamily: SouthCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Iraqw = &Language{
		Name:      "iraqw",
		Subfamily: SouthCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Alagwa = &Language{
		Name:      "alagwa",
		Subfamily: SouthCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Burunge = &Language{
		Name:      "burunge",
		Subfamily: SouthCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Aasax = &Language{
		Name:      "aasax",
		Subfamily: SouthCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Kwadza = &Language{
		Name:      "kwadza",
		Subfamily: SouthCushiticSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	Egyptian = &Language{
		Name:      "egyptian",
		Subfamily: EgyptianSubfamily,
		IsLiving:  true,
		WordBase:  BerberWordBase,
	}
	// semitic
	Akkadian = &Language{
		Name:      "akkadian",
		Subfamily: EastSemiticSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Eblaite = &Language{
		Name:      "eblaite",
		Subfamily: EastSemiticSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Kishite = &Language{
		Name:      "kishite",
		Subfamily: EastSemiticSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Ashkenazi = &Language{
		Name:      "ashkenazi",
		Subfamily: HebrewSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Sephardi = &Language{
		Name:      "sephardi",
		Subfamily: HebrewSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Hebrew = &Language{
		Name:      "hebrew",
		Subfamily: HebrewSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Samaritan = &Language{
		Name:      "samaritan",
		Subfamily: HebrewSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Ammonite = &Language{
		Name:      "ammonite",
		Subfamily: CanaaniteSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Edomite = &Language{
		Name:      "edomite",
		Subfamily: CanaaniteSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Moabite = &Language{
		Name:      "moabite",
		Subfamily: CanaaniteSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Phoenician = &Language{
		Name:      "phoenician",
		Subfamily: CanaaniteSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Samalian = &Language{
		Name:      "samalian",
		Subfamily: CanaaniteSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Aramaic = &Language{
		Name:      "aramaic",
		Subfamily: AramaicSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Nabataean = &Language{
		Name:      "nabataean",
		Subfamily: AramaicSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Amorite = &Language{
		Name:      "amorite",
		Subfamily: AramaicSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Himyaritic = &Language{
		Name:      "himyaritic",
		Subfamily: AramaicSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Sutean = &Language{
		Name:      "sutean",
		Subfamily: AramaicSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Syriac = &Language{
		Name:      "syriac",
		Subfamily: AramaicSubfamily,
		IsLiving:  true,
		WordBase:  MesopotamianWordBase,
	}
	Arabic = &Language{
		Name:      "arabic",
		Subfamily: ArabicSubfamily,
		IsLiving:  true,
		WordBase:  ArabicWordBase,
	}
	Geez = &Language{
		Name:      "geez",
		Subfamily: EthiopicSubfamily,
		IsLiving:  true,
		WordBase:  ArabicWordBase,
	}
	Tigrinya = &Language{
		Name:      "tigrinya",
		Subfamily: EthiopicSubfamily,
		IsLiving:  true,
		WordBase:  ArabicWordBase,
	}
	Amharic = &Language{
		Name:      "amharic",
		Subfamily: EthiopicSubfamily,
		IsLiving:  true,
		WordBase:  ArabicWordBase,
	}
	Argobba = &Language{
		Name:      "argobba",
		Subfamily: EthiopicSubfamily,
		IsLiving:  true,
		WordBase:  ArabicWordBase,
	}
	Harari = &Language{
		Name:      "harari",
		Subfamily: EthiopicSubfamily,
		IsLiving:  true,
		WordBase:  ArabicWordBase,
	}
	Gafat = &Language{
		Name:      "gafat",
		Subfamily: EthiopicSubfamily,
		IsLiving:  true,
		WordBase:  ArabicWordBase,
	}
	Soddo = &Language{
		Name:      "soddo",
		Subfamily: EthiopicSubfamily,
		IsLiving:  true,
		WordBase:  ArabicWordBase,
	}
	Mesmes = &Language{
		Name:      "mesmes",
		Subfamily: EthiopicSubfamily,
		IsLiving:  true,
		WordBase:  ArabicWordBase,
	}
	Muher = &Language{
		Name:      "muher",
		Subfamily: EthiopicSubfamily,
		IsLiving:  true,
		WordBase:  ArabicWordBase,
	}
	WestGurage = &Language{
		Name:      "west_gurage",
		Subfamily: EthiopicSubfamily,
		IsLiving:  true,
		WordBase:  ArabicWordBase,
	}
	Mesqan = &Language{
		Name:      "mesqan",
		Subfamily: EthiopicSubfamily,
		IsLiving:  true,
		WordBase:  ArabicWordBase,
	}
	SebatBet = &Language{
		Name:      "sebat_bet",
		Subfamily: EthiopicSubfamily,
		IsLiving:  true,
		WordBase:  ArabicWordBase,
	}
	Baṭḥari = &Language{
		Name:      "bathari",
		IsLiving:  true,
		Subfamily: SouthArabianSubfamily,
		WordBase:  ArabicWordBase,
	}
	Harsusi = &Language{
		Name:      "harsusi",
		IsLiving:  true,
		Subfamily: SouthArabianSubfamily,
		WordBase:  ArabicWordBase,
	}
	Hobyot = &Language{
		Name:      "hobyot",
		IsLiving:  true,
		Subfamily: SouthArabianSubfamily,
		WordBase:  ArabicWordBase,
	}
	Mehri = &Language{
		Name:      "mehri",
		IsLiving:  true,
		Subfamily: SouthArabianSubfamily,
		WordBase:  ArabicWordBase,
	}
	Shehri = &Language{
		Name:      "shehri",
		IsLiving:  true,
		Subfamily: SouthArabianSubfamily,
		WordBase:  ArabicWordBase,
	}
	Soqotri = &Language{
		Name:      "soqotri",
		IsLiving:  true,
		Subfamily: SouthArabianSubfamily,
		WordBase:  ArabicWordBase,
	}
	Faifi = &Language{
		Name:      "faifi",
		IsLiving:  true,
		Subfamily: SouthArabianSubfamily,
		WordBase:  ArabicWordBase,
	}
	Hadramautic = &Language{
		Name:      "hadramautic",
		IsLiving:  true,
		Subfamily: SouthArabianSubfamily,
		WordBase:  ArabicWordBase,
	}
	Minaean = &Language{
		Name:      "minaean",
		IsLiving:  true,
		Subfamily: SouthArabianSubfamily,
		WordBase:  ArabicWordBase,
	}
	Qatabanian = &Language{
		Name:      "qatabanian",
		IsLiving:  true,
		Subfamily: SouthArabianSubfamily,
		WordBase:  ArabicWordBase,
	}
	Razihi = &Language{
		Name:      "razihi",
		IsLiving:  true,
		Subfamily: SouthArabianSubfamily,
		WordBase:  ArabicWordBase,
	}
	Sabaean = &Language{
		Name:      "sabaean",
		IsLiving:  true,
		Subfamily: SouthArabianSubfamily,
		WordBase:  ArabicWordBase,
	}
	// Omotic
	SouthOmotic = &Language{
		Name:      "south_omotic",
		IsLiving:  true,
		Subfamily: OmoticSubfamily,
		WordBase:  ArabicWordBase,
	}
	Mao = &Language{
		Name:      "mao",
		IsLiving:  true,
		Subfamily: OmoticSubfamily,
		WordBase:  ArabicWordBase,
	}
	Dizoid = &Language{
		Name:      "dizoid",
		IsLiving:  true,
		Subfamily: OmoticSubfamily,
		WordBase:  ArabicWordBase,
	}
	Gonga = &Language{
		Name:      "gonga",
		IsLiving:  true,
		Subfamily: OmoticSubfamily,
		WordBase:  ArabicWordBase,
	}
)

var AllAfroAsiaticLanguages = []*Language{
	Kabyle,
	Tamazight,
	Shilha,
	SenhajaDeSrair,
	Ghomara,
	Riffian,
	AytSeghrouchen,
	AytWarayn,
	Shenwa,
	Shawiya,
	Zenaga,
	Siwi,
	Nafusi,
	Sokna,
	Ghadames,
	Awjila,
	Tuareg,
	Beja,
	Awngi,
	Bilen,
	Qimant,
	Xamtanga,
	Gawwada,
	Tsamai,
	Dihina,
	Dobase,
	Burji,
	Sidamo,
	Gedeo,
	HadiyyaLibido,
	KambaataAlaba,
	SahoAfar,
	OmoTana,
	Oromoid,
	Dullay,
	Yaaku,
	Gorowa,
	Iraqw,
	Alagwa,
	Burunge,
	Aasax,
	Kwadza,
	Egyptian,
	Akkadian,
	Eblaite,
	Kishite,
	Ashkenazi,
	Sephardi,
	Hebrew,
	Samaritan,
	Ammonite,
	Edomite,
	Moabite,
	Phoenician,
	Samalian,
	Aramaic,
	Nabataean,
	Amorite,
	Himyaritic,
	Sutean,
	Syriac,
	Arabic,
	Geez,
	Tigrinya,
	Amharic,
	Argobba,
	Harari,
	Gafat,
	Soddo,
	Mesmes,
	Muher,
	WestGurage,
	Mesqan,
	SebatBet,
	Baṭḥari,
	Harsusi,
	Hobyot,
	Mehri,
	Shehri,
	Soqotri,
	Faifi,
	Hadramautic,
	Minaean,
	Qatabanian,
	Razihi,
	Sabaean,
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

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
		Name:        "albanian",
		Subfamily:   AlbanianSubfamily,
		IsLiving:    true,
		WordBaseRef: AlbanianWordBaseRef,
	}
	Armenian = &Language{
		Name:        "armenian",
		Subfamily:   ArmenianSubfamily,
		IsLiving:    true,
		WordBaseRef: ArmenianWordBaseRef,
	}
	// Slavic
	Ruthenian = &Language{
		Name:        "ruthenian",
		Subfamily:   RuthenianSubfamily,
		IsLiving:    false,
		WordBaseRef: RuthenianWordBaseRef,
	}
	Belarusian = &Language{
		Name:        "belarusian",
		Subfamily:   RuthenianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding belarusian to word bases
	}
	Ukrainian = &Language{
		Name:        "ukrainian",
		Subfamily:   RuthenianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding ukrainian to word bases
	}
	Russian = &Language{
		Name:        "russian",
		Subfamily:   MoscovianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding russian to word bases
	}
	Slovenian = &Language{
		Name:        "slovenian",
		Subfamily:   WesternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding slovenian to word bases
	}
	Croatian = &Language{
		Name:        "croation",
		Subfamily:   WesternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding croation to word bases
	}
	Serbian = &Language{
		Name:        "serbian",
		Subfamily:   WesternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding serbian to word bases
	}
	Bosnian = &Language{
		Name:        "bosnian",
		Subfamily:   WesternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding bosnian to word bases
	}
	Bulgarian = &Language{
		Name:        "bulgarian",
		Subfamily:   EasternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding bulgarian to word bases
	}
	Macedonian = &Language{
		Name:        "macedonian",
		Subfamily:   EasternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding macedonian to word bases
	}
	Polish = &Language{
		Name:        "polish",
		Subfamily:   LechiticSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding polish to word bases
	}
	Silesian = &Language{
		Name:        "silesian",
		Subfamily:   LechiticSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding silesian to word bases
	}
	LowerSorbian = &Language{
		Name:        "lower_sorbian",
		Subfamily:   SorbianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding lower_sorbian to word bases
	}
	UpperSorbian = &Language{
		Name:        "upper_sorbian",
		Subfamily:   SorbianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding upper_sorbian to word bases
	}
	Czech = &Language{
		Name:        "czech",
		Subfamily:   CzechSlovakSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding czech to word bases
	}
	Slovak = &Language{
		Name:        "slovak",
		Subfamily:   CzechSlovakSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding slovak to word bases
	}
	// Baltic
	Latvian = &Language{
		Name:        "latvian",
		Subfamily:   EasternBalticSubfamily,
		IsLiving:    true,
		WordBaseRef: LithuanianWordBaseRef, // @TODO can be changed after adding latvian to word bases
	}
	Latgalian = &Language{
		Name:        "latgalian",
		Subfamily:   EasternBalticSubfamily,
		IsLiving:    true,
		WordBaseRef: LithuanianWordBaseRef, // @TODO can be changed after adding latgalian to word bases
	}
	Lithuanian = &Language{
		Name:        "lithuanian",
		Subfamily:   EasternBalticSubfamily,
		IsLiving:    true,
		WordBaseRef: LithuanianWordBaseRef,
	}
	Semigallian = &Language{
		Name:        "semigallian",
		Subfamily:   EasternBalticSubfamily,
		IsLiving:    false,
		WordBaseRef: LithuanianWordBaseRef, // @TODO can be changed after adding semigallian to word bases
	}
	// Celtic
	Celtiberian = &Language{
		Name:        "celtiberian",
		Subfamily:   CelticSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Gallaecian = &Language{
		Name:        "gallaecian",
		Subfamily:   CelticSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Noric = &Language{
		Name:        "noric",
		Subfamily:   CelticSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Pictish = &Language{
		Name:        "pictish",
		Subfamily:   BrittonicSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Cumbric = &Language{
		Name:        "cumbric",
		Subfamily:   BrittonicSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Welsh = &Language{
		Name:        "welsh",
		Subfamily:   WesternBrittonicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	Cornish = &Language{
		Name:        "cornish",
		Subfamily:   SouthWesternBrittonicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	Breton = &Language{
		Name:        "breton",
		Subfamily:   SouthWesternBrittonicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	Irish = &Language{
		Name:        "irish",
		Subfamily:   GoidelicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	Manx = &Language{
		Name:        "manx",
		Subfamily:   GoidelicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	ScottishGaelic = &Language{
		Name:        "scottish_gaelic",
		Subfamily:   GoidelicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	// Dacian
	Dacian = &Language{
		Name:      "dacian",
		Subfamily: DacianSubfamily,
		IsLiving:  false,
	}
	// Germanic
	Danish = &Language{
		Name:        "danish",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding danish to word bases
	}
	Faroese = &Language{
		Name:        "faroese",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding faroese to word bases
	}
	Icelandic = &Language{
		Name:        "icelandic",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding icelandic to word bases
	}
	Norwegian = &Language{
		Name:        "norwegian",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding norwegian to word bases
	}
	Swedish = &Language{
		Name:        "swedish",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef,
	}
	Dalecarlian = &Language{
		Name:        "dalecarlian",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding dalecarlian to word bases
	}
	Gutnish = &Language{
		Name:        "gutnish",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding gutnish to word bases
	}
	English = &Language{
		Name:        "english",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: EnglishWordBaseRef,
	}
	Scots = &Language{
		Name:        "scots",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: EnglishWordBaseRef, // @TODO can be changed after adding scots to word bases
	}
	Yola = &Language{
		Name:        "yola",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: EnglishWordBaseRef, // @TODO can be changed after adding yola to word bases
	}
	Frisian = &Language{
		Name:        "frisian",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: EnglishWordBaseRef, // @TODO can be changed after adding frisian to word bases
	}
	German = &Language{
		Name:        "german",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: GermanWordBaseRef,
	}
	Dutch = &Language{
		Name:        "dutch",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding dutch to word bases
	}
	Luxembourgish = &Language{
		Name:        "luxembourgish",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding luxembourgish to word bases
	}
	Yiddish = &Language{
		Name:        "yiddish",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding yiddish to word bases
	}
	Gothic = &Language{
		Name:        "gothic",
		Subfamily:   EastGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding gothic to word bases
	}
	Vandalic = &Language{
		Name:        "vandalic",
		Subfamily:   EastGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding vandalic to word bases
	}
	Burgundian = &Language{
		Name:        "burgundian",
		Subfamily:   EastGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding burgundian to word bases
	}
	Alamannian = &Language{
		Name:        "alamannian",
		Subfamily:   ElbeGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding alamannian to word bases
	}
	Bavarian = &Language{
		Name:        "bavarian",
		Subfamily:   ElbeGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding bavarian to word bases
	}
	Langobardian = &Language{
		Name:        "lombardian",
		Subfamily:   ElbeGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding lombardian to word bases
	}
	Frankish = &Language{
		Name:        "frankish",
		Subfamily:   WeserRhineGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding frankish to word bases
	}
	Anglic = &Language{
		Name:        "anglic",
		Subfamily:   NorthSeaGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding anglic to word bases
	}
	Saxon = &Language{
		Name:        "saxon",
		Subfamily:   NorthSeaGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding saxon to word bases
	}
	// Greek
	Greek = &Language{
		Name:        "greek",
		Subfamily:   GreekSubfamily,
		IsLiving:    true,
		WordBaseRef: GreekWordBaseRef,
	}
	// Illyrian
	Illyrian = &Language{
		Name:        "illyrian",
		Subfamily:   IllyrianSubfamily,
		IsLiving:    false,
		WordBaseRef: IllyrianWordBaseRef,
	}
	// IndoIranian
	Pashai = &Language{
		Name:        "pashai",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Chitrali = &Language{
		Name:        "chitrali",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Shina = &Language{
		Name:        "shina",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Kohistani = &Language{
		Name:        "kohistani",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Kashmiri = &Language{
		Name:        "kashmiri",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Punjabi = &Language{
		Name:        "punjabi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Sindhi = &Language{
		Name:        "sindhi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Rajasthani = &Language{
		Name:        "rajasthani",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Gujarati = &Language{
		Name:        "gujarati",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Bhili = &Language{
		Name:        "bhili",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Khandeshi = &Language{
		Name:        "khandeshi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	HimachaliDogri = &Language{
		Name:        "himachali_dogri",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	GirhwaliKumaoni = &Language{
		Name:        "girhwali_kumaoni",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Nepali = &Language{
		Name:        "nepali",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Hindi = &Language{
		Name:        "hindi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Bihari = &Language{
		Name:        "bihari",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	BengaliAssamese = &Language{
		Name:        "bengali_assamese",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Odia = &Language{
		Name:        "odia",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Halbi = &Language{
		Name:        "halbi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	MarathiKonkani = &Language{
		Name:        "marathi_konkani",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	SinhalaMaldivian = &Language{
		Name:        "sinhala_maldivian",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	// Iranian
	Farsi = &Language{
		Name:        "farsi",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Avestan = &Language{
		Name:        "avestan",
		Subfamily:   IranianSubfamily,
		IsLiving:    false,
		WordBaseRef: IranianWordBaseRef,
	}
	Dari = &Language{
		Name:        "dari",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Tajik = &Language{
		Name:        "tajik",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Pashto = &Language{
		Name:        "pashto",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Kurdish = &Language{
		Name:        "kurdish",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Luri = &Language{
		Name:        "luri",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Balochi = &Language{
		Name:        "balochi",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Scythian = &Language{
		Name:        "scythian",
		Subfamily:   IranianSubfamily,
		IsLiving:    false,
		WordBaseRef: IranianWordBaseRef,
	}
	// Nuriastani
	KamkataVari = &Language{
		Name:        "kamkata_vari",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into KamkataVari when available
	}
	VasiVari = &Language{
		Name:        "vasi_vari",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into VasiVari when available
		IsLiving:    true,
	}
	Askunu = &Language{
		Name:        "askunu",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into Askunu when available
		IsLiving:    true,
	}
	Waigali = &Language{
		Name:        "waigali",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into Waigali when available
		IsLiving:    true,
	}
	Tregami = &Language{
		Name:        "tregami",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into Tregami when available
		IsLiving:    true,
	}
	Zemiaki = &Language{
		Name:        "zemiaki",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into Zemiaki when available
		IsLiving:    true,
	}
	// Italic
	Venetic = &Language{
		Name:        "venetic",
		Subfamily:   ItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into venetic when available
		IsLiving:    false,
	}
	Sicel = &Language{
		Name:        "sicel",
		Subfamily:   ItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into sicel when available
		IsLiving:    false,
	}
	Lusitanian = &Language{
		Name:        "lusitanian",
		Subfamily:   ItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into lusitanian when available
		IsLiving:    false,
	}
	Latin = &Language{
		Name:        "latin",
		Subfamily:   LatinoFaliscanSubfamily,
		WordBaseRef: LatinWordBaseRef, // @TODO change it into latin when available
		IsLiving:    true,
	}
	Faliscan = &Language{
		Name:        "faliscan",
		Subfamily:   LatinoFaliscanSubfamily,
		WordBaseRef: LatinWordBaseRef, // @TODO change it into faliscan when available
		IsLiving:    false,
	}
	Lanuvain = &Language{
		Name:        "lanuvain",
		Subfamily:   LatinoFaliscanSubfamily,
		WordBaseRef: LatinWordBaseRef, // @TODO change it into lanuvain when available
		IsLiving:    false,
	}
	Venetian = &Language{
		Name:        "venetian",
		Subfamily:   RomanceSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into venetian when available
		IsLiving:    true,
	}
	Sardinian = &Language{
		Name:        "sardinian",
		Subfamily:   RomanceSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into sardinian when available
		IsLiving:    true,
	}
	Portuguese = &Language{
		Name:        "portuguese",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: PortugueseWordBaseRef,
		IsLiving:    true,
	}
	Galician = &Language{
		Name:        "galician",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: PortugueseWordBaseRef,
		IsLiving:    true,
	}
	Asturleonese = &Language{
		Name:        "asturleonese",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Mirandese = &Language{
		Name:        "mirandese",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: PortugueseWordBaseRef,
		IsLiving:    true,
	}
	Spanish = &Language{
		Name:        "spanish",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Aragonese = &Language{
		Name:        "aragonese",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Ladino = &Language{
		Name:        "ladino",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Catalan = &Language{
		Name:        "catalan",
		Subfamily:   OccitanoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Occitan = &Language{
		Name:        "occitan",
		Subfamily:   OccitanoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Gascon = &Language{
		Name:        "gascon",
		Subfamily:   OccitanoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	French = &Language{
		Name:        "french",
		Subfamily:   GalloRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	FrancoProvencal = &Language{
		Name:        "franco_provencal",
		Subfamily:   GalloRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Romansh = &Language{
		Name:        "romansh",
		Subfamily:   RhaetoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Ladin = &Language{
		Name:        "ladin",
		Subfamily:   RhaetoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Friulian = &Language{
		Name:        "friulian",
		Subfamily:   RhaetoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Piedmontense = &Language{
		Name:        "piedmontense",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Ligurian = &Language{
		Name:        "ligurian",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Lombard = &Language{
		Name:        "lombard",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Emilian = &Language{
		Name:        "emilian",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Romagnol = &Language{
		Name:        "romagnol",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Italian = &Language{
		Name:        "italian",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Sicilian = &Language{
		Name:        "sicilian",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Neapolitan = &Language{
		Name:        "neapolitan",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Dalmatian = &Language{
		Name:        "dalmatian",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Istriot = &Language{
		Name:        "istriot",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Romanian = &Language{
		Name:        "romanian",
		Subfamily:   EasternRomanceSubfamily,
		WordBaseRef: LatinWordBaseRef,
		IsLiving:    true,
	}
	Aromanian = &Language{
		Name:        "aromanian",
		Subfamily:   EasternRomanceSubfamily,
		WordBaseRef: LatinWordBaseRef,
		IsLiving:    true,
	}
	MaglenoRomanian = &Language{
		Name:        "magleno_romanian",
		Subfamily:   EasternRomanceSubfamily,
		WordBaseRef: LatinWordBaseRef,
		IsLiving:    true,
	}
	IstroRomanian = &Language{
		Name:        "istro_romanian",
		Subfamily:   EasternRomanceSubfamily,
		WordBaseRef: LatinWordBaseRef,
		IsLiving:    true,
	}
	Oscan = &Language{
		Name:        "oscan",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Umbrian = &Language{
		Name:        "umbrian",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Volscian = &Language{
		Name:        "volscian",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Sabine = &Language{
		Name:        "sabine",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	SouthPicene = &Language{
		Name:        "south_picene",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Marsian = &Language{
		Name:        "marsian",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Paeligni = &Language{
		Name:        "paeligni",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Hernican = &Language{
		Name:        "hernican",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Marrucinian = &Language{
		Name:        "marrucinian",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	PreSamnite = &Language{
		Name:        "pre_samnite",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	// Tysenian
	Rhaetic = &Language{
		Name:        "rhaetic",
		Subfamily:   TysenianSubfamily,
		WordBaseRef: EtruscanWordBaseRef,
		IsLiving:    false,
	}
	Etruscan = &Language{
		Name:        "etruscan",
		Subfamily:   TysenianSubfamily,
		WordBaseRef: EtruscanWordBaseRef,
		IsLiving:    false,
	}
	// Tocharian
	Tocharian = &Language{
		Name:        "tocharian",
		Subfamily:   TocharianSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    false,
	}
)

var AllIndoEuropeanLanguages = []*Language{
	Albanian,
	Armenian,
	Ruthenian,
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
		Name:        "basque",
		Subfamily:   BasqueSubfamily,
		IsLiving:    true,
		WordBaseRef: BasqueWordBaseRef,
	}
)
var AllBasqueLanguages = []*Language{Basque}

// EastAsian
var (
	Japanese = &Language{
		Name:        "japanese",
		Subfamily:   JaponicSubfamily,
		IsLiving:    true,
		WordBaseRef: JapaneseWordBaseRef,
	}
	Korean = &Language{
		Name:        "korean",
		Subfamily:   KoreanicSubfamily,
		IsLiving:    true,
		WordBaseRef: KoreanWordBaseRef,
	}
)

var AllEastAsianLanguages = []*Language{Japanese, Korean}

// Kartvelian
var (
	Georgian = &Language{
		Name:        "Georgian",
		Subfamily:   KartoZanSubfamily,
		WordBaseRef: GeorgianWordBaseRef,
		IsLiving:    true,
	}
	Mingrelian = &Language{
		Name:        "Mingrelian",
		Subfamily:   KartoZanSubfamily,
		WordBaseRef: GeorgianWordBaseRef,
		IsLiving:    true,
	}
	Laz = &Language{
		Name:        "Laz",
		Subfamily:   KartoZanSubfamily,
		WordBaseRef: GeorgianWordBaseRef,
		IsLiving:    true,
	}
	Svan = &Language{
		Name:        "Svan",
		Subfamily:   SvanSubfamily,
		WordBaseRef: GeorgianWordBaseRef,
		IsLiving:    true,
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
		Name:        "chinese",
		Subfamily:   ChineseSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	Cantonese = &Language{
		Name:        "cantonese",
		Subfamily:   ChineseSubfamily,
		IsLiving:    true,
		WordBaseRef: CantoneseWordBaseRef,
	}
	Bai = &Language{
		Name:        "bai",
		Subfamily:   GreaterBaiSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	Caijia = &Language{
		Name:        "caijia",
		Subfamily:   GreaterBaiSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	Longjia = &Language{
		Name:        "longjia",
		Subfamily:   GreaterBaiSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	Luren = &Language{
		Name:        "luren",
		Subfamily:   GreaterBaiSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	// LoloBurmese
	Muangphe = &Language{
		Name:        "muangphe",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mango = &Language{
		Name:        "mango",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Manga = &Language{
		Name:        "manga",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Maang = &Language{
		Name:        "maang",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mondzi = &Language{
		Name:        "mondzi",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Maza = &Language{
		Name:        "maza",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mauphu = &Language{
		Name:        "mauphu",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Motang = &Language{
		Name:        "motang",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mongphu = &Language{
		Name:        "mongphu",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Burmese = &Language{
		Name:        "burmish",
		Subfamily:   BurmishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Jinuo = &Language{
		Name:        "jinuo",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Sangkong = &Language{
		Name:        "sangkong",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Bisu = &Language{
		Name:        "bisu",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Phunoi = &Language{
		Name:        "phunoi",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Pyen = &Language{
		Name:        "pyen",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Sila = &Language{
		Name:        "sila",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Phana = &Language{
		Name:        "phana",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Akeu = &Language{
		Name:        "akeu",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Hani = &Language{
		Name:        "hani",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Piyo = &Language{
		Name:        "piyo",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Enu = &Language{
		Name:        "enu",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mpi = &Language{
		Name:        "mpi",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Kaduo = &Language{
		Name:        "kaduo",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Laha = &Language{
		Name:        "laha",
		Subfamily:   LahoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Kucong = &Language{
		Name:        "kucong",
		Subfamily:   LahoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Namuyi = &Language{
		Name:        "namuyi",
		Subfamily:   NaxishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Shixing = &Language{
		Name:        "shixing",
		Subfamily:   NaxishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Naish = &Language{
		Name:        "naish",
		Subfamily:   NaxishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Nusu = &Language{
		Name:        "nusu",
		Subfamily:   NusoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Zauzou = &Language{
		Name:        "zauzou",
		Subfamily:   NusoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Katso = &Language{
		Name:        "katso",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Samu = &Language{
		Name:        "samu",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Sanie = &Language{
		Name:        "sanie",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Sadu = &Language{
		Name:        "sadu",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Meuma = &Language{
		Name:        "meuma",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Lisu = &Language{
		Name:        "lisu",
		Subfamily:   LisoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Lolopo = &Language{
		Name:        "lolopo",
		Subfamily:   LisoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Tangut = &Language{
		Name:        "tangut",
		Subfamily:   QiangicSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Tibetic = &Language{
		Name:        "tibetic",
		Subfamily:   BodishSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
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
		Name:        "vietnamese",
		Subfamily:   VietMuongSubfamily,
		IsLiving:    true,
		WordBaseRef: VietnameseWordBaseRef,
	}
	Kri = &Language{
		Name:        "kri",
		Subfamily:   VieticSubfamily,
		IsLiving:    true,
		WordBaseRef: VietnameseWordBaseRef,
	}
	Khmer = &Language{
		Name:        "khmer",
		Subfamily:   KhmericSubfamily,
		IsLiving:    true,
		WordBaseRef: VietnameseWordBaseRef,
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
		Name:        "mari",
		Subfamily:   FinnoPermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Merya = &Language{
		Name:        "merya",
		Subfamily:   FinnoPermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Komi = &Language{
		Name:        "komi",
		Subfamily:   PermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Permyak = &Language{
		Name:        "permyak",
		Subfamily:   PermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Udmurt = &Language{
		Name:        "udmurt",
		Subfamily:   PermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Meshchera = &Language{
		Name:        "meshchera",
		Subfamily:   PermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Estonian = &Language{
		Name:        "estonian",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Livonian = &Language{
		Name:        "livonian",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Votic = &Language{
		Name:        "votic",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Finnish = &Language{
		Name:        "finnish",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Ingrian = &Language{
		Name:        "ingrian",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Karelian = &Language{
		Name:        "karelian",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Ludic = &Language{
		Name:        "ludic",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Veps = &Language{
		Name:        "veps",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Sami = &Language{
		Name:        "sami",
		Subfamily:   SamiSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Erzya = &Language{
		Name:        "erzya",
		Subfamily:   MordvinSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Moksha = &Language{
		Name:        "moksha",
		Subfamily:   MordvinSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Hungarian = &Language{
		Name:        "hungarian",
		Subfamily:   UgricSubfamily,
		IsLiving:    true,
		WordBaseRef: HungarianWordBaseRef,
	}
	Khanty = &Language{
		Name:        "khanty",
		Subfamily:   UgricSubfamily,
		IsLiving:    true,
		WordBaseRef: HungarianWordBaseRef,
	}
	Mansi = &Language{
		Name:        "mansi",
		Subfamily:   UgricSubfamily,
		IsLiving:    true,
		WordBaseRef: HungarianWordBaseRef,
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
		Name:        "gagauz",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Turkish = &Language{
		Name:        "turkish",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Azerbaijani = &Language{
		Name:        "azerbaijani",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Turkmen = &Language{
		Name:        "turkmen",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Qashqai = &Language{
		Name:        "qashqai",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Chaharmahali = &Language{
		Name:        "chaharmahali",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Khorasani = &Language{
		Name:        "khorasani",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Salar = &Language{
		Name:        "salar",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Bashkir = &Language{
		Name:        "bashkir",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Tatar = &Language{
		Name:        "tatar",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	CrimeanTatar = &Language{
		Name:        "crimean_tatar",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	KarachayBalkar = &Language{
		Name:        "karachay_balkar",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Kumyk = &Language{
		Name:        "kumyk",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Karaim = &Language{
		Name:        "karaim",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Krymchak = &Language{
		Name:        "krymchak",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Urum = &Language{
		Name:        "urum",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Kazakh = &Language{
		Name:        "kazakh",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Karakalpak = &Language{
		Name:        "karakalpak",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Nogai = &Language{
		Name:        "nogai",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	SiberianTatar = &Language{
		Name:        "siberian_tatar",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Baraba = &Language{
		Name:        "baraba",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	SouthernAltai = &Language{
		Name:        "southern_altai",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Teleut = &Language{
		Name:        "teleut",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Kyrgyz = &Language{
		Name:        "kyrgyz",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Uzbek = &Language{
		Name:        "uzbek",
		Subfamily:   KarlukSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Uyghur = &Language{
		Name:        "uyghur",
		Subfamily:   KarlukSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Aynu = &Language{
		Name:        "aynu",
		Subfamily:   KarlukSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Ili = &Language{
		Name:        "ili",
		Subfamily:   KarlukSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Khalaj = &Language{
		Name:        "khalaj",
		Subfamily:   ShazTurkicSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Chuvash = &Language{
		Name:        "chuvash",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Bulgar = &Language{
		Name:        "bulgar",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Sabir = &Language{
		Name:        "sabir",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Khazar = &Language{
		Name:        "khazar",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Hunnic = &Language{
		Name:        "hunnic",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Tuoba = &Language{
		Name:        "tuoba",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Avar = &Language{
		Name:        "avar",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Cuman = &Language{
		Name:        "cuman",
		Subfamily:   KipchakSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
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
		Name:        "dagur",
		Subfamily:   DagurSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Moghol = &Language{
		Name:        "moghol",
		Subfamily:   MogholSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Khamnigan = &Language{
		Name:        "khamnigan",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Buryat = &Language{
		Name:        "buryat",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Mongolian = &Language{
		Name:        "mongolian",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Khalkha = &Language{
		Name:        "khalkha",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Chakhar = &Language{
		Name:        "chakhar",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Khorchin = &Language{
		Name:        "khorchin",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Ordos = &Language{
		Name:        "ordos",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Oirat = &Language{
		Name:        "oirat",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Monguor = &Language{
		Name:        "monguor",
		Subfamily:   SouthernMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Bonan = &Language{
		Name:        "bonan",
		Subfamily:   SouthernMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Santa = &Language{
		Name:        "santa",
		Subfamily:   SouthernMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Kangjia = &Language{
		Name:        "kangjia",
		Subfamily:   SouthernMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Khitan = &Language{
		Name:        "khitan",
		Subfamily:   KhitanSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
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
		Name:        "manchu",
		Subfamily:   JurchenicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Xibe = &Language{
		Name:        "xibe",
		Subfamily:   JurchenicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
)

var AllTungusicLanguages = []*Language{Manchu, Xibe}

// Dravidian
var (
	Tamil = &Language{
		Name:        "tamil",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Malayalam = &Language{
		Name:        "malayalam",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Irula = &Language{
		Name:        "irula",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kodava = &Language{
		Name:        "kodava",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kurumba = &Language{
		Name:        "kurumba",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Toda = &Language{
		Name:        "toda",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kota = &Language{
		Name:        "kota",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kannada = &Language{
		Name:        "kannada",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Badaga = &Language{
		Name:        "badaga",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Telugu = &Language{
		Name:        "telugu",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Chenchu = &Language{
		Name:        "chenchu",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Konda = &Language{
		Name:        "konda",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	MukhaDora = &Language{
		Name:        "mukha_dora",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Manda = &Language{
		Name:        "manda",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Pengo = &Language{
		Name:        "pengo",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kuvi = &Language{
		Name:        "kuvi",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kui = &Language{
		Name:        "kui",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Gondi = &Language{
		Name:        "gondi",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kolami = &Language{
		Name:        "kolami",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Naiki = &Language{
		Name:        "naiki",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Gadaba = &Language{
		Name:        "gadaba",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Ollari = &Language{
		Name:        "ollari",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kondekor = &Language{
		Name:        "kondekor",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Duruwa = &Language{
		Name:        "duruwa",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Oraon = &Language{
		Name:        "oraon",
		Subfamily:   NorthernDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kisan = &Language{
		Name:        "kisan",
		Subfamily:   NorthernDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	KumarbhagPaharia = &Language{
		Name:        "kumarbhag_paharia",
		Subfamily:   NorthernDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	SauriaPaharia = &Language{
		Name:        "sauria_paharia",
		Subfamily:   NorthernDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Brahui = &Language{
		Name:        "brahui",
		Subfamily:   NorthernDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
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
		Name:        "kru",
		Subfamily:   KruSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Kwa = &Language{
		Name:        "kwa",
		Subfamily:   VoltaCongoSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Gur = &Language{
		Name:        "gur",
		Subfamily:   VoltaCongoSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Soninke = &Language{
		Name:        "soninke",
		Subfamily:   MandeSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Manding = &Language{
		Name:        "manding",
		Subfamily:   MandeSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Senegambian = &Language{
		Name:        "senegambian",
		Subfamily:   AtlanticCongoSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Swahili = &Language{
		Name:        "swahili",
		Subfamily:   AtlanticCongoSubfamily,
		IsLiving:    true,
		WordBaseRef: SwahiliWordBaseRef,
	}
	Tebu = &Language{
		Name:        "tebu",
		Subfamily:   SaharanSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Hausa = &Language{
		Name:        "hausa",
		Subfamily:   WestChadicSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Nobiin = &Language{
		Name:        "nobiin",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Kenzi = &Language{
		Name:        "kenzi",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Midob = &Language{
		Name:        "midob",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Birgid = &Language{
		Name:        "birgid",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	HillNubian = &Language{
		Name:        "hill_nubian",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Daju = &Language{
		Name:        "daju",
		Subfamily:   DajuSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Somali = &Language{
		Name:        "somali",
		Subfamily:   SomalicSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	// Songhay
	Korandje = &Language{
		Name:        "korandje",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	KoyraChiini = &Language{
		Name:        "koyra_chiini",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tadaksahak = &Language{
		Name:        "tadaksahak",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tasawaq = &Language{
		Name:        "tasawaq",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tagdal = &Language{
		Name:        "tagdal",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	TondiSongwayKiini = &Language{
		Name:        "tondi_songway_kiini",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	HumburiSenni = &Language{
		Name:        "humburi_senni",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	KoyraboroSenni = &Language{
		Name:        "koyraboro_senni",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Zarma = &Language{
		Name:        "zarma",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	SonghoyboroCiine = &Language{
		Name:        "songhoyboro_ciine",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Dendi = &Language{
		Name:        "dendi",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
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
		Name:        "oodham",
		Subfamily:   TepimanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	PimaBajo = &Language{
		Name:        "pima_bajo",
		Subfamily:   TepimanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Tepehuan = &Language{
		Name:        "tepehuan",
		Subfamily:   TepimanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Southern = &Language{
		Name:        "southern",
		Subfamily:   TepimanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Tepecano = &Language{
		Name:        "tepecano",
		Subfamily:   TepimanSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Tarahumara = &Language{
		Name:        "tarahumara",
		Subfamily:   TarahumaranSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	UpriverGuarijio = &Language{
		Name:        "upriver_guarijio",
		Subfamily:   TarahumaranSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	DownriverGuarijio = &Language{
		Name:        "downriver_guarijio",
		Subfamily:   TarahumaranSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Tubar = &Language{
		Name:        "tubar",
		Subfamily:   TarahumaranSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Yaqui = &Language{
		Name:        "yaqui",
		Subfamily:   CahitaSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Mayo = &Language{
		Name:        "mayo",
		Subfamily:   CahitaSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Opata = &Language{
		Name:        "opata",
		Subfamily:   OpatanSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Eudeve = &Language{
		Name:        "eudeve",
		Subfamily:   OpatanSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Cora = &Language{
		Name:        "Cora",
		Subfamily:   CoracholSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Huichol = &Language{
		Name:        "Huichol",
		Subfamily:   CoracholSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Pochutec = &Language{
		Name:        "pochutec",
		Subfamily:   AztecanSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Pipil = &Language{
		Name:        "pipil",
		Subfamily:   AztecanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Nahuatl = &Language{
		Name:        "nahuatl",
		Subfamily:   AztecanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
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
	Name:        "inuit",
	Subfamily:   EskimoSubfamily,
	IsLiving:    true,
	WordBaseRef: InuitWordBaseRef,
}

var AllEskimoAleutLanguages = []*Language{Inuit}

// Austonesian
var (
	Hawaiian = &Language{
		Name:        "hawaiian",
		Subfamily:   MarquesicSubfamily,
		IsLiving:    true,
		WordBaseRef: HawaiianWordBaseRef,
	}
)

var AllAustonesianLanguages = []*Language{Hawaiian}

// Quechua
var (
	Ancash = &Language{
		Name:        "Ancash",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	AltoPativilcaAltoMaranonAltoHuallaga = &Language{
		Name:        "alto_pativilca_alto_maranon_alto_huallaga",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Yaru = &Language{
		Name:        "yaru",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Wanka = &Language{
		Name:        "wanka",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	YauyosChincha = &Language{
		Name:        "yauyos_chincha",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Pacaraos = &Language{
		Name:        "pacaraos",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Lambayeque = &Language{
		Name:        "lambayeque",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Cajamarca = &Language{
		Name:        "cajamarca",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Lincha = &Language{
		Name:        "lincha",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Laraos = &Language{
		Name:        "laraos",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Ayacucho = &Language{
		Name:        "ayacucho",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Cusco = &Language{
		Name:        "cusco",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Puno = &Language{
		Name:        "puno",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	NorthernBolivian = &Language{
		Name:        "northern_bolivian",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	SouthernBolivia = &Language{
		Name:        "southern_bolivia",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	SantiagoDelEstero = &Language{
		Name:        "santiago_del_estero",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
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
		Name:        "kabyle",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tamazight = &Language{
		Name:        "tamazight",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Shilha = &Language{
		Name:        "shilha",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	SenhajaDeSrair = &Language{
		Name:        "senhaja_de_srair",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Ghomara = &Language{
		Name:        "ghomara",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Riffian = &Language{
		Name:        "riffian",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	AytSeghrouchen = &Language{
		Name:        "ayt_seghrouchen",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	AytWarayn = &Language{
		Name:        "ayt_warayn",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Shenwa = &Language{
		Name:        "shenwa",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Shawiya = &Language{
		Name:        "shawiya",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Zenaga = &Language{
		Name:        "zenaga",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Siwi = &Language{
		Name:        "siwi",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Nafusi = &Language{
		Name:        "nafusi",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Sokna = &Language{
		Name:        "sokna",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Ghadames = &Language{
		Name:        "ghadames",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Awjila = &Language{
		Name:        "awjila",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tuareg = &Language{
		Name:        "tuareg",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Numidian = &Language{
		Name:        "numidian",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	// Cushitic
	Beja = &Language{
		Name:        "beja",
		Subfamily:   CushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Awngi = &Language{
		Name:        "awngi",
		Subfamily:   AgawSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Bilen = &Language{
		Name:        "bilen",
		Subfamily:   AgawSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Qimant = &Language{
		Name:        "qimant",
		Subfamily:   AgawSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Xamtanga = &Language{
		Name:        "xamtanga",
		Subfamily:   AgawSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Gawwada = &Language{
		Name:        "gawwada",
		Subfamily:   DullaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tsamai = &Language{
		Name:        "tsamai",
		Subfamily:   DullaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Dihina = &Language{
		Name:        "dihina",
		Subfamily:   DullaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Dobase = &Language{
		Name:        "dobase",
		Subfamily:   DullaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Burji = &Language{
		Name:        "burji",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Sidamo = &Language{
		Name:        "sidamo",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Gedeo = &Language{
		Name:        "gedeo",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	HadiyyaLibido = &Language{
		Name:        "hadiyya_libido",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	KambaataAlaba = &Language{
		Name:        "kambaata_alaba",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	SahoAfar = &Language{
		Name:        "saho_afar",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	OmoTana = &Language{
		Name:        "omo_tana",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Oromoid = &Language{
		Name:        "oromoid",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Dullay = &Language{
		Name:        "dullay",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Yaaku = &Language{
		Name:        "yaaku",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Gorowa = &Language{
		Name:        "gorowa",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Iraqw = &Language{
		Name:        "iraqw",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Alagwa = &Language{
		Name:        "alagwa",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Burunge = &Language{
		Name:        "burunge",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Aasax = &Language{
		Name:        "aasax",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Kwadza = &Language{
		Name:        "kwadza",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Egyptian = &Language{
		Name:        "egyptian",
		Subfamily:   EgyptianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	// semitic
	Akkadian = &Language{
		Name:        "akkadian",
		Subfamily:   EastSemiticSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Eblaite = &Language{
		Name:        "eblaite",
		Subfamily:   EastSemiticSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Kishite = &Language{
		Name:        "kishite",
		Subfamily:   EastSemiticSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Ashkenazi = &Language{
		Name:        "ashkenazi",
		Subfamily:   HebrewSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Sephardi = &Language{
		Name:        "sephardi",
		Subfamily:   HebrewSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Hebrew = &Language{
		Name:        "hebrew",
		Subfamily:   HebrewSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Samaritan = &Language{
		Name:        "samaritan",
		Subfamily:   HebrewSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Ammonite = &Language{
		Name:        "ammonite",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Edomite = &Language{
		Name:        "edomite",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Moabite = &Language{
		Name:        "moabite",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Phoenician = &Language{
		Name:        "phoenician",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Punic = &Language{
		Name:        "punic",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Samalian = &Language{
		Name:        "samalian",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Aramaic = &Language{
		Name:        "aramaic",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Nabataean = &Language{
		Name:        "nabataean",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Amorite = &Language{
		Name:        "amorite",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Himyaritic = &Language{
		Name:        "himyaritic",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Sutean = &Language{
		Name:        "sutean",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Syriac = &Language{
		Name:        "syriac",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Arabic = &Language{
		Name:        "arabic",
		Subfamily:   ArabicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Geez = &Language{
		Name:        "geez",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Tigrinya = &Language{
		Name:        "tigrinya",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Amharic = &Language{
		Name:        "amharic",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Argobba = &Language{
		Name:        "argobba",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Harari = &Language{
		Name:        "harari",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Gafat = &Language{
		Name:        "gafat",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Soddo = &Language{
		Name:        "soddo",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Mesmes = &Language{
		Name:        "mesmes",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Muher = &Language{
		Name:        "muher",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	WestGurage = &Language{
		Name:        "west_gurage",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Mesqan = &Language{
		Name:        "mesqan",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	SebatBet = &Language{
		Name:        "sebat_bet",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Baṭḥari = &Language{
		Name:        "bathari",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Harsusi = &Language{
		Name:        "harsusi",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Hobyot = &Language{
		Name:        "hobyot",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Mehri = &Language{
		Name:        "mehri",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Shehri = &Language{
		Name:        "shehri",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Soqotri = &Language{
		Name:        "soqotri",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Faifi = &Language{
		Name:        "faifi",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Hadramautic = &Language{
		Name:        "hadramautic",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Minaean = &Language{
		Name:        "minaean",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Qatabanian = &Language{
		Name:        "qatabanian",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Razihi = &Language{
		Name:        "razihi",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Sabaean = &Language{
		Name:        "sabaean",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	// Omotic
	SouthOmotic = &Language{
		Name:        "south_omotic",
		IsLiving:    true,
		Subfamily:   OmoticSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Mao = &Language{
		Name:        "mao",
		IsLiving:    true,
		Subfamily:   OmoticSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Dizoid = &Language{
		Name:        "dizoid",
		IsLiving:    true,
		Subfamily:   OmoticSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Gonga = &Language{
		Name:        "gonga",
		IsLiving:    true,
		Subfamily:   OmoticSubfamily,
		WordBaseRef: ArabicWordBaseRef,
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
	Numidian,
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
	Punic,
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
		Name:        "common_eldarin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding common_eldarin to word bases
	}
	Quenya = &Language{
		Name:        "quenya",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding quenya to word bases
	}
	Quendya = &Language{
		Name:        "quendya",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding quendya to word bases
	}
	ExilicQuendya = &Language{
		Name:        "exilic_quendya",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding exilic_quendya to word bases
	}
	CommonTelerin = &Language{
		Name:        "common_telerin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding common_telerin to word bases
	}
	Telerin = &Language{
		Name:        "telerin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding telerin to word bases
	}
	Sindarin = &Language{
		Name:        "sindarin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef,
	}
	Nandorin = &Language{
		Name:        "nandorin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding nandorin to word bases
	}
	Avarin = &Language{
		Name:        "avarin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding avarin to word bases
	}
	Aldmeris = &Language{
		Name:        "aldmeris",
		Subfamily:   AldmerisSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding aldmeris to word bases
	}
	Dunmeris = &Language{
		Name:        "dunmeris",
		Subfamily:   AldmerisSubfamily,
		IsLiving:    true,
		WordBaseRef: DunmerisWordBaseRef,
	}
	Bosmeris = &Language{
		Name:        "bosmeris",
		Subfamily:   AldmerisSubfamily,
		IsLiving:    true,
		WordBaseRef: DunmerisWordBaseRef, // @TODO can be changed after adding bosmeris to word bases
	}
	Dwemeris = &Language{
		Name:        "dwemeris",
		Subfamily:   DwermerisSubfamily,
		IsLiving:    true,
		WordBaseRef: DwemerisWordBaseRef,
	}
	Goblins = &Language{
		Name:        "goblins",
		Subfamily:   OrcishSubfamily,
		IsLiving:    true,
		WordBaseRef: GoblinsWordBaseRef,
	}
	Orcish = &Language{
		Name:        "orcish",
		Subfamily:   OrcishSubfamily,
		IsLiving:    true,
		WordBaseRef: OrcWordBaseRef,
	}
	Giantish = &Language{
		Name:        "giantish",
		Subfamily:   GiantishSubfamily,
		IsLiving:    true,
		WordBaseRef: GiantWordBaseRef,
	}
	Draconic = &Language{
		Name:        "draconic",
		Subfamily:   DraconicSubfamily,
		IsLiving:    true,
		WordBaseRef: DraconicWordBaseRef,
	}
	Arachnid = &Language{
		Name:        "arachnid",
		Subfamily:   ArachnidSubfamily,
		IsLiving:    true,
		WordBaseRef: ArachnidWordBaseRef,
	}
	Serpents = &Language{
		Name:        "serpents",
		Subfamily:   SerpentsSubfamily,
		IsLiving:    true,
		WordBaseRef: SerpentsWordBaseRef,
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

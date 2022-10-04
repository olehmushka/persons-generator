package language

import (
	"persons_generator/core/tools"

	"github.com/google/uuid"
)

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
		ID:          uuid.New(),
		Name:        "albanian",
		Subfamily:   AlbanianSubfamily,
		IsLiving:    true,
		WordBaseRef: AlbanianWordBaseRef,
	}
	Armenian = &Language{
		ID:          uuid.New(),
		Name:        "armenian",
		Subfamily:   ArmenianSubfamily,
		IsLiving:    true,
		WordBaseRef: ArmenianWordBaseRef,
	}
	// Slavic
	Ruthenian = &Language{
		ID:          uuid.New(),
		Name:        "ruthenian",
		Subfamily:   RuthenianSubfamily,
		IsLiving:    false,
		WordBaseRef: RuthenianWordBaseRef,
	}
	Belarusian = &Language{
		ID:          uuid.New(),
		Name:        "belarusian",
		Subfamily:   RuthenianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding belarusian to word bases
	}
	Ukrainian = &Language{
		ID:          uuid.New(),
		Name:        "ukrainian",
		Subfamily:   RuthenianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding ukrainian to word bases
	}
	Russian = &Language{
		ID:          uuid.New(),
		Name:        "russian",
		Subfamily:   MoscovianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding russian to word bases
	}
	Slovenian = &Language{
		ID:          uuid.New(),
		Name:        "slovenian",
		Subfamily:   WesternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding slovenian to word bases
	}
	Croatian = &Language{
		ID:          uuid.New(),
		Name:        "croation",
		Subfamily:   WesternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding croation to word bases
	}
	Serbian = &Language{
		ID:          uuid.New(),
		Name:        "serbian",
		Subfamily:   WesternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding serbian to word bases
	}
	Bosnian = &Language{
		ID:          uuid.New(),
		Name:        "bosnian",
		Subfamily:   WesternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding bosnian to word bases
	}
	Bulgarian = &Language{
		ID:          uuid.New(),
		Name:        "bulgarian",
		Subfamily:   EasternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding bulgarian to word bases
	}
	Macedonian = &Language{
		ID:          uuid.New(),
		Name:        "macedonian",
		Subfamily:   EasternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding macedonian to word bases
	}
	Polish = &Language{
		ID:          uuid.New(),
		Name:        "polish",
		Subfamily:   LechiticSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding polish to word bases
	}
	Silesian = &Language{
		ID:          uuid.New(),
		Name:        "silesian",
		Subfamily:   LechiticSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding silesian to word bases
	}
	LowerSorbian = &Language{
		ID:          uuid.New(),
		Name:        "lower_sorbian",
		Subfamily:   SorbianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding lower_sorbian to word bases
	}
	UpperSorbian = &Language{
		ID:          uuid.New(),
		Name:        "upper_sorbian",
		Subfamily:   SorbianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding upper_sorbian to word bases
	}
	Czech = &Language{
		ID:          uuid.New(),
		Name:        "czech",
		Subfamily:   CzechSlovakSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding czech to word bases
	}
	Slovak = &Language{
		ID:          uuid.New(),
		Name:        "slovak",
		Subfamily:   CzechSlovakSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding slovak to word bases
	}
	// Baltic
	Latvian = &Language{
		ID:          uuid.New(),
		Name:        "latvian",
		Subfamily:   EasternBalticSubfamily,
		IsLiving:    true,
		WordBaseRef: LithuanianWordBaseRef, // @TODO can be changed after adding latvian to word bases
	}
	Latgalian = &Language{
		ID:          uuid.New(),
		Name:        "latgalian",
		Subfamily:   EasternBalticSubfamily,
		IsLiving:    true,
		WordBaseRef: LithuanianWordBaseRef, // @TODO can be changed after adding latgalian to word bases
	}
	Lithuanian = &Language{
		ID:          uuid.New(),
		Name:        "lithuanian",
		Subfamily:   EasternBalticSubfamily,
		IsLiving:    true,
		WordBaseRef: LithuanianWordBaseRef,
	}
	Semigallian = &Language{
		ID:          uuid.New(),
		Name:        "semigallian",
		Subfamily:   EasternBalticSubfamily,
		IsLiving:    false,
		WordBaseRef: LithuanianWordBaseRef, // @TODO can be changed after adding semigallian to word bases
	}
	// Celtic
	Celtiberian = &Language{
		ID:          uuid.New(),
		Name:        "celtiberian",
		Subfamily:   CelticSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Gallaecian = &Language{
		ID:          uuid.New(),
		Name:        "gallaecian",
		Subfamily:   CelticSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Noric = &Language{
		ID:          uuid.New(),
		Name:        "noric",
		Subfamily:   CelticSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Pictish = &Language{
		ID:          uuid.New(),
		Name:        "pictish",
		Subfamily:   BrittonicSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Cumbric = &Language{
		ID:          uuid.New(),
		Name:        "cumbric",
		Subfamily:   BrittonicSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Welsh = &Language{
		ID:          uuid.New(),
		Name:        "welsh",
		Subfamily:   WesternBrittonicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	Cornish = &Language{
		ID:          uuid.New(),
		Name:        "cornish",
		Subfamily:   SouthWesternBrittonicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	Breton = &Language{
		ID:          uuid.New(),
		Name:        "breton",
		Subfamily:   SouthWesternBrittonicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	Irish = &Language{
		ID:          uuid.New(),
		Name:        "irish",
		Subfamily:   GoidelicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	Manx = &Language{
		ID:          uuid.New(),
		Name:        "manx",
		Subfamily:   GoidelicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	ScottishGaelic = &Language{
		ID:          uuid.New(),
		Name:        "scottish_gaelic",
		Subfamily:   GoidelicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	// Dacian
	Dacian = &Language{
		ID:        uuid.New(),
		Name:      "dacian",
		Subfamily: DacianSubfamily,
		IsLiving:  false,
	}
	// Germanic
	Danish = &Language{
		ID:          uuid.New(),
		Name:        "danish",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding danish to word bases
	}
	Faroese = &Language{
		ID:          uuid.New(),
		Name:        "faroese",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding faroese to word bases
	}
	Icelandic = &Language{
		ID:          uuid.New(),
		Name:        "icelandic",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding icelandic to word bases
	}
	Norwegian = &Language{
		ID:          uuid.New(),
		Name:        "norwegian",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding norwegian to word bases
	}
	Swedish = &Language{
		ID:          uuid.New(),
		Name:        "swedish",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef,
	}
	Dalecarlian = &Language{
		ID:          uuid.New(),
		Name:        "dalecarlian",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding dalecarlian to word bases
	}
	Gutnish = &Language{
		ID:          uuid.New(),
		Name:        "gutnish",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding gutnish to word bases
	}
	English = &Language{
		ID:          uuid.New(),
		Name:        "english",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: EnglishWordBaseRef,
	}
	Scots = &Language{
		ID:          uuid.New(),
		Name:        "scots",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: EnglishWordBaseRef, // @TODO can be changed after adding scots to word bases
	}
	Yola = &Language{
		ID:          uuid.New(),
		Name:        "yola",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: EnglishWordBaseRef, // @TODO can be changed after adding yola to word bases
	}
	Frisian = &Language{
		ID:          uuid.New(),
		Name:        "frisian",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: EnglishWordBaseRef, // @TODO can be changed after adding frisian to word bases
	}
	German = &Language{
		ID:          uuid.New(),
		Name:        "german",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: GermanWordBaseRef,
	}
	Dutch = &Language{
		ID:          uuid.New(),
		Name:        "dutch",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding dutch to word bases
	}
	Luxembourgish = &Language{
		ID:          uuid.New(),
		Name:        "luxembourgish",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding luxembourgish to word bases
	}
	Yiddish = &Language{
		ID:          uuid.New(),
		Name:        "yiddish",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding yiddish to word bases
	}
	Gothic = &Language{
		ID:          uuid.New(),
		Name:        "gothic",
		Subfamily:   EastGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding gothic to word bases
	}
	Vandalic = &Language{
		ID:          uuid.New(),
		Name:        "vandalic",
		Subfamily:   EastGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding vandalic to word bases
	}
	Burgundian = &Language{
		ID:          uuid.New(),
		Name:        "burgundian",
		Subfamily:   EastGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding burgundian to word bases
	}
	Alamannian = &Language{
		ID:          uuid.New(),
		Name:        "alamannian",
		Subfamily:   ElbeGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding alamannian to word bases
	}
	Bavarian = &Language{
		ID:          uuid.New(),
		Name:        "bavarian",
		Subfamily:   ElbeGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding bavarian to word bases
	}
	Langobardian = &Language{
		ID:          uuid.New(),
		Name:        "lombardian",
		Subfamily:   ElbeGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding lombardian to word bases
	}
	Frankish = &Language{
		ID:          uuid.New(),
		Name:        "frankish",
		Subfamily:   WeserRhineGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding frankish to word bases
	}
	Anglic = &Language{
		ID:          uuid.New(),
		Name:        "anglic",
		Subfamily:   NorthSeaGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding anglic to word bases
	}
	Saxon = &Language{
		ID:          uuid.New(),
		Name:        "saxon",
		Subfamily:   NorthSeaGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding saxon to word bases
	}
	// Greek
	Greek = &Language{
		ID:          uuid.New(),
		Name:        "greek",
		Subfamily:   GreekSubfamily,
		IsLiving:    true,
		WordBaseRef: GreekWordBaseRef,
	}
	// Illyrian
	Illyrian = &Language{
		ID:          uuid.New(),
		Name:        "illyrian",
		Subfamily:   IllyrianSubfamily,
		IsLiving:    false,
		WordBaseRef: IllyrianWordBaseRef,
	}
	// IndoIranian
	Pashai = &Language{
		ID:          uuid.New(),
		Name:        "pashai",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Chitrali = &Language{
		ID:          uuid.New(),
		Name:        "chitrali",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Shina = &Language{
		ID:          uuid.New(),
		Name:        "shina",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Kohistani = &Language{
		ID:          uuid.New(),
		Name:        "kohistani",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Kashmiri = &Language{
		ID:          uuid.New(),
		Name:        "kashmiri",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Punjabi = &Language{
		ID:          uuid.New(),
		Name:        "punjabi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Sindhi = &Language{
		ID:          uuid.New(),
		Name:        "sindhi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Rajasthani = &Language{
		ID:          uuid.New(),
		Name:        "rajasthani",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Gujarati = &Language{
		ID:          uuid.New(),
		Name:        "gujarati",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Bhili = &Language{
		ID:          uuid.New(),
		Name:        "bhili",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Khandeshi = &Language{
		ID:          uuid.New(),
		Name:        "khandeshi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	HimachaliDogri = &Language{
		ID:          uuid.New(),
		Name:        "himachali_dogri",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	GirhwaliKumaoni = &Language{
		ID:          uuid.New(),
		Name:        "girhwali_kumaoni",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Nepali = &Language{
		ID:          uuid.New(),
		Name:        "nepali",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Hindi = &Language{
		ID:          uuid.New(),
		Name:        "hindi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Bihari = &Language{
		ID:          uuid.New(),
		Name:        "bihari",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	BengaliAssamese = &Language{
		ID:          uuid.New(),
		Name:        "bengali_assamese",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Odia = &Language{
		ID:          uuid.New(),
		Name:        "odia",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Halbi = &Language{
		ID:          uuid.New(),
		Name:        "halbi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	MarathiKonkani = &Language{
		ID:          uuid.New(),
		Name:        "marathi_konkani",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	SinhalaMaldivian = &Language{
		ID:          uuid.New(),
		Name:        "sinhala_maldivian",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	// Iranian
	Farsi = &Language{
		ID:          uuid.New(),
		Name:        "farsi",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Avestan = &Language{
		ID:          uuid.New(),
		Name:        "avestan",
		Subfamily:   IranianSubfamily,
		IsLiving:    false,
		WordBaseRef: IranianWordBaseRef,
	}
	Dari = &Language{
		ID:          uuid.New(),
		Name:        "dari",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Tajik = &Language{
		ID:          uuid.New(),
		Name:        "tajik",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Pashto = &Language{
		ID:          uuid.New(),
		Name:        "pashto",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Kurdish = &Language{
		ID:          uuid.New(),
		Name:        "kurdish",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Luri = &Language{
		ID:          uuid.New(),
		Name:        "luri",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Balochi = &Language{
		ID:          uuid.New(),
		Name:        "balochi",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Scythian = &Language{
		ID:          uuid.New(),
		Name:        "scythian",
		Subfamily:   IranianSubfamily,
		IsLiving:    false,
		WordBaseRef: IranianWordBaseRef,
	}
	// Nuriastani
	KamkataVari = &Language{
		ID:          uuid.New(),
		Name:        "kamkata_vari",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into KamkataVari when available
	}
	VasiVari = &Language{
		ID:          uuid.New(),
		Name:        "vasi_vari",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into VasiVari when available
		IsLiving:    true,
	}
	Askunu = &Language{
		ID:          uuid.New(),
		Name:        "askunu",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into Askunu when available
		IsLiving:    true,
	}
	Waigali = &Language{
		ID:          uuid.New(),
		Name:        "waigali",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into Waigali when available
		IsLiving:    true,
	}
	Tregami = &Language{
		ID:          uuid.New(),
		Name:        "tregami",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into Tregami when available
		IsLiving:    true,
	}
	Zemiaki = &Language{
		ID:          uuid.New(),
		Name:        "zemiaki",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into Zemiaki when available
		IsLiving:    true,
	}
	// Italic
	Venetic = &Language{
		ID:          uuid.New(),
		Name:        "venetic",
		Subfamily:   ItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into venetic when available
		IsLiving:    false,
	}
	Sicel = &Language{
		ID:          uuid.New(),
		Name:        "sicel",
		Subfamily:   ItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into sicel when available
		IsLiving:    false,
	}
	Lusitanian = &Language{
		ID:          uuid.New(),
		Name:        "lusitanian",
		Subfamily:   ItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into lusitanian when available
		IsLiving:    false,
	}
	Latin = &Language{
		ID:          uuid.New(),
		Name:        "latin",
		Subfamily:   LatinoFaliscanSubfamily,
		WordBaseRef: LatinWordBaseRef, // @TODO change it into latin when available
		IsLiving:    true,
	}
	Faliscan = &Language{
		ID:          uuid.New(),
		Name:        "faliscan",
		Subfamily:   LatinoFaliscanSubfamily,
		WordBaseRef: LatinWordBaseRef, // @TODO change it into faliscan when available
		IsLiving:    false,
	}
	Lanuvain = &Language{
		ID:          uuid.New(),
		Name:        "lanuvain",
		Subfamily:   LatinoFaliscanSubfamily,
		WordBaseRef: LatinWordBaseRef, // @TODO change it into lanuvain when available
		IsLiving:    false,
	}
	Venetian = &Language{
		ID:          uuid.New(),
		Name:        "venetian",
		Subfamily:   RomanceSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into venetian when available
		IsLiving:    true,
	}
	Sardinian = &Language{
		ID:          uuid.New(),
		Name:        "sardinian",
		Subfamily:   RomanceSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into sardinian when available
		IsLiving:    true,
	}
	Portuguese = &Language{
		ID:          uuid.New(),
		Name:        "portuguese",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: PortugueseWordBaseRef,
		IsLiving:    true,
	}
	Galician = &Language{
		ID:          uuid.New(),
		Name:        "galician",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: PortugueseWordBaseRef,
		IsLiving:    true,
	}
	Asturleonese = &Language{
		ID:          uuid.New(),
		Name:        "asturleonese",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Mirandese = &Language{
		ID:          uuid.New(),
		Name:        "mirandese",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: PortugueseWordBaseRef,
		IsLiving:    true,
	}
	Spanish = &Language{
		ID:          uuid.New(),
		Name:        "spanish",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Aragonese = &Language{
		ID:          uuid.New(),
		Name:        "aragonese",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Ladino = &Language{
		ID:          uuid.New(),
		Name:        "ladino",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Catalan = &Language{
		ID:          uuid.New(),
		Name:        "catalan",
		Subfamily:   OccitanoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Occitan = &Language{
		ID:          uuid.New(),
		Name:        "occitan",
		Subfamily:   OccitanoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Gascon = &Language{
		ID:          uuid.New(),
		Name:        "gascon",
		Subfamily:   OccitanoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	French = &Language{
		ID:          uuid.New(),
		Name:        "french",
		Subfamily:   GalloRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	FrancoProvencal = &Language{
		ID:          uuid.New(),
		Name:        "franco_provencal",
		Subfamily:   GalloRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Romansh = &Language{
		ID:          uuid.New(),
		Name:        "romansh",
		Subfamily:   RhaetoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Ladin = &Language{
		ID:          uuid.New(),
		Name:        "ladin",
		Subfamily:   RhaetoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Friulian = &Language{
		ID:          uuid.New(),
		Name:        "friulian",
		Subfamily:   RhaetoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Piedmontense = &Language{
		ID:          uuid.New(),
		Name:        "piedmontense",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Ligurian = &Language{
		ID:          uuid.New(),
		Name:        "ligurian",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Lombard = &Language{
		ID:          uuid.New(),
		Name:        "lombard",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Emilian = &Language{
		ID:          uuid.New(),
		Name:        "emilian",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Romagnol = &Language{
		ID:          uuid.New(),
		Name:        "romagnol",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Italian = &Language{
		ID:          uuid.New(),
		Name:        "italian",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Sicilian = &Language{
		ID:          uuid.New(),
		Name:        "sicilian",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Neapolitan = &Language{
		ID:          uuid.New(),
		Name:        "neapolitan",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Dalmatian = &Language{
		ID:          uuid.New(),
		Name:        "dalmatian",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Istriot = &Language{
		ID:          uuid.New(),
		Name:        "istriot",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Romanian = &Language{
		ID:          uuid.New(),
		Name:        "romanian",
		Subfamily:   EasternRomanceSubfamily,
		WordBaseRef: LatinWordBaseRef,
		IsLiving:    true,
	}
	Aromanian = &Language{
		ID:          uuid.New(),
		Name:        "aromanian",
		Subfamily:   EasternRomanceSubfamily,
		WordBaseRef: LatinWordBaseRef,
		IsLiving:    true,
	}
	MaglenoRomanian = &Language{
		ID:          uuid.New(),
		Name:        "magleno_romanian",
		Subfamily:   EasternRomanceSubfamily,
		WordBaseRef: LatinWordBaseRef,
		IsLiving:    true,
	}
	IstroRomanian = &Language{
		ID:          uuid.New(),
		Name:        "istro_romanian",
		Subfamily:   EasternRomanceSubfamily,
		WordBaseRef: LatinWordBaseRef,
		IsLiving:    true,
	}
	Oscan = &Language{
		ID:          uuid.New(),
		Name:        "oscan",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Umbrian = &Language{
		ID:          uuid.New(),
		Name:        "umbrian",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Volscian = &Language{
		ID:          uuid.New(),
		Name:        "volscian",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Sabine = &Language{
		ID:          uuid.New(),
		Name:        "sabine",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	SouthPicene = &Language{
		ID:          uuid.New(),
		Name:        "south_picene",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Marsian = &Language{
		ID:          uuid.New(),
		Name:        "marsian",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Paeligni = &Language{
		ID:          uuid.New(),
		Name:        "paeligni",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Hernican = &Language{
		ID:          uuid.New(),
		Name:        "hernican",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Marrucinian = &Language{
		ID:          uuid.New(),
		Name:        "marrucinian",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	PreSamnite = &Language{
		ID:          uuid.New(),
		Name:        "pre_samnite",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	// Tysenian
	Rhaetic = &Language{
		ID:          uuid.New(),
		Name:        "rhaetic",
		Subfamily:   TysenianSubfamily,
		WordBaseRef: EtruscanWordBaseRef,
		IsLiving:    false,
	}
	Etruscan = &Language{
		ID:          uuid.New(),
		Name:        "etruscan",
		Subfamily:   TysenianSubfamily,
		WordBaseRef: EtruscanWordBaseRef,
		IsLiving:    false,
	}
	// Tocharian
	Tocharian = &Language{
		ID:          uuid.New(),
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
		ID:          uuid.New(),
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
		ID:          uuid.New(),
		Name:        "japanese",
		Subfamily:   JaponicSubfamily,
		IsLiving:    true,
		WordBaseRef: JapaneseWordBaseRef,
	}
	Korean = &Language{
		ID:          uuid.New(),
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
		ID:          uuid.New(),
		Name:        "Georgian",
		Subfamily:   KartoZanSubfamily,
		WordBaseRef: GeorgianWordBaseRef,
		IsLiving:    true,
	}
	Mingrelian = &Language{
		ID:          uuid.New(),
		Name:        "Mingrelian",
		Subfamily:   KartoZanSubfamily,
		WordBaseRef: GeorgianWordBaseRef,
		IsLiving:    true,
	}
	Laz = &Language{
		ID:          uuid.New(),
		Name:        "Laz",
		Subfamily:   KartoZanSubfamily,
		WordBaseRef: GeorgianWordBaseRef,
		IsLiving:    true,
	}
	Svan = &Language{
		ID:          uuid.New(),
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
		ID:          uuid.New(),
		Name:        "chinese",
		Subfamily:   ChineseSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	Cantonese = &Language{
		ID:          uuid.New(),
		Name:        "cantonese",
		Subfamily:   ChineseSubfamily,
		IsLiving:    true,
		WordBaseRef: CantoneseWordBaseRef,
	}
	Bai = &Language{
		ID:          uuid.New(),
		Name:        "bai",
		Subfamily:   GreaterBaiSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	Caijia = &Language{
		ID:          uuid.New(),
		Name:        "caijia",
		Subfamily:   GreaterBaiSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	Longjia = &Language{
		ID:          uuid.New(),
		Name:        "longjia",
		Subfamily:   GreaterBaiSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	Luren = &Language{
		ID:          uuid.New(),
		Name:        "luren",
		Subfamily:   GreaterBaiSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	// LoloBurmese
	Muangphe = &Language{
		ID:          uuid.New(),
		Name:        "muangphe",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mango = &Language{
		ID:          uuid.New(),
		Name:        "mango",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Manga = &Language{
		ID:          uuid.New(),
		Name:        "manga",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Maang = &Language{
		ID:          uuid.New(),
		Name:        "maang",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mondzi = &Language{
		ID:          uuid.New(),
		Name:        "mondzi",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Maza = &Language{
		ID:          uuid.New(),
		Name:        "maza",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mauphu = &Language{
		ID:          uuid.New(),
		Name:        "mauphu",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Motang = &Language{
		ID:          uuid.New(),
		Name:        "motang",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mongphu = &Language{
		ID:          uuid.New(),
		Name:        "mongphu",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Burmese = &Language{
		ID:          uuid.New(),
		Name:        "burmish",
		Subfamily:   BurmishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Jinuo = &Language{
		ID:          uuid.New(),
		Name:        "jinuo",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Sangkong = &Language{
		ID:          uuid.New(),
		Name:        "sangkong",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Bisu = &Language{
		ID:          uuid.New(),
		Name:        "bisu",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Phunoi = &Language{
		ID:          uuid.New(),
		Name:        "phunoi",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Pyen = &Language{
		ID:          uuid.New(),
		Name:        "pyen",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Sila = &Language{
		ID:          uuid.New(),
		Name:        "sila",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Phana = &Language{
		ID:          uuid.New(),
		Name:        "phana",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Akeu = &Language{
		ID:          uuid.New(),
		Name:        "akeu",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Hani = &Language{
		ID:          uuid.New(),
		Name:        "hani",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Piyo = &Language{
		ID:          uuid.New(),
		Name:        "piyo",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Enu = &Language{
		ID:          uuid.New(),
		Name:        "enu",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mpi = &Language{
		ID:          uuid.New(),
		Name:        "mpi",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Kaduo = &Language{
		ID:          uuid.New(),
		Name:        "kaduo",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Laha = &Language{
		ID:          uuid.New(),
		Name:        "laha",
		Subfamily:   LahoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Kucong = &Language{
		ID:          uuid.New(),
		Name:        "kucong",
		Subfamily:   LahoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Namuyi = &Language{
		ID:          uuid.New(),
		Name:        "namuyi",
		Subfamily:   NaxishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Shixing = &Language{
		ID:          uuid.New(),
		Name:        "shixing",
		Subfamily:   NaxishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Naish = &Language{
		ID:          uuid.New(),
		Name:        "naish",
		Subfamily:   NaxishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Nusu = &Language{
		ID:          uuid.New(),
		Name:        "nusu",
		Subfamily:   NusoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Zauzou = &Language{
		ID:          uuid.New(),
		Name:        "zauzou",
		Subfamily:   NusoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Katso = &Language{
		ID:          uuid.New(),
		Name:        "katso",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Samu = &Language{
		ID:          uuid.New(),
		Name:        "samu",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Sanie = &Language{
		ID:          uuid.New(),
		Name:        "sanie",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Sadu = &Language{
		ID:          uuid.New(),
		Name:        "sadu",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Meuma = &Language{
		ID:          uuid.New(),
		Name:        "meuma",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Lisu = &Language{
		ID:          uuid.New(),
		Name:        "lisu",
		Subfamily:   LisoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Lolopo = &Language{
		ID:          uuid.New(),
		Name:        "lolopo",
		Subfamily:   LisoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Tangut = &Language{
		ID:          uuid.New(),
		Name:        "tangut",
		Subfamily:   QiangicSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Tibetic = &Language{
		ID:          uuid.New(),
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
		ID:          uuid.New(),
		Name:        "vietnamese",
		Subfamily:   VietMuongSubfamily,
		IsLiving:    true,
		WordBaseRef: VietnameseWordBaseRef,
	}
	Kri = &Language{
		ID:          uuid.New(),
		Name:        "kri",
		Subfamily:   VieticSubfamily,
		IsLiving:    true,
		WordBaseRef: VietnameseWordBaseRef,
	}
	Khmer = &Language{
		ID:          uuid.New(),
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
		ID:          uuid.New(),
		Name:        "mari",
		Subfamily:   FinnoPermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Merya = &Language{
		ID:          uuid.New(),
		Name:        "merya",
		Subfamily:   FinnoPermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Komi = &Language{
		ID:          uuid.New(),
		Name:        "komi",
		Subfamily:   PermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Permyak = &Language{
		ID:          uuid.New(),
		Name:        "permyak",
		Subfamily:   PermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Udmurt = &Language{
		ID:          uuid.New(),
		Name:        "udmurt",
		Subfamily:   PermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Meshchera = &Language{
		ID:          uuid.New(),
		Name:        "meshchera",
		Subfamily:   PermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Estonian = &Language{
		ID:          uuid.New(),
		Name:        "estonian",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Livonian = &Language{
		ID:          uuid.New(),
		Name:        "livonian",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Votic = &Language{
		ID:          uuid.New(),
		Name:        "votic",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Finnish = &Language{
		ID:          uuid.New(),
		Name:        "finnish",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Ingrian = &Language{
		ID:          uuid.New(),
		Name:        "ingrian",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Karelian = &Language{
		ID:          uuid.New(),
		Name:        "karelian",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Ludic = &Language{
		ID:          uuid.New(),
		Name:        "ludic",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Veps = &Language{
		ID:          uuid.New(),
		Name:        "veps",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Sami = &Language{
		ID:          uuid.New(),
		Name:        "sami",
		Subfamily:   SamiSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Erzya = &Language{
		ID:          uuid.New(),
		Name:        "erzya",
		Subfamily:   MordvinSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Moksha = &Language{
		ID:          uuid.New(),
		Name:        "moksha",
		Subfamily:   MordvinSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Hungarian = &Language{
		ID:          uuid.New(),
		Name:        "hungarian",
		Subfamily:   UgricSubfamily,
		IsLiving:    true,
		WordBaseRef: HungarianWordBaseRef,
	}
	Khanty = &Language{
		ID:          uuid.New(),
		Name:        "khanty",
		Subfamily:   UgricSubfamily,
		IsLiving:    true,
		WordBaseRef: HungarianWordBaseRef,
	}
	Mansi = &Language{
		ID:          uuid.New(),
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
		ID:          uuid.New(),
		Name:        "gagauz",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Turkish = &Language{
		ID:          uuid.New(),
		Name:        "turkish",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Azerbaijani = &Language{
		ID:          uuid.New(),
		Name:        "azerbaijani",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Turkmen = &Language{
		ID:          uuid.New(),
		Name:        "turkmen",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Qashqai = &Language{
		ID:          uuid.New(),
		Name:        "qashqai",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Chaharmahali = &Language{
		ID:          uuid.New(),
		Name:        "chaharmahali",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Khorasani = &Language{
		ID:          uuid.New(),
		Name:        "khorasani",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Salar = &Language{
		ID:          uuid.New(),
		Name:        "salar",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Bashkir = &Language{
		ID:          uuid.New(),
		Name:        "bashkir",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Tatar = &Language{
		ID:          uuid.New(),
		Name:        "tatar",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	CrimeanTatar = &Language{
		ID:          uuid.New(),
		Name:        "crimean_tatar",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	KarachayBalkar = &Language{
		ID:          uuid.New(),
		Name:        "karachay_balkar",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Kumyk = &Language{
		ID:          uuid.New(),
		Name:        "kumyk",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Karaim = &Language{
		ID:          uuid.New(),
		Name:        "karaim",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Krymchak = &Language{
		ID:          uuid.New(),
		Name:        "krymchak",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Urum = &Language{
		ID:          uuid.New(),
		Name:        "urum",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Kazakh = &Language{
		ID:          uuid.New(),
		Name:        "kazakh",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Karakalpak = &Language{
		ID:          uuid.New(),
		Name:        "karakalpak",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Nogai = &Language{
		ID:          uuid.New(),
		Name:        "nogai",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	SiberianTatar = &Language{
		ID:          uuid.New(),
		Name:        "siberian_tatar",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Baraba = &Language{
		ID:          uuid.New(),
		Name:        "baraba",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	SouthernAltai = &Language{
		ID:          uuid.New(),
		Name:        "southern_altai",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Teleut = &Language{
		ID:          uuid.New(),
		Name:        "teleut",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Kyrgyz = &Language{
		ID:          uuid.New(),
		Name:        "kyrgyz",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Uzbek = &Language{
		ID:          uuid.New(),
		Name:        "uzbek",
		Subfamily:   KarlukSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Uyghur = &Language{
		ID:          uuid.New(),
		Name:        "uyghur",
		Subfamily:   KarlukSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Aynu = &Language{
		ID:          uuid.New(),
		Name:        "aynu",
		Subfamily:   KarlukSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Ili = &Language{
		ID:          uuid.New(),
		Name:        "ili",
		Subfamily:   KarlukSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Khalaj = &Language{
		ID:          uuid.New(),
		Name:        "khalaj",
		Subfamily:   ShazTurkicSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Chuvash = &Language{
		ID:          uuid.New(),
		Name:        "chuvash",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Bulgar = &Language{
		ID:          uuid.New(),
		Name:        "bulgar",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Sabir = &Language{
		ID:          uuid.New(),
		Name:        "sabir",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Khazar = &Language{
		ID:          uuid.New(),
		Name:        "khazar",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Hunnic = &Language{
		ID:          uuid.New(),
		Name:        "hunnic",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Tuoba = &Language{
		ID:          uuid.New(),
		Name:        "tuoba",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Avar = &Language{
		ID:          uuid.New(),
		Name:        "avar",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Cuman = &Language{
		ID:          uuid.New(),
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
		ID:          uuid.New(),
		Name:        "dagur",
		Subfamily:   DagurSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Moghol = &Language{
		ID:          uuid.New(),
		Name:        "moghol",
		Subfamily:   MogholSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Khamnigan = &Language{
		ID:          uuid.New(),
		Name:        "khamnigan",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Buryat = &Language{
		ID:          uuid.New(),
		Name:        "buryat",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Mongolian = &Language{
		ID:          uuid.New(),
		Name:        "mongolian",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Khalkha = &Language{
		ID:          uuid.New(),
		Name:        "khalkha",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Chakhar = &Language{
		ID:          uuid.New(),
		Name:        "chakhar",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Khorchin = &Language{
		ID:          uuid.New(),
		Name:        "khorchin",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Ordos = &Language{
		ID:          uuid.New(),
		Name:        "ordos",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Oirat = &Language{
		ID:          uuid.New(),
		Name:        "oirat",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Monguor = &Language{
		ID:          uuid.New(),
		Name:        "monguor",
		Subfamily:   SouthernMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Bonan = &Language{
		ID:          uuid.New(),
		Name:        "bonan",
		Subfamily:   SouthernMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Santa = &Language{
		ID:          uuid.New(),
		Name:        "santa",
		Subfamily:   SouthernMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Kangjia = &Language{
		ID:          uuid.New(),
		Name:        "kangjia",
		Subfamily:   SouthernMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Khitan = &Language{
		ID:          uuid.New(),
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
		ID:          uuid.New(),
		Name:        "manchu",
		Subfamily:   JurchenicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Xibe = &Language{
		ID:          uuid.New(),
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
		ID:          uuid.New(),
		Name:        "tamil",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Malayalam = &Language{
		ID:          uuid.New(),
		Name:        "malayalam",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Irula = &Language{
		ID:          uuid.New(),
		Name:        "irula",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kodava = &Language{
		ID:          uuid.New(),
		Name:        "kodava",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kurumba = &Language{
		ID:          uuid.New(),
		Name:        "kurumba",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Toda = &Language{
		ID:          uuid.New(),
		Name:        "toda",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kota = &Language{
		ID:          uuid.New(),
		Name:        "kota",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kannada = &Language{
		ID:          uuid.New(),
		Name:        "kannada",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Badaga = &Language{
		ID:          uuid.New(),
		Name:        "badaga",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Telugu = &Language{
		ID:          uuid.New(),
		Name:        "telugu",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Chenchu = &Language{
		ID:          uuid.New(),
		Name:        "chenchu",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Konda = &Language{
		ID:          uuid.New(),
		Name:        "konda",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	MukhaDora = &Language{
		ID:          uuid.New(),
		Name:        "mukha_dora",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Manda = &Language{
		ID:          uuid.New(),
		Name:        "manda",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Pengo = &Language{
		ID:          uuid.New(),
		Name:        "pengo",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kuvi = &Language{
		ID:          uuid.New(),
		Name:        "kuvi",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kui = &Language{
		ID:          uuid.New(),
		Name:        "kui",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Gondi = &Language{
		ID:          uuid.New(),
		Name:        "gondi",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kolami = &Language{
		ID:          uuid.New(),
		Name:        "kolami",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Naiki = &Language{
		ID:          uuid.New(),
		Name:        "naiki",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Gadaba = &Language{
		ID:          uuid.New(),
		Name:        "gadaba",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Ollari = &Language{
		ID:          uuid.New(),
		Name:        "ollari",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kondekor = &Language{
		ID:          uuid.New(),
		Name:        "kondekor",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Duruwa = &Language{
		ID:          uuid.New(),
		Name:        "duruwa",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Oraon = &Language{
		ID:          uuid.New(),
		Name:        "oraon",
		Subfamily:   NorthernDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kisan = &Language{
		ID:          uuid.New(),
		Name:        "kisan",
		Subfamily:   NorthernDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	KumarbhagPaharia = &Language{
		ID:          uuid.New(),
		Name:        "kumarbhag_paharia",
		Subfamily:   NorthernDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	SauriaPaharia = &Language{
		ID:          uuid.New(),
		Name:        "sauria_paharia",
		Subfamily:   NorthernDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Brahui = &Language{
		ID:          uuid.New(),
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
		ID:          uuid.New(),
		Name:        "kru",
		Subfamily:   KruSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Kwa = &Language{
		ID:          uuid.New(),
		Name:        "kwa",
		Subfamily:   VoltaCongoSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Gur = &Language{
		ID:          uuid.New(),
		Name:        "gur",
		Subfamily:   VoltaCongoSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Soninke = &Language{
		ID:          uuid.New(),
		Name:        "soninke",
		Subfamily:   MandeSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Manding = &Language{
		ID:          uuid.New(),
		Name:        "manding",
		Subfamily:   MandeSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Senegambian = &Language{
		ID:          uuid.New(),
		Name:        "senegambian",
		Subfamily:   AtlanticCongoSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Swahili = &Language{
		ID:          uuid.New(),
		Name:        "swahili",
		Subfamily:   AtlanticCongoSubfamily,
		IsLiving:    true,
		WordBaseRef: SwahiliWordBaseRef,
	}
	Tebu = &Language{
		ID:          uuid.New(),
		Name:        "tebu",
		Subfamily:   SaharanSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Hausa = &Language{
		ID:          uuid.New(),
		Name:        "hausa",
		Subfamily:   WestChadicSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Nobiin = &Language{
		ID:          uuid.New(),
		Name:        "nobiin",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Kenzi = &Language{
		ID:          uuid.New(),
		Name:        "kenzi",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Midob = &Language{
		ID:          uuid.New(),
		Name:        "midob",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Birgid = &Language{
		ID:          uuid.New(),
		Name:        "birgid",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	HillNubian = &Language{
		ID:          uuid.New(),
		Name:        "hill_nubian",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Daju = &Language{
		ID:          uuid.New(),
		Name:        "daju",
		Subfamily:   DajuSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Somali = &Language{
		ID:          uuid.New(),
		Name:        "somali",
		Subfamily:   SomalicSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	// Songhay
	Korandje = &Language{
		ID:          uuid.New(),
		Name:        "korandje",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	KoyraChiini = &Language{
		ID:          uuid.New(),
		Name:        "koyra_chiini",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tadaksahak = &Language{
		ID:          uuid.New(),
		Name:        "tadaksahak",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tasawaq = &Language{
		ID:          uuid.New(),
		Name:        "tasawaq",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tagdal = &Language{
		ID:          uuid.New(),
		Name:        "tagdal",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	TondiSongwayKiini = &Language{
		ID:          uuid.New(),
		Name:        "tondi_songway_kiini",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	HumburiSenni = &Language{
		ID:          uuid.New(),
		Name:        "humburi_senni",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	KoyraboroSenni = &Language{
		ID:          uuid.New(),
		Name:        "koyraboro_senni",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Zarma = &Language{
		ID:          uuid.New(),
		Name:        "zarma",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	SonghoyboroCiine = &Language{
		ID:          uuid.New(),
		Name:        "songhoyboro_ciine",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Dendi = &Language{
		ID:          uuid.New(),
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
		ID:          uuid.New(),
		Name:        "oodham",
		Subfamily:   TepimanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	PimaBajo = &Language{
		ID:          uuid.New(),
		Name:        "pima_bajo",
		Subfamily:   TepimanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Tepehuan = &Language{
		ID:          uuid.New(),
		Name:        "tepehuan",
		Subfamily:   TepimanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Southern = &Language{
		ID:          uuid.New(),
		Name:        "southern",
		Subfamily:   TepimanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Tepecano = &Language{
		ID:          uuid.New(),
		Name:        "tepecano",
		Subfamily:   TepimanSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Tarahumara = &Language{
		ID:          uuid.New(),
		Name:        "tarahumara",
		Subfamily:   TarahumaranSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	UpriverGuarijio = &Language{
		ID:          uuid.New(),
		Name:        "upriver_guarijio",
		Subfamily:   TarahumaranSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	DownriverGuarijio = &Language{
		ID:          uuid.New(),
		Name:        "downriver_guarijio",
		Subfamily:   TarahumaranSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Tubar = &Language{
		ID:          uuid.New(),
		Name:        "tubar",
		Subfamily:   TarahumaranSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Yaqui = &Language{
		ID:          uuid.New(),
		Name:        "yaqui",
		Subfamily:   CahitaSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Mayo = &Language{
		ID:          uuid.New(),
		Name:        "mayo",
		Subfamily:   CahitaSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Opata = &Language{
		ID:          uuid.New(),
		Name:        "opata",
		Subfamily:   OpatanSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Eudeve = &Language{
		ID:          uuid.New(),
		Name:        "eudeve",
		Subfamily:   OpatanSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Cora = &Language{
		ID:          uuid.New(),
		Name:        "Cora",
		Subfamily:   CoracholSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Huichol = &Language{
		ID:          uuid.New(),
		Name:        "Huichol",
		Subfamily:   CoracholSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Pochutec = &Language{
		ID:          uuid.New(),
		Name:        "pochutec",
		Subfamily:   AztecanSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Pipil = &Language{
		ID:          uuid.New(),
		Name:        "pipil",
		Subfamily:   AztecanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Nahuatl = &Language{
		ID:          uuid.New(),
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
		ID:          uuid.New(),
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
		ID:          uuid.New(),
		Name:        "Ancash",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	AltoPativilcaAltoMaranonAltoHuallaga = &Language{
		ID:          uuid.New(),
		Name:        "alto_pativilca_alto_maranon_alto_huallaga",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Yaru = &Language{
		ID:          uuid.New(),
		Name:        "yaru",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Wanka = &Language{
		ID:          uuid.New(),
		Name:        "wanka",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	YauyosChincha = &Language{
		ID:          uuid.New(),
		Name:        "yauyos_chincha",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Pacaraos = &Language{
		ID:          uuid.New(),
		Name:        "pacaraos",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Lambayeque = &Language{
		ID:          uuid.New(),
		Name:        "lambayeque",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Cajamarca = &Language{
		ID:          uuid.New(),
		Name:        "cajamarca",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Lincha = &Language{
		ID:          uuid.New(),
		Name:        "lincha",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Laraos = &Language{
		ID:          uuid.New(),
		Name:        "laraos",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Ayacucho = &Language{
		ID:          uuid.New(),
		Name:        "ayacucho",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Cusco = &Language{
		ID:          uuid.New(),
		Name:        "cusco",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Puno = &Language{
		ID:          uuid.New(),
		Name:        "puno",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	NorthernBolivian = &Language{
		ID:          uuid.New(),
		Name:        "northern_bolivian",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	SouthernBolivia = &Language{
		ID:          uuid.New(),
		Name:        "southern_bolivia",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	SantiagoDelEstero = &Language{
		ID:          uuid.New(),
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
		ID:          uuid.New(),
		Name:        "kabyle",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tamazight = &Language{
		ID:          uuid.New(),
		Name:        "tamazight",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Shilha = &Language{
		ID:          uuid.New(),
		Name:        "shilha",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	SenhajaDeSrair = &Language{
		ID:          uuid.New(),
		Name:        "senhaja_de_srair",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Ghomara = &Language{
		ID:          uuid.New(),
		Name:        "ghomara",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Riffian = &Language{
		ID:          uuid.New(),
		Name:        "riffian",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	AytSeghrouchen = &Language{
		ID:          uuid.New(),
		Name:        "ayt_seghrouchen",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	AytWarayn = &Language{
		ID:          uuid.New(),
		Name:        "ayt_warayn",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Shenwa = &Language{
		ID:          uuid.New(),
		Name:        "shenwa",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Shawiya = &Language{
		ID:          uuid.New(),
		Name:        "shawiya",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Zenaga = &Language{
		ID:          uuid.New(),
		Name:        "zenaga",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Siwi = &Language{
		ID:          uuid.New(),
		Name:        "siwi",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Nafusi = &Language{
		ID:          uuid.New(),
		Name:        "nafusi",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Sokna = &Language{
		ID:          uuid.New(),
		Name:        "sokna",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Ghadames = &Language{
		ID:          uuid.New(),
		Name:        "ghadames",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Awjila = &Language{
		ID:          uuid.New(),
		Name:        "awjila",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tuareg = &Language{
		ID:          uuid.New(),
		Name:        "tuareg",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Numidian = &Language{
		ID:          uuid.New(),
		Name:        "numidian",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	// Cushitic
	Beja = &Language{
		ID:          uuid.New(),
		Name:        "beja",
		Subfamily:   CushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Awngi = &Language{
		ID:          uuid.New(),
		Name:        "awngi",
		Subfamily:   AgawSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Bilen = &Language{
		ID:          uuid.New(),
		Name:        "bilen",
		Subfamily:   AgawSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Qimant = &Language{
		ID:          uuid.New(),
		Name:        "qimant",
		Subfamily:   AgawSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Xamtanga = &Language{
		ID:          uuid.New(),
		Name:        "xamtanga",
		Subfamily:   AgawSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Gawwada = &Language{
		ID:          uuid.New(),
		Name:        "gawwada",
		Subfamily:   DullaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tsamai = &Language{
		ID:          uuid.New(),
		Name:        "tsamai",
		Subfamily:   DullaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Dihina = &Language{
		ID:          uuid.New(),
		Name:        "dihina",
		Subfamily:   DullaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Dobase = &Language{
		ID:          uuid.New(),
		Name:        "dobase",
		Subfamily:   DullaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Burji = &Language{
		ID:          uuid.New(),
		Name:        "burji",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Sidamo = &Language{
		ID:          uuid.New(),
		Name:        "sidamo",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Gedeo = &Language{
		ID:          uuid.New(),
		Name:        "gedeo",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	HadiyyaLibido = &Language{
		ID:          uuid.New(),
		Name:        "hadiyya_libido",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	KambaataAlaba = &Language{
		ID:          uuid.New(),
		Name:        "kambaata_alaba",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	SahoAfar = &Language{
		ID:          uuid.New(),
		Name:        "saho_afar",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	OmoTana = &Language{
		ID:          uuid.New(),
		Name:        "omo_tana",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Oromoid = &Language{
		ID:          uuid.New(),
		Name:        "oromoid",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Dullay = &Language{
		ID:          uuid.New(),
		Name:        "dullay",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Yaaku = &Language{
		ID:          uuid.New(),
		Name:        "yaaku",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Gorowa = &Language{
		ID:          uuid.New(),
		Name:        "gorowa",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Iraqw = &Language{
		ID:          uuid.New(),
		Name:        "iraqw",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Alagwa = &Language{
		ID:          uuid.New(),
		Name:        "alagwa",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Burunge = &Language{
		ID:          uuid.New(),
		Name:        "burunge",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Aasax = &Language{
		ID:          uuid.New(),
		Name:        "aasax",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Kwadza = &Language{
		ID:          uuid.New(),
		Name:        "kwadza",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Egyptian = &Language{
		ID:          uuid.New(),
		Name:        "egyptian",
		Subfamily:   EgyptianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	// semitic
	Akkadian = &Language{
		ID:          uuid.New(),
		Name:        "akkadian",
		Subfamily:   EastSemiticSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Eblaite = &Language{
		ID:          uuid.New(),
		Name:        "eblaite",
		Subfamily:   EastSemiticSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Kishite = &Language{
		ID:          uuid.New(),
		Name:        "kishite",
		Subfamily:   EastSemiticSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Ashkenazi = &Language{
		ID:          uuid.New(),
		Name:        "ashkenazi",
		Subfamily:   HebrewSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Sephardi = &Language{
		ID:          uuid.New(),
		Name:        "sephardi",
		Subfamily:   HebrewSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Hebrew = &Language{
		ID:          uuid.New(),
		Name:        "hebrew",
		Subfamily:   HebrewSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Samaritan = &Language{
		ID:          uuid.New(),
		Name:        "samaritan",
		Subfamily:   HebrewSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Ammonite = &Language{
		ID:          uuid.New(),
		Name:        "ammonite",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Edomite = &Language{
		ID:          uuid.New(),
		Name:        "edomite",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Moabite = &Language{
		ID:          uuid.New(),
		Name:        "moabite",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Phoenician = &Language{
		ID:          uuid.New(),
		Name:        "phoenician",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Punic = &Language{
		ID:          uuid.New(),
		Name:        "punic",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Samalian = &Language{
		ID:          uuid.New(),
		Name:        "samalian",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Aramaic = &Language{
		ID:          uuid.New(),
		Name:        "aramaic",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Nabataean = &Language{
		ID:          uuid.New(),
		Name:        "nabataean",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Amorite = &Language{
		ID:          uuid.New(),
		Name:        "amorite",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Himyaritic = &Language{
		ID:          uuid.New(),
		Name:        "himyaritic",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Sutean = &Language{
		ID:          uuid.New(),
		Name:        "sutean",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Syriac = &Language{
		ID:          uuid.New(),
		Name:        "syriac",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Arabic = &Language{
		ID:          uuid.New(),
		Name:        "arabic",
		Subfamily:   ArabicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Geez = &Language{
		ID:          uuid.New(),
		Name:        "geez",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Tigrinya = &Language{
		ID:          uuid.New(),
		Name:        "tigrinya",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Amharic = &Language{
		ID:          uuid.New(),
		Name:        "amharic",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Argobba = &Language{
		ID:          uuid.New(),
		Name:        "argobba",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Harari = &Language{
		ID:          uuid.New(),
		Name:        "harari",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Gafat = &Language{
		ID:          uuid.New(),
		Name:        "gafat",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Soddo = &Language{
		ID:          uuid.New(),
		Name:        "soddo",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Mesmes = &Language{
		ID:          uuid.New(),
		Name:        "mesmes",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Muher = &Language{
		ID:          uuid.New(),
		Name:        "muher",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	WestGurage = &Language{
		ID:          uuid.New(),
		Name:        "west_gurage",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Mesqan = &Language{
		ID:          uuid.New(),
		Name:        "mesqan",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	SebatBet = &Language{
		ID:          uuid.New(),
		Name:        "sebat_bet",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Baṭḥari = &Language{
		ID:          uuid.New(),
		Name:        "bathari",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Harsusi = &Language{
		ID:          uuid.New(),
		Name:        "harsusi",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Hobyot = &Language{
		ID:          uuid.New(),
		Name:        "hobyot",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Mehri = &Language{
		ID:          uuid.New(),
		Name:        "mehri",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Shehri = &Language{
		ID:          uuid.New(),
		Name:        "shehri",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Soqotri = &Language{
		ID:          uuid.New(),
		Name:        "soqotri",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Faifi = &Language{
		ID:          uuid.New(),
		Name:        "faifi",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Hadramautic = &Language{
		ID:          uuid.New(),
		Name:        "hadramautic",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Minaean = &Language{
		ID:          uuid.New(),
		Name:        "minaean",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Qatabanian = &Language{
		ID:          uuid.New(),
		Name:        "qatabanian",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Razihi = &Language{
		ID:          uuid.New(),
		Name:        "razihi",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Sabaean = &Language{
		ID:          uuid.New(),
		Name:        "sabaean",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	// Omotic
	SouthOmotic = &Language{
		ID:          uuid.New(),
		Name:        "south_omotic",
		IsLiving:    true,
		Subfamily:   OmoticSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Mao = &Language{
		ID:          uuid.New(),
		Name:        "mao",
		IsLiving:    true,
		Subfamily:   OmoticSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Dizoid = &Language{
		ID:          uuid.New(),
		Name:        "dizoid",
		IsLiving:    true,
		Subfamily:   OmoticSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Gonga = &Language{
		ID:          uuid.New(),
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
		ID:          uuid.New(),
		Name:        "common_eldarin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding common_eldarin to word bases
	}
	Quenya = &Language{
		ID:          uuid.New(),
		Name:        "quenya",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding quenya to word bases
	}
	Quendya = &Language{
		ID:          uuid.New(),
		Name:        "quendya",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding quendya to word bases
	}
	ExilicQuendya = &Language{
		ID:          uuid.New(),
		Name:        "exilic_quendya",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding exilic_quendya to word bases
	}
	CommonTelerin = &Language{
		ID:          uuid.New(),
		Name:        "common_telerin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding common_telerin to word bases
	}
	Telerin = &Language{
		ID:          uuid.New(),
		Name:        "telerin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding telerin to word bases
	}
	Sindarin = &Language{
		ID:          uuid.New(),
		Name:        "sindarin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef,
	}
	Nandorin = &Language{
		ID:          uuid.New(),
		Name:        "nandorin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding nandorin to word bases
	}
	Avarin = &Language{
		ID:          uuid.New(),
		Name:        "avarin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding avarin to word bases
	}
	Aldmeris = &Language{
		ID:          uuid.New(),
		Name:        "aldmeris",
		Subfamily:   AldmerisSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding aldmeris to word bases
	}
	Dunmeris = &Language{
		ID:          uuid.New(),
		Name:        "dunmeris",
		Subfamily:   AldmerisSubfamily,
		IsLiving:    true,
		WordBaseRef: DunmerisWordBaseRef,
	}
	Bosmeris = &Language{
		ID:          uuid.New(),
		Name:        "bosmeris",
		Subfamily:   AldmerisSubfamily,
		IsLiving:    true,
		WordBaseRef: DunmerisWordBaseRef, // @TODO can be changed after adding bosmeris to word bases
	}
	Dwemeris = &Language{
		ID:          uuid.New(),
		Name:        "dwemeris",
		Subfamily:   DwermerisSubfamily,
		IsLiving:    true,
		WordBaseRef: DwemerisWordBaseRef,
	}
	Goblins = &Language{
		ID:          uuid.New(),
		Name:        "goblins",
		Subfamily:   OrcishSubfamily,
		IsLiving:    true,
		WordBaseRef: GoblinsWordBaseRef,
	}
	Orcish = &Language{
		ID:          uuid.New(),
		Name:        "orcish",
		Subfamily:   OrcishSubfamily,
		IsLiving:    true,
		WordBaseRef: OrcWordBaseRef,
	}
	Giantish = &Language{
		ID:          uuid.New(),
		Name:        "giantish",
		Subfamily:   GiantishSubfamily,
		IsLiving:    true,
		WordBaseRef: GiantWordBaseRef,
	}
	Draconic = &Language{
		ID:          uuid.New(),
		Name:        "draconic",
		Subfamily:   DraconicSubfamily,
		IsLiving:    true,
		WordBaseRef: DraconicWordBaseRef,
	}
	Arachnid = &Language{
		ID:          uuid.New(),
		Name:        "arachnid",
		Subfamily:   ArachnidSubfamily,
		IsLiving:    true,
		WordBaseRef: ArachnidWordBaseRef,
	}
	Serpents = &Language{
		ID:          uuid.New(),
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

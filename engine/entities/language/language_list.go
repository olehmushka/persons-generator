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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "albanian",
		Subfamily:   AlbanianSubfamily,
		IsLiving:    true,
		WordBaseRef: AlbanianWordBaseRef,
	}
	Armenian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "armenian",
		Subfamily:   ArmenianSubfamily,
		IsLiving:    true,
		WordBaseRef: ArmenianWordBaseRef,
	}
	// Slavic
	Ruthenian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ruthenian",
		Subfamily:   RuthenianSubfamily,
		IsLiving:    false,
		WordBaseRef: RuthenianWordBaseRef,
	}
	Belarusian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "belarusian",
		Subfamily:   RuthenianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding belarusian to word bases
	}
	Ukrainian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ukrainian",
		Subfamily:   RuthenianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding ukrainian to word bases
	}
	Russian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "russian",
		Subfamily:   MoscovianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding russian to word bases
	}
	Slovenian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "slovenian",
		Subfamily:   WesternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding slovenian to word bases
	}
	Croatian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "croation",
		Subfamily:   WesternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding croation to word bases
	}
	Serbian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "serbian",
		Subfamily:   WesternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding serbian to word bases
	}
	Bosnian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "bosnian",
		Subfamily:   WesternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding bosnian to word bases
	}
	Bulgarian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "bulgarian",
		Subfamily:   EasternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding bulgarian to word bases
	}
	Macedonian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "macedonian",
		Subfamily:   EasternSouthSlavicSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding macedonian to word bases
	}
	Polish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "polish",
		Subfamily:   LechiticSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding polish to word bases
	}
	Silesian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "silesian",
		Subfamily:   LechiticSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding silesian to word bases
	}
	LowerSorbian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "lower_sorbian",
		Subfamily:   SorbianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding lower_sorbian to word bases
	}
	UpperSorbian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "upper_sorbian",
		Subfamily:   SorbianSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding upper_sorbian to word bases
	}
	Czech = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "czech",
		Subfamily:   CzechSlovakSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding czech to word bases
	}
	Slovak = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "slovak",
		Subfamily:   CzechSlovakSubfamily,
		IsLiving:    true,
		WordBaseRef: RuthenianWordBaseRef, // @TODO can be changed after adding slovak to word bases
	}
	// Baltic
	Latvian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "latvian",
		Subfamily:   EasternBalticSubfamily,
		IsLiving:    true,
		WordBaseRef: LithuanianWordBaseRef, // @TODO can be changed after adding latvian to word bases
	}
	Latgalian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "latgalian",
		Subfamily:   EasternBalticSubfamily,
		IsLiving:    true,
		WordBaseRef: LithuanianWordBaseRef, // @TODO can be changed after adding latgalian to word bases
	}
	Lithuanian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "lithuanian",
		Subfamily:   EasternBalticSubfamily,
		IsLiving:    true,
		WordBaseRef: LithuanianWordBaseRef,
	}
	Semigallian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "semigallian",
		Subfamily:   EasternBalticSubfamily,
		IsLiving:    false,
		WordBaseRef: LithuanianWordBaseRef, // @TODO can be changed after adding semigallian to word bases
	}
	// Celtic
	Celtiberian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "celtiberian",
		Subfamily:   CelticSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Gallaecian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "gallaecian",
		Subfamily:   CelticSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Noric = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "noric",
		Subfamily:   CelticSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Pictish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "pictish",
		Subfamily:   BrittonicSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Cumbric = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "cumbric",
		Subfamily:   BrittonicSubfamily,
		IsLiving:    false,
		WordBaseRef: CelticWordBaseRef,
	}
	Welsh = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "welsh",
		Subfamily:   WesternBrittonicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	Cornish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "cornish",
		Subfamily:   SouthWesternBrittonicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	Breton = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "breton",
		Subfamily:   SouthWesternBrittonicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	Irish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "irish",
		Subfamily:   GoidelicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	Manx = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "manx",
		Subfamily:   GoidelicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	ScottishGaelic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "scottish_gaelic",
		Subfamily:   GoidelicSubfamily,
		IsLiving:    true,
		WordBaseRef: CelticWordBaseRef,
	}
	// Dacian
	Dacian = &Language{
		Origin:    NativeOrigin,
		ID:        uuid.New().String(),
		Name:      "dacian",
		Subfamily: DacianSubfamily,
		IsLiving:  false,
	}
	// Germanic
	Danish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "danish",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding danish to word bases
	}
	Faroese = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "faroese",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding faroese to word bases
	}
	Icelandic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "icelandic",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding icelandic to word bases
	}
	Norwegian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "norwegian",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding norwegian to word bases
	}
	Swedish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "swedish",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef,
	}
	Dalecarlian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "dalecarlian",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding dalecarlian to word bases
	}
	Gutnish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "gutnish",
		Subfamily:   NorthGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: NordicWordBaseRef, // @TODO can be changed after adding gutnish to word bases
	}
	English = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "english",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: EnglishWordBaseRef,
	}
	Scots = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "scots",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: EnglishWordBaseRef, // @TODO can be changed after adding scots to word bases
	}
	Yola = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "yola",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: EnglishWordBaseRef, // @TODO can be changed after adding yola to word bases
	}
	Frisian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "frisian",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: EnglishWordBaseRef, // @TODO can be changed after adding frisian to word bases
	}
	German = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "german",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: GermanWordBaseRef,
	}
	Dutch = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "dutch",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding dutch to word bases
	}
	Luxembourgish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "luxembourgish",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding luxembourgish to word bases
	}
	Yiddish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "yiddish",
		Subfamily:   WestGermanicSubfamily,
		IsLiving:    true,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding yiddish to word bases
	}
	Gothic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "gothic",
		Subfamily:   EastGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding gothic to word bases
	}
	Vandalic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "vandalic",
		Subfamily:   EastGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding vandalic to word bases
	}
	Burgundian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "burgundian",
		Subfamily:   EastGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding burgundian to word bases
	}
	Alamannian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "alamannian",
		Subfamily:   ElbeGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding alamannian to word bases
	}
	Bavarian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "bavarian",
		Subfamily:   ElbeGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding bavarian to word bases
	}
	Langobardian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "lombardian",
		Subfamily:   ElbeGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding lombardian to word bases
	}
	Frankish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "frankish",
		Subfamily:   WeserRhineGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding frankish to word bases
	}
	Anglic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "anglic",
		Subfamily:   NorthSeaGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding anglic to word bases
	}
	Saxon = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "saxon",
		Subfamily:   NorthSeaGermanicSubfamily,
		IsLiving:    false,
		WordBaseRef: GermanWordBaseRef, // @TODO can be changed after adding saxon to word bases
	}
	// Greek
	Greek = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "greek",
		Subfamily:   GreekSubfamily,
		IsLiving:    true,
		WordBaseRef: GreekWordBaseRef,
	}
	// Illyrian
	Illyrian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "illyrian",
		Subfamily:   IllyrianSubfamily,
		IsLiving:    false,
		WordBaseRef: IllyrianWordBaseRef,
	}
	// IndoIranian
	Pashai = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "pashai",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Chitrali = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "chitrali",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Shina = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "shina",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Kohistani = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kohistani",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Kashmiri = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kashmiri",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Punjabi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "punjabi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Sindhi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sindhi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Rajasthani = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "rajasthani",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Gujarati = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "gujarati",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Bhili = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "bhili",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Khandeshi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "khandeshi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	HimachaliDogri = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "himachali_dogri",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	GirhwaliKumaoni = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "girhwali_kumaoni",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Nepali = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "nepali",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Hindi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "hindi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Bihari = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "bihari",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	BengaliAssamese = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "bengali_assamese",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Odia = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "odia",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	Halbi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "halbi",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	MarathiKonkani = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "marathi_konkani",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	SinhalaMaldivian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sinhala_maldivian",
		Subfamily:   IndoAryanSubfamily,
		IsLiving:    true,
		WordBaseRef: HindiWordBaseRef,
	}
	// Iranian
	Farsi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "farsi",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Avestan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "avestan",
		Subfamily:   IranianSubfamily,
		IsLiving:    false,
		WordBaseRef: IranianWordBaseRef,
	}
	Dari = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "dari",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Tajik = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tajik",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Pashto = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "pashto",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Kurdish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kurdish",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Luri = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "luri",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Balochi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "balochi",
		Subfamily:   IranianSubfamily,
		IsLiving:    true,
		WordBaseRef: IranianWordBaseRef,
	}
	Scythian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "scythian",
		Subfamily:   IranianSubfamily,
		IsLiving:    false,
		WordBaseRef: IranianWordBaseRef,
	}
	// Nuriastani
	KamkataVari = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kamkata_vari",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into KamkataVari when available
	}
	VasiVari = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "vasi_vari",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into VasiVari when available
		IsLiving:    true,
	}
	Askunu = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "askunu",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into Askunu when available
		IsLiving:    true,
	}
	Waigali = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "waigali",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into Waigali when available
		IsLiving:    true,
	}
	Tregami = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tregami",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into Tregami when available
		IsLiving:    true,
	}
	Zemiaki = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "zemiaki",
		Subfamily:   NuriastaniSubfamily,
		WordBaseRef: KannadaWordBaseRef, // @TODO change it into Zemiaki when available
		IsLiving:    true,
	}
	// Italic
	Venetic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "venetic",
		Subfamily:   ItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into venetic when available
		IsLiving:    false,
	}
	Sicel = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sicel",
		Subfamily:   ItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into sicel when available
		IsLiving:    false,
	}
	Lusitanian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "lusitanian",
		Subfamily:   ItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into lusitanian when available
		IsLiving:    false,
	}
	Latin = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "latin",
		Subfamily:   LatinoFaliscanSubfamily,
		WordBaseRef: LatinWordBaseRef, // @TODO change it into latin when available
		IsLiving:    true,
	}
	Faliscan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "faliscan",
		Subfamily:   LatinoFaliscanSubfamily,
		WordBaseRef: LatinWordBaseRef, // @TODO change it into faliscan when available
		IsLiving:    false,
	}
	Lanuvain = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "lanuvain",
		Subfamily:   LatinoFaliscanSubfamily,
		WordBaseRef: LatinWordBaseRef, // @TODO change it into lanuvain when available
		IsLiving:    false,
	}
	Venetian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "venetian",
		Subfamily:   RomanceSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into venetian when available
		IsLiving:    true,
	}
	Sardinian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sardinian",
		Subfamily:   RomanceSubfamily,
		WordBaseRef: ItalianWordBaseRef, // @TODO change it into sardinian when available
		IsLiving:    true,
	}
	Portuguese = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "portuguese",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: PortugueseWordBaseRef,
		IsLiving:    true,
	}
	Galician = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "galician",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: PortugueseWordBaseRef,
		IsLiving:    true,
	}
	Asturleonese = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "asturleonese",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Mirandese = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "mirandese",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: PortugueseWordBaseRef,
		IsLiving:    true,
	}
	Spanish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "spanish",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Aragonese = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "aragonese",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Ladino = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ladino",
		Subfamily:   IberoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Catalan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "catalan",
		Subfamily:   OccitanoRomanceSubfamily,
		WordBaseRef: SpanishWordBaseRef,
		IsLiving:    true,
	}
	Occitan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "occitan",
		Subfamily:   OccitanoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Gascon = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "gascon",
		Subfamily:   OccitanoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	French = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "french",
		Subfamily:   GalloRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	FrancoProvencal = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "franco_provencal",
		Subfamily:   GalloRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Romansh = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "romansh",
		Subfamily:   RhaetoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Ladin = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ladin",
		Subfamily:   RhaetoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Friulian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "friulian",
		Subfamily:   RhaetoRomanceSubfamily,
		WordBaseRef: FrenchWordBaseRef,
		IsLiving:    true,
	}
	Piedmontense = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "piedmontense",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Ligurian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ligurian",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Lombard = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "lombard",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Emilian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "emilian",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Romagnol = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "romagnol",
		Subfamily:   GalloItalicSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Italian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "italian",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Sicilian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sicilian",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Neapolitan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "neapolitan",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Dalmatian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "dalmatian",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Istriot = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "istriot",
		Subfamily:   ItaloDalmatianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    true,
	}
	Romanian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "romanian",
		Subfamily:   EasternRomanceSubfamily,
		WordBaseRef: LatinWordBaseRef,
		IsLiving:    true,
	}
	Aromanian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "aromanian",
		Subfamily:   EasternRomanceSubfamily,
		WordBaseRef: LatinWordBaseRef,
		IsLiving:    true,
	}
	MaglenoRomanian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "magleno_romanian",
		Subfamily:   EasternRomanceSubfamily,
		WordBaseRef: LatinWordBaseRef,
		IsLiving:    true,
	}
	IstroRomanian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "istro_romanian",
		Subfamily:   EasternRomanceSubfamily,
		WordBaseRef: LatinWordBaseRef,
		IsLiving:    true,
	}
	Oscan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "oscan",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Umbrian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "umbrian",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Volscian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "volscian",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Sabine = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sabine",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	SouthPicene = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "south_picene",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Marsian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "marsian",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Paeligni = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "paeligni",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Hernican = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "hernican",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	Marrucinian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "marrucinian",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	PreSamnite = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "pre_samnite",
		Subfamily:   OscoUmbrianSubfamily,
		WordBaseRef: ItalianWordBaseRef,
		IsLiving:    false,
	}
	// Tysenian
	Rhaetic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "rhaetic",
		Subfamily:   TysenianSubfamily,
		WordBaseRef: EtruscanWordBaseRef,
		IsLiving:    false,
	}
	Etruscan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "etruscan",
		Subfamily:   TysenianSubfamily,
		WordBaseRef: EtruscanWordBaseRef,
		IsLiving:    false,
	}
	// Tocharian
	Tocharian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "japanese",
		Subfamily:   JaponicSubfamily,
		IsLiving:    true,
		WordBaseRef: JapaneseWordBaseRef,
	}
	Korean = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "Georgian",
		Subfamily:   KartoZanSubfamily,
		WordBaseRef: GeorgianWordBaseRef,
		IsLiving:    true,
	}
	Mingrelian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "Mingrelian",
		Subfamily:   KartoZanSubfamily,
		WordBaseRef: GeorgianWordBaseRef,
		IsLiving:    true,
	}
	Laz = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "Laz",
		Subfamily:   KartoZanSubfamily,
		WordBaseRef: GeorgianWordBaseRef,
		IsLiving:    true,
	}
	Svan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "chinese",
		Subfamily:   ChineseSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	Cantonese = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "cantonese",
		Subfamily:   ChineseSubfamily,
		IsLiving:    true,
		WordBaseRef: CantoneseWordBaseRef,
	}
	Bai = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "bai",
		Subfamily:   GreaterBaiSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	Caijia = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "caijia",
		Subfamily:   GreaterBaiSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	Longjia = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "longjia",
		Subfamily:   GreaterBaiSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	Luren = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "luren",
		Subfamily:   GreaterBaiSubfamily,
		IsLiving:    true,
		WordBaseRef: ChineseWordBaseRef,
	}
	// LoloBurmese
	Muangphe = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "muangphe",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mango = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "mango",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Manga = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "manga",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Maang = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "maang",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mondzi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "mondzi",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Maza = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "maza",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mauphu = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "mauphu",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Motang = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "motang",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mongphu = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "mongphu",
		Subfamily:   MondzishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Burmese = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "burmish",
		Subfamily:   BurmishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Jinuo = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "jinuo",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Sangkong = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sangkong",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Bisu = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "bisu",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Phunoi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "phunoi",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Pyen = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "pyen",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Sila = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sila",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Phana = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "phana",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Akeu = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "akeu",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Hani = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "hani",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Piyo = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "piyo",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Enu = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "enu",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Mpi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "mpi",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Kaduo = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kaduo",
		Subfamily:   HanoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Laha = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "laha",
		Subfamily:   LahoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Kucong = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kucong",
		Subfamily:   LahoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Namuyi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "namuyi",
		Subfamily:   NaxishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Shixing = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "shixing",
		Subfamily:   NaxishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Naish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "naish",
		Subfamily:   NaxishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Nusu = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "nusu",
		Subfamily:   NusoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Zauzou = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "zauzou",
		Subfamily:   NusoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Katso = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "katso",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Samu = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "samu",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Sanie = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sanie",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Sadu = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sadu",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Meuma = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "meuma",
		Subfamily:   KazhuoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Lisu = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "lisu",
		Subfamily:   LisoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Lolopo = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "lolopo",
		Subfamily:   LisoishSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Tangut = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tangut",
		Subfamily:   QiangicSubfamily,
		WordBaseRef: ChineseWordBaseRef,
		IsLiving:    true,
	}
	Tibetic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "vietnamese",
		Subfamily:   VietMuongSubfamily,
		IsLiving:    true,
		WordBaseRef: VietnameseWordBaseRef,
	}
	Kri = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kri",
		Subfamily:   VieticSubfamily,
		IsLiving:    true,
		WordBaseRef: VietnameseWordBaseRef,
	}
	Khmer = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "mari",
		Subfamily:   FinnoPermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Merya = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "merya",
		Subfamily:   FinnoPermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Komi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "komi",
		Subfamily:   PermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Permyak = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "permyak",
		Subfamily:   PermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Udmurt = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "udmurt",
		Subfamily:   PermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Meshchera = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "meshchera",
		Subfamily:   PermicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Estonian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "estonian",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Livonian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "livonian",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Votic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "votic",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Finnish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "finnish",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Ingrian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ingrian",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Karelian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "karelian",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Ludic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ludic",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Veps = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "veps",
		Subfamily:   BaltoFinnicSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Sami = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sami",
		Subfamily:   SamiSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Erzya = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "erzya",
		Subfamily:   MordvinSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Moksha = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "moksha",
		Subfamily:   MordvinSubfamily,
		IsLiving:    true,
		WordBaseRef: FinnicWordBaseRef,
	}
	Hungarian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "hungarian",
		Subfamily:   UgricSubfamily,
		IsLiving:    true,
		WordBaseRef: HungarianWordBaseRef,
	}
	Khanty = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "khanty",
		Subfamily:   UgricSubfamily,
		IsLiving:    true,
		WordBaseRef: HungarianWordBaseRef,
	}
	Mansi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "gagauz",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Turkish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "turkish",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Azerbaijani = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "azerbaijani",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Turkmen = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "turkmen",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Qashqai = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "qashqai",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Chaharmahali = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "chaharmahali",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Khorasani = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "khorasani",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Salar = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "salar",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Bashkir = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "bashkir",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Tatar = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tatar",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	CrimeanTatar = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "crimean_tatar",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	KarachayBalkar = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "karachay_balkar",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Kumyk = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kumyk",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Karaim = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "karaim",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Krymchak = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "krymchak",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Urum = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "urum",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Kazakh = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kazakh",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Karakalpak = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "karakalpak",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Nogai = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "nogai",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	SiberianTatar = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "siberian_tatar",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Baraba = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "baraba",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	SouthernAltai = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "southern_altai",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Teleut = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "teleut",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Kyrgyz = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kyrgyz",
		Subfamily:   KipchakSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Uzbek = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "uzbek",
		Subfamily:   KarlukSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Uyghur = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "uyghur",
		Subfamily:   KarlukSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Aynu = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "aynu",
		Subfamily:   KarlukSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Ili = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ili",
		Subfamily:   KarlukSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Khalaj = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "khalaj",
		Subfamily:   ShazTurkicSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Chuvash = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "chuvash",
		Subfamily:   OghuzSubfamily,
		IsLiving:    true,
		WordBaseRef: TurkishWordBaseRef,
	}
	Bulgar = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "bulgar",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Sabir = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sabir",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Khazar = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "khazar",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Hunnic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "hunnic",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Tuoba = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tuoba",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Avar = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "avar",
		Subfamily:   OghuzSubfamily,
		IsLiving:    false,
		WordBaseRef: TurkishWordBaseRef,
	}
	Cuman = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "dagur",
		Subfamily:   DagurSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Moghol = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "moghol",
		Subfamily:   MogholSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Khamnigan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "khamnigan",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Buryat = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "buryat",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Mongolian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "mongolian",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Khalkha = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "khalkha",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Chakhar = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "chakhar",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Khorchin = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "khorchin",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Ordos = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ordos",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Oirat = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "oirat",
		Subfamily:   CentralMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Monguor = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "monguor",
		Subfamily:   SouthernMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Bonan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "bonan",
		Subfamily:   SouthernMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Santa = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "santa",
		Subfamily:   SouthernMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Kangjia = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kangjia",
		Subfamily:   SouthernMongolicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Khitan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "manchu",
		Subfamily:   JurchenicSubfamily,
		IsLiving:    true,
		WordBaseRef: MongolianWordBaseRef,
	}
	Xibe = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tamil",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Malayalam = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "malayalam",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Irula = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "irula",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kodava = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kodava",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kurumba = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kurumba",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Toda = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "toda",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kota = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kota",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kannada = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kannada",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Badaga = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "badaga",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Telugu = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "telugu",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Chenchu = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "chenchu",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Konda = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "konda",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	MukhaDora = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "mukha_dora",
		Subfamily:   SouthDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Manda = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "manda",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Pengo = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "pengo",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kuvi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kuvi",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kui = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kui",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Gondi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "gondi",
		Subfamily:   SouthCentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kolami = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kolami",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Naiki = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "naiki",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Gadaba = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "gadaba",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Ollari = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ollari",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kondekor = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kondekor",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Duruwa = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "duruwa",
		Subfamily:   CentralDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Oraon = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "oraon",
		Subfamily:   NorthernDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Kisan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kisan",
		Subfamily:   NorthernDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	KumarbhagPaharia = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kumarbhag_paharia",
		Subfamily:   NorthernDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	SauriaPaharia = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sauria_paharia",
		Subfamily:   NorthernDravidianSubfamily,
		IsLiving:    true,
		WordBaseRef: KannadaWordBaseRef,
	}
	Brahui = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kru",
		Subfamily:   KruSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Kwa = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kwa",
		Subfamily:   VoltaCongoSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Gur = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "gur",
		Subfamily:   VoltaCongoSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Soninke = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "soninke",
		Subfamily:   MandeSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Manding = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "manding",
		Subfamily:   MandeSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Senegambian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "senegambian",
		Subfamily:   AtlanticCongoSubfamily,
		IsLiving:    true,
		WordBaseRef: NigerianWordBaseRef,
	}
	Swahili = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "swahili",
		Subfamily:   AtlanticCongoSubfamily,
		IsLiving:    true,
		WordBaseRef: SwahiliWordBaseRef,
	}
	Tebu = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tebu",
		Subfamily:   SaharanSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Hausa = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "hausa",
		Subfamily:   WestChadicSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Nobiin = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "nobiin",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Kenzi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kenzi",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Midob = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "midob",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Birgid = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "birgid",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	HillNubian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "hill_nubian",
		Subfamily:   NubianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Daju = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "daju",
		Subfamily:   DajuSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Somali = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "somali",
		Subfamily:   SomalicSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	// Songhay
	Korandje = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "korandje",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	KoyraChiini = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "koyra_chiini",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tadaksahak = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tadaksahak",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tasawaq = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tasawaq",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tagdal = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tagdal",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	TondiSongwayKiini = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tondi_songway_kiini",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	HumburiSenni = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "humburi_senni",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	KoyraboroSenni = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "koyraboro_senni",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Zarma = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "zarma",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	SonghoyboroCiine = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "songhoyboro_ciine",
		Subfamily:   SonghaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Dendi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "oodham",
		Subfamily:   TepimanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	PimaBajo = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "pima_bajo",
		Subfamily:   TepimanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Tepehuan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tepehuan",
		Subfamily:   TepimanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Southern = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "southern",
		Subfamily:   TepimanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Tepecano = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tepecano",
		Subfamily:   TepimanSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Tarahumara = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tarahumara",
		Subfamily:   TarahumaranSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	UpriverGuarijio = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "upriver_guarijio",
		Subfamily:   TarahumaranSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	DownriverGuarijio = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "downriver_guarijio",
		Subfamily:   TarahumaranSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Tubar = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tubar",
		Subfamily:   TarahumaranSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Yaqui = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "yaqui",
		Subfamily:   CahitaSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Mayo = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "mayo",
		Subfamily:   CahitaSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Opata = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "opata",
		Subfamily:   OpatanSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Eudeve = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "eudeve",
		Subfamily:   OpatanSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Cora = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "Cora",
		Subfamily:   CoracholSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Huichol = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "Huichol",
		Subfamily:   CoracholSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Pochutec = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "pochutec",
		Subfamily:   AztecanSubfamily,
		IsLiving:    false,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Pipil = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "pipil",
		Subfamily:   AztecanSubfamily,
		IsLiving:    true,
		WordBaseRef: NahuatlWordBaseRef,
	}
	Nahuatl = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
	Origin:      NativeOrigin,
	Name:        "inuit",
	Subfamily:   EskimoSubfamily,
	IsLiving:    true,
	WordBaseRef: InuitWordBaseRef,
}

var AllEskimoAleutLanguages = []*Language{Inuit}

// Austonesian
var (
	Hawaiian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "Ancash",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	AltoPativilcaAltoMaranonAltoHuallaga = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "alto_pativilca_alto_maranon_alto_huallaga",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Yaru = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "yaru",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Wanka = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "wanka",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	YauyosChincha = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "yauyos_chincha",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Pacaraos = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "pacaraos",
		Subfamily:   QuechuaISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Lambayeque = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "lambayeque",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Cajamarca = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "cajamarca",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Lincha = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "lincha",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Laraos = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "laraos",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Ayacucho = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ayacucho",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Cusco = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "cusco",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	Puno = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "puno",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	NorthernBolivian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "northern_bolivian",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	SouthernBolivia = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "southern_bolivia",
		Subfamily:   QuechuaIISubfamily,
		IsLiving:    true,
		WordBaseRef: QuechuaWordBaseRef,
	}
	SantiagoDelEstero = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kabyle",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tamazight = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tamazight",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Shilha = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "shilha",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	SenhajaDeSrair = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "senhaja_de_srair",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Ghomara = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ghomara",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Riffian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "riffian",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	AytSeghrouchen = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ayt_seghrouchen",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	AytWarayn = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ayt_warayn",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Shenwa = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "shenwa",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Shawiya = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "shawiya",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Zenaga = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "zenaga",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Siwi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "siwi",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Nafusi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "nafusi",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Sokna = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sokna",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Ghadames = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ghadames",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Awjila = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "awjila",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tuareg = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tuareg",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Numidian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "numidian",
		Subfamily:   BerberSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	// Cushitic
	Beja = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "beja",
		Subfamily:   CushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Awngi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "awngi",
		Subfamily:   AgawSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Bilen = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "bilen",
		Subfamily:   AgawSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Qimant = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "qimant",
		Subfamily:   AgawSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Xamtanga = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "xamtanga",
		Subfamily:   AgawSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Gawwada = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "gawwada",
		Subfamily:   DullaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Tsamai = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tsamai",
		Subfamily:   DullaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Dihina = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "dihina",
		Subfamily:   DullaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Dobase = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "dobase",
		Subfamily:   DullaySubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Burji = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "burji",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Sidamo = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sidamo",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Gedeo = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "gedeo",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	HadiyyaLibido = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "hadiyya_libido",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	KambaataAlaba = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kambaata_alaba",
		Subfamily:   HighlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	SahoAfar = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "saho_afar",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	OmoTana = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "omo_tana",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Oromoid = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "oromoid",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Dullay = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "dullay",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Yaaku = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "yaaku",
		Subfamily:   LowlandEastCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Gorowa = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "gorowa",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Iraqw = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "iraqw",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Alagwa = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "alagwa",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Burunge = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "burunge",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Aasax = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "aasax",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Kwadza = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kwadza",
		Subfamily:   SouthCushiticSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	Egyptian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "egyptian",
		Subfamily:   EgyptianSubfamily,
		IsLiving:    true,
		WordBaseRef: BerberWordBaseRef,
	}
	// semitic
	Akkadian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "akkadian",
		Subfamily:   EastSemiticSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Eblaite = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "eblaite",
		Subfamily:   EastSemiticSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Kishite = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "kishite",
		Subfamily:   EastSemiticSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Ashkenazi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ashkenazi",
		Subfamily:   HebrewSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Sephardi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sephardi",
		Subfamily:   HebrewSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Hebrew = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "hebrew",
		Subfamily:   HebrewSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Samaritan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "samaritan",
		Subfamily:   HebrewSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Ammonite = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "ammonite",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Edomite = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "edomite",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Moabite = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "moabite",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Phoenician = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "phoenician",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Punic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "punic",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Samalian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "samalian",
		Subfamily:   CanaaniteSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Aramaic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "aramaic",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Nabataean = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "nabataean",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Amorite = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "amorite",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Himyaritic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "himyaritic",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Sutean = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sutean",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Syriac = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "syriac",
		Subfamily:   AramaicSubfamily,
		IsLiving:    true,
		WordBaseRef: MesopotamianWordBaseRef,
	}
	Arabic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "arabic",
		Subfamily:   ArabicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Geez = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "geez",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Tigrinya = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "tigrinya",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Amharic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "amharic",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Argobba = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "argobba",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Harari = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "harari",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Gafat = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "gafat",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Soddo = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "soddo",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Mesmes = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "mesmes",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Muher = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "muher",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	WestGurage = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "west_gurage",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Mesqan = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "mesqan",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	SebatBet = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sebat_bet",
		Subfamily:   EthiopicSubfamily,
		IsLiving:    true,
		WordBaseRef: ArabicWordBaseRef,
	}
	Baṭḥari = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "bathari",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Harsusi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "harsusi",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Hobyot = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "hobyot",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Mehri = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "mehri",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Shehri = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "shehri",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Soqotri = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "soqotri",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Faifi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "faifi",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Hadramautic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "hadramautic",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Minaean = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "minaean",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Qatabanian = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "qatabanian",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Razihi = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "razihi",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Sabaean = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sabaean",
		IsLiving:    true,
		Subfamily:   SouthArabianSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	// Omotic
	SouthOmotic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "south_omotic",
		IsLiving:    true,
		Subfamily:   OmoticSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Mao = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "mao",
		IsLiving:    true,
		Subfamily:   OmoticSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Dizoid = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "dizoid",
		IsLiving:    true,
		Subfamily:   OmoticSubfamily,
		WordBaseRef: ArabicWordBaseRef,
	}
	Gonga = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "common_eldarin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding common_eldarin to word bases
	}
	Quenya = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "quenya",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding quenya to word bases
	}
	Quendya = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "quendya",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding quendya to word bases
	}
	ExilicQuendya = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "exilic_quendya",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding exilic_quendya to word bases
	}
	CommonTelerin = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "common_telerin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding common_telerin to word bases
	}
	Telerin = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "telerin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding telerin to word bases
	}
	Sindarin = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "sindarin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef,
	}
	Nandorin = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "nandorin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding nandorin to word bases
	}
	Avarin = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "avarin",
		Subfamily:   QuendianSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding avarin to word bases
	}
	Aldmeris = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "aldmeris",
		Subfamily:   AldmerisSubfamily,
		IsLiving:    true,
		WordBaseRef: SindarinWordBaseRef, // @TODO can be changed after adding aldmeris to word bases
	}
	Dunmeris = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "dunmeris",
		Subfamily:   AldmerisSubfamily,
		IsLiving:    true,
		WordBaseRef: DunmerisWordBaseRef,
	}
	Bosmeris = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "bosmeris",
		Subfamily:   AldmerisSubfamily,
		IsLiving:    true,
		WordBaseRef: DunmerisWordBaseRef, // @TODO can be changed after adding bosmeris to word bases
	}
	Dwemeris = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "dwemeris",
		Subfamily:   DwermerisSubfamily,
		IsLiving:    true,
		WordBaseRef: DwemerisWordBaseRef,
	}
	Goblins = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "goblins",
		Subfamily:   OrcishSubfamily,
		IsLiving:    true,
		WordBaseRef: GoblinsWordBaseRef,
	}
	Orcish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "orcish",
		Subfamily:   OrcishSubfamily,
		IsLiving:    true,
		WordBaseRef: OrcWordBaseRef,
	}
	Giantish = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "giantish",
		Subfamily:   GiantishSubfamily,
		IsLiving:    true,
		WordBaseRef: GiantWordBaseRef,
	}
	Draconic = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "draconic",
		Subfamily:   DraconicSubfamily,
		IsLiving:    true,
		WordBaseRef: DraconicWordBaseRef,
	}
	Arachnid = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
		Name:        "arachnid",
		Subfamily:   ArachnidSubfamily,
		IsLiving:    true,
		WordBaseRef: ArachnidWordBaseRef,
	}
	Serpents = &Language{
		Origin:      NativeOrigin,
		ID:          uuid.New().String(),
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

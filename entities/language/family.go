package language

type Family string

const (
	NigerCongoFamily    Family = "niger_congo_lang_family"
	AustronesianFamily  Family = "austronesian_lang_family"
	SinoTibetanFamily   Family = "sino_tibetan_lang_family"
	IndoEuropeanFamily  Family = "indo_european_lang_family"
	KartvelianFamily    Family = "kartvelian_lang_family"
	BasqueFamily        Family = "basque_lang_family"
	JaponicFamily       Family = "japonic_lang_family"
	KoreanicFamily      Family = "koreanic_lang_family"
	UralicFamily        Family = "uralic_lang_family"
	TurkicFamily        Family = "turkic_lang_family"
	MongolicFamily      Family = "mongolic_lang_family"
	AfroAsiaticFamily   Family = "afro_asiatic_lang_family"
	NiloSaharanFamily   Family = "nilo_saharan_lang_family"
	OtoMangueanFamily   Family = "oto_manguean_lang_family"
	UtoAztecanFamily    Family = "uto_aztecan_lang_family"
	AustroasiaticFamily Family = "austroasiatic_lang_family"
	TaiKadaiFamily      Family = "tai_kadai_lang_family"
	DravidianFamily     Family = "dravidian_lang_family"
	TupianFamily        Family = "tupian_lang_family"
	EskimoAleutFamily   Family = "eskimo_aleut_lang_family"
	QuechuaFamily       Family = "quechua_lang_family"
	// Fantasy
	ElvenFamily    Family = "elven_lang_family"
	OrcishFamily   Family = "orcish_lang_family"
	GiantishFamily Family = "giantish_lang_family"
	DraconicFamily Family = "draconic_lang_family"
	ArachnidFamily Family = "arachnid_lang_family"
	SerpentsFamily Family = "serpents_lang_family"
)

var AllFamilies = []Family{
	NigerCongoFamily,
	AustronesianFamily,
	SinoTibetanFamily,
	IndoEuropeanFamily,
	KartvelianFamily,
	BasqueFamily,
	JaponicFamily,
	KoreanicFamily,
	UralicFamily,
	TurkicFamily,
	MongolicFamily,
	AfroAsiaticFamily,
	NiloSaharanFamily,
	OtoMangueanFamily,
	UtoAztecanFamily,
	AustroasiaticFamily,
	TaiKadaiFamily,
	DravidianFamily,
	TupianFamily,
	EskimoAleutFamily,
	QuechuaFamily,
	ElvenFamily,
	OrcishFamily,
	GiantishFamily,
	DraconicFamily,
	ArachnidFamily,
	SerpentsFamily,
}

package culture

import "persons_generator/tools"

var AllAbsCultures = tools.Merge(AllAncientAbsCultures, AllMedievalAbsCultures)

// Ancient
var (
	AbsAncientBelgae = &AbstructCulture{
		Name: "abs_ancient_belgae",
		Root: EuropeContinentalRoot,
	}
	AbsAncientCeltIberian = &AbstructCulture{
		Name: "abs_ancient_celt_iberian",
		Root: EuropeContinentalRoot,
	}
	AbsAncientGaelic = &AbstructCulture{
		Name: "abs_ancient_gaelic",
		Root: EuropeContinentalRoot,
	}
	AbsAncientGallic = &AbstructCulture{
		Name: "abs_ancient_gallic",
		Root: EuropeContinentalRoot,
	}
	AbsAncientGermanic = &AbstructCulture{
		Name: "abs_ancient_germanic",
		Root: EuropeContinentalRoot,
	}
	AbsAncientIberian = &AbstructCulture{
		Name: "abs_ancient_iberian",
		Root: MediterraneanRoot,
	}
	AbsAncientOccidental = &AbstructCulture{
		Name: "abs_ancient_occidental",
		Root: EuropeContinentalRoot,
	}
	AbsAncientPannonian = &AbstructCulture{
		Name: "abs_ancient_pannonian",
		Root: EuropeContinentalRoot,
	}
	AbsAncientPretani = &AbstructCulture{
		Name: "abs_ancient_pretani",
		Root: EuropeContinentalRoot,
	}
	AbsAncientVeneti = &AbstructCulture{
		Name: "abs_ancient_veneti",
		Root: EuropeContinentalRoot,
	}
	AbsAncientItalic = &AbstructCulture{
		Name: "abs_ancient_italic",
		Root: MediterraneanRoot,
	}
	AbsAncientDacian = &AbstructCulture{
		Name: "abs_ancient_dacian",
		Root: EuropeContinentalRoot,
	}
	AbsAncientHellenistic = &AbstructCulture{
		Name: "abs_ancient_hellenistic",
		Root: MediterraneanRoot,
	}
	AbsAncientIllyrian = &AbstructCulture{
		Name: "abs_ancient_illyrian",
		Root: MediterraneanRoot,
	}
	AbsAncientNumidian = &AbstructCulture{
		Name: "abs_ancient_numidian",
		Root: AfricanRoot,
	}
	AbsAncientLevantine = &AbstructCulture{
		Name: "abs_ancient_levantine",
		Root: SemiticRoot,
	}
	AbsAncientAksumite = &AbstructCulture{
		Name: "abs_ancient_aksumite",
		Root: EgyptianRoot,
	}
	AbsAncientArabian = &AbstructCulture{
		Name: "abs_ancient_arabian",
		Root: SemiticRoot,
	}
	AbsAncientEgyptian = &AbstructCulture{
		Name: "abs_ancient_egyptian",
		Root: EgyptianRoot,
	}
	AbsAncientNubian = &AbstructCulture{
		Name: "abs_ancient_nubian",
		Root: AfricanRoot,
	}
	AbsAncientAnatolian = &AbstructCulture{
		Name: "abs_ancient_anatolian",
		Root: CaucasianRoot,
	}
	AbsAncientAramaic = &AbstructCulture{
		Name: "abs_ancient_aramaic",
		Root: SemiticRoot,
	}
	AbsAncientBactrian = &AbstructCulture{
		Name: "abs_ancient_bactrian",
		Root: PersianRoot,
	}
	AbsAncientCaucasian = &AbstructCulture{
		Name: "abs_ancient_caucasian",
		Root: CaucasianRoot,
	}
	AbsAncientIranian = &AbstructCulture{
		Name: "abs_ancient_iranian",
		Root: PersianRoot,
	}
	AbsAncientScythian = &AbstructCulture{
		Name: "abs_ancient_scythian",
		Root: EuropeContinentalRoot,
	}
	AbsAncientAryan = &AbstructCulture{
		Name: "abs_ancient_aryan",
		Root: IndianRoot,
	}
	AbsAncientDravidian = &AbstructCulture{
		Name: "abs_ancient_dravidian",
		Root: IndianRoot,
	}
	AbsAncientPracyan = &AbstructCulture{
		Name: "abs_ancient_pracyan",
		Root: IndianRoot,
	}
	AbsAncientTibetan = &AbstructCulture{
		Name: "abs_ancient_tibetan",
		Root: IndianRoot,
	}
)

var AllAncientAbsCultures = []*AbstructCulture{
	AbsAncientBelgae,
	AbsAncientCeltIberian,
	AbsAncientGaelic,
	AbsAncientGallic,
	AbsAncientGermanic,
	AbsAncientIberian,
	AbsAncientOccidental,
	AbsAncientPannonian,
	AbsAncientPretani,
	AbsAncientVeneti,
	AbsAncientItalic,
	AbsAncientDacian,
	AbsAncientHellenistic,
	AbsAncientIllyrian,
	AbsAncientNumidian,
	AbsAncientLevantine,
	AbsAncientAksumite,
	AbsAncientArabian,
	AbsAncientEgyptian,
	AbsAncientNubian,
	AbsAncientAnatolian,
	AbsAncientAramaic,
	AbsAncientBactrian,
	AbsAncientCaucasian,
	AbsAncientIranian,
	AbsAncientScythian,
	AbsAncientAryan,
	AbsAncientDravidian,
	AbsAncientPracyan,
	AbsAncientTibetan,
}

// Medieval
var (
	AbsMedievalAkan = &AbstructCulture{
		Name: "abs_medieval_akan",
		Root: AfricanRoot,
	}
	AbsMedievalArabic = &AbstructCulture{
		Name: "abs_medieval_arabic",
		Root: SemiticRoot,
	}
	AbsMedievalBaltic = &AbstructCulture{
		Name: "abs_medieval_baltic",
		Root: EuropeContinentalRoot,
	}
	AbsMedievalBaltoFinnic = &AbstructCulture{
		Name: "abs_medieval_balto_finnic",
		Root: EuropeContinentalRoot,
	}
	AbsMedievalBerber = &AbstructCulture{
		Name: "abs_medieval_berber",
		Root: BerberRoot,
	}
	AbsMedievalBrythonic = &AbstructCulture{
		Name: "abs_medieval_brythonic",
		Root: EuropeContinentalRoot,
	}
	AbsMedievalBurman = &AbstructCulture{
		Name: "abs_medieval_burman",
		Root: IndianRoot,
	}
	AbsMedievalByzantine = &AbstructCulture{
		Name: "abs_medieval_byzantine",
		Root: MediterraneanRoot,
	}
	AbsMedievalCentralAfrican = &AbstructCulture{
		Name: "abs_medieval_central_african",
		Root: AfricanRoot,
	}
	AbsMedievalCentralGermanic = &AbstructCulture{
		Name: "abs_medieval_central_germanic",
		Root: EuropeContinentalRoot,
	}
	AbsMedievalChinese = &AbstructCulture{
		Name: "abs_medieval_chinese",
		Root: ChineseRoot,
	}
	AbsMedievalDravidian = &AbstructCulture{
		Name: "abs_medieval_dravidian",
		Root: IndianRoot,
	}
	AbsMedievalEastAfrican = &AbstructCulture{
		Name: "abs_medieval_east_african",
		Root: AfricanRoot,
	}
	AbsMedievalEastSlavic = &AbstructCulture{
		Name: "abs_medieval_east_slavic",
		Root: EuropeContinentalRoot,
	}
	AbsMedievalFrankish = &AbstructCulture{
		Name: "abs_medieval_frankish",
		Root: EuropeContinentalRoot,
	}
	AbsMedievalGoidelic = &AbstructCulture{
		Name: "abs_medieval_goidelic",
		Root: EuropeContinentalRoot,
	}
	AbsMedievalGuineanUplander = &AbstructCulture{
		Name: "abs_medieval_guinean_uplander",
		Root: AfricanRoot,
	}
	AbsMedievalHornAfrican = &AbstructCulture{
		Name: "abs_medieval_horn_african",
		Root: AfricanRoot,
	}
	AbsMedievalIberian = &AbstructCulture{
		Name: "abs_medieval_iberian",
		Root: MediterraneanRoot,
	}
	AbsMedievalIndoAryan = &AbstructCulture{
		Name: "abs_medieval_indo_aryan",
		Root: IndianRoot,
	}
	AbsMedievalIranian = &AbstructCulture{
		Name: "abs_medieval_iranian",
		Root: PersianRoot,
	}
	AbsMedievalIsraelite = &AbstructCulture{
		Name: "abs_medieval_israelite",
		Root: SemiticRoot,
	}
	AbsMedievalLatin = &AbstructCulture{
		Name: "abs_medieval_latin",
		Root: MediterraneanRoot,
	}
	AbsMedievalMagyar = &AbstructCulture{
		Name: "abs_medieval_magyar",
		Root: EuropeContinentalRoot,
	}
	AbsMedievalMongolic = &AbstructCulture{
		Name: "abs_medieval_mongolic",
		Root: MongolianRoot,
	}
	AbsMedievalNigerDelta = &AbstructCulture{
		Name: "abs_medieval_niger_delta",
		Root: AfricanRoot,
	}
	AbsMedievalNorthGermanic = &AbstructCulture{
		Name: "abs_medieval_north_germanic",
		Root: NordicEuropeRoot,
	}
	AbsMedievalQiangic = &AbstructCulture{
		Name: "abs_medieval_qiangic",
		Root: MongolianRoot,
	}
	AbsMedievalSahelian = &AbstructCulture{
		Name: "abs_medieval_sahelian",
		Root: AfricanRoot,
	}
	AbsMedievalSenegambian = &AbstructCulture{
		Name: "abs_medieval_senegambian",
		Root: AfricanRoot,
	}
	AbsMedievalSouthSlavic = &AbstructCulture{
		Name: "abs_medieval_south_slavic",
		Root: EuropeContinentalRoot,
	}
	AbsMedievalTibetan = &AbstructCulture{
		Name: "abs_medieval_tibetan",
		Root: IndianRoot,
	}
	AbsMedievalTocharian = &AbstructCulture{
		Name: "abs_medieval_tocharian",
		Root: PersianRoot,
	}
	AbsMedievalTurkic = &AbstructCulture{
		Name: "abs_medieval_turkic",
		Root: TurkicRoot,
	}
	AbsMedievalUgroPermian = &AbstructCulture{
		Name: "abs_medieval_ugro_permian",
		Root: EuropeContinentalRoot,
	}
	AbsMedievalVlach = &AbstructCulture{
		Name: "abs_medieval_vlach",
		Root: EuropeContinentalRoot,
	}
	AbsMedievalVolgaFinnic = &AbstructCulture{
		Name: "abs_medieval_volga_finnic",
		Root: EuropeContinentalRoot,
	}
	AbsMedievalWestGermanic = &AbstructCulture{
		Name: "abs_medieval_west_germanic",
		Root: EuropeContinentalRoot,
	}
	AbsMedievalWestSlavic = &AbstructCulture{
		Name: "abs_medieval_west_slavic",
		Root: EuropeContinentalRoot,
	}
)

var AllMedievalAbsCultures = []*AbstructCulture{
	AbsMedievalAkan,
	AbsMedievalArabic,
	AbsMedievalBaltic,
	AbsMedievalBaltoFinnic,
	AbsMedievalBerber,
	AbsMedievalBrythonic,
	AbsMedievalBurman,
	AbsMedievalByzantine,
	AbsMedievalCentralAfrican,
	AbsMedievalCentralGermanic,
	AbsMedievalChinese,
	AbsMedievalDravidian,
	AbsMedievalEastAfrican,
	AbsMedievalEastSlavic,
	AbsMedievalFrankish,
	AbsMedievalGoidelic,
	AbsMedievalGuineanUplander,
	AbsMedievalHornAfrican,
	AbsMedievalIberian,
	AbsMedievalIndoAryan,
	AbsMedievalIranian,
	AbsMedievalIsraelite,
	AbsMedievalLatin,
	AbsMedievalMagyar,
	AbsMedievalMongolic,
	AbsMedievalNigerDelta,
	AbsMedievalNorthGermanic,
	AbsMedievalQiangic,
	AbsMedievalSahelian,
	AbsMedievalSenegambian,
	AbsMedievalSouthSlavic,
	AbsMedievalTibetan,
	AbsMedievalTocharian,
	AbsMedievalTurkic,
	AbsMedievalUgroPermian,
	AbsMedievalVlach,
	AbsMedievalVolgaFinnic,
	AbsMedievalWestGermanic,
	AbsMedievalWestSlavic,
}

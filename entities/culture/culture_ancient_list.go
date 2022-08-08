package culture

import (
	g "persons_generator/entities/gender"
	"persons_generator/entities/influence"
	"persons_generator/entities/language"
	"persons_generator/tools"
)

// Antique
var AllAntiqueCultures = tools.Merge(
	AllAncientBelgaeCultures,
	AllAncientCeltIberianCultures,
	AllAncientGaelicCultures,
	AllAncientGallicCultures,
	AllAncientGermanicCultures,
	AllAncientIberianCultures,
	AllAncientOccidentalCultures,
	AllAncientPannonianCultures,
	AllAncientPretaniCultures,
	AllAncientVenetiCultures,
	AllAncientItalicCultures,
	AllAncientDacianCultures,
	AllAncientHellenisticCultures,
	AllAncientIllyrianCultures,
	AllAncientNumidianCultures,
	AllAncientLevantineCultures,
	AllAncientAksumiteCultures,
	AllAncientArabianCultures,
	AllAncientEgyptianCultures,
	AllAncientNubianCultures,
	AllAncientAnatolianCultures,
	AllAncientAramaicCultures,
	AllAncientBactrianCultures,
	AllAncientCaucasianCultures,
	AllAncientIranianCultures,
	AllAncientScythianCultures,
	AllAncientAryanCultures,
	AllAncientDravidianCultures,
	AllAncientPracyanCultures,
	AllAncientTibetanCultures,
)

// AncientBelgae
var (
// -- celtic
// Aduatacian
// Bellovacian
// Eburonian
// Nervian
// Sennonian
// Treverian
// Menapian
// Morinian
// Veliocassian
)
var AllAncientBelgaeCultures = []*Culture{}

// AncientCeltIberian
var (
// -- celtic
// Asturian
// Callaecian
// Caristian
// Carpetanian
// Celtiberian
// Celtician
// Lobetanian
// Lusitanian
// Sedetanian
// Vaccaeian
// Vardulian
// Vettonian
)
var AllAncientCeltIberianCultures = []*Culture{}

// AncientGaelic
var (
// -- celtic
// Hibernian
// Manavian
)
var AllAncientGaelicCultures = []*Culture{}

// AncientGallic
var (
// -- celtic
// Aremorican
// Aulercian
// Arverni
// Biturigan
// Carnuti
// Haedui
// Helvetian
// Lemovician
// Pictonian
// Santonian
// Vindelician
// Volcae
// Lepontic
// Salluvian
)
var AllAncientGallicCultures = []*Culture{}

// AncientGermanic
var (
// -- germanic
// Anglian
// Bastarnae
// Cimbrian
// Frisian
// Gothonic
// Gutoni
// Helveconian
// Herulian
// Ingvaeonic
// Irminonic
// Istvaonic
// Lemovian
// Raumarician
// Reudingian
// Rugian
// Saxonian
// Suebian
// Teutonian
// Vandal
)
var AllAncientGermanicCultures = []*Culture{}

// AncientIberian
var (
// -- iberian
// Ausetanian
// Bastetanian
// Contestanian
// Cynetian
// Edetanian
// Ilercavonian
// Indiketian
// Lacetanian
// Oretanian
// Turdetanian
// Turdulian
)
var AllAncientIberianCultures = []*Culture{}

// AncientOccidental
var (
// -- iberian
// Aquitani
// Autrigonian
// Cantabrian
// Ilergetian
// Vasconian
// Corsian
// Nuragic
// Talaiotic
)
var AllAncientOccidentalCultures = []*Culture{}

// AncientPannonian
var (
// -- celtic
// Boian
// Carnian
// Eraviscian
// Letobician
// Scordiscian
// Noric
)
var AllAncientPannonianCultures = []*Culture{}

// AncientPretani
var (
// -- britannic
// Brigantic
// Caledonian
// Cantian
// Coritani
// Cornovian
// Damnonian
// Demetian
// Dobunnian
// Dumnonian
// Durotrigan
// Iceni
// Ordovitian
// Pretani
// Taexalian
// Trinovantian
// Votadinian
)
var AllAncientPretaniCultures = []*Culture{}

// AncientVeneti
var (
// -- britannic
// Aestian
)
var AllAncientVenetiCultures = []*Culture{}

// AncientItalic
var (
	// -- roman
	// Bruttian
	// Ligurian
	// Lucanian
	// Siculian
	// Umbrian
	// Venetic
	AncientEtruscan = &Culture{
		Name:            "ancient_etruscan",
		Abstuct:         AbsAncientItalic,
		Root:            MediterraneanRoot,
		Language:        language.Etruscan,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions: []*Tradition{
			MysticalAncestorsTradition,
			PracticedPiratesTradition,
			CityKeepersTradition,
			LegalisticTradition,
			ParochialismTradition,
		},
	}
	// Rhaetian
	AncientRoman = &Culture{
		Name:            "ancient_roman",
		Abstuct:         AbsAncientItalic,
		Root:            MediterraneanRoot,
		Language:        language.Latin,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions: []*Tradition{
			FormationFightingExpertsTradition,
			HereditaryHierarchyTradition,
			LegalisticTradition,
			MysticalAncestorsTradition,
			RefinedPoetryTradition,
		},
	}
	AncientSabellian = &Culture{
		Name:            "ancient_sabellian",
		Abstuct:         AbsAncientItalic,
		Root:            MediterraneanRoot,
		Language:        language.Oscan,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HillDwellersTradition, MetalworkersTradition},
	}
)

var AllAncientItalicCultures = []*Culture{
	AncientEtruscan,
	AncientRoman,
	AncientSabellian,
}

// AncientDacian
var (
	// -- levantine
	AncientDacian = &Culture{
		Name:            "ancient_dacian",
		Abstuct:         AbsAncientDacian,
		Root:            EuropeContinentalRoot,
		Language:        language.Dacian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MetalworkersTradition, TribalUnityTradition, StalwartDefendersTradition},
	}
)
var AllAncientDacianCultures = []*Culture{AncientDacian}

// AncientHellenistic
var (
	// -- greek
	AncientAchaean = &Culture{
		Name:            "ancient_achaean",
		Abstuct:         AbsAncientHellenistic,
		Root:            MediterraneanRoot,
		Language:        language.Greek,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CityKeepersTradition, ParochialismTradition, RefinedPoetryTradition},
	}
	// Aegean
	// Aeolian
	// Aetolian
	// Arcadian
	AncientAthenian = &Culture{
		Name:            "ancient_athenian",
		Abstuct:         AbsAncientHellenistic,
		Root:            MediterraneanRoot,
		Language:        language.Greek,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			PhilosopherCultureTradition,
			FerventTempleBuildersTradition,
			ParochialismTradition,
			FormationFightingExpertsTradition,
			StalwartDefendersTradition,
		},
	}
	// Boeotian
	AncientCypriot = &Culture{
		Name:            "ancient_cypriot",
		Abstuct:         AbsAncientHellenistic,
		Root:            MediterraneanRoot,
		Language:        language.Greek,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CultureBlendingTradition, MalleableSubjectsTradition},
	}
	// Euboean
	AncientHellenistic = &Culture{
		Name:            "ancient_hellenistic",
		Abstuct:         AbsAncientHellenistic,
		Root:            MediterraneanRoot,
		Language:        language.Greek,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{PhilosopherCultureTradition, ParochialismTradition, CultureBlendingTradition, FormationFightingExpertsTradition},
	}
	AncientIonian = &Culture{
		Name:            "ancient_ionian",
		Abstuct:         AbsAncientHellenistic,
		Root:            MediterraneanRoot,
		Language:        language.Greek,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			CityKeepersTradition,
			SeafarersTradition,
			ParochialismTradition,
			XenophilicTradition,
		},
	}
	// Massalian
	// Pontic
	// Propontic
	AncientTroan = &Culture{
		Name:            "ancient_troan",
		Abstuct:         AbsAncientHellenistic,
		Root:            MediterraneanRoot,
		Language:        language.Greek,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			ParochialismTradition,
			HereditaryHierarchyTradition,
			StalwartDefendersTradition,
		},
	}
	// Argolian
	AncientLacedaemonian = &Culture{
		Name:            "ancient_lacedaemonian",
		Abstuct:         AbsAncientHellenistic,
		Root:            MediterraneanRoot,
		Language:        language.Greek,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			SwordsForHireTradition,
			WarriorCultureTradition,
			FormationFightingExpertsTradition,
			StalwartDefendersTradition,
			SpartanTradition,
		},
	}
	// Bosporan
	// Cretan
	// Cyrenaican
	// Epirote
	// Italiotian
	// Siceliote
	AncientMacedonian = &Culture{
		Name:            "ancient_macedonian",
		Abstuct:         AbsAncientHellenistic,
		Root:            MediterraneanRoot,
		Language:        language.Greek,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HillDwellersTradition, FormationFightingExpertsTradition},
	}
	// Thessalian
)

var AllAncientHellenisticCultures = []*Culture{
	AncientAchaean,
	AncientAthenian,
	AncientCypriot,
	AncientHellenistic,
	AncientIonian,
	AncientTroan,
	AncientLacedaemonian,
	AncientMacedonian,
}

// AncientIllyrian
var (
	// -- illyrian
	AncientIllyrian = &Culture{
		Name:            "ancient_illyrian",
		Abstuct:         AbsAncientIllyrian,
		Root:            MediterraneanRoot,
		Language:        language.Illyrian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{PracticedPiratesTradition, CoastalWarriorsTradition, DexterousFishermenTradition},
	}
)
var AllAncientIllyrianCultures = []*Culture{AncientIllyrian}

// AncientNumidian
var (
	// -- north_african
	AncientBerber = &Culture{
		Name:            "ancient_berber",
		Abstuct:         AbsAncientNumidian,
		Root:            AfricanRoot,
		Language:        language.Zenaga,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DrylandDwellersTradition, TribalUnityTradition, CaravaneersTradition, AfricanToleranceTradition},
	}
	// Gaetulian
	// Garamantic
	AncientLibyan = &Culture{
		Name:            "ancient_libyan",
		Abstuct:         AbsAncientNumidian,
		Root:            AfricanRoot,
		Language:        language.Numidian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DrylandDwellersTradition, WarriorsOfTheDryTradition, AfricanToleranceTradition},
	}
	// Massaesylian
	// Massylian
	AncientMaurian = &Culture{
		Name:            "ancient_maurian",
		Abstuct:         AbsAncientNumidian,
		Root:            AfricanRoot,
		Language:        language.Tuareg,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DrylandDwellersTradition, AfricanToleranceTradition},
	}
	AncientNumidian = &Culture{
		Name:            "ancient_numidian",
		Abstuct:         AbsAncientNumidian,
		Root:            AfricanRoot,
		Language:        language.Numidian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			DrylandDwellersTradition,
			TribalUnityTradition,
			StalwartDefendersTradition,
			WarriorsOfTheDryTradition,
			AfricanToleranceTradition,
			ChivalryTradition,
		},
	}
)
var AllAncientNumidianCultures = []*Culture{AncientBerber, AncientLibyan, AncientMaurian, AncientNumidian}

// AncientLevantine
var (
	AncientPunic = &Culture{
		Name:            "ancient_punic",
		Abstuct:         AbsAncientLevantine,
		Proto:           []*Culture{AncientPhoenician},
		Root:            SemiticRoot,
		Language:        language.Punic,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			MaritimeMercantilismTradition,
			RulingCasteTradition,
			XenophilicTradition,
			StrongBelieversTradition,
		},
	}
	AncientHebrew = &Culture{
		Name:            "ancient_hebrew",
		Abstuct:         AbsAncientLevantine,
		Root:            SemiticRoot,
		Language:        language.Hebrew,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			DefensiveTacticsTradition,
			PhilosopherCultureTradition,
			BoundByFaithTradition,
			FerventTempleBuildersTradition,
			ForebearingTradition,
			XenophilicTradition,
		},
	}
	AncientPhoenician = &Culture{
		Name:            "ancient_phoenician",
		Abstuct:         AbsAncientLevantine,
		Root:            SemiticRoot,
		Language:        language.Phoenician,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			SeafarersTradition,
			MaritimeMercantilismTradition,
			CityKeepersTradition,
			ParochialismTradition,
			CultureBlendingTradition,
			ConcubinesTradition,
		},
	}
	AncientNabatean = &Culture{
		Name:            "ancient_nabatean",
		Abstuct:         AbsAncientLevantine,
		Root:            SemiticRoot,
		Language:        language.Nabataean,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			DesertRibatsTradition,
			PastorialistsTradition,
			EsteemedHospitalityTradition,
			CaravaneersTradition,
			XenophilicTradition,
			PolygamousTradition,
		},
	}
)
var AllAncientLevantineCultures = []*Culture{AncientPunic, AncientHebrew, AncientPhoenician, AncientNabatean}

// AncientAksumite
var (
// -- ethiopian
// Aksumite
// Macrobian
)
var AllAncientAksumiteCultures = []*Culture{}

// AncientArabian
var (
// -- arabian
// Hadhrami
// Himjar
// Minaean
// Qatabanian
// Sabaean
// Makan
// Qedarite
// Ta'if
// Thamudi
)
var AllAncientArabianCultures = []*Culture{}

// AncientEgyptian
var (
	// -- egyptian
	AncientEgyptian = &Culture{
		Name:            "ancient_egyptian",
		Abstuct:         AbsAncientEgyptian,
		Root:            EgyptianRoot,
		Language:        language.Egyptian,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			AgrarianTradition,
			StaunchTraditionalistsTradition,
			RulingCasteTradition,
			FerventTempleBuildersTradition,
			HereditaryHierarchyTradition,
			StrongBelieversTradition,
			ConcubinesTradition,
		},
	}
)
var AllAncientEgyptianCultures = []*Culture{AncientEgyptian}

// AncientNubian
var (
// -- egyptian
// Blemmyan
// Meroitic
)
var AllAncientNubianCultures = []*Culture{}

// AncientAnatolian
var (
	// -- anatolian
	AncientArmenian = &Culture{
		Name:            "ancient_armenian",
		Abstuct:         AbsAncientAnatolian,
		Root:            CaucasianRoot,
		Language:        language.Armenian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			SacredMountainsTradition,
			HighlandWarriorsTradition,
			StalwartDefendersTradition,
			ExpertArtisansTradition,
		},
	}

// Cabalian
// Carian
// Cataonian
// Cennataean
// Isaurian
// Lalasian
// Lycaonian
// Lycian
// Lydian
// Milyadian
// Morimenian
// Mysian
// Oroandian
// Pamphylian
// Paphlagonian
// Phrygian
// Pisidian
// Cappadocian
// Cilician
)
var AllAncientAnatolianCultures = []*Culture{AncientArmenian}

// AncientAramaic
var (
	// -- levantine
	AncientAramaic = &Culture{
		Name:            "ancient_aramaic",
		Abstuct:         AbsAncientAramaic,
		Root:            MesopotamianRoot,
		Language:        language.Aramaic,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CultureBlendingTradition, StrongBelieversTradition, ConcubinesTradition},
	}
	AncientAssyrian = &Culture{
		Name:            "ancient_assyrian",
		Abstuct:         AbsAncientAramaic,
		Root:            MesopotamianRoot,
		Language:        language.Akkadian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{WarriorCultureTradition, BattlefieldLootersTradition, QuarrelsomeTradition, ConcubinesTradition},
	}
	AncientBabylonian = &Culture{
		Name:            "ancient_babylonian",
		Abstuct:         AbsAncientAramaic,
		Root:            MesopotamianRoot,
		Language:        language.Akkadian,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			FerventTempleBuildersTradition,
			PhilosopherCultureTradition,
			ParochialismTradition,
			RecognitionOfTalentTradition,
			MalleableSubjectsTradition,
			CharismaticTradition,
			ConcubinesTradition,
		},
	}
)
var AllAncientAramaicCultures = []*Culture{AncientAramaic, AncientAssyrian, AncientBabylonian}

// AncientBactrian
var (
// -- bactrian
// Bactrian
// Margian
// Sogdian
// Khotanese
// Phryni
// Shule
// Tayuan
// Tocharian
// Wusun
// Yuezhi
)
var AllAncientBactrianCultures = []*Culture{}

// AncientCaucasian
var (
// -- persian
// Albanian
// Colchian
// Ibero
)
var AllAncientCaucasianCultures = []*Culture{}

// AncientIranian
var (
	// -- persian
	// Agartian
	// Amardian
	// Cossian
	// Elamite
	// Gedrosian
	// Hyrcanian
	// Median
	// Pactyan
	// Parecanian
	// Sarangian
	// Sattagydian
	// Utian
	// Uxian
	// Cadusian
	// Carmanian
	// Parthian
	AncientPersian = &Culture{
		Name:            "ancient_persian",
		Abstuct:         AbsAncientIranian,
		Root:            PersianRoot,
		Language:        language.Farsi,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			DrylandDwellersTradition,
			ImmortalsTradition,
			GardenArchitectsTradition,
			PhilosopherCultureTradition,
			RefinedPoetryTradition,
		},
	}
)
var AllAncientIranianCultures = []*Culture{AncientPersian}

// AncientScythian
var (
	// -- scythian
	// Agathyrsian
	// Kharesmian
	// Legian
	// Maeotian
	// Sakan
	// Sindi
	// Siracian
	// Thyssagetian
	// Dahae
	AncientSarmatian = &Culture{
		Name:            "ancient_sarmatian",
		Abstuct:         AbsAncientScythian,
		Root:            PersianRoot,
		Language:        language.Scythian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			HorseLordsTradition,
			SteppeToleranceTradition,
			HitAndRunTacticiansTradition,
			ChivalryTradition,
			ConcubinesTradition,
		},
	}
	AncientScythian = &Culture{
		Name:            "ancient_scythian",
		Abstuct:         AbsAncientScythian,
		Root:            PersianRoot,
		Language:        language.Scythian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			HorseLordsTradition,
			SteppeToleranceTradition,
			TribalUnityTradition,
			BattlefieldLootersTradition,
			QuarrelsomeTradition,
			ChivalryTradition,
			ConcubinesTradition,
		},
	}
)
var AllAncientScythianCultures = []*Culture{AncientSarmatian, AncientScythian}

// AncientAryan
var (
// -- mauryan
// Avanti
// Lankan
// Saurashtran
// Shauraseni
// Vidharban
// Dardic
// Maharashtran
// Gandhari
// Sindhi
)
var AllAncientAryanCultures = []*Culture{}

// AncientDravidian
var (
// -- dravidian
// Kannadan
// Tamil
// Telugu
)
var AllAncientDravidianCultures = []*Culture{}

// AncientPracyan
var (
// -- mauryan
// Atavi
// Kalingan
// Kamarupi
// Bangli
// Magadhi
// Nepala
)
var AllAncientPracyanCultures = []*Culture{}

// AncientTibetan
var (
// -- tibetan
// Sumpa
// Tsang
// Yarlung
// Zhangzhung
)
var AllAncientTibetanCultures = []*Culture{}

package culture

import (
	g "persons_generator/entities/gender"
	"persons_generator/entities/influence"
	"persons_generator/entities/language"
	"persons_generator/tools"
)

var AllCultures = tools.Merge(
	AllAkanCultures,
	AllArabicCultures,
	AllBalticCultures,
	AllBaltoFinnicCultures,
	AllBerberCultures,
	AllBrythonicCultures,
	AllBurmanCultures,
	AllByzantineCultures,
	AllCentralAfricanCultures,
	AllCentralGermanicCultures,
	AllChineseCultures,
	AllDravidianCultures,
	AllEastAfricanCultures,
	AllEastSlavicCultures,
	AllFrankishCultures,
	AllGoidelicCultures,
	AllGuineanUplanderCultures,
	AllHornAfricanCultures,
	AllIberianCultures,
	AllIndoAryanCultures,
	AllIranianCultures,
	AllIsraeliteCultures,
	AllLatinCultures,
	AllMagyarCultures,
	AllMongolicCultures,
	AllNigerDeltaCultures,
	AllNorthGermanicCultures,
	AllQiangicCultures,
	AllSahelianCultures,
	AllSenegambianCultures,
	AllSouthSlavicCultures,
	AllTibetanCultures,
	AllTocharianCultures,
	AllTurkicCultures,
	AllUgroPermianCultures,
	AllVlachCultures,
	AllVolgaFinnicCultures,
	AllWestGermanicCultures,
	AllWestSlavicCultures,
)

// Akan
var (
	AkanCulture = &Culture{
		Name:            "akan_akan",
		Root:            AfricanRoot,
		Language:        language.Kwa,
		Ethos:           Egalitarian,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.FemaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{BushHuntingTradition, MatriarchalTradition, MysticalAncestorsTradition, ParochialismTradition},
	}
	GuanCulture = &Culture{
		Name:            "akan_guan",
		Root:            AfricanRoot,
		Language:        language.Kwa,
		Ethos:           Communal,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{BushHuntingTradition, DrylandDwellersTradition, FrequentFestivitiesTradition, StrongBelieversTradition},
	}
	KruCulture = &Culture{
		Name:            "akan_kru",
		Root:            AfricanRoot,
		Language:        language.Kru,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{BushHuntingTradition, DexterousFishermenTradition, PracticedPiratesTradition},
	}
)

var AllAkanCultures = []*Culture{AkanCulture, GuanCulture, KruCulture}

// Arabic
var (
	ArabicBedouinCulture = &Culture{
		Name:            "arabic_bedouin",
		Root:            SemiticRoot,
		Language:        language.Arabic,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, DesertTravelersTradition, MubarizunTradition, TribalUnityTradition},
	}
	ArabicEgyptianCulture = &Culture{
		Name:            "arabic_egyptian",
		Root:            SemiticRoot,
		Language:        language.Arabic,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{AgrarianTradition, MubarizunTradition, PhilosopherCultureTradition, StrongBelieversTradition},
	}
	ArabicMaghrebiCulture = &Culture{
		Proto:           []*Culture{ArabicBedouinCulture, BaranisCulture},
		Name:            "arabic_maghrebi",
		Root:            SemiticRoot,
		Language:        language.Arabic,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{AfricanToleranceTradition, DexterousFishermenTradition, DrylandDwellersTradition, MubarizunTradition, XenophilicTradition},
	}
	ArabicMashriqiCulture = &Culture{
		Proto:           []*Culture{ArabicBedouinCulture},
		Name:            "arabic_mashriqi",
		Root:            SemiticRoot,
		Language:        language.Arabic,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{DrylandDwellersTradition, MedicinalHerbalistsTradition, MubarizunTradition, PhilosopherCultureTradition},
	}
	ArabicYemeniCulture = &Culture{
		Name:            "arabic_yemeni",
		Root:            SemiticRoot,
		Language:        language.Arabic,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{MaritimeMercantilismTradition, MountainHerdingTradition, MountaineersTradition, ReligionBlendingTradition},
	}
)

var AllArabicCultures = []*Culture{ArabicBedouinCulture, ArabicEgyptianCulture, ArabicMaghrebiCulture, ArabicMashriqiCulture, ArabicYemeniCulture}

// Baltic

var (
	LatgalianCulture = &Culture{
		Name:            "baltic_latgalian",
		Root:            EuropeContinentalRoot,
		Language:        language.Latgalian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{ForestWardensTradition, SacredGrovesTradition, StaunchTraditionalistsTradition},
	}
	LithuanianCulture = &Culture{
		Name:            "baltic_lithuanian",
		Root:            EuropeContinentalRoot,
		Language:        language.Lithuanian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{ForestWardensTradition, HitAndRunTacticiansTradition, SacredGrovesTradition, StrongBelieversTradition},
	}
	PrussianCulture = &Culture{
		Name:            "baltic_prussian",
		Root:            EuropeContinentalRoot,
		Language:        language.Lithuanian,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CharismaticTradition, EquitableTradition, ForestWardensTradition, SacredGrovesTradition},
	}
)

var AllBalticCultures = []*Culture{LatgalianCulture, LithuanianCulture, PrussianCulture}

// BaltoFinnic

var (
	EstonianCulture = &Culture{
		Name:            "balto_finnic_estonian",
		Root:            EuropeContinentalRoot,
		Language:        language.Estonian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MusicalTheoristsTradition, CoastalWarriorsTradition, HirdsTradition, StaunchTraditionalistsTradition},
	}
	FinnishCulture = &Culture{
		Name:            "balto_finnic_finnish",
		Root:            EuropeContinentalRoot,
		Language:        language.Finnish,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DexterousFishermenTradition, ForestWardensTradition, IsolationistTradition},
	}
	KarelianCulture = &Culture{
		Name:            "balto_finnic_karelian",
		Root:            EuropeContinentalRoot,
		Language:        language.Karelian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ForestWardensTradition, StalwartDefendersTradition, SwordsForHireTradition},
	}
	SamiCulture = &Culture{
		Name:            "balto_finnic_sami",
		Root:            EuropeContinentalRoot,
		Language:        language.Sami,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AdaptiveSkirmishersTradition, ForestWardensTradition, WinterWarriorsTradition},
	}
	VepsianCulture = &Culture{
		Name:            "balto_finnic_vepsian",
		Root:            EuropeContinentalRoot,
		Language:        language.Veps,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{IsolationistTradition, XenophilicTradition},
	}
)

var AllBaltoFinnicCultures = []*Culture{EstonianCulture, FinnishCulture, KarelianCulture, SamiCulture, VepsianCulture}

// Berber

var (
	BaranisCulture = &Culture{
		Name:            "berber_baranis",
		Root:            BerberRoot,
		Language:        language.Tuareg,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{AfricanToleranceTradition, FamilyBusinessTradition, MountainHerdingTradition},
	}
	ButrCulture = &Culture{
		Name:            "berber_butr",
		Root:            BerberRoot,
		Language:        language.Tuareg,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{AfricanToleranceTradition, DesertRibatsTradition, EqualInheritanceTradition, SaharanNomadsTradition},
	}
	GuancheCulture = &Culture{
		Name:            "berber_guanche",
		Root:            BerberRoot,
		Language:        language.Tuareg,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{StalwartDefendersTradition, MysticalAncestorsTradition, IsolationistTradition},
	}
	ZaghawaCulture = &Culture{
		Name:            "berber_zaghawa",
		Root:            BerberRoot,
		Language:        language.Tuareg,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{AfricanToleranceTradition, DesertRibatsTradition, SaharanNomadsTradition},
	}
)

var AllBerberCultures = []*Culture{BaranisCulture, ButrCulture, GuancheCulture, ZaghawaCulture}

// Brythonic

var AllBrythonicCultures = []*Culture{}

// Burman

var AllBurmanCultures = []*Culture{}

// Byzantine

var AllByzantineCultures = []*Culture{}

// CentralAfrican

var AllCentralAfricanCultures = []*Culture{}

// CentralGermanic

var AllCentralGermanicCultures = []*Culture{}

// Chinese

var AllChineseCultures = []*Culture{}

// Dravidian

var AllDravidianCultures = []*Culture{}

// EastAfrican

var AllEastAfricanCultures = []*Culture{}

// EastSlavic

var AllEastSlavicCultures = []*Culture{}

// Frankish

var AllFrankishCultures = []*Culture{}

// Goidelic

var AllGoidelicCultures = []*Culture{}

// GuineanUplander

var AllGuineanUplanderCultures = []*Culture{}

// HornAfrican

var AllHornAfricanCultures = []*Culture{}

// Iberian

var AllIberianCultures = []*Culture{}

// IndoAryan

var AllIndoAryanCultures = []*Culture{}

// Iranian

var AllIranianCultures = []*Culture{}

// Israelite

var AllIsraeliteCultures = []*Culture{}

// Latin

var AllLatinCultures = []*Culture{}

// Magyar

var AllMagyarCultures = []*Culture{}

// Mongolic

var AllMongolicCultures = []*Culture{}

// NigerDelta

var AllNigerDeltaCultures = []*Culture{}

// NorthGermanic

var AllNorthGermanicCultures = []*Culture{}

// Qiangic

var AllQiangicCultures = []*Culture{}

// Sahelian

var AllSahelianCultures = []*Culture{}

// Senegambian

var AllSenegambianCultures = []*Culture{}

// SouthSlavic

var AllSouthSlavicCultures = []*Culture{}

// Tibetan

var AllTibetanCultures = []*Culture{}

// Tocharian

var AllTocharianCultures = []*Culture{}

// Turkic

var AllTurkicCultures = []*Culture{}

// UgroPermian

var AllUgroPermianCultures = []*Culture{}

// Vlach

var AllVlachCultures = []*Culture{}

// VolgaFinnic

var AllVolgaFinnicCultures = []*Culture{}

// WestGermanic

var AllWestGermanicCultures = []*Culture{}

// WestSlavic

var AllWestSlavicCultures = []*Culture{}

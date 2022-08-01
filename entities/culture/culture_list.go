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

var (
	BretonCulture = &Culture{
		Name:            "brythonic_breton",
		Root:            EuropeContinentalRoot,
		Language:        language.Breton,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ChansonDeGesteTradition, DexterousFishermenTradition, StorytellersTradition, SwordsForHireTradition},
	}
	CornishCulture = &Culture{
		Name:            "brythonic_cornish",
		Root:            EuropeContinentalRoot,
		Language:        language.Cornish,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AncientMinersTradition, DexterousFishermenTradition, StorytellersTradition},
	}
	CumbrianCulture = &Culture{
		Name:            "brythonic_cumbrian",
		Root:            EuropeContinentalRoot,
		Language:        language.Cumbric,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ConcubinesTradition, HighlandWarriorsTradition, RefinedPoetryTradition, StalwartDefendersTradition},
	}
	PictishCulture = &Culture{
		Name:            "brythonic_pictish",
		Root:            EuropeContinentalRoot,
		Language:        language.Pictish,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ConcubinesTradition, HighlandWarriorsTradition, HillDwellersTradition, RefinedPoetryTradition},
	}
	WelshCulture = &Culture{
		Name:            "brythonic_welsh",
		Root:            EuropeContinentalRoot,
		Language:        language.Welsh,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ConcubinesTradition, HighlandWarriorsTradition, LongbowCompetitionsTradition, RefinedPoetryTradition},
	}
)

var AllBrythonicCultures = []*Culture{BretonCulture, CornishCulture, CumbrianCulture, PictishCulture, WelshCulture}

// Burman

var (
	BurmeseCulture = &Culture{
		Name:            "burman_burmese",
		Root:            IndianRoot,
		Language:        language.Burmese,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{FerventTempleBuildersTradition, FrequentFestivitiesTradition, ReligionBlendingTradition, RoyalArmyTradition},
	}
	MonCulture = &Culture{
		Name:            "burman_mon",
		Root:            IndianRoot,
		Language:        language.Burmese,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{MysticalAncestorsTradition, ReligiousPatronageTradition, CultureBlendingTradition},
	}
)

var AllBurmanCultures = []*Culture{BurmeseCulture, MonCulture}

// Byzantine

var (
	AlanCulture = &Culture{
		Name:            "byzantine_alan",
		Root:            TurkicRoot,
		Language:        language.Scythian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, MaritalCeremoniesTradition, MountaineersTradition, SteppeToleranceTradition},
	}
	ArmenianCulture = &Culture{
		Name:            "byzantine_armenian",
		Root:            CaucasianRoot,
		Language:        language.Armenian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{FerventTempleBuildersTradition, HighlandWarriorsTradition, HillDwellersTradition, StrongBelieversTradition},
	}
	GeorgianCulture = &Culture{
		Name:            "byzantine_georgian",
		Root:            CaucasianRoot,
		Language:        language.Georgian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CastleKeepersTradition, CaucasianWolvesTradition, HereditaryHierarchyTradition, MetalworkersTradition},
	}
	GreekCulture = &Culture{
		Name:            "byzantine_greek",
		Root:            MediterraneanRoot,
		Language:        language.Greek,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{FormationFightingExpertsTradition, LegalisticTradition},
	}
	SyriacCulture = &Culture{
		Name:            "byzantine_syriac",
		Root:            SemiticRoot,
		Language:        language.Aramaic,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, LoyalSubjectsTradition, PhilosopherCultureTradition, StrongBelieversTradition},
	}
)

var AllByzantineCultures = []*Culture{AlanCulture, ArmenianCulture, GeorgianCulture, GreekCulture, SyriacCulture}

// CentralAfrican

var (
	HausaCulture = &Culture{
		Name:            "central_african_hausa",
		Root:            AfricanRoot,
		Language:        language.Hausa,
		Ethos:           Egalitarian,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, FamilyBusinessTradition, ParochialismTradition},
	}
	KanuriCulture = &Culture{
		Name:            "central_african_kanuri",
		Root:            AfricanRoot,
		Language:        language.Tebu,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HitAndRunTacticiansTradition, QuarrelsomeTradition},
	}
	NupeCulture = &Culture{
		Name:            "central_african_nupe",
		Root:            AfricanRoot,
		Language:        language.Kwa,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MusicalTheoristsTradition, ParochialismTradition},
	}
	SaoCulture = &Culture{
		Name:            "central_african_sao",
		Root:            AfricanRoot,
		Language:        language.Hausa,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, CityKeepersTradition, DrylandDwellersTradition, ParochialismTradition},
	}
)

var AllCentralAfricanCultures = []*Culture{HausaCulture, KanuriCulture, NupeCulture, SaoCulture}

// CentralGermanic

var (
	BavarianCulture = &Culture{
		Name:            "germanic_bavarian",
		Root:            EuropeContinentalRoot,
		Language:        language.German,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HereditaryHierarchyTradition, StandAndFightTradition, StrongBelieversTradition},
	}
	DutchCulture = &Culture{
		Name:            "germanic_dutch",
		Proto:           []*Culture{FrisianCulture /* @TODO add frankish */},
		Root:            EuropeContinentalRoot,
		Language:        language.Dutch,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, MaritimeMercantilismTradition, ParochialismTradition, PoldersTradition},
	}
	FranconianCulture = &Culture{
		Name:            "germanic_franconian",
		Proto:           []*Culture{ /* @TODO add frankish */ },
		Root:            EuropeContinentalRoot,
		Language:        language.German,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AstuteDiplomatsTradition, CastleKeepersTradition, HereditaryHierarchyTradition},
	}
	FrisianCulture = &Culture{
		Name:            "germanic_frisian",
		Root:            EuropeContinentalRoot,
		Language:        language.Frisian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{BattlefieldLootersTradition, DexterousFishermenTradition, StrongBelieversTradition},
	}
	GermanCulture = &Culture{
		Name:            "germanic_german",
		Root:            EuropeContinentalRoot,
		Language:        language.German,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AstuteDiplomatsTradition, HereditaryHierarchyTradition, StandAndFightTradition},
	}
	LangobardCulture = &Culture{
		Name:            "germanic_langobard",
		Root:            EuropeContinentalRoot,
		Language:        language.Lombardian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{IsolationistTradition, MartialAdmirationTradition, StandAndFightTradition, MalleableInvadersTradition},
	}
	SaxonCulture = &Culture{
		Name:            "germanic_saxon",
		Proto:           []*Culture{ /* @TODO add old_saxon */ },
		Root:            EuropeContinentalRoot,
		Language:        language.Saxon,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ReligiousPatronageTradition, RulingCasteTradition, StandAndFightTradition},
	}
	SwabianCulture = &Culture{
		Name:            "germanic_swabian",
		Root:            EuropeContinentalRoot,
		Language:        language.German,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{IndustriousTradition, MountainHomesTradition, ParochialismTradition},
	}
)

var AllCentralGermanicCultures = []*Culture{BavarianCulture, DutchCulture, FranconianCulture, FrisianCulture, GermanCulture, LangobardCulture, SaxonCulture, SwabianCulture}

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

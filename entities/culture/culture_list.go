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

var HanCulture = &Culture{
	Name:            "chinese_han",
	Root:            ChineseRoot,
	Language:        language.Chinese,
	Ethos:           Courtly,
	MartialCustom:   g.OnlyMen,
	GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
	Traditions:      []*Tradition{CourtEunuchsTradition, FamilyBusinessTradition, MysticalAncestorsTradition, PhilosopherCultureTradition},
}

var AllChineseCultures = []*Culture{HanCulture}

// Dravidian

var (
	GondCulture = &Culture{
		Name:            "dravidian_gond",
		Root:            IndianRoot,
		Language:        language.Telugu,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{JungleWarriorsTradition, SacredGrovesTradition, CultureBlendingTradition},
	}
	KannadaCulture = &Culture{
		Name:            "dravidian_kannada",
		Root:            IndianRoot,
		Language:        language.Kannada,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FerventTempleBuildersTradition, GarudaWarriorsTradition, RefinedPoetryTradition, RulingCasteTradition},
	}
	TamilCulture = &Culture{
		Name:            "dravidian_tamil",
		Root:            IndianRoot,
		Language:        language.Tamil,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MetalworkersTradition, ParochialismTradition, SeafarersTradition},
	}
	TeluguCulture = &Culture{
		Name:            "dravidian_telugu",
		Root:            IndianRoot,
		Language:        language.Telugu,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MetalworkersTradition, RecognitionOfTalentTradition, WarriorCultureTradition},
	}
)

var AllDravidianCultures = []*Culture{GondCulture, KannadaCulture, TamilCulture, TeluguCulture}

// EastAfrican

var (
	DajuCulture = &Culture{
		Name:            "east_african_daju",
		Root:            AfricanRoot,
		Language:        language.Daju,
		Ethos:           Stoic,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DrylandDwellersTradition, WarriorCultureTradition, WarriorQueensTradition},
	}
	EthiopianCulture = &Culture{
		Name:            "east_african_ethiopian",
		Root:            AfricanRoot,
		Language:        language.Amharic,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FerventTempleBuildersTradition, IsolationistTradition, StrongBelieversTradition},
	}
	NubianCulture = &Culture{
		Name:            "east_african_nubian",
		Root:            AfricanRoot,
		Language:        language.HillNubian,
		Ethos:           Spiritual,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, AstuteDiplomatsTradition, LandOfTheBowTradition, WarriorQueensTradition, XenophilicTradition},
	}
	WelaytaCulture = &Culture{
		Name:            "east_african_welayta",
		Root:            AfricanRoot,
		Language:        language.SouthOmotic,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FrequentFestivitiesTradition, MountainSkirmishingTradition},
	}
)

var AllEastAfricanCultures = []*Culture{DajuCulture, EthiopianCulture, NubianCulture, WelaytaCulture}

// EastSlavic

var (
	IlmenianCulture = &Culture{
		Name:            "east_slavic_ilmenian",
		Root:            EuropeContinentalRoot,
		Language:        language.Ruthenian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, DruzhinaTradition, ForestFolkTradition, HitAndRunTacticiansTradition},
	}
	RuthenianCulture = &Culture{
		Name:            "east_slavic_ruthenian",
		Root:            EuropeContinentalRoot,
		Language:        language.Ruthenian,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DruzhinaTradition, AgrarianTradition, CharismaticTradition, CityKeepersTradition},
	}
	MoscovianCulture = &Culture{
		Name:            "east_slavic_moscovian",
		Root:            EuropeContinentalRoot,
		Language:        language.Russian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DruzhinaTradition, MendicantMysticsTradition, ForestFolkTradition},
	}
	SeverianCulture = &Culture{
		Name:            "east_slavic_severian",
		Root:            EuropeContinentalRoot,
		Language:        language.Ruthenian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DruzhinaTradition, ForestFolkTradition, SacredGrovesTradition},
	}
	VolhynianCulture = &Culture{
		Name:            "east_slavic_volhynian",
		Root:            EuropeContinentalRoot,
		Language:        language.Ruthenian,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DruzhinaTradition, AgrarianTradition, AstuteDiplomatsTradition, CityKeepersTradition},
	}
)

var AllEastSlavicCultures = []*Culture{IlmenianCulture, RuthenianCulture, MoscovianCulture, SeverianCulture, VolhynianCulture}

// Frankish

var (
	FrankishCulture = &Culture{
		Name:            "frankish_frankish",
		Root:            EuropeContinentalRoot,
		Language:        language.Frankish,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HereditaryHierarchyTradition, StandAndFightTradition, WarriorCultureTradition},
	}
	FrenchCulture = &Culture{
		Name:            "frankish_french",
		Proto:           []*Culture{FrankishCulture},
		Root:            EuropeContinentalRoot,
		Language:        language.French,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ChansonDeGesteTradition, ChivalryTradition, HereditaryHierarchyTradition},
	}
	NormanCulture = &Culture{
		Name:            "frankish_norman",
		Proto:           []*Culture{FrenchCulture /* @TODO /Norse */},
		Root:            EuropeContinentalRoot,
		Language:        language.French,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ChansonDeGesteTradition, HereditaryHierarchyTradition, StandAndFightTradition},
	}
	OccitanCulture = &Culture{
		Name:            "frankish_occitan",
		Root:            EuropeContinentalRoot,
		Language:        language.Occitan,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ChansonDeGesteTradition, ChivalryTradition, EqualInheritanceTradition, ParochialismTradition},
	}
)

var AllFrankishCultures = []*Culture{FrankishCulture, FrenchCulture, NormanCulture, OccitanCulture}

// Goidelic

var (
	GaelicCulture = &Culture{
		Name:            "goidelic_gaelic",
		Proto:           []*Culture{IrishCulture, PictishCulture},
		Root:            EuropeContinentalRoot,
		Language:        language.ScottishGaelic,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ConcubinesTradition, DexterousFishermenTradition, HighlandWarriorsTradition, HillDwellersTradition, StrongKinshipTradition},
	}
	IrishCulture = &Culture{
		Name:            "goidelic_irish",
		Root:            EuropeContinentalRoot,
		Language:        language.Irish,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MaritimeMercantilismTradition, MonasticCommunitiesTradition, PastorialistsTradition, PolygamousTradition, RefinedPoetryTradition},
	}
)

var AllGoidelicCultures = []*Culture{GaelicCulture, IrishCulture}

// GuineanUplander

var (
	BoboCulture = &Culture{
		Name:            "guinean_uplander_bobo",
		Root:            AfricanRoot,
		Language:        language.Soninke,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{StalwartDefendersTradition, UplandSkirmishingTradition, SorcerousMetallurgyTradition},
	}
	GurCulture = &Culture{
		Name:            "guinean_uplander_gur",
		Root:            AfricanRoot,
		Language:        language.Gur,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AncientMinersTradition, UplandSkirmishingTradition, SorcerousMetallurgyTradition},
	}
	MalinkeCulture = &Culture{
		Name:            "guinean_uplander_malinke",
		Root:            AfricanRoot,
		Language:        language.Manding,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CharitableTradition, UplandSkirmishingTradition, SorcerousMetallurgyTradition},
	}
	MarkaCulture = &Culture{
		Name:            "guinean_uplander_marka",
		Root:            AfricanRoot,
		Language:        language.Manding,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MusicalTheoristsTradition, UplandSkirmishingTradition, WarriorsOfTheDryTradition},
	}
	MelCulture = &Culture{
		Name:            "guinean_uplander_mel",
		Root:            AfricanRoot,
		Language:        language.Manding,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{StalwartDefendersTradition, UplandSkirmishingTradition, StaunchTraditionalistsTradition},
	}
)

var AllGuineanUplanderCultures = []*Culture{BoboCulture, GurCulture, MalinkeCulture, MarkaCulture, MelCulture}

// HornAfrican

var (
	AfarCulture = &Culture{
		Name:            "horn_african_afar",
		Root:            AfricanRoot,
		Language:        language.SahoAfar,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{EyeForAnEyeTradition, MonogamousTradition, MountainSkirmishingTradition, WarriorCultureTradition},
	}
	BejaCulture = &Culture{
		Name:            "horn_african_beja",
		Root:            AfricanRoot,
		Language:        language.Beja,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DesertTravelersTradition, IsolationistTradition, MountainSkirmishingTradition, WarriorsOfTheDryTradition},
	}
	SomaliCulture = &Culture{
		Name:            "horn_african_somali",
		Root:            AfricanRoot,
		Language:        language.Somali,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DrylandDwellersTradition, MaritimeMercantilismTradition, MountainSkirmishingTradition, RefinedPoetryTradition},
	}
)

var AllHornAfricanCultures = []*Culture{AfarCulture, BejaCulture, SomaliCulture}

// Iberian

var (
	AndalusianCulture = &Culture{
		Name:            "iberian_andalusian",
		Proto:           []*Culture{VisigothicCulture, ArabicBedouinCulture},
		Root:            MediterraneanRoot,
		Language:        language.Arabic,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			HitAndRunTacticiansTradition,
			ExpertArtisansTradition,
			MalleableSubjectsTradition,
			RitualizedFriendshipTradition,
			TabletopWarriorsTradition,
			XenophilicTradition,
		},
	}
	AragoneseCulture = &Culture{
		Name:            "iberian_aragonese",
		Proto:           []*Culture{BasqueCulture, CatalanCulture},
		Root:            MediterraneanRoot,
		Language:        language.Aragonese,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			MartialAdmirationTradition,
			MountaineersTradition,
			VisigothicCodesTradition,
			RitualizedFriendshipTradition,
			StateRansomingTradition,
			MaritalCeremoniesTradition,
		},
	}
	AsturleoneseCulture = &Culture{
		Name:            "iberian_asturleonese",
		Proto:           []*Culture{VisigothicCulture},
		Root:            MediterraneanRoot,
		Language:        language.Asturleonese,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FerventTempleBuildersTradition, HitAndRunTacticiansTradition, MountaineersTradition, RitualizedFriendshipTradition},
	}
	BasqueCulture = &Culture{
		Name:            "iberian_basque",
		Root:            MediterraneanRoot,
		Language:        language.Basque,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MaritalCeremoniesTradition, MountaineersTradition, VisigothicCodesTradition, RitualizedFriendshipTradition},
	}
	CastilianCulture = &Culture{
		Name:            "iberian_castilian",
		Proto:           []*Culture{VisigothicCulture},
		Root:            MediterraneanRoot,
		Language:        language.Spanish,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			CastleKeepersTradition,
			ChivalryTradition,
			MartialAdmirationTradition,
			RitualizedFriendshipTradition,
			TabletopWarriorsTradition,
			HitAndRunTacticiansTradition,
		},
	}
	CatalanCulture = &Culture{
		Name:            "iberian_catalan",
		Proto:           []*Culture{VisigothicCulture, OccitanCulture},
		Root:            MediterraneanRoot,
		Language:        language.Catalan,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			MaritimeMercantilismTradition,
			ParochialismTradition,
			RefinedPoetryTradition,
			VisigothicCodesTradition,
			RitualizedFriendshipTradition,
		},
	}
	GalicianCulture = &Culture{
		Name:            "iberian_galician",
		Proto:           []*Culture{VisigothicCulture},
		Root:            MediterraneanRoot,
		Language:        language.Galician,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MonasticCommunitiesTradition, DexterousFishermenTradition, HighlandWarriorsTradition, RitualizedFriendshipTradition},
	}
	PortugueseCulture = &Culture{
		Name:            "iberian_portuguese",
		Root:            MediterraneanRoot,
		Language:        language.Portuguese,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ChivalryTradition, FerventTempleBuildersTradition, MartialAdmirationTradition, RitualizedFriendshipTradition},
	}
	SuebiCulture = &Culture{
		Name:            "iberian_suebi",
		Root:            MediterraneanRoot,
		Language:        language.German,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, StandAndFightTradition},
	}
	VisigothicCulture = &Culture{
		Name:            "iberian_visigothic",
		Root:            MediterraneanRoot,
		Language:        language.Spanish,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{VisigothicCodesTradition, HitAndRunTacticiansTradition, MartialAdmirationTradition, MalleableSubjectsTradition, RitualizedFriendshipTradition},
	}
)

var AllIberianCultures = []*Culture{
	AndalusianCulture,
	AragoneseCulture,
	AsturleoneseCulture,
	BasqueCulture,
	CastilianCulture,
	CatalanCulture,
	GalicianCulture,
	PortugueseCulture,
	SuebiCulture,
	VisigothicCulture,
}

// IndoAryan

var (
	BengaliCulture = &Culture{
		Name:            "indo_aryan_bengali",
		Root:            IndianRoot,
		Language:        language.BengaliAssamese,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, PhilosopherCultureTradition, ReligiousPatronageTradition},
	}
	GujaratiCulture = &Culture{
		Name:            "indo_aryan_gujarati",
		Root:            IndianRoot,
		Language:        language.Gujarati,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ExpertArtisansTradition, ReligionBlendingTradition, SeafarersTradition, VegetariansTradition},
	}
	KamrupiCulture = &Culture{
		Name:            "indo_aryan_kamrupi",
		Root:            IndianRoot,
		Language:        language.Hindi,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CollectiveLandsTradition, JungleWarriorsTradition, CultureBlendingTradition},
	}
	KannaujiCulture = &Culture{
		Name:            "indo_aryan_kannauji",
		Root:            IndianRoot,
		Language:        language.Hindi,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, CityKeepersTradition, CourtEunuchsTradition, RulingCasteTradition},
	}
	KashmiriCulture = &Culture{
		Name:            "indo_aryan_kashmiri",
		Root:            IndianRoot,
		Language:        language.Kashmiri,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{EqualInheritanceTradition, HimalayanSettlersTradition, RefinedPoetryTradition},
	}
	MalviCulture = &Culture{
		Name:            "indo_aryan_malvi",
		Root:            IndianRoot,
		Language:        language.SinhalaMaldivian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, FerventTempleBuildersTradition, KhadgaPujaTradition, MysticalAncestorsTradition},
	}
	MarathiCulture = &Culture{
		Name:            "indo_aryan_marathi",
		Root:            IndianRoot,
		Language:        language.MarathiKonkani,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{IndustriousTradition, ModestTradition, PastorialistsTradition},
	}
	NepaliCulture = &Culture{
		Name:            "indo_aryan_nepali",
		Root:            IndianRoot,
		Language:        language.Nepali,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HimalayanSettlersTradition, MartialAdmirationTradition, ReligiousPatronageTradition},
	}
	OriyaCulture = &Culture{
		Name:            "indo_aryan_oriya",
		Root:            IndianRoot,
		Language:        language.Hindi,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CulinaryArtistsTradition, LordsOfTheElephantTradition, SeafarersTradition},
	}
	PunjabiCulture = &Culture{
		Name:            "indo_aryan_punjabi",
		Root:            IndianRoot,
		Language:        language.Punjabi,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, MaritalCeremoniesTradition, StorytellersTradition},
	}
	RajasthaniCulture = &Culture{
		Name:            "indo_aryan_rajasthani",
		Root:            IndianRoot,
		Language:        language.Rajasthani,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{KhadgaPujaTradition, MysticalAncestorsTradition, QuarrelsomeTradition, WarriorCultureTradition},
	}
	SindhiCulture = &Culture{
		Name:            "indo_aryan_sindhi",
		Root:            IndianRoot,
		Language:        language.Sindhi,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MendicantMysticsTradition, RefinedPoetryTradition, SeafarersTradition},
	}
	SinhalaCulture = &Culture{
		Name:            "indo_aryan_sinhala",
		Root:            IndianRoot,
		Language:        language.SinhalaMaldivian,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CityKeepersTradition, FerventTempleBuildersTradition, RulingCasteTradition},
	}
)

var AllIndoAryanCultures = []*Culture{
	BengaliCulture,
	GujaratiCulture,
	KamrupiCulture,
	KannaujiCulture,
	KashmiriCulture,
	MalviCulture,
	MarathiCulture,
	NepaliCulture,
	OriyaCulture,
	PunjabiCulture,
	RajasthaniCulture,
	SindhiCulture,
	SinhalaCulture,
}

// Iranian

var (
	AfghanCulture = &Culture{
		Name:            "iranian_afghan",
		Root:            PersianRoot,
		Language:        language.Pashto,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{EsteemedHospitalityTradition, FutuwaaTradition, LoyalSubjectsTradition, MountainHomesTradition},
	}
	BalochCulture = &Culture{
		Name:            "iranian_baloch",
		Root:            PersianRoot,
		Language:        language.Balochi,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{DesertTravelersTradition, FutuwaaTradition, EsteemedHospitalityTradition, IsolationistTradition},
	}
	DaylamiteCulture = &Culture{
		Name:            "iranian_daylamite",
		Root:            PersianRoot,
		Language:        language.Farsi,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{FutuwaaTradition, MountaineersTradition, StalwartDefendersTradition, SwordsForHireTradition},
	}
	KhwarezmianCulture = &Culture{
		Name:            "iranian_khwarezmian",
		Root:            PersianRoot,
		Language:        language.Scythian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, DrylandDwellersTradition, FutuwaaTradition, IsolationistTradition},
	}
	KurdishCulture = &Culture{
		Name:            "iranian_kurdish",
		Root:            PersianRoot,
		Language:        language.Kurdish,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{EyeForAnEyeTradition, FutuwaaTradition, MountainHomesTradition, SwordsForHireTradition},
	}
	PersianCulture = &Culture{
		Name:            "iranian_persian",
		Root:            PersianRoot,
		Language:        language.Farsi,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			DrylandDwellersTradition,
			FutuwaaTradition,
			GardenArchitectsTradition,
			PhilosopherCultureTradition,
			RefinedPoetryTradition,
		},
	}
	SakaCulture = &Culture{
		Name:            "iranian_saka",
		Root:            PersianRoot,
		Language:        language.Scythian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, DrylandDwellersTradition, ExpertArtisansTradition, HorseLordsTradition},
	}
	SogdianCulture = &Culture{
		Name:            "iranian_sogdian",
		Root:            PersianRoot,
		Language:        language.Scythian,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, FutuwaaTradition, ParochialismTradition, ReligionBlendingTradition},
	}
	TajikCulture = &Culture{
		Name:            "iranian_tajik",
		Root:            PersianRoot,
		Language:        language.Pashto,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, DrylandDwellersTradition, FutuwaaTradition, PhilosopherCultureTradition},
	}
)

var AllIranianCultures = []*Culture{AfghanCulture, BalochCulture, DaylamiteCulture, KhwarezmianCulture, KurdishCulture, PersianCulture, SakaCulture, SogdianCulture, TajikCulture}

// Israelite

var (
	AshkenaziCulture = &Culture{
		Name:            "istaelite_ashkenazi",
		Root:            SemiticRoot,
		Language:        language.Ashkenazi,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			DiasporicTradition,
			PhilosopherCultureTradition,
			BoundByFaithTradition,
			DefensiveTacticsTradition,
			ForebearingTradition,
		},
	}
	KochinimCulture = &Culture{
		Name:            "istaelite_kochinim",
		Proto:           []*Culture{SephardiCulture, TamilCulture},
		Root:            SemiticRoot,
		Language:        language.Hebrew,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			DiasporicTradition,
			PhilosopherCultureTradition,
			FerventTempleBuildersTradition,
			XenophilicTradition,
		},
	}
	RadhaniteCulture = &Culture{
		Name:            "istaelite_radhanite",
		Root:            SemiticRoot,
		Language:        language.Hebrew,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			DiasporicTradition,
			PhilosopherCultureTradition,
			DefensiveTacticsTradition,
			MaritimeMercantilismTradition,
			XenophilicTradition,
		},
	}
	SephardiCulture = &Culture{
		Name:            "istaelite_sephardi",
		Root:            SemiticRoot,
		Language:        language.Sephardi,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			DiasporicTradition,
			PhilosopherCultureTradition,
			CaravaneersTradition,
			DefensiveTacticsTradition,
			MaritimeMercantilismTradition,
		},
	}
)

var AllIsraeliteCultures = []*Culture{AshkenaziCulture, KochinimCulture, RadhaniteCulture, SephardiCulture}

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

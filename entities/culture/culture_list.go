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
		Proto:           []*Culture{FrisianCulture, FrankishCulture},
		Root:            EuropeContinentalRoot,
		Language:        language.Dutch,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, MaritimeMercantilismTradition, ParochialismTradition, PoldersTradition},
	}
	FranconianCulture = &Culture{
		Name:            "germanic_franconian",
		Proto:           []*Culture{FrankishCulture},
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
		Language:        language.Langobardian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{IsolationistTradition, MartialAdmirationTradition, StandAndFightTradition, MalleableInvadersTradition},
	}
	SaxonCulture = &Culture{
		Name:            "germanic_saxon",
		Proto:           []*Culture{OldSaxonCulture},
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
		Proto:           []*Culture{FrenchCulture, NorseCulture},
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
		Name:            "israelite_ashkenazi",
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
		Name:            "israelite_kochinim",
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
		Name:            "israelite_radhanite",
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
		Name:            "israelite_sephardi",
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

var (
	CisalpineCulture = &Culture{
		Name:            "latin_cisalpine",
		Proto:           []*Culture{LombardCulture, FrankishCulture},
		Root:            MediterraneanRoot,
		Language:        language.Italian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MaritimeMercantilismTradition, MountainHomesTradition, RepublicanLegacyTradition},
	}
	ItalianCulture = &Culture{
		Name:            "latin_italian",
		Proto:           []*Culture{RomanCulture},
		Root:            MediterraneanRoot,
		Language:        language.Italian,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FormationFightingExpertsTradition, RefinedPoetryTradition, RepublicanLegacyTradition},
	}
	LombardCulture = &Culture{
		Name:            "latin_lombard",
		Proto:           []*Culture{ItalianCulture, LangobardCulture},
		Root:            MediterraneanRoot,
		Language:        language.Lombard,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{IsolationistTradition, MartialAdmirationTradition, RepublicanLegacyTradition, StandAndFightTradition},
	}
	RomanCulture = &Culture{
		Name:            "latin_roman",
		Root:            MediterraneanRoot,
		Language:        language.Latin,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FormationFightingExpertsTradition, HereditaryHierarchyTradition, LegalisticTradition, RefinedPoetryTradition},
	}
	SardinianCulture = &Culture{
		Name:            "latin_sardinian",
		Proto:           []*Culture{RomanCulture},
		Root:            MediterraneanRoot,
		Language:        language.Sardinian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{IsolationistTradition, StalwartDefendersTradition},
	}
	SicilianCulture = &Culture{
		Name:            "latin_sicilian",
		Proto:           []*Culture{LombardCulture, GreekCulture},
		Root:            MediterraneanRoot,
		Language:        language.Sicilian,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{RefinedPoetryTradition, RepublicanLegacyTradition, RulingCasteTradition, SwordsForHireTradition, XenophilicTradition},
	}
)

var AllLatinCultures = []*Culture{CisalpineCulture, ItalianCulture, LombardCulture, RomanCulture, SardinianCulture, SicilianCulture}

// Magyar

var (
	HungarianCulture = &Culture{
		Name:            "magyar_hungarian",
		Proto:           []*Culture{MogyerCulture},
		Root:            MediterraneanRoot,
		Language:        language.Hungarian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, KonniRaidsTradition, StrongBelieversTradition},
	}
	MogyerCulture = &Culture{
		Name:            "magyar_mogyer",
		Root:            MediterraneanRoot,
		Language:        language.Hungarian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, StrongBelieversTradition, TribalUnityTradition},
	}
)

var AllMagyarCultures = []*Culture{HungarianCulture, MogyerCulture}

// Mongolic

var (
	BuryatCulture = &Culture{
		Name:            "mongolic_buryat",
		Root:            MongolianRoot,
		Language:        language.Buryat,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, MysticalAncestorsTradition, SacredMountainsTradition, SteppeToleranceTradition},
	}
	JurchenCulture = &Culture{
		Name:            "mongolic_jurchen",
		Root:            MongolianRoot,
		Language:        language.Xibe,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, HorseBreedersTradition, HorseLordsTradition, MalleableInvadersTradition},
	}
	KeraitCulture = &Culture{
		Name:            "mongolic_kerait",
		Root:            MongolianRoot,
		Language:        language.Mongolian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseBreedersTradition, HorseLordsTradition, ProlificHuntersTradition, SteppeToleranceTradition},
	}
	KhitanCulture = &Culture{
		Name:            "mongolic_khitan",
		Root:            MongolianRoot,
		Language:        language.Khitan,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, PastorialistsTradition, SacredHuntsTradition, MalleableInvadersTradition},
	}
	MongolCulture = &Culture{
		Name:            "mongolic_mongol",
		Root:            MongolianRoot,
		Language:        language.Mongolian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, ProlificHuntersTradition, SteppeToleranceTradition, MalleableInvadersTradition},
	}
	NaimanCulture = &Culture{
		Name:            "mongolic_naiman",
		Root:            MongolianRoot,
		Language:        language.Mongolian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CharismaticTradition, HorseLordsTradition, PastorialistsTradition, SteppeToleranceTradition},
	}
	OiratCulture = &Culture{
		Name:            "mongolic_oirat",
		Root:            MongolianRoot,
		Language:        language.Oirat,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, LoyalSubjectsTradition, ProlificHuntersTradition, StrongBelieversTradition},
	}
	OngudCulture = &Culture{
		Name:            "mongolic_ongud",
		Proto:           []*Culture{ShatuoCulture},
		Root:            MongolianRoot,
		Language:        language.Turkish,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AstuteDiplomatsTradition, HorseLordsTradition, StalwartDefendersTradition, SteppeToleranceTradition},
	}
	TuyuhunCulture = &Culture{
		Name:            "mongolic_tuyuhun",
		Root:            MongolianRoot,
		Language:        language.Mongolian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DrylandDwellersTradition, HorseBreedersTradition, HorseLordsTradition, PastorialistsTradition},
	}
)

var AllMongolicCultures = []*Culture{BuryatCulture, JurchenCulture, KeraitCulture, KhitanCulture, MongolCulture, NaimanCulture, OiratCulture, OngudCulture, TuyuhunCulture}

// NigerDelta

var (
	EdoCulture = &Culture{
		Name:            "niger_delta_edo",
		Root:            AfricanRoot,
		Language:        language.Kwa,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{BushHuntingTradition, HiddenCitiesTradition, ParochialismTradition},
	}
	EweCulture = &Culture{
		Name:            "niger_delta_ewe",
		Root:            AfricanRoot,
		Language:        language.Kwa,
		Ethos:           Spiritual,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{BushHuntingTradition, LegalisticTradition, ReligionBlendingTradition},
	}
	IgboCulture = &Culture{
		Name:            "niger_delta_igbo",
		Root:            AfricanRoot,
		Language:        language.Kwa,
		Ethos:           Spiritual,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{BushHuntingTradition, HiddenCitiesTradition, RecognitionOfTalentTradition},
	}
	YorubaCulture = &Culture{
		Name:            "niger_delta_yoruba",
		Root:            AfricanRoot,
		Language:        language.Kwa,
		Ethos:           Communal,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{BushHuntingTradition, EqualInheritanceTradition, HiddenCitiesTradition, PhilosopherCultureTradition},
	}
)

var AllNigerDeltaCultures = []*Culture{EdoCulture, EweCulture, IgboCulture, YorubaCulture}

// NorthGermanic

var (
	DanishCulture = &Culture{
		Name:            "north_germanic_danish",
		Root:            NordicEuropeRoot,
		Language:        language.Danish,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			EyeForAnEyeTradition,
			HereditaryHierarchyTradition,
			RunestoneRaisersTradition,
			TingMeetTradition,
			CoastalWarriorsTradition,
			HirdsTradition,
		},
	}
	NorseCulture = &Culture{
		Name:            "north_germanic_norse",
		Root:            NordicEuropeRoot,
		Language:        language.Swedish,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			TingMeetTradition,
			CoastalWarriorsTradition,
			MalleableInvadersTradition,
			NorthernStoriesTradition,
			PerformativeHonorTradition,
			HirdsTradition,
			RefinedPoetryTradition,
			RunestoneRaisersTradition,
		},
	}
	NorwegianCulture = &Culture{
		Name:            "north_germanic_norwegian",
		Root:            NordicEuropeRoot,
		Language:        language.Norwegian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			MaritimeMercantilismTradition,
			RunestoneRaisersTradition,
			StorytellersTradition,
			TingMeetTradition,
			CoastalWarriorsTradition,
			HirdsTradition,
		},
	}
	SwedishCulture = &Culture{
		Name:            "north_germanic_swedish",
		Proto:           []*Culture{NorseCulture},
		Root:            NordicEuropeRoot,
		Language:        language.Swedish,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions: []*Tradition{
			RunestoneRaisersTradition,
			StrongBelieversTradition,
			TingMeetTradition,
			CoastalWarriorsTradition,
			HirdsTradition,
		},
	}
)

var AllNorthGermanicCultures = []*Culture{DanishCulture, NorseCulture, NorwegianCulture, SwedishCulture}

// Qiangic

var (
	QiangCulture = &Culture{
		Name:            "qiangic_qiang",
		Root:            MongolianRoot,
		Language:        language.Tangut,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{EqualInheritanceTradition, HimalayanSettlersTradition, MedicinalHerbalistsTradition, SacredMountainsTradition},
	}
	TangutCulture = &Culture{
		Name:            "qiangic_tangut",
		Root:            MongolianRoot,
		Language:        language.Tangut,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, CourtEunuchsTradition, HimalayanSettlersTradition, ReligionBlendingTradition, ReligiousPatronageTradition},
	}
)

var AllQiangicCultures = []*Culture{QiangCulture, TangutCulture}

// Sahelian

var (
	BozoCulture = &Culture{
		Name:            "sahelian_bozo",
		Root:            AfricanRoot,
		Language:        language.Soninke,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{IsolationistTradition, ReligionBlendingTradition, WetlandersTradition},
	}
	GawCulture = &Culture{
		Name:            "sahelian_gaw",
		Root:            AfricanRoot,
		Language:        language.Zarma,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, DrylandDwellersTradition},
	}
	MossiCulture = &Culture{
		Name:            "sahelian_mossi",
		Root:            AfricanRoot,
		Language:        language.Gur,
		Ethos:           Bellicose,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AdaptiveSkirmishersTradition, HereditaryHierarchyTradition, RulingCasteTradition},
	}
	SonghaiCulture = &Culture{
		Name:            "sahelian_songhai",
		Root:            AfricanRoot,
		Language:        language.Korandje,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{DrylandDwellersTradition, RulingCasteTradition, WarriorCultureTradition},
	}
	SoninkeCulture = &Culture{
		Name:            "sahelian_soninke",
		Root:            AfricanRoot,
		Language:        language.Soninke,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, CityKeepersTradition, DrylandDwellersTradition, StrongBelieversTradition},
	}
	SorkoCulture = &Culture{
		Name:            "sahelian_sorko",
		Root:            AfricanRoot,
		Language:        language.KoyraboroSenni,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{DrylandDwellersTradition, XenophilicTradition},
	}
)

var AllSahelianCultures = []*Culture{BozoCulture, GawCulture, MossiCulture, SonghaiCulture, SoninkeCulture, SorkoCulture}

// Senegambian

var (
	PulaarCulture = &Culture{
		Name:            "senegambian_pulaar",
		Root:            AfricanRoot,
		Language:        language.Senegambian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CharismaticTradition, ForebearingTradition},
	}
	SererCulture = &Culture{
		Name:            "senegambian_serer",
		Root:            AfricanRoot,
		Language:        language.Senegambian,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HereditaryHierarchyTradition, StalwartDefendersTradition, StrongBelieversTradition},
	}
	WolofCulture = &Culture{
		Name:            "senegambian_wolof",
		Root:            AfricanRoot,
		Language:        language.Senegambian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{RulingCasteTradition, StrongBelieversTradition},
	}
)

var AllSenegambianCultures = []*Culture{PulaarCulture, SererCulture, WolofCulture}

// SouthSlavic

var (
	BosnianCulture = &Culture{
		Name:            "south_slavic_bosnian",
		Proto:           []*Culture{CroatianCulture, SerbianCulture},
		Root:            EuropeContinentalRoot,
		Language:        language.Bosnian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{IsolationistTradition, MendicantMysticsTradition, MonasticCommunitiesTradition},
	}
	BulgarianCulture = &Culture{
		Name:            "south_slavic_bulgarian",
		Proto:           []*Culture{CroatianCulture, BolgharCulture},
		Root:            EuropeContinentalRoot,
		Language:        language.Bulgarian,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{KonniRaidsTradition, MercifulBlindingsTradition, RulingCasteTradition, StandAndFightTradition},
	}
	CroatianCulture = &Culture{
		Name:            "south_slavic_croatian",
		Root:            EuropeContinentalRoot,
		Language:        language.Croatian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{AstuteDiplomatsTradition, KonniRaidsTradition, HereditaryHierarchyTradition},
	}
	SerbianCulture = &Culture{
		Name:            "south_slavic_serbian",
		Root:            EuropeContinentalRoot,
		Language:        language.Serbian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{KonniRaidsTradition, ReligiousPatronageTradition, StalwartDefendersTradition, StaunchTraditionalistsTradition},
	}
)

var AllSouthSlavicCultures = []*Culture{BosnianCulture, BulgarianCulture, CroatianCulture, SerbianCulture}

// Tibetan

var (
	BodpaCulture = &Culture{
		Name:            "tibetian_bodpa",
		Root:            IndianRoot,
		Language:        language.Tibetic,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HimalayanSettlersTradition, HorseBreedersTradition, ReligiousPatronageTradition},
	}
	KiratiCulture = &Culture{
		Name:            "tibetian_kirati",
		Root:            IndianRoot,
		Language:        language.Tibetic,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HimalayanSettlersTradition, ReligionBlendingTradition},
	}
	LhomonCulture = &Culture{
		Name:            "tibetian_lhomon",
		Root:            IndianRoot,
		Language:        language.Tibetic,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HimalayanSettlersTradition, MedicinalHerbalistsTradition, MysticalAncestorsTradition},
	}
	SumpaCulture = &Culture{
		Name:            "tibetian_sumpa",
		Root:            IndianRoot,
		Language:        language.Tibetic,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{EqualInheritanceTradition, HimalayanSettlersTradition, MountaineersTradition},
	}
	TsangpaCulture = &Culture{
		Name:            "tibetian_tsangpa",
		Root:            IndianRoot,
		Language:        language.Tibetic,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HimalayanSettlersTradition, ReligiousPatronageTradition, SacredMountainsTradition},
	}
	ZhangzhungCulture = &Culture{
		Name:            "tibetian_zhangzhung",
		Root:            IndianRoot,
		Language:        language.Tibetic,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HimalayanSettlersTradition, SacredMountainsTradition},
	}
)

var AllTibetanCultures = []*Culture{BodpaCulture, KiratiCulture, LhomonCulture, SumpaCulture, TsangpaCulture, ZhangzhungCulture}

// Tocharian

var TocharianCulture = &Culture{
	Name:            "tocharian",
	Root:            PersianRoot,
	Language:        language.Tocharian,
	Ethos:           Spiritual,
	MartialCustom:   g.OnlyMen,
	GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
	Traditions:      []*Tradition{CaravaneersTradition, PhilosopherCultureTradition, ReligiousPatronageTradition, XenophilicTradition},
}

var AllTocharianCultures = []*Culture{TocharianCulture}

// Turkic

var (
	AvarCulture = &Culture{
		Name:            "turkic_avar",
		Root:            TurkicRoot,
		Language:        language.Avar,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, RecognitionOfTalentTradition, RulingCasteTradition, CultureBlendingTradition},
	}
	BashkirCulture = &Culture{
		Name:            "turkic_bashkir",
		Root:            TurkicRoot,
		Language:        language.Bashkir,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{ForestFolkTradition, HorseLordsTradition, SacredMountainsTradition, MalleableInvadersTradition},
	}
	BolgharCulture = &Culture{
		Name:            "turkic_bolghar",
		Root:            TurkicRoot,
		Language:        language.Bulgar,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{AstuteDiplomatsTradition, HorseLordsTradition, RulingCasteTradition, MalleableInvadersTradition},
	}
	ChuvashCulture = &Culture{
		Name:            "turkic_chuvash",
		Proto:           []*Culture{BolgharCulture},
		Root:            TurkicRoot,
		Language:        language.Chuvash,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, SacredHuntsTradition, StrongBelieversTradition},
	}
	CumanCulture = &Culture{
		Name:            "turkic_cuman",
		Root:            TurkicRoot,
		Language:        language.Cuman,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, SwordsForHireTradition, WarriorCultureTradition, MalleableInvadersTradition},
	}
	KarlukCulture = &Culture{
		Name:            "turkic_karluk",
		Root:            TurkicRoot,
		Language:        language.Uzbek,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, SacredMountainsTradition, SteppeToleranceTradition, MalleableInvadersTradition},
	}
	KhazarCulture = &Culture{
		Name:            "turkic_khazar",
		Root:            TurkicRoot,
		Language:        language.Khazar,
		Ethos:           Egalitarian,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, StandAndFightTradition, SteppeToleranceTradition},
	}
	KimekCulture = &Culture{
		Name:            "turkic_kimek",
		Root:            TurkicRoot,
		Language:        language.Karaim,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, PastorialistsTradition, ProlificHuntersTradition, StalwartDefendersTradition},
	}
	KipchakCulture = &Culture{
		Name:            "turkic_kipchak",
		Root:            TurkicRoot,
		Language:        language.Karaim,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, ProlificHuntersTradition, QuarrelsomeTradition, SwordsForHireTradition},
	}
	KirghizCulture = &Culture{
		Name:            "turkic_kirghiz",
		Root:            TurkicRoot,
		Language:        language.Kyrgyz,
		Ethos:           Egalitarian,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, ProlificHuntersTradition, QuarrelsomeTradition},
	}
	LaktanCulture = &Culture{
		Name:            "turkic_laktan",
		Root:            TurkicRoot,
		Language:        language.Turkish,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, ForestFolkTradition, ProlificHuntersTradition, StalwartDefendersTradition},
	}
	OghuzCulture = &Culture{
		Name:            "turkic_oghuz",
		Root:            TurkicRoot,
		Language:        language.Turkish,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, RulingCasteTradition, SwordsForHireTradition, WarriorCultureTradition},
	}
	PechenegCulture = &Culture{
		Name:            "turkic_pecheneg",
		Root:            TurkicRoot,
		Language:        language.Khazar,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, QuarrelsomeTradition, SwordsForHireTradition, WarriorCultureTradition},
	}
	ShatuoCulture = &Culture{
		Name:            "turkic_shatuo",
		Root:            TurkicRoot,
		Language:        language.Chinese,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CourtEunuchsTradition, RulingCasteTradition, MalleableInvadersTradition},
	}
	UriankhaiCulture = &Culture{
		Name:            "turkic_uriankhai",
		Root:            TurkicRoot,
		Language:        language.Turkish,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, ForestFolkTradition, MendicantMysticsTradition, ProlificHuntersTradition},
	}
	UyghurCulture = &Culture{
		Name:            "turkic_uyghur",
		Root:            TurkicRoot,
		Language:        language.Uyghur,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, CharitableTradition, ReligiousPatronageTradition, SteppeToleranceTradition, XenophilicTradition},
	}
	YughurCulture = &Culture{
		Name:            "turkic_yughur",
		Proto:           []*Culture{UyghurCulture},
		Root:            TurkicRoot,
		Language:        language.Turkish,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, PastorialistsTradition, SteppeToleranceTradition},
	}
)

var AllTurkicCultures = []*Culture{
	AvarCulture,
	BashkirCulture,
	BolgharCulture,
	ChuvashCulture,
	CumanCulture,
	KarlukCulture,
	KhazarCulture,
	KimekCulture,
	KipchakCulture,
	KirghizCulture,
	LaktanCulture,
	OghuzCulture,
	PechenegCulture,
	ShatuoCulture,
	UriankhaiCulture,
	UyghurCulture,
	YughurCulture,
}

// UgroPermian

var (
	BjarmianCulture = &Culture{
		Name:            "ugro_permian_bjarmian",
		Root:            EuropeContinentalRoot,
		Language:        language.Finnish,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, CharismaticTradition, ForestWardensTradition, XenophilicTradition},
	}
	OstyakCulture = &Culture{
		Name:            "ugro_permian_ostyak",
		Root:            EuropeContinentalRoot,
		Language:        language.Permyak,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{IsolationistTradition, StrongBelieversTradition},
	}
	PermianCulture = &Culture{
		Name:            "ugro_permian_permian",
		Root:            EuropeContinentalRoot,
		Language:        language.Permyak,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{ForestWardensTradition, FrugalArmorsmithsTradition},
	}
)

var AllUgroPermianCultures = []*Culture{BjarmianCulture, OstyakCulture, PermianCulture}

// Vlach

var VlachCulture = &Culture{
	Name:            "vlach",
	Proto:           []*Culture{RomanCulture},
	Root:            MediterraneanRoot,
	Language:        language.Dacian,
	Ethos:           Communal,
	MartialCustom:   g.OnlyMen,
	GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
	Traditions:      []*Tradition{KonniRaidsTradition, PastorialistsTradition, StrongBelieversTradition, XenophilicTradition, StaunchTraditionalistsTradition},
}

var AllVlachCultures = []*Culture{VlachCulture}

// VolgaFinnic

var (
	MariCulture = &Culture{
		Name:            "volga_finnic_mari",
		Root:            EuropeContinentalRoot,
		Language:        language.Mari,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FrequentFestivitiesTradition, MusicalTheoristsTradition},
	}
	MeryaCulture = &Culture{
		Name:            "volga_finnic_merya",
		Root:            EuropeContinentalRoot,
		Language:        language.Merya,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{StorytellersTradition, CultureBlendingTradition},
	}
	MeshcheraCulture = &Culture{
		Name:            "volga_finnic_meshchera",
		Root:            EuropeContinentalRoot,
		Language:        language.Meshchera,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MusicalTheoristsTradition, ReligionBlendingTradition, CultureBlendingTradition},
	}
	MordvinCulture = &Culture{
		Name:            "volga_finnic_mordvin",
		Root:            EuropeContinentalRoot,
		Language:        language.Moksha,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FrequentFestivitiesTradition, MusicalTheoristsTradition, ReligionBlendingTradition},
	}
	MuromaCulture = &Culture{
		Name:            "volga_finnic_muroma",
		Root:            EuropeContinentalRoot,
		Language:        language.Mari,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FrequentFestivitiesTradition, StorytellersTradition, CultureBlendingTradition},
	}
)

var AllVolgaFinnicCultures = []*Culture{MariCulture, MeryaCulture, MeshcheraCulture, MordvinCulture, MuromaCulture}

// WestGermanic

var (
	AngloSaxonCulture = &Culture{
		Name:            "west_germanic_anglo_saxon",
		Proto:           []*Culture{OldSaxonCulture},
		Root:            EuropeContinentalRoot,
		Language:        language.English,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CityKeepersTradition, HirdsTradition, TheWitenagemotTradition},
	}
	EnglishCulture = &Culture{
		Name:            "west_germanic_english",
		Root:            EuropeContinentalRoot,
		Language:        language.English,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CastleKeepersTradition, ChivalryTradition, HereditaryHierarchyTradition, LongbowCompetitionsTradition},
	}
	OldSaxonCulture = &Culture{
		Name:            "west_germanic_old_saxon",
		Root:            EuropeContinentalRoot,
		Language:        language.Saxon,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HirdsTradition, SeafarersTradition, TingMeetTradition, MalleableInvadersTradition},
	}
	ScotsCulture = &Culture{
		Name:            "west_germanic_scots",
		Proto:           []*Culture{AngloSaxonCulture, CumbrianCulture},
		Root:            EuropeContinentalRoot,
		Language:        language.Scots,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CityKeepersTradition, HereditaryHierarchyTradition, StrongKinshipTradition, XenophilicTradition},
	}
)

var AllWestGermanicCultures = []*Culture{AngloSaxonCulture, EnglishCulture, OldSaxonCulture, ScotsCulture}

// WestSlavic

var (
	CarantanianCulture = &Culture{
		Name:            "west_slavic_carantanian",
		Root:            EuropeContinentalRoot,
		Language:        language.Croatian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MartialAdmirationTradition, MountaineerRuralismTradition},
	}
	CzechCulture = &Culture{
		Name:            "west_slavic_czech",
		Root:            EuropeContinentalRoot,
		Language:        language.Czech,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CastleKeepersTradition, IndustriousTradition, LifeIsJustAJokeTradition, MountaineerRuralismTradition},
	}
	PolabianCulture = &Culture{
		Name:            "west_slavic_polabian",
		Root:            EuropeContinentalRoot,
		Language:        language.Polish,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HereditaryHierarchyTradition, KonniRaidsTradition, StalwartDefendersTradition},
	}
	PolishCulture = &Culture{
		Name:            "west_slavic_polish",
		Root:            EuropeContinentalRoot,
		Language:        language.Polish,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{KonniRaidsTradition, XenophilicTradition, StaunchTraditionalistsTradition},
	}
	PomeranianCulture = &Culture{
		Name:            "west_slavic_pomeranian",
		Root:            EuropeContinentalRoot,
		Language:        language.Polish,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{KonniRaidsTradition, MaritimeMercantilismTradition, StaunchTraditionalistsTradition},
	}
	SlovienCulture = &Culture{
		Name:            "west_slavic_slovien",
		Root:            EuropeContinentalRoot,
		Language:        language.Slovenian,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{EquitableTradition, MountaineerRuralismTradition},
	}
)

var AllWestSlavicCultures = []*Culture{CarantanianCulture, CzechCulture, PolabianCulture, PolishCulture, PomeranianCulture, SlovienCulture}

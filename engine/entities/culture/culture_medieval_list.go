package culture

import (
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/influence"
	"persons_generator/engine/entities/language"

	"github.com/google/uuid"
)

// Akan
var (
	AkanCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "akan_akan",
		Abstuct:         AbsMedievalAkan,
		Root:            AfricanRoot,
		Language:        language.Kwa,
		Ethos:           Egalitarian,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.FemaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{BushHuntingTradition, MatriarchalTradition, MysticalAncestorsTradition, ParochialismTradition},
	}
	GuanCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "akan_guan",
		Abstuct:         AbsMedievalAkan,
		Root:            AfricanRoot,
		Language:        language.Kwa,
		Ethos:           Communal,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{BushHuntingTradition, DrylandDwellersTradition, FrequentFestivitiesTradition, StrongBelieversTradition},
	}
	KruCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "akan_kru",
		Abstuct:         AbsMedievalAkan,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "arabic_bedouin",
		Abstuct:         AbsAncientArabian,
		Root:            SemiticRoot,
		Language:        language.Arabic,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, DesertTravelersTradition, MubarizunTradition, TribalUnityTradition},
	}
	ArabicEgyptianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "arabic_egyptian",
		Abstuct:         AbsAncientArabian,
		Root:            SemiticRoot,
		Language:        language.Arabic,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{AgrarianTradition, MubarizunTradition, PhilosopherCultureTradition, StrongBelieversTradition},
	}
	ArabicMaghrebiCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Proto:           []*Culture{ArabicBedouinCulture, BaranisCulture},
		Name:            "arabic_maghrebi",
		Abstuct:         AbsAncientArabian,
		Root:            SemiticRoot,
		Language:        language.Arabic,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{AfricanToleranceTradition, DexterousFishermenTradition, DrylandDwellersTradition, MubarizunTradition, XenophilicTradition},
	}
	ArabicMashriqiCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Proto:           []*Culture{ArabicBedouinCulture},
		Name:            "arabic_mashriqi",
		Abstuct:         AbsAncientArabian,
		Root:            SemiticRoot,
		Language:        language.Arabic,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{DrylandDwellersTradition, MedicinalHerbalistsTradition, MubarizunTradition, PhilosopherCultureTradition},
	}
	ArabicYemeniCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "arabic_yemeni",
		Abstuct:         AbsAncientArabian,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "baltic_latgalian",
		Abstuct:         AbsMedievalBaltic,
		Root:            EuropeContinentalRoot,
		Language:        language.Latgalian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{ForestWardensTradition, SacredGrovesTradition, StaunchTraditionalistsTradition},
	}
	LithuanianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "baltic_lithuanian",
		Abstuct:         AbsMedievalBaltic,
		Root:            EuropeContinentalRoot,
		Language:        language.Lithuanian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{ForestWardensTradition, HitAndRunTacticiansTradition, SacredGrovesTradition, StrongBelieversTradition},
	}
	PrussianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "baltic_prussian",
		Abstuct:         AbsMedievalBaltic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "balto_finnic_estonian",
		Abstuct:         AbsMedievalBaltoFinnic,
		Root:            EuropeContinentalRoot,
		Language:        language.Estonian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MusicalTheoristsTradition, CoastalWarriorsTradition, HirdsTradition, StaunchTraditionalistsTradition},
	}
	FinnishCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "balto_finnic_finnish",
		Abstuct:         AbsMedievalBaltoFinnic,
		Root:            EuropeContinentalRoot,
		Language:        language.Finnish,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DexterousFishermenTradition, ForestWardensTradition, IsolationistTradition},
	}
	KarelianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "balto_finnic_karelian",
		Abstuct:         AbsMedievalBaltoFinnic,
		Root:            EuropeContinentalRoot,
		Language:        language.Karelian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ForestWardensTradition, StalwartDefendersTradition, SwordsForHireTradition},
	}
	SamiCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "balto_finnic_sami",
		Abstuct:         AbsMedievalBaltoFinnic,
		Root:            EuropeContinentalRoot,
		Language:        language.Sami,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AdaptiveSkirmishersTradition, ForestWardensTradition, WinterWarriorsTradition},
	}
	VepsianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "balto_finnic_vepsian",
		Abstuct:         AbsMedievalBaltoFinnic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "berber_baranis",
		Abstuct:         AbsMedievalBerber,
		Root:            BerberRoot,
		Language:        language.Tuareg,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{AfricanToleranceTradition, FamilyBusinessTradition, MountainHerdingTradition},
	}
	ButrCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "berber_butr",
		Abstuct:         AbsMedievalBerber,
		Root:            BerberRoot,
		Language:        language.Tuareg,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{AfricanToleranceTradition, DesertRibatsTradition, EqualInheritanceTradition, SaharanNomadsTradition},
	}
	GuancheCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "berber_guanche",
		Abstuct:         AbsMedievalBerber,
		Root:            BerberRoot,
		Language:        language.Tuareg,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{StalwartDefendersTradition, MysticalAncestorsTradition, IsolationistTradition},
	}
	ZaghawaCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "berber_zaghawa",
		Abstuct:         AbsMedievalBerber,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "brythonic_breton",
		Abstuct:         AbsMedievalBrythonic,
		Root:            EuropeContinentalRoot,
		Language:        language.Breton,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ChansonDeGesteTradition, DexterousFishermenTradition, StorytellersTradition, SwordsForHireTradition},
	}
	CornishCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "brythonic_cornish",
		Abstuct:         AbsMedievalBrythonic,
		Root:            EuropeContinentalRoot,
		Language:        language.Cornish,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AncientMinersTradition, DexterousFishermenTradition, StorytellersTradition},
	}
	CumbrianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "brythonic_cumbrian",
		Abstuct:         AbsMedievalBrythonic,
		Root:            EuropeContinentalRoot,
		Language:        language.Cumbric,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ConcubinesTradition, HighlandWarriorsTradition, RefinedPoetryTradition, StalwartDefendersTradition},
	}
	PictishCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "brythonic_pictish",
		Abstuct:         AbsMedievalBrythonic,
		Root:            EuropeContinentalRoot,
		Language:        language.Pictish,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ConcubinesTradition, HighlandWarriorsTradition, HillDwellersTradition, RefinedPoetryTradition},
	}
	WelshCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "brythonic_welsh",
		Abstuct:         AbsMedievalBrythonic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "burman_burmese",
		Abstuct:         AbsMedievalBurman,
		Root:            IndianRoot,
		Language:        language.Burmese,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{FerventTempleBuildersTradition, FrequentFestivitiesTradition, ReligionBlendingTradition, RoyalArmyTradition},
	}
	MonCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "burman_mon",
		Abstuct:         AbsMedievalBurman,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "byzantine_alan",
		Abstuct:         AbsMedievalByzantine,
		Root:            TurkicRoot,
		Language:        language.Scythian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, MaritalCeremoniesTradition, MountaineersTradition, SteppeToleranceTradition},
	}
	ArmenianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "byzantine_armenian",
		Abstuct:         AbsMedievalByzantine,
		Root:            CaucasianRoot,
		Language:        language.Armenian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{FerventTempleBuildersTradition, HighlandWarriorsTradition, HillDwellersTradition, StrongBelieversTradition},
	}
	GeorgianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "byzantine_georgian",
		Abstuct:         AbsMedievalByzantine,
		Root:            CaucasianRoot,
		Language:        language.Georgian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CastleKeepersTradition, CaucasianWolvesTradition, HereditaryHierarchyTradition, MetalworkersTradition},
	}
	GreekCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "byzantine_greek",
		Abstuct:         AbsMedievalByzantine,
		Root:            MediterraneanRoot,
		Language:        language.Greek,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{FormationFightingExpertsTradition, LegalisticTradition},
	}
	SyriacCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "byzantine_syriac",
		Abstuct:         AbsMedievalByzantine,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "african_hausa",
		Abstuct:         AbsMedievalCentralAfrican,
		Root:            AfricanRoot,
		Language:        language.Hausa,
		Ethos:           Egalitarian,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, FamilyBusinessTradition, ParochialismTradition},
	}
	KanuriCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "african_kanuri",
		Abstuct:         AbsMedievalCentralAfrican,
		Root:            AfricanRoot,
		Language:        language.Tebu,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HitAndRunTacticiansTradition, QuarrelsomeTradition},
	}
	NupeCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "african_nupe",
		Abstuct:         AbsMedievalCentralAfrican,
		Root:            AfricanRoot,
		Language:        language.Kwa,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MusicalTheoristsTradition, ParochialismTradition},
	}
	SaoCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "african_sao",
		Abstuct:         AbsMedievalCentralAfrican,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_bavarian",
		Abstuct:         AbsMedievalCentralGermanic,
		Root:            EuropeContinentalRoot,
		Language:        language.German,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HereditaryHierarchyTradition, StandAndFightTradition, StrongBelieversTradition},
	}
	DutchCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_dutch",
		Proto:           []*Culture{FrisianCulture, FrankishCulture},
		Abstuct:         AbsMedievalCentralGermanic,
		Root:            EuropeContinentalRoot,
		Language:        language.Dutch,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, MaritimeMercantilismTradition, ParochialismTradition, PoldersTradition},
	}
	FranconianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_franconian",
		Proto:           []*Culture{FrankishCulture},
		Abstuct:         AbsMedievalCentralGermanic,
		Root:            EuropeContinentalRoot,
		Language:        language.German,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AstuteDiplomatsTradition, CastleKeepersTradition, HereditaryHierarchyTradition},
	}
	FrisianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_frisian",
		Abstuct:         AbsMedievalCentralGermanic,
		Root:            EuropeContinentalRoot,
		Language:        language.Frisian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{BattlefieldLootersTradition, DexterousFishermenTradition, StrongBelieversTradition},
	}
	GermanCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_german",
		Abstuct:         AbsMedievalCentralGermanic,
		Root:            EuropeContinentalRoot,
		Language:        language.German,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AstuteDiplomatsTradition, HereditaryHierarchyTradition, StandAndFightTradition},
	}
	LangobardCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_langobard",
		Abstuct:         AbsMedievalCentralGermanic,
		Root:            EuropeContinentalRoot,
		Language:        language.Langobardian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{IsolationistTradition, MartialAdmirationTradition, StandAndFightTradition, MalleableInvadersTradition},
	}
	SaxonCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_saxon",
		Proto:           []*Culture{OldSaxonCulture},
		Abstuct:         AbsMedievalCentralGermanic,
		Root:            EuropeContinentalRoot,
		Language:        language.Saxon,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ReligiousPatronageTradition, RulingCasteTradition, StandAndFightTradition},
	}
	SwabianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_swabian",
		Abstuct:         AbsMedievalCentralGermanic,
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
	ID:              uuid.NewString(),
	Origin:          NativeOrigin,
	Name:            "chinese_han",
	Abstuct:         AbsMedievalChinese,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "dravidian_gond",
		Abstuct:         AbsAncientDravidian,
		Root:            IndianRoot,
		Language:        language.Telugu,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{JungleWarriorsTradition, SacredGrovesTradition, CultureBlendingTradition},
	}
	KannadaCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "dravidian_kannada",
		Abstuct:         AbsAncientDravidian,
		Root:            IndianRoot,
		Language:        language.Kannada,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FerventTempleBuildersTradition, GarudaWarriorsTradition, RefinedPoetryTradition, RulingCasteTradition},
	}
	TamilCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "dravidian_tamil",
		Abstuct:         AbsAncientDravidian,
		Root:            IndianRoot,
		Language:        language.Tamil,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MetalworkersTradition, ParochialismTradition, SeafarersTradition},
	}
	TeluguCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "dravidian_telugu",
		Abstuct:         AbsAncientDravidian,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "african_daju",
		Abstuct:         AbsMedievalEastAfrican,
		Root:            AfricanRoot,
		Language:        language.Daju,
		Ethos:           Stoic,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DrylandDwellersTradition, WarriorCultureTradition, WarriorQueensTradition},
	}
	EthiopianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "african_ethiopian",
		Abstuct:         AbsMedievalEastAfrican,
		Root:            AfricanRoot,
		Language:        language.Amharic,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FerventTempleBuildersTradition, IsolationistTradition, StrongBelieversTradition},
	}
	NubianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "african_nubian",
		Abstuct:         AbsMedievalEastAfrican,
		Root:            AfricanRoot,
		Language:        language.HillNubian,
		Ethos:           Spiritual,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, AstuteDiplomatsTradition, LandOfTheBowTradition, WarriorQueensTradition, XenophilicTradition},
	}
	WelaytaCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "african_welayta",
		Abstuct:         AbsMedievalEastAfrican,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_ilmenian",
		Abstuct:         AbsMedievalEastSlavic,
		Root:            EuropeContinentalRoot,
		Language:        language.Ruthenian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, DruzhinaTradition, ForestFolkTradition, HitAndRunTacticiansTradition},
	}
	RuthenianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_ruthenian",
		Abstuct:         AbsMedievalEastSlavic,
		Root:            EuropeContinentalRoot,
		Language:        language.Ruthenian,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DruzhinaTradition, AgrarianTradition, CharismaticTradition, CityKeepersTradition},
	}
	MoscovianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_moscovian",
		Abstuct:         AbsMedievalEastSlavic,
		Root:            EuropeContinentalRoot,
		Language:        language.Russian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DruzhinaTradition, MendicantMysticsTradition, ForestFolkTradition},
	}
	SeverianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_severian",
		Proto:           []*Culture{RuthenianCulture},
		Abstuct:         AbsMedievalEastSlavic,
		Root:            EuropeContinentalRoot,
		Language:        language.Ruthenian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DruzhinaTradition, ForestFolkTradition, SacredGrovesTradition},
	}
	VolhynianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_volhynian",
		Proto:           []*Culture{RuthenianCulture},
		Abstuct:         AbsMedievalEastSlavic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "frankish_frankish",
		Abstuct:         AbsMedievalFrankish,
		Root:            EuropeContinentalRoot,
		Language:        language.Frankish,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HereditaryHierarchyTradition, StandAndFightTradition, WarriorCultureTradition},
	}
	FrenchCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "frankish_french",
		Proto:           []*Culture{FrankishCulture},
		Abstuct:         AbsMedievalFrankish,
		Root:            EuropeContinentalRoot,
		Language:        language.French,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ChansonDeGesteTradition, ChivalryTradition, HereditaryHierarchyTradition},
	}
	NormanCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "frankish_norman",
		Proto:           []*Culture{FrenchCulture, NorseCulture},
		Abstuct:         AbsMedievalFrankish,
		Root:            EuropeContinentalRoot,
		Language:        language.French,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ChansonDeGesteTradition, HereditaryHierarchyTradition, StandAndFightTradition},
	}
	OccitanCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "frankish_occitan",
		Abstuct:         AbsMedievalFrankish,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "goidelic_gaelic",
		Proto:           []*Culture{IrishCulture, PictishCulture},
		Abstuct:         AbsMedievalGoidelic,
		Root:            EuropeContinentalRoot,
		Language:        language.ScottishGaelic,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ConcubinesTradition, DexterousFishermenTradition, HighlandWarriorsTradition, HillDwellersTradition, StrongKinshipTradition},
	}
	IrishCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "goidelic_irish",
		Abstuct:         AbsMedievalGoidelic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "guinean_bobo",
		Abstuct:         AbsMedievalGuineanUplander,
		Root:            AfricanRoot,
		Language:        language.Soninke,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{StalwartDefendersTradition, UplandSkirmishingTradition, SorcerousMetallurgyTradition},
	}
	GurCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "guinean_gur",
		Abstuct:         AbsMedievalGuineanUplander,
		Root:            AfricanRoot,
		Language:        language.Gur,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AncientMinersTradition, UplandSkirmishingTradition, SorcerousMetallurgyTradition},
	}
	MalinkeCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "guinean_malinke",
		Abstuct:         AbsMedievalGuineanUplander,
		Root:            AfricanRoot,
		Language:        language.Manding,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CharitableTradition, UplandSkirmishingTradition, SorcerousMetallurgyTradition},
	}
	MarkaCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "guinean_marka",
		Abstuct:         AbsMedievalGuineanUplander,
		Root:            AfricanRoot,
		Language:        language.Manding,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MusicalTheoristsTradition, UplandSkirmishingTradition, WarriorsOfTheDryTradition},
	}
	MelCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "guinean_mel",
		Abstuct:         AbsMedievalGuineanUplander,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "african_afar",
		Abstuct:         AbsMedievalHornAfrican,
		Root:            AfricanRoot,
		Language:        language.SahoAfar,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{EyeForAnEyeTradition, MonogamousTradition, MountainSkirmishingTradition, WarriorCultureTradition},
	}
	BejaCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "african_beja",
		Abstuct:         AbsMedievalHornAfrican,
		Root:            AfricanRoot,
		Language:        language.Beja,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{DesertTravelersTradition, IsolationistTradition, MountainSkirmishingTradition, WarriorsOfTheDryTradition},
	}
	SomaliCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "african_somali",
		Abstuct:         AbsMedievalHornAfrican,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iberian_andalusian",
		Proto:           []*Culture{VisigothicCulture, ArabicBedouinCulture},
		Abstuct:         AbsMedievalIberian,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iberian_aragonese",
		Proto:           []*Culture{BasqueCulture, CatalanCulture},
		Abstuct:         AbsMedievalIberian,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iberian_asturleonese",
		Proto:           []*Culture{VisigothicCulture},
		Abstuct:         AbsMedievalIberian,
		Root:            MediterraneanRoot,
		Language:        language.Asturleonese,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FerventTempleBuildersTradition, HitAndRunTacticiansTradition, MountaineersTradition, RitualizedFriendshipTradition},
	}
	BasqueCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iberian_basque",
		Abstuct:         AbsMedievalIberian,
		Root:            MediterraneanRoot,
		Language:        language.Basque,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MaritalCeremoniesTradition, MountaineersTradition, VisigothicCodesTradition, RitualizedFriendshipTradition},
	}
	CastilianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iberian_castilian",
		Proto:           []*Culture{VisigothicCulture},
		Abstuct:         AbsMedievalIberian,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iberian_catalan",
		Proto:           []*Culture{VisigothicCulture, OccitanCulture},
		Abstuct:         AbsMedievalIberian,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iberian_galician",
		Proto:           []*Culture{VisigothicCulture},
		Abstuct:         AbsMedievalIberian,
		Root:            MediterraneanRoot,
		Language:        language.Galician,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MonasticCommunitiesTradition, DexterousFishermenTradition, HighlandWarriorsTradition, RitualizedFriendshipTradition},
	}
	PortugueseCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iberian_portuguese",
		Abstuct:         AbsMedievalIberian,
		Root:            MediterraneanRoot,
		Language:        language.Portuguese,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ChivalryTradition, FerventTempleBuildersTradition, MartialAdmirationTradition, RitualizedFriendshipTradition},
	}
	SuebiCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iberian_suebi",
		Abstuct:         AbsMedievalIberian,
		Root:            MediterraneanRoot,
		Language:        language.German,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, StandAndFightTradition},
	}
	VisigothicCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iberian_visigothic",
		Abstuct:         AbsMedievalIberian,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "indo_aryan_bengali",
		Abstuct:         AbsMedievalIndoAryan,
		Root:            IndianRoot,
		Language:        language.BengaliAssamese,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, PhilosopherCultureTradition, ReligiousPatronageTradition},
	}
	GujaratiCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "indo_aryan_gujarati",
		Abstuct:         AbsMedievalIndoAryan,
		Root:            IndianRoot,
		Language:        language.Gujarati,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{ExpertArtisansTradition, ReligionBlendingTradition, SeafarersTradition, VegetariansTradition},
	}
	KamrupiCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "indo_aryan_kamrupi",
		Abstuct:         AbsMedievalIndoAryan,
		Root:            IndianRoot,
		Language:        language.Hindi,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CollectiveLandsTradition, JungleWarriorsTradition, CultureBlendingTradition},
	}
	KannaujiCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "indo_aryan_kannauji",
		Abstuct:         AbsMedievalIndoAryan,
		Root:            IndianRoot,
		Language:        language.Hindi,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, CityKeepersTradition, CourtEunuchsTradition, RulingCasteTradition},
	}
	KashmiriCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "indo_aryan_kashmiri",
		Abstuct:         AbsMedievalIndoAryan,
		Root:            IndianRoot,
		Language:        language.Kashmiri,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{EqualInheritanceTradition, HimalayanSettlersTradition, RefinedPoetryTradition},
	}
	MalviCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "indo_aryan_malvi",
		Abstuct:         AbsMedievalIndoAryan,
		Root:            IndianRoot,
		Language:        language.SinhalaMaldivian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, FerventTempleBuildersTradition, KhadgaPujaTradition, MysticalAncestorsTradition},
	}
	MarathiCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "indo_aryan_marathi",
		Abstuct:         AbsMedievalIndoAryan,
		Root:            IndianRoot,
		Language:        language.MarathiKonkani,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{IndustriousTradition, ModestTradition, PastorialistsTradition},
	}
	NepaliCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "indo_aryan_nepali",
		Abstuct:         AbsMedievalIndoAryan,
		Root:            IndianRoot,
		Language:        language.Nepali,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HimalayanSettlersTradition, MartialAdmirationTradition, ReligiousPatronageTradition},
	}
	OriyaCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "indo_aryan_oriya",
		Abstuct:         AbsMedievalIndoAryan,
		Root:            IndianRoot,
		Language:        language.Hindi,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CulinaryArtistsTradition, LordsOfTheElephantTradition, SeafarersTradition},
	}
	PunjabiCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "indo_aryan_punjabi",
		Abstuct:         AbsMedievalIndoAryan,
		Root:            IndianRoot,
		Language:        language.Punjabi,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, MaritalCeremoniesTradition, StorytellersTradition},
	}
	RajasthaniCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "indo_aryan_rajasthani",
		Abstuct:         AbsMedievalIndoAryan,
		Root:            IndianRoot,
		Language:        language.Rajasthani,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{KhadgaPujaTradition, MysticalAncestorsTradition, QuarrelsomeTradition, WarriorCultureTradition},
	}
	SindhiCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "indo_aryan_sindhi",
		Abstuct:         AbsMedievalIndoAryan,
		Root:            IndianRoot,
		Language:        language.Sindhi,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MendicantMysticsTradition, RefinedPoetryTradition, SeafarersTradition},
	}
	SinhalaCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "indo_aryan_sinhala",
		Abstuct:         AbsMedievalIndoAryan,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iranian_afghan",
		Abstuct:         AbsMedievalIranian,
		Root:            PersianRoot,
		Language:        language.Pashto,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{EsteemedHospitalityTradition, FutuwaaTradition, LoyalSubjectsTradition, MountainHomesTradition},
	}
	BalochCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iranian_baloch",
		Abstuct:         AbsMedievalIranian,
		Root:            PersianRoot,
		Language:        language.Balochi,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{DesertTravelersTradition, FutuwaaTradition, EsteemedHospitalityTradition, IsolationistTradition},
	}
	DaylamiteCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iranian_daylamite",
		Abstuct:         AbsMedievalIranian,
		Root:            PersianRoot,
		Language:        language.Farsi,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{FutuwaaTradition, MountaineersTradition, StalwartDefendersTradition, SwordsForHireTradition},
	}
	KhwarezmianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iranian_khwarezmian",
		Abstuct:         AbsMedievalIranian,
		Root:            PersianRoot,
		Language:        language.Scythian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, DrylandDwellersTradition, FutuwaaTradition, IsolationistTradition},
	}
	KurdishCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iranian_kurdish",
		Abstuct:         AbsMedievalIranian,
		Root:            PersianRoot,
		Language:        language.Kurdish,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{EyeForAnEyeTradition, FutuwaaTradition, MountainHomesTradition, SwordsForHireTradition},
	}
	PersianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iranian_persian",
		Abstuct:         AbsMedievalIranian,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iranian_saka",
		Abstuct:         AbsMedievalIranian,
		Root:            PersianRoot,
		Language:        language.Scythian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, DrylandDwellersTradition, ExpertArtisansTradition, HorseLordsTradition},
	}
	SogdianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iranian_sogdian",
		Abstuct:         AbsMedievalIranian,
		Root:            PersianRoot,
		Language:        language.Scythian,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, FutuwaaTradition, ParochialismTradition, ReligionBlendingTradition},
	}
	TajikCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "iranian_tajik",
		Abstuct:         AbsMedievalIranian,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "israelite_ashkenazi",
		Abstuct:         AbsMedievalIsraelite,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "israelite_kochinim",
		Proto:           []*Culture{SephardiCulture, TamilCulture},
		Abstuct:         AbsMedievalIsraelite,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "israelite_radhanite",
		Abstuct:         AbsMedievalIsraelite,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "israelite_sephardi",
		Abstuct:         AbsMedievalIsraelite,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "latin_cisalpine",
		Proto:           []*Culture{LombardCulture, FrankishCulture},
		Abstuct:         AbsMedievalLatin,
		Root:            MediterraneanRoot,
		Language:        language.Italian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MaritimeMercantilismTradition, MountainHomesTradition, RepublicanLegacyTradition},
	}
	ItalianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "latin_italian",
		Proto:           []*Culture{RomanCulture},
		Abstuct:         AbsMedievalLatin,
		Root:            MediterraneanRoot,
		Language:        language.Italian,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FormationFightingExpertsTradition, RefinedPoetryTradition, RepublicanLegacyTradition},
	}
	LombardCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "latin_lombard",
		Proto:           []*Culture{ItalianCulture, LangobardCulture},
		Abstuct:         AbsMedievalLatin,
		Root:            MediterraneanRoot,
		Language:        language.Lombard,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{IsolationistTradition, MartialAdmirationTradition, RepublicanLegacyTradition, StandAndFightTradition},
	}
	RomanCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "latin_roman",
		Abstuct:         AbsMedievalLatin,
		Root:            MediterraneanRoot,
		Language:        language.Latin,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FormationFightingExpertsTradition, HereditaryHierarchyTradition, LegalisticTradition, RefinedPoetryTradition},
	}
	SardinianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "latin_sardinian",
		Proto:           []*Culture{RomanCulture},
		Abstuct:         AbsMedievalLatin,
		Root:            MediterraneanRoot,
		Language:        language.Sardinian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{IsolationistTradition, StalwartDefendersTradition},
	}
	SicilianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "latin_sicilian",
		Proto:           []*Culture{LombardCulture, GreekCulture},
		Abstuct:         AbsMedievalLatin,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "magyar_hungarian",
		Proto:           []*Culture{MogyerCulture},
		Abstuct:         AbsMedievalMagyar,
		Root:            MediterraneanRoot,
		Language:        language.Hungarian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, KonniRaidsTradition, StrongBelieversTradition},
	}
	MogyerCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "magyar_mogyer",
		Abstuct:         AbsMedievalMagyar,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "mongolic_buryat",
		Abstuct:         AbsMedievalMongolic,
		Root:            MongolianRoot,
		Language:        language.Buryat,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, MysticalAncestorsTradition, SacredMountainsTradition, SteppeToleranceTradition},
	}
	JurchenCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "mongolic_jurchen",
		Abstuct:         AbsMedievalMongolic,
		Root:            MongolianRoot,
		Language:        language.Xibe,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AgrarianTradition, HorseBreedersTradition, HorseLordsTradition, MalleableInvadersTradition},
	}
	KeraitCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "mongolic_kerait",
		Abstuct:         AbsMedievalMongolic,
		Root:            MongolianRoot,
		Language:        language.Mongolian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseBreedersTradition, HorseLordsTradition, ProlificHuntersTradition, SteppeToleranceTradition},
	}
	KhitanCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "mongolic_khitan",
		Abstuct:         AbsMedievalMongolic,
		Root:            MongolianRoot,
		Language:        language.Khitan,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, PastorialistsTradition, SacredHuntsTradition, MalleableInvadersTradition},
	}
	MongolCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "mongolic_mongol",
		Abstuct:         AbsMedievalMongolic,
		Root:            MongolianRoot,
		Language:        language.Mongolian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, ProlificHuntersTradition, SteppeToleranceTradition, MalleableInvadersTradition},
	}
	NaimanCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "mongolic_naiman",
		Abstuct:         AbsMedievalMongolic,
		Root:            MongolianRoot,
		Language:        language.Mongolian,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CharismaticTradition, HorseLordsTradition, PastorialistsTradition, SteppeToleranceTradition},
	}
	OiratCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "mongolic_oirat",
		Abstuct:         AbsMedievalMongolic,
		Root:            MongolianRoot,
		Language:        language.Oirat,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, LoyalSubjectsTradition, ProlificHuntersTradition, StrongBelieversTradition},
	}
	OngudCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "mongolic_ongud",
		Proto:           []*Culture{ShatuoCulture},
		Abstuct:         AbsMedievalMongolic,
		Root:            MongolianRoot,
		Language:        language.Turkish,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AstuteDiplomatsTradition, HorseLordsTradition, StalwartDefendersTradition, SteppeToleranceTradition},
	}
	TuyuhunCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "mongolic_tuyuhun",
		Abstuct:         AbsMedievalMongolic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "niger_delta_edo",
		Abstuct:         AbsMedievalNigerDelta,
		Root:            AfricanRoot,
		Language:        language.Kwa,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{BushHuntingTradition, HiddenCitiesTradition, ParochialismTradition},
	}
	EweCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "niger_delta_ewe",
		Abstuct:         AbsMedievalNigerDelta,
		Root:            AfricanRoot,
		Language:        language.Kwa,
		Ethos:           Spiritual,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{BushHuntingTradition, LegalisticTradition, ReligionBlendingTradition},
	}
	IgboCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "niger_delta_igbo",
		Abstuct:         AbsMedievalNigerDelta,
		Root:            AfricanRoot,
		Language:        language.Kwa,
		Ethos:           Spiritual,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{BushHuntingTradition, HiddenCitiesTradition, RecognitionOfTalentTradition},
	}
	YorubaCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "niger_delta_yoruba",
		Abstuct:         AbsMedievalNigerDelta,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_danish",
		Abstuct:         AbsMedievalNorthGermanic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_norse",
		Abstuct:         AbsMedievalNorthGermanic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_norwegian",
		Abstuct:         AbsMedievalNorthGermanic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_swedish",
		Proto:           []*Culture{NorseCulture},
		Abstuct:         AbsMedievalNorthGermanic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "qiangic_qiang",
		Abstuct:         AbsMedievalQiangic,
		Root:            MongolianRoot,
		Language:        language.Tangut,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{EqualInheritanceTradition, HimalayanSettlersTradition, MedicinalHerbalistsTradition, SacredMountainsTradition},
	}
	TangutCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "qiangic_tangut",
		Abstuct:         AbsMedievalQiangic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "sahelian_bozo",
		Abstuct:         AbsMedievalSahelian,
		Root:            AfricanRoot,
		Language:        language.Soninke,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{IsolationistTradition, ReligionBlendingTradition, WetlandersTradition},
	}
	GawCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "sahelian_gaw",
		Abstuct:         AbsMedievalSahelian,
		Root:            AfricanRoot,
		Language:        language.Zarma,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, DrylandDwellersTradition},
	}
	MossiCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "sahelian_mossi",
		Abstuct:         AbsMedievalSahelian,
		Root:            AfricanRoot,
		Language:        language.Gur,
		Ethos:           Bellicose,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{AdaptiveSkirmishersTradition, HereditaryHierarchyTradition, RulingCasteTradition},
	}
	SonghaiCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "sahelian_songhai",
		Abstuct:         AbsMedievalSahelian,
		Root:            AfricanRoot,
		Language:        language.Korandje,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{DrylandDwellersTradition, RulingCasteTradition, WarriorCultureTradition},
	}
	SoninkeCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "sahelian_soninke",
		Abstuct:         AbsMedievalSahelian,
		Root:            AfricanRoot,
		Language:        language.Soninke,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, CityKeepersTradition, DrylandDwellersTradition, StrongBelieversTradition},
	}
	SorkoCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "sahelian_sorko",
		Abstuct:         AbsMedievalSahelian,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "senegambian_pulaar",
		Abstuct:         AbsMedievalSenegambian,
		Root:            AfricanRoot,
		Language:        language.Senegambian,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CharismaticTradition, ForebearingTradition},
	}
	SererCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "senegambian_serer",
		Abstuct:         AbsMedievalSenegambian,
		Root:            AfricanRoot,
		Language:        language.Senegambian,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HereditaryHierarchyTradition, StalwartDefendersTradition, StrongBelieversTradition},
	}
	WolofCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "senegambian_wolof",
		Abstuct:         AbsMedievalSenegambian,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_bosnian",
		Proto:           []*Culture{CroatianCulture, SerbianCulture},
		Abstuct:         AbsMedievalSouthSlavic,
		Root:            EuropeContinentalRoot,
		Language:        language.Bosnian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{IsolationistTradition, MendicantMysticsTradition, MonasticCommunitiesTradition},
	}
	BulgarianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_bulgarian",
		Proto:           []*Culture{CroatianCulture, BolgharCulture},
		Abstuct:         AbsMedievalSouthSlavic,
		Root:            EuropeContinentalRoot,
		Language:        language.Bulgarian,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{KonniRaidsTradition, MercifulBlindingsTradition, RulingCasteTradition, StandAndFightTradition},
	}
	CroatianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_croatian",
		Abstuct:         AbsMedievalSouthSlavic,
		Root:            EuropeContinentalRoot,
		Language:        language.Croatian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{AstuteDiplomatsTradition, KonniRaidsTradition, HereditaryHierarchyTradition},
	}
	SerbianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_serbian",
		Abstuct:         AbsMedievalSouthSlavic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "tibetian_bodpa",
		Abstuct:         AbsMedievalTibetan,
		Root:            IndianRoot,
		Language:        language.Tibetic,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HimalayanSettlersTradition, HorseBreedersTradition, ReligiousPatronageTradition},
	}
	KiratiCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "tibetian_kirati",
		Abstuct:         AbsMedievalTibetan,
		Root:            IndianRoot,
		Language:        language.Tibetic,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HimalayanSettlersTradition, ReligionBlendingTradition},
	}
	LhomonCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "tibetian_lhomon",
		Abstuct:         AbsMedievalTibetan,
		Root:            IndianRoot,
		Language:        language.Tibetic,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HimalayanSettlersTradition, MedicinalHerbalistsTradition, MysticalAncestorsTradition},
	}
	SumpaCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "tibetian_sumpa",
		Abstuct:         AbsMedievalTibetan,
		Root:            IndianRoot,
		Language:        language.Tibetic,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{EqualInheritanceTradition, HimalayanSettlersTradition, MountaineersTradition},
	}
	TsangpaCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "tibetian_tsangpa",
		Abstuct:         AbsMedievalTibetan,
		Root:            IndianRoot,
		Language:        language.Tibetic,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HimalayanSettlersTradition, ReligiousPatronageTradition, SacredMountainsTradition},
	}
	ZhangzhungCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "tibetian_zhangzhung",
		Abstuct:         AbsMedievalTibetan,
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
	ID:              uuid.NewString(),
	Origin:          NativeOrigin,
	Name:            "tocharian",
	Abstuct:         AbsMedievalTocharian,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_avar",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Avar,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, RecognitionOfTalentTradition, RulingCasteTradition, CultureBlendingTradition},
	}
	BashkirCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_bashkir",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Bashkir,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{ForestFolkTradition, HorseLordsTradition, SacredMountainsTradition, MalleableInvadersTradition},
	}
	BolgharCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_bolghar",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Bulgar,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{AstuteDiplomatsTradition, HorseLordsTradition, RulingCasteTradition, MalleableInvadersTradition},
	}
	ChuvashCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_chuvash",
		Proto:           []*Culture{BolgharCulture},
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Chuvash,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, SacredHuntsTradition, StrongBelieversTradition},
	}
	CumanCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_cuman",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Cuman,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, SwordsForHireTradition, WarriorCultureTradition, MalleableInvadersTradition},
	}
	KarlukCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_karluk",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Uzbek,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, SacredMountainsTradition, SteppeToleranceTradition, MalleableInvadersTradition},
	}
	KhazarCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_khazar",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Khazar,
		Ethos:           Egalitarian,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, StandAndFightTradition, SteppeToleranceTradition},
	}
	KimekCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_kimek",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Karaim,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, PastorialistsTradition, ProlificHuntersTradition, StalwartDefendersTradition},
	}
	KipchakCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_kipchak",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Karaim,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, ProlificHuntersTradition, QuarrelsomeTradition, SwordsForHireTradition},
	}
	KirghizCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_kirghiz",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Kyrgyz,
		Ethos:           Egalitarian,
		MartialCustom:   g.MenAndWomen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, ProlificHuntersTradition, QuarrelsomeTradition},
	}
	LaktanCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_laktan",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Turkish,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, ForestFolkTradition, ProlificHuntersTradition, StalwartDefendersTradition},
	}
	OghuzCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_oghuz",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Turkish,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, RulingCasteTradition, SwordsForHireTradition, WarriorCultureTradition},
	}
	PechenegCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_pecheneg",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Khazar,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, QuarrelsomeTradition, SwordsForHireTradition, WarriorCultureTradition},
	}
	ShatuoCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_shatuo",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Chinese,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CourtEunuchsTradition, RulingCasteTradition, MalleableInvadersTradition},
	}
	UriankhaiCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_uriankhai",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Turkish,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{HorseLordsTradition, ForestFolkTradition, MendicantMysticsTradition, ProlificHuntersTradition},
	}
	UyghurCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_uyghur",
		Abstuct:         AbsMedievalTurkic,
		Root:            TurkicRoot,
		Language:        language.Uyghur,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, CharitableTradition, ReligiousPatronageTradition, SteppeToleranceTradition, XenophilicTradition},
	}
	YughurCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "turkic_yughur",
		Proto:           []*Culture{UyghurCulture},
		Abstuct:         AbsMedievalTurkic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "ugro_permian_bjarmian",
		Abstuct:         AbsMedievalUgroPermian,
		Root:            EuropeContinentalRoot,
		Language:        language.Finnish,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{CaravaneersTradition, CharismaticTradition, ForestWardensTradition, XenophilicTradition},
	}
	OstyakCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "ugro_permian_ostyak",
		Abstuct:         AbsMedievalUgroPermian,
		Root:            EuropeContinentalRoot,
		Language:        language.Permyak,
		Ethos:           Spiritual,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.StrongInfluence),
		Traditions:      []*Tradition{IsolationistTradition, StrongBelieversTradition},
	}
	PermianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "ugro_permian_permian",
		Abstuct:         AbsMedievalUgroPermian,
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
	ID:              uuid.NewString(),
	Origin:          NativeOrigin,
	Name:            "vlach",
	Proto:           []*Culture{RomanCulture},
	Abstuct:         AbsMedievalVlach,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "volga_finnic_mari",
		Abstuct:         AbsMedievalVolgaFinnic,
		Root:            EuropeContinentalRoot,
		Language:        language.Mari,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FrequentFestivitiesTradition, MusicalTheoristsTradition},
	}
	MeryaCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "volga_finnic_merya",
		Abstuct:         AbsMedievalVolgaFinnic,
		Root:            EuropeContinentalRoot,
		Language:        language.Merya,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{StorytellersTradition, CultureBlendingTradition},
	}
	MeshcheraCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "volga_finnic_meshchera",
		Abstuct:         AbsMedievalVolgaFinnic,
		Root:            EuropeContinentalRoot,
		Language:        language.Meshchera,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MusicalTheoristsTradition, ReligionBlendingTradition, CultureBlendingTradition},
	}
	MordvinCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "volga_finnic_mordvin",
		Abstuct:         AbsMedievalVolgaFinnic,
		Root:            EuropeContinentalRoot,
		Language:        language.Moksha,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{FrequentFestivitiesTradition, MusicalTheoristsTradition, ReligionBlendingTradition},
	}
	MuromaCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "volga_finnic_muroma",
		Abstuct:         AbsMedievalVolgaFinnic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_anglo_saxon",
		Proto:           []*Culture{OldSaxonCulture},
		Abstuct:         AbsMedievalWestGermanic,
		Root:            EuropeContinentalRoot,
		Language:        language.English,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CityKeepersTradition, HirdsTradition, TheWitenagemotTradition},
	}
	EnglishCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_english",
		Abstuct:         AbsMedievalWestGermanic,
		Root:            EuropeContinentalRoot,
		Language:        language.English,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CastleKeepersTradition, ChivalryTradition, HereditaryHierarchyTradition, LongbowCompetitionsTradition},
	}
	OldSaxonCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_old_saxon",
		Abstuct:         AbsMedievalWestGermanic,
		Root:            EuropeContinentalRoot,
		Language:        language.Saxon,
		Ethos:           Bellicose,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HirdsTradition, SeafarersTradition, TingMeetTradition, MalleableInvadersTradition},
	}
	ScotsCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "germanic_scots",
		Proto:           []*Culture{AngloSaxonCulture, CumbrianCulture},
		Abstuct:         AbsMedievalWestGermanic,
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
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_carantanian",
		Abstuct:         AbsMedievalWestSlavic,
		Root:            EuropeContinentalRoot,
		Language:        language.Croatian,
		Ethos:           Communal,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{MartialAdmirationTradition, MountaineerRuralismTradition},
	}
	CzechCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_czech",
		Abstuct:         AbsMedievalWestSlavic,
		Root:            EuropeContinentalRoot,
		Language:        language.Czech,
		Ethos:           Bureaucratic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{CastleKeepersTradition, IndustriousTradition, LifeIsJustAJokeTradition, MountaineerRuralismTradition},
	}
	PolabianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_polabian",
		Abstuct:         AbsMedievalWestSlavic,
		Root:            EuropeContinentalRoot,
		Language:        language.Polish,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{HereditaryHierarchyTradition, KonniRaidsTradition, StalwartDefendersTradition},
	}
	PolishCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_polish",
		Abstuct:         AbsMedievalWestSlavic,
		Root:            EuropeContinentalRoot,
		Language:        language.Polish,
		Ethos:           Courtly,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{KonniRaidsTradition, XenophilicTradition, StaunchTraditionalistsTradition},
	}
	PomeranianCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_pomeranian",
		Abstuct:         AbsMedievalWestSlavic,
		Root:            EuropeContinentalRoot,
		Language:        language.Polish,
		Ethos:           Stoic,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.MaleDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{KonniRaidsTradition, MaritimeMercantilismTradition, StaunchTraditionalistsTradition},
	}
	SlovienCulture = &Culture{
		ID:              uuid.NewString(),
		Origin:          NativeOrigin,
		Name:            "slavic_slovien",
		Abstuct:         AbsMedievalWestSlavic,
		Root:            EuropeContinentalRoot,
		Language:        language.Slovenian,
		Ethos:           Egalitarian,
		MartialCustom:   g.OnlyMen,
		GenderDominance: g.NewDominanceWithParams(g.EqualityDominance, influence.ModerateInfluence),
		Traditions:      []*Tradition{EquitableTradition, MountaineerRuralismTradition},
	}
)

var AllWestSlavicCultures = []*Culture{CarantanianCulture, CzechCulture, PolabianCulture, PolishCulture, PomeranianCulture, SlovienCulture}

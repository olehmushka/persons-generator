package culture

import "persons_generator/tools"

var AllTraditions = tools.Merge(AllRealmTraditions, AllWarfareTraditions, AllSocialTraditions, AllRitualTraditions, AllRegionalTraditions)

// Realm Traditions
var (
	// It is a common idea in this culture that the purest expression of superiority is the ability to subdue your enemy without fighting them.
	AstuteDiplomatsTradition = &Tradition{
		Name:             "astute_diplomats",
		PreferredEthoses: []*Ethos{Bureaucratic, Communal, Courtly, Egalitarian, Spiritual, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// The idea that individuals personally own land is alien to this culture; land belongs to all people, for the common good.
	CollectiveLandsTradition = &Tradition{
		Name:             "collective_lands",
		PreferredEthoses: []*Ethos{Communal, Egalitarian, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// In this culture rulers are known, and expected, to welcome anyone who might come to their court with open arms.
	// Hosting guests and visitors in a spectacular fashion.
	EsteemedHospitalityTradition = &Tradition{
		Name:             "esteemed_hospitality",
		PreferredEthoses: []*Ethos{Communal, Courtly, Spiritual},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture has a long tradition of constructing
	// and maintaining various types of gardens, building themselves a small paradise.
	GardenArchitectsTradition = &Tradition{
		Name:             "garden_architects",
		PreferredEthoses: []*Ethos{Communal, Courtly, Spiritual},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture prefers to keep to itself, and doesn't often look outside its own sphere.
	IsolationistTradition = &Tradition{
		Name:             "isolationist",
		PreferredEthoses: []*Ethos{Communal, Spiritual, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// For this culture, the union of two people in marriage is considered a highly public and ritualistic affair.
	MaritalCeremoniesTradition = &Tradition{
		Name:             "marital_ceremonies",
		PreferredEthoses: []*Ethos{Communal, Courtly, Spiritual},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture has enjoyed easy access to iron for centuries.
	// Over time they have developed their understanding of metal to such a degree that
	// their name has become a byword for durable and high quality arms and armor.
	MetalworkersTradition = &Tradition{
		Name:             "metalworkers",
		PreferredEthoses: []*Ethos{Bureaucratic, Communal, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// Rulers of this culture are used to lording over subjects that are not their own culture.
	// They know how to effectively suppress revolts, though at the cost of public perception.
	RulingCasteTradition = &Tradition{
		Name:             "ruling_caste",
		PreferredEthoses: []*Ethos{Courtly, Spiritual},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture is distinctly agrarian, their lands producing ample crops for hungry armies.
	// While not having to struggle for food means many peasants for the levy,
	// it also means that each soldier is not as motivated as those from harsher regions.
	AgrarianTradition = &Tradition{
		Name:             "agrarian",
		PreferredEthoses: []*Ethos{Bureaucratic, Communal, Courtly, Egalitarian, Spiritual, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// Offense is not a good defense. A good defense is having the most castles on the face of the known world.
	CastleKeepersTradition = &Tradition{
		Name:             "castle_keepers",
		PreferredEthoses: []*Ethos{Bellicose, Bureaucratic, Spiritual},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "city_keepers" {
					return false
				}
			}
			return true
		},
	}
	// Cities are the beating urban heart of this culture, and they wish for every one of their metropolises to become enviable gems known across the world.
	CityKeepersTradition = &Tradition{
		Name:             "city_keepers",
		PreferredEthoses: []*Ethos{Bureaucratic, Courtly, Egalitarian},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "castle_keepers" {
					return false
				}
			}
			return true
		},
	}
	// This culture makes extensive use of eunuchs at court, occupying positions ranging from domestic servants to bureaucratic administrators and even military commanders.
	// Men with no desire are easy to trust.
	CourtEunuchsTradition = &Tradition{
		Name:             "court_eunuchs",
		PreferredEthoses: []*Ethos{Communal, Courtly, Spiritual},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture is accustomed to living in dry climates, and know where to find water and how to work the lands.
	DrylandDwellersTradition = &Tradition{
		Name:             "dryland_dwellers",
		PreferredEthoses: []*Ethos{Bureaucratic, Communal, Courtly, Egalitarian, Spiritual, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture has a history of rulers dividing their lands between all their children, even those who normally would be excluded.
	EqualInheritanceTradition = &Tradition{
		Name:             "equal_inheritance",
		PreferredEthoses: []*Ethos{Egalitarian},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "warrior_queens" {
					return false
				}
			}
			if c.GenderDominance.IsEquality() {
				return true
			}

			return false
		},
	}
	// This culture encourages businesses to develop along familial lines, each generation picking up the tools and
	// techniques of their trade from the last, building a strong tradition of ancestral professionalism
	FamilyBusinessTradition = &Tradition{
		Name:             "family_business",
		PreferredEthoses: []*Ethos{Communal, Courtly},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture finds constructing temples the most direct path to the divine. Fervent rulers are expected to build many, and grand, temples.
	FerventTempleBuildersTradition = &Tradition{
		Name:             "fervent_temple_builders",
		PreferredEthoses: []*Ethos{Communal, Spiritual, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture lives close to vast forests, and knows how to forage its bounties.
	ForestFolkTradition = &Tradition{
		Name:             "forest_folk",
		PreferredEthoses: []*Ethos{Bureaucratic, Communal, Courtly, Egalitarian, Spiritual, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// It is ingrained in this culture that leadership in passed on by blood, nobility is bound by stone in castles, and that the feudal structure is unshakable.
	// They do not easily forget those who break the feudal code.
	HereditaryHierarchyTradition = &Tradition{
		Name:             "hereditary_hierarchy",
		PreferredEthoses: []*Ethos{Courtly, Spiritual},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// Whether they are up in the jungle trees, or protected by the organic maze, foreigners may not even know this culture's cities exist, let alone have a clue how to besiege them.
	HiddenCitiesTradition = &Tradition{
		Name:             "hidden_cities",
		PreferredEthoses: []*Ethos{Bureaucratic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "jungle_dwellers" {
					return false
				}
			}
			return true
		},
	}
	// This culture is at home in hills, and know how to work its lands effectively.
	HillDwellersTradition = &Tradition{
		Name:             "hill_dwellers",
		PreferredEthoses: []*Ethos{Bureaucratic, Communal, Courtly, Egalitarian, Spiritual, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture lives close to the jungle, and knows how to harvest its riches.
	JungleDwellersTradition = &Tradition{
		Name:             "jungle_dwellers",
		PreferredEthoses: []*Ethos{Bureaucratic, Communal, Courtly, Egalitarian, Spiritual, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "hidden_cities" {
					return false
				}
			}
			return true
		},
	}
	// This culture regards the rule of law and its codification as being the most important parameter of a civilized society.
	LegalisticTradition = &Tradition{
		Name:             "legalistic",
		PreferredEthoses: []*Ethos{Communal, Courtly, Egalitarian},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// The world may be ruled by armies, but this culture knows that it is truly controlled by whoever dominates the flow of gold across the seas.
	MaritimeMercantilismTradition = &Tradition{
		Name:             "maritime_mercantilism",
		PreferredEthoses: []*Ethos{Bureaucratic, Egalitarian, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture enforces a matriarchal hierarchy, where the ruling class is overwhelmingly comprised of women.
	MatriarchalTradition = &Tradition{
		Name:             "matriarchal",
		PreferredEthoses: []*Ethos{Egalitarian},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == " warrior_queens" {
					return false
				}
			}
			if c.GenderDominance.IsMaleDominate() || c.GenderDominance.IsEquality() {
				return false
			}

			return true
		},
	}
	// This culture has a long history of sending noble progeny to serve the faith. Revered political thinkers often arise from the monastic orders.
	MonasticCommunitiesTradition = &Tradition{
		Name:             "monastic_communities",
		PreferredEthoses: []*Ethos{Communal, Spiritual},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture knows how to live and thrive on plateaus near the harsh slopes of mountains.
	MountainHomesTradition = &Tradition{
		Name:             "mountain_homes",
		PreferredEthoses: []*Ethos{Bureaucratic, Communal, Courtly, Egalitarian, Spiritual, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// City residents of this culture are fiercely competitive and independent.
	// They invest a lot of money and energy to make sure that their city is the grandest.
	ParochialismTradition = &Tradition{
		Name:             "parochialism",
		PreferredEthoses: []*Ethos{Communal, Courtly, Spiritual},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture is at home in wide and open terrain, where they herd large groups of animals.
	PastorialistsTradition = &Tradition{
		Name:             "pastorialists",
		PreferredEthoses: []*Ethos{Bellicose, Communal, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture has a long tradition of reclaiming land from the surrounding sea.
	// This allows them to utilize what would normally be a shallow seabed for farming and construction.
	PoldersTradition = &Tradition{
		Name:             "polders",
		PreferredEthoses: []*Ethos{Bureaucratic, Communal, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture has a long history of republican values, where the urban classes are just as important as the rural landholders.
	RepublicanLegacyTradition = &Tradition{
		Name:             "republican_legacy",
		PreferredEthoses: []*Ethos{Communal, Courtly, Spiritual},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "parochialism" {
					return false
				}
			}
			return true
		},
	}
	// This culture emphasizes the familial and communal bonds its people share, pushing them to remember their shared heritage.
	TribalUnityTradition = &Tradition{
		Name:             "tribal_unity",
		PreferredEthoses: []*Ethos{Communal, Spiritual, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture has mastered the bogs and marshes.
	// Though life is at times a battle against mold, they have learned to use the peat to their advantage.
	WetlandersTradition = &Tradition{
		Name:             "wetlanders",
		PreferredEthoses: []*Ethos{Bureaucratic, Communal, Courtly, Egalitarian, Spiritual, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// This culture has always been surrounded by raw materials, ore, and uncut gems.
	// They have an affinity for finding sites for mining excavations.
	AncientMinersTradition = &Tradition{
		Name:             "ancient_miners",
		PreferredEthoses: []*Ethos{},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// We have always been drawn to the shoreline.
	// The sting of sea air, the crying gulls, the scent of smoked fish...
	// what warrior would ever opt to live far from the water and its freedoms?
	CoastalWarriorsTradition = &Tradition{
		Name:             "coastal_warriors",
		PreferredEthoses: []*Ethos{Bellicose},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "warrior_culture":
					fallthrough
				case "seafarers":
					return false
				}
			}

			return true
		},
	}
	// Members of this culture have often lived alongside people of other cultures.
	// Through exposure, they have become adept at incorporating foreign languages, traditions and customs into their own lives.
	CultureBlendingTradition = &Tradition{
		Name:             "culture_blending",
		PreferredEthoses: []*Ethos{Communal, Egalitarian},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// Tradition is tradition, it is immovable and unchangeable.
	// We must be ready and willing to stand up for what makes us who we are.
	StaunchTraditionalistsTradition = &Tradition{
		Name:             "staunch_traditionalists",
		PreferredEthoses: []*Ethos{Communal, Spiritual, Stoic},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// Who's to say what freedom is? One king is much like another regardless of culture.
	// True freedom comes from learning how to live with any liege.
	MalleableSubjectsTradition = &Tradition{
		Name:             "malleable_subjects",
		PreferredEthoses: []*Ethos{Courtly, Egalitarian},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
	// Every soldier goes into battle ready to risk life and limb for their liege!
	// Does it not behoove the liege to offer them some guarantee of safety, regardless of station, should they be captured in the field?
	StateRansomingTradition = &Tradition{
		Name:             "state_ransoming",
		PreferredEthoses: []*Ethos{Bellicose, Communal},
		Type:             RealmTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return true
		},
	}
)

var AllRealmTraditions = []*Tradition{
	AstuteDiplomatsTradition,
	CollectiveLandsTradition,
	EsteemedHospitalityTradition,
	GardenArchitectsTradition,
	IsolationistTradition,
	MaritalCeremoniesTradition,
	FamilyBusinessTradition,
	FerventTempleBuildersTradition,
	ForestFolkTradition,
	HereditaryHierarchyTradition,
	HiddenCitiesTradition,
	HillDwellersTradition,
	JungleDwellersTradition,
	LegalisticTradition,
	MaritimeMercantilismTradition,
	MatriarchalTradition,
	MonasticCommunitiesTradition,
	MountainHomesTradition,
	ParochialismTradition,
	PastorialistsTradition,
	PoldersTradition,
	RepublicanLegacyTradition,
	TribalUnityTradition,
	WetlandersTradition,
	AncientMinersTradition,
	CoastalWarriorsTradition,
	CultureBlendingTradition,
	StaunchTraditionalistsTradition,
	MalleableSubjectsTradition,
	StateRansomingTradition,
}

// Warfare Traditions

var (
	// This culture knows how to effectively field skirmishers in any environment.
	AdaptiveSkirmishersTradition = &Tradition{
		Name:              "adaptive_skirmishers",
		PreferredEthoses:  []*Ethos{Bellicose, Spiritual, Stoic},
		Type:              WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// In this culture battles are not fought for prestige, but for profit.
	// Who cares about one's standing if you're unable to pay for the war?
	BattlefieldLootersTradition = &Tradition{
		Name:              "battlefield_looters",
		PreferredEthoses:  []*Ethos{Bellicose, Bureaucratic},
		Type:              WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture knows how to create synergy between different types of units by use of efficient protective formations.
	FormationFightingExpertsTradition = &Tradition{
		Name:              "formation_fighting_experts",
		PreferredEthoses:  []*Ethos{Bellicose, Communal, Courtly},
		Type:              WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// While a set of high-quality armor might save one life, having ten decent sets might win a battle.
	FrugalArmorsmithsTradition = &Tradition{
		Name:              "frugal_armorsmiths",
		PreferredEthoses:  []*Ethos{Bellicose, Communal, Stoic},
		Type:              WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// "Kill me if you wish, but if you let me live, I shall serve you loyally."
	RecognitionOfTalentTradition = &Tradition{
		Name:              "recognition_of_talent",
		PreferredEthoses:  []*Ethos{Bellicose, Courtly, Egalitarian},
		Type:              WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture doesn't look down upon those who can no longer fight due to injury, instead they are celebrated as heroes and used as teachers.
	ReverenceForVeteransTradition = &Tradition{
		Name:              "Reverence_for_veterans",
		PreferredEthoses:  []*Ethos{Bellicose, Egalitarian, Stoic},
		Type:              WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Defending that which is one's own is of paramount importance to this culture.
	StalwartDefendersTradition = &Tradition{
		Name:              "stalwart_defenders",
		PreferredEthoses:  []*Ethos{Bellicose, Courtly, Stoic},
		Type:              WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture forsakes having elite troops, and instead favors massed armies.
	StrengthInNumbersTradition = &Tradition{
		Name:              "strength_in_numbers",
		PreferredEthoses:  []*Ethos{Bellicose, Spiritual},
		Type:              WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture views mercenary work favorably and encourages warriors to seek glory as mercenaries in-between wars.
	SwordsForHireTradition = &Tradition{
		Name:              "swords_for_hire",
		PreferredEthoses:  []*Ethos{Bellicose, Communal, Courtly},
		Type:              WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture believes that if you choose to pursue theological studies, you must also be able to defend your faith.
	WarriorPriestsTradition = &Tradition{
		Name:              "warrior_priests",
		PreferredEthoses:  []*Ethos{Bellicose, Spiritual},
		Type:              WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture believes that if you've proven yourself capable as a warrior, you should be allowed to fight — no matter who you are.
	WarriorsByMeritTradition = &Tradition{
		Name:              "warriors_by_merit",
		PreferredEthoses:  []*Ethos{Bellicose, Communal, Egalitarian},
		Type:              WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture is well-versed at fighting in forests.
	ForestFightersTradition = &Tradition{
		Name:             "forest_fighters",
		PreferredEthoses: []*Ethos{Bellicose, Communal, Stoic},
		Type:             WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "pacifists" {
					return false
				}
			}
			return true
		},
	}
	// Warriors of this culture fight well in the slopes and valleys of their hilly homes.
	HighlandWarriorsTradition = &Tradition{
		Name:             "highland_warriors",
		PreferredEthoses: []*Ethos{Bellicose, Communal, Stoic},
		Type:             WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "pacifists":
					fallthrough
				case "upland_skirmishing":
					return false
				}
			}
			return true
		},
	}
	// This culture has mastered the use of lightly-armored units to hit the enemy hard, and then fall back.
	HitAndRunTacticiansTradition = &Tradition{
		Name:             "hit_and_run_tacticians",
		PreferredEthoses: []*Ethos{Bellicose, Egalitarian, Spiritual},
		Type:             WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "konni_raids" {
					return false
				}
			}
			return true
		},
	}
	// This culture has painstakingly accumulated knowledge and experience in the fine art and science of horse breeding.
	// Whether destriers or coursers, the horses of these people are renowned for their superiority.
	HorseBreedersTradition = &Tradition{
		Name:              "horse_breeders",
		PreferredEthoses:  []*Ethos{Bellicose, Communal, Stoic},
		Type:              WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Warriors of this culture know well how to traverse treacherous jungles.
	JungleWarriorsTradition = &Tradition{
		Name:             "jungle_warriors",
		PreferredEthoses: []*Ethos{Bellicose, Communal, Stoic},
		Type:             WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "pacifists":
					fallthrough
				case "bush_hunting":
					return false
				}
			}

			return true
		},
	}
	// This culture favors a bow more powerful than most could draw, and
	// practicing with them once a week is as dear as any ritual of worship everywhere from the most rural villages to the largest urban metropoles.
	LongbowCompetitionsTradition = &Tradition{
		Name:              "longbow_competitions",
		PreferredEthoses:  []*Ethos{Bureaucratic, Stoic},
		Type:              WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Soldiers of this culture carry all they need to traverse mountains.
	MountaineersTradition = &Tradition{
		Name:             "mountaineers",
		PreferredEthoses: []*Ethos{Bellicose, Communal, Stoic},
		Type:             WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "pacifists":
					fallthrough
				case "mountain_skirmishing":
					fallthrough
				case "caucasian_wolves":
					return false
				}
			}
			return true
		},
	}
	// This culture places a strong emphasis on having a few, well-trained warriors. If you're not the best-of-the-best, you're not welcome to serve.
	OnlyTheStrongTradition = &Tradition{
		Name:             "only_the_strong",
		PreferredEthoses: []*Ethos{Bellicose, Stoic},
		Type:             WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return c.Ethos != Egalitarian
		},
	}
	// Border conflicts are common for rulers of this culture. Land often changes hands in unjust ways.
	QuarrelsomeTradition = &Tradition{
		Name:             "quarrelsome",
		PreferredEthoses: []*Ethos{Bellicose, Spiritual},
		Type:             WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "pacifists" {
					return false
				}
			}

			return true
		},
	}
	// Warriors from this culture are unyielding and unshakable.
	// They do not fall back - even in the face of overwhelming odds, for better or for worse.
	StandAndFightTradition = &Tradition{
		Name:             "stand_and_fight",
		PreferredEthoses: []*Ethos{Bellicose, Spiritual, Stoic},
		Type:             WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "defensive_tactics" {
					return false
				}
			}

			return true
		},
	}
	// This culture has mastered the art of fighting in very dry climates.
	WarriorsOfTheDryTradition = &Tradition{
		Name:             "warriors_of_the_dry",
		PreferredEthoses: []*Ethos{Bellicose, Communal, Stoic},
		Type:             WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "pacifists" {
					return false
				}
			}

			return true
		},
	}
	// This culture is used to the reality that harsh winds and bitter cold brings.
	WinterWarriorsTradition = &Tradition{
		Name:             "winter_warriors",
		PreferredEthoses: []*Ethos{Bellicose, Communal, Stoic},
		Type:             WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "pacifists" {
					return false
				}
			}

			return true
		},
	}
	// We will take your lands, your children, your traditions, and your future!
	MalleableInvadersTradition = &Tradition{
		Name:             "malleable_invaders",
		PreferredEthoses: []*Ethos{Bellicose, Bureaucratic, Egalitarian},
		Type:             WarfareTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "pacifists" {
					return false
				}
			}

			return true
		},
	}
)

var AllWarfareTraditions = []*Tradition{
	AdaptiveSkirmishersTradition,
	BattlefieldLootersTradition,
	FormationFightingExpertsTradition,
	FrugalArmorsmithsTradition,
	RecognitionOfTalentTradition,
	ReverenceForVeteransTradition,
	StalwartDefendersTradition,
	StrengthInNumbersTradition,
	SwordsForHireTradition,
	WarriorPriestsTradition,
	WarriorsByMeritTradition,
	ForestFightersTradition,
	HighlandWarriorsTradition,
	HitAndRunTacticiansTradition,
	HorseBreedersTradition,
	JungleWarriorsTradition,
	LongbowCompetitionsTradition,
	MountaineersTradition,
	OnlyTheStrongTradition,
	QuarrelsomeTradition,
	StandAndFightTradition,
	WarriorsOfTheDryTradition,
	WinterWarriorsTradition,
	MalleableInvadersTradition,
}

// Social Traditions

var (
	// The people of this culture excel in diplomatic matters, and social etiquette is valued above all else.
	CharismaticTradition = &Tradition{
		Name:              "charismatic",
		PreferredEthoses:  []*Ethos{Communal, Egalitarian, Stoic},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Regardless of faith, people of this culture are motivated to support those less fortunate.
	CharitableTradition = &Tradition{
		Name:              "charitable",
		PreferredEthoses:  []*Ethos{Communal, Spiritual},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// In this culture truth and impartial justice is valued highly.
	EquitableTradition = &Tradition{
		Name:              "equitable",
		PreferredEthoses:  []*Ethos{Egalitarian, Stoic},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// In this culture slights are not to be forgotten, nor forgiven. Its people can carry grudges for long, and vengeance is carried out with a passion.
	EyeForAnEyeTradition = &Tradition{
		Name:              "eye_for_an_eye",
		PreferredEthoses:  []*Ethos{Bellicose, Communal},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// In this culture being patient and restrained in the face of adversity is common.
	ForebearingTradition = &Tradition{
		Name:              "forebearing",
		PreferredEthoses:  []*Ethos{Spiritual, Stoic},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture teaches and enshrines ideas of industriousness and hard work on behalf of one's community.
	IndustriousTradition = &Tradition{
		Name:              "industrious",
		PreferredEthoses:  []*Ethos{Bellicose, Communal, Stoic},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture would chuckle on the chopping block with an axe above its head.
	LifeIsJustAJokeTradition = &Tradition{
		Name:              "life_is_just_a_joke",
		PreferredEthoses:  []*Ethos{Communal, Stoic},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// For this culture, serving one's liege and country is both noble and just - a duty and a privilege, rather than an avaricious arrangement or an unwanted burden.
	LoyalSubjectsTradition = &Tradition{
		Name:              "loyal_subjects",
		PreferredEthoses:  []*Ethos{Bellicose, Communal, Stoic},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// The holy people of this culture believe strongly that the faith should be ministered in the field, and their most devout frequently wander the world.
	MendicantMysticsTradition = &Tradition{
		Name:              "mendicant_mystics",
		PreferredEthoses:  []*Ethos{Communal, Spiritual},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture values modesty, one should not take up too much space or think oneself better than others.
	ModestTradition = &Tradition{
		Name:              "modest",
		PreferredEthoses:  []*Ethos{Stoic, Spiritual},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// The people of this culture are particularly zealous and dedicated to their faith.
	StrongBelieversTradition = &Tradition{
		Name:              "strong_believers",
		PreferredEthoses:  []*Ethos{Communal, Spiritual},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture has embraced chivalry and chivalric codes of conduct as a social method of regulating behavior.
	// Martial prowess, duty, honor and morality are prized, as is bad poetry and romantic literature.
	ChivalryTradition = &Tradition{
		Name:              "chivalry",
		PreferredEthoses:  []*Ethos{Bellicose, Courtly, Egalitarian},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Anyone can fish, but to do so with such skill that any catch is bountiful, regardless of tide or wind, is a rarer talent.
	// No one knows the coast, land or sea, better than these fishermen.
	DexterousFishermenTradition = &Tradition{
		Name:              "dexterous_fishermen",
		PreferredEthoses:  []*Ethos{Bureaucratic, Communal, Courtly, Egalitarian, Spiritual, Stoic},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Those who are willing to take up sword and fight for their culture are worthy of admiration. No matter the odds.
	MartialAdmirationTradition = &Tradition{
		Name:             "martial_admiration",
		PreferredEthoses: []*Ethos{Bellicose, Stoic},
		Type:             SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "chanson_de_geste":
					fallthrough
				case "futuwaa":
					fallthrough
				case "pacifists":
					fallthrough
				case "immortals":
					fallthrough
				case "druzhina":
					return false
				}
			}

			return true
		},
	}
	// For this culture, the use of force can never be justified. Only by pursuing a path of non-violence can people truly live in peace.
	PacifistsTradition = &Tradition{
		Name:             "pacifists",
		PreferredEthoses: []*Ethos{Egalitarian, Spiritual},
		Type:             SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			return c.Ethos != Bellicose
		},
	}
	// This culture promotes thought and self-reflection.
	PhilosopherCultureTradition = &Tradition{
		Name:             "philosopher_culture",
		PreferredEthoses: []*Ethos{Communal, Courtly, Spiritual},
		Type:             SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			if c.Ethos == Bellicose {
				return false
			}
			for _, t := range c.Traditions {
				if t.Name == "warrior_culture" {
					return false
				}
			}

			return true
		},
	}
	// Storming a port, setting it ablaze, and taking everything that isn't nailed down may not be noble, but it is profitable.
	PracticedPiratesTradition = &Tradition{
		Name:              "practiced_pirates",
		PreferredEthoses:  []*Ethos{Bellicose},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Driven by economic necessity, or perhaps social expectation, this culture has refined their hunting practices to precise and almost beautiful art.
	ProlificHuntersTradition = &Tradition{
		Name:             "prolific_hunters",
		PreferredEthoses: []*Ethos{Bellicose, Spiritual, Stoic},
		Type:             SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "vegetarians":
					fallthrough
				case "pacifists":
					return false
				}
			}

			return true
		},
	}
	// Poetry is considered a noble art in this culture, and many spend their time piecing words together with meaning and thought.
	RefinedPoetryTradition = &Tradition{
		Name:             "refined_poetry",
		PreferredEthoses: []*Ethos{Communal, Courtly, Spiritual},
		Type:             SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "northern_stories" {
					return false
				}
			}

			return true
		},
	}
	// For this culture, the call of the sea is too strong to resist, and
	// they live to sail like a dream on a crystal clear ocean, or ride on the crest of a wild raging storm.
	SeafarersTradition = &Tradition{
		Name:             "seafarers",
		PreferredEthoses: []*Ethos{Bellicose, Bureaucratic, Spiritual},
		Type:             SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "coastal_warriors" {
					return false
				}
			}

			return true
		},
	}
	// This culture doesn't prize the trappings of power - the ceremony, the wealth, the pointless decoration.
	// This culture is more interested in power itself, preferably in the form of towering walls and keeps.
	SpartanTradition = &Tradition{
		Name:              "spartan",
		PreferredEthoses:  []*Ethos{Communal, Stoic},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture has a strong and persistent oral tradition.
	// The past is preserved through ritualistic storytelling, where the heroes and legends of the past are passed down through generations.
	StorytellersTradition = &Tradition{
		Name:             "storytellers",
		PreferredEthoses: []*Ethos{Communal, Courtly, Stoic},
		Type:             SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "northern_stories" {
					return false
				}
			}

			return true
		},
	}
	// his culture values martial prowess and strength above everything else.
	// Children are brought up knowing how to fight, and are discouraged from scholarly pursuits.
	// Weakness is not tolerated.
	WarriorCultureTradition = &Tradition{
		Name:             "warrior_culture",
		PreferredEthoses: []*Ethos{Bellicose, Spiritual},
		Type:             SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "coastal_warriors":
					fallthrough
				case "pacifists":
					return false
				}
			}

			return true
		},
	}
	// This culture embraces everyone and is genuinely fascinated by all cultures.
	XenophilicTradition = &Tradition{
		Name:              "xenophilic",
		PreferredEthoses:  []*Ethos{Communal, Egalitarian},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture has developed a strong aversion to consuming the flesh of animals, and practices vegetarianism throughout their society.
	VegetariansTradition = &Tradition{
		Name:             "vegetarians",
		PreferredEthoses: []*Ethos{Communal, Spiritual, Stoic},
		Type:             SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "sacred_hunts":
					fallthrough
				case "prolific_hunters":
					return false
				}
			}
			return true
		},
	}
	// People of this culture have lost the homeland they once held dear.
	// Some among them travel the world searching for a place where they may begin anew.
	DiasporicTradition = &Tradition{
		Name:              "diasporic",
		PreferredEthoses:  AllEthoses,
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Life is lived, told, and remembered through the lens of tales.
	// What we say, and what people say about us, matters as much as what we do.
	// We remember that, and it shows in our stories.
	NorthernStoriesTradition = &Tradition{
		Name:             "northern_stories",
		PreferredEthoses: []*Ethos{Bellicose, Bureaucratic},
		Type:             SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "refined_poetry":
					fallthrough
				case "storytellers":
					return false
				}
			}

			return true
		},
	}
	// Honor is born, yes, but honor is also earned and lost through deeds.
	// By pursuing and avenging slights, or failing to, soldiers fade to scum while whelps grow into warriors.
	PerformativeHonorTradition = &Tradition{
		Name:             "performative_honor",
		PreferredEthoses: []*Ethos{Bellicose},
		Type:             SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "trials_by_combat" {
					return false
				}
			}

			return true
		},
	}
	// Though not all are born fighters, with the right mettle and a strong sword-arm, any can show themselves a warrior at heart and so become one in deed.
	TheRightToProveTradition = &Tradition{
		Name:             "the_right_to_prove",
		PreferredEthoses: []*Ethos{Bellicose, Communal, Egalitarian},
		Type:             SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "performative_honor" {
					return false
				}
			}

			return true
		},
	}
	// Words are a coward's substitute for weapons.
	// Why leave justice to decrepit laws and corrupt magistrates when disputes can be settled faster and fairer with the blade?
	TrialsByCombatTradition = &Tradition{
		Name:             "trials_by_combat",
		PreferredEthoses: []*Ethos{Bellicose, Courtly, Stoic},
		Type:             SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "performative_honor" {
					return false
				}
			}
			return true
		},
	}
	// Children in this culture often carry on the profession of their parents, accumulating extensive knowledge and skill for their trade across generations.
	ExpertArtisansTradition = &Tradition{
		Name:              "expert_artisans",
		PreferredEthoses:  []*Ethos{Communal, Courtly, Spiritual},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture treats music almost as the language of the divine, and thus many individuals take up the noble and
	// celebrated pursuit of musical study from a young age.
	MusicalTheoristsTradition = &Tradition{
		Name:              "musical_theorists",
		PreferredEthoses:  []*Ethos{Communal, Courtly, Spiritual},
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Blood may be thicker than water, but wine makes for better living.
	RitualizedFriendshipTradition = &Tradition{
		Name:              "ritualized_friendship",
		PreferredEthoses:  AllEthoses,
		Type:              SocialTraditionType,
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
)

var AllSocialTraditions = []*Tradition{
	CharismaticTradition,
	CharitableTradition,
	EquitableTradition,
	EyeForAnEyeTradition,
	ForebearingTradition,
	IndustriousTradition,
	LifeIsJustAJokeTradition,
	LoyalSubjectsTradition,
	MendicantMysticsTradition,
	ModestTradition,
	StrongBelieversTradition,
	ChivalryTradition,
	DexterousFishermenTradition,
	MartialAdmirationTradition,
	PacifistsTradition,
	PhilosopherCultureTradition,
	PracticedPiratesTradition,
	ProlificHuntersTradition,
	RefinedPoetryTradition,
	SeafarersTradition,
	SpartanTradition,
	StorytellersTradition,
	WarriorCultureTradition,
	XenophilicTradition,
	VegetariansTradition,
	DiasporicTradition,
	NorthernStoriesTradition,
	PerformativeHonorTradition,
	TheRightToProveTradition,
	TrialsByCombatTradition,
	ExpertArtisansTradition,
	MusicalTheoristsTradition,
	RitualizedFriendshipTradition,
}

// Ritual Traditions

var (
	// n this culture being of a different faith means that you're not eligible for succession.
	BoundByFaithTradition = &Tradition{
		Name:              "bound_by_faith",
		Type:              RitualTraditionType,
		PreferredEthoses:  []*Ethos{Spiritual},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// The frequency and exuberance with which this culture throws festivities would be considered vulgar to others, but a party is a party!
	FrequentFestivitiesTradition = &Tradition{
		Name:              "frequent_festivities",
		Type:              RitualTraditionType,
		PreferredEthoses:  []*Ethos{Communal, Courtly, Stoic},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// To learn the words of a neighbor brings you one step closer to the language of the Divine.
	LinguistsTradition = &Tradition{
		Name:              "linguists",
		Type:              RitualTraditionType,
		PreferredEthoses:  []*Ethos{Bureaucratic, Egalitarian, Spiritual},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture long ago developed knowledge of the medicinal properties of the plants and trees;
	// to them most ailments are treatable with the right poultice, salve or stew.
	MedicinalHerbalistsTradition = &Tradition{
		Name:              "medicinal_herbalists",
		Type:              RitualTraditionType,
		PreferredEthoses:  []*Ethos{Bureaucratic, Stoic},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Members of this culture believe that the best missionary is one carrying a sword.
	// While support for holy wars are widespread, motives are scrutinized as to make sure the Divine powers would approve.
	ByTheSwordTradition = &Tradition{
		Name:              "by_the_sword",
		Type:              RitualTraditionType,
		PreferredEthoses:  []*Ethos{Spiritual},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Rulers of this culture often keep a number of concubines in their household, regardless of their faith.
	ConcubinesTradition = &Tradition{
		Name:             "concubines",
		Type:             RitualTraditionType,
		PreferredEthoses: AllEthoses,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "monogamous":
					fallthrough
				case "polygamous":
					return false
				}
			}

			return true
		},
	}
	// Food isn't just sustenance for this culture; it is both an art and a ritual, a focal point for family and community alike.
	CulinaryArtistsTradition = &Tradition{
		Name:              "culinary_artists",
		Type:              RitualTraditionType,
		PreferredEthoses:  []*Ethos{Communal, Courtly, Spiritual},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// No matter how far beyond forgiveness one has gone, this culture thinks it is better to gouge out an eye than run through a heart.
	MercifulBlindingsTradition = &Tradition{
		Name:              "merciful_blindings",
		Type:              RitualTraditionType,
		PreferredEthoses:  []*Ethos{Courtly, Spiritual},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// In this culture one keeps a single spouse, regardless of faith.
	MonogamousTradition = &Tradition{
		Name:             "monogamous",
		Type:             RitualTraditionType,
		PreferredEthoses: AllEthoses,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "concubines":
					fallthrough
				case "polygamous":
					return false
				}
			}

			return true
		},
	}
	// In this culture one tend to have multiple spouses, regardless of faith.
	PolygamousTradition = &Tradition{
		Name:             "polygamous",
		Type:             RitualTraditionType,
		PreferredEthoses: AllEthoses,
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "concubines":
					fallthrough
				case "monogamous":
					return false
				}
			}

			return true
		},
	}
	// Members of this culture often live together with those sharing different faiths and beliefs,
	// and do well in adopting foreign elements into their own worship.
	ReligionBlendingTradition = &Tradition{
		Name:             "religion_blending",
		Type:             RitualTraditionType,
		PreferredEthoses: []*Ethos{Communal, Courtly, Spiritual},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "steppe_tolerance" {
					return false
				}
			}

			return true
		},
	}
	// For the salvation of the soul, one must commit to works of religious significance.
	// You cannot purchase redemption, but having a temple of priests sing your praises helps!
	ReligiousPatronageTradition = &Tradition{
		Name:              "religious_patronage",
		Type:              RitualTraditionType,
		PreferredEthoses:  []*Ethos{Communal, Courtly, Spiritual},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Death, conquest, the acquisition of rank — are these not the things that define life in the nobility?
	// Where our people go, so go the markers of our lives' worth.
	RunestoneRaisersTradition = &Tradition{
		Name:              "runestone_raisers",
		Type:              RitualTraditionType,
		PreferredEthoses:  []*Ethos{Bellicose, Bureaucratic},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Trees are considered sacred by this culture. The forest is a living being that ought to be cared for—not destroyed.
	SacredGrovesTradition = &Tradition{
		Name:              "sacred_groves",
		Type:              RitualTraditionType,
		PreferredEthoses:  []*Ethos{Spiritual},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Hunting is neither a sport nor a pastime,
	// it is nothing less than the purest expression of the human experience,
	// where one may connect with the divine by demonstrating their mastery over nature.
	SacredHuntsTradition = &Tradition{
		Name:             "sacred_hunts",
		Type:             RitualTraditionType,
		PreferredEthoses: []*Ethos{Bellicose, Spiritual, Stoic},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "vegetarians":
					fallthrough
				case "pacifists":
					return false
				}
			}

			return true
		},
	}
	// This culture considers mountains to be connections to the divine, and treats them with reverence and respect.
	SacredMountainsTradition = &Tradition{
		Name:              "sacred_mountains",
		Type:              RitualTraditionType,
		PreferredEthoses:  []*Ethos{Spiritual},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// There are as many alchemists in this culture attempting to turn metal into gold as there are smiths forging weapons.
	// What does it matter if this culture's wealth and weaponry are enhanced by witchcraft?
	SorcerousMetallurgyTradition = &Tradition{
		Name:              "sorcerous_metallurgy",
		Type:              RitualTraditionType,
		PreferredEthoses:  []*Ethos{Communal, Spiritual},
		CompatibilityFunc: func(c *Culture) bool { return true }, // Lead to accepted witchcraft
	}
	// Any commander can gamble and get lucky on the day of battle.
	// A real strategist keeps their mind honed through practice and pretend, so that they have no need of 'luck'.
	TabletopWarriorsTradition = &Tradition{
		Name:              "tabletop_warriors",
		Type:              RitualTraditionType,
		PreferredEthoses:  []*Ethos{Bellicose, Courtly},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Lineage is very important to this culture,
	// to the point where ancestors have become mythical and legendary beings with many who claim to be their descendants.
	MysticalAncestorsTradition = &Tradition{
		Name:              "mystical_ancestors",
		Type:              RitualTraditionType,
		PreferredEthoses:  []*Ethos{Spiritual},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
)

var AllRitualTraditions = []*Tradition{
	BoundByFaithTradition,
	FrequentFestivitiesTradition,
	LinguistsTradition,
	MedicinalHerbalistsTradition,
	ByTheSwordTradition,
	ConcubinesTradition,
	CulinaryArtistsTradition,
	MercifulBlindingsTradition,
	MonogamousTradition,
	PolygamousTradition,
	ReligionBlendingTradition,
	ReligiousPatronageTradition,
	RunestoneRaisersTradition,
	SacredGrovesTradition,
	SacredHuntsTradition,
	SacredMountainsTradition,
	SorcerousMetallurgyTradition,
	TabletopWarriorsTradition,
	MysticalAncestorsTradition,
}

// Regional Traditions

var (
	// Though some would decry it as weakness, a realm founded on understanding and tolerance is a realm built on sturdy foundations.
	AfricanToleranceTradition = &Tradition{
		Name:             "african_tolerance",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Communal, Egalitarian},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "steppe_tolerance" {
					return false
				}
			}

			return true
		},
	}
	// The bush is a rough and rugged place, but it offers a bounty of opportunity for those who know where to look and how to exploit it.
	// Training our archers in these techniques will give them an advantage when engaging enemies in these areas.
	BushHuntingTradition = &Tradition{
		Name:             "bush_hunting",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Communal, Egalitarian},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "jungle_warriors" {
					return false
				}
			}

			return true
		},
	}
	// Having spent centuries making themselves a home around the Caucasian gates, this culture has grown well at home with everything that mountain warfare entails.
	CaucasianWolvesTradition = &Tradition{
		Name:             "caucasian_wolves",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Bellicose, Communal, Stoic},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "mountain_skirmishing":
					fallthrough
				case "mountaineers":
					return false
				}
			}

			return true
		},
	}
	// This culture values the heroic deeds of long gone ancestors.
	// The romantic retelling of the lives of knights such as Guillaume, or Roland, will inspire generations to come.
	ChansonDeGesteTradition = &Tradition{
		Name:             "chanson_de_geste",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Bellicose, Courtly, Stoic},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "druzhina":
					fallthrough
				case "martial_admiration":
					fallthrough
				case "immortals":
					fallthrough
				case "futuwaa":
					return false
				}
			}

			return true
		},
	}
	// Our people have been exiled and driven out time and time again.
	// While tragic, these experiences have led us to become exceptionally skilled in rear-guard tactics which minimize casualties.
	DefensiveTacticsTradition = &Tradition{
		Name:             "defensive_tactics",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Communal},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "stand_and_fight" {
					return false
				}
			}

			return true
		},
	}
	// This culture has a tradition of fortified desert retreats, where the pious can live in connection with the open sky and the sands.
	// This lifestyle breeds able and pious desert warriors and attracts mystics.
	DesertRibatsTradition = &Tradition{
		Name:             "desert_ribats",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Bellicose, Spiritual, Stoic},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "pacifists" {
					return false
				}
			}

			return true
		},
	}
	// This culture is no stranger to the harsh environment of the desert and has mastered the use of camels to aid them in travel, warfare, and everyday life.
	DesertTravelersTradition = &Tradition{
		Name:              "desert_travelers",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Spiritual, Stoic},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Varangian soldiers have long served as imposing bodyguards for our monarchs, but they can just as easily serve as dedicated heavy infantry in our armies.
	DruzhinaTradition = &Tradition{
		Name:             "druzhina",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Bellicose, Stoic},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "chanson_de_geste":
					fallthrough
				case "martial_admiration":
					fallthrough
				case "immortals":
					fallthrough
				case "futuwaa":
					return false
				}
			}

			return true
		},
	}
	// Our ancestors have lived in these forests for generations, but now they are under threat by outsiders.
	// Centuries of accumulated experience fighting in forests will aid us in protecting our ancestral homeland.
	ForestWardensTradition = &Tradition{
		Name:             "forest_wardens",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Bellicose, Communal, Stoic},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "forest_folk" {
					return false
				}
			}

			return true
		},
	}
	// Our youth have begun forming clubs that promote prowess, vigor, and moral behavior.
	// By endorsing and supporting these clubs, we will ensure a supply of able-bodied soldiers we can rally to our cause.
	FutuwaaTradition = &Tradition{
		Name:             "futuwaa",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Bellicose, Communal, Courtly},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "chanson_de_geste":
					fallthrough
				case "martial_admiration":
					fallthrough
				case "immortals":
					fallthrough
				case "druzhina":
					return false
				}
			}

			return true
		},
	}
	// Distinguished warriors in this culture are designated as Garudas, and are expected to fight until death.
	// Even those who follow a Garuda are expected to die if their leader does, either in battle or by their own hand.
	GarudaWarriorsTradition = &Tradition{
		Name:              "garuda_warriors",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Bellicose, Stoic},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture has made its home atop the 'roof of the world', where the air is fresh but cold, and the winters long and hard.
	// They pride themselves that few of the cultures below would be able to live as well as this people does, among the peaks and plateaus where otherwise only sheep can thrive.
	HimalayanSettlersTradition = &Tradition{
		Name:              "himalayan_settlers",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Communal, Spiritual},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture has mastered the horse and its use in warfare.
	HorseLordsTradition = &Tradition{
		Name:              "horse_lords",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Bellicose, Communal},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Immortals was an elite heavy infantry unit of 10,000 soldiers in the army of the Achaemenid Empire.
	// The unit served in a dual capacity through its role as imperial guard alongside
	// its contribution to the ranks of the Persian Empire's standing army.
	ImmortalsTradition = &Tradition{
		Name:             "immortals",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Bellicose, Courtly},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "chanson_de_geste":
					fallthrough
				case "martial_admiration":
					fallthrough
				case "druzhina":
					fallthrough
				case "futuwaa":
					return false
				}
			}

			return true
		},
	}
	// This culture has embraced the martial art and ritual worship of two-handed broadswords letting them foster strong and able warriors.
	// The heavy swords allow them to mow down infantry and cavalry alike.
	KhadgaPujaTradition = &Tradition{
		Name:              "khadga_puja",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Bellicose, Stoic},
		CompatibilityFunc: func(c *Culture) bool { return c.Ethos != Egalitarian },
	}
	// Light cavalry which can strike quickly at exposed enemies before darting back away, our Konni can be formed into regiments which specialize in harassment and raiding.
	KonniRaidsTradition = &Tradition{
		Name:             "konni_raids",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Stoic},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "hit_and_run_tacticians" {
					return false
				}
			}

			return true
		},
	}
	// This culture venerates elephants, likening them to royalty.
	// They are masterful at capturing them, training them, and using them for devastating effect in war.
	LordsOfTheElephantTradition = &Tradition{
		Name:              "lords_of_the_elephant",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Bellicose, Courtly, Stoic},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture has made the cooler mountains their homes in an otherwise dry and arid region.
	// The herders and farmers of these uplands have make for resilient and robust warriors,
	// and over time their toil has made this people flourish far above the naked desert.
	MountainHerdingTradition = &Tradition{
		Name:             "mountain_herding",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Bellicose, Communal, Stoic},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "pacifists" {
					return false
				}
			}

			return true
		},
	}
	// The great mountains in the Horn of Africa pose logistical challenges to many types of armies,
	// but our skirmishers can easily adapt to these conditions to rout enemy invaders.
	MountainSkirmishingTradition = &Tradition{
		Name:             "mountain_skirmishing",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Communal, Spiritual},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "caucasian_wolves":
					fallthrough
				case "mountaineers":
					return false
				}
			}

			return true
		},
	}
	// From their origins in the hills of the Bohemian Forest, to the crags of the Ore and Tatra mountains,
	// most in this culture know not only how to traverse the rocky lands, but handle life on a mountain with ease.
	MountaineerRuralismTradition = &Tradition{
		Name:              "mountaineer_ruralism",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Stoic},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Modeled after the legendary champions of the Rashidun army, our Mubarizun soldiers are trained to excel in both formation fighting as well as single combat.
	MubarizunTradition = &Tradition{
		Name:             "mubarizun",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Bellicose, Communal, Stoic},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "pacifists":
					fallthrough
				case "warriors_of_the_dry":
					return false
				}
			}

			return true
		},
	}
	// We have a long tradition of a standing army made to protect the capital.
	// This has given us access to tough soldiers that work exceedingly well with supporting troops, such as War Elephants.
	RoyalArmyTradition = &Tradition{
		Name:              "royal_army",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Spiritual},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture is intimately accustomed with the Saharan desert, which it has traversed for centuries, connecting the Mediterranean with the sub-Saharan kingdoms of the Sahel.
	SaharanNomadsTradition = &Tradition{
		Name:              "saharan_nomads",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Spiritual, Stoic},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// Those who live on the Steppe will always have much in common with those who share their way of life.
	// Strangers may worship differently, but they still live in the saddle.
	SteppeToleranceTradition = &Tradition{
		Name:             "steppe_tolerance",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Bellicose, Communal, Egalitarian},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "african_tolerance":
					fallthrough
				case "religion_blending":
					return false
				}
			}

			return true
		},
	}
	// This culture has a strong tradition of kinship within the extended family, or clan. In times of need a ruler will always be able to rely on their kin.
	StrongKinshipTradition = &Tradition{
		Name:              "strong_kinship",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Bellicose, Bureaucratic, Stoic},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// This culture has a long history of holding assemblies comprised of
	// the land's ruling potentates, debating politics, law, finance, and, most importantly, how best to... advise the ruler.
	TheWitenagemotTradition = &Tradition{
		Name:              "the_witenagemot",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Bureaucratic, Stoic},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// In Scandinavia, the distances and wilderness render centralized authority ineffective.
	// When disputes must be resolved and decisions made, the Ting-meet is gathered, an assembly of peers presided over by a lawspeaker.
	TingMeetTradition = &Tradition{
		Name:              "ting_meet",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Bellicose, Bureaucratic},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// The hills of the wooded savanna have been our homeland for generations.
	// Our ancestors showed us how to utilize this terrain, and with raiders threatening our borders, we can use this knowledge in our defense.
	UplandSkirmishingTradition = &Tradition{
		Name:             "upland_skirmishing",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Communal, Egalitarian},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				if t.Name == "highland_warriors" {
					return false
				}
			}

			return true
		},
	}
	// Although many of the older traditions of the Visigoths have long been wiped from the world, the children of the Pyrenees remember the ancient ways,
	// and how land was divided between sons and daughters practically but fairly.
	VisigothicCodesTradition = &Tradition{
		Name:              "visigothic_codes",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Egalitarian},
		CompatibilityFunc: func(c *Culture) bool { return true }, // @TODO produce equality
	}
	// Those who live in East Africa well remember the great Warrior Queens who once stood against invaders and forged kingdoms on the Kush.
	// Our neighbors may insist on the superiority of males, but we track our bloodlines through our mothers,
	// and the warcry of powerful Nubian women send these foreign "men" running home.
	WarriorQueensTradition = &Tradition{
		Name:              "warrior_queens",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Egalitarian},
		CompatibilityFunc: func(c *Culture) bool { return true }, // @TODO produce equality
	}
	// This culture is well-integrated into the commercial practices of the region, and its people are strongly associated with the caravan trade.
	CaravaneersTradition = &Tradition{
		Name:              "caravaneers",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Egalitarian},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// In a land of plenty, enriched by the Nile and their own ingenuity, the Nubian have prospered for generations.
	// For almost as long, the bow has been the weapon of choice to deter the raiders and brigands of the surrounding deserts.
	LandOfTheBowTradition = &Tradition{
		Name:              "land_of_the_bow",
		Type:              RegionalTraditionType,
		PreferredEthoses:  []*Ethos{Bellicose, Spiritual},
		CompatibilityFunc: func(c *Culture) bool { return true },
	}
	// The tradition of keeping armed retinues in service to a household has served us well.
	// By formalizing and expanding this system we can field entire regiments of huscarls in our armies.
	HirdsTradition = &Tradition{
		Name:             "hirds",
		Type:             RegionalTraditionType,
		PreferredEthoses: []*Ethos{Bellicose, Spiritual},
		CompatibilityFunc: func(c *Culture) bool {
			for _, t := range c.Traditions {
				switch t.Name {
				case "warrior_culture":
					fallthrough
				case "pacifists":
					return false
				}
			}

			return true
		},
	}
)

var AllRegionalTraditions = []*Tradition{
	AfricanToleranceTradition,
	BushHuntingTradition,
	CaucasianWolvesTradition,
	ChansonDeGesteTradition,
	DefensiveTacticsTradition,
	DesertRibatsTradition,
	DesertTravelersTradition,
	DruzhinaTradition,
	ForestWardensTradition,
	FutuwaaTradition,
	GarudaWarriorsTradition,
	HimalayanSettlersTradition,
	HorseLordsTradition,
	ImmortalsTradition,
	KhadgaPujaTradition,
	KonniRaidsTradition,
	LordsOfTheElephantTradition,
	MountainHerdingTradition,
	MountainSkirmishingTradition,
	MountaineerRuralismTradition,
	MubarizunTradition,
	RoyalArmyTradition,
	SaharanNomadsTradition,
	SteppeToleranceTradition,
	StrongKinshipTradition,
	TheWitenagemotTradition,
	TingMeetTradition,
	UplandSkirmishingTradition,
	VisigothicCodesTradition,
	WarriorQueensTradition,
	CaravaneersTradition,
	LandOfTheBowTradition,
	HirdsTradition,
}

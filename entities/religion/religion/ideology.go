package religion

type Ideology struct {
	religion *Religion
}

func NewIdeology(r *Religion) *Ideology {
	i := &Ideology{religion: r}

	return i
}

/*

	HighGoal:
		TransformingIntoLikenessOfGod (only for monotheism)
		AttainHeavenAndResidesThereWithGod (only for monotheism)
		BringManToCompletenessAndToCloseRelationshipWithGod (only for monotheism)
		BringHolinessDownToTheWorld
		StopReincarnation (must have reincarnation)
		NeverStopReincarnation (must have reincarnation)
		ProduceChildren
		InvestigateMyself
		GetMaxPleasure
		SpreadReligion
		LovePeople
		BecomePerfectAndSaints
		FightWithEvil
		FightForEvil



	Doctrines
		God:
			GodIsGood bool
			GodIsEvil bool
			GodSourceOfMoralLaw bool
			Gnosticism bool
			Afterlife
				DoesExist bool
				ForAll (Good, Bad, Depends)
				ForBelievers (Good, Bad, Depends)
				ForTopBelievers (Good, Bad, Depends)
				ForUntrueBelievers (Good, Bad, Depends)
				ForAtheists (Good, Bad, Depends)
		Human:
			Purity bool
			LiveIsSacred bool
			SanctityOfNature           bool
			SacredChildbirth           bool
			Karma bool
			PeopleHaveSoul bool
			Polyamory                  bool
			Asceticism                 bool
			BadThingForGoodPurpose bool
			HumanNatureIsEvil bool
			Raider                     bool // The strong should take from the weak
			Hedonism                   bool
			CanBeSaint bool





	Rules:
		DressCode           bool
		BeHumble           bool
		LearnKeyScriptures           bool
		LiveUnderGround            bool
		PrayWithFrequency bool
		TaxNonbelievers            bool


	teology:
		Messiah                    bool
		Prophets                   bool
		Legalism                   boolLegalism                   bool
		ReligiousLaw               bool
		AncestorWorship            bool
		Reincarnation              bool
		SanctionedFalseConversions bool // If your life is threatened, it is acceptable to confess a false faith, as long as you keep the truth in your heart.
		SunWorship                 bool
		MoonWorship                bool
		TreeConnection             bool // Trees are the essence of life, and we must be near them.
		AnimalConnection           bool
		Blindsight                 bool // Only the blind can perceive true reality
		Astrology                  bool

	clerics:
		MendicantPreachers         bool
		Monasticism                bool

	rituals:
		RitualHospitality          bool
*/

// temples, shelters, holy scripture, symbolics, initiation
// religious art, religious literature, reliquary, sacred places
// preachers, forced conversion, masters, missionary outreach

// monotheism

// rich ideology: cult property, indulgences, tithe, primacy of the pope, tax to support the poor
// poor ideology: holiness, celibacy, flagellantism

// developed ideology: seminaries, libraries, feed the world, patronship
// undeveloped ideology: secret sect, sacrificial offesring, rites and chants

// militant ideology: holy army, warrior monks, defenders of Faith
// peaceful ideology: non-violent resistance, swords to plowshares, no more killing, world religion

// primitive ideology: great fast,, repentance, martyrdom
// organized ideology: confession, commandments, great afterlife, devotional code

// divine law, sacrament, justice of heaven

// polytheism

// rich ideology: buying your sins, inviolacy of temples, god of king cult, cult of justice, god of prosperity and wealth, sacred stone, patron saint of merchants
// poor ideology: oracles, diviners, libations, holy animals, all father's watchmen, ordeals,
// ... mother goddess

// developed ideology: seminaries, libraries, schools of philosophers, kalokagathos, galdbarok, runes, volhvs, periapts
// undeveloped ideology: hecatombs, pharmakos, inferiae, god of erebus, trickster, samhain

// militant ideology: god of war cult, gymnasium schools, holy army. marching hymns, agoge, berserkers, heavenly palace
// peaceful ideology: temple healers, sport contest, oath of healer, veneration of the ancestors, blood oath, truce

// primitive ideology: symposium, music, theatre,, liturgical drama, mead of poetry, visas, sages
// organized ideology: theogony, scapegoat, psychopomp, spoken word, norns, tapestry, prophesy, honorable death

// , deed of expiation, mysteries, nithing

// deism

// rich ideology: emperor's doctrine
// poor ideology: animal messengers

// developed ideology: seminaries, calligraphy, sanctuary of mind
// undeveloped ideology: phantoms, living pillars, demonic deal

// militant ideology: armor of god, warrior's path
// peaceful ideology:

// primitive ideology:
// organized ideology: calendar of feasts, prayer spells, spirits of death

// mono no aware, cleansing rituals

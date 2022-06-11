package religion

type Attributes struct {
	Temples      *Temples
	Symbolics    *Symbolics
	SacredPlaces *SacredPlaces
	Education    *Education
	Conversion   *Conversion
	Ideology     *Ideology
	Acts         *Acts
}

type Temples struct {
	CultProperty       bool // monotheism
	InviolacyOfTemples bool // monotheism
}

type Symbolics struct {
	Reliquary *Reliquary
	Art       *Art
}

type Reliquary struct {
	SacredStone bool // polytheism
}

type Art struct{}

type SacredPlaces struct{}

type Education struct {
	HasSeminaries         bool // monotheism, polytheism
	HasLibraries          bool // monotheism, polytheism
	Patronship            bool // monotheism
	SchoolsOfPhilosophers bool // polytheism
}

type Conversion struct {
	Preachers          bool
	ForcedConversion   bool
	Masters            bool
	MissionaryOutreach bool
}

type Ideology struct {
	Indulgences         bool // monotheism
	BuyingYourSins      bool // polytheism
	Tithe               bool // monotheism
	TaxToSupportThePoor bool // monotheism

	Holiness      bool // monotheism
	Celibacy      bool // monotheism
	Flagellantism bool // monotheism

	FeedTheWorld bool // monotheism
	SecretSect   bool

	HolyArmy             bool
	WarriorMonks         bool // monotheism
	DefendersAfFaith     bool
	NonViolentResistance bool
	Martyrdom            bool // monotheism
	Commandments         bool // to holy scripture
	DevotionalCode       bool // to holy scripture
	GreatAfterlife       bool
	DivineLaw            bool

	GodOfKingCult            bool // polytheism
	CultOfJustice            bool // polytheism
	GodOfProsperityAndWealth bool // polytheism
	PatronSaintOfMerchants   bool // polytheism
	GodOfErebus              bool // polytheism

	Oracles  bool // to clerics
	Diviners bool // to clerics
	Volhvs   bool // to clerics

	HolyAnimals  bool // polytheism
	Ordeals      bool
	KaloKagathos bool // an ideal of gentlemanly personal conduct, especially in a military context
	MagicBooks   bool // polytheism
	Runes        bool // to holy scripture
	Amulets      bool // polytheism
	Trickster    bool // polytheism
}

type Acts struct {
	RitesAndChants   bool
	GreatFast        bool
	Repentance       bool // monotheism
	Confession       bool // monotheism
	MassiveSacrifice bool // polytheism
	Pharmakos        bool // polytheism
	Inferiae         bool // polytheism
	Samhain          bool // polytheism to celebrations
}

// polytheism

// developed ideology: seminaries, libraries, schools of philosophers, kalokagathos, galdrabok, runes, volhvs, periapts
// undeveloped ideology: hecatombs, pharmakos, inferiae, god of erebus, trickster, samhain

// militant ideology: god of war cult, gymnasium schools, holy army. marching hymns, agoge, berserkers, heavenly palace
// peaceful ideology: temple healers, sport contest, oath of healer, veneration of the ancestors, blood oath, truce

// primitive ideology: symposium, music, theatre,, liturgical drama, mead of poetry, visas, sages
// organized ideology: theogony, scapegoat, psychopomp, spoken word, norns, tapestry, prophesy, honorable death

// mother goddess, deed of expiation, mysteries, nithing

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

package religion

type Doctrines struct {
	Main *MainDoctrine

	FullTolerance              bool
	Messiah                    bool
	Prophets                   bool
	Asceticism                 bool
	Astrology                  bool
	Esotericism                bool
	Legalism                   bool
	MendicantPreachers         bool
	Monasticism                bool
	Polyamory                  bool
	ReligiousLaw               bool
	SanctityOfNature           bool
	TaxNonbelievers            bool
	UnrelentingFaith           bool
	AncestorWorship            bool
	Pacifism                   bool
	Reincarnation              bool
	RitualHospitality          bool
	SacredChildbirth           bool
	SanctionedFalseConversions bool // If your life is threatened, it is acceptable to confess a false faith, as long as you keep the truth in your heart.
	SunWorship                 bool
	MoonWorship                bool
	PainIsVirtue               bool
	Darkness                   bool
	LiveUnderGround            bool
	TreeConnection             bool // Trees are the essence of life, and we must be near them.
	AnimalConnection           bool
	Blindsight                 bool // Only the blind can perceive true reality
	Raider                     bool // The strong should take from the weak
	Proselytizer               bool // It is our duty to spread our beliefs
	Hedonism                   bool
}

type MainDoctrine struct {
	Monotheism   bool // 60%
	Polytheism   bool // 60%
	DeityDualism bool // 60%
	Deism        bool // 20% Deism is the belief in the existence of God solely based on rational thought without any reliance on revealed religions or religious authority
	Henothism    bool // 40% is the worship of a single, supreme god while not denying the existence or possible existence of other lower deities.[
	Monolatry    bool // 55% is the belief in the existence of many gods, but with the consistent worship of only one deity
	Omnism       bool // 20%
}

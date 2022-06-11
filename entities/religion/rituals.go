package religion

type Rituals struct {
	RitualCelebrations *Celebrations
	Funerals           *Funerals
	Sacrifices         *Sacrifices
	Adorcism           bool
	Exorcism           bool
	Orgy               bool
	SmokeCircle        bool
	GladiatorDuel      bool
	Blinding           bool
	RitualCannibalism  bool
}

type Celebrations struct {
	DanceParty      bool
	DrumParty       bool
	SocialFestivals bool
	TreeCelebration bool
}

type Sacrifices struct {
	RitualSuicide   bool
	AnimalSacrifice bool
	HumanSacrifice  bool
	PlantsSacrifice bool
}

type Funerals struct {
	SkyBurials         bool
	CaveBurials        bool
	UndergroundBurials bool
}

// @todo fill it and use it
type Initiations struct{}

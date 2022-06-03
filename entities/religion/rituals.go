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

/*
Burn circle: A moral leader gives a condemnatory speech, and the people burn a hated symbol like an effigy or flag.
Smoke circle: People burn a special giant bong or incense shrine and inhale the smoke together.
Skylantern festival: People build and release beautiful wooden lanterns that float into the sky.
Gladiator duel: Slaves or prisoners fight in a ring to entertain spectators. You can supply weapons and they will be used. If a fighter dies, the crowd will come away even more excited.
Sacrifice: A moral leader kills an animal or prisoner before a crowd.
Cannibal feast: Colonists gather to devour a scrumptious platter of human meat.
Scarification: The moral guide inducts a new believer by aesthetically scarring them while others watch.
Blinding: The moral guide inducts a new believer by blinding them while others watch.
Funeral: People gather around a grave to remember someone they lost, while a moral guide gives a speech.
Some rituals aren’t specific to belief systems:

Leader speeches: The colony’s leader gives a rousing speech intended to boost morale. If their social skills fail, it might have the opposite effect!
Trial: The leader accuses someone of heinous crimes. The accuser and accused argue the case in front of witnesses. If convicted, the accused can be punished without social consequences.
Public execution: A killing of a prisoner. This goes best after a trial where the prisoner is found guilty of something.
Tree connection: A majestic ceremony where a colonist connects with a Gauranlen tree, giving them the ability to command the dryads inside.
*/

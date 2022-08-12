package culture

var (
	// This culture considers conflict and violence to be necessary states of existence;
	// ingrained in its people is the idea that one should stand up and fight for their own.
	Bellicose = &Ethos{
		Name:             "bellicose",
		IsDiplomatic:     false,
		IsWarlike:        true,
		IsAdministrative: false,
		IsIntrigue:       true,
		IsScholarly:      false,
	}
	// No culture lasts longer than its oldest record,
	// and this culture believes in accounting so meticulous as to guarantee immortality.
	Bureaucratic = &Ethos{
		Name:             "bureaucratic",
		IsDiplomatic:     false,
		IsWarlike:        true,
		IsAdministrative: true,
		IsIntrigue:       false,
		IsScholarly:      false,
	}
	// This culture values the bonds of community above all else,
	// fostering great loyalty and dedication by working together towards common goals.
	Communal = &Ethos{
		Name:             "communal",
		IsDiplomatic:     false,
		IsWarlike:        false,
		IsAdministrative: true,
		IsIntrigue:       true,
		IsScholarly:      false,
	}
	// he ceremony of court life is so integral to this culture that
	// it forms a core part of people's social behavior.
	// A place for everyone, and everyone in their place.
	Courtly = &Ethos{
		Name:             "courtly",
		IsDiplomatic:     true,
		IsWarlike:        false,
		IsAdministrative: false,
		IsIntrigue:       false,
		IsScholarly:      true,
	}
	// Intolerance and isolationism may be the way for others,
	// but this culture recognizes that accepting difference is far better than annihilating it.
	Egalitarian = &Ethos{
		Name:             "egalitarian",
		IsDiplomatic:     true,
		IsWarlike:        false,
		IsAdministrative: true,
		IsIntrigue:       false,
		IsScholarly:      false,
	}
	// While some cultures turn to war and others to worldly knowledge,
	// this culture places its trust in the only constant throughout its history - the divine.
	// Spirituality is the only way forward in a harsh and uncaring world.
	Spiritual = &Ethos{
		Name:             "spiritual",
		IsDiplomatic:     false,
		IsWarlike:        false,
		IsAdministrative: true,
		IsIntrigue:       false,
		IsScholarly:      true,
	}
	// This culture believes in standing strong like a mountain,
	// taking any and all hardships that life may throw at them and
	// enduring them all with grim determination and an indomitable spirit.
	Stoic = &Ethos{
		Name:             "stoic",
		IsDiplomatic:     true,
		IsWarlike:        true,
		IsAdministrative: false,
		IsIntrigue:       false,
		IsScholarly:      false,
	}
)

var AllEthoses = []*Ethos{
	Bellicose,
	Bureaucratic,
	Communal,
	Courtly,
	Egalitarian,
	Spiritual,
	Stoic,
}

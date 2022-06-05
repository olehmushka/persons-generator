package religion

type Taboos struct {
	RaisingTaboos          *RaisingTaboos
	EatingTaboos           *EatingTaboos
	SexualTaboos           *SexualTaboos
	KillingTaboos          *KillingTaboos
	BodyModificationTaboos *BodyModificationTaboos
	Witchcraft             Acceptance
	Nudism                 Acceptance
}

type RaisingTaboos struct {
	RaisingAnimals Acceptance
	RaisingPlants  Acceptance
	RaisingFungus  Acceptance
}

type EatingTaboos struct {
	EatAnimals               Acceptance
	EatInsects               Acceptance
	EatFungus                Acceptance
	Cannibalism              Acceptance
	DrinkingBlood            Acceptance
	DrinkingStrongAlcohol    Acceptance
	DrinkingNotStrongAlcohol Acceptance
	UseNicotine              Acceptance
	UseCannabis              Acceptance
	UseHallucinogens         Acceptance
	UseCNSStimulants         Acceptance
	UseOpium                 Acceptance
}

type SexualTaboos struct {
	SameSexRelations Acceptance
	MaleAdultery     Acceptance
	FemaleAdultery   Acceptance
	Deviancy         Acceptance
}

type KillingTaboos struct {
	Killing          Acceptance
	KillAnimals      Acceptance
	KillHollyAnimals Acceptance
	KillHumans       Acceptance
	Kinslaying       Acceptance
	Suicide          Acceptance
}

type BodyModificationTaboos struct {
	Tattoo Acceptance
}

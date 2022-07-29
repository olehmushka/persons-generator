package gender

import (
	"fmt"

	"persons_generator/entities/influence"
	pm "persons_generator/probability_machine"
)

type Dominance struct {
	Dominance GenderDominance
	Influence influence.Influence
}

func NewDominance() *Dominance {
	return &Dominance{
		Dominance: generateDominance(),
		Influence: generateInfluence(),
	}
}

func (gd *Dominance) Print() {
	fmt.Printf("Dominated gender=%s(%s)\n", gd.Dominance, gd.Influence)
}

func (gd *Dominance) IsMaleDominate() bool {
	return gd.Dominance == MaleDominance
}

func (gd *Dominance) IsEquality() bool {
	return gd.Dominance == EqualityDominance
}

func (gd *Dominance) IsFemaleDominate() bool {
	return gd.Dominance == FemaleDominance
}

func (gd *Dominance) GetCoef() float64 {
	const (
		strongGenderInfluence   = 1.25
		moderateGenderInfluence = 1
		weakGenderInfluence     = 0.75
	)

	switch gd.Influence {
	case influence.StrongInfluence:
		return strongGenderInfluence
	case influence.ModerateInfluence:
		return moderateGenderInfluence
	case influence.WeakInfluence:
		return weakGenderInfluence
	}

	return 1
}

func generateDominance() GenderDominance {
	return GenderDominance(pm.GetRandomFromSeveral(map[string]float64{
		string(MaleDominance):     pm.RandFloat64InRange(0.25, 0.35),
		string(EqualityDominance): pm.RandFloat64InRange(0.2, 0.3),
		string(FemaleDominance):   pm.RandFloat64InRange(0.15, 0.25),
	}))
}

func generateInfluence() influence.Influence {
	var (
		strong   = 0.15
		moderate = 0.15
		weak     = 0.1
	)

	return influence.GetInfluenceByProbability(strong, moderate, weak)
}

type GenderDominance string

const (
	MaleDominance     GenderDominance = "male"
	EqualityDominance GenderDominance = "equality"
	FemaleDominance   GenderDominance = "female"
)

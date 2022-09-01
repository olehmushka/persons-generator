package gender

import (
	"fmt"

	"persons_generator/engine/entities/influence"
	pm "persons_generator/engine/probability_machine"
)

type Dominance struct {
	Dominance GenderDominance     `json:"dominance"`
	Influence influence.Influence `json:"influence"`
}

func NewDominance() (*Dominance, error) {
	d, err := generateDominance()
	if err != nil {
		return nil, err
	}
	return &Dominance{
		Dominance: d,
		Influence: generateInfluence(),
	}, nil
}

func NewDominanceWithParams(d GenderDominance, i influence.Influence) *Dominance {
	return &Dominance{
		Dominance: d,
		Influence: i,
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

func generateDominance() (GenderDominance, error) {
	maleDominance, err := pm.RandFloat64InRange(0.25, 0.35)
	if err != nil {
		return "", err
	}
	equalityDominance, err := pm.RandFloat64InRange(0.2, 0.3)
	if err != nil {
		return "", err
	}
	femaleDominance, err := pm.RandFloat64InRange(0.15, 0.25)
	if err != nil {
		return "", err
	}

	return GetGenderDominanceByProbability(maleDominance, equalityDominance, femaleDominance), nil
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

func (gd GenderDominance) String() string {
	return string(gd)
}

const (
	MaleDominance     GenderDominance = "male"
	EqualityDominance GenderDominance = "equality"
	FemaleDominance   GenderDominance = "female"
)

func GetGenderDominanceByProbability(male, equality, female float64) GenderDominance {
	return GenderDominance(pm.GetRandomFromSeveral(map[string]float64{
		string(MaleDominance):     male,
		string(EqualityDominance): equality,
		string(FemaleDominance):   female,
	}))
}

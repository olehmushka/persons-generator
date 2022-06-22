package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
)

type GenderDominance struct {
	religion *Religion

	Dominance Dominance
	Influence Influence
}

func NewGenderDominance(r *Religion) *GenderDominance {
	gd := &GenderDominance{religion: r}
	gd.Dominance = gd.generateDominance()
	gd.Influence = gd.generateInfluence()

	return gd
}

func (gd *GenderDominance) Print() {
	fmt.Printf("Dominated gender=%s(%s)\n", gd.Dominance, gd.Influence)
}

func (gd *GenderDominance) IsMaleDominate() bool {
	return gd.Dominance == MaleDominance
}

func (gd *GenderDominance) IsEquality() bool {
	return gd.Dominance == EqualityDominance
}

func (gd *GenderDominance) IsFemaleDominate() bool {
	return gd.Dominance == FemaleDominance
}

func (gd *GenderDominance) GetCoef() float64 {
	const (
		strongGenderInfluence   = 1.25
		moderateGenderInfluence = 1
		weakGenderInfluence     = 0.75
	)

	switch gd.Influence {
	case StrongInfluence:
		return strongGenderInfluence
	case ModerateInfluence:
		return moderateGenderInfluence
	case WeakInfluence:
		return weakGenderInfluence
	}

	return 1
}

func (gd *GenderDominance) generateDominance() Dominance {
	var (
		male     = 0.3
		equality = 0.25
		female   = 0.15
	)
	switch {
	case gd.religion.Type.IsMonotheism():
		male += 0.1
		equality += 0.11
	case gd.religion.Type.IsPolytheism():
		male += 0.1
		equality += 0.1
		female += 0.1
	case gd.religion.Type.IsDeityDualism():
		male += 0.05
		equality += 0.07
	case gd.religion.Type.IsDeism():
		male += 0.05
		equality += 0.05
	case gd.religion.Type.IsAtheism():
		male += 0.7
		equality += 0.1
		female += 0.03
	}

	m := map[string]float64{
		string(MaleDominance):     pm.PrepareProbability(male),
		string(EqualityDominance): pm.PrepareProbability(equality),
		string(FemaleDominance):   pm.PrepareProbability(female),
	}

	return Dominance(pm.GetRandomFromSeveral(m))
}

func (gd *GenderDominance) generateInfluence() Influence {
	var (
		strong   = 0.15
		moderate = 0.15
		weak     = 0.1
	)
	switch {
	case gd.religion.Type.IsMonotheism():
		strong += 0.09
		moderate += 0.05
		weak += 0.005
	case gd.religion.Type.IsPolytheism():
		strong += 0.05
		moderate += 0.09
		weak += 0.03
	case gd.religion.Type.IsDeityDualism():
		strong += 0.1
		moderate += 0.05
		weak += 0.005
	case gd.religion.Type.IsDeism():
		strong += 0.05
		moderate += 0.05
		weak += 0.08
	case gd.religion.Type.IsAtheism():
		strong += 0.01
		moderate += 0.05
		weak += 0.1
	}

	if gd.IsEquality() {
		moderate += 0.25
	}

	return geInfluenceByProbability(strong, moderate, weak)
}

type Dominance string

const (
	MaleDominance     Dominance = "Male"
	EqualityDominance Dominance = "Equality"
	FemaleDominance   Dominance = "Female"
)

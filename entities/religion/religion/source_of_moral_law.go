package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
)

type SourceOfMoralLaw string

const (
	DeitySourceOfMoralLaw  SourceOfMoralLaw = "Deity"
	NoneSourceOfMoralLaw   SourceOfMoralLaw = "None"
	HumanSourceOfMoralLaw  SourceOfMoralLaw = "Human"
	NatureSourceOfMoralLaw SourceOfMoralLaw = "Nature"
)

func (d *Doctrine) generateSourceOfMoralLaw() SourceOfMoralLaw {
	var (
		deity  = pm.RandFloat64InRange(0.03, 0.15)
		none   = pm.RandFloat64InRange(0.01, 0.05)
		human  = pm.RandFloat64InRange(0.02, 0.09)
		nature = pm.RandFloat64InRange(0.025, 0.095)
		rmCoef = pm.RandFloat64InRange(0.01, 0.02)
	)
	human += rmCoef * d.religion.metadata.Hedonism

	if d.Deity.Nature.Goodness.Goodness == Good {
		var goodnessCoef float64
		switch d.Deity.Nature.Goodness.Level {
		case Major:
			goodnessCoef = 1.1
		case Middle:
			goodnessCoef = 1
		case Minor:
			goodnessCoef = 0.4
		}
		deity += 0.1 * goodnessCoef
	}
	for _, trait := range d.Deity.Nature.Traits {
		if trait.Name == "IsJust" {
			deity += 0.5
		}
	}

	if d.Human.Nature.Goodness.Goodness == Good {
		var goodnessCoef float64
		switch d.Human.Nature.Goodness.Level {
		case Major:
			goodnessCoef = 1.1
		case Middle:
			goodnessCoef = 1
		case Minor:
			goodnessCoef = 0.4
		}
		human += 0.1 * goodnessCoef
	}

	return getSourceOfMoralLawByProbability(deity, none, human, nature)
}

func (soml SourceOfMoralLaw) Print() {
	fmt.Printf("Source of moral law is %s\n", soml)
}

func getSourceOfMoralLawByProbability(deity, none, human, nature float64) SourceOfMoralLaw {
	m := map[string]float64{
		string(DeitySourceOfMoralLaw):  pm.PrepareProbability(deity),
		string(NoneSourceOfMoralLaw):   pm.PrepareProbability(none),
		string(HumanSourceOfMoralLaw):  pm.PrepareProbability(human),
		string(NatureSourceOfMoralLaw): pm.PrepareProbability(nature),
	}
	return SourceOfMoralLaw(pm.GetRandomFromSeveral(m))
}

func (s SourceOfMoralLaw) IsDeity() bool {
	return s == DeitySourceOfMoralLaw
}

func (s SourceOfMoralLaw) IsNone() bool {
	return s == NoneSourceOfMoralLaw
}

func (s SourceOfMoralLaw) IsHuman() bool {
	return s == HumanSourceOfMoralLaw
}

func (s SourceOfMoralLaw) IsNature() bool {
	return s == NatureSourceOfMoralLaw
}

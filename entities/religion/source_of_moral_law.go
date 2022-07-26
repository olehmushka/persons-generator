package religion

import (
	"fmt"

	pm "persons_generator/probability_machine"
)

type SourceOfMoralLaw string

const (
	DeitySourceOfMoralLaw  SourceOfMoralLaw = "Deity"
	NoneSourceOfMoralLaw   SourceOfMoralLaw = "None"
	HumanSourceOfMoralLaw  SourceOfMoralLaw = "Human"
	NatureSourceOfMoralLaw SourceOfMoralLaw = "Nature"
)

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

func (d *Doctrine) generateSourceOfMoralLaw() SourceOfMoralLaw {
	if d.religion.Type.IsAtheism() {
		return NoneSourceOfMoralLaw
	}

	var (
		deity  = pm.RandFloat64InRange(0.1, 0.2)
		none   = pm.RandFloat64InRange(0.01, 0.1)
		human  = pm.RandFloat64InRange(0.05, 0.15)
		nature = pm.RandFloat64InRange(0.05, 0.15)
		rmCoef = d.religion.M.LowBaseCoef
	)
	if d.religion.Type.IsMonotheism() {
		deity += pm.RandFloat64InRange(0.05, 0.1)
	}
	if d.religion.metadata.IsHedonistic() {
		human += rmCoef * pm.RandFloat64InRange(0.05, 0.15)
	}

	if d.Deity.Nature.Goodness.Goodness == Good {
		deity += pm.RandFloat64InRange(0.05, 0.15) * d.Deity.Nature.Goodness.Level.GetLevelCoef()
	}
	for _, trait := range d.Deity.Nature.Traits {
		if trait.Name == "IsJust" {
			deity += pm.RandFloat64InRange(0.01, 0.1)
		}
	}

	if d.Human.Nature.Goodness.Goodness == Good {
		human += pm.RandFloat64InRange(0.05, 0.15) * d.Human.Nature.Goodness.Level.GetLevelCoef()
	}

	return getSourceOfMoralLawByProbability(deity, none, human, nature)
}

func (soml SourceOfMoralLaw) Print() {
	fmt.Printf("Source of moral law is %s\n", soml)
}

func getSourceOfMoralLawByProbability(deity, none, human, nature float64) SourceOfMoralLaw {
	return SourceOfMoralLaw(pm.GetRandomFromSeveral(map[string]float64{
		string(DeitySourceOfMoralLaw):  pm.PrepareProbability(deity),
		string(NoneSourceOfMoralLaw):   pm.PrepareProbability(none),
		string(HumanSourceOfMoralLaw):  pm.PrepareProbability(human),
		string(NatureSourceOfMoralLaw): pm.PrepareProbability(nature),
	}))
}

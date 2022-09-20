package religion

import (
	"fmt"

	"persons_generator/core/wrapped_error"
	pm "persons_generator/engine/probability_machine"
)

type SourceOfMoralLaw string

func (soml SourceOfMoralLaw) String() string {
	return string(soml)
}

const (
	DeitySourceOfMoralLaw  SourceOfMoralLaw = "deity"
	NoneSourceOfMoralLaw   SourceOfMoralLaw = "none"
	HumanSourceOfMoralLaw  SourceOfMoralLaw = "human"
	NatureSourceOfMoralLaw SourceOfMoralLaw = "nature"
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

func (d *Doctrine) generateSourceOfMoralLaw() (SourceOfMoralLaw, error) {
	if d.religion.Type.IsAtheism() {
		return NoneSourceOfMoralLaw, nil
	}

	deity, err := pm.RandFloat64InRange(0.1, 0.2)
	if err != nil {
		return "", err
	}
	none, err := pm.RandFloat64InRange(0.01, 0.1)
	if err != nil {
		return "", err
	}
	human, err := pm.RandFloat64InRange(0.05, 0.15)
	if err != nil {
		return "", err
	}
	nature, err := pm.RandFloat64InRange(0.05, 0.15)
	if err != nil {
		return "", err
	}
	rmCoef := d.religion.M.LowBaseCoef
	if d.religion.Type.IsMonotheism() {
		deityP, err := pm.RandFloat64InRange(0.05, 0.1)
		if err != nil {
			return "", err
		}
		deity += deityP
	}
	if d.religion.Metadata.IsHedonistic() {
		humanP, err := pm.RandFloat64InRange(0.05, 0.15)
		if err != nil {
			return "", err
		}
		human += rmCoef * humanP
	}

	if d.Deity.Nature.Goodness.Goodness == Good {
		deityP, err := pm.RandFloat64InRange(0.05, 0.15)
		if err != nil {
			return "", err
		}
		deity += deityP * d.Deity.Nature.Goodness.Level.GetCoef()
	}
	for _, trait := range d.Deity.Nature.Traits {
		if trait.Name == IsJustDeityTrait {
			deityP, err := pm.RandFloat64InRange(0.01, 0.1)
			if err != nil {
				return "", err
			}
			deity += deityP
		}
	}

	if d.Human.Nature.Goodness.Goodness == Good {
		humanP, err := pm.RandFloat64InRange(0.05, 0.15)
		if err != nil {
			return "", err
		}
		human += humanP * d.Human.Nature.Goodness.Level.GetCoef()
	}

	return getSourceOfMoralLawByProbability(deity, none, human, nature)
}

func (soml SourceOfMoralLaw) Print() {
	fmt.Printf("Source of moral law is %s\n", soml)
}

func getSourceOfMoralLawByProbability(deity, none, human, nature float64) (SourceOfMoralLaw, error) {
	soml, err := pm.GetRandomFromSeveral(map[string]float64{
		string(DeitySourceOfMoralLaw):  pm.PrepareProbability(deity),
		string(NoneSourceOfMoralLaw):   pm.PrepareProbability(none),
		string(HumanSourceOfMoralLaw):  pm.PrepareProbability(human),
		string(NatureSourceOfMoralLaw): pm.PrepareProbability(nature),
	})
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generate source of moral law")
	}

	return SourceOfMoralLaw(soml), nil
}

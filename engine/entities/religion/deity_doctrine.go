package religion

import (
	"fmt"

	pm "persons_generator/engine/probability_machine"
)

type DeityDoctrine struct {
	religion *Religion `json:"-" bson:"-"`
	doctrine *Doctrine `json:"-" bson:"-"`

	Nature *DeityNature `json:"nature" bson:"nature"`
}

func (d *Doctrine) generateDeityDoctrine() (*DeityDoctrine, error) {
	dd := &DeityDoctrine{religion: d.religion, doctrine: d}
	nature, err := dd.generateDeityNature()
	if err != nil {
		return nil, err
	}
	dd.Nature = nature

	return dd, nil
}

func (dd *DeityDoctrine) SerializeNatureTraits() []string {
	if dd == nil || dd.Nature == nil {
		return []string{}
	}

	return extractTraitNames(dd.Nature.Traits)
}

func (dd *DeityDoctrine) Print() {
	dd.Nature.Print()
}

type DeityNature struct {
	religion      *Religion      `json:"-"`
	deityDoctrine *DeityDoctrine `json:"-"`

	Goodness GoodnessNature `json:"goodness"`
	Traits   []*trait       `json:"traits"`
}

func (dn *DeityNature) Print() {
	fmt.Printf("Deity(s) is/are %s and level of it is %s\n", dn.Goodness.Goodness, dn.Goodness.Level)
	if len(dn.Traits) > 0 {
		fmt.Printf("Deity Traits (religion_name=%s):\n", dn.religion.Name)
		for _, trait := range dn.Traits {
			fmt.Printf(" - %s\n", trait.Name)
		}
	}
}

func (dd *DeityDoctrine) generateDeityNature() (*DeityNature, error) {
	dn := &DeityNature{religion: dd.religion, deityDoctrine: dd}
	goodness, err := dn.generateGoodness()
	if err != nil {
		return nil, err
	}
	dn.Goodness = goodness
	if !dn.religion.Type.IsAtheism() {
		ts, err := generateTraits(dd.religion, dn.getAllDeityNatureTraits(), generateTraitsOpts{
			Label: "DeityNature.generateTraits",
			Min:   1,
			Max:   5,
		})
		if err != nil {
			return nil, err
		}
		dn.Traits = ts
	}

	return dn, nil
}

func (dd *DeityNature) generateGoodness() (GoodnessNature, error) {
	if dd.religion.Type.IsAtheism() {
		return GoodnessNature{
			Level:    Minor,
			Goodness: Neutral,
		}, nil
	}
	out := GoodnessNature{}
	goodness, err := dd.getGoodnessByReligionMetadata()
	if err != nil {
		return GoodnessNature{}, err
	}
	out.Goodness = goodness
	level, err := dd.getGoodnessLevelByReligionMetadata()
	if err != nil {
		return GoodnessNature{}, err
	}
	out.Level = level

	return out, nil
}

func (dn *DeityNature) getGoodnessByReligionMetadata() (Goodness, error) {
	rmCoef, err := pm.RandFloat64InRange(0.05, 0.15)
	if err != nil {
		return "", err
	}
	good, err := pm.RandFloat64InRange(0.15, 0.25)
	if err != nil {
		return "", err
	}
	neutral, err := pm.RandFloat64InRange(0.1, 0.2)
	if err != nil {
		return "", err
	}
	evil, err := pm.RandFloat64InRange(0.05, 0.15)
	if err != nil {
		return "", err
	}

	if dn.religion.Metadata.IsHedonistic() || dn.religion.Metadata.IsAltruistic() {
		pGood, err := pm.RandFloat64InRange(0.05, 0.1)
		if err != nil {
			return "", err
		}
		good += pGood * rmCoef
		pNeutral, err := pm.RandFloat64InRange(0.05, 0.1)
		if err != nil {
			return "", err
		}
		neutral += pNeutral * rmCoef
	}
	if dn.religion.Metadata.IsChthonic() {
		p, err := pm.RandFloat64InRange(0.1, 0.25)
		if err != nil {
			return "", err
		}
		evil += p * rmCoef
	}

	for _, goal := range dn.deityDoctrine.doctrine.HighGoal.Goals {
		switch goal.Name {
		case FightAgainstEvilHighGoal:
			p, err := pm.RandFloat64InRange(0.15, 0.25)
			if err != nil {
				return "", err
			}
			good += p * rmCoef
		case FightForEvilHighGoal:
			p, err := pm.RandFloat64InRange(0.15, 0.25)
			if err != nil {
				return "", err
			}
			evil += p
		}
	}

	return getGoodnessByProbability(good, neutral, evil)
}

func (dn *DeityNature) getGoodnessLevelByReligionMetadata() (Level, error) {
	rmCoef, err := pm.RandFloat64InRange(0.05, 0.15)
	if err != nil {
		return "", err
	}
	major, err := pm.RandFloat64InRange(0.1, 0.2)
	if err != nil {
		return "", err
	}
	middle, err := pm.RandFloat64InRange(0.1, 0.2)
	if err != nil {
		return "", err
	}
	minor, err := pm.RandFloat64InRange(0.1, 0.2)
	if err != nil {
		return "", err
	}

	if dn.religion.Metadata.IsAuthoritaristic() {
		p, err := pm.RandFloat64InRange(0.1, 0.25)
		if err != nil {
			return "", err
		}
		major += p * rmCoef
	}
	if dn.religion.Metadata.IsLiberal() {
		p, err := pm.RandFloat64InRange(0.1, 0.25)
		if err != nil {
			return "", err
		}
		minor += p * rmCoef
	}

	return getLevelByProbability(major, middle, minor)
}

func (dd *DeityNature) getAllDeityNatureTraits() []*trait {
	return []*trait{
		{
			Name: IsJustDeityTrait,
			ReligionMetadata: &religionMetadata{
				Lawful: 1,
				Simple: 0.25,
			},
			BaseCoef: dd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: IsWorldCreatorDeityTrait,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Lawful:       0.25,
				Simple:       0.25,
			},
			BaseCoef: dd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: IsImmortalDeityTrait,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Lawful:       0.25,
			},
			BaseCoef: dd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				var addCoef float64
				if r.Type.IsMonotheism() {
					c, err := pm.RandFloat64InRange(0.15, 0.25)
					if err != nil {
						return false, err
					}
					addCoef += c
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: IsTranscendentalDeityTrait,
			ReligionMetadata: &religionMetadata{
				Complicated: 1,
			},
			BaseCoef: dd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				var addCoef float64
				switch {
				case r.Type.IsAtheism():
					return false, nil
				case r.Type.IsMonotheism():
					c, err := pm.RandFloat64InRange(0.15, 0.25)
					if err != nil {
						return false, err
					}
					addCoef += c
				case r.Type.IsDeism():
					fallthrough
				case r.Type.IsDeityDualism():
					c, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += c
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: TakePartInHumanLifeDeityTrait,
			ReligionMetadata: &religionMetadata{
				Naturalistic:    0.25,
				Authoritaristic: 1,
			},
			BaseCoef: dd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: CommunicateWithHumansDeityTrait,
			ReligionMetadata: &religionMetadata{
				Naturalistic:    0.25,
				Authoritaristic: 1,
			},
			BaseCoef: dd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: IsGnosticDeityDeityTrait, // God(s) created spiritual world but Demiurge(s) created physical world. Body is sinful
			ReligionMetadata: &religionMetadata{
				Chthonic:    0.25,
				Complicated: 1,
			},
			BaseCoef: dd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				var addCoef float64
				if r.Type.IsDeityDualism() {
					c, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += c
				}
				if dd.Goodness.Goodness == Evil {
					c, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += c
				}
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func GetDeityDoctrineSimilarityCoef(d1, d2 *DeityDoctrine) float64 {
	if d1 == nil && d2 == nil {
		return 1
	}
	if d1 == nil || d2 == nil {
		return 0
	}

	similarityTraits := []struct {
		enable bool
		value  float64
		coef   float64
	}{
		{
			enable: true,
			value:  GetGoodnessNatureSimilarityCoef(d1.Nature.Goodness, d2.Nature.Goodness),
			coef:   0.5,
		},
		{
			enable: true,
			value:  GetTraitsSimilarityCoef(d1.Nature.Traits, d2.Nature.Traits),
			coef:   0.5,
		},
	}

	var out float64
	for _, t := range similarityTraits {
		if t.enable {
			out += t.value * t.coef
		}
	}

	return out
}

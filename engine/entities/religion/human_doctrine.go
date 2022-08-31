package religion

import (
	"fmt"

	pm "persons_generator/engine/probability_machine"
)

type HumanDoctrine struct {
	religion *Religion
	doctrine *Doctrine

	Nature *HumanNature
}

func (d *Doctrine) generateHumanDoctrine() (*HumanDoctrine, error) {
	hd := &HumanDoctrine{religion: d.religion, doctrine: d}
	n, err := hd.generateHumanNature()
	if err != nil {
		return nil, err
	}
	hd.Nature = n

	return hd, nil
}

func (hd *HumanDoctrine) Print() {
	hd.Nature.Print()
}

type HumanNature struct {
	religion      *Religion
	humanDoctrine *HumanDoctrine

	Goodness GoodnessNature
	Traits   []*trait
}

func (hd *HumanDoctrine) generateHumanNature() (*HumanNature, error) {
	hn := &HumanNature{religion: hd.religion, humanDoctrine: hd}
	goodness, err := hn.generateGoodness()
	if err != nil {
		return nil, err
	}
	hn.Goodness = goodness
	ts, err := generateTraits(hd.religion, hn.getAllHumanNatureTraits(), generateTraitsOpts{
		Label: "HumanNature.generateTraits",
		Min:   1,
		Max:   2,
	})
	if err != nil {
		return nil, err
	}
	hn.Traits = ts

	return hn, nil
}

func (hn *HumanNature) Print() {
	fmt.Printf("Human is %s and level of it is %s\n", hn.Goodness.Goodness, hn.Goodness.Level)
	if len(hn.Traits) > 0 {
		fmt.Printf("Human Traits (religion_name=%s):\n", hn.religion.Name)
		for _, trait := range hn.Traits {
			fmt.Printf(" - %s\n", trait.Name)
		}
	}
}

func (hn *HumanNature) generateGoodness() (GoodnessNature, error) {
	out := GoodnessNature{}
	goodness, err := hn.getGoodnessByReligionMetadata()
	if err != nil {
		return out, err
	}
	out.Goodness = goodness

	level, err := hn.getGoodnessLevelByReligionMetadata()
	if err != nil {
		return out, err
	}
	out.Level = level

	return out, nil
}

func (hn *HumanNature) getGoodnessByReligionMetadata() (Goodness, error) {
	rmCoef, err := pm.RandFloat64InRange(0.05, 0.15)
	if err != nil {
		return "", err
	}
	good, err := pm.RandFloat64InRange(0.1, 0.2)
	if err != nil {
		return "", err
	}
	neutral, err := pm.RandFloat64InRange(0.1, 0.2)
	if err != nil {
		return "", err
	}
	evil, err := pm.RandFloat64InRange(0.1, 0.2)
	if err != nil {
		return "", err
	}

	if hn.religion.metadata.IsHedonistic() || hn.religion.metadata.IsAltruistic() {
		goodP, err := pm.RandFloat64InRange(0.05, 0.1)
		if err != nil {
			return "", err
		}
		good += goodP * rmCoef
		neutralP, err := pm.RandFloat64InRange(0.05, 0.1)
		if err != nil {
			return "", err
		}
		neutral += neutralP * rmCoef
	}
	if hn.religion.metadata.IsChthonic() {
		evilP, err := pm.RandFloat64InRange(0.1, 0.25)
		if err != nil {
			return "", err
		}
		evil += evilP * rmCoef
	}

	for _, goal := range hn.humanDoctrine.doctrine.HighGoal.Goals {
		if goal.Name == InvestigateMyselfHighGoal {
			goodP, err := pm.RandFloat64InRange(0.01, 0.1)
			if err != nil {
				return "", err
			}
			good += goodP
		}
	}

	return getGoodnessByProbability(good, neutral, evil), nil
}

func (hn *HumanNature) getGoodnessLevelByReligionMetadata() (Level, error) {
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
	if hn.religion.metadata.IsLiberal() {
		minorP, err := pm.RandFloat64InRange(0.1, 0.25)
		if err != nil {
			return "", err
		}
		minor += minorP * rmCoef
	}

	return getLevelByProbability(major, middle, minor), nil
}

func (hn *HumanNature) getAllHumanNatureTraits() []*trait {
	return []*trait{
		{
			Name: HasSoulHumanTrait,
			_religionMetadata: &religionMetadata{
				Naturalistic:    0.5,
				Individualistic: 0.75,
			},
			baseCoef: hn.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: CanBeSaintHumanTrait,
			_religionMetadata: &religionMetadata{
				Ascetic:         0.5,
				Individualistic: 1,
			},
			baseCoef: hn.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				var addCoef float64
				for _, goal := range hn.humanDoctrine.doctrine.HighGoal.Goals {
					if goal == nil {
						continue
					}
					if goal.Name == BecomePerfectAndSaintsHighGoal {
						coef, err := pm.RandFloat64InRange(0.15, 0.25)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

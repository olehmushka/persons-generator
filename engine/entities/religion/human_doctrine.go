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

func (d *Doctrine) generateHumanDoctrine() *HumanDoctrine {
	hd := &HumanDoctrine{religion: d.religion, doctrine: d}
	hd.Nature = hd.generateHumanNature()

	return hd
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

func (hd *HumanDoctrine) generateHumanNature() *HumanNature {
	hn := &HumanNature{religion: hd.religion, humanDoctrine: hd}
	hn.Goodness = hn.generateGoodness()
	hn.Traits = generateTraits(hd.religion, hn.getAllHumanNatureTraits(), generateTraitsOpts{
		Label: "HumanNature.generateTraits",
		Min:   1,
		Max:   2,
	})

	return hn
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

func (hn *HumanNature) generateGoodness() GoodnessNature {
	return GoodnessNature{
		Goodness: hn.getGoodnessByReligionMetadata(),
		Level:    hn.getGoodnessLevelByReligionMetadata(),
	}
}

func (hn *HumanNature) getGoodnessByReligionMetadata() Goodness {
	var (
		rmCoef  = pm.RandFloat64InRange(0.05, 0.15)
		good    = pm.RandFloat64InRange(0.1, 0.2)
		neutral = pm.RandFloat64InRange(0.1, 0.2)
		evil    = pm.RandFloat64InRange(0.1, 0.2)
	)
	if hn.religion.metadata.IsHedonistic() || hn.religion.metadata.IsAltruistic() {
		good += pm.RandFloat64InRange(0.05, 0.1) * rmCoef
		neutral += pm.RandFloat64InRange(0.05, 0.1) * rmCoef
	}
	if hn.religion.metadata.IsChthonic() {
		evil += pm.RandFloat64InRange(0.1, 0.25) * rmCoef
	}

	for _, goal := range hn.humanDoctrine.doctrine.HighGoal.Goals {
		if goal.Name == "InvestigateMyself" {
			good += pm.RandFloat64InRange(0.01, 0.1)
		}
	}

	return getGoodnessByProbability(good, neutral, evil)
}

func (hn *HumanNature) getGoodnessLevelByReligionMetadata() Level {
	var (
		rmCoef = pm.RandFloat64InRange(0.05, 0.15)
		major  = pm.RandFloat64InRange(0.1, 0.2)
		middle = pm.RandFloat64InRange(0.1, 0.2)
		minor  = pm.RandFloat64InRange(0.1, 0.2)
	)
	if hn.religion.metadata.IsLiberal() {
		minor += pm.RandFloat64InRange(0.1, 0.25) * rmCoef
	}

	return getLevelByProbability(major, middle, minor)
}

func (hn *HumanNature) getAllHumanNatureTraits() []*trait {
	return []*trait{
		{
			Name: "HasSoul",
			_religionMetadata: &religionMetadata{
				Naturalistic:    0.5,
				Individualistic: 0.75,
			},
			baseCoef: hn.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "CanBeSaint",
			_religionMetadata: &religionMetadata{
				Ascetic:         0.5,
				Individualistic: 1,
			},
			baseCoef: hn.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) bool {
				var addCoef float64
				for _, goal := range hn.humanDoctrine.doctrine.HighGoal.Goals {
					if goal == nil {
						continue
					}
					if goal.Name == "BecomePerfectAndSaints" {
						addCoef += pm.RandFloat64InRange(0.15, 0.25)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

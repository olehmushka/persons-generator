package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
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
	Traits   []*humanNatureTrait
}

func (hd *HumanDoctrine) generateHumanNature() *HumanNature {
	hn := &HumanNature{religion: hd.religion, humanDoctrine: hd}
	hn.Goodness = hn.generateGoodness()
	hn.Traits = hn.generateTraits(1, 2)

	return hn
}

func (hn *HumanNature) Print() {
	fmt.Printf("Human is %s and level of it is %s\n", hn.Goodness.Goodness, hn.Goodness.Level)
	fmt.Printf("Human Traits (religion_name=%s):\n", hn.religion.Name)
	for _, trait := range hn.Traits {
		fmt.Printf(" - %s\n", trait.Name)
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

func (hn *HumanNature) generateTraits(min, max int) []*humanNatureTrait {
	if min < 0 {
		panic("[HumanNature.generateTraits] min can not be less than 0")
	}
	allTraits := hn.getAllHumanNatureTraits()
	if max > len(allTraits) {
		panic("[HumanNature.generateTraits] max can not be greater than allTraits length")
	}

	traits := make([]*humanNatureTrait, 0, len(allTraits))
	for count := 0; count < 20; count++ {
		for _, trait := range allTraits {
			if trait.Calc(hn.religion, trait, traits) {
				traits = append(traits, &humanNatureTrait{
					_religionMetadata: trait._religionMetadata,
					Name:              trait.Name,
					Calc:              trait.Calc,
				})
			}
			if len(traits) == max {
				break
			}
		}
		if len(traits) == max {
			break
		}
		if len(traits) >= min {
			break
		}
	}

	for _, trait := range traits {
		hn.religion.UpdateMetadata(UpdateReligionMetadata(*hn.religion.metadata, *trait._religionMetadata))
	}

	return traits
}

type humanNatureTrait struct {
	_religionMetadata *religionMetadata
	baseCoef          float64
	Name              string
	Calc              func(r *Religion, self *humanNatureTrait, selectedTraits []*humanNatureTrait) bool
}

func (hn *HumanNature) getAllHumanNatureTraits() []*humanNatureTrait {
	return []*humanNatureTrait{
		{
			Name: "HasSoul",
			_religionMetadata: &religionMetadata{
				Naturalistic:    0.5,
				Individualistic: 0.75,
			},
			baseCoef: hn.religion.M.BaseCoef,
			Calc: func(r *Religion, self *humanNatureTrait, _ []*humanNatureTrait) bool {
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
			Calc: func(r *Religion, self *humanNatureTrait, selectedTraits []*humanNatureTrait) bool {
				var addCoef float64
				for _, goal := range hn.humanDoctrine.doctrine.HighGoal.Goals {
					if goal.Name == "BecomePerfectAndSaints" {
						addCoef += pm.RandFloat64InRange(0.15, 0.25)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

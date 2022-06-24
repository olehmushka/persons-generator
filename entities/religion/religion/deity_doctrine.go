package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
)

type DeityDoctrine struct {
	religion *Religion
	doctrine *Doctrine

	Nature *DeityNature
}

func (d *Doctrine) generateDeityDoctrine() *DeityDoctrine {
	dd := &DeityDoctrine{religion: d.religion, doctrine: d}
	dd.Nature = dd.generateDeityNature()

	return dd
}

func (dd *DeityDoctrine) Print() {
	dd.Nature.Print()
}

type DeityNature struct {
	religion      *Religion
	deityDoctrine *DeityDoctrine

	Goodness GoodnessNature
	Traits   []*deityNatureTrait
}

func (dn *DeityNature) Print() {
	fmt.Printf("Deity(s) is/are %s and level of it is %s\n", dn.Goodness.Goodness, dn.Goodness.Level)
	fmt.Printf("Deity Traits (religion_name=%s):\n", dn.religion.Name)
	for _, trait := range dn.Traits {
		fmt.Printf(" - %s\n", trait.Name)
	}
}

func (dd *DeityDoctrine) generateDeityNature() *DeityNature {
	dn := &DeityNature{religion: dd.religion, deityDoctrine: dd}
	dn.Goodness = dn.generateGoodness()
	if !dn.religion.Type.IsAtheism() {
		dn.Traits = dn.generateTraits(1, 5)
	}

	return dn
}

func (dd *DeityNature) generateGoodness() GoodnessNature {
	if dd.religion.Type.IsAtheism() {
		return GoodnessNature{
			Level:    Minor,
			Goodness: Neutral,
		}
	}

	return GoodnessNature{
		Goodness: dd.getGoodnessByReligionMetadata(),
		Level:    dd.getGoodnessLevelByReligionMetadata(),
	}
}

func (dn *DeityNature) getGoodnessByReligionMetadata() Goodness {
	var (
		rmCoef  = pm.RandFloat64InRange(0.05, 0.15)
		good    = pm.RandFloat64InRange(0.15, 0.25)
		neutral = pm.RandFloat64InRange(0.1, 0.2)
		evil    = pm.RandFloat64InRange(0.05, 0.15)
	)
	if dn.religion.metadata.IsHedonistic() || dn.religion.metadata.IsAltruistic() {
		good += pm.RandFloat64InRange(0.05, 0.1) * rmCoef
		neutral += pm.RandFloat64InRange(0.05, 0.1) * rmCoef
	}
	if dn.religion.metadata.IsChthonic() {
		evil += pm.RandFloat64InRange(0.1, 0.25) * rmCoef
	}

	for _, goal := range dn.deityDoctrine.doctrine.HighGoal.Goals {
		switch goal.Name {
		case "FightAgainstEvil":
			good += pm.RandFloat64InRange(0.15, 0.25)
		case "FightForEvil":
			evil += pm.RandFloat64InRange(0.15, 0.25)
		}
	}

	return getGoodnessByProbability(good, neutral, evil)
}

func (dn *DeityNature) getGoodnessLevelByReligionMetadata() Level {
	var (
		rmCoef = pm.RandFloat64InRange(0.05, 0.15)
		major  = pm.RandFloat64InRange(0.1, 0.2)
		middle = pm.RandFloat64InRange(0.1, 0.2)
		minor  = pm.RandFloat64InRange(0.1, 0.2)
	)
	if dn.religion.metadata.IsAuthoritaristic() {
		major += pm.RandFloat64InRange(0.1, 0.25) * rmCoef
	}
	if dn.religion.metadata.IsLiberal() {
		minor += pm.RandFloat64InRange(0.1, 0.25) * rmCoef
	}

	return getLevelByProbability(major, middle, minor)
}

func (dd *DeityNature) generateTraits(min, max int) []*deityNatureTrait {
	if min < 0 {
		panic("[DeityNature.generateTraits] min can not be less than 0")
	}
	if dd.religion.Type.IsAtheism() {
		return []*deityNatureTrait{}
	}
	allTraits := dd.getAllDeityNatureTraits()
	if max > len(allTraits) {
		panic("[DeityNature.generateTraits] max can not be greater than allTraits length")
	}

	traits := make([]*deityNatureTrait, 0, len(allTraits))
	for count := 0; count < 20; count++ {
		for _, trait := range allTraits {
			if trait.Calc(dd.religion, trait, traits) {
				traits = append(traits, &deityNatureTrait{
					_religionMetadata: trait._religionMetadata,
					baseCoef:          trait.baseCoef,
					Name:              trait.Name,
					Calc:              trait.Calc,
				})
				if len(traits) == max {
					break
				}
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
		dd.religion.UpdateMetadata(UpdateReligionMetadata(*dd.religion.metadata, *trait._religionMetadata))
	}

	return traits
}

type deityNatureTrait struct {
	_religionMetadata *religionMetadata
	baseCoef          float64
	Name              string
	Calc              func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool
}

func (dd *DeityNature) getAllDeityNatureTraits() []*deityNatureTrait {
	return []*deityNatureTrait{
		{
			Name: "IsJust",
			_religionMetadata: &religionMetadata{
				Lawful: 1,
				Simple: 0.25,
			},
			baseCoef: dd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "IsWorldCreator",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Lawful:       0.25,
				Simple:       0.25,
			},
			baseCoef: dd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "IsImmortal",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Lawful:       0.25,
			},
			baseCoef: dd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				var addCoef float64
				if r.Type.IsMonotheism() {
					addCoef += pm.RandFloat64InRange(0.15, 0.25)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "IsTranscendental",
			_religionMetadata: &religionMetadata{
				Complicated: 1,
			},
			baseCoef: dd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				var addCoef float64
				if r.Type.IsMonotheism() || r.Type.IsDeism() || r.Type.IsDeityDualism() {
					addCoef += pm.RandFloat64InRange(0.05, 0.15)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "TakePartInHumanLife",
			_religionMetadata: &religionMetadata{
				Naturalistic:    0.25,
				Authoritaristic: 1,
			},
			baseCoef: dd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "CommunicateWithHumans",
			_religionMetadata: &religionMetadata{
				Naturalistic:    0.25,
				Authoritaristic: 1,
			},
			baseCoef: dd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "IsGnosticDeity", // God(s) created spiritual world but Demiurge(s) created physical world. Body is sinful
			_religionMetadata: &religionMetadata{
				Chthonic:    0.25,
				Complicated: 1,
			},
			baseCoef: dd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				var addCoef float64
				if r.Type.IsDeityDualism() {
					addCoef += pm.RandFloat64InRange(0.05, 0.15)
				}
				if dd.Goodness.Goodness == Evil {
					addCoef += pm.RandFloat64InRange(0.05, 0.15)
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

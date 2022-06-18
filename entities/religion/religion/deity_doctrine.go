package religion

import "fmt"

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
	fmt.Printf("Deity Doctrine (religion=%s)\n", dd.religion.Name)
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
	dn.Traits = dn.generateTraits(1, 5)

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
		rmCoef  = 0.1
		good    = 0.2 + rmCoef*dn.religion.metadata.Hedonism
		neutral = 0.15 + rmCoef*dn.religion.metadata.Hedonism
		evil    = 0.1 + rmCoef*dn.religion.metadata.Chthonic
	)
	for _, goal := range dn.deityDoctrine.doctrine.HighGoal.Goals {
		switch goal.Name {
		case "FightAgainstEvil":
			good += 0.2
		case "FightForEvil":
			evil += 0.2
		}
	}

	return getGoodnessByProbability(good, neutral, evil)
}

func (dd *DeityNature) getGoodnessLevelByReligionMetadata() Level {
	var (
		rmCoef = 0.1
		major  = 0.15 + rmCoef*dd.religion.metadata.Fanaticism
		middle = 0.15
		minor  = 0.15
	)
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
		for i, trait := range allTraits {
			if trait.Calc(dd.religion, trait, traits) {
				traits = append(traits, &deityNatureTrait{
					_religionMetadata: trait._religionMetadata,
					Index:             i,
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

	return traits
}

type deityNatureTrait struct {
	_religionMetadata *updateReligionMetadata
	Index             int
	Name              string
	Calc              func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool
}

func (dd *DeityNature) getAllDeityNatureTraits() []*deityNatureTrait {
	return []*deityNatureTrait{
		{
			Name: "IsJust",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.02),
				OutsideDirected:  Float64(0.06),
				Strictness:       Float64(0.1),
				Ascetic:          Float64(0.01),
				Elitaric:         Float64(0.01),
				Organized:        Float64(0.01),
			},
			Calc: func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool {
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
		{
			Name: "IsWorldCreator",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.09),
				OutsideDirected:  Float64(0.09),
				Organized:        Float64(0.02),
			},
			Calc: func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool {
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
		{
			Name: "IsImmortal",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.05),
				OutsideDirected:  Float64(0.05),
			},
			Calc: func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool {
				baseCoef := 0.9
				if r.Type.IsMonotheism() {
					baseCoef = 1.1
				}
				if CalculateProbabilityFromReligionMetadata(baseCoef, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
		{
			Name: "IsTranscendental",
			_religionMetadata: &updateReligionMetadata{
				Elitaric: Float64(0.07),
			},
			Calc: func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool {
				var baseCoef float64 = 1
				if r.Type.IsMonotheism() || r.Type.IsDeism() || r.Type.IsDeityDualism() {
					baseCoef = 1.1
				}
				if CalculateProbabilityFromReligionMetadata(baseCoef, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
		{
			Name: "TakePartInHumanLife",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				OutsideDirected:  Float64(0.2),
				Fanaticism:       Float64(0.05),
			},
			Calc: func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool {
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
		{
			Name: "CommunicateWithHumans",
			_religionMetadata: &updateReligionMetadata{
				Centralized:      Float64(0.05),
				RealLifeOriented: Float64(0.1),
				OutsideDirected:  Float64(0.2),
				Fanaticism:       Float64(0.05),
			},
			Calc: func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool {
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
		{
			Name: "IsGnosticDeity", // God(s) created spiritual world but Demiurge(s) created physical world. Body is sinful
			_religionMetadata: &updateReligionMetadata{
				Decentralized:     Float64(0.03),
				AfterlifeOriented: Float64(0.03),
				Chthonic:          Float64(0.01),
				Elitaric:          Float64(0.02),
			},
			Calc: func(r *Religion, self *deityNatureTrait, selectedTraits []*deityNatureTrait) bool {
				var baseCoef float64 = 1
				if r.Type.IsDeityDualism() {
					baseCoef += 0.1
				}
				if dd.Goodness.Goodness == Evil {
					baseCoef += 0.1
				}
				if CalculateProbabilityFromReligionMetadata(baseCoef, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
	}
}

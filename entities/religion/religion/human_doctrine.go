package religion

import "fmt"

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
	fmt.Printf("Human Doctrine (religion=%s)\n", hd.religion.Name)
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
	hn.Traits = hn.generateTraits()

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
		rmCoef  = 0.1
		good    = 0.1 + rmCoef*hn.religion.metadata.Hedonism
		neutral = 0.1 + rmCoef*hn.religion.metadata.Hedonism
		evil    = 0.1 + rmCoef*hn.religion.metadata.Chthonic
	)
	for _, goal := range hn.humanDoctrine.doctrine.HighGoal.Goals {
		switch goal.Name {
		case "ProduceChildren":
			good += 0.05
		case "InvestigateMyself":
			good += 0.05
		}
	}

	return getGoodnessByProbability(good, neutral, evil)
}

func (hn *HumanNature) getGoodnessLevelByReligionMetadata() Level {
	var (
		major  = 0.1
		middle = 0.1
		minor  = 0.1
	)
	return getLevelByProbability(major, middle, minor)
}

func (hn *HumanNature) generateTraits() []*humanNatureTrait {
	allTraits := hn.getAllHumanNatureTraits()
	traits := make([]*humanNatureTrait, 0, len(allTraits))
	for i, trait := range allTraits {
		if trait.Calc(hn.religion, trait, traits) {
			traits = append(traits, &humanNatureTrait{
				_religionMetadata: trait._religionMetadata,
				Index:             i,
				Name:              trait.Name,
				Calc:              trait.Calc,
			})
		}
	}

	return traits
}

type humanNatureTrait struct {
	_religionMetadata *updateReligionMetadata
	Index             int
	Name              string
	Calc              func(r *Religion, self *humanNatureTrait, selectedTraits []*humanNatureTrait) bool
}

func (hn *HumanNature) getAllHumanNatureTraits() []*humanNatureTrait {
	return []*humanNatureTrait{
		{
			Name: "HasSoul",
			_religionMetadata: &updateReligionMetadata{
				AfterlifeOriented: Float64(0.09),
				InsideDirected:    Float64(0.09),
				Strictness:        Float64(0.01),
			},
			Calc: func(r *Religion, self *humanNatureTrait, _ []*humanNatureTrait) bool {
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
		{
			Name: "CanBeSaint",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.05),
				InsideDirected:   Float64(0.06),
				Strictness:       Float64(0.08),
				Ascetic:          Float64(0.08),
			},
			Calc: func(r *Religion, self *humanNatureTrait, selectedTraits []*humanNatureTrait) bool {
				var baseCoef float64 = 1
				for _, goal := range hn.humanDoctrine.doctrine.HighGoal.Goals {
					if goal.Name == "BecomePerfectAndSaints" {
						baseCoef += 0.2
					}
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

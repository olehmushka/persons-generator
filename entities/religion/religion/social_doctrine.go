package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
)

type SocialDoctrine struct {
	religion *Religion
	doctrine *Doctrine

	Traits []*socialTrait
}

func (d *Doctrine) generateSocialDoctrine() *SocialDoctrine {
	sd := &SocialDoctrine{religion: d.religion, doctrine: d}
	sd.Traits = sd.generateTraits(1, 5)

	return sd
}

func (sd *SocialDoctrine) Print() {
	fmt.Printf("Social Traits (religion_name=%s):\n", sd.religion.Name)
	for _, trait := range sd.Traits {
		fmt.Printf(" - %s\n", trait.Name)
	}
}

func (sd *SocialDoctrine) generateTraits(min, max int) []*socialTrait {
	if min < 0 {
		panic("[SocialDoctrine.generateTraits] min can not be less than 0")
	}
	allTraits := sd.getAllSocialTraits()
	if max > len(allTraits) {
		panic("[SocialDoctrine.generateTraits] max can not be greater than allTraits length")
	}

	traits := make([]*socialTrait, 0, len(allTraits))
	for count := 0; count < 20; count++ {
		for i, trait := range allTraits {
			if trait.Calc(sd.religion, trait, traits) {
				traits = append(traits, &socialTrait{
					_religionMetadata: trait._religionMetadata,
					Index:             i,
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
		sd.religion.UpdateMetadata(UpdateReligionMetadata(*sd.religion.metadata, *trait._religionMetadata))
	}

	return traits
}

type socialTrait struct {
	_religionMetadata *updateReligionMetadata
	baseCoef          float64
	Index             int
	Name              string
	Calc              func(r *Religion, self *socialTrait, selectedTraits []*socialTrait) bool
}

func (sd *SocialDoctrine) getAllSocialTraits() []*socialTrait {
	return []*socialTrait{
		{
			Name: "Purity",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				InsideDirected:   Float64(0.02),
				Strictness:       Float64(0.05),
				Ascetic:          Float64(0.5),
				Elitaric:         Float64(0.09),
				Organized:        Float64(0.08),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "LiveIsSacred",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				OutsideDirected:  Float64(0.1),
				Strictness:       Float64(0.01),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				var addCoef float64
				for _, goal := range sd.doctrine.HighGoal.Goals {
					switch goal.Name {
					case "LovePeople":
						addCoef += pm.RandFloat64InRange(0.03, 0.1)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SanctityOfNature",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				OutsideDirected:  Float64(0.1),
				Strictness:       Float64(0.01),
				Primitive:        Float64(0.03),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SacredChildbirth",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.08),
				OutsideDirected:  Float64(0.08),
				Primitive:        Float64(0.03),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				var addCoef float64
				for _, goal := range sd.doctrine.HighGoal.Goals {
					switch goal.Name {
					case "ProduceChildren":
						addCoef += pm.RandFloat64InRange(0.075, 0.1)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Karma",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.05),
				OutsideDirected:  Float64(0.07),
				Strictness:       Float64(0.04),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				var addCoef float64
				if r.HasReincarnation() {
					addCoef += pm.RandFloat64InRange(0.01, 0.125)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Polyamory",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.05),
				OutsideDirected:  Float64(0.05),
				Hedonism:         Float64(0.1),
				Chthonic:         Float64(0.015),
				Primitive:        Float64(0.015),
			},
			baseCoef: pm.RandFloat64InRange(0.7, 1),
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{Log: true})
			},
		},
		{
			Name: "BadThingForGoodPurpose",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.06),
				OutsideDirected:  Float64(0.06),
			},
			baseCoef: pm.RandFloat64InRange(0.7, 1),
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				var addCoef float64
				for _, trait := range sd.doctrine.Deity.Nature.Traits {
					switch trait.Name {
					case "IsJust":
						addCoef -= pm.RandFloat64InRange(0.01, 0.1)
					}
				}

				return CalculateProbabilityFromReligionMetadata(pm.PrepareProbability(self.baseCoef+addCoef), r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Raider",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				OutsideDirected:  Float64(0.1),
				Chthonic:         Float64(0.1),
				Primitive:        Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.7, 1),
			Calc: func(r *Religion, self *socialTrait, allTraits []*socialTrait) bool {
				var addCoef float64
				for _, goal := range sd.doctrine.HighGoal.Goals {
					switch goal.Name {
					case "LovePeople":
						addCoef -= pm.RandFloat64InRange(0.05, 0.1)
					case "FightForEvil":
						addCoef += pm.RandFloat64InRange(0.075, 0.1)
					}
				}
				for _, trait := range allTraits {
					if trait.Name == "Karma" {
						addCoef -= pm.RandFloat64InRange(0.05, 0.1)
					}
				}

				return CalculateProbabilityFromReligionMetadata(pm.PrepareProbability(self.baseCoef+addCoef), r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Agoge",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				OutsideDirected:  Float64(0.1),
				Elitaric:         Float64(0.01),
				Organized:        Float64(0.065),
			},
			baseCoef: pm.RandFloat64InRange(0.85, 1.1),
			Calc: func(r *Religion, self *socialTrait, allTraits []*socialTrait) bool {
				var addCoef float64
				if !r.IsPacifistic() && r.IsAggressive() {
					addCoef += pm.RandFloat64InRange(0.015, 0.065)
				}

				return CalculateProbabilityFromReligionMetadata(pm.PrepareProbability(self.baseCoef+addCoef), r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Berserkers",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				OutsideDirected:  Float64(0.1),
				Fanaticism:       Float64(0.085),
				Chthonic:         Float64(0.085),
			},
			baseCoef: pm.RandFloat64InRange(0.85, 1.1),
			Calc: func(r *Religion, self *socialTrait, allTraits []*socialTrait) bool {
				var addCoef float64
				if !r.IsPacifistic() && r.IsAggressive() {
					addCoef += pm.RandFloat64InRange(0.03, 0.07)
				}

				return CalculateProbabilityFromReligionMetadata(pm.PrepareProbability(self.baseCoef+addCoef), r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HonorableDeath",
			_religionMetadata: &updateReligionMetadata{
				AfterlifeOriented: Float64(0.1),
				Fanaticism:        Float64(0.1),
				Chthonic:          Float64(0.09),
			},
			baseCoef: pm.RandFloat64InRange(0.85, 1.1),
			Calc: func(r *Religion, self *socialTrait, allTraits []*socialTrait) bool {
				var addCoef float64
				if r.IsAggressive() {
					addCoef += pm.RandFloat64InRange(0.01, 0.03)
				}

				return CalculateProbabilityFromReligionMetadata(pm.PrepareProbability(self.baseCoef+addCoef), r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "DeedOfExpiation",
			_religionMetadata: &updateReligionMetadata{
				AfterlifeOriented: Float64(0.1),
				Fanaticism:        Float64(0.1),
				Chthonic:          Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.85, 1.1),
			Calc: func(r *Religion, self *socialTrait, allTraits []*socialTrait) bool {
				return CalculateProbabilityFromReligionMetadata(pm.PrepareProbability(self.baseCoef), r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "WarriorsPath",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.095),
				Fanaticism:       Float64(0.1),
				Chthonic:         Float64(0.01),
				Elitaric:         Float64(0.025),
			},
			baseCoef: pm.RandFloat64InRange(0.85, 1.1),
			Calc: func(r *Religion, self *socialTrait, allTraits []*socialTrait) bool {
				var addCoef float64
				if !r.IsPacifistic() && r.IsAggressive() {
					addCoef += pm.RandFloat64InRange(0.035, 0.075)
				}

				return CalculateProbabilityFromReligionMetadata(pm.PrepareProbability(self.baseCoef+addCoef), r, *self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (sd *SocialDoctrine) GetNaturalisticCriterias() float64 {
	if len(sd.Traits) == 0 {
		return 0
	}

	var criterias float64
	for _, trait := range sd.Traits {
		switch trait.Name {
		case "SanctityOfNature":
			criterias += 1
		case "SacredChildbirth":
			criterias += 0.5
		}
	}

	return criterias
}

func (sd *SocialDoctrine) GetPacifisticCriterias() float64 {
	if len(sd.Traits) == 0 {
		return 0
	}

	var criterias float64
	for _, trait := range sd.Traits {
		switch trait.Name {
		case "LiveIsSacred":
			criterias += 1
		case "SacredChildbirth":
			criterias += 0.5
		}
	}

	return criterias
}

func (sd *SocialDoctrine) GetAggressiveCriterias() float64 {
	if len(sd.Traits) == 0 {
		return 0
	}

	var criterias float64
	for _, trait := range sd.Traits {
		switch trait.Name {
		case "Raider":
			fallthrough
		case "Agoge":
			fallthrough
		case "Berserkers":
			fallthrough
		case "HonorableDeath":
			fallthrough
		case "WarriorsPath":
			criterias += 1
		}
	}

	return criterias
}

func (sd *SocialDoctrine) GetSexualActiveCriterias() float64 {
	if len(sd.Traits) == 0 {
		return 0
	}

	var criterias float64
	for _, trait := range sd.Traits {
		switch trait.Name {
		case "Polyamory":
			criterias += 1
		}
	}

	return criterias
}

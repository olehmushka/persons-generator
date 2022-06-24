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
		for _, trait := range allTraits {
			if trait.Calc(sd.religion, trait, traits) {
				traits = append(traits, &socialTrait{
					_religionMetadata: trait._religionMetadata,
					baseCoef:          trait.baseCoef,
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
	_religionMetadata *religionMetadata
	baseCoef          float64
	Name              string
	Calc              func(r *Religion, self *socialTrait, selectedTraits []*socialTrait) bool
}

func (sd *SocialDoctrine) getAllSocialTraits() []*socialTrait {
	return []*socialTrait{
		{
			Name: "Purity",
			_religionMetadata: &religionMetadata{
				Lawful:          0.25,
				Ascetic:         0.25,
				Authoritaristic: 0.5,
				Individualistic: 1,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "LiveIsSacred",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Altruistic:   0.75,
				Pacifistic:   1,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				var addCoef float64
				for _, goal := range sd.doctrine.HighGoal.Goals {
					switch goal.Name {
					case "LovePeople":
						addCoef += pm.RandFloat64InRange(0.03, 0.1)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SanctityOfNature",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.25,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SacredChildbirth",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				SexualActive: 0.25,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				var addCoef float64
				for _, goal := range sd.doctrine.HighGoal.Goals {
					switch goal.Name {
					case "ProduceChildren":
						addCoef += pm.RandFloat64InRange(0.1, 0.25)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Karma",
			_religionMetadata: &religionMetadata{
				Pacifistic:  0.5,
				Complicated: 0.5,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				var addCoef float64
				if r.HasReincarnation() {
					addCoef += pm.RandFloat64InRange(0.1, 0.25)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Polyamory",
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Hedonistic:   0.75,
				Liberal:      0.5,
			},
			baseCoef: sd.religion.M.BaseCoef - pm.RandFloat64InRange(0, 0.15),
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{Log: true})
			},
		},
		{
			Name: "BadThingForGoodPurpose",
			_religionMetadata: &religionMetadata{
				Chthonic: 0.25,
				Liberal:  1,
			},
			baseCoef: sd.religion.M.BaseCoef - pm.RandFloat64InRange(0, 0.15),
			Calc: func(r *Religion, self *socialTrait, _ []*socialTrait) bool {
				var addCoef float64
				for _, trait := range sd.doctrine.Deity.Nature.Traits {
					switch trait.Name {
					case "IsJust":
						addCoef -= pm.RandFloat64InRange(0.01, 0.1)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Raider",
			_religionMetadata: &religionMetadata{
				Chthonic:   1,
				Aggressive: 1,
			},
			baseCoef: sd.religion.M.LowBaseCoef,
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

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Agoge",
			_religionMetadata: &religionMetadata{
				Aggressive:      1,
				Authoritaristic: 0.5,
				Collectivistic:  0.5,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *socialTrait, allTraits []*socialTrait) bool {
				var addCoef float64
				if !r.metadata.IsPacifistic() && r.metadata.IsAggressive() {
					addCoef += pm.RandFloat64InRange(0.05, 0.15)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Berserkers",
			_religionMetadata: &religionMetadata{
				Chthonic:       0.5,
				Aggressive:     1,
				Collectivistic: 0.5,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *socialTrait, allTraits []*socialTrait) bool {
				var addCoef float64
				if !r.metadata.IsPacifistic() && r.metadata.IsAggressive() {
					addCoef += pm.RandFloat64InRange(0.05, 0.15)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HonorableDeath",
			_religionMetadata: &religionMetadata{
				Chthonic:        0.5,
				Aggressive:      1,
				Individualistic: 0.25,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *socialTrait, allTraits []*socialTrait) bool {
				var addCoef float64
				if r.metadata.IsAggressive() {
					addCoef += pm.RandFloat64InRange(0.05, 0.15)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "DeedOfExpiation",
			_religionMetadata: &religionMetadata{
				Chthonic:        0.5,
				Aggressive:      1,
				Individualistic: 0.25,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *socialTrait, allTraits []*socialTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "WarriorsPath",
			_religionMetadata: &religionMetadata{
				Chthonic:        0.25,
				Aggressive:      1,
				Individualistic: 0.5,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *socialTrait, allTraits []*socialTrait) bool {
				var addCoef float64
				if !r.metadata.IsPacifistic() && r.metadata.IsAggressive() {
					addCoef += pm.RandFloat64InRange(0.05, 0.15)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Kalokagathos",
			_religionMetadata: &religionMetadata{
				Ascetic:         0.25,
				Authoritaristic: 0.5,
				Individualistic: 1,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *socialTrait, selectedTraits []*socialTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
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

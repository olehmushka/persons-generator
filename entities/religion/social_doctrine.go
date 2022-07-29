package religion

import (
	"fmt"

	pm "persons_generator/probability_machine"
)

type SocialDoctrine struct {
	religion *Religion
	doctrine *Doctrine

	Traits []*trait
}

func (d *Doctrine) generateSocialDoctrine() *SocialDoctrine {
	sd := &SocialDoctrine{religion: d.religion, doctrine: d}
	sd.Traits = generateTraits(sd.religion, sd.getAllSocialTraits(), generateTraitsOpts{
		Label: "SocialDoctrine.generateTraits",
		Min:   1,
		Max:   5,
	})

	return sd
}

func (sd *SocialDoctrine) Print() {
	if len(sd.Traits) > 0 {
		fmt.Printf("Social Traits (religion_name=%s):\n", sd.religion.Name)
		for _, trait := range sd.Traits {
			fmt.Printf(" - %s\n", trait.Name)
		}
		return
	}
	fmt.Printf("Has not SocialDoctrines (religion_name=%s):\n", sd.religion.Name)
}

func (sd *SocialDoctrine) getAllSocialTraits() []*trait {
	return []*trait{
		{
			Name: "Purity",
			_religionMetadata: &religionMetadata{
				Lawful:          0.25,
				Ascetic:         0.25,
				Authoritaristic: 0.5,
				Individualistic: 1,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
				Complicated: 0.75,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			baseCoef: sd.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "BadThingForGoodPurpose",
			_religionMetadata: &religionMetadata{
				Chthonic: 0.25,
				Liberal:  1,
			},
			baseCoef: sd.religion.M.BaseCoef - pm.RandFloat64InRange(0, 0.15),
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, allTraits []*trait) bool {
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
			Calc: func(r *Religion, self *trait, allTraits []*trait) bool {
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
			Calc: func(r *Religion, self *trait, allTraits []*trait) bool {
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
			Calc: func(r *Religion, self *trait, allTraits []*trait) bool {
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
			Calc: func(r *Religion, self *trait, allTraits []*trait) bool {
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
			Calc: func(r *Religion, self *trait, allTraits []*trait) bool {
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
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

package religion

import (
	"fmt"

	pm "persons_generator/engine/probability_machine"
)

type SocialDoctrine struct {
	religion *Religion
	doctrine *Doctrine

	Traits []*trait
}

func (d *Doctrine) generateSocialDoctrine() (*SocialDoctrine, error) {
	sd := &SocialDoctrine{religion: d.religion, doctrine: d}
	ts, err := generateTraits(sd.religion, sd.getAllSocialTraits(), generateTraitsOpts{
		Label: "SocialDoctrine.generateTraits",
		Min:   1,
		Max:   5,
	})
	if err != nil {
		return nil, err
	}
	sd.Traits = ts

	return sd, nil
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
			Name: PuritySocialTrait,
			_religionMetadata: &religionMetadata{
				Lawful:          0.25,
				Ascetic:         0.25,
				Authoritaristic: 0.5,
				Individualistic: 1,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: LiveIsSacredSocialTrait,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Altruistic:   0.75,
				Pacifistic:   1,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				for _, goal := range sd.doctrine.HighGoal.Goals {
					switch goal.Name {
					case LovePeopleHighGoal:
						coef, err := pm.RandFloat64InRange(0.03, 0.1)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SanctityOfNatureSocialTrait,
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.25,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SacredChildbirthSocialTrait,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				SexualActive: 0.25,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				for _, goal := range sd.doctrine.HighGoal.Goals {
					switch goal.Name {
					case ProduceChildrenHighGoal:
						coef, err := pm.RandFloat64InRange(0.1, 0.25)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: KarmaSocialTrait,
			_religionMetadata: &religionMetadata{
				Pacifistic:  0.5,
				Complicated: 0.75,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.HasReincarnation() {
					coef, err := pm.RandFloat64InRange(0.1, 0.25)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: PolyamorySocialTrait,
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Hedonistic:   0.75,
				Liberal:      0.5,
			},
			baseCoef: sd.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BadThingForGoodPurposeSocialTrait,
			_religionMetadata: &religionMetadata{
				Chthonic: 0.25,
				Liberal:  1,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				for _, trait := range sd.doctrine.Deity.Nature.Traits {
					switch trait.Name {
					case IsJustDeityTrait:
						coef, err := pm.RandFloat64InRange(0.01, 0.1)
						if err != nil {
							return false, err
						}
						addCoef -= coef
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: RaiderSocialTrait,
			_religionMetadata: &religionMetadata{
				Chthonic:   1,
				Aggressive: 1,
			},
			baseCoef: sd.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, allTraits []*trait) (bool, error) {
				var addCoef float64
				for _, goal := range sd.doctrine.HighGoal.Goals {
					switch goal.Name {
					case LovePeopleHighGoal:
						coef, err := pm.RandFloat64InRange(0.05, 0.1)
						if err != nil {
							return false, err
						}
						addCoef -= coef
					case FightForEvilHighGoal:
						coef, err := pm.RandFloat64InRange(0.075, 0.1)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}
				for _, trait := range allTraits {
					if trait.Name == KarmaSocialTrait {
						coef, err := pm.RandFloat64InRange(0.05, 0.1)
						if err != nil {
							return false, err
						}
						addCoef -= coef
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: AgogeSocialTrait,
			_religionMetadata: &religionMetadata{
				Aggressive:      1,
				Authoritaristic: 0.5,
				Collectivistic:  0.5,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, allTraits []*trait) (bool, error) {
				var addCoef float64
				if !r.metadata.IsPacifistic() && r.metadata.IsAggressive() {
					coef, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BerserkersSocialTrait,
			_religionMetadata: &religionMetadata{
				Chthonic:       0.5,
				Aggressive:     1,
				Collectivistic: 0.5,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, allTraits []*trait) (bool, error) {
				var addCoef float64
				if !r.metadata.IsPacifistic() && r.metadata.IsAggressive() {
					coef, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HonorableDeathSocialTrait,
			_religionMetadata: &religionMetadata{
				Chthonic:        0.5,
				Aggressive:      1,
				Individualistic: 0.25,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, allTraits []*trait) (bool, error) {
				var addCoef float64
				if r.metadata.IsAggressive() {
					coef, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DeedOfExpiationSocialTrait,
			_religionMetadata: &religionMetadata{
				Chthonic:        0.5,
				Aggressive:      1,
				Individualistic: 0.25,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, allTraits []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: WarriorsPathSocialTrait,
			_religionMetadata: &religionMetadata{
				Chthonic:        0.25,
				Aggressive:      1,
				Individualistic: 0.5,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, allTraits []*trait) (bool, error) {
				var addCoef float64
				if !r.metadata.IsPacifistic() && r.metadata.IsAggressive() {
					coef, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: KalokagathosSocialTrait,
			_religionMetadata: &religionMetadata{
				Ascetic:         0.25,
				Authoritaristic: 0.5,
				Individualistic: 1,
			},
			baseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

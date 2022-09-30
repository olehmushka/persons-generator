package religion

import (
	"fmt"

	pm "persons_generator/engine/probability_machine"
)

type SocialDoctrine struct {
	religion *Religion `json:"-" bson:"-"`
	doctrine *Doctrine `json:"-" bson:"-"`

	Traits []*trait `json:"traits" bson:"traits"`
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

func (sd *SocialDoctrine) SerializeTraits() []string {
	if sd == nil {
		return []string{}
	}

	return extractTraitNames(sd.Traits)
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
			ReligionMetadata: &religionMetadata{
				Lawful:          0.25,
				Ascetic:         0.25,
				Authoritaristic: 0.5,
				Individualistic: 1,
			},
			BaseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: LiveIsSacredSocialTrait,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Altruistic:   0.75,
				Pacifistic:   1,
			},
			BaseCoef: sd.religion.M.BaseCoef,
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

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SanctityOfNatureSocialTrait,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.25,
			},
			BaseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SacredChildbirthSocialTrait,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				SexualActive: 0.25,
			},
			BaseCoef: sd.religion.M.BaseCoef,
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

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: KarmaSocialTrait,
			ReligionMetadata: &religionMetadata{
				Pacifistic:  0.5,
				Complicated: 0.75,
			},
			BaseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.HasReincarnation() {
					coef, err := pm.RandFloat64InRange(0.1, 0.25)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: PolyamorySocialTrait,
			ReligionMetadata: &religionMetadata{
				SexualActive: 1,
				Hedonistic:   0.75,
				Liberal:      0.5,
			},
			BaseCoef: sd.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BadThingForGoodPurposeSocialTrait,
			ReligionMetadata: &religionMetadata{
				Chthonic: 0.25,
				Liberal:  1,
			},
			BaseCoef: sd.religion.M.BaseCoef,
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

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: RaiderSocialTrait,
			ReligionMetadata: &religionMetadata{
				Chthonic:   1,
				Aggressive: 1,
			},
			BaseCoef: sd.religion.M.LowBaseCoef,
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

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: AgogeSocialTrait,
			ReligionMetadata: &religionMetadata{
				Aggressive:      1,
				Authoritaristic: 0.5,
				Collectivistic:  0.5,
			},
			BaseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, allTraits []*trait) (bool, error) {
				var addCoef float64
				if !r.Metadata.IsPacifistic() && r.Metadata.IsAggressive() {
					coef, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BerserkersSocialTrait,
			ReligionMetadata: &religionMetadata{
				Chthonic:       0.5,
				Aggressive:     1,
				Collectivistic: 0.5,
			},
			BaseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, allTraits []*trait) (bool, error) {
				var addCoef float64
				if !r.Metadata.IsPacifistic() && r.Metadata.IsAggressive() {
					coef, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HonorableDeathSocialTrait,
			ReligionMetadata: &religionMetadata{
				Chthonic:        0.5,
				Aggressive:      1,
				Individualistic: 0.25,
			},
			BaseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, allTraits []*trait) (bool, error) {
				var addCoef float64
				if r.Metadata.IsAggressive() {
					coef, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DeedOfExpiationSocialTrait,
			ReligionMetadata: &religionMetadata{
				Chthonic:        0.5,
				Aggressive:      1,
				Individualistic: 0.25,
			},
			BaseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, allTraits []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: WarriorsPathSocialTrait,
			ReligionMetadata: &religionMetadata{
				Chthonic:        0.25,
				Aggressive:      1,
				Individualistic: 0.5,
			},
			BaseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, allTraits []*trait) (bool, error) {
				var addCoef float64
				if !r.Metadata.IsPacifistic() && r.Metadata.IsAggressive() {
					coef, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: KalokagathosSocialTrait,
			ReligionMetadata: &religionMetadata{
				Ascetic:         0.25,
				Authoritaristic: 0.5,
				Individualistic: 1,
			},
			BaseCoef: sd.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func GetSocialDoctrineSimilarityCoef(d1, d2 *SocialDoctrine) float64 {
	if d1 == nil && d2 == nil {
		return 1
	}
	if d1 == nil || d2 == nil {
		return 0
	}

	return GetTraitsSimilarityCoef(d1.Traits, d2.Traits)
}

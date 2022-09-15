package religion

import (
	"fmt"
)

type HighGoal struct {
	religion *Religion `json:"-"`

	Goals []*trait `json:"goals"`
}

func (d *Doctrine) generateHighGoal() (*HighGoal, error) {
	hg := &HighGoal{religion: d.religion}
	gs, err := generateTraits(hg.religion, hg.getAllGoals(), generateTraitsOpts{
		Label: "HighGoal.generateGoals",
		Min:   1,
		Max:   3,
	})
	if err != nil {
		return nil, err
	}
	hg.Goals = gs

	return hg, nil
}

func (hg *HighGoal) Print() {
	if len(hg.Goals) > 0 {
		fmt.Printf("HighGoals (religion_name=%s):\n", hg.religion.Name)
		for _, goal := range hg.Goals {
			fmt.Printf(" - %s\n", goal.Name)
		}
		return
	}
	fmt.Printf("Has not High Goals (religion_name=%s):\n", hg.religion.Name)
}

func (hg *HighGoal) ContainsReincarnation() bool {
	for _, goal := range hg.Goals {
		switch goal.Name {
		case StopReincarnationHighGoal:
			fallthrough
		case NeverStopReincarnationHighGoal:
			return true
		}
	}
	return false
}

func (hg *HighGoal) getAllGoals() []*trait {
	return []*trait{
		{
			Name: TransformingIntoLikenessOfGodHighGoal,
			ReligionMetadata: &religionMetadata{
				Individualistic: 1,
				Ascetic:         0.5,
			},
			BaseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if !r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: AttainHeavenAndResidesThereWithGodHighGoal,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Ascetic:      0.1,
			},
			BaseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BringManToCompletenessAndToCloseRelationshipWithGodHighGoal,
			ReligionMetadata: &religionMetadata{
				Naturalistic:    0.1,
				Individualistic: 1,
			},
			BaseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if !r.Type.IsMonotheism() && !r.Type.IsDeityDualism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BringHolinessDownToTheWorldHighGoal,
			ReligionMetadata: &religionMetadata{
				Altruistic:     0.5,
				Collectivistic: 0.5,
			},
			BaseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: StopReincarnationHighGoal,
			ReligionMetadata: &religionMetadata{
				Ascetic:         0.5,
				Individualistic: 1,
			},
			BaseCoef: hg.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) (bool, error) {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == NeverStopReincarnationHighGoal {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: NeverStopReincarnationHighGoal,
			ReligionMetadata: &religionMetadata{
				Ascetic:         0.5,
				Individualistic: 1,
			},
			BaseCoef: hg.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) (bool, error) {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == StopReincarnationHighGoal {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ProduceChildrenHighGoal,
			ReligionMetadata: &religionMetadata{
				Naturalistic:   1,
				SexualActive:   0.75,
				Collectivistic: 0.1,
			},
			BaseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: InvestigateMyselfHighGoal,
			ReligionMetadata: &religionMetadata{
				Individualistic: 1,
			},
			BaseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GetMaxPleasureHighGoal,
			ReligionMetadata: &religionMetadata{
				SexualActive: 0.25,
				Liberal:      0.25,
				Hedonistic:   1,
			},
			BaseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Metadata.IsAscetic() && !r.Metadata.IsHedonistic() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SpreadReligionHighGoal,
			ReligionMetadata: &religionMetadata{
				Aggressive:     0.25,
				Collectivistic: 1,
			},
			BaseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: LovePeopleHighGoal,
			ReligionMetadata: &religionMetadata{
				SexualActive: 0.1,
				Altruistic:   1,
				Pacifistic:   0.25,
			},
			BaseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BecomePerfectAndSaintsHighGoal,
			ReligionMetadata: &religionMetadata{
				Ascetic:         1,
				Individualistic: 1,
			},
			BaseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: FightAgainstEvilHighGoal,
			ReligionMetadata: &religionMetadata{
				Lawful:     0.1,
				Aggressive: 0.75,
			},
			BaseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) (bool, error) {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == FightForEvilHighGoal {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: FightForEvilHighGoal,
			ReligionMetadata: &religionMetadata{
				Chthonic:   1,
				Aggressive: 0.75,
			},
			BaseCoef: hg.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) (bool, error) {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == FightAgainstEvilHighGoal {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: StopArmageddonHighGoal,
			ReligionMetadata: &religionMetadata{
				Altruistic: 0.5,
			},
			BaseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) (bool, error) {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == StartArmageddonHighGoal {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: StartArmageddonHighGoal,
			ReligionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			BaseCoef: hg.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) (bool, error) {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == StopArmageddonHighGoal {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func GetHighGoalSimilarityCoef(h1, h2 *HighGoal) float64 {
	if h1 == nil && h2 == nil {
		return 1
	}
	if h1 == nil || h2 == nil {
		return 0
	}

	return GetTraitsSimilarityCoef(h1.Goals, h2.Goals)
}

package religion

import (
	"fmt"
)

type HighGoal struct {
	religion *Religion

	Goals []*trait
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
			_religionMetadata: &religionMetadata{
				Individualistic: 1,
				Ascetic:         0.5,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if !r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: AttainHeavenAndResidesThereWithGodHighGoal,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Ascetic:      0.1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BringManToCompletenessAndToCloseRelationshipWithGodHighGoal,
			_religionMetadata: &religionMetadata{
				Naturalistic:    0.1,
				Individualistic: 1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if !r.Type.IsMonotheism() && !r.Type.IsDeityDualism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BringHolinessDownToTheWorldHighGoal,
			_religionMetadata: &religionMetadata{
				Altruistic:     0.5,
				Collectivistic: 0.5,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: StopReincarnationHighGoal,
			_religionMetadata: &religionMetadata{
				Ascetic:         0.5,
				Individualistic: 1,
			},
			baseCoef: hg.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) (bool, error) {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == NeverStopReincarnationHighGoal {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: NeverStopReincarnationHighGoal,
			_religionMetadata: &religionMetadata{
				Ascetic:         0.5,
				Individualistic: 1,
			},
			baseCoef: hg.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) (bool, error) {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == StopReincarnationHighGoal {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ProduceChildrenHighGoal,
			_religionMetadata: &religionMetadata{
				Naturalistic:   1,
				SexualActive:   0.75,
				Collectivistic: 0.1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: InvestigateMyselfHighGoal,
			_religionMetadata: &religionMetadata{
				Individualistic: 1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GetMaxPleasureHighGoal,
			_religionMetadata: &religionMetadata{
				SexualActive: 0.25,
				Liberal:      0.25,
				Hedonistic:   1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.metadata.IsAscetic() && !r.metadata.IsHedonistic() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SpreadReligionHighGoal,
			_religionMetadata: &religionMetadata{
				Aggressive:     0.25,
				Collectivistic: 1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: LovePeopleHighGoal,
			_religionMetadata: &religionMetadata{
				SexualActive: 0.1,
				Altruistic:   1,
				Pacifistic:   0.25,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BecomePerfectAndSaintsHighGoal,
			_religionMetadata: &religionMetadata{
				Ascetic:         1,
				Individualistic: 1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: FightAgainstEvilHighGoal,
			_religionMetadata: &religionMetadata{
				Lawful:     0.1,
				Aggressive: 0.75,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) (bool, error) {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == FightForEvilHighGoal {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: FightForEvilHighGoal,
			_religionMetadata: &religionMetadata{
				Chthonic:   1,
				Aggressive: 0.75,
			},
			baseCoef: hg.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) (bool, error) {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == FightAgainstEvilHighGoal {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: StopArmageddonHighGoal,
			_religionMetadata: &religionMetadata{
				Altruistic: 0.5,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) (bool, error) {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == StartArmageddonHighGoal {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: StartArmageddonHighGoal,
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: hg.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) (bool, error) {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == StopArmageddonHighGoal {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

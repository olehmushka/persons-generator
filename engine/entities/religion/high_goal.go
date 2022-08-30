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
		case "stop_reincarnation":
			return true
		case "never_stop_reincarnation":
			return true
		}
	}
	return false
}

func (hg *HighGoal) getAllGoals() []*trait {
	return []*trait{
		{
			Name: "transforming_into_likeness_of_god",
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
			Name: "attain_heaven_and_resides_there_with_god",
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
			Name: "bring_man_to_completeness_and_to_close_relationship_with_god",
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
			Name: "bring_holiness_down_to_the_world",
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
			Name: "stop_reincarnation",
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
					if goal.Name == "never_stop_reincarnation" {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "never_stop_reincarnation",
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
					if goal.Name == "stop_reincarnation" {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "produce_children",
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
			Name: "investigate_myself",
			_religionMetadata: &religionMetadata{
				Individualistic: 1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "get_max_pleasure",
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
			Name: "spread_religion",
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
			Name: "love_people",
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
			Name: "become_perfect_and_saints",
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
			Name: "fight_against_evil",
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
					if goal.Name == "fight_for_evil" {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "fight_for_evil",
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
					if goal.Name == "fight_against_evil" {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "stop_armageddon",
			_religionMetadata: &religionMetadata{
				Altruistic: 0.5,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) (bool, error) {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == "start_armageddon" {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "start_armageddon",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: hg.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) (bool, error) {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == "stop_armageddon" {
						return false, nil
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

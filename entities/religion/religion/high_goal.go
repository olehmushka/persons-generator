package religion

import (
	"fmt"
)

type HighGoal struct {
	religion *Religion

	Goals []*highGoal
}

func (d *Doctrine) generateHighGoal() *HighGoal {
	hg := &HighGoal{religion: d.religion}
	hg.Goals = hg.generateGoals(1, 3)

	return hg
}

func (hg *HighGoal) Print() {
	fmt.Printf("HighGoals (religion_name=%s):\n", hg.religion.Name)
	for _, goal := range hg.Goals {
		fmt.Printf(" - %s\n", goal.Name)
	}
}

func (hg *HighGoal) ContainsReincarnation() bool {
	for _, goal := range hg.Goals {
		switch goal.Name {
		case "StopReincarnation":
			return true
		case "NeverStopReincarnation":
			return true
		}
	}
	return false
}

type highGoal struct {
	_religionMetadata *religionMetadata
	baseCoef          float64

	Name string
	Calc func(r *Religion, self *highGoal, selectedGoals []*highGoal) bool
}

func (hg *HighGoal) generateGoals(min, max int) []*highGoal {
	if min < 0 {
		panic("[HighGoal.generateGoals] min can not be less than 0")
	}
	allGoals := hg.getAllGoals()
	if max > len(allGoals) {
		panic("[HighGoal.generateGoals] max can not be greater than allGoals length")
	}

	goals := make([]*highGoal, 0, len(allGoals))
	for count := 0; count < 20; count++ {
		for _, goal := range allGoals {
			if goal.Calc(hg.religion, goal, goals) {
				goals = append(goals, &highGoal{
					_religionMetadata: goal._religionMetadata,
					baseCoef:          goal.baseCoef,
					Name:              goal.Name,
					Calc:              goal.Calc,
				})
			}
			if len(goals) == max {
				break
			}
		}
		if len(goals) == max {
			break
		}
		if len(goals) >= min {
			break
		}
	}

	for _, goal := range goals {
		hg.religion.UpdateMetadata(UpdateReligionMetadata(*hg.religion.metadata, *goal._religionMetadata))
	}

	return goals
}

func (hg *HighGoal) getAllGoals() []*highGoal {
	return []*highGoal{
		{
			Name: "TransformingIntoLikenessOfGod",
			_religionMetadata: &religionMetadata{
				Individualistic: 1,
				Ascetic:         0.5,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				if !r.Type.IsMonotheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "AttainHeavenAndResidesThereWithGod",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Ascetic:      0.1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "BringManToCompletenessAndToCloseRelationshipWithGod",
			_religionMetadata: &religionMetadata{
				Naturalistic:    0.1,
				Individualistic: 1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				if !r.Type.IsMonotheism() && !r.Type.IsDeityDualism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "BringHolinessDownToTheWorld",
			_religionMetadata: &religionMetadata{
				Altruistic:     0.5,
				Collectivistic: 0.5,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "StopReincarnation",
			_religionMetadata: &religionMetadata{
				Ascetic:         0.5,
				Individualistic: 1,
			},
			baseCoef: hg.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *highGoal, selectedGoals []*highGoal) bool {
				for _, goal := range selectedGoals {
					if goal.Name == "NeverStopReincarnation" {
						return false
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "NeverStopReincarnation",
			_religionMetadata: &religionMetadata{
				Ascetic:         0.5,
				Individualistic: 1,
			},
			baseCoef: hg.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *highGoal, selectedGoals []*highGoal) bool {
				for _, goal := range selectedGoals {
					if goal.Name == "StopReincarnation" {
						return false
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "ProduceChildren",
			_religionMetadata: &religionMetadata{
				Naturalistic:   1,
				SexualActive:   0.75,
				Collectivistic: 0.1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "InvestigateMyself",
			_religionMetadata: &religionMetadata{
				Individualistic: 1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "GetMaxPleasure",
			_religionMetadata: &religionMetadata{
				SexualActive: 0.25,
				Liberal:      0.25,
				Hedonistic:   1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				if r.metadata.IsAscetic() && !r.metadata.IsHedonistic() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SpreadReligion",
			_religionMetadata: &religionMetadata{
				Aggressive:     0.25,
				Collectivistic: 1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "LovePeople",
			_religionMetadata: &religionMetadata{
				SexualActive: 0.1,
				Altruistic:   1,
				Pacifistic:   0.25,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "BecomePerfectAndSaints",
			_religionMetadata: &religionMetadata{
				Ascetic:         1,
				Individualistic: 1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "FightAgainstEvil",
			_religionMetadata: &religionMetadata{
				Lawful:     0.1,
				Aggressive: 0.75,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *highGoal, selectedGoals []*highGoal) bool {
				for _, goal := range selectedGoals {
					if goal.Name == "FightForEvil" {
						return false
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "FightForEvil",
			_religionMetadata: &religionMetadata{
				Chthonic:   1,
				Aggressive: 0.75,
			},
			baseCoef: hg.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *highGoal, selectedGoals []*highGoal) bool {
				for _, goal := range selectedGoals {
					if goal.Name == "FightAgainstEvil" {
						return false
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "StopArmageddon",
			_religionMetadata: &religionMetadata{
				Altruistic: 0.5,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *highGoal, selectedGoals []*highGoal) bool {
				for _, goal := range selectedGoals {
					if goal.Name == "StartArmageddon" {
						return false
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "StartArmageddon",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: hg.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *highGoal, selectedGoals []*highGoal) bool {
				for _, goal := range selectedGoals {
					if goal.Name == "StopArmageddon" {
						return false
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

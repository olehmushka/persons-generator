package religion

import (
	"fmt"
)

type HighGoal struct {
	religion *Religion

	Goals []*trait
}

func (d *Doctrine) generateHighGoal() *HighGoal {
	hg := &HighGoal{religion: d.religion}
	hg.Goals = generateTraits(hg.religion, hg.getAllGoals(), generateTraitsOpts{
		Label: "HighGoal.generateGoals",
		Min:   1,
		Max:   3,
	})

	return hg
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
		case "StopReincarnation":
			return true
		case "NeverStopReincarnation":
			return true
		}
	}
	return false
}

func (hg *HighGoal) getAllGoals() []*trait {
	return []*trait{
		{
			Name: "TransformingIntoLikenessOfGod",
			_religionMetadata: &religionMetadata{
				Individualistic: 1,
				Ascetic:         0.5,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) bool {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
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
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) bool {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "InvestigateMyself",
			_religionMetadata: &religionMetadata{
				Individualistic: 1,
			},
			baseCoef: hg.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) bool {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
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
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) bool {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
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
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) bool {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
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
			Calc: func(r *Religion, self *trait, selectedGoals []*trait) bool {
				for _, goal := range selectedGoals {
					if goal == nil {
						continue
					}
					if goal.Name == "StopArmageddon" {
						return false
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

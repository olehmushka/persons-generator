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
	_religionMetadata *updateReligionMetadata

	Index int
	Name  string
	Calc  func(r *Religion, self *highGoal, selectedGoals []*highGoal) bool
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
		for i, goal := range allGoals {
			if goal.Calc(hg.religion, goal, goals) {
				goals = append(goals, &highGoal{
					_religionMetadata: goal._religionMetadata,
					Index:             i,
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
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.05),
				InsideDirected:   Float64(0.09),
				Strictness:       Float64(0.05),
				Ascetic:          Float64(0.01),
				Elitaric:         Float64(0.01),
			},
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				if !r.Type.IsMonotheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "AttainHeavenAndResidesThereWithGod",
			_religionMetadata: &updateReligionMetadata{
				AfterlifeOriented: Float64(0.1),
				InsideDirected:    Float64(0.05),
				Fanaticism:        Float64(0.05),
				Strictness:        Float64(0.09),
			},
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "BringManToCompletenessAndToCloseRelationshipWithGod",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.09),
				InsideDirected:   Float64(0.08),
				Strictness:       Float64(0.05),
				Ascetic:          Float64(0.01),
			},
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				if !r.Type.IsMonotheism() && !r.Type.IsDeityDualism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "BringHolinessDownToTheWorld",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				OutsideDirected:  Float64(0.1),
				Fanaticism:       Float64(0.07),
				Strictness:       Float64(0.05),
				Ascetic:          Float64(0.03),
				Organized:        Float64(0.02),
			},
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				return CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "StopReincarnation",
			_religionMetadata: &updateReligionMetadata{
				AfterlifeOriented: Float64(0.1),
				Fanaticism:        Float64(0.02),
				Strictness:        Float64(0.08),
				Ascetic:           Float64(0.07),
			},
			Calc: func(r *Religion, self *highGoal, selectedGoals []*highGoal) bool {
				for _, goal := range selectedGoals {
					if goal.Name == "NeverStopReincarnation" {
						return false
					}
				}

				return CalculateProbabilityFromReligionMetadata(0.6, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "NeverStopReincarnation",
			_religionMetadata: &updateReligionMetadata{
				AfterlifeOriented: Float64(0.1),
				Fanaticism:        Float64(0.03),
				Strictness:        Float64(0.08),
				Ascetic:           Float64(0.07),
			},
			Calc: func(r *Religion, self *highGoal, selectedGoals []*highGoal) bool {
				for _, goal := range selectedGoals {
					if goal.Name == "StopReincarnation" {
						return false
					}
				}

				return CalculateProbabilityFromReligionMetadata(0.4, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "ProduceChildren",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				OutsideDirected:  Float64(0.1),
				Primitive:        Float64(0.02),
			},
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				return CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "InvestigateMyself",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				InsideDirected:   Float64(0.1),
				Hedonism:         Float64(0.05),
				Primitive:        Float64(0.03),
			},
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				return CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "GetMaxPleasure",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				InsideDirected:   Float64(0.05),
				Hedonism:         Float64(0.1),
				Primitive:        Float64(0.1),
			},
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				return CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SpreadReligion",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				OutsideDirected:  Float64(0.1),
				Fanaticism:       Float64(0.9),
				Organized:        Float64(0.6),
			},
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				return CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "LovePeople",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				OutsideDirected:  Float64(0.1),
				Primitive:        Float64(0.1),
			},
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				return CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "BecomePerfectAndSaints",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				InsideDirected:   Float64(0.1),
				Strictness:       Float64(0.1),
				Ascetic:          Float64(0.1),
			},
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				return CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "FightAgainstEvil",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.6),
				Fanaticism:       Float64(0.1),
				Primitive:        Float64(0.01),
			},
			Calc: func(r *Religion, self *highGoal, selectedGoals []*highGoal) bool {
				for _, goal := range selectedGoals {
					if goal.Name == "FightForEvil" {
						return false
					}
				}

				return CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "FightForEvil",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.07),
				Fanaticism:       Float64(0.1),
				Chthonic:         Float64(0.1),
				Primitive:        Float64(0.01),
			},
			Calc: func(r *Religion, self *highGoal, selectedGoals []*highGoal) bool {
				for _, goal := range selectedGoals {
					if goal.Name == "FightAgainstEvil" {
						return false
					}
				}

				return CalculateProbabilityFromReligionMetadata(0.5, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "StopArmageddon",
			_religionMetadata: &updateReligionMetadata{
				OutsideDirected: Float64(0.02),
				Fanaticism:      Float64(0.1),
				Chthonic:        Float64(0.01),
				Primitive:       Float64(0.01),
			},
			Calc: func(r *Religion, self *highGoal, _ []*highGoal) bool {
				return CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (hg *HighGoal) GetNaturalisticCriterias() float64 {
	if len(hg.Goals) == 0 {
		return 0
	}

	var criterias float64
	for _, goal := range hg.Goals {
		switch goal.Name {
		case "ProduceChildren":
			criterias += 0.5
		}
	}

	return criterias
}

func (hg *HighGoal) GetPacifisticCriterias() float64 {
	if len(hg.Goals) == 0 {
		return 0
	}

	var criterias float64
	for _, goal := range hg.Goals {
		switch goal.Name {
		case "LovePeople":
			criterias += 1
		}
	}

	return criterias
}

func (hg *HighGoal) GetAggressiveCriterias() float64 {
	if len(hg.Goals) == 0 {
		return 0
	}

	var criterias float64
	for _, goal := range hg.Goals {
		switch goal.Name {
		case "FightAgainstEvil":
			criterias += 0.5
		case "FightForEvil":
			criterias += 0.5
		}
	}

	return criterias
}

func (hg *HighGoal) GetSexualActiveCriterias() float64 {
	if len(hg.Goals) == 0 {
		return 0
	}

	var criterias float64
	for _, goal := range hg.Goals {
		switch goal.Name {
		case "ProduceChildren":
			fallthrough
		case "GetMaxPleasure":
			criterias += 0.5
		}
	}

	return criterias
}

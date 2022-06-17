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
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
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
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
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
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
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
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
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

				if CalculateProbabilityFromReligionMetadata(0.6, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
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

				if CalculateProbabilityFromReligionMetadata(0.4, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
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
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
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
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
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
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
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
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
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
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
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
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
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

				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
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

				if CalculateProbabilityFromReligionMetadata(0.5, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
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
				if CalculateProbabilityFromReligionMetadata(1, r, *self._religionMetadata, CalcProbOpts{}) {
					r.UpdateMetadata(UpdateReligionMetadata(*r.metadata, *self._religionMetadata))
					return true
				}

				return false
			},
		},
	}
}

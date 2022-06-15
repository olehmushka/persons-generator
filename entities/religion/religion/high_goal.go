package religion

import "fmt"

type HighGoal struct {
	religion *Religion

	Goals []*highGoal
}

func (d *Doctrine) generateHighGoal() *HighGoal {
	hg := &HighGoal{religion: d.religion}
	hg.Goals = hg.generateGoals()

	return hg
}

func (hg *HighGoal) Print() {
	fmt.Printf("HighGoals (religion_name=%s):\n", hg.religion.Name)
	for _, goal := range hg.Goals {
		fmt.Printf(" - %s\n", goal.Name)
	}
}

type highGoal struct {
	_religionMetadata *religionMetadata

	Index int
	Name  string
	Calc  func(r *Religion, self *highGoal) bool
}

func (hg *HighGoal) generateGoals() []*highGoal {
	allGoals := hg.getAllGoals()
	goals := make([]*highGoal, 0, len(allGoals))
	for i, goal := range allGoals {
		if goal.Calc(hg.religion, goal) {
			goals = append(goals, &highGoal{
				_religionMetadata: goal._religionMetadata,
				Index:             i,
				Calc:              goal.Calc,
			})
		}
	}

	return goals
}

func (hg *HighGoal) getAllGoals() []*highGoal {
	return []*highGoal{
		{
			Name: "TransformingIntoLikenessOfGod",
			_religionMetadata: &religionMetadata{
				RealLifeOriented:  0.2,
				AfterlifeOriented: 0.05,
				InsideDirected:    0.3,
				OutsideDirected:   0.2,
				Strictness:        0.05,
				Ascetic:           0.01,
				Elitaric:          0.01,
			},
			Calc: func(r *Religion, self *highGoal) bool {
				if !r.Type.IsMonotheism() {
					return false
				}

				return false
			},
		},
	}
}

/*
	HighGoal:
		TransformingIntoLikenessOfGod (only for monotheism)
		AttainHeavenAndResidesThereWithGod (only for monotheism)
		BringManToCompletenessAndToCloseRelationshipWithGod (only for monotheism)
		BringHolinessDownToTheWorld
		StopReincarnation (must have reincarnation)
		NeverStopReincarnation (must have reincarnation)
		ProduceChildren
		InvestigateMyself
		GetMaxPleasure
		SpreadReligion
		LovePeople
		BecomePerfectAndSaints
		FightWithEvil
		FightForEvil


*/

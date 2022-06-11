package religion

type HighGoal struct {
	religion *Religion

	goals []*highGoal
}

func (d *Doctrine) generateHighGoal() *HighGoal {
	hg := &HighGoal{religion: d.religion}
	hg.goals = hg.generateGoals()

	return hg
}

type highGoal struct {
	_religionMetadata *religionMetadata
	_metadata         *highCoalMetadata

	Index int
	Name  string
	Calc  func(r *Religion, self *highGoal) bool
}

type highCoalMetadata struct {
	IsCalculated bool
	Value        bool
}

func (hg *HighGoal) generateGoals() []*highGoal {
	return []*highGoal{}
}

func (hg *HighGoal) getAllGoals() []*highGoal {
	return []*highGoal{}
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

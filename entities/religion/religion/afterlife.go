package religion

import pm "persons_generator/probability-machine"

type Afterlife struct {
	religion *Religion

	IsExists     bool
	Participants *AfterlifeParticipants
}

func (d *Doctrine) generateAfterlife() *Afterlife {
	al := &Afterlife{religion: d.religion}

	return al
}

type AfterlifeParticipants struct {
	religion *Religion

	ForTopBelievers    AfterlifeOption
	ForBelievers       AfterlifeOption
	ForUntrueBelievers AfterlifeOption
	ForAtheists        AfterlifeOption
	ForAll             AfterlifeOption
}

func (al *Afterlife) generateAfterlifeParticipants() *AfterlifeParticipants {
	alp := &AfterlifeParticipants{religion: al.religion}

	return alp
}

type AfterlifeOption string

const (
	GoodAfterlife    AfterlifeOption = "Good"
	DependsAfterlife AfterlifeOption = "Depends"
	BadAfterlife     AfterlifeOption = "Bad"
)

func getAfterlifeOptionByProbability(good, depends, bad float64) AfterlifeOption {
	m := map[string]float64{
		string(GoodAfterlife):    pm.PrepareProbability(good),
		string(DependsAfterlife): pm.PrepareProbability(depends),
		string(BadAfterlife):     pm.PrepareProbability(bad),
	}
	return AfterlifeOption(pm.GetRandomFromSeveral(m))
}

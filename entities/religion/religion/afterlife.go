package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
)

type Afterlife struct {
	religion *Religion
	doctrine *Doctrine

	IsExists     bool
	Participants *AfterlifeParticipants
}

func (al *Afterlife) Print() {
	fmt.Printf("Afterlife Doctrine (religion=%s)\n", al.religion.Name)
	if !al.IsExists {
		fmt.Printf("Afterlife does not exists (religion=%s)\n", al.religion.Name)
		return
	}
	fmt.Printf("Afterlife exists (religion=%s)\n", al.religion.Name)
	al.Participants.Print()
}

func (d *Doctrine) generateAfterlife() *Afterlife {
	al := &Afterlife{religion: d.religion, doctrine: d}
	al.IsExists = al.generateIsExistsAfterlife()
	if al.IsExists {
		al.Participants = al.generateAfterlifeParticipants()
	}

	return al
}

func (al *Afterlife) generateIsExistsAfterlife() bool {
	if al.religion.Type.IsAtheism() {
		return false
	}

	var probability float64
	switch {
	case al.religion.Type.IsMonotheism():
		probability += pm.RandFloat64InRange(0.75, 1)
	case al.religion.Type.IsPolytheism():
		probability += pm.RandFloat64InRange(0.3, 1)
	case al.religion.Type.IsDeityDualism():
		probability += pm.RandFloat64InRange(0.5, 1)
	case al.religion.Type.IsDeism():
		probability += pm.RandFloat64InRange(0.05, 1)
	}
	probability += al.religion.metadata.AfterlifeOriented
	for _, trait := range al.doctrine.Human.Nature.Traits {
		if trait.Name == "HasSoul" {
			probability += pm.RandFloat64InRange(0.25, 0.75)
		}
	}

	return pm.GetRandomBool(probability)
}

type AfterlifeParticipants struct {
	religion  *Religion
	afterlife *Afterlife

	ForTopBelievers    AfterlifeOption
	ForBelievers       AfterlifeOption
	ForUntrueBelievers AfterlifeOption
	ForAtheists        AfterlifeOption
}

func (al *Afterlife) generateAfterlifeParticipants() *AfterlifeParticipants {
	alp := &AfterlifeParticipants{religion: al.religion, afterlife: al}
	alp.ForTopBelievers = al.generateForTopBelieversAfterlife(alp)
	alp.ForBelievers = al.generateForBelieversAfterlife(alp)
	alp.ForUntrueBelievers = al.generateForUntrueBelieversAfterlife(alp)
	alp.ForAtheists = al.generateForAtheistsAfterlife(alp)

	return alp
}

func (alp *AfterlifeParticipants) Print() {
	fmt.Printf("Afterlife for top believers is %s\n", alp.ForTopBelievers)
	fmt.Printf("Afterlife for believers is %s\n", alp.ForBelievers)
	fmt.Printf("Afterlife for untrue believers is %s\n", alp.ForUntrueBelievers)
	fmt.Printf("Afterlife for atheists is %s\n", alp.ForAtheists)
}

type AfterlifeOption string

const (
	GoodAfterlife    AfterlifeOption = "Good"
	DependsAfterlife AfterlifeOption = "Depends"
	BadAfterlife     AfterlifeOption = "Bad"
)

func (al *Afterlife) generateForTopBelieversAfterlife(alp *AfterlifeParticipants) AfterlifeOption {
	var (
		good    = pm.RandFloat64InRange(0.5, 1)
		depends = pm.RandFloat64InRange(0.3, 0.8)
		bad     = pm.RandFloat64InRange(0, 0.4)
	)

	return getAfterlifeOptionByProbability(good, depends, bad)
}

func (al *Afterlife) generateForBelieversAfterlife(alp *AfterlifeParticipants) AfterlifeOption {
	var (
		good    = pm.RandFloat64InRange(0.5, 0.9)
		depends = pm.RandFloat64InRange(0.3, 0.75)
		bad     = pm.RandFloat64InRange(0, 0.45)
	)

	return getAfterlifeOptionByProbability(good, depends, bad)
}

func (al *Afterlife) generateForUntrueBelieversAfterlife(alp *AfterlifeParticipants) AfterlifeOption {
	var (
		good    = pm.RandFloat64InRange(0.1, 0.5)
		depends = pm.RandFloat64InRange(0.1, 0.6)
		bad     = pm.RandFloat64InRange(0.1, 0.7)
	)
	if al.religion.Type.IsPolytheism() {
		if al.religion.Type.IsMonolatryPolytheism() {
			good += pm.RandFloat64InRange(0.01, 0.1)
		}
		if al.religion.Type.IsOmnismPolytheism() {
			good += pm.RandFloat64InRange(0.05, 0.2)
		}
	}

	return getAfterlifeOptionByProbability(good, depends, bad)
}

func (al *Afterlife) generateForAtheistsAfterlife(alp *AfterlifeParticipants) AfterlifeOption {
	var (
		good    = pm.RandFloat64InRange(0.05, 0.35)
		depends = pm.RandFloat64InRange(0.05, 0.45)
		bad     = pm.RandFloat64InRange(0.1, 0.8)
	)
	if al.religion.Type.IsDeism() {
		good += pm.RandFloat64InRange(0.01, 0.1)
		depends += pm.RandFloat64InRange(0.01, 0.1)
	}

	return getAfterlifeOptionByProbability(good, depends, bad)
}

func getAfterlifeOptionByProbability(good, depends, bad float64) AfterlifeOption {
	m := map[string]float64{
		string(GoodAfterlife):    pm.PrepareProbability(good),
		string(DependsAfterlife): pm.PrepareProbability(depends),
		string(BadAfterlife):     pm.PrepareProbability(bad),
	}
	return AfterlifeOption(pm.GetRandomFromSeveral(m))
}

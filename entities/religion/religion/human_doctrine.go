package religion

import "fmt"

type HumanDoctrine struct {
	religion *Religion

	Nature *HumanNature
}

func (d *Doctrine) generateHumanDoctrine() *HumanDoctrine {
	hd := &HumanDoctrine{religion: d.religion}
	hd.Nature = hd.generateHumanNature()

	return hd
}

func (hd *HumanDoctrine) Print() {
	fmt.Printf("Human Doctrine (religion=%s)\n", hd.religion.Name)
	hd.Nature.Print()
}

type HumanNature struct {
	religion *Religion

	Goodness   GoodnessNature
	HasSoul    bool
	CanBeSaint bool
}

func (hd *HumanDoctrine) generateHumanNature() *HumanNature {
	hn := &HumanNature{religion: hd.religion}

	return hn
}

func (hn *HumanNature) Print() {
	fmt.Printf("Human is %s and level of it is %s\n", hn.Goodness.Goodness, hn.Goodness.Level)
	humanNatureTraits := make([]string, 0)
	if hn.HasSoul {
		humanNatureTraits = append(humanNatureTraits, "has soul")
	}
	if hn.CanBeSaint {
		humanNatureTraits = append(humanNatureTraits, "can be saint")
	}

	for _, trait := range humanNatureTraits {
		fmt.Printf(" - human %s\n", trait)
	}
}

package religion

type HumanDoctrine struct {
	religion *Religion

	Nature *HumanNature
}

func (d *Doctrine) generateHumanDoctrine() *HumanDoctrine {
	hd := &HumanDoctrine{religion: d.religion}
	hd.Nature = hd.generateHumanNature()

	return hd
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

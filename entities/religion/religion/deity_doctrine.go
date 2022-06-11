package religion

type DeityDoctrine struct {
	religion *Religion

	Nature *DeityNature
}

func (d *Doctrine) generateDeityDoctrine() *DeityDoctrine {
	dd := &DeityDoctrine{religion: d.religion}

	return dd
}

type DeityNature struct {
	religion *Religion

	Goodness                  GoodnessNature
	IsJust                    bool
	IsWorldCreator            bool
	IsImmortal                bool
	TakePartInHumanLife       bool
	DoesCommunicateWithHumans bool
	Gnosticism                bool // God(s) created spiritual world but Demiurge(s) created physical world. Body is sinful
}

package religion

import "fmt"

type DeityDoctrine struct {
	religion *Religion

	Nature *DeityNature
}

func (d *Doctrine) generateDeityDoctrine() *DeityDoctrine {
	dd := &DeityDoctrine{religion: d.religion}

	return dd
}

func (dd *DeityDoctrine) Print() {
	fmt.Printf("Deity Doctrine (religion=%s)\n", dd.religion.Name)
	dd.Nature.Print()
}

type DeityNature struct {
	religion *Religion

	Goodness                  GoodnessNature
	IsJust                    bool
	IsWorldCreator            bool
	IsImmortal                bool
	IsTranscendental          bool
	TakePartInHumanLife       bool
	DoesCommunicateWithHumans bool
	Gnosticism                bool // God(s) created spiritual world but Demiurge(s) created physical world. Body is sinful
}

func (dn *DeityNature) Print() {
	fmt.Printf("Deity(s) is/are %s and level of it is %s\n", dn.Goodness.Goodness, dn.Goodness.Level)
	deityNatureTraits := make([]string, 0)
	if dn.IsJust {
		deityNatureTraits = append(deityNatureTraits, "is just")
	}
	if dn.IsImmortal {
		deityNatureTraits = append(deityNatureTraits, "is immortal")
	}
	if dn.IsWorldCreator {
		deityNatureTraits = append(deityNatureTraits, "is world creator")
	}
	if dn.IsTranscendental {
		deityNatureTraits = append(deityNatureTraits, "is transcendental")
	}
	if dn.TakePartInHumanLife {
		deityNatureTraits = append(deityNatureTraits, "take part in human life")
	}
	if dn.DoesCommunicateWithHumans {
		deityNatureTraits = append(deityNatureTraits, "communicate with humans")
	}
	if dn.Gnosticism {
		deityNatureTraits = append(deityNatureTraits, "gnostic deity")
	}

	for _, trait := range deityNatureTraits {
		fmt.Printf(" - deity %s\n", trait)
	}
}

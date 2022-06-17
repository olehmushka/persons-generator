package religion

import "fmt"

type Rituals struct {
	religion *Religion

	Initiation *InitiationRituals
}

func (t *Theology) generateRituals() *Rituals {
	rs := &Rituals{religion: t.religion}
	rs.Initiation = rs.generateInitiationRituals()

	return rs
}

func (rs *Rituals) Print() {
	fmt.Printf("Rituals (religion_name=%s):\n", rs.religion.Name)
}

type InitiationRituals struct {
	religion *Religion
}

func (rs *Rituals) generateInitiationRituals() *InitiationRituals {
	ir := &InitiationRituals{rs.religion}

	return ir
}

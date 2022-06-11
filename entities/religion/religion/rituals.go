package religion

type Rituals struct {
	religion *Religion

	Initiation *InitiationRituals
}

func NewRituals(r *Religion) *Rituals {
	rs := &Rituals{religion: r}
	rs.Initiation = rs.generateInitiationRituals()

	return rs
}

func (rs *Rituals) generateInitiationRituals() *InitiationRituals {
	ir := &InitiationRituals{}
	return ir
}

type InitiationRituals struct{}

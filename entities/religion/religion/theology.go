package religion

type Theology struct {
	religion *Religion
}

func NewTheology(r *Religion) *Theology {
	t := &Theology{religion: r}

	return t
}

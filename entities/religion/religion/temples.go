package religion

type Temples struct {
	religion *Religion
}

func NewTemples(r *Religion) *Temples {
	ts := &Temples{religion: r}

	return ts
}

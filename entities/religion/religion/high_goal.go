package religion

type HighGoal struct {
	religion *Religion
}

func NewHighGoal(r *Religion) *HighGoal {
	hg := &HighGoal{religion: r}

	return hg
}

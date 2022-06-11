package religion

type Holydays struct {
	religion *Religion
}

func NewHolydays(r *Religion) *Holydays {
	hs := &Holydays{religion: r}

	return hs
}

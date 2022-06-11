package religion

type HolyScripture struct {
	religion *Religion
}

func NewHolyScripture(r *Religion) *HolyScripture {
	hs := &HolyScripture{religion: r}

	return hs
}

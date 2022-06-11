package religion

type Religion struct {
	Type            *Type
	GenderDominance *GenderDominance
}

func NewReligion() *Religion {
	r := &Religion{}
	r.Type = NewType(r)
	r.GenderDominance = NewGenderDominance(r)

	return r
}

func NewReligions(n int) []*Religion {
	religions := make([]*Religion, n)
	for i := range religions {
		religions[i] = NewReligion()
	}

	return religions
}

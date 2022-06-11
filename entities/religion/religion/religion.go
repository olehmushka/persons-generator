package religion

type Religion struct {
	metadata *religionMetadata

	Type            *Type
	GenderDominance *GenderDominance
	Doctrine        *Doctrine
}

func NewReligion() *Religion {
	r := &Religion{}
	r.Type = NewType(r)
	r.GenderDominance = NewGenderDominance(r)
	r.Doctrine = NewDoctrine(r)

	return r
}

func NewReligions(n int) []*Religion {
	religions := make([]*Religion, n)
	for i := range religions {
		religions[i] = NewReligion()
	}

	return religions
}

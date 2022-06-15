package religion

import "fmt"

type Religion struct {
	metadata *religionMetadata

	Name            string
	Type            *Type
	GenderDominance *GenderDominance
	Doctrine        *Doctrine
}

func NewReligion() *Religion {
	r := &Religion{}
	r.Type = NewType(r)
	r.GenderDominance = NewGenderDominance(r)
	r.metadata = r.generateMetadata()
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

func (r *Religion) UpdateMetadata(rm *religionMetadata) {
	r.metadata = rm
}

func (r *Religion) Print() {
	fmt.Printf("Religion (name=%s)\n", r.Name)
	r.Type.Print()
	r.GenderDominance.Print()
	fmt.Printf("=====================\n\n")
}

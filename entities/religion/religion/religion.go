package religion

import "fmt"

type Religion struct {
	metadata *religionMetadata

	Name            string
	Type            *Type
	GenderDominance *GenderDominance
	Doctrine        *Doctrine
	Attributes      *Attributes
	Theology        *Theology
}

func NewReligion() *Religion {
	r := &Religion{}
	r.Type = NewType(r)
	r.GenderDominance = NewGenderDominance(r)
	r.metadata = r.generateMetadata()
	r.Doctrine = NewDoctrine(r)
	r.Attributes = NewAttributes(r)
	r.Theology = NewTheology(r)

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
	r.Doctrine.Print()
	r.Attributes.Print()
	r.Theology.Print()
	fmt.Printf("=====================\n\n")
}

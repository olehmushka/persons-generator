package religion

import (
	"fmt"

	nb "persons_generator/entities/name_base"
	ng "persons_generator/name_generator"
	pm "persons_generator/probability_machine"
)

type Religion struct {
	M        Metadata
	metadata *religionMetadata

	Name            string
	Type            *Type
	GenderDominance *GenderDominance
	Doctrine        *Doctrine
	Attributes      *Attributes
	Theology        *Theology
}

func NewReligion() *Religion {
	nameBase := nb.GetRandomNmeBase()
	r := &Religion{
		M: Metadata{
			LowBaseCoef:  pm.RandFloat64InRange(0.45, 0.75),
			BaseCoef:     pm.RandFloat64InRange(0.95, 1.05),
			HighBaseCoef: pm.RandFloat64InRange(1, 1.25),

			MaxMetadataValue: pm.RandFloat64InRange(8, 10),
		},
		Name: ng.GetBase(
			nameBase.Index,
			nb.NameBasesToStrSlices(nb.AllNameBases),
			nameBase.Min,
			nameBase.Max,
			nameBase.Dupl,
		),
	}
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

	fmt.Printf("\n\nMetadata:\n%+v\n", r.metadata)
	fmt.Printf("=====================\n\n")
}

/* ********************************************************************************************************** */

func (r *Religion) HasReincarnation() bool {
	if r.Doctrine != nil {
		if r.Doctrine.HighGoal != nil {
			if r.Doctrine.HighGoal.ContainsReincarnation() {
				return true
			}
		}
	}

	if r.Theology != nil {
		if r.Theology.HasReincarnation() {
			return true
		}
	}

	return false
}

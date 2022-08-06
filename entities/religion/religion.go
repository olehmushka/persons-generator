package religion

import (
	"fmt"

	"persons_generator/entities/culture"
	g "persons_generator/entities/gender"
	pm "persons_generator/probability_machine"
	"persons_generator/tools"
)

type Religion struct {
	M        Metadata
	metadata *religionMetadata

	Name            string
	Type            *Type
	GenderDominance *g.Dominance
	Doctrine        *Doctrine
	Attributes      *Attributes
	Theology        *Theology
}

func NewReligion(c *culture.Culture) *Religion {
	r := &Religion{
		M: Metadata{
			LowBaseCoef:  pm.RandFloat64InRange(0.45, 0.75),
			BaseCoef:     pm.RandFloat64InRange(0.95, 1.05),
			HighBaseCoef: pm.RandFloat64InRange(1, 1.25),

			MaxMetadataValue: pm.RandFloat64InRange(8, 10),
		},
		Name: c.Language.GetReligionName(),
	}
	r.Type = NewType(r)
	r.GenderDominance = c.GenderDominance
	r.metadata = r.generateMetadata()
	r.Doctrine = NewDoctrine(r)
	r.Attributes = NewAttributes(r, c)
	r.Theology = NewTheology(r, c)

	return r
}

func NewReligions(cultures []*culture.Culture) []*Religion {
	religions := make([]*Religion, len(cultures))
	for i := range religions {
		religions[i] = NewReligion(cultures[i])
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

func GetReligionByName(name string, list []*Religion) *Religion {
	if name == "" || len(list) == 0 {
		return nil
	}

	return tools.Search(list, func(r *Religion) string { return r.Name }, name)
}

func MapReligionNames(religions []*Religion) []string {
	if len(religions) == 0 {
		return []string{}
	}

	out := make([]string, len(religions))
	for i := range out {
		out[i] = religions[i].Name
	}

	return out
}

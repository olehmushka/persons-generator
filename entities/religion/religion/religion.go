package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
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
	r := &Religion{
		M: Metadata{
			LowBaseCoef:  pm.RandFloat64InRange(0.45, 0.75),
			BaseCoef:     pm.RandFloat64InRange(0.95, 1.05),
			HighBaseCoef: pm.RandFloat64InRange(1, 1.25),
		},
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

func (r *Religion) IsLawful() bool {
	var criterias float64

	if r.Doctrine != nil {
		if !r.Doctrine.SourceOfMoralLaw.IsNone() {
			criterias += 0.5
		}
		if r.Doctrine.Deity != nil {
			if r.Doctrine.Deity.Nature != nil {
				if len(r.Doctrine.Deity.Nature.Traits) != 0 {
					for _, trait := range r.Doctrine.Deity.Nature.Traits {
						if trait.Name == "IsJust" {
							criterias += 0.5
						}
					}
				}
			}
		}
	}

	if r.Attributes != nil {
		if r.Attributes.HolyScripture != nil {
			if len(r.Attributes.HolyScripture.Traits) != 0 {
				for _, trait := range r.Attributes.HolyScripture.Traits {
					switch trait.Name {
					case "Commandments":
						fallthrough
					case "DivineLaw":
						criterias += 1
					}
				}
			}
		}
	}

	if r.Theology != nil {
		if len(r.Theology.Traits) != 0 {
			for _, trait := range r.Theology.Traits {
				if trait.Name == "ReligiousLaw" {
					criterias += 1.5
				}
			}
		}
	}

	return criterias >= 1.5
}

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

func (r *Religion) IsNaturalistic() bool {
	var criterias float64
	if r.Doctrine != nil {
		if r.Doctrine.HighGoal != nil {
			criterias += r.Doctrine.HighGoal.GetNaturalisticCriterias()
		}
		if r.Doctrine.Social != nil {
			criterias += r.Doctrine.Social.GetNaturalisticCriterias()
		}
		if r.Doctrine.SourceOfMoralLaw.IsNature() {
			criterias += 1
		}
	}
	if r.Theology != nil {
		criterias += r.Theology.GetNaturalisticCriterias()
	}

	return criterias > 3
}

func (r *Religion) IsPacifistic() bool {
	var criterias float64
	if r.Doctrine != nil {
		if r.Doctrine.HighGoal != nil {
			criterias += r.Doctrine.HighGoal.GetPacifisticCriterias()
		}
		if r.Doctrine.Social != nil {
			criterias += r.Doctrine.Social.GetPacifisticCriterias()
		}
	}
	// if r.Theology != nil {
	// 	criterias += r.Theology.GetNaturalisticCriterias()
	// }

	return criterias > 1.5
}

func (r *Religion) IsAggressive() bool {
	var criterias float64
	if r.Doctrine != nil {
		if r.Doctrine.HighGoal != nil {
			criterias += r.Doctrine.HighGoal.GetAggressiveCriterias()
		}
		if r.Doctrine.Social != nil {
			criterias += r.Doctrine.Social.GetAggressiveCriterias()
		}
	}
	if r.Theology != nil {
		criterias += r.Theology.GetAggressiveCriterias()
	}

	return criterias > 1.5
}

func (r *Religion) IsSexualActive() bool {
	var criterias float64
	if r.Doctrine != nil {
		if r.Doctrine.HighGoal != nil {
			criterias += r.Doctrine.HighGoal.GetSexualActiveCriterias()
		}
		if r.Doctrine.Social != nil {
			criterias += r.Doctrine.Social.GetSexualActiveCriterias()
		}
	}
	// if r.Theology != nil {
	// 	criterias += r.Theology.GetNaturalisticCriterias()
	// }

	return criterias > 1.5
}

func (r *Religion) IsPlutocratic() bool {
	var criterias float64
	// if r.Doctrine != nil {
	// 	if r.Doctrine.HighGoal != nil {
	// 		criterias += r.Doctrine.HighGoal.GetSexualActiveCriterias()
	// 	}
	// 	if r.Doctrine.Social != nil {
	// 		criterias += r.Doctrine.Social.GetSexualActiveCriterias()
	// 	}
	// }
	if r.Theology != nil {
		criterias += r.Theology.GetPlutocracyCriterias()
	}

	return criterias > 1.5
}

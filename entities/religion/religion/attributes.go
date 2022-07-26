package religion

import (
	"fmt"

	pm "persons_generator/probability_machine"
)

type Attributes struct {
	religion *Religion

	Traits        []*trait
	Clerics       *Clerics
	HolyScripture *HolyScripture
	Temples       *Temples
}

func NewAttributes(r *Religion) *Attributes {
	attrs := &Attributes{religion: r}
	attrs.Traits = generateTraits(r, attrs.getAllAttributeTraits(), generateTraitsOpts{
		Label: "Attributes.generateTraits",
		Min:   1,
		Max:   len(attrs.getAllAttributeTraits()),
	})
	attrs.Clerics = attrs.generateClerics()
	attrs.HolyScripture = attrs.generateHolyScripture()
	attrs.Temples = attrs.generateTemples()

	return attrs
}

func (as *Attributes) Print() {
	fmt.Printf("Attributes (religion_name=%s):\n", as.religion.Name)
	if len(as.Traits) > 0 {
		fmt.Printf("Attributes Traits (religion_name=%s):\n", as.religion.Name)
		for _, trait := range as.Traits {
			fmt.Printf(" - %s\n", trait.Name)
		}
	}
	as.Clerics.Print()
	as.HolyScripture.Print()
	as.Temples.Print()
}

func (as *Attributes) getAllAttributeTraits() []*trait {
	return []*trait{
		{
			Name: "Amulets",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Chthonic:     0.1,
				Simple:       0.25,
			},
			baseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "ReligionMusic",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Simple:       0.25,
			},
			baseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "ReligionTheatre",
			_religionMetadata: &religionMetadata{
				Naturalistic:   0.75,
				Collectivistic: 0.75,
				Simple:         0.25,
			},
			baseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "ReligionPoetry",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Simple:       0.25,
			},
			baseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "ReligionTapestry",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Simple:       0.25,
			},
			baseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Calligraphy",
			_religionMetadata: &religionMetadata{
				Individualistic: 0.5,
			},
			baseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Idols",
			_religionMetadata: &religionMetadata{
				Chthonic: 0.5,
			},
			baseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				switch {
				case r.Type.IsMonotheism():
					addCoef -= pm.RandFloat64InRange(0.01, 0.1)
				case r.Type.IsPolytheism():
					addCoef += pm.RandFloat64InRange(1, 2)
				case r.Type.IsAtheism():
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

package religion

import (
	"fmt"

	"persons_generator/engine/entities/culture"
	pm "persons_generator/engine/probability_machine"
)

type Attributes struct {
	religion *Religion

	Traits        []*trait
	Clerics       *Clerics
	HolyScripture *HolyScripture
	Temples       *Temples
}

func NewAttributes(r *Religion, c *culture.Culture) (*Attributes, error) {
	attrs := &Attributes{religion: r}
	ts, err := generateTraits(r, attrs.getAllAttributeTraits(), generateTraitsOpts{
		Label: "Attributes.generateTraits",
		Min:   1,
		Max:   len(attrs.getAllAttributeTraits()),
	})
	if err != nil {
		return nil, err
	}
	attrs.Traits = ts
	cs, err := attrs.generateClerics()
	if err != nil {
		return nil, err
	}
	attrs.Clerics = cs
	hs, err := attrs.generateHolyScripture()
	if err != nil {
		return nil, err
	}
	attrs.HolyScripture = hs
	temples, err := attrs.generateTemples(c)
	if err != nil {
		return nil, err
	}
	attrs.Temples = temples

	return attrs, nil
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
			Name: AmuletsAttribute,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Chthonic:     0.1,
				Simple:       0.25,
			},
			baseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ReligionMusicAttribute,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Simple:       0.25,
			},
			baseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ReligionTheatreAttribute,
			_religionMetadata: &religionMetadata{
				Naturalistic:   0.75,
				Collectivistic: 0.75,
				Simple:         0.25,
			},
			baseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ReligionPoetryAttribute,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Simple:       0.25,
			},
			baseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ReligionTapestryAttribute,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Simple:       0.25,
			},
			baseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: CalligraphyAttribute,
			_religionMetadata: &religionMetadata{
				Individualistic: 0.5,
			},
			baseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: IdolsAttribute,
			_religionMetadata: &religionMetadata{
				Chthonic: 0.5,
			},
			baseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				switch {
				case r.Type.IsMonotheism():
					p, err := pm.RandFloat64InRange(0.01, 0.1)
					if err != nil {
						return false, err
					}
					addCoef -= p
				case r.Type.IsPolytheism():
					p, err := pm.RandFloat64InRange(1, 2)
					if err != nil {
						return false, err
					}
					addCoef += p
				case r.Type.IsAtheism():
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

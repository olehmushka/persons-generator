package religion

import (
	"fmt"

	"persons_generator/engine/entities/culture"
	pm "persons_generator/engine/probability_machine"
)

type Attributes struct {
	religion *Religion `json:"-" bson:"-"`

	Traits        []*trait       `json:"traits" bson:"traits"`
	Clerics       *Clerics       `json:"clerics" bson:"clerics"`
	HolyScripture *HolyScripture `json:"holy_scripture" bson:"holy_scripture"`
	Temples       *Temples       `json:"temples" bson:"temples"`
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

func (as *Attributes) SerializeTraits() []string {
	if as == nil {
		return []string{}
	}

	return extractTraitNames(as.Traits)
}

func (as *Attributes) getAllAttributeTraits() []*trait {
	return []*trait{
		{
			Name: AmuletsAttribute,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Chthonic:     0.1,
				Simple:       0.25,
			},
			BaseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ReligionMusicAttribute,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Simple:       0.25,
			},
			BaseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ReligionTheatreAttribute,
			ReligionMetadata: &religionMetadata{
				Naturalistic:   0.75,
				Collectivistic: 0.75,
				Simple:         0.25,
			},
			BaseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ReligionPoetryAttribute,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Simple:       0.25,
			},
			BaseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ReligionTapestryAttribute,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Simple:       0.25,
			},
			BaseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: CalligraphyAttribute,
			ReligionMetadata: &religionMetadata{
				Individualistic: 0.5,
			},
			BaseCoef: as.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: IdolsAttribute,
			ReligionMetadata: &religionMetadata{
				Chthonic: 0.5,
			},
			BaseCoef: as.religion.M.BaseCoef,
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

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func GetAttributesSimilarityCoef(a1, a2 *Attributes) float64 {
	if a1 == nil && a2 == nil {
		return 1
	}
	if a1 == nil || a2 == nil {
		return 0
	}

	similarityTraits := []struct {
		value float64
		coef  float64
	}{
		{
			value: GetTraitsSimilarityCoef(a1.Traits, a2.Traits),
			coef:  0.25,
		},
		{
			value: GetClericsSimilarityCoef(a1.Clerics, a2.Clerics),
			coef:  0.25,
		},
		{
			value: GetHolyScriptureSimilarityCoef(a1.HolyScripture, a2.HolyScripture),
			coef:  0.25,
		},
		{
			value: GetTemplesSimilarityCoef(a1.Temples, a2.Temples),
			coef:  0.25,
		},
	}

	var out float64
	for _, t := range similarityTraits {
		out += t.value * t.coef
	}

	return out
}

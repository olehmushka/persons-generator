package religion

import (
	"fmt"
	pm "persons_generator/engine/probability_machine"
)

type Conversion struct {
	religion *Religion `json:"-" bson:"-"`
	theology *Theology `json:"-" bson:"-"`

	Traits []*trait `json:"traits" bson:"traits"`
}

func (t *Theology) generateConversion() (*Conversion, error) {
	c := &Conversion{religion: t.religion, theology: t}
	ts, err := generateTraits(c.religion, c.getAllConversionTraits(), generateTraitsOpts{
		Label: "Conversion.generateTraits",
		Min:   1,
		Max:   3,
	})
	if err != nil {
		return nil, err
	}
	c.Traits = ts

	return c, nil
}

func (c *Conversion) IsZero() bool {
	return c == nil
}

func (c *Conversion) Serialize() []string {
	if c == nil {
		return []string{}
	}

	return extractTraitNames(c.Traits)
}

func (c *Conversion) Print() {
	if len(c.Traits) > 0 {
		fmt.Printf("Conversion (religion_name=%s):\n", c.religion.Name)
		for _, trait := range c.Traits {
			fmt.Printf(" - %s\n", trait.Name)
		}
	}
}

func (c *Conversion) getAllConversionTraits() []*trait {
	return []*trait{
		{
			Name: MissionaryOutreachConversionWay,
			ReligionMetadata: &religionMetadata{
				Altruistic:     0.25,
				Educational:    0.25,
				Collectivistic: 0.1,
			},
			BaseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: MasterConversionWay,
			ReligionMetadata: &religionMetadata{
				Altruistic: 0.1,
				Simple:     0.5,
			},
			BaseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ForcedConversionConversionWay,
			ReligionMetadata: &religionMetadata{
				Aggressive:      1,
				Authoritaristic: 1,
			},
			BaseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: InitiationWithFireConversionWay,
			ReligionMetadata: &religionMetadata{
				Chthonic:        0.25,
				Aggressive:      0.5,
				Authoritaristic: 1,
			},
			BaseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: InitiatedInBloodConversionWay,
			ReligionMetadata: &religionMetadata{
				Chthonic:        0.75,
				Authoritaristic: 1,
			},
			BaseCoef: c.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: PreachersConversionWay,
			ReligionMetadata: &religionMetadata{
				Altruistic: 0.75,
				Pacifistic: 0.1,
			},
			BaseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.Attributes != nil && r.Attributes.Clerics != nil && len(r.Attributes.Clerics.Functions) > 0 {
					for _, fn := range r.Attributes.Clerics.Functions {
						if fn == nil {
							continue
						}
						if fn.Name == TeachClericFunction {
							coef, err := pm.RandFloat64InRange(0.05, 0.1)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: WitchcraftConversionWay,
			ReligionMetadata: &religionMetadata{
				Chthonic: 0.5,
			},
			BaseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.Theology != nil && r.Theology.Taboos != nil && len(r.Theology.Taboos.Taboos) > 0 {
					for _, taboo := range r.Theology.Taboos.Taboos {
						if taboo == nil {
							continue
						}
						if taboo.Name == WitchcraftTabooName {
							switch {
							case taboo.Acceptance.IsAccepted():
								coef, err := pm.RandFloat64InRange(0.01, 0.1)
								if err != nil {
									return false, err
								}
								addCoef += coef
							case taboo.Acceptance.IsShunned():
								coef, err := pm.RandFloat64InRange(0.01, 0.1)
								if err != nil {
									return false, err
								}
								addCoef -= coef
							case taboo.Acceptance.IsCriminal():
								return false, nil
							}
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DarkPactsConversionWay,
			ReligionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			BaseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.Theology != nil && r.Theology.Taboos != nil && len(r.Theology.Taboos.Taboos) > 0 {
					for _, taboo := range r.Theology.Taboos.Taboos {
						if taboo == nil {
							continue
						}
						if taboo.Name == WitchcraftTabooName {
							switch {
							case taboo.Acceptance.IsAccepted():
								coef, err := pm.RandFloat64InRange(0.01, 0.1)
								if err != nil {
									return false, err
								}
								addCoef += coef
							case taboo.Acceptance.IsShunned():
								coef, err := pm.RandFloat64InRange(0.03, 0.13)
								if err != nil {
									return false, err
								}
								addCoef -= coef
							case taboo.Acceptance.IsCriminal():
								coef, err := pm.RandFloat64InRange(0.08, 0.18)
								if err != nil {
									return false, err
								}
								addCoef -= coef
							}
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func GetConversionSimilarityCoef(c1, c2 *Conversion) float64 {
	if c1.IsZero() && c2.IsZero() {
		return 1
	}
	if c1.IsZero() || c2.IsZero() {
		return 0
	}

	return GetTraitsSimilarityCoef(c1.Traits, c2.Traits)
}

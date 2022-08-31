package religion

import (
	"fmt"
	pm "persons_generator/engine/probability_machine"
)

type Conversion struct {
	religion *Religion
	theology *Theology

	Traits []*trait
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
			_religionMetadata: &religionMetadata{
				Altruistic:     0.25,
				Educational:    0.25,
				Collectivistic: 0.1,
			},
			baseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: MasterConversionWay,
			_religionMetadata: &religionMetadata{
				Altruistic: 0.1,
				Simple:     0.5,
			},
			baseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ForcedConversionConversionWay,
			_religionMetadata: &religionMetadata{
				Aggressive:      1,
				Authoritaristic: 1,
			},
			baseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: InitiationWithFireConversionWay,
			_religionMetadata: &religionMetadata{
				Chthonic:        0.25,
				Aggressive:      0.5,
				Authoritaristic: 1,
			},
			baseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: InitiatedInBloodConversionWay,
			_religionMetadata: &religionMetadata{
				Chthonic:        0.75,
				Authoritaristic: 1,
			},
			baseCoef: c.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: PreachersConversionWay,
			_religionMetadata: &religionMetadata{
				Altruistic: 0.75,
				Pacifistic: 0.1,
			},
			baseCoef: c.religion.M.BaseCoef,
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

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: WitchcraftConversionWay,
			_religionMetadata: &religionMetadata{
				Chthonic: 0.5,
			},
			baseCoef: c.religion.M.BaseCoef,
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

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DarkPactsConversionWay,
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: c.religion.M.BaseCoef,
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

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

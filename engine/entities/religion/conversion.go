package religion

import (
	"fmt"
)

type Conversion struct {
	religion *Religion

	Traits []*trait
}

func (t *Theology) generateConversion() (*Conversion, error) {
	c := &Conversion{religion: t.religion}
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
			Name: "missionary_outreach",
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
			Name: "master",
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
			Name: "forced_conversion",
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
			Name: "initiation_with_fire",
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
			Name: "initiated_in_blood",
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
			Name: "preachers",
			_religionMetadata: &religionMetadata{
				Altruistic: 0.75,
				Pacifistic: 0.1,
			},
			baseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "witchcraft",
			_religionMetadata: &religionMetadata{
				Chthonic: 0.5,
			},
			baseCoef: c.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "dark_pacts",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: c.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

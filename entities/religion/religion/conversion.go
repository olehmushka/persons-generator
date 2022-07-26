package religion

import (
	"fmt"
)

type Conversion struct {
	religion *Religion

	Traits []*trait
}

func (t *Theology) generateConversion() *Conversion {
	c := &Conversion{religion: t.religion}
	c.Traits = generateTraits(c.religion, c.getAllConversionTraits(), generateTraitsOpts{
		Label: "Conversion.generateTraits",
		Min:   1,
		Max:   3,
	})

	return c
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
			Name: "MissionaryOutreach",
			_religionMetadata: &religionMetadata{
				Altruistic:     0.25,
				Educational:    0.25,
				Collectivistic: 0.1,
			},
			baseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Master",
			_religionMetadata: &religionMetadata{
				Altruistic: 0.1,
				Simple:     0.5,
			},
			baseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "ForcedConversion",
			_religionMetadata: &religionMetadata{
				Aggressive:      1,
				Authoritaristic: 1,
			},
			baseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "InitiationWithFire",
			_religionMetadata: &religionMetadata{
				Chthonic:        0.25,
				Aggressive:      0.5,
				Authoritaristic: 1,
			},
			baseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "InitiatedInBlood",
			_religionMetadata: &religionMetadata{
				Chthonic:        0.75,
				Authoritaristic: 1,
			},
			baseCoef: c.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Preachers",
			_religionMetadata: &religionMetadata{
				Altruistic: 0.75,
				Pacifistic: 0.1,
			},
			baseCoef: c.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Witchcraft",
			_religionMetadata: &religionMetadata{
				Chthonic: 0.5,
			},
			baseCoef: c.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "DarkPacts",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: c.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

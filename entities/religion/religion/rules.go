package religion

import "fmt"

type Rules struct {
	religion *Religion

	Rules []*trait
}

func (t *Theology) generateRules() *Rules {
	rs := &Rules{religion: t.religion}
	rs.Rules = generateTraits(rs.religion, rs.getAllRules(), generateTraitsOpts{
		Label: "Rules.generateRules",
		Min:   1,
		Max:   5,
	})

	return rs
}

func (rs *Rules) Print() {
	if len(rs.Rules) > 0 {
		fmt.Printf("Rules (religion_name=%s):\n", rs.religion.Name)
		for _, rule := range rs.Rules {
			fmt.Printf(" - %s\n", rule.Name)
		}
		return
	}
	fmt.Printf("Has not Rules (religion_name=%s):\n", rs.religion.Name)
}

func (rs *Rules) getAllRules() []*trait {
	return []*trait{
		{
			Name: "PrayWithFrequency",
			_religionMetadata: &religionMetadata{
				Lawful:  0.5,
				Ascetic: 0.25,
				Simple:  0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "TaxToSupportPoor",
			_religionMetadata: &religionMetadata{
				Plutocratic: 0.25,
				Altruistic:  1,
				Lawful:      0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Tithe",
			_religionMetadata: &religionMetadata{
				Plutocratic: 1,
				Lawful:      0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "DressCode",
			_religionMetadata: &religionMetadata{
				Lawful:          0.5,
				Ascetic:         0.5,
				Authoritaristic: 0.5,
				Collectivistic:  0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "BeHumble",
			_religionMetadata: &religionMetadata{
				Altruistic: 1,
				Simple:     1,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "BeHonest",
			_religionMetadata: &religionMetadata{
				Altruistic: 0.5,
				Simple:     1,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "LearnKeyScriptures",
			_religionMetadata: &religionMetadata{
				Educational:     0.75,
				Authoritaristic: 0.25,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "TaxNonbelievers",
			_religionMetadata: &religionMetadata{
				Plutocratic:     0.5,
				Authoritaristic: 0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

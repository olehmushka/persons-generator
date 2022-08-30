package religion

import "fmt"

type Rules struct {
	religion *Religion

	Rules []*trait
}

func (t *Theology) generateRules() (*Rules, error) {
	rs := &Rules{religion: t.religion}
	rules, err := generateTraits(rs.religion, rs.getAllRules(), generateTraitsOpts{
		Label: "Rules.generateRules",
		Min:   1,
		Max:   5,
	})
	if err != nil {
		return nil, err
	}
	rs.Rules = rules

	return rs, nil
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
			Name: "pray_with_frequency",
			_religionMetadata: &religionMetadata{
				Lawful:  0.5,
				Ascetic: 0.25,
				Simple:  0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "tax_to_support_poor",
			_religionMetadata: &religionMetadata{
				Plutocratic: 0.25,
				Altruistic:  1,
				Lawful:      0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "tithe",
			_religionMetadata: &religionMetadata{
				Plutocratic: 1,
				Lawful:      0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "dress_code",
			_religionMetadata: &religionMetadata{
				Lawful:          0.5,
				Ascetic:         0.5,
				Authoritaristic: 0.5,
				Collectivistic:  0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "be_humble",
			_religionMetadata: &religionMetadata{
				Altruistic: 1,
				Simple:     1,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "be_honest",
			_religionMetadata: &religionMetadata{
				Altruistic: 0.5,
				Simple:     1,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "learn_key_scriptures",
			_religionMetadata: &religionMetadata{
				Educational:     0.75,
				Authoritaristic: 0.25,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "tax_nonbelievers",
			_religionMetadata: &religionMetadata{
				Plutocratic:     0.5,
				Authoritaristic: 0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

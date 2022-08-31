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
			Name: PrayWithFrequencyRule,
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
			Name: TaxToSupportPoorRule,
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
			Name: TitheRule,
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
			Name: DressCodeRule,
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
			Name: BeHumbleRule,
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
			Name: BeHonestRule,
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
			Name: LearnKeyScripturesRule,
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
			Name: TaxNonbelieversRule,
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

package religion

import "fmt"

type Rules struct {
	religion *Religion

	Rules []*Rule
}

func (t *Theology) generateRules() *Rules {
	rs := &Rules{religion: t.religion}
	rs.Rules = rs.generateRules(1, 5)

	return rs
}

func (rs *Rules) Print() {
	fmt.Printf("Rules (religion_name=%s):\n", rs.religion.Name)
	for _, rule := range rs.Rules {
		fmt.Printf(" - %s\n", rule.Name)
	}
}

type Rule struct {
	_religionMetadata *religionMetadata
	baseCoef          float64

	Name string
	Calc func(r *Religion, self *Rule, selectedRules []*Rule) bool
}

func (rs *Rules) generateRules(min, max int) []*Rule {
	if min < 0 {
		panic("[Rules.generateRules] min can not be less than 0")
	}
	allRules := rs.getAllRules()
	if max > len(allRules) {
		panic("[Rules.generateRules] max can not be greater than allRules length")
	}

	cults := make([]*Rule, 0, len(allRules))
	for count := 0; count < 20; count++ {
		for _, rule := range allRules {
			if rule.Calc(rs.religion, rule, cults) {
				cults = append(cults, &Rule{
					_religionMetadata: rule._religionMetadata,
					baseCoef:          rule.baseCoef,
					Name:              rule.Name,
					Calc:              rule.Calc,
				})
			}
			if len(cults) == max {
				break
			}
		}
		if len(cults) == max {
			break
		}
		if len(cults) >= min {
			break
		}
	}

	for _, cult := range cults {
		rs.religion.UpdateMetadata(UpdateReligionMetadata(*rs.religion.metadata, *cult._religionMetadata))
	}

	return cults
}

func (rs *Rules) getAllRules() []*Rule {
	return []*Rule{
		{
			Name: "PrayWithFrequency",
			_religionMetadata: &religionMetadata{
				Lawful:  0.5,
				Ascetic: 0.25,
				Simple:  0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Rule, _ []*Rule) bool {
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
			Calc: func(r *Religion, self *Rule, _ []*Rule) bool {
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
			Calc: func(r *Religion, self *Rule, _ []*Rule) bool {
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
			Calc: func(r *Religion, self *Rule, _ []*Rule) bool {
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
			Calc: func(r *Religion, self *Rule, _ []*Rule) bool {
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
			Calc: func(r *Religion, self *Rule, _ []*Rule) bool {
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
			Calc: func(r *Religion, self *Rule, _ []*Rule) bool {
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
			Calc: func(r *Religion, self *Rule, _ []*Rule) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

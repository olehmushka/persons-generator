package religion

import "fmt"

type Rules struct {
	religion *Religion `json:"-" bson:"-"`

	Rules []*trait `json:"rules" bson:"rules"`
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

func (rs *Rules) IsZero() bool {
	return rs == nil
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
			ReligionMetadata: &religionMetadata{
				Lawful:  0.5,
				Ascetic: 0.25,
				Simple:  0.5,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: TaxToSupportPoorRule,
			ReligionMetadata: &religionMetadata{
				Plutocratic: 0.25,
				Altruistic:  1,
				Lawful:      0.5,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: TitheRule,
			ReligionMetadata: &religionMetadata{
				Plutocratic: 1,
				Lawful:      0.5,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DressCodeRule,
			ReligionMetadata: &religionMetadata{
				Lawful:          0.5,
				Ascetic:         0.5,
				Authoritaristic: 0.5,
				Collectivistic:  0.5,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BeHumbleRule,
			ReligionMetadata: &religionMetadata{
				Altruistic: 1,
				Simple:     1,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BeHonestRule,
			ReligionMetadata: &religionMetadata{
				Altruistic: 0.5,
				Simple:     1,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: LearnKeyScripturesRule,
			ReligionMetadata: &religionMetadata{
				Educational:     0.75,
				Authoritaristic: 0.25,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: TaxNonbelieversRule,
			ReligionMetadata: &religionMetadata{
				Plutocratic:     0.5,
				Authoritaristic: 0.5,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func GetRulesSimilarityCoef(r1, r2 *Rules) float64 {
	if r1.IsZero() && r2.IsZero() {
		return 1
	}
	if r1.IsZero() || r2.IsZero() {
		return 0
	}

	return GetTraitsSimilarityCoef(r1.Rules, r2.Rules)
}

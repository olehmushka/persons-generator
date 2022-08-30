package religion

import (
	"fmt"

	"persons_generator/engine/entities/culture"
	pm "persons_generator/engine/probability_machine"
)

type Theology struct {
	religion *Religion

	Traits            []*trait
	Cults             []*trait
	Rules             *Rules
	Taboos            *Taboos
	Rituals           *Rituals
	Holydays          *Holydays
	Conversion        *Conversion
	MarriageTradition *MarriageTradition
}

func NewTheology(r *Religion, c *culture.Culture) (*Theology, error) {
	t := &Theology{religion: r}
	ts, err := generateTraits(r, t.getAllTheologyTraits(), generateTraitsOpts{
		Label: "Theology.generateTraits",
		Min:   1,
		Max:   5,
	})
	if err != nil {
		return nil, err
	}
	t.Traits = ts

	minCults, maxCults := t.getCultsRange()
	cults, err := t.generateCults(minCults, maxCults)
	if err != nil {
		return nil, err
	}
	t.Cults = cults

	rules, err := t.generateRules()
	if err != nil {
		return nil, err
	}
	t.Rules = rules

	taboos, err := t.generateTaboos(c)
	if err != nil {
		return nil, err
	}
	t.Taboos = taboos

	rituals, err := t.generateRituals()
	if err != nil {
		return nil, err
	}
	t.Rituals = rituals

	holydays, err := t.generateHolydays()
	if err != nil {
		return nil, err
	}
	t.Holydays = holydays

	conversion, err := t.generateConversion()
	if err != nil {
		return nil, err
	}
	t.Conversion = conversion

	marriageTradition, err := t.generateMarriageTradition(c)
	if err != nil {
		return nil, err
	}
	t.MarriageTradition = marriageTradition

	return t, nil
}

func (t *Theology) Print() {
	fmt.Printf("Theology (religion_name=%s):\n", t.religion.Name)
	fmt.Printf("Theology Traits (religion_name=%s):\n", t.religion.Name)
	for _, trait := range t.Traits {
		fmt.Printf(" - %s\n", trait.Name)
	}
	fmt.Printf("Cults (religion_name=%s):\n", t.religion.Name)
	for _, cult := range t.Cults {
		fmt.Printf(" - %s\n", cult.Name)
	}
	t.Rules.Print()
	t.Taboos.Print()
	t.Rituals.Print()
	t.Holydays.Print()
	t.Conversion.Print()
	t.MarriageTradition.Print()
}

func (t *Theology) getAllTheologyTraits() []*trait {
	return []*trait{
		{
			Name: "messiah",
			_religionMetadata: &religionMetadata{
				Lawful:          0.25,
				Authoritaristic: 0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				switch {
				case r.Type.IsMonotheism():
					addCoef += 0.1
				case r.Type.IsPolytheism():
					addCoef += 0.01
				case r.Type.IsDeityDualism():
					addCoef += 0.05
				case r.Type.IsDeism():
					addCoef += 0.005
				default:
					return false, nil
				}
				if t.religion.Attributes != nil {
					if t.religion.Attributes.HolyScripture != nil {
						if t.religion.Attributes.HolyScripture.HasHolyScripture {
							coef, err := pm.RandFloat64InRange(0.05, 0.2)
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
			Name: "prophets",
			_religionMetadata: &religionMetadata{
				Lawful:          0.5,
				Authoritaristic: 0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "religious_law",
			_religionMetadata: &religionMetadata{
				Lawful: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.metadata.IsLawful() {
					return true, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "reincarnation",
			_religionMetadata: &religionMetadata{
				Individualistic: 0.75,
				Complicated:     0.75,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.HasReincarnation() {
					return true, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "sanctioned_false_conversions", // If your life is threatened, it is acceptable to confess a false faith, as long as you keep the truth in your heart.
			_religionMetadata: &religionMetadata{
				Pacifistic:      0.25,
				Liberal:         1,
				Individualistic: 0.75,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "tree_connection", // Trees are the essence of life, and we must be near them.
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Chthonic:     0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.metadata.IsNaturalistic() {
					return true, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "animal_connection",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Chthonic:     0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.metadata.IsNaturalistic() {
					return true, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "blindsight", // Only the blind can perceive true reality
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
				Ascetic:  0.5,
			},
			baseCoef: t.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "astrology",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Educational:  0.5,
				Complicated:  1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "repentance",
			_religionMetadata: &religionMetadata{
				Lawful:  0.25,
				Ascetic: 0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "martyrdom",
			_religionMetadata: &religionMetadata{
				Ascetic:        0.25,
				Collectivistic: 0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				var addCoef float64
				for _, trait := range selectedTraits {
					switch trait.Name {
					case "messiah":
						coef, err := pm.RandFloat64InRange(0.01, 0.125)
						if err != nil {
							return false, err
						}
						addCoef += coef
					case "prophets":
						coef, err := pm.RandFloat64InRange(0.1, 0.2)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "indulgences",
			_religionMetadata: &religionMetadata{
				Plutocratic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "holiness",
			_religionMetadata: &religionMetadata{
				Ascetic:         0.75,
				Individualistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				for _, goal := range r.Doctrine.HighGoal.Goals {
					switch goal.Name {
					case "bring_holiness_down_to_the_world":
						coef, err := pm.RandFloat64InRange(0.03, 0.1)
						if err != nil {
							return false, err
						}
						addCoef += coef
					case "become_perfect_and_saints":
						coef, err := pm.RandFloat64InRange(0.05, 0.15)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "celibacy",
			_religionMetadata: &religionMetadata{
				SexualStrictness: 1,
				Ascetic:          1,
				Individualistic:  0.5,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.Attributes.Clerics.HasClerics {
					if r.Attributes.Clerics.Limitations.Marriage.IsDisallowed() {
						coef, err := pm.RandFloat64InRange(0.05, 0.2)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "flagellantism",
			_religionMetadata: &religionMetadata{
				Ascetic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				var addCoef float64
				for _, trait := range selectedTraits {
					switch trait.Name {
					case "repentance":
						fallthrough
					case "martyrdom":
						coef, err := pm.RandFloat64InRange(0.05, 0.15)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "feed_the_world",
			_religionMetadata: &religionMetadata{
				Altruistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "holy_army",
			_religionMetadata: &religionMetadata{
				Chthonic:       0.25,
				Aggressive:     1,
				Collectivistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.metadata.IsAggressive() {
					coef, err := pm.RandFloat64InRange(0.05, 0.125)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "defenders_of_faith",
			_religionMetadata: &religionMetadata{
				Lawful:     0.25,
				Aggressive: 0.75,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "non_violent_resistance",
			_religionMetadata: &religionMetadata{
				Pacifistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.metadata.IsAggressive() {
					p, err := pm.RandFloat64InRange(0, 0.05)
					if err != nil {
						return false, err
					}

					return pm.GetRandomBool(p)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "no_more_killing",
			_religionMetadata: &religionMetadata{
				Altruistic: 1,
				Pacifistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.metadata.IsAggressive() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "animal_messengers",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "sanctuary_of_mind",
			_religionMetadata: &religionMetadata{
				Individualistic: 1,
				Complicated:     0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "phantoms",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (t *Theology) HasReincarnation() bool {
	if len(t.Traits) == 0 {
		return false
	}

	for _, trait := range t.Traits {
		if trait == nil {
			continue
		}
		if trait.Name == "reincarnation" {
			return true
		}
	}

	return false
}

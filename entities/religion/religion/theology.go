package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
)

type Theology struct {
	religion *Religion

	Traits            []*theologyTrait
	Cults             []*Cult
	Rules             *Rules
	Taboos            *Taboos
	Rituals           *Rituals
	Holydays          *Holydays
	Conversion        *Conversion
	MarriageTradition *MarriageTradition
}

func NewTheology(r *Religion) *Theology {
	t := &Theology{religion: r}
	t.Traits = t.generateTraits(1, 5)
	minCults, maxCults := t.getCultsRange()
	t.Cults = t.generateCults(minCults, maxCults)
	t.Rules = t.generateRules()
	// t.Taboos = t.generateTaboos()
	// t.Rituals = t.generateRituals()
	// t.Holydays = t.generateHolydays()
	// t.Conversion = t.generateConversion()
	// t.MarriageTradition = t.generateMarriageTradition()

	return t
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
	// t.Taboos.Print()
	// t.Rituals.Print()
	// t.Holydays.Print()
	// t.Conversion.Print()
	// t.MarriageTradition.Print()
}

func (t *Theology) generateTraits(min, max int) []*theologyTrait {
	if min < 0 {
		panic("[Theology.generateTraits] min can not be less than 0")
	}
	allTraits := t.getAllTheologyTraits()
	if max > len(allTraits) {
		panic("[Theology.generateTraits] max can not be greater than allTraits length")
	}

	traits := make([]*theologyTrait, 0, len(allTraits))
	for count := 0; count < 20; count++ {
		for _, trait := range allTraits {
			if trait.Calc(t.religion, trait, traits) {
				traits = append(traits, &theologyTrait{
					_religionMetadata: trait._religionMetadata,
					baseCoef:          trait.baseCoef,
					Name:              trait.Name,
					Calc:              trait.Calc,
				})
			}
			if len(traits) == max {
				break
			}
		}
		if len(traits) == max {
			break
		}
		if len(traits) >= min {
			break
		}
	}

	for _, trait := range traits {
		t.religion.UpdateMetadata(UpdateReligionMetadata(*t.religion.metadata, *trait._religionMetadata))
		if trait.Name == "Reincarnation" {
			t.religion.Doctrine.Afterlife.updateAfterlife(false)
		}
	}

	return traits
}

type theologyTrait struct {
	_religionMetadata *religionMetadata
	baseCoef          float64
	Name              string
	Calc              func(r *Religion, self *theologyTrait, selectedTraits []*theologyTrait) bool
}

func (t *Theology) getAllTheologyTraits() []*theologyTrait {
	return []*theologyTrait{
		{
			Name: "Messiah",
			_religionMetadata: &religionMetadata{
				Lawful:          0.25,
				Authoritaristic: 0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
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
					return false
				}
				if t.religion.Attributes != nil {
					if t.religion.Attributes.HolyScripture != nil {
						if t.religion.Attributes.HolyScripture.HasHolyScripture {
							addCoef += pm.RandFloat64InRange(0.05, 0.2)
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Prophets",
			_religionMetadata: &religionMetadata{
				Lawful:          0.5,
				Authoritaristic: 0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, selectedTraits []*theologyTrait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "ReligiousLaw",
			_religionMetadata: &religionMetadata{
				Lawful: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				var addCoef float64
				if r.metadata.IsLawful() {
					return true
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Reincarnation",
			_religionMetadata: &religionMetadata{
				Individualistic: 0.75,
				Complicated:     0.75,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				if r.HasReincarnation() {
					return true
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SanctionedFalseConversions", // If your life is threatened, it is acceptable to confess a false faith, as long as you keep the truth in your heart.
			_religionMetadata: &religionMetadata{
				Pacifistic:      0.25,
				Liberal:         1,
				Individualistic: 0.75,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "TreeConnection", // Trees are the essence of life, and we must be near them.
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Chthonic:     0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				var addCoef float64
				if r.metadata.IsNaturalistic() {
					return true
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "AnimalConnection",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Chthonic:     0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				var addCoef float64
				if r.metadata.IsNaturalistic() {
					return true
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Blindsight", // Only the blind can perceive true reality
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
				Ascetic:  0.5,
			},
			baseCoef: t.religion.M.BaseCoef - pm.RandFloat64InRange(0.1, 0.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Astrology",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Educational:  0.5,
				Complicated:  1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Repentance",
			_religionMetadata: &religionMetadata{
				Lawful:  0.25,
				Ascetic: 0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Martyrdom",
			_religionMetadata: &religionMetadata{
				Ascetic:        0.25,
				Collectivistic: 0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, selectedTraits []*theologyTrait) bool {
				var addCoef float64
				for _, trait := range selectedTraits {
					switch trait.Name {
					case "Messiah":
						addCoef += pm.RandFloat64InRange(0.01, 0.125)
					case "Prophets":
						addCoef += pm.RandFloat64InRange(0.1, 0.2)
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Indulgences",
			_religionMetadata: &religionMetadata{
				Plutocratic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Holiness",
			_religionMetadata: &religionMetadata{
				Ascetic:         0.75,
				Individualistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				var addCoef float64
				for _, goal := range r.Doctrine.HighGoal.Goals {
					switch goal.Name {
					case "BringHolinessDownToTheWorld":
						addCoef += pm.RandFloat64InRange(0.03, 0.1)
					case "BecomePerfectAndSaints":
						addCoef += pm.RandFloat64InRange(0.05, 0.15)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Celibacy",
			_religionMetadata: &religionMetadata{
				Ascetic:         1,
				Individualistic: 0.5,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				var addCoef float64
				if r.Attributes.Clerics.HasClerics {
					if r.Attributes.Clerics.Limitations.Marriage.IsDisallowed() {
						addCoef += pm.RandFloat64InRange(0.05, 0.2)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Flagellantism",
			_religionMetadata: &religionMetadata{
				Ascetic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, selectedTraits []*theologyTrait) bool {
				var addCoef float64
				for _, trait := range selectedTraits {
					switch trait.Name {
					case "Repentance":
						fallthrough
					case "Martyrdom":
						addCoef += pm.RandFloat64InRange(0.05, 0.15)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "FeedTheWorld",
			_religionMetadata: &religionMetadata{
				Altruistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HolyArmy",
			_religionMetadata: &religionMetadata{
				Chthonic:       0.25,
				Aggressive:     1,
				Collectivistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				var addCoef float64
				if r.metadata.IsAggressive() {
					addCoef += pm.RandFloat64InRange(0.05, 0.125)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "DefendersOfFaith",
			_religionMetadata: &religionMetadata{
				Lawful:     0.25,
				Aggressive: 0.75,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "NonViolentResistance",
			_religionMetadata: &religionMetadata{
				Pacifistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				if r.metadata.IsAggressive() {
					return pm.GetRandomBool(pm.RandFloat64InRange(0, 0.05))
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "NoMoreKilling",
			_religionMetadata: &religionMetadata{
				Altruistic: 1,
				Pacifistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				if r.metadata.IsAggressive() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "AnimalMessengers",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SanctuaryOfMind",
			_religionMetadata: &religionMetadata{
				Individualistic: 1,
				Complicated:     0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Phantoms",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
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
		if trait.Name == "Reincarnation" {
			return true
		}
	}

	return false
}

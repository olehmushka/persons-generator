package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
)

type Theology struct {
	religion *Religion

	// traits []trait
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
	t.Cults = t.generateCults(1, pm.RandIntInRange(2, len(t.getAllCults())))
	t.Rules = t.generateRules()
	t.Taboos = t.generateTaboos()
	t.Rituals = t.generateRituals()
	t.Holydays = t.generateHolydays()
	t.Conversion = t.generateConversion()
	t.MarriageTradition = t.generateMarriageTradition()

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
	t.Taboos.Print()
	t.Rituals.Print()
	t.Holydays.Print()
	t.Conversion.Print()
	t.MarriageTradition.Print()
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
		for i, trait := range allTraits {
			if trait.Calc(t.religion, trait, traits) {
				traits = append(traits, &theologyTrait{
					_religionMetadata: trait._religionMetadata,
					baseCoef:          trait.baseCoef,
					Index:             i,
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
	}

	return traits
}

type theologyTrait struct {
	_religionMetadata *updateReligionMetadata
	baseCoef          float64
	Index             int
	Name              string
	Calc              func(r *Religion, self *theologyTrait, selectedTraits []*theologyTrait) bool
}

func (t *Theology) getAllTheologyTraits() []*theologyTrait {
	return []*theologyTrait{
		{
			Name: "Messiah",
			_religionMetadata: &updateReligionMetadata{
				Centralized:      Float64(0.05),
				RealLifeOriented: Float64(0.09),
				Fanaticism:       Float64(0.06),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
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

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Prophets",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.09),
				Fanaticism:       Float64(0.075),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, selectedTraits []*theologyTrait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "ReligiousLaw",
			_religionMetadata: &updateReligionMetadata{
				Centralized:      Float64(0.1),
				RealLifeOriented: Float64(0.075),
				OutsideDirected:  Float64(0.05),
				Strictness:       Float64(0.075),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				var addCoef float64
				if r.IsLawful() {
					addCoef += 0.5
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Reincarnation",
			_religionMetadata: &updateReligionMetadata{
				AfterlifeOriented: Float64(0.1),
				Strictness:        Float64(0.01),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				if r.HasReincarnation() {
					return true
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SanctionedFalseConversions", // If your life is threatened, it is acceptable to confess a false faith, as long as you keep the truth in your heart.
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				Hedonism:         Float64(0.05),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.1),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "TreeConnection", // Trees are the essence of life, and we must be near them.
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.045),
				Chthonic:         Float64(0.005),
				Primitive:        Float64(0.09),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				var addCoef float64
				if r.IsNaturalistic() {
					addCoef += 0.25
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "AnimalConnection",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.045),
				Chthonic:         Float64(0.0055),
				Primitive:        Float64(0.09),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				var addCoef float64
				if r.IsNaturalistic() {
					addCoef += 0.25
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Blindsight", // Only the blind can perceive true reality
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.045),
				InsideDirected:   Float64(0.09),
				Fanaticism:       Float64(0.08),
				Chthonic:         Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.5, 1),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Astrology",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.05),
				Elitaric:         Float64(0.05),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.15),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				var addCoef float64
				if r.IsNaturalistic() {
					addCoef += 0.05
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Repentance",
			_religionMetadata: &updateReligionMetadata{
				AfterlifeOriented: Float64(0.035),
				Strictness:        Float64(0.1),
				Ascetic:           Float64(0.025),
				Primitive:         Float64(0.01),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Martyrdom",
			_religionMetadata: &updateReligionMetadata{
				AfterlifeOriented: Float64(0.09),
				Fanaticism:        Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Indulgences",
			_religionMetadata: &updateReligionMetadata{
				Centralized:       Float64(0.05),
				AfterlifeOriented: Float64(0.05),
				Organized:         Float64(0.05),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				if r.Type.IsAtheism() {
					return false
				}

				var addCoef float64
				if r.IsPlutocratic() {
					addCoef += pm.RandFloat64InRange(0.15, 0.3)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Holiness",
			_religionMetadata: &updateReligionMetadata{
				AfterlifeOriented: Float64(0.065),
				InsideDirected:    Float64(0.05),
				Strictness:        Float64(0.075),
				Ascetic:           Float64(0.05),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
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

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Celibacy",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.075),
				Fanaticism:       Float64(0.01),
				Ascetic:          Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				var addCoef float64
				if r.Attributes.Clerics.HasClerics {
					if r.Attributes.Clerics.Limitations.Marriage.IsDisallowed() {
						addCoef += pm.RandFloat64InRange(0.05, 0.1)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Flagellantism",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.075),
				Fanaticism:       Float64(0.1),
				Ascetic:          Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.75, 1.1),
			Calc: func(r *Religion, self *theologyTrait, selectedTraits []*theologyTrait) bool {
				var addCoef float64
				for _, trait := range selectedTraits {
					switch trait.Name {
					case "Repentance":
						addCoef += pm.RandFloat64InRange(0.025, 0.075)
					case "Martyrdom":
						addCoef += pm.RandFloat64InRange(0.05, 0.1)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "FeedTheWorld",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.09),
				OutsideDirected:  Float64(0.09),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HolyArmy",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.045),
				Fanaticism:       Float64(0.1),
				Organized:        Float64(0.095),
			},
			baseCoef: pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				var addCoef float64
				if r.IsAggressive() {
					addCoef += pm.RandFloat64InRange(0.05, 0.125)
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "DefendersOfFaith",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "NonViolentResistance",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "NoMoreKilling",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "AnimalMessengers",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "SanctuaryOfMind",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Phantoms",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
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

func (t *Theology) GetNaturalisticCriterias() float64 {
	if len(t.Traits) == 0 {
		return 0
	}

	var criterias float64
	for _, trait := range t.Traits {
		switch trait.Name {
		case "TreeConnection":
			fallthrough
		case "AnimalConnection":
			criterias += 1
		case "Astrology":
			criterias += 0.5
		case "AnimalMessengers":
			criterias += 0.5
		}
	}

	return criterias
}

func (t *Theology) GetAggressiveCriterias() float64 {
	if len(t.Traits) == 0 {
		return 0
	}

	var criterias float64
	for _, trait := range t.Traits {
		switch trait.Name {
		case "HolyArmy":
			criterias += 1
		}
	}

	return criterias
}

func (t *Theology) GetPlutocracyCriterias() float64 {
	if len(t.Traits) == 0 {
		return 0
	}

	var criterias float64
	for _, trait := range t.Traits {
		switch trait.Name {
		case "Indulgences":
			criterias += 1
		}
	}

	return criterias
}

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
			Name:              "Messiah",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Prophets",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "ReligiousLaw",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Reincarnation",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "SanctionedFalseConversions", // If your life is threatened, it is acceptable to confess a false faith, as long as you keep the truth in your heart.
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "TreeConnection", // Trees are the essence of life, and we must be near them.
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "AnimalConnection",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Blindsight", // Only the blind can perceive true reality
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Astrology",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Repentance",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Martyrdom",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Indulgences",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Holiness",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Celibacy",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Flagellantism",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "FeedTheWorld",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "HolyArmy",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.8, 1.2),
			Calc: func(r *Religion, self *theologyTrait, _ []*theologyTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
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
		// trickster, idols, animal messengers, sanctuary of mind
		// phantoms, demonic deal, spirits of death
	}
}

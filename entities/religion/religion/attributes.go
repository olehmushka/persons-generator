package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
)

type Attributes struct {
	religion *Religion

	Traits        []*attributeTrait
	Clerics       *Clerics
	HolyScripture *HolyScripture
	Temples       *Temples
}

func NewAttributes(r *Religion) *Attributes {
	attrs := &Attributes{religion: r}
	attrs.Traits = attrs.generateTraits(0, len(attrs.Traits))
	attrs.Clerics = attrs.generateClerics()
	attrs.HolyScripture = attrs.generateHolyScripture()
	attrs.Temples = attrs.generateTemples()

	return attrs
}

func (as *Attributes) Print() {
	fmt.Printf("Attributes (religion_name=%s):\n", as.religion.Name)
	fmt.Printf("Attributes Traits (religion_name=%s):\n", as.religion.Name)
	for _, trait := range as.Traits {
		fmt.Printf(" - %s\n", trait.Name)
	}
	as.Clerics.Print()
	as.HolyScripture.Print()
	as.Temples.Print()
}

func (as *Attributes) generateTraits(min, max int) []*attributeTrait {
	if min < 0 {
		panic("[Attributes.generateTraits] min can not be less than 0")
	}
	allTraits := as.getAllAttributeTraits()
	if max > len(allTraits) {
		panic("[Attributes.generateTraits] max can not be greater than allTraits length")
	}

	traits := make([]*attributeTrait, 0, len(allTraits))
	for count := 0; count < 20; count++ {
		for _, trait := range allTraits {
			if trait.Calc(as.religion, trait, traits) {
				traits = append(traits, &attributeTrait{
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
		as.religion.UpdateMetadata(UpdateReligionMetadata(*as.religion.metadata, *trait._religionMetadata))
	}

	return traits
}

type attributeTrait struct {
	_religionMetadata *updateReligionMetadata
	baseCoef          float64
	Name              string
	Calc              func(r *Religion, self *attributeTrait, selectedTraits []*attributeTrait) bool
}

func (as *Attributes) getAllAttributeTraits() []*attributeTrait {
	return []*attributeTrait{
		{
			Name: "Amulets",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.02),
				InsideDirected:   Float64(0.02),
				OutsideDirected:  Float64(0.02),
				Primitive:        Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *attributeTrait, _ []*attributeTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "ReligionMusic",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented:  Float64(0.01),
				AfterlifeOriented: Float64(0.01),
				OutsideDirected:   Float64(0.05),
				Primitive:         Float64(0.05),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *attributeTrait, _ []*attributeTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "ReligionTheatre",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented:  Float64(0.01),
				AfterlifeOriented: Float64(0.01),
				OutsideDirected:   Float64(0.05),
				Primitive:         Float64(0.05),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *attributeTrait, _ []*attributeTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "ReligionPoetry",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented:  Float64(0.01),
				AfterlifeOriented: Float64(0.01),
				OutsideDirected:   Float64(0.05),
				Primitive:         Float64(0.04),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *attributeTrait, _ []*attributeTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "ReligionTapestry",
			_religionMetadata: &updateReligionMetadata{
				OutsideDirected: Float64(0.05),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *attributeTrait, _ []*attributeTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Calligraphy",
			_religionMetadata: &updateReligionMetadata{
				OutsideDirected: Float64(0.02),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *attributeTrait, _ []*attributeTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Idols",
			_religionMetadata: &updateReligionMetadata{
				OutsideDirected: Float64(0.02),
				Primitive:       Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.95, 1.05),
			Calc: func(r *Religion, self *attributeTrait, _ []*attributeTrait) bool {
				var addCoef float64
				switch {
				case r.Type.IsMonotheism():
					addCoef -= 0.05
				case r.Type.IsPolytheism():
					addCoef += 0.1
				case r.Type.IsAtheism():
					return false
				}

				return CalculateProbabilityFromReligionMetadata(pm.PrepareProbability(self.baseCoef+addCoef), r, *self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

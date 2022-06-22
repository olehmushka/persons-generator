package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
)

type HolyScripture struct {
	religion *Religion
	attrs    *Attributes

	HasHolyScripture bool
	Traits           []*holyScriptureTrait
}

func (as *Attributes) generateHolyScripture() *HolyScripture {
	hs := &HolyScripture{religion: as.religion, attrs: as}
	hs.HasHolyScripture = hs.generateHasHolyScripture()
	if hs.HasHolyScripture {
		hs.Traits = hs.generateTraits(1, 3)
	}

	return hs
}

func (hs *HolyScripture) generateHasHolyScripture() bool {
	primaryProbability := pm.RandFloat64InRange(0.5, 0.7)
	if hs.religion.metadata.Organized > 0 {
		primaryProbability += pm.RandFloat64InRange(0.05, 0.1)
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProbability))
}

func (hs *HolyScripture) Print() {
	if !hs.HasHolyScripture {
		fmt.Printf("Has not HolyScripture (religion_name=%s):\n", hs.religion.Name)
		return
	}
	fmt.Printf("HolyScripture Traits (religion_name=%s):\n", hs.religion.Name)
	for _, trait := range hs.Traits {
		fmt.Printf(" - %s\n", trait.Name)
	}
}

type holyScriptureTrait struct {
	_religionMetadata *updateReligionMetadata
	baseCoef          float64
	Name              string
	Calc              func(r *Religion, self *holyScriptureTrait, selectedTraits []*holyScriptureTrait) bool
}

func (hs *HolyScripture) generateTraits(min, max int) []*holyScriptureTrait {
	if min < 0 {
		panic("[HolyScripture.generateTraits] min can not be less than 0")
	}
	allTraits := hs.getAllHolyScriptureTraits()
	if max > len(allTraits) {
		panic("[HolyScripture.generateTraits] max can not be greater than allTraits length")
	}

	traits := make([]*holyScriptureTrait, 0, len(allTraits))
	for count := 0; count < 20; count++ {
		for _, trait := range allTraits {
			if trait.Calc(hs.religion, trait, traits) {
				traits = append(traits, &holyScriptureTrait{
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
		hs.religion.UpdateMetadata(UpdateReligionMetadata(*hs.religion.metadata, *trait._religionMetadata))
	}

	return traits
}

func (hs *HolyScripture) getAllHolyScriptureTraits() []*holyScriptureTrait {
	return []*holyScriptureTrait{
		{
			Name: "Commandments",
			_religionMetadata: &updateReligionMetadata{
				Centralized: Float64(0.09),
				Strictness:  Float64(0.09),
				Organized:   Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *holyScriptureTrait, _ []*holyScriptureTrait) bool {
				var addCoef float64
				switch {
				case r.Type.IsMonotheism():
					addCoef += pm.RandFloat64InRange(0.055, 0.105)
				case r.Type.IsDeityDualism():
					addCoef += pm.RandFloat64InRange(0.04, 0.09)
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "DevotionalCode",
			_religionMetadata: &updateReligionMetadata{
				Fanaticism: Float64(0.02),
				Strictness: Float64(0.085),
				Elitaric:   Float64(0.02),
				Organized:  Float64(0.01),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *holyScriptureTrait, _ []*holyScriptureTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "DivineLaw",
			_religionMetadata: &updateReligionMetadata{
				Centralized: Float64(0.03),
				Fanaticism:  Float64(0.09),
				Strictness:  Float64(0.1),
				Organized:   Float64(0.01),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *holyScriptureTrait, selectedTraits []*holyScriptureTrait) bool {
				var addCoef float64
				for _, trait := range selectedTraits {
					if trait.Name == "Commandments" {
						addCoef += 0.075
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "MagicBooks",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.02),
				Primitive:        Float64(0.05),
				Chthonic:         Float64(0.07),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *holyScriptureTrait, _ []*holyScriptureTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Runes",
			_religionMetadata: &updateReligionMetadata{
				Elitaric:  Float64(0.05),
				Organized: Float64(0.05),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *holyScriptureTrait, _ []*holyScriptureTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "MarchingHymns",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.075),
				OutsideDirected:  Float64(0.08),
				Fanaticism:       Float64(0.06),
				Primitive:        Float64(0.01),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *holyScriptureTrait, _ []*holyScriptureTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Theogony",
			_religionMetadata: &updateReligionMetadata{
				Centralized: Float64(0.04),
				Fanaticism:  Float64(0.01),
				Elitaric:    Float64(0.03),
				Organized:   Float64(0.05),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *holyScriptureTrait, _ []*holyScriptureTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

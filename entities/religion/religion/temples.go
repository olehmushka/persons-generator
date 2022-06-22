package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
)

type Temples struct {
	religion *Religion
	attrs    *Attributes

	HasTemples      bool
	HasSacredPlaces bool
	Traits          []*templeTrait
}

func (as *Attributes) generateTemples() *Temples {
	ts := &Temples{religion: as.religion, attrs: as}
	ts.HasTemples = ts.generateHasTemples()
	ts.HasSacredPlaces = ts.generateSacredPlaces()
	if ts.HasTemples || ts.HasSacredPlaces {
		ts.Traits = ts.generateTraits(1, len(ts.Traits)-2)
	}

	return ts
}

func (ts *Temples) Print() {
	fmt.Printf("Temples (religion_name=%s):\n", ts.religion.Name)

	if !ts.HasTemples {
		fmt.Printf("Has not Temples (religion_name=%s):\n", ts.religion.Name)
	}
	if !ts.HasSacredPlaces {
		fmt.Printf("Has not HasSacredPlaces (religion_name=%s):\n", ts.religion.Name)
	}
	if !ts.HasTemples && !ts.HasSacredPlaces {
		return
	}
	fmt.Printf("Temples Traits (religion_name=%s):\n", ts.religion.Name)
	for _, trait := range ts.Traits {
		fmt.Printf(" - %s\n", trait.Name)
	}
}

func (ts *Temples) generateHasTemples() bool {
	primaryProbability := pm.RandFloat64InRange(0.3, 0.5)
	if ts.religion.metadata.Organized > 0 {
		primaryProbability += pm.RandFloat64InRange(0.05, 0.15)
	}
	if ts.attrs.Clerics.HasClerics {
		primaryProbability += pm.RandFloat64InRange(0.1, 0.2)
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProbability))
}

func (ts *Temples) generateSacredPlaces() bool {
	primaryProbability := pm.RandFloat64InRange(0.5, 0.7)
	if ts.religion.metadata.Organized > 0 {
		primaryProbability += pm.RandFloat64InRange(0.05, 0.1)
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProbability))
}

type templeTrait struct {
	_religionMetadata *updateReligionMetadata
	baseCoef          float64
	Name              string
	Calc              func(r *Religion, self *templeTrait, selectedTraits []*templeTrait) bool
}

func (ts *Temples) generateTraits(min, max int) []*templeTrait {
	if min < 0 {
		panic("[Temples.generateTraits] min can not be less than 0")
	}
	allTraits := ts.getAllTemplesTraits()
	if max > len(allTraits) {
		panic("[Temples.generateTraits] max can not be greater than allTraits length")
	}

	traits := make([]*templeTrait, 0, len(allTraits))
	for count := 0; count < 20; count++ {
		for _, trait := range allTraits {
			if trait.Calc(ts.religion, trait, traits) {
				traits = append(traits, &templeTrait{
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
		ts.religion.UpdateMetadata(UpdateReligionMetadata(*ts.religion.metadata, *trait._religionMetadata))
	}

	return traits
}

func (hs *Temples) getAllTemplesTraits() []*templeTrait {
	return []*templeTrait{
		{
			Name: "CultProperty",
			_religionMetadata: &updateReligionMetadata{
				Centralized:      Float64(0.09),
				RealLifeOriented: Float64(0.05),
				Organized:        Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *templeTrait, _ []*templeTrait) bool {
				var addCoef float64
				for _, trait := range hs.attrs.HolyScripture.Traits {
					if trait.Name == "DivineLaw" {
						addCoef += pm.RandFloat64InRange(0.01, 0.07)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "InviolacyOfTemples",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				Organized:        Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *templeTrait, _ []*templeTrait) bool {
				if !hs.HasTemples {
					return false
				}
				var addCoef float64
				for _, trait := range hs.attrs.HolyScripture.Traits {
					if trait.Name == "DivineLaw" {
						addCoef += pm.RandFloat64InRange(0.03, 0.09)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HasSeminaries",
			_religionMetadata: &updateReligionMetadata{
				Centralized: Float64(0.09),
				Strictness:  Float64(0.01),
				Organized:   Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *templeTrait, _ []*templeTrait) bool {
				var addCoef float64
				if hs.attrs.HolyScripture.HasHolyScripture {
					addCoef += 0.1
				}
				if hs.attrs.Clerics.HasClerics {
					for _, trait := range hs.attrs.Clerics.Traits {
						if trait.Name == "Patronship" {
							addCoef += pm.RandFloat64InRange(0.01, 0.05)
						}
					}
					for _, function := range hs.attrs.Clerics.Functions {
						if function.Name == "Teach" {
							addCoef += pm.RandFloat64InRange(0.05, 0.15)
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HasLibraries",
			_religionMetadata: &updateReligionMetadata{
				Centralized: Float64(0.01),
				Organized:   Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *templeTrait, _ []*templeTrait) bool {
				var addCoef float64
				if hs.attrs.HolyScripture.HasHolyScripture {
					addCoef += 0.105
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HasSchoolsOfPhilosophers",
			_religionMetadata: &updateReligionMetadata{
				Centralized: Float64(0.01),
				Elitaric:    Float64(0.03),
				Organized:   Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *templeTrait, _ []*templeTrait) bool {
				var addCoef float64
				if hs.attrs.HolyScripture.HasHolyScripture {
					addCoef += 0.105
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HasGymnasiumSchools",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.06),
				Strictness:       Float64(0.01),
				Organized:        Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *templeTrait, _ []*templeTrait) bool {
				var addCoef float64
				if hs.attrs.Clerics.HasClerics {
					for _, trait := range hs.attrs.Clerics.Traits {
						if trait.Name == "Patronship" {
							addCoef += pm.RandFloat64InRange(0.02, 0.06)
						}
					}
					for _, function := range hs.attrs.Clerics.Functions {
						if function.Name == "Teach" {
							addCoef += pm.RandFloat64InRange(0.02, 0.1)
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HasShelters",
			_religionMetadata: &updateReligionMetadata{
				RealLifeOriented: Float64(0.1),
				Fanaticism:       Float64(0.1),
				Organized:        Float64(0.05),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *templeTrait, _ []*templeTrait) bool {
				var addCoef float64
				if hs.attrs.Clerics.HasClerics {
					for _, trait := range hs.attrs.Clerics.Traits {
						if trait.Name == "Patronship" {
							addCoef += pm.RandFloat64InRange(0.065, 0.105)
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "TempleHealers",
			_religionMetadata: &updateReligionMetadata{
				Decentralized:    Float64(0.035),
				RealLifeOriented: Float64(0.065),
				Fanaticism:       Float64(0.05),
				Primitive:        Float64(0.075),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *templeTrait, _ []*templeTrait) bool {
				if !hs.HasTemples {
					return false
				}
				var addCoef float64
				if hs.attrs.Clerics.HasClerics {
					for _, function := range hs.attrs.Clerics.Functions {
						if function.Name == "Heal" {
							addCoef += pm.RandFloat64InRange(0.1, 0.2)
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SacredStones",
			_religionMetadata: &updateReligionMetadata{
				Decentralized: Float64(0.04),
				Chthonic:      Float64(0.035),
				Primitive:     Float64(0.1),
			},
			baseCoef: pm.RandFloat64InRange(0.9, 1.1),
			Calc: func(r *Religion, self *templeTrait, _ []*templeTrait) bool {
				if !hs.HasSacredPlaces {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

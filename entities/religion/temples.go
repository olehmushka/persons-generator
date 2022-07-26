package religion

import (
	"fmt"

	pm "persons_generator/probability_machine"
)

type Temples struct {
	religion *Religion
	attrs    *Attributes

	HasTemples      bool
	HasSacredPlaces bool
	Traits          []*trait
}

func (as *Attributes) generateTemples() *Temples {
	ts := &Temples{religion: as.religion, attrs: as}
	ts.HasTemples = ts.generateHasTemples()
	ts.HasSacredPlaces = ts.generateSacredPlaces()
	if ts.HasTemples || ts.HasSacredPlaces {
		ts.Traits = generateTraits(ts.religion, ts.getAllTemplesTraits(), generateTraitsOpts{
			Label: "Temples.generateTraits",
			Min:   1,
			Max:   len(ts.getAllTemplesTraits()) - 2,
		})
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
	primaryProbability := pm.RandFloat64InRange(0.35, 0.55)
	if ts.attrs.Clerics.HasClerics {
		primaryProbability += pm.RandFloat64InRange(0.15, 0.25)
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProbability))
}

func (ts *Temples) generateSacredPlaces() bool {
	primaryProbability := pm.RandFloat64InRange(0.5, 0.7)
	if ts.religion.metadata.IsNaturalistic() {
		primaryProbability += pm.RandFloat64InRange(0.05, 0.15)
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProbability))
}

func (ts *Temples) getAllTemplesTraits() []*trait {
	return []*trait{
		{
			Name: "CultProperty",
			_religionMetadata: &religionMetadata{
				Plutocratic: 0.75,
				Lawful:      0.25,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				for _, trait := range ts.attrs.HolyScripture.Traits {
					if trait.Name == "DivineLaw" {
						addCoef += pm.RandFloat64InRange(0.04, 0.1)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "InviolacyOfTemples",
			_religionMetadata: &religionMetadata{
				Plutocratic: 1,
				Lawful:      0.25,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				if !ts.HasTemples {
					return false
				}

				var addCoef float64
				for _, trait := range ts.attrs.HolyScripture.Traits {
					if trait.Name == "DivineLaw" {
						addCoef += pm.RandFloat64InRange(0.04, 0.1)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HasSeminaries",
			_religionMetadata: &religionMetadata{
				Educational:    1,
				Collectivistic: 0.25,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				if ts.attrs.HolyScripture.HasHolyScripture {
					addCoef += pm.RandFloat64InRange(0.05, 0.15)
				}
				if ts.attrs.Clerics.HasClerics {
					for _, trait := range ts.attrs.Clerics.Traits {
						if trait.Name == "Patronship" {
							addCoef += pm.RandFloat64InRange(0.04, 0.1)
						}
					}
					for _, function := range ts.attrs.Clerics.Functions {
						if function.Name == "Teach" {
							addCoef += pm.RandFloat64InRange(0.05, 0.15)
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HasLibraries",
			_religionMetadata: &religionMetadata{
				Educational: 1,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				if ts.attrs.HolyScripture.HasHolyScripture {
					addCoef += pm.RandFloat64InRange(0.05, 0.15)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HasSchoolsOfPhilosophers",
			_religionMetadata: &religionMetadata{
				Educational: 1,
				Complicated: 0.5,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				if ts.attrs.HolyScripture.HasHolyScripture {
					addCoef += pm.RandFloat64InRange(0.05, 0.15)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HasGymnasiumSchools",
			_religionMetadata: &religionMetadata{
				Educational:     0.75,
				Authoritaristic: 0.25,
				Collectivistic:  0.75,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				if ts.attrs.Clerics.HasClerics {
					for _, trait := range ts.attrs.Clerics.Traits {
						if trait.Name == "Patronship" {
							addCoef += pm.RandFloat64InRange(0.02, 0.06)
						}
					}
					for _, function := range ts.attrs.Clerics.Functions {
						if function.Name == "Teach" {
							addCoef += pm.RandFloat64InRange(0.02, 0.1)
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HasShelters",
			_religionMetadata: &religionMetadata{
				Altruistic:     1,
				Collectivistic: 0.75,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				if ts.attrs.Clerics.HasClerics {
					for _, trait := range ts.attrs.Clerics.Traits {
						if trait.Name == "Patronship" {
							addCoef += pm.RandFloat64InRange(0.065, 0.105)
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "TempleHealers",
			_religionMetadata: &religionMetadata{
				Altruistic: 1,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				if !ts.HasTemples {
					return false
				}
				var addCoef float64
				if ts.attrs.Clerics.HasClerics {
					for _, function := range ts.attrs.Clerics.Functions {
						switch function.Name {
						case "Heal":
							addCoef += pm.RandFloat64InRange(0.1, 0.2)
						case "Oracle":
							addCoef += pm.RandFloat64InRange(0.01, 0.065)
						case "Diviner":
							addCoef += pm.RandFloat64InRange(0.025, 0.1)
						case "Druid":
							addCoef += pm.RandFloat64InRange(0.015, 0.9)
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SacredStones",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.5,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				if !ts.HasSacredPlaces {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "TempleProstitution",
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Plutocratic:  0.75,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				if r.Theology != nil && r.Theology.Taboos != nil && len(r.Theology.Taboos.Taboos) > 0 {
					for _, taboo := range r.Theology.Taboos.Taboos {
						if taboo == nil {
							continue
						}
						if taboo.Name == "Prostitution" {
							switch taboo.Acceptance {
							case Accepted:
								addCoef += pm.RandFloat64InRange(0.1, 0.2)
							case Shunned:
								addCoef += -pm.RandFloat64InRange(0.01, 0.1)
							case Criminal:
								return false
							}
						}
					}
				}
				if addCoef <= 0 {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

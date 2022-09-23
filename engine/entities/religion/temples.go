package religion

import (
	"fmt"

	"persons_generator/engine/entities/culture"
	pm "persons_generator/engine/probability_machine"
)

type Temples struct {
	religion *Religion   `json:"-"`
	attrs    *Attributes `json:"-"`

	HasTemples      bool     `json:"has_temples"`
	HasSacredPlaces bool     `json:"has_sacred_places"`
	Traits          []*trait `json:"traits"`
}

func (as *Attributes) generateTemples(c *culture.Culture) (*Temples, error) {
	ts := &Temples{religion: as.religion, attrs: as}
	hasTemples, err := ts.generateHasTemples(c)
	if err != nil {
		return nil, err
	}
	ts.HasTemples = hasTemples
	hasSacredPlaces, err := ts.generateSacredPlaces()
	if err != nil {
		return nil, err
	}
	ts.HasSacredPlaces = hasSacredPlaces
	if ts.HasTemples || ts.HasSacredPlaces {
		traits, err := generateTraits(ts.religion, ts.getAllTemplesTraits(), generateTraitsOpts{
			Label: "Temples.generateTraits",
			Min:   1,
			Max:   len(ts.getAllTemplesTraits()) - 2,
		})
		if err != nil {
			return nil, err
		}
		ts.Traits = traits
	}

	return ts, nil
}

func (ts *Temples) SerializeTraits() []string {
	if ts == nil || !ts.HasSacredPlaces || !ts.HasTemples {
		return []string{}
	}

	return extractTraitNames(ts.Traits)
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

func (ts *Temples) generateHasTemples(c *culture.Culture) (bool, error) {
	if hasTemples := getHasTemplesByCulture(c); hasTemples != nil {
		return *hasTemples, nil
	}

	primaryProbability, err := pm.RandFloat64InRange(0.35, 0.55)
	if err != nil {
		return false, err
	}
	if ts.attrs.Clerics.HasClerics {
		p, err := pm.RandFloat64InRange(0.15, 0.25)
		if err != nil {
			return false, err
		}
		primaryProbability += p
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProbability))
}

func getHasTemplesByCulture(c *culture.Culture) *bool {
	if c == nil {
		return nil
	}
	tPtr := true

	for _, t := range c.Traditions {
		if t == nil {
			continue
		}

		if t.Name == culture.FerventTempleBuildersTradition.Name {
			return &tPtr
		}
	}

	return nil
}

func (ts *Temples) generateSacredPlaces() (bool, error) {
	primaryProbability, err := pm.RandFloat64InRange(0.5, 0.7)
	if err != nil {
		return false, err
	}
	if ts.religion.Metadata.IsNaturalistic() {
		p, err := pm.RandFloat64InRange(0.05, 0.15)
		if err != nil {
			return false, err
		}
		primaryProbability += p
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProbability))
}

func (ts *Temples) getAllTemplesTraits() []*trait {
	return []*trait{
		{
			Name: CultPropertyTempleTrait,
			ReligionMetadata: &religionMetadata{
				Plutocratic: 0.75,
				Lawful:      0.25,
			},
			BaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				for _, trait := range ts.attrs.HolyScripture.Traits {
					if trait.Name == DivineLawHolyScriptureTrait {
						coef, err := pm.RandFloat64InRange(0.04, 0.1)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: InviolacyOfTemplesTempleTrait,
			ReligionMetadata: &religionMetadata{
				Plutocratic: 1,
				Lawful:      0.25,
			},
			BaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if !ts.HasTemples {
					return false, nil
				}

				var addCoef float64
				for _, trait := range ts.attrs.HolyScripture.Traits {
					if trait.Name == DivineLawHolyScriptureTrait {
						coef, err := pm.RandFloat64InRange(0.04, 0.1)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HasSeminariesTempleTrait,
			ReligionMetadata: &religionMetadata{
				Educational:    1,
				Collectivistic: 0.25,
			},
			BaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if ts.attrs.HolyScripture.HasHolyScripture {
					coef, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}
				if ts.attrs.Clerics.HasClerics {
					for _, trait := range ts.attrs.Clerics.Traits {
						if trait.Name == PatronshipClericTrait {
							coef, err := pm.RandFloat64InRange(0.04, 0.1)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
					for _, function := range ts.attrs.Clerics.Functions {
						if function.Name == TeachClericFunction {
							coef, err := pm.RandFloat64InRange(0.05, 0.15)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HasLibrariesTempleTrait,
			ReligionMetadata: &religionMetadata{
				Educational: 1,
			},
			BaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if ts.attrs.HolyScripture.HasHolyScripture {
					coef, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HasSchoolsOfPhilosophersTempleTrait,
			ReligionMetadata: &religionMetadata{
				Educational: 1,
				Complicated: 0.5,
			},
			BaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if ts.attrs.HolyScripture.HasHolyScripture {
					coef, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HasGymnasiumSchoolsTempleTrait,
			ReligionMetadata: &religionMetadata{
				Educational:     0.75,
				Authoritaristic: 0.25,
				Collectivistic:  0.75,
			},
			BaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if ts.attrs.Clerics.HasClerics {
					for _, trait := range ts.attrs.Clerics.Traits {
						if trait.Name == PatronshipClericTrait {
							coef, err := pm.RandFloat64InRange(0.02, 0.06)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
					for _, function := range ts.attrs.Clerics.Functions {
						if function.Name == TeachClericFunction {
							coef, err := pm.RandFloat64InRange(0.02, 0.1)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HasSheltersTempleTrait,
			ReligionMetadata: &religionMetadata{
				Altruistic:     1,
				Collectivistic: 0.75,
			},
			BaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if ts.attrs.Clerics.HasClerics {
					for _, trait := range ts.attrs.Clerics.Traits {
						if trait.Name == PatronshipClericTrait {
							coef, err := pm.RandFloat64InRange(0.065, 0.105)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: TempleHealersTempleTrait,
			ReligionMetadata: &religionMetadata{
				Altruistic: 1,
			},
			BaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if !ts.HasTemples {
					return false, nil
				}
				var addCoef float64
				if ts.attrs.Clerics.HasClerics {
					for _, function := range ts.attrs.Clerics.Functions {
						switch function.Name {
						case HealClericFuction:
							coef, err := pm.RandFloat64InRange(0.1, 0.2)
							if err != nil {
								return false, err
							}
							addCoef += coef
						case OracleClericFunction:
							coef, err := pm.RandFloat64InRange(0.01, 0.065)
							if err != nil {
								return false, err
							}
							addCoef += coef
						case DivinerClericFunction:
							coef, err := pm.RandFloat64InRange(0.025, 0.1)
							if err != nil {
								return false, err
							}
							addCoef += coef
						case DruidClericFunction:
							coef, err := pm.RandFloat64InRange(0.015, 0.9)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SacredStonesTempleTrait,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.5,
			},
			BaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if !ts.HasSacredPlaces {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: TempleProstitutionTempleTrait,
			ReligionMetadata: &religionMetadata{
				SexualActive: 1,
				Plutocratic:  0.75,
			},
			BaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.Theology != nil && r.Theology.Taboos != nil && len(r.Theology.Taboos.Taboos) > 0 {
					for _, taboo := range r.Theology.Taboos.Taboos {
						if taboo == nil {
							continue
						}
						if taboo.Name == ProstitutionTabooName {
							switch taboo.Acceptance {
							case Accepted:
								coef, err := pm.RandFloat64InRange(0.1, 0.2)
								if err != nil {
									return false, err
								}
								addCoef += coef
							case Shunned:
								coef, err := pm.RandFloat64InRange(0.01, 0.1)
								if err != nil {
									return false, err
								}
								addCoef -= coef
							case Criminal:
								return false, nil
							}
						}
					}
				}
				if addCoef <= 0 {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func GetTemplesSimilarityCoef(t1, t2 *Temples) float64 {
	if t1 == nil && t2 == nil {
		return 1
	}
	if t1 == nil || t2 == nil {
		return 0
	}

	if t1.HasTemples != t2.HasTemples && t1.HasSacredPlaces != t2.HasSacredPlaces {
		return 0
	}
	if !t1.HasTemples && !t1.HasSacredPlaces {
		return 1
	}

	return (GetTraitsSimilarityCoef(t1.Traits, t2.Traits) + 1) / 2
}

package religion

import (
	"fmt"

	"persons_generator/engine/entities/culture"
	pm "persons_generator/engine/probability_machine"
)

type Temples struct {
	religion *Religion
	attrs    *Attributes

	HasTemples      bool
	HasSacredPlaces bool
	Traits          []*trait
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
	if ts.religion.metadata.IsNaturalistic() {
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
			Name: "cult_property",
			_religionMetadata: &religionMetadata{
				Plutocratic: 0.75,
				Lawful:      0.25,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				for _, trait := range ts.attrs.HolyScripture.Traits {
					if trait.Name == "divine_law" {
						coef, err := pm.RandFloat64InRange(0.04, 0.1)
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
			Name: "inviolacy_of_temples",
			_religionMetadata: &religionMetadata{
				Plutocratic: 1,
				Lawful:      0.25,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if !ts.HasTemples {
					return false, nil
				}

				var addCoef float64
				for _, trait := range ts.attrs.HolyScripture.Traits {
					if trait.Name == "divine_law" {
						coef, err := pm.RandFloat64InRange(0.04, 0.1)
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
			Name: "has_seminaries",
			_religionMetadata: &religionMetadata{
				Educational:    1,
				Collectivistic: 0.25,
			},
			baseCoef: ts.religion.M.BaseCoef,
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
						if trait.Name == "patronship" {
							coef, err := pm.RandFloat64InRange(0.04, 0.1)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
					for _, function := range ts.attrs.Clerics.Functions {
						if function.Name == "teach" {
							coef, err := pm.RandFloat64InRange(0.05, 0.15)
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
			Name: "has_libraries",
			_religionMetadata: &religionMetadata{
				Educational: 1,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if ts.attrs.HolyScripture.HasHolyScripture {
					coef, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "has_schools_of_philosophers",
			_religionMetadata: &religionMetadata{
				Educational: 1,
				Complicated: 0.5,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if ts.attrs.HolyScripture.HasHolyScripture {
					coef, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "has_gymnasium_schools",
			_religionMetadata: &religionMetadata{
				Educational:     0.75,
				Authoritaristic: 0.25,
				Collectivistic:  0.75,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if ts.attrs.Clerics.HasClerics {
					for _, trait := range ts.attrs.Clerics.Traits {
						if trait.Name == "patronship" {
							coef, err := pm.RandFloat64InRange(0.02, 0.06)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
					for _, function := range ts.attrs.Clerics.Functions {
						if function.Name == "teach" {
							coef, err := pm.RandFloat64InRange(0.02, 0.1)
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
			Name: "has_shelters",
			_religionMetadata: &religionMetadata{
				Altruistic:     1,
				Collectivistic: 0.75,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if ts.attrs.Clerics.HasClerics {
					for _, trait := range ts.attrs.Clerics.Traits {
						if trait.Name == "patronship" {
							coef, err := pm.RandFloat64InRange(0.065, 0.105)
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
			Name: "temple_healers",
			_religionMetadata: &religionMetadata{
				Altruistic: 1,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if !ts.HasTemples {
					return false, nil
				}
				var addCoef float64
				if ts.attrs.Clerics.HasClerics {
					for _, function := range ts.attrs.Clerics.Functions {
						switch function.Name {
						case "heal":
							coef, err := pm.RandFloat64InRange(0.1, 0.2)
							if err != nil {
								return false, err
							}
							addCoef += coef
						case "oracle":
							coef, err := pm.RandFloat64InRange(0.01, 0.065)
							if err != nil {
								return false, err
							}
							addCoef += coef
						case "diviner":
							coef, err := pm.RandFloat64InRange(0.025, 0.1)
							if err != nil {
								return false, err
							}
							addCoef += coef
						case "druid":
							coef, err := pm.RandFloat64InRange(0.015, 0.9)
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
			Name: "sacred_stones",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.5,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if !ts.HasSacredPlaces {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "temple_prostitution",
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Plutocratic:  0.75,
			},
			baseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.Theology != nil && r.Theology.Taboos != nil && len(r.Theology.Taboos.Taboos) > 0 {
					for _, taboo := range r.Theology.Taboos.Taboos {
						if taboo == nil {
							continue
						}
						if taboo.Name == "prostitution" {
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

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

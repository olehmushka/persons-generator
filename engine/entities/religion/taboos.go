package religion

import (
	"fmt"

	"persons_generator/engine/entities/culture"
	pm "persons_generator/engine/probability_machine"
)

type Taboos struct {
	religion *Religion
	theology *Theology

	Taboos []*Taboo
}

func (t *Theology) generateTaboos(c *culture.Culture) (*Taboos, error) {
	ts := &Taboos{religion: t.religion, theology: t}
	taboos, err := ts.generateTaboos(c)
	if err != nil {
		return nil, err
	}
	ts.Taboos = taboos

	return ts, nil
}

func (ts *Taboos) Print() {
	fmt.Printf("Taboos (religion_name=%s):\n", ts.religion.Name)
	for _, taboo := range ts.Taboos {
		fmt.Printf(" - %s (%s)\n", taboo.Name, taboo.Acceptance)
	}
}

type Taboo struct {
	_religionMetadata                                   *religionMetadata
	acceptedBaseCoef, shunnedBaseCoef, criminalBaseCoef float64

	Name       string
	Acceptance Acceptance
	Calc       func(r *Religion, self *Taboo, selectedTaboos []*Taboo) (Acceptance, error)
}

func (ts *Taboos) generateTaboos(c *culture.Culture) ([]*Taboo, error) {
	allTaboos := ts.getAllTaboos(c)
	taboos := make([]*Taboo, len(allTaboos))
	for i, taboo := range allTaboos {
		a, err := taboo.Calc(ts.religion, taboo, taboos)
		if err != nil {
			return nil, err
		}
		taboos[i] = &Taboo{
			_religionMetadata: taboo._religionMetadata,
			acceptedBaseCoef:  taboo.acceptedBaseCoef,
			shunnedBaseCoef:   taboo.shunnedBaseCoef,
			criminalBaseCoef:  taboo.criminalBaseCoef,
			Name:              taboo.Name,
			Acceptance:        a,
			Calc:              taboo.Calc,
		}

		md, err := UpdateReligionMetadataWithAcceptance(ts.religion, *ts.religion.metadata, *taboos[i]._religionMetadata, taboos[i].Acceptance)
		if err != nil {
			return nil, err
		}
		ts.religion.UpdateMetadata(md)
	}

	return taboos, nil
}

func (ts *Taboos) getAllTaboos(c *culture.Culture) []*Taboo {
	return []*Taboo{
		{
			Name: RaisingAnimalsTabooName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Plutocratic:  0.1,
			},
			acceptedBaseCoef: ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:  ts.religion.M.LowBaseCoef,
			criminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: RaisingPlantsTabooName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Altruistic:   0.5,
				Plutocratic:  0.25,
			},
			acceptedBaseCoef: ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:  ts.religion.M.LowBaseCoef,
			criminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: RaisingFungusTabooName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Chthonic:     0.1,
				Plutocratic:  0.1,
			},
			acceptedBaseCoef: ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: EatAnimalsTabooName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Chthonic:     0.1,
				Hedonistic:   0.1,
			},
			acceptedBaseCoef: ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: EatSomeAnimalsTabooName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Lawful:       0.25,
			},
			acceptedBaseCoef: ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) (Acceptance, error) {
				var shunnedAddCoef, criminalAddCoef float64
				for _, taboo := range selectedTaboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == EatAnimalsTabooName {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoefP, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return "", err
							}
							shunnedAddCoef += shunnedAddCoefP
							criminalAddCoefP, err := pm.RandFloat64InRange(0.03, 0.1)
							if err != nil {
								return "", err
							}
							criminalAddCoef += criminalAddCoefP
						case Criminal:
							return Criminal, nil
						}
					}
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: EatInsectsTabooName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Chthonic:     0.1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: EatFungusTabooName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
			},
			acceptedBaseCoef: ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: CannibalismTabooName,
			_religionMetadata: &religionMetadata{
				Chthonic:   1,
				Aggressive: 0.25,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.HighBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DrinkingBloodTabooName,
			_religionMetadata: &religionMetadata{
				Chthonic:   1,
				Aggressive: 0.25,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DrinkingStrongAlcoholTabooName,
			_religionMetadata: &religionMetadata{
				SexualActive: 0.1,
				Chthonic:     0.25,
				Aggressive:   0.1,
				Hedonistic:   1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DrinkingNotStrongAlcoholTabooName,
			_religionMetadata: &religionMetadata{
				SexualActive: 0.25,
				Chthonic:     0.1,
				Hedonistic:   1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) (Acceptance, error) {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				for _, taboo := range selectedTaboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == DrinkingStrongAlcoholTabooName {
						switch taboo.Acceptance {
						case Accepted:
							acceptedAddCoefP, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return "", err
							}
							acceptedAddCoef += acceptedAddCoefP
						case Shunned:
							acceptedAddCoefP, err := pm.RandFloat64InRange(0.01, 0.03)
							if err != nil {
								return "", err
							}
							acceptedAddCoef += acceptedAddCoefP
							shunnedAddCoefP, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return "", err
							}
							shunnedAddCoef += shunnedAddCoefP
						case Criminal:
							shunnedAddCoefP, err := pm.RandFloat64InRange(0.01, 0.03)
							if err != nil {
								return "", err
							}
							shunnedAddCoef += shunnedAddCoefP
							criminalAddCoefP, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return "", err
							}
							criminalAddCoef += criminalAddCoefP
						}
					}
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: UseNicotineTabooName,
			_religionMetadata: &religionMetadata{
				Chthonic: 0.1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: UseCannabisTabooName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Pacifistic:   0.1,
				Hedonistic:   1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: UseHallucinogensTabooName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Educational:  0.1,
				Pacifistic:   0.1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: UseCMSStimulantsTabooName,
			_religionMetadata: &religionMetadata{
				Chthonic:   0.1,
				Aggressive: 0.25,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: UseOpiumTabooName,
			_religionMetadata: &religionMetadata{
				Chthonic:   0.25,
				Hedonistic: 0.75,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SameSexRelationsTabooName,
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Chthonic:     0.1,
				Hedonistic:   0.5,
				Liberal:      0.5,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.HighBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: MaleAdulteryTabooName,
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Hedonistic:   0.5,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.HighBaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				switch {
				case r.GenderDominance.IsMaleDominate():
					acceptedAddCoefP, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return "", err
					}
					acceptedAddCoef += acceptedAddCoefP
					shunnedAddCoefP, err := pm.RandFloat64InRange(0.1, 0.2)
					if err != nil {
						return "", err
					}
					shunnedAddCoef += shunnedAddCoefP
				case r.GenderDominance.IsFemaleDominate():
					shunnedAddCoefP, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return "", err
					}
					shunnedAddCoef += shunnedAddCoefP
					criminalAddCoefP, err := pm.RandFloat64InRange(0.1, 0.2)
					if err != nil {
						return "", err
					}
					criminalAddCoef += criminalAddCoefP
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: FemaleAdulteryTabooName,
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Hedonistic:   0.5,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.HighBaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				switch {
				case r.GenderDominance.IsMaleDominate():
					shunnedAddCoefP, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return "", err
					}
					shunnedAddCoef += shunnedAddCoefP
					criminalAddCoefP, err := pm.RandFloat64InRange(0.1, 0.2)
					if err != nil {
						return "", err
					}
					criminalAddCoef += criminalAddCoefP
				case r.GenderDominance.IsFemaleDominate():
					acceptedAddCoefP, err := pm.RandFloat64InRange(0.05, 0.15)
					if err != nil {
						return "", err
					}
					acceptedAddCoef += acceptedAddCoefP
					shunnedAddCoefP, err := pm.RandFloat64InRange(0.1, 0.2)
					if err != nil {
						return "", err
					}
					shunnedAddCoef += shunnedAddCoefP
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SexualDeviancyTabooName,
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Chthonic:     0.25,
				Hedonistic:   0.25,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ProstitutionTabooName,
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Plutocratic:  0.5,
				Hedonistic:   0.25,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: KillAnyoneTabooName,
			_religionMetadata: &religionMetadata{
				Chthonic:   0.1,
				Aggressive: 1,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				var shunnedAddCoef, criminalAddCoef float64
				if len(ts.theology.Traits) > 0 {
					for _, trait := range ts.theology.Traits {
						if trait == nil {
							continue
						}
						if trait.Name == NoMoreKillingTheologyTrait {
							shunnedAddCoefP, err := pm.RandFloat64InRange(0.05, 0.2)
							if err != nil {
								return "", err
							}
							shunnedAddCoef += shunnedAddCoefP
							criminalAddCoefP, err := pm.RandFloat64InRange(0.1, 0.25)
							if err != nil {
								return "", err
							}
							criminalAddCoef += criminalAddCoefP
						}
					}
				}

				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: KillAnimalsTabooName,
			_religionMetadata: &religionMetadata{
				Chthonic:   0.1,
				Aggressive: 0.1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) (Acceptance, error) {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				for _, taboo := range selectedTaboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == KillAnyoneTabooName {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoefP, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return "", err
							}
							shunnedAddCoef += shunnedAddCoefP
							criminalAddCoefP, err := pm.RandFloat64InRange(0.01, 0.03)
							if err != nil {
								return "", err
							}
							criminalAddCoef += criminalAddCoefP
						case Criminal:
							return Criminal, nil
						}
					}
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: KillHollyAnimalsTabooName,
			_religionMetadata: &religionMetadata{
				Chthonic:   0.1,
				Aggressive: 0.1,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.HighBaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) (Acceptance, error) {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				for _, taboo := range selectedTaboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == KillAnyoneTabooName {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoefP, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return "", err
							}
							shunnedAddCoef += shunnedAddCoefP
							criminalAddCoefP, err := pm.RandFloat64InRange(0.01, 0.03)
							if err != nil {
								return "", err
							}
							criminalAddCoef += criminalAddCoefP
						case Criminal:
							return Criminal, nil
						}
					}
					if taboo.Name == KillAnimalsTabooName {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoefP, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return "", err
							}
							shunnedAddCoef += shunnedAddCoefP
							criminalAddCoefP, err := pm.RandFloat64InRange(0.01, 0.03)
							if err != nil {
								return "", err
							}
							criminalAddCoef += criminalAddCoefP
						case Criminal:
							return Criminal, nil
						}
					}
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: KillHumansTabooName,
			_religionMetadata: &religionMetadata{
				Chthonic:   0.25,
				Aggressive: 0.75,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) (Acceptance, error) {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				for _, taboo := range selectedTaboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == KillAnyoneTabooName {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoefP, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return "", err
							}
							shunnedAddCoef += shunnedAddCoefP
							criminalAddCoefP, err := pm.RandFloat64InRange(0.01, 0.03)
							if err != nil {
								return "", err
							}
							criminalAddCoef += criminalAddCoefP
						case Criminal:
							return Criminal, nil
						}
					}
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: KinslayingTabooName,
			_religionMetadata: &religionMetadata{
				Chthonic:   0.75,
				Aggressive: 1,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.HighBaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) (Acceptance, error) {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				for _, taboo := range selectedTaboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == KillAnyoneTabooName {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoefP, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return "", err
							}
							shunnedAddCoef += shunnedAddCoefP
							criminalAddCoefP, err := pm.RandFloat64InRange(0.01, 0.03)
							if err != nil {
								return "", err
							}
							criminalAddCoef += criminalAddCoefP
						case Criminal:
							return Criminal, nil
						}
					}
					if taboo.Name == KillHumansTabooName {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoefP, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return "", err
							}
							shunnedAddCoef += shunnedAddCoefP
							criminalAddCoefP, err := pm.RandFloat64InRange(0.01, 0.03)
							if err != nil {
								return "", err
							}
							criminalAddCoef += criminalAddCoefP
						case Criminal:
							return Criminal, nil
						}
					}
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SuicideTabooName,
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) (Acceptance, error) {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				for _, taboo := range selectedTaboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == KillAnyoneTabooName {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoefP, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return "", err
							}
							shunnedAddCoef += shunnedAddCoefP
							criminalAddCoefP, err := pm.RandFloat64InRange(0.01, 0.03)
							if err != nil {
								return "", err
							}
							criminalAddCoef += criminalAddCoefP
						case Criminal:
							return Criminal, nil
						}
					}
					if taboo.Name == KillHumansTabooName {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoefP, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return "", err
							}
							shunnedAddCoef += shunnedAddCoefP
							criminalAddCoefP, err := pm.RandFloat64InRange(0.01, 0.03)
							if err != nil {
								return "", err
							}
							criminalAddCoef += criminalAddCoefP
						case Criminal:
							return Criminal, nil
						}
					}
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: TattoosTabooName,
			_religionMetadata: &religionMetadata{
				Chthonic: 0.1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: WitchcraftTabooName,
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.HighBaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				if c != nil && len(c.Traditions) > 0 {
					for _, t := range c.Traditions {
						if t == nil {
							continue
						}
						if t.Name == culture.SorcerousMetallurgyTradition.Name {
							return Accepted, nil
						}
					}
				}

				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: NudismTabooName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				SexualActive: 0.1,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.HighBaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

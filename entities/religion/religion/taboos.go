package religion

import (
	"fmt"

	pm "persons_generator/probability_machine"
)

type Taboos struct {
	religion *Religion
	theology *Theology

	Taboos []*Taboo
}

func (t *Theology) generateTaboos() *Taboos {
	ts := &Taboos{religion: t.religion, theology: t}
	ts.Taboos = ts.generateTaboos()

	return ts
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
	Calc       func(r *Religion, self *Taboo, selectedTaboos []*Taboo) Acceptance
}

func (ts *Taboos) generateTaboos() []*Taboo {
	allTaboos := ts.getAllTaboos()
	taboos := make([]*Taboo, len(allTaboos))
	for i, taboo := range allTaboos {
		taboos[i] = &Taboo{
			_religionMetadata: taboo._religionMetadata,
			acceptedBaseCoef:  taboo.acceptedBaseCoef,
			shunnedBaseCoef:   taboo.shunnedBaseCoef,
			criminalBaseCoef:  taboo.criminalBaseCoef,
			Name:              taboo.Name,
			Acceptance:        taboo.Calc(ts.religion, taboo, taboos),
			Calc:              taboo.Calc,
		}
		ts.religion.UpdateMetadata(UpdateReligionMetadataWithAcceptance(ts.religion, *ts.religion.metadata, *taboos[i]._religionMetadata, taboos[i].Acceptance))
	}

	return taboos
}

func (ts *Taboos) getAllTaboos() []*Taboo {
	return []*Taboo{
		{
			Name: "RaisingAnimals",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Plutocratic:  0.1,
			},
			acceptedBaseCoef: ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:  ts.religion.M.LowBaseCoef,
			criminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{
					Log:   true,
					Label: "Taboos.RaisingAnimals",
				})
			},
		},
		{
			Name: "RaisingPlants",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Altruistic:   0.5,
				Plutocratic:  0.25,
			},
			acceptedBaseCoef: ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:  ts.religion.M.LowBaseCoef,
			criminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{
					Log:   true,
					Label: "Taboos.RaisingPlants",
				})
			},
		},
		{
			Name: "RaisingFungus",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Chthonic:     0.1,
				Plutocratic:  0.1,
			},
			acceptedBaseCoef: ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{
					Log:   true,
					Label: "Taboos.RaisingFungus",
				})
			},
		},
		{
			Name: "EatAnimals",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Chthonic:     0.1,
				Hedonistic:   0.1,
			},
			acceptedBaseCoef: ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "EatSomeAnimals",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Lawful:       0.25,
			},
			acceptedBaseCoef: ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) Acceptance {
				var shunnedAddCoef, criminalAddCoef float64
				for _, taboo := range selectedTaboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == "EatAnimals" {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoef += pm.RandFloat64InRange(0.01, 0.05)
							criminalAddCoef += pm.RandFloat64InRange(0.03, 0.1)
						case Criminal:
							return Criminal
						}
					}
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "EatInsects",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Chthonic:     0.1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "EatFungus",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
			},
			acceptedBaseCoef: ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Cannibalism",
			_religionMetadata: &religionMetadata{
				Chthonic:   1,
				Aggressive: 0.25,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.HighBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "DrinkingBlood",
			_religionMetadata: &religionMetadata{
				Chthonic:   1,
				Aggressive: 0.25,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "DrinkingStrongAlcohol",
			_religionMetadata: &religionMetadata{
				SexualActive: 0.1,
				Chthonic:     0.25,
				Aggressive:   0.1,
				Hedonistic:   1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "DrinkingNotStrongAlcohol",
			_religionMetadata: &religionMetadata{
				SexualActive: 0.25,
				Chthonic:     0.1,
				Hedonistic:   1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) Acceptance {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				for _, taboo := range selectedTaboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == "DrinkingStrongAlcohol" {
						switch taboo.Acceptance {
						case Accepted:
							acceptedAddCoef += pm.RandFloat64InRange(0.01, 0.05)
						case Shunned:
							acceptedAddCoef += pm.RandFloat64InRange(0.01, 0.03)
							shunnedAddCoef += pm.RandFloat64InRange(0.01, 0.05)
						case Criminal:
							shunnedAddCoef += pm.RandFloat64InRange(0.01, 0.03)
							criminalAddCoef += pm.RandFloat64InRange(0.01, 0.05)
						}
					}
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "UseNicotine",
			_religionMetadata: &religionMetadata{
				Chthonic: 0.1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "UseCannabis",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Pacifistic:   0.1,
				Hedonistic:   1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "UseHallucinogens",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Educational:  0.1,
				Pacifistic:   0.1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "UseCNSStimulants",
			_religionMetadata: &religionMetadata{
				Chthonic:   0.1,
				Aggressive: 0.25,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "UseOpium",
			_religionMetadata: &religionMetadata{
				Chthonic:   0.25,
				Hedonistic: 0.75,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SameSexRelations",
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Chthonic:     0.1,
				Hedonistic:   0.5,
				Liberal:      0.5,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.HighBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "MaleAdultery",
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Hedonistic:   0.5,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.HighBaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				switch {
				case r.GenderDominance.IsMaleDominate():
					acceptedAddCoef += pm.RandFloat64InRange(0.05, 0.15)
					shunnedAddCoef += pm.RandFloat64InRange(0.1, 0.2)
				case r.GenderDominance.IsFemaleDominate():
					shunnedAddCoef += pm.RandFloat64InRange(0.05, 0.15)
					criminalAddCoef += pm.RandFloat64InRange(0.1, 0.2)
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "FemaleAdultery",
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Hedonistic:   0.5,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.HighBaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				switch {
				case r.GenderDominance.IsMaleDominate():
					shunnedAddCoef += pm.RandFloat64InRange(0.05, 0.15)
					criminalAddCoef += pm.RandFloat64InRange(0.1, 0.2)
				case r.GenderDominance.IsFemaleDominate():
					acceptedAddCoef += pm.RandFloat64InRange(0.05, 0.15)
					shunnedAddCoef += pm.RandFloat64InRange(0.1, 0.2)
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SexualDeviancy",
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Chthonic:     0.25,
				Hedonistic:   0.25,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Prostitution",
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Plutocratic:  0.5,
				Hedonistic:   0.25,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "KillAnyone",
			_religionMetadata: &religionMetadata{
				Chthonic:   0.1,
				Aggressive: 1,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				var shunnedAddCoef, criminalAddCoef float64
				if len(ts.theology.Traits) > 0 {
					for _, trait := range ts.theology.Traits {
						if trait == nil {
							continue
						}
						if trait.Name == "NoMoreKilling" {
							shunnedAddCoef += pm.RandFloat64InRange(0.05, 0.2)
							criminalAddCoef += pm.RandFloat64InRange(0.1, 0.25)
						}
					}
				}

				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "KillAnimals",
			_religionMetadata: &religionMetadata{
				Chthonic:   0.1,
				Aggressive: 0.1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) Acceptance {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				for _, taboo := range selectedTaboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == "KillAnyone" {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoef += pm.RandFloat64InRange(0.01, 0.05)
							criminalAddCoef += pm.RandFloat64InRange(0.01, 0.03)
						case Criminal:
							return Criminal
						}
					}
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "KillHollyAnimals",
			_religionMetadata: &religionMetadata{
				Chthonic:   0.1,
				Aggressive: 0.1,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.HighBaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) Acceptance {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				for _, taboo := range selectedTaboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == "KillAnyone" {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoef += pm.RandFloat64InRange(0.01, 0.05)
							criminalAddCoef += pm.RandFloat64InRange(0.01, 0.03)
						case Criminal:
							return Criminal
						}
					}
					if taboo.Name == "KillAnimals" {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoef += pm.RandFloat64InRange(0.01, 0.05)
							criminalAddCoef += pm.RandFloat64InRange(0.01, 0.03)
						case Criminal:
							return Criminal
						}
					}
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "KillHumans",
			_religionMetadata: &religionMetadata{
				Chthonic:   0.25,
				Aggressive: 0.75,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) Acceptance {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				for _, taboo := range selectedTaboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == "KillAnyone" {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoef += pm.RandFloat64InRange(0.01, 0.05)
							criminalAddCoef += pm.RandFloat64InRange(0.01, 0.03)
						case Criminal:
							return Criminal
						}
					}
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Kinslaying",
			_religionMetadata: &religionMetadata{
				Chthonic:   0.75,
				Aggressive: 1,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.HighBaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) Acceptance {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				for _, taboo := range selectedTaboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == "KillAnyone" {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoef += pm.RandFloat64InRange(0.01, 0.05)
							criminalAddCoef += pm.RandFloat64InRange(0.01, 0.03)
						case Criminal:
							return Criminal
						}
					}
					if taboo.Name == "KillHumans" {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoef += pm.RandFloat64InRange(0.01, 0.05)
							criminalAddCoef += pm.RandFloat64InRange(0.01, 0.03)
						case Criminal:
							return Criminal
						}
					}
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Suicide",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) Acceptance {
				var acceptedAddCoef, shunnedAddCoef, criminalAddCoef float64
				for _, taboo := range selectedTaboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == "KillAnyone" {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoef += pm.RandFloat64InRange(0.01, 0.05)
							criminalAddCoef += pm.RandFloat64InRange(0.01, 0.03)
						case Criminal:
							return Criminal
						}
					}
					if taboo.Name == "KillHumans" {
						switch taboo.Acceptance {
						case Shunned:
							shunnedAddCoef += pm.RandFloat64InRange(0.01, 0.05)
							criminalAddCoef += pm.RandFloat64InRange(0.01, 0.03)
						case Criminal:
							return Criminal
						}
					}
				}
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef+acceptedAddCoef, self.shunnedBaseCoef+shunnedAddCoef, self.criminalBaseCoef+criminalAddCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Tattoos",
			_religionMetadata: &religionMetadata{
				Chthonic: 0.1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.BaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Witchcraft",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			acceptedBaseCoef: ts.religion.M.BaseCoef,
			shunnedBaseCoef:  ts.religion.M.HighBaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Nudism",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				SexualActive: 0.1,
			},
			acceptedBaseCoef: ts.religion.M.LowBaseCoef,
			shunnedBaseCoef:  ts.religion.M.HighBaseCoef,
			criminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

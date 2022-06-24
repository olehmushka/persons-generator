package religion

import "fmt"

type Taboos struct {
	religion *Religion

	Taboos []*Taboo
}

func (t *Theology) generateTaboos() *Taboos {
	ts := &Taboos{religion: t.religion}
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
		ts.religion.UpdateMetadata(UpdateReligionMetadata(*ts.religion.metadata, *taboos[i]._religionMetadata))
	}

	return taboos
}

func (ts *Taboos) getAllTaboos() []*Taboo {
	return []*Taboo{
		{
			Name:              "RaisingAnimals",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:   ts.religion.M.LowBaseCoef,
			criminalBaseCoef:  ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "RaisingPlants",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:   ts.religion.M.LowBaseCoef,
			criminalBaseCoef:  ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "RaisingFungus",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.HighBaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "EatAnimals",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "EatSomeAnimals",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, selectedTaboos []*Taboo) Acceptance {
				// for _, taboo := range selectedTaboos {
				// 	if taboo.
				// }
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "EatInsects",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "EatFungus",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Cannibalism",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "DrinkingBlood",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "DrinkingStrongAlcohol",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "DrinkingNotStrongAlcohol",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "UseNicotine",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "UseCannabis",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "UseHallucinogens",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "UseCNSStimulants",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "UseOpium",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "SameSexRelations",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "MaleAdultery",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "FemaleAdultery",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "SexualDeviancy",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "KillAnyone",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "KillAnimals",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "KillHollyAnimals",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "KillHumans",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Kinslaying",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Suicide",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Tattoos",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Witchcraft",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Nudism",
			_religionMetadata: &religionMetadata{},
			acceptedBaseCoef:  ts.religion.M.BaseCoef,
			shunnedBaseCoef:   ts.religion.M.BaseCoef,
			criminalBaseCoef:  ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) Acceptance {
				return CalculateAcceptanceFromReligionMetadata(self.acceptedBaseCoef, self.shunnedBaseCoef, self.criminalBaseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

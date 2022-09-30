package religion

import (
	"fmt"
	"math"

	"persons_generator/engine/entities/culture"
	pm "persons_generator/engine/probability_machine"
)

type Taboos struct {
	religion *Religion `json:"-" bson:"-"`
	theology *Theology `json:"-" bson:"-"`

	Taboos []*Taboo `json:"taboos" bson:"taboos"`
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

func (ts *Taboos) IsZero() bool {
	return ts == nil
}

func (ts *Taboos) Print() {
	fmt.Printf("Taboos (religion_name=%s):\n", ts.religion.Name)
	for _, taboo := range ts.Taboos {
		fmt.Printf(" - %s (%s)\n", taboo.Name, taboo.Acceptance)
	}
}

type Taboo struct {
	ReligionMetadata *religionMetadata `json:"religion_metadata" bson:"religion_metadata"`
	AcceptedBaseCoef float64           `json:"accepted_base_coef" bson:"accepted_base_coef"`
	ShunnedBaseCoef  float64           `json:"shunned_base_coef" bson:"shunned_base_coef"`
	CriminalBaseCoef float64           `json:"criminal_base_coef" bson:"criminal_base_coef"`

	Name       string                                                                      `json:"name" bson:"name"`
	Acceptance Acceptance                                                                  `json:"acceptance" bson:"acceptance"`
	Calc       func(r *Religion, self *Taboo, selectedTaboos []*Taboo) (Acceptance, error) `json:"-" bson:"-"`
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
			ReligionMetadata: taboo.ReligionMetadata,
			AcceptedBaseCoef: taboo.AcceptedBaseCoef,
			ShunnedBaseCoef:  taboo.ShunnedBaseCoef,
			CriminalBaseCoef: taboo.CriminalBaseCoef,
			Name:             taboo.Name,
			Acceptance:       a,
			Calc:             taboo.Calc,
		}

		md, err := UpdateReligionMetadataWithAcceptance(ts.religion, *ts.religion.Metadata, *taboos[i].ReligionMetadata, taboos[i].Acceptance)
		if err != nil {
			return nil, err
		}
		ts.religion.UpdateMetadata(md)
	}

	return taboos, nil
}

func (t *Taboo) IsZero() bool {
	return t == nil
}

func (ts *Taboos) getAllTaboos(c *culture.Culture) []*Taboo {
	return []*Taboo{
		{
			Name: RaisingAnimalsTabooName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Plutocratic:  0.1,
			},
			AcceptedBaseCoef: ts.religion.M.HighBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.LowBaseCoef,
			CriminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: RaisingPlantsTabooName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Altruistic:   0.5,
				Plutocratic:  0.25,
			},
			AcceptedBaseCoef: ts.religion.M.HighBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.LowBaseCoef,
			CriminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: RaisingFungusTabooName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Chthonic:     0.1,
				Plutocratic:  0.1,
			},
			AcceptedBaseCoef: ts.religion.M.HighBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: EatAnimalsTabooName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Chthonic:     0.1,
				Hedonistic:   0.1,
			},
			AcceptedBaseCoef: ts.religion.M.HighBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: EatSomeAnimalsTabooName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Lawful:       0.25,
			},
			AcceptedBaseCoef: ts.religion.M.HighBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
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
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef+shunnedAddCoef, self.CriminalBaseCoef+criminalAddCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: EatInsectsTabooName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Chthonic:     0.1,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: EatFungusTabooName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.1,
			},
			AcceptedBaseCoef: ts.religion.M.HighBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: CannibalismTabooName,
			ReligionMetadata: &religionMetadata{
				Chthonic:   1,
				Aggressive: 0.25,
			},
			AcceptedBaseCoef: ts.religion.M.LowBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.HighBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DrinkingBloodTabooName,
			ReligionMetadata: &religionMetadata{
				Chthonic:   1,
				Aggressive: 0.25,
			},
			AcceptedBaseCoef: ts.religion.M.LowBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DrinkingStrongAlcoholTabooName,
			ReligionMetadata: &religionMetadata{
				SexualActive: 0.1,
				Chthonic:     0.25,
				Aggressive:   0.1,
				Hedonistic:   1,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DrinkingNotStrongAlcoholTabooName,
			ReligionMetadata: &religionMetadata{
				SexualActive: 0.25,
				Chthonic:     0.1,
				Hedonistic:   1,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.LowBaseCoef,
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
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef+acceptedAddCoef, self.ShunnedBaseCoef+shunnedAddCoef, self.CriminalBaseCoef+criminalAddCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: UseNicotineTabooName,
			ReligionMetadata: &religionMetadata{
				Chthonic: 0.1,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: UseCannabisTabooName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Pacifistic:   0.1,
				Hedonistic:   1,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: UseHallucinogensTabooName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Educational:  0.1,
				Pacifistic:   0.1,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: UseCMSStimulantsTabooName,
			ReligionMetadata: &religionMetadata{
				Chthonic:   0.1,
				Aggressive: 0.25,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: UseOpiumTabooName,
			ReligionMetadata: &religionMetadata{
				Chthonic:   0.25,
				Hedonistic: 0.75,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SameSexRelationsTabooName,
			ReligionMetadata: &religionMetadata{
				SexualActive: 1,
				Chthonic:     0.1,
				Hedonistic:   0.5,
				Liberal:      0.5,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.HighBaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: MaleAdulteryTabooName,
			ReligionMetadata: &religionMetadata{
				SexualActive: 1,
				Hedonistic:   0.5,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.HighBaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
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
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef+acceptedAddCoef, self.ShunnedBaseCoef+shunnedAddCoef, self.CriminalBaseCoef+criminalAddCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: FemaleAdulteryTabooName,
			ReligionMetadata: &religionMetadata{
				SexualActive: 1,
				Hedonistic:   0.5,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.HighBaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
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
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef+acceptedAddCoef, self.ShunnedBaseCoef+shunnedAddCoef, self.CriminalBaseCoef+criminalAddCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SexualDeviancyTabooName,
			ReligionMetadata: &religionMetadata{
				SexualActive: 1,
				Chthonic:     0.25,
				Hedonistic:   0.25,
			},
			AcceptedBaseCoef: ts.religion.M.LowBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ProstitutionTabooName,
			ReligionMetadata: &religionMetadata{
				SexualActive: 1,
				Plutocratic:  0.5,
				Hedonistic:   0.25,
			},
			AcceptedBaseCoef: ts.religion.M.LowBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: KillAnyoneTabooName,
			ReligionMetadata: &religionMetadata{
				Chthonic:   0.1,
				Aggressive: 1,
			},
			AcceptedBaseCoef: ts.religion.M.LowBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
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

				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: KillAnimalsTabooName,
			ReligionMetadata: &religionMetadata{
				Chthonic:   0.1,
				Aggressive: 0.1,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
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
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef+acceptedAddCoef, self.ShunnedBaseCoef+shunnedAddCoef, self.CriminalBaseCoef+criminalAddCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: KillHollyAnimalsTabooName,
			ReligionMetadata: &religionMetadata{
				Chthonic:   0.1,
				Aggressive: 0.1,
			},
			AcceptedBaseCoef: ts.religion.M.LowBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.HighBaseCoef,
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
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef+acceptedAddCoef, self.ShunnedBaseCoef+shunnedAddCoef, self.CriminalBaseCoef+criminalAddCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: KillHumansTabooName,
			ReligionMetadata: &religionMetadata{
				Chthonic:   0.25,
				Aggressive: 0.75,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
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
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef+acceptedAddCoef, self.ShunnedBaseCoef+shunnedAddCoef, self.CriminalBaseCoef+criminalAddCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: KinslayingTabooName,
			ReligionMetadata: &religionMetadata{
				Chthonic:   0.75,
				Aggressive: 1,
			},
			AcceptedBaseCoef: ts.religion.M.LowBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.HighBaseCoef,
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
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef+acceptedAddCoef, self.ShunnedBaseCoef+shunnedAddCoef, self.CriminalBaseCoef+criminalAddCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SuicideTabooName,
			ReligionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			AcceptedBaseCoef: ts.religion.M.LowBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
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
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef+acceptedAddCoef, self.ShunnedBaseCoef+shunnedAddCoef, self.CriminalBaseCoef+criminalAddCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: TattoosTabooName,
			ReligionMetadata: &religionMetadata{
				Chthonic: 0.1,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.BaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: WitchcraftTabooName,
			ReligionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			AcceptedBaseCoef: ts.religion.M.BaseCoef,
			ShunnedBaseCoef:  ts.religion.M.HighBaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
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

				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: NudismTabooName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				SexualActive: 0.1,
			},
			AcceptedBaseCoef: ts.religion.M.LowBaseCoef,
			ShunnedBaseCoef:  ts.religion.M.HighBaseCoef,
			CriminalBaseCoef: ts.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Taboo, _ []*Taboo) (Acceptance, error) {
				return CalculateAcceptanceFromReligionMetadata(self.AcceptedBaseCoef, self.ShunnedBaseCoef, self.CriminalBaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func GetTabooSimilarityCoef(t1, t2 *Taboo) float64 {
	if t1.IsZero() && t2.IsZero() {
		return 1
	}
	if t1.IsZero() || t2.IsZero() {
		return 0
	}

	return 1 - math.Abs(t1.Acceptance.Value()-t2.Acceptance.Value())
}

func GetTaboosSimilarityCoef(t1, t2 *Taboos) float64 {
	if t1.IsZero() && t2.IsZero() {
		return 1
	}
	if t1.IsZero() || t2.IsZero() {
		return 0
	}
	if len(t1.Taboos) != len(t2.Taboos) {
		return 0
	}
	var (
		out      float64
		unitCoef = 1 / float64(len(t1.Taboos))
	)
	for i, coef := range t1.Taboos {
		out += GetTabooSimilarityCoef(coef, t2.Taboos[i]) * unitCoef
	}

	return out
}

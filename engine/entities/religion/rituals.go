package religion

import (
	"fmt"

	pm "persons_generator/engine/probability_machine"
)

type Rituals struct {
	religion *Religion `json:"-"`
	theology *Theology `json:"-"`

	Initiation []*trait `json:"initiations"`
	Funeral    []*trait `json:"funeral"`
	Sacrifice  []*trait `json:"sacrifice"`
	Holyday    []*trait `json:"holyday"`
}

func (t *Theology) generateRituals() (*Rituals, error) {
	rs := &Rituals{religion: t.religion, theology: t}
	initiation, err := generateTraits(rs.religion, rs.getAllInitiationRituals(), generateTraitsOpts{
		Label: "Rituals.generateInitiationRituals",
		Min:   1,
		Max:   3,
	})
	if err != nil {
		return nil, err
	}
	rs.Initiation = initiation
	funeral, err := generateTraits(rs.religion, rs.getAllFuneralRituals(), generateTraitsOpts{
		Label: "Rituals.generateFuneralRituals",
		Min:   1,
		Max:   3,
	})
	if err != nil {
		return nil, err
	}
	rs.Funeral = funeral
	sacrifice, err := generateTraits(rs.religion, rs.getAllSacrificeRituals(), generateTraitsOpts{
		Label: "Rituals.generateSacrificeRituals",
		Min:   0,
		Max:   3,
	})
	if err != nil {
		return nil, err
	}
	rs.Sacrifice = sacrifice
	holyday, err := generateTraits(rs.religion, rs.getAllHolydayRituals(), generateTraitsOpts{
		Label: "Rituals.generateHolydayRituals",
		Min:   0,
		Max:   3,
	})
	if err != nil {
		return nil, err
	}
	rs.Holyday = holyday

	return rs, nil
}

func (rs *Rituals) IsZero() bool {
	return rs == nil
}

func (rs *Rituals) Print() {
	fmt.Printf("Rituals (religion_name=%s):\n", rs.religion.Name)
	if len(rs.Initiation) > 0 {
		for _, ritual := range rs.Initiation {
			fmt.Printf(" - %s\n", ritual.Name)
		}
	}
	if len(rs.Funeral) > 0 {
		for _, ritual := range rs.Funeral {
			fmt.Printf(" - %s\n", ritual.Name)
		}
	}
	if len(rs.Sacrifice) > 0 {
		for _, ritual := range rs.Sacrifice {
			fmt.Printf(" - %s\n", ritual.Name)
		}
	}
	if len(rs.Holyday) > 0 {
		for _, ritual := range rs.Holyday {
			fmt.Printf(" - %s\n", ritual.Name)
		}
	}
}

/*

ordeals,
nithing,
prayer spells

PersonalRituals:
	- OathOfHealer
	- BloodOath
	- Pilgrims


SituativeRituals:
	- Confession
	- BuyingYourSins
	- Pharmakos
	- Inferiae
	- Scapegoat
	- CleansingRituals
	- Adorcism
	- Exorcism
	- Blinding
	- DemonicDeal
*/

func (rs *Rituals) getAllInitiationRituals() []*trait { // @TODO: finish it
	return []*trait{
		{
			Name: MarriageWithDeityRitual,
			ReligionMetadata: &religionMetadata{
				SexualActive:    0.5,
				Lawful:          0.1,
				Individualistic: 0.5,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BaptismRitual,
			ReligionMetadata: &religionMetadata{
				Individualistic: 0.5,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: InitialTattooRitual,
			ReligionMetadata: &religionMetadata{
				Chthonic:        0.1,
				Individualistic: 0.5,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.Theology != nil && r.Theology.Taboos != nil && len(r.Theology.Taboos.Taboos) > 0 {
					for _, taboo := range r.Theology.Taboos.Taboos {
						if taboo == nil {
							continue
						}
						if taboo.Name == TattoosTabooName {
							switch taboo.Acceptance {
							case Accepted:
								coef, err := pm.RandFloat64InRange(0.01, 0.1)
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
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: OathOfLoyaltyRitual,
			ReligionMetadata: &religionMetadata{
				Lawful:          0.75,
				Individualistic: 0.5,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: InitiationByFastRitual,
			ReligionMetadata: &religionMetadata{
				Ascetic:         1,
				Individualistic: 0.25,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: InitiationByFireRitual,
			ReligionMetadata: &religionMetadata{
				Chthonic:        0.25,
				Aggressive:      0.1,
				Individualistic: 0.5,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (rs *Rituals) getAllFuneralRituals() []*trait {
	return []*trait{
		{
			Name: SkyBurialsRitual,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Lawful:       0.25,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: CaveBurialsRitual,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Lawful:       0.25,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: UndergroundBurialsRitual,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Lawful:       0.25,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (rs *Rituals) getAllSacrificeRituals() []*trait {
	return []*trait{
		{
			Name: RitualSuicideRitual,
			ReligionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				for _, taboo := range rs.theology.Taboos.Taboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == SuicideTabooName {
						if taboo.Acceptance.IsCriminal() {
							return false, nil
						}
						if taboo.Acceptance.IsShunned() {
							coef, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return false, err
							}
							addCoef -= coef
						}
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: AnimalSacrificeRitual,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Simple:       0.5,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				for _, taboo := range rs.theology.Taboos.Taboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == KillAnimalsTabooName {
						if taboo.Acceptance.IsCriminal() {
							return false, nil
						}
						if taboo.Acceptance.IsShunned() {
							coef, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return false, err
							}
							addCoef -= coef
						}
					}
				}
				if r.Type.IsPolytheism() {
					coef, err := pm.RandFloat64InRange(0.1, 0.2)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HecatombsRitual,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Simple:       0.5,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedRituals []*trait) (bool, error) {
				var addCoef float64
				for _, ritual := range selectedRituals {
					if ritual == nil {
						continue
					}
					if ritual.Name == AnimalSacrificeRitual {
						coef, err := pm.RandFloat64InRange(0.1, 0.2)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}
				if addCoef == 0 {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HumanSacrificeRitual,
			ReligionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				for _, taboo := range rs.theology.Taboos.Taboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == KillHumansTabooName {
						if taboo.Acceptance.IsCriminal() {
							return false, nil
						}
						if taboo.Acceptance.IsShunned() {
							coef, err := pm.RandFloat64InRange(0.05, 0.15)
							if err != nil {
								return false, err
							}
							addCoef -= coef
						}
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ChildSacrificeRitual,
			ReligionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			BaseCoef: rs.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, selectedRituals []*trait) (bool, error) {
				var addCoef float64
				for _, ritual := range selectedRituals {
					if ritual == nil {
						continue
					}
					if ritual.Name == HumanSacrificeRitual {
						coef, err := pm.RandFloat64InRange(0.01, 0.1)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}
				if addCoef == 0 {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: PlantsSacrificeRitual,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (rs *Rituals) getAllHolydayRituals() []*trait {
	return []*trait{
		{
			Name: GreatFastRitual,
			ReligionMetadata: &religionMetadata{
				Ascetic:        1,
				Collectivistic: 0.25,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: RitesAndChantsRitual,
			ReligionMetadata: &religionMetadata{
				Hedonistic: 0.75,
				Simple:     1,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: OrgyRitual,
			ReligionMetadata: &religionMetadata{
				SexualActive: 1,
				Hedonistic:   1,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				for _, taboo := range rs.theology.Taboos.Taboos {
					if taboo.Name == MaleAdulteryTabooName || taboo.Name == FemaleAdulteryTabooName {
						if taboo.Acceptance.IsCriminal() {
							coef, err := pm.RandFloat64InRange(0.1, 0.3)
							if err != nil {
								return false, err
							}
							addCoef -= coef
						}
						if taboo.Acceptance.IsShunned() {
							coef, err := pm.RandFloat64InRange(0.05, 0.1)
							if err != nil {
								return false, err
							}
							addCoef -= coef
						}
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SmokeCircleRitual,
			ReligionMetadata: &religionMetadata{
				Naturalistic:   0.25,
				Collectivistic: 0.5,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				for _, taboo := range rs.theology.Taboos.Taboos {
					if taboo.Name == UseNicotineTabooName || taboo.Name == UseCannabisTabooName || taboo.Name == UseHallucinogensTabooName {
						switch {
						case taboo.Acceptance.IsAccepted():
							coef, err := pm.RandFloat64InRange(0.05, 0.15)
							if err != nil {
								return false, err
							}
							addCoef += coef
						case taboo.Acceptance.IsShunned():
							coef, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return false, err
							}
							addCoef -= coef
						case taboo.Acceptance.IsCriminal():
							coef, err := pm.RandFloat64InRange(0.05, 0.1)
							if err != nil {
								return false, err
							}
							addCoef -= coef
						}
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GladiatorDuelRitual,
			ReligionMetadata: &religionMetadata{
				Aggressive: 1,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				for _, taboo := range rs.theology.Taboos.Taboos {
					if taboo.Name == KillHumansTabooName {
						switch {
						case taboo.Acceptance.IsAccepted():
							coef, err := pm.RandFloat64InRange(0.05, 0.15)
							if err != nil {
								return false, err
							}
							addCoef += coef
						case taboo.Acceptance.IsShunned():
							coef, err := pm.RandFloat64InRange(0.01, 0.05)
							if err != nil {
								return false, err
							}
							addCoef -= coef
						case taboo.Acceptance.IsCriminal():
							return false, nil
						}
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: RitualCannibalismRitual,
			ReligionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			BaseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				for _, taboo := range rs.theology.Taboos.Taboos {
					if taboo.Name == CannibalismTabooName {
						switch {
						case taboo.Acceptance.IsAccepted():
							coef, err := pm.RandFloat64InRange(0.05, 0.15)
							if err != nil {
								return false, err
							}
							addCoef += coef
						case taboo.Acceptance.IsShunned():
							coef, err := pm.RandFloat64InRange(0.03, 0.075)
							if err != nil {
								return false, err
							}
							addCoef -= coef
						case taboo.Acceptance.IsCriminal():
							return false, nil
						}
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func GetRitualsSimilarityCoef(r1, r2 *Rituals) float64 {
	if r1.IsZero() && r2.IsZero() {
		return 1
	}
	if r1.IsZero() || r2.IsZero() {
		return 0
	}

	similarityTraits := []struct {
		value float64
		coef  float64
	}{
		{
			value: GetTraitsSimilarityCoef(r1.Initiation, r2.Initiation),
			coef:  0.25,
		},
		{
			value: GetTraitsSimilarityCoef(r1.Funeral, r2.Funeral),
			coef:  0.25,
		},
		{
			value: GetTraitsSimilarityCoef(r1.Sacrifice, r2.Sacrifice),
			coef:  0.25,
		},
		{
			value: GetTraitsSimilarityCoef(r1.Holyday, r2.Holyday),
			coef:  0.25,
		},
	}

	var out float64
	for _, t := range similarityTraits {
		out += t.value * t.coef
	}

	return out
}

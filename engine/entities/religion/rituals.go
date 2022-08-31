package religion

import (
	"fmt"

	pm "persons_generator/engine/probability_machine"
)

type Rituals struct {
	religion *Religion
	theology *Theology

	Initiation []*trait
	Funeral    []*trait
	Sacrifice  []*trait
	Holyday    []*trait
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
			_religionMetadata: &religionMetadata{
				SexualActive:    0.5,
				Lawful:          0.1,
				Individualistic: 0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BaptismRitual,
			_religionMetadata: &religionMetadata{
				Individualistic: 0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: InitialTattooRitual,
			_religionMetadata: &religionMetadata{
				Chthonic:        0.1,
				Individualistic: 0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
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
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: OathOfLoyaltyRitual,
			_religionMetadata: &religionMetadata{
				Lawful:          0.75,
				Individualistic: 0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: InitiationByFastRitual,
			_religionMetadata: &religionMetadata{
				Ascetic:         1,
				Individualistic: 0.25,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: InitiationByFireRitual,
			_religionMetadata: &religionMetadata{
				Chthonic:        0.25,
				Aggressive:      0.1,
				Individualistic: 0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (rs *Rituals) getAllFuneralRituals() []*trait {
	return []*trait{
		{
			Name: SkyBurialsRitual,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Lawful:       0.25,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: CaveBurialsRitual,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Lawful:       0.25,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: UndergroundBurialsRitual,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Lawful:       0.25,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (rs *Rituals) getAllSacrificeRituals() []*trait {
	return []*trait{
		{
			Name: RitualSuicideRitual,
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: rs.religion.M.BaseCoef,
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
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: AnimalSacrificeRitual,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Simple:       0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
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
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HecatombsRitual,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Simple:       0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
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

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HumanSacrificeRitual,
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: rs.religion.M.BaseCoef,
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
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ChildSacrificeRitual,
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: rs.religion.M.LowBaseCoef,
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

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: PlantsSacrificeRitual,
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (rs *Rituals) getAllHolydayRituals() []*trait {
	return []*trait{
		{
			Name: GreatFastRitual,
			_religionMetadata: &religionMetadata{
				Ascetic:        1,
				Collectivistic: 0.25,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: RitesAndChantsRitual,
			_religionMetadata: &religionMetadata{
				Hedonistic: 0.75,
				Simple:     1,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: OrgyRitual,
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Hedonistic:   1,
			},
			baseCoef: rs.religion.M.BaseCoef,
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
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SmokeCircleRitual,
			_religionMetadata: &religionMetadata{
				Naturalistic:   0.25,
				Collectivistic: 0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
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
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GladiatorDuelRitual,
			_religionMetadata: &religionMetadata{
				Aggressive: 1,
			},
			baseCoef: rs.religion.M.BaseCoef,
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
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: RitualCannibalismRitual,
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: rs.religion.M.BaseCoef,
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
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

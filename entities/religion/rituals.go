package religion

import (
	"fmt"

	pm "persons_generator/probability_machine"
)

type Rituals struct {
	religion *Religion
	theology *Theology

	Initiation []*trait
	Funeral    []*trait
	Sacrifice  []*trait
	Holyday    []*trait
}

func (t *Theology) generateRituals() *Rituals {
	rs := &Rituals{religion: t.religion, theology: t}
	rs.Initiation = generateTraits(rs.religion, rs.getAllInitiationRituals(), generateTraitsOpts{
		Label: "Rituals.generateInitiationRituals",
		Min:   1,
		Max:   3,
	})
	rs.Funeral = generateTraits(rs.religion, rs.getAllFuneralRituals(), generateTraitsOpts{
		Label: "Rituals.generateFuneralRituals",
		Min:   1,
		Max:   3,
	})
	rs.Sacrifice = generateTraits(rs.religion, rs.getAllSacrificeRituals(), generateTraitsOpts{
		Label: "Rituals.generateSacrificeRituals",
		Min:   0,
		Max:   3,
	})
	rs.Holyday = generateTraits(rs.religion, rs.getAllHolydayRituals(), generateTraitsOpts{
		Label: "Rituals.generateHolydayRituals",
		Min:   0,
		Max:   3,
	})

	return rs
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
			Name: "MarriageWithDeity",
			_religionMetadata: &religionMetadata{
				SexualActive:    0.5,
				Lawful:          0.1,
				Individualistic: 0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Baptism",
			_religionMetadata: &religionMetadata{
				Individualistic: 0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "InitialTattoo",
			_religionMetadata: &religionMetadata{
				Chthonic:        0.1,
				Individualistic: 0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				if r.Theology != nil && r.Theology.Taboos != nil && len(r.Theology.Taboos.Taboos) > 0 {
					for _, taboo := range r.Theology.Taboos.Taboos {
						if taboo == nil {
							continue
						}
						if taboo.Name == "Tattoos" {
							switch taboo.Acceptance {
							case Accepted:
								addCoef += pm.RandFloat64InRange(0.01, 0.1)
							case Shunned:
								addCoef += -pm.RandFloat64InRange(0.01, 0.1)
							case Criminal:
								return false
							}
						}
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "OathOfLoyalty",
			_religionMetadata: &religionMetadata{
				Lawful:          0.75,
				Individualistic: 0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "InitiationByFast",
			_religionMetadata: &religionMetadata{
				Ascetic:         1,
				Individualistic: 0.25,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "InitiationByFire",
			_religionMetadata: &religionMetadata{
				Chthonic:        0.25,
				Aggressive:      0.1,
				Individualistic: 0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (rs *Rituals) getAllFuneralRituals() []*trait {
	return []*trait{
		{
			Name: "SkyBurials",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Lawful:       0.25,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "CaveBurials",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Lawful:       0.25,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "UndergroundBurials",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Lawful:       0.25,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (rs *Rituals) getAllSacrificeRituals() []*trait {
	return []*trait{
		{
			Name: "RitualSuicide",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				for _, taboo := range rs.theology.Taboos.Taboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == "Suicide" {
						if taboo.Acceptance.IsCriminal() {
							return false
						}
						if taboo.Acceptance.IsShunned() {
							addCoef += -pm.RandFloat64InRange(0.01, 0.05)
						}
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "AnimalSacrifice",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Simple:       0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				for _, taboo := range rs.theology.Taboos.Taboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == "KillAnimals" {
						if taboo.Acceptance.IsCriminal() {
							return false
						}
						if taboo.Acceptance.IsShunned() {
							addCoef += -pm.RandFloat64InRange(0.01, 0.05)
						}
					}
				}
				if r.Type.IsPolytheism() {
					addCoef += pm.RandFloat64InRange(0.1, 0.2)
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Hecatombs",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Simple:       0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedRituals []*trait) bool {
				var addCoef float64
				for _, ritual := range selectedRituals {
					if ritual == nil {
						continue
					}
					if ritual.Name == "AnimalSacrifice" {
						addCoef += pm.RandFloat64InRange(0.1, 0.2)
					}
				}
				if addCoef == 0 {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HumanSacrifice",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				for _, taboo := range rs.theology.Taboos.Taboos {
					if taboo == nil {
						continue
					}
					if taboo.Name == "KillHumans" {
						if taboo.Acceptance.IsCriminal() {
							return false
						}
						if taboo.Acceptance.IsShunned() {
							addCoef += -pm.RandFloat64InRange(0.05, 0.15)
						}
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "ChildSacrifice",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: rs.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, selectedRituals []*trait) bool {
				var addCoef float64
				for _, ritual := range selectedRituals {
					if ritual == nil {
						continue
					}
					if ritual.Name == "HumanSacrifice" {
						addCoef += pm.RandFloat64InRange(0.01, 0.1)
					}
				}
				if addCoef == 0 {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "PlantsSacrifice",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (rs *Rituals) getAllHolydayRituals() []*trait {
	return []*trait{
		{
			Name: "GreatFast",
			_religionMetadata: &religionMetadata{
				Ascetic:        1,
				Collectivistic: 0.25,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "RitesAndChants",
			_religionMetadata: &religionMetadata{
				Hedonistic: 0.75,
				Simple:     1,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Orgy",
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Hedonistic:   1,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				for _, taboo := range rs.theology.Taboos.Taboos {
					if taboo.Name == "MaleAdultery" || taboo.Name == "FemaleAdultery" {
						if taboo.Acceptance.IsCriminal() {
							addCoef += -pm.RandFloat64InRange(0.1, 0.3)
						}
						if taboo.Acceptance.IsShunned() {
							addCoef += -pm.RandFloat64InRange(0.05, 0.1)
						}
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SmokeCircle",
			_religionMetadata: &religionMetadata{
				Naturalistic:   0.25,
				Collectivistic: 0.5,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				for _, taboo := range rs.theology.Taboos.Taboos {
					if taboo.Name == "UseNicotine" || taboo.Name == "UseCannabis" || taboo.Name == "UseHallucinogens" {
						switch {
						case taboo.Acceptance.IsAccepted():
							addCoef += pm.RandFloat64InRange(0.05, 0.15)
						case taboo.Acceptance.IsShunned():
							addCoef += -pm.RandFloat64InRange(0.01, 0.05)
						case taboo.Acceptance.IsCriminal():
							addCoef += -pm.RandFloat64InRange(0.05, 0.1)
						}
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "GladiatorDuel",
			_religionMetadata: &religionMetadata{
				Aggressive: 1,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				for _, taboo := range rs.theology.Taboos.Taboos {
					if taboo.Name == "KillHumans" {
						switch {
						case taboo.Acceptance.IsAccepted():
							addCoef += pm.RandFloat64InRange(0.05, 0.15)
						case taboo.Acceptance.IsShunned():
							addCoef += -pm.RandFloat64InRange(0.01, 0.05)
						case taboo.Acceptance.IsCriminal():
							return false
						}
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "RitualCannibalism",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: rs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				for _, taboo := range rs.theology.Taboos.Taboos {
					if taboo.Name == "Cannibalism" {
						switch {
						case taboo.Acceptance.IsAccepted():
							addCoef += pm.RandFloat64InRange(0.05, 0.15)
						case taboo.Acceptance.IsShunned():
							addCoef += -pm.RandFloat64InRange(0.03, 0.075)
						case taboo.Acceptance.IsCriminal():
							return false
						}
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

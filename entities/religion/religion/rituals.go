package religion

import "fmt"

type Rituals struct {
	religion *Religion

	Initiation []*Ritual
	Funeral    []*Ritual
	Sacrifice  []*Ritual
}

func (t *Theology) generateRituals() *Rituals {
	rs := &Rituals{religion: t.religion}
	rs.Initiation = rs.generateSomeRituals(rs.getAllInitiationRituals(), 0, len(rs.getAllInitiationRituals()))
	rs.Funeral = rs.generateSomeRituals(rs.getAllFuneralRituals(), 0, len(rs.getAllFuneralRituals()))
	rs.Sacrifice = rs.generateSomeRituals(rs.getAllSacrificeRituals(), 0, len(rs.getAllSacrificeRituals()))

	return rs
}

func (rs *Rituals) Print() {
	fmt.Printf("Rituals (religion_name=%s):\n", rs.religion.Name)
}

/*
ordeals, great fast, confession
sacrificial offering, rites and chants
buying your sins, hecatombs, pharmakos, inferiae
scapegoat
nithing, cleansing rituals
oath of healer, blood oath
prayer spells

Orgy, Adorcism, Exorcism
SmokeCircle, GladiatorDuel
Blinding, RitualCannibalism
*/

type Ritual struct {
	_religionMetadata *updateReligionMetadata
	baseCoef          float64

	Name string
	Calc func(r *Religion, self *Ritual, selectedRituals []*Ritual) bool
}

func (rs *Rituals) generateSomeRituals(ritualsToSelect []*Ritual, min, max int) []*Ritual {
	if min < 0 {
		panic("[Rituals.generateInitiation] min can not be less than 0")
	}
	if max > len(ritualsToSelect) {
		panic("[Rituals.generateInitiation] max can not be greater than ritualsToSelect length")
	}

	rituals := make([]*Ritual, 0, len(ritualsToSelect))
	for count := 0; count < 20; count++ {
		for _, ritual := range ritualsToSelect {
			if ritual.Calc(rs.religion, ritual, rituals) {
				rituals = append(rituals, &Ritual{
					_religionMetadata: ritual._religionMetadata,
					Name:              ritual.Name,
					Calc:              ritual.Calc,
				})
			}
			if len(rituals) == max {
				break
			}
		}
		if len(rituals) == max {
			break
		}
		if len(rituals) >= min {
			break
		}
	}

	for _, trait := range rituals {
		rs.religion.UpdateMetadata(UpdateReligionMetadata(*rs.religion.metadata, *trait._religionMetadata))
	}

	return rituals
}

func (rs *Rituals) getAllInitiationRituals() []*Ritual {
	return []*Ritual{}
}

func (rs *Rituals) getAllFuneralRituals() []*Ritual {
	return []*Ritual{
		{
			Name:              "SkyBurials",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          1,
			Calc: func(r *Religion, self *Ritual, _ []*Ritual) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "CaveBurials",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          1,
			Calc: func(r *Religion, self *Ritual, _ []*Ritual) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "UndergroundBurials",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          1,
			Calc: func(r *Religion, self *Ritual, _ []*Ritual) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (rs *Rituals) getAllSacrificeRituals() []*Ritual {
	return []*Ritual{
		{
			Name:              "RitualSuicide",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          1,
			Calc: func(r *Religion, self *Ritual, _ []*Ritual) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "AnimalSacrifice",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          1,
			Calc: func(r *Religion, self *Ritual, _ []*Ritual) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "HumanSacrifice",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          1,
			Calc: func(r *Religion, self *Ritual, _ []*Ritual) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "PlantsSacrifice",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          1,
			Calc: func(r *Religion, self *Ritual, _ []*Ritual) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

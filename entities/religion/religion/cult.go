package religion

import pm "persons_generator/probability-machine"

func (t *Theology) getCultsRange() (int, int) {
	var min, max int
	switch {
	case t.religion.Type.IsMonotheism():
		max = 3
	case t.religion.Type.IsPolytheism():
		min = 1
		max = len(t.getAllCults())
	case t.religion.Type.IsDeityDualism():
		max = 4
	case t.religion.Type.IsDeism():
		min = 1
		max = len(t.getAllCults())
	case t.religion.Type.IsAtheism():
	default:
	}

	return min, max
}

type Cult struct {
	_religionMetadata *religionMetadata
	baseCoef          float64

	Name string
	Calc func(r *Religion, self *Cult, selectedCults []*Cult) bool
}

func (t *Theology) generateCults(min, max int) []*Cult {
	if min < 0 {
		panic("[Theology.generateCults] min can not be less than 0")
	}
	allCults := t.getAllCults()
	if max > len(allCults) {
		panic("[Theology.generateCults] max can not be greater than allCults length")
	}

	cults := make([]*Cult, 0, len(allCults))
	for count := 0; count < 20; count++ {
		for _, cult := range allCults {
			if cult.Calc(t.religion, cult, cults) {
				cults = append(cults, &Cult{
					_religionMetadata: cult._religionMetadata,
					baseCoef:          cult.baseCoef,
					Name:              cult.Name,
					Calc:              cult.Calc,
				})
			}
			if len(cults) == max {
				break
			}
		}
		if len(cults) == max {
			break
		}
		if len(cults) >= min {
			break
		}
	}

	for _, cult := range cults {
		t.religion.UpdateMetadata(UpdateReligionMetadata(*t.religion.metadata, *cult._religionMetadata))
	}

	return cults
}

func (t *Theology) getAllCults() []*Cult {
	return []*Cult{
		{
			Name: "GodOfKingCult",
			_religionMetadata: &religionMetadata{
				Plutocratic:     0.75,
				Authoritaristic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				if r.Type.IsMonotheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "GodOfProsperityAndWealth",
			_religionMetadata: &religionMetadata{
				Plutocratic: 0.75,
				Hedonistic:  0.75,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				if r.Type.IsMonotheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "CultOfJustice",
			_religionMetadata: &religionMetadata{
				Lawful: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				if r.Type.IsMonotheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "GodOfWar",
			_religionMetadata: &religionMetadata{
				Chthonic:   0.5,
				Aggressive: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				if r.Type.IsMonotheism() {
					return false
				}
				if r.metadata.IsPacifistic() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "PatronSaintOfMerchants",
			_religionMetadata: &religionMetadata{
				Plutocratic: 0.5,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "GodOfDeath",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				if r.Type.IsMonotheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "GodOfWisdom",
			_religionMetadata: &religionMetadata{
				Lawful:      0.25,
				Educational: 0.75,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				if r.Type.IsMonotheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "GodOfSea",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Simple:       0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				if r.Type.IsMonotheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "GodOfLove",
			_religionMetadata: &religionMetadata{
				SexualActive: 1,
				Pacifistic:   0.25,
				Hedonistic:   0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				if r.Type.IsMonotheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SunWorship",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Simple:       0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "MoonWorship",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Chthonic:     0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "GodOfWine",
			_religionMetadata: &religionMetadata{
				Hedonistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				if r.Type.IsMonotheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "GodOfThunder",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				if r.Type.IsMonotheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "MotherGoddess",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				if r.Type.IsMonotheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HolyAnimals",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.5,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HolyPlants",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.5,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HolyFungus",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Chthonic:     0.25,
			},
			baseCoef: t.religion.M.BaseCoef - pm.RandFloat64InRange(0.01, 0.1),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HolyParasite",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: t.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{Log: true, Label: "Cults.HolyParasite"})
			},
		},
		{
			Name: "HolyInsects",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Chthonic:     1,
			},
			baseCoef: t.religion.M.BaseCoef - pm.RandFloat64InRange(0.01, 0.1),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HolyStones",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HolyRiver",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HolyLake",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HolyMountain",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "AncestorWorship",
			_religionMetadata: &religionMetadata{
				Lawful:          0.25,
				Authoritaristic: 0.25,
				Simple:          0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Trickster",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Chthonic:     0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				if r.Type.IsMonotheism() {
					return false
				}
				var addCoef float64
				if r.metadata.IsLawful() {
					addCoef += (-1) * pm.RandFloat64InRange(0.01, 0.1)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SpiritsOfDeath",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

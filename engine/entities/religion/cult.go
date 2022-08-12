package religion

import pm "persons_generator/engine/probability_machine"

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

func (t *Theology) generateCults(min, max int) []*trait {
	return generateTraits(t.religion, t.getAllCults(), generateTraitsOpts{
		Label: "Theology.generateCults",
		Min:   min,
		Max:   max,
	})
}

func (t *Theology) getAllCults() []*trait {
	return []*trait{
		{
			Name: "GodOfKingCult",
			_religionMetadata: &religionMetadata{
				Plutocratic:     0.75,
				Authoritaristic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "GodOfDeath",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				if r.Type.IsMonotheism() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "GodOfLove",
			_religionMetadata: &religionMetadata{
				SexualActive: 0.75,
				Pacifistic:   0.25,
				Hedonistic:   0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "GodOfWine",
			_religionMetadata: &religionMetadata{
				Hedonistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HolyParasite",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: t.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HolyInsects",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Chthonic:     1,
			},
			baseCoef: t.religion.M.BaseCoef - pm.RandFloat64InRange(0.01, 0.1),
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
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
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

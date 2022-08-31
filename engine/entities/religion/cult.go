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

func (t *Theology) generateCults(min, max int) ([]*trait, error) {
	return generateTraits(t.religion, t.getAllCults(), generateTraitsOpts{
		Label: "Theology.generateCults",
		Min:   min,
		Max:   max,
	})
}

func (t *Theology) getAllCults() []*trait {
	return []*trait{
		{
			Name: GodOfKingCultCultName,
			_religionMetadata: &religionMetadata{
				Plutocratic:     0.75,
				Authoritaristic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfProsperityAndWealthCultName,
			_religionMetadata: &religionMetadata{
				Plutocratic: 0.75,
				Hedonistic:  0.75,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: CultOfJusticeCultName,
			_religionMetadata: &religionMetadata{
				Lawful: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfWarCultName,
			_religionMetadata: &religionMetadata{
				Chthonic:   0.5,
				Aggressive: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}
				if r.metadata.IsPacifistic() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: PatronSaintOfMerchantsCultName,
			_religionMetadata: &religionMetadata{
				Plutocratic: 0.5,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfDeathCultName,
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfWisdomCultName,
			_religionMetadata: &religionMetadata{
				Lawful:      0.25,
				Educational: 0.75,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfSeaCultName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Simple:       0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfLoveCultName,
			_religionMetadata: &religionMetadata{
				SexualActive: 0.75,
				Pacifistic:   0.25,
				Hedonistic:   0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SunWorshipCultName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Simple:       0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: MoonWorshipCultName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Chthonic:     0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfWineCultName,
			_religionMetadata: &religionMetadata{
				Hedonistic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfThunderCultName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: MotherGoddessCultName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Simple:       0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyAnimalsCultName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.5,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyPlantsCultName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.5,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyFungusCultName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Chthonic:     0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				addCoef, err := pm.RandFloat64InRange(0.01, 0.1)
				if err != nil {
					return false, err
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef-addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyParasiteCultName,
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: t.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyInsectsCultName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Chthonic:     1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				addCoef, err := pm.RandFloat64InRange(0.01, 0.1)
				if err != nil {
					return false, err
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef-addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyStonesCultName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyRiverCultName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyLakeCultName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyMountainCultName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: AncestorWorshipCultName,
			_religionMetadata: &religionMetadata{
				Lawful:          0.25,
				Authoritaristic: 0.25,
				Simple:          0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: TricksterCultName,
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Chthonic:     0.25,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}
				var addCoef float64
				if r.metadata.IsLawful() {
					c, err := pm.RandFloat64InRange(0.01, 0.1)
					if err != nil {
						return false, err
					}
					addCoef -= c
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SpiritsOfDeathCultName,
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

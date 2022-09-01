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
			ReligionMetadata: &religionMetadata{
				Plutocratic:     0.75,
				Authoritaristic: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfProsperityAndWealthCultName,
			ReligionMetadata: &religionMetadata{
				Plutocratic: 0.75,
				Hedonistic:  0.75,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: CultOfJusticeCultName,
			ReligionMetadata: &religionMetadata{
				Lawful: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfWarCultName,
			ReligionMetadata: &religionMetadata{
				Chthonic:   0.5,
				Aggressive: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}
				if r.Metadata.IsPacifistic() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: PatronSaintOfMerchantsCultName,
			ReligionMetadata: &religionMetadata{
				Plutocratic: 0.5,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfDeathCultName,
			ReligionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfWisdomCultName,
			ReligionMetadata: &religionMetadata{
				Lawful:      0.25,
				Educational: 0.75,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfSeaCultName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Simple:       0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfLoveCultName,
			ReligionMetadata: &religionMetadata{
				SexualActive: 0.75,
				Pacifistic:   0.25,
				Hedonistic:   0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SunWorshipCultName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Simple:       0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: MoonWorshipCultName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Chthonic:     0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfWineCultName,
			ReligionMetadata: &religionMetadata{
				Hedonistic: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: GodOfThunderCultName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.75,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: MotherGoddessCultName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Simple:       0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyAnimalsCultName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.5,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyPlantsCultName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.5,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyFungusCultName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Chthonic:     0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				addCoef, err := pm.RandFloat64InRange(0.01, 0.1)
				if err != nil {
					return false, err
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef-addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyParasiteCultName,
			ReligionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			BaseCoef: t.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyInsectsCultName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Chthonic:     1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				addCoef, err := pm.RandFloat64InRange(0.01, 0.1)
				if err != nil {
					return false, err
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef-addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyStonesCultName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyRiverCultName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyLakeCultName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyMountainCultName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: AncestorWorshipCultName,
			ReligionMetadata: &religionMetadata{
				Lawful:          0.25,
				Authoritaristic: 0.25,
				Simple:          0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: TricksterCultName,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Chthonic:     0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsMonotheism() {
					return false, nil
				}
				var addCoef float64
				if r.Metadata.IsLawful() {
					c, err := pm.RandFloat64InRange(0.01, 0.1)
					if err != nil {
						return false, err
					}
					addCoef -= c
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SpiritsOfDeathCultName,
			ReligionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

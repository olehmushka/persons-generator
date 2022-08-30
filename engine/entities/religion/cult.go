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
			Name: "god_of_king_cult",
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
			Name: "god_of_prosperity_and_wealth",
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
			Name: "cult_of_justice",
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
			Name: "god_of_war",
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
			Name: "Patron_saint_of_merchants",
			_religionMetadata: &religionMetadata{
				Plutocratic: 0.5,
			},
			baseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "god_of_death",
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
			Name: "god_of_wisdom",
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
			Name: "God_of_sea",
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
			Name: "god_of_love",
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
			Name: "sun_worship",
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
			Name: "moon_worship",
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
			Name: "god_of_wine",
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
			Name: "god_of_thunder",
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
			Name: "mother_goddess",
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
			Name: "holy_animals",
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
			Name: "holy_plants",
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
			Name: "holy_fungus",
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
			Name: "holy_parasite",
			_religionMetadata: &religionMetadata{
				Chthonic: 1,
			},
			baseCoef: t.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "holy_insects",
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
			Name: "holy_stones",
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
			Name: "holy_river",
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
			Name: "holy_lake",
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
			Name: "holy_mountain",
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
			Name: "ancestor_worship",
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
			Name: "trickster",
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
					addCoef += (-1) * c
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "spirits_of_death",
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

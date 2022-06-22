package religion

import pm "persons_generator/probability-machine"

type Cult struct {
	_religionMetadata *updateReligionMetadata
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
			Name:              "GodOfKingCult",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "GodOfProsperityAndWealth",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "CultOfJustice",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "GodOfWar",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "PatronSaintOfMerchants",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "GodOfDeath",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "GodOfWisdom",
			_religionMetadata: &updateReligionMetadata{},
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "GodOfSea",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "GodOfLove",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "SunWorship",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "MoonWorship",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "GodOfWine",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "GodOfThunder",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "MotherGoddess",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "HolyAnimals",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "HolyPlants",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "HolyFungus",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "HolyParasite",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "HolyInsects",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				baseCoef := pm.RandFloat64InRange(0.5, 0.9)

				return CalculateProbabilityFromReligionMetadata(baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "HolyStones",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "HolyRiver",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "HolyLake",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "HolyMountain",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "AncestorWorship",
			_religionMetadata: &updateReligionMetadata{},
			baseCoef:          pm.RandFloat64InRange(0.5, 0.9),
			Calc: func(r *Religion, self *Cult, _ []*Cult) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

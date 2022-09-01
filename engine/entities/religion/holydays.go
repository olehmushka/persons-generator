package religion

import (
	"fmt"

	pm "persons_generator/engine/probability_machine"
)

type Holydays struct {
	religion *Religion `json:"-"`
	theology *Theology `json:"-"`

	Holydays []*trait `json:"holydays"`
}

func (t *Theology) generateHolydays() (*Holydays, error) {
	hs := &Holydays{religion: t.religion, theology: t}
	ts, err := generateTraits(t.religion, hs.getAllHolydays(), generateTraitsOpts{
		Label: "Holydays.generateHolydays",
		Min:   1,
		Max:   len(hs.getAllHolydays()),
	})
	if err != nil {
		return nil, err
	}
	hs.Holydays = ts

	return hs, nil
}

func (hs *Holydays) Print() {
	if len(hs.Holydays) > 0 {
		fmt.Printf("Holydays (religion_name=%s):\n", hs.religion.Name)
		for _, holyday := range hs.Holydays {
			fmt.Printf(" - %s\n", holyday.Name)
		}
	}
}

func (hs *Holydays) getAllHolydays() []*trait {
	return []*trait{
		{
			Name: SamhainHolyday,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Chthonic:     0.1,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: MysteriesHolyday,
			ReligionMetadata: &religionMetadata{
				Hedonistic:     0.1,
				Collectivistic: 0.25,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SummerSolsticeHolyday,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Educational:  0.1,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedHolydays []*trait) (bool, error) {
				var addCoef float64
				if len(hs.theology.Cults) > 0 {
					for _, cult := range hs.theology.Cults {
						if cult == nil {
							continue
						}
						if cult.Name == SunWorshipCultName || cult.Name == MoonWorshipCultName {
							coef, err := pm.RandFloat64InRange(0.05, 0.15)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
				}

				if len(hs.theology.Traits) > 0 {
					for _, trait := range hs.theology.Traits {
						if trait.Name == AstrologyTheologyTrait {
							coef, err := pm.RandFloat64InRange(0.1, 0.2)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
				}

				for _, holyday := range selectedHolydays {
					if holyday == nil {
						continue
					}
					if holyday.Name == WinterSolsticeHolyday {
						coef, err := pm.RandFloat64InRange(0.25, 0.5)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: WinterSolsticeHolyday,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Educational:  0.1,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedHolydays []*trait) (bool, error) {
				var addCoef float64
				if len(hs.theology.Cults) > 0 {
					for _, cult := range hs.theology.Cults {
						if cult == nil {
							continue
						}
						if cult.Name == SunWorshipCultName || cult.Name == MoonWorshipCultName {
							coef, err := pm.RandFloat64InRange(0.05, 0.15)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
				}

				if len(hs.theology.Traits) > 0 {
					for _, trait := range hs.theology.Traits {
						if trait.Name == AstrologyTheologyTrait {
							coef, err := pm.RandFloat64InRange(0.1, 0.2)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
				}

				for _, holyday := range selectedHolydays {
					if holyday == nil {
						continue
					}
					if holyday.Name == SummerSolsticeHolyday {
						coef, err := pm.RandFloat64InRange(0.25, 0.5)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: CalendarOfFeastsHolyday,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Educational:  0.5,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: NewYearHolyday,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.5,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DancePartyHolyday,
			ReligionMetadata: &religionMetadata{
				Simple:         0.75,
				Collectivistic: 0.5,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DrumPartyHolyday,
			ReligionMetadata: &religionMetadata{
				Simple:         0.75,
				Collectivistic: 0.5,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SocialFestivalsHolyday,
			ReligionMetadata: &religionMetadata{
				Collectivistic: 1,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: TreeCelebrationHolyday,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.75,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if len(hs.theology.Traits) > 0 {
					for _, trait := range hs.theology.Traits {
						if trait.Name == TreeConnectionTheologyTrait {
							coef, err := pm.RandFloat64InRange(0.1, 0.2)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

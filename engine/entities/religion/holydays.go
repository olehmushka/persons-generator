package religion

import (
	"fmt"

	pm "persons_generator/engine/probability_machine"
)

type Holydays struct {
	religion *Religion
	theology *Theology

	Holydays []*trait
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
			Name: "samhain",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Chthonic:     0.1,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "mysteries",
			_religionMetadata: &religionMetadata{
				Hedonistic:     0.1,
				Collectivistic: 0.25,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "summer_solstice",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Educational:  0.1,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if len(hs.theology.Cults) > 0 {
					for _, cult := range hs.theology.Cults {
						if cult == nil {
							continue
						}
						if cult.Name == "sun_worship" || cult.Name == "moon_worship" {
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
						if trait.Name == "astrology" {
							coef, err := pm.RandFloat64InRange(0.1, 0.2)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "winter_solstice",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Educational:  0.1,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedHolydays []*trait) (bool, error) {
				var addCoef float64
				if len(hs.theology.Cults) > 0 {
					for _, cult := range hs.theology.Cults {
						if cult == nil {
							continue
						}
						if cult.Name == "sun_worship" || cult.Name == "moon_worship" {
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
						if trait.Name == "astrology" {
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
					if holyday.Name == "summer_solstice" {
						coef, err := pm.RandFloat64InRange(0.25, 0.5)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "calendar_of_feasts",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Educational:  0.5,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "new_year",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.5,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "dance_party",
			_religionMetadata: &religionMetadata{
				Simple:         0.75,
				Collectivistic: 0.5,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "drum_party",
			_religionMetadata: &religionMetadata{
				Simple:         0.75,
				Collectivistic: 0.5,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "social_festivals",
			_religionMetadata: &religionMetadata{
				Collectivistic: 1,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "tree_celebration",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.75,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if len(hs.theology.Traits) > 0 {
					for _, trait := range hs.theology.Traits {
						if trait.Name == "tree_connection" {
							coef, err := pm.RandFloat64InRange(0.1, 0.2)
							if err != nil {
								return false, err
							}
							addCoef += coef
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

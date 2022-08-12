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

func (t *Theology) generateHolydays() *Holydays {
	hs := &Holydays{religion: t.religion, theology: t}
	hs.Holydays = generateTraits(t.religion, hs.getAllHolydays(), generateTraitsOpts{
		Label: "Holydays.generateHolydays",
		Min:   1,
		Max:   len(hs.getAllHolydays()),
	})

	return hs
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
			Name: "Samhain",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Chthonic:     0.1,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Mysteries",
			_religionMetadata: &religionMetadata{
				Hedonistic:     0.1,
				Collectivistic: 0.25,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SummerSolstice",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Educational:  0.1,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				if len(hs.theology.Cults) > 0 {
					for _, cult := range hs.theology.Cults {
						if cult == nil {
							continue
						}
						if cult.Name == "SunWorship" || cult.Name == "MoonWorship" {
							addCoef += pm.RandFloat64InRange(0.05, 0.15)
						}
					}
				}

				if len(hs.theology.Traits) > 0 {
					for _, trait := range hs.theology.Traits {
						if trait.Name == "Astrology" {
							addCoef += pm.RandFloat64InRange(0.1, 0.2)
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "WinterSolstice",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.75,
				Educational:  0.1,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedHolydays []*trait) bool {
				var addCoef float64
				if len(hs.theology.Cults) > 0 {
					for _, cult := range hs.theology.Cults {
						if cult == nil {
							continue
						}
						if cult.Name == "SunWorship" || cult.Name == "MoonWorship" {
							addCoef += pm.RandFloat64InRange(0.05, 0.15)
						}
					}
				}

				if len(hs.theology.Traits) > 0 {
					for _, trait := range hs.theology.Traits {
						if trait.Name == "Astrology" {
							addCoef += pm.RandFloat64InRange(0.1, 0.2)
						}
					}
				}

				for _, holyday := range selectedHolydays {
					if holyday == nil {
						continue
					}
					if holyday.Name == "SummerSolstice" {
						addCoef += pm.RandFloat64InRange(0.25, 0.5)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "CalendarOfFeasts",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.1,
				Educational:  0.5,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "NewYear",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.5,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "DanceParty",
			_religionMetadata: &religionMetadata{
				Simple:         0.75,
				Collectivistic: 0.5,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "DrumParty",
			_religionMetadata: &religionMetadata{
				Simple:         0.75,
				Collectivistic: 0.5,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "SocialFestivals",
			_religionMetadata: &religionMetadata{
				Collectivistic: 1,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "TreeCelebration",
			_religionMetadata: &religionMetadata{
				Naturalistic: 1,
				Simple:       0.75,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				if len(hs.theology.Traits) > 0 {
					for _, trait := range hs.theology.Traits {
						if trait.Name == "TreeConnection" {
							addCoef += pm.RandFloat64InRange(0.1, 0.2)
						}
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

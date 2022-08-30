package religion

import (
	"fmt"

	pm "persons_generator/engine/probability_machine"
)

type HolyScripture struct {
	religion *Religion
	attrs    *Attributes

	HasHolyScripture bool
	Traits           []*trait
}

func (as *Attributes) generateHolyScripture() (*HolyScripture, error) {
	hs := &HolyScripture{religion: as.religion, attrs: as}
	hasHolyScripture, err := hs.generateHasHolyScripture()
	if err != nil {
		return nil, err
	}
	hs.HasHolyScripture = hasHolyScripture
	if hs.HasHolyScripture {
		ts, err := generateTraits(as.religion, hs.getAllHolyScriptureTraits(), generateTraitsOpts{
			Label: "HolyScripture.generateTraits",
			Min:   1,
			Max:   3,
		})
		if err != nil {
			return nil, err
		}
		hs.Traits = ts
	}

	return hs, nil
}

func (hs *HolyScripture) generateHasHolyScripture() (bool, error) {
	primaryProbability, err := pm.RandFloat64InRange(0.5, 0.7)
	if err != nil {
		return false, err
	}
	if hs.religion.metadata.IsLawful() {
		p, err := pm.RandFloat64InRange(0.05, 0.1)
		if err != nil {
			return false, err
		}
		primaryProbability += p
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProbability))
}

func (hs *HolyScripture) Print() {
	if !hs.HasHolyScripture {
		fmt.Printf("Has not HolyScripture (religion_name=%s):\n", hs.religion.Name)
		return
	}
	if len(hs.Traits) > 0 {
		fmt.Printf("HolyScripture Traits (religion_name=%s):\n", hs.religion.Name)
		for _, trait := range hs.Traits {
			fmt.Printf(" - %s\n", trait.Name)
		}
		return
	}
	fmt.Printf("Has not HolyScripture (religion_name=%s):\n", hs.religion.Name)
}

func (hs *HolyScripture) getAllHolyScriptureTraits() []*trait {
	return []*trait{
		{
			Name: "commandments",
			_religionMetadata: &religionMetadata{
				Lawful: 1,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				switch {
				case r.Type.IsMonotheism():
					coef, err := pm.RandFloat64InRange(0.055, 0.105)
					if err != nil {
						return false, err
					}
					addCoef += coef
				case r.Type.IsDeityDualism():
					coef, err := pm.RandFloat64InRange(0.04, 0.09)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "devotional_code",
			_religionMetadata: &religionMetadata{
				Individualistic: 0.25,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "divine_law",
			_religionMetadata: &religionMetadata{
				Lawful:          1,
				Authoritaristic: 0.5,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				var addCoef float64
				for _, trait := range selectedTraits {
					if trait == nil {
						continue
					}
					if trait.Name == "commandments" {
						coef, err := pm.RandFloat64InRange(0.05, 0.1)
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
			Name: "magic_books",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Chthonic:     1,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "runes",
			_religionMetadata: &religionMetadata{
				Educational: 0.5,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "marching_hymns",
			_religionMetadata: &religionMetadata{
				Aggressive:     0.75,
				Collectivistic: 0.75,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "theogony",
			_religionMetadata: &religionMetadata{
				Lawful:      0.25,
				Educational: 0.25,
				Complicated: 0.5,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

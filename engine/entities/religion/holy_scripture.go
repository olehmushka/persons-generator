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

func (as *Attributes) generateHolyScripture() *HolyScripture {
	hs := &HolyScripture{religion: as.religion, attrs: as}
	hs.HasHolyScripture = hs.generateHasHolyScripture()
	if hs.HasHolyScripture {
		hs.Traits = generateTraits(as.religion, hs.getAllHolyScriptureTraits(), generateTraitsOpts{
			Label: "HolyScripture.generateTraits",
			Min:   1,
			Max:   3,
		})
	}

	return hs
}

func (hs *HolyScripture) generateHasHolyScripture() bool {
	primaryProbability := pm.RandFloat64InRange(0.5, 0.7)
	if hs.religion.metadata.IsLawful() {
		primaryProbability += pm.RandFloat64InRange(0.05, 0.1)
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
			Name: "Commandments",
			_religionMetadata: &religionMetadata{
				Lawful: 1,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				var addCoef float64
				switch {
				case r.Type.IsMonotheism():
					addCoef += pm.RandFloat64InRange(0.055, 0.105)
				case r.Type.IsDeityDualism():
					addCoef += pm.RandFloat64InRange(0.04, 0.09)
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "DevotionalCode",
			_religionMetadata: &religionMetadata{
				Individualistic: 0.25,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "DivineLaw",
			_religionMetadata: &religionMetadata{
				Lawful:          1,
				Authoritaristic: 0.5,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) bool {
				var addCoef float64
				for _, trait := range selectedTraits {
					if trait == nil {
						continue
					}
					if trait.Name == "Commandments" {
						addCoef += pm.RandFloat64InRange(0.05, 0.1)
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "MagicBooks",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Chthonic:     1,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Runes",
			_religionMetadata: &religionMetadata{
				Educational: 0.5,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "MarchingHymns",
			_religionMetadata: &religionMetadata{
				Aggressive:     0.75,
				Collectivistic: 0.75,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Theogony",
			_religionMetadata: &religionMetadata{
				Lawful:      0.25,
				Educational: 0.25,
				Complicated: 0.5,
			},
			baseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

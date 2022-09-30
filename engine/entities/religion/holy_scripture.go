package religion

import (
	"fmt"

	pm "persons_generator/engine/probability_machine"
)

type HolyScripture struct {
	religion *Religion   `json:"-" bson:"-"`
	attrs    *Attributes `json:"-" bson:"-"`

	HasHolyScripture bool     `json:"has_holy_scripture" bson:"has_holy_scripture"`
	Traits           []*trait `json:"traits" bson:"traits"`
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

func (hs *HolyScripture) SerializeTraits() []string {
	if hs == nil || !hs.HasHolyScripture {
		return []string{}
	}

	return extractTraitNames(hs.Traits)
}

func (hs *HolyScripture) generateHasHolyScripture() (bool, error) {
	primaryProbability, err := pm.RandFloat64InRange(0.5, 0.7)
	if err != nil {
		return false, err
	}
	if hs.religion.Metadata.IsLawful() {
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
			Name: CommandmentsHolyScriptureTrait,
			ReligionMetadata: &religionMetadata{
				Lawful: 1,
			},
			BaseCoef: hs.religion.M.BaseCoef,
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

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DevotionalCodeHolyScriptureTrait,
			ReligionMetadata: &religionMetadata{
				Individualistic: 0.25,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DivineLawHolyScriptureTrait,
			ReligionMetadata: &religionMetadata{
				Lawful:          1,
				Authoritaristic: 0.5,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				var addCoef float64
				for _, trait := range selectedTraits {
					if trait == nil {
						continue
					}
					if trait.Name == CommandmentsHolyScriptureTrait {
						coef, err := pm.RandFloat64InRange(0.05, 0.1)
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
			Name: MagicBooksHolyScriptureTrait,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Chthonic:     1,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: RunesHolyScriptureTrait,
			ReligionMetadata: &religionMetadata{
				Educational: 0.5,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: MarchingHymnsHolyScriptureTrait,
			ReligionMetadata: &religionMetadata{
				Aggressive:     0.75,
				Collectivistic: 0.75,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: TheogonyHolyScriptureTrait,
			ReligionMetadata: &religionMetadata{
				Lawful:      0.25,
				Educational: 0.25,
				Complicated: 0.5,
			},
			BaseCoef: hs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func GetHolyScriptureSimilarityCoef(hs1, hs2 *HolyScripture) float64 {
	if hs1 == nil && hs2 == nil {
		return 1
	}
	if hs1 == nil || hs2 == nil {
		return 0
	}

	if hs1.HasHolyScripture != hs2.HasHolyScripture {
		return 0
	}
	if !hs1.HasHolyScripture {
		return 1
	}

	return (GetTraitsSimilarityCoef(hs1.Traits, hs2.Traits) + 1) / 2
}

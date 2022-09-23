package religion

import (
	"fmt"

	"persons_generator/engine/entities/culture"
	pm "persons_generator/engine/probability_machine"
)

type Theology struct {
	religion *Religion `json:"-"`

	Traits            []*trait           `json:"traits"`
	Cults             []*trait           `json:"cults"`
	Rules             *Rules             `json:"rules"`
	Taboos            *Taboos            `json:"taboos"`
	Rituals           *Rituals           `json:"rituals"`
	Holydays          *Holydays          `json:"holydays"`
	Conversion        *Conversion        `json:"conversion"`
	MarriageTradition *MarriageTradition `json:"marriage_tradition"`
}

func NewTheology(r *Religion, c *culture.Culture) (*Theology, error) {
	t := &Theology{religion: r}
	ts, err := generateTraits(r, t.getAllTheologyTraits(), generateTraitsOpts{
		Label: "Theology.generateTraits",
		Min:   1,
		Max:   5,
	})
	if err != nil {
		return nil, err
	}
	t.Traits = ts

	minCults, maxCults := t.getCultsRange()
	cults, err := t.generateCults(minCults, maxCults)
	if err != nil {
		return nil, err
	}
	t.Cults = cults

	rules, err := t.generateRules()
	if err != nil {
		return nil, err
	}
	t.Rules = rules

	taboos, err := t.generateTaboos(c)
	if err != nil {
		return nil, err
	}
	t.Taboos = taboos

	rituals, err := t.generateRituals()
	if err != nil {
		return nil, err
	}
	t.Rituals = rituals

	holydays, err := t.generateHolydays()
	if err != nil {
		return nil, err
	}
	t.Holydays = holydays

	conversion, err := t.generateConversion()
	if err != nil {
		return nil, err
	}
	t.Conversion = conversion

	marriageTradition, err := t.generateMarriageTradition(c)
	if err != nil {
		return nil, err
	}
	t.MarriageTradition = marriageTradition

	return t, nil
}

func (t *Theology) IsZero() bool {
	return t == nil
}

func (t *Theology) SerializeTraits() []string {
	if t == nil {
		return []string{}
	}

	return extractTraitNames(t.Traits)
}

func (t *Theology) SerializeCults() []string {
	if t == nil {
		return []string{}
	}

	return extractTraitNames(t.Cults)
}

func (t *Theology) SerializeRules() []string {
	if t == nil || t.Rules == nil {
		return []string{}
	}

	return extractTraitNames(t.Rules.Rules)
}

func (t *Theology) SerializeTaboos() []map[string]string {
	if t == nil || t.Taboos == nil || len(t.Taboos.Taboos) == 0 {
		return []map[string]string{}
	}

	out := make([]map[string]string, len(t.Taboos.Taboos))
	for i := range out {
		out[i] = map[string]string{
			t.Taboos.Taboos[i].Name: t.Taboos.Taboos[i].Acceptance.String(),
		}
	}

	return out
}

func (t *Theology) SerializeRituals() []string {
	if t == nil {
		return []string{}
	}

	out := make([]string, 0, len(t.Rituals.Initiation)+len(t.Rituals.Funeral)+len(t.Rituals.Holyday)+len(t.Rituals.Sacrifice))
	out = append(out, extractTraitNames(t.Rituals.Initiation)...)
	out = append(out, extractTraitNames(t.Rituals.Funeral)...)
	out = append(out, extractTraitNames(t.Rituals.Holyday)...)
	out = append(out, extractTraitNames(t.Rituals.Sacrifice)...)

	return out
}

func (t *Theology) Print() {
	fmt.Printf("Theology (religion_name=%s):\n", t.religion.Name)
	fmt.Printf("Theology Traits (religion_name=%s):\n", t.religion.Name)
	for _, trait := range t.Traits {
		fmt.Printf(" - %s\n", trait.Name)
	}
	fmt.Printf("Cults (religion_name=%s):\n", t.religion.Name)
	for _, cult := range t.Cults {
		fmt.Printf(" - %s\n", cult.Name)
	}
	t.Rules.Print()
	t.Taboos.Print()
	t.Rituals.Print()
	t.Holydays.Print()
	t.Conversion.Print()
	t.MarriageTradition.Print()
}

func (t *Theology) getAllTheologyTraits() []*trait {
	return []*trait{
		{
			Name: MessiahTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Lawful:          0.25,
				Authoritaristic: 0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				switch {
				case r.Type.IsMonotheism():
					addCoef += 0.1
				case r.Type.IsPolytheism():
					addCoef += 0.01
				case r.Type.IsDeityDualism():
					addCoef += 0.05
				case r.Type.IsDeism():
					addCoef += 0.005
				default:
					return false, nil
				}
				if t.religion.Attributes != nil {
					if t.religion.Attributes.HolyScripture != nil {
						if t.religion.Attributes.HolyScripture.HasHolyScripture {
							coef, err := pm.RandFloat64InRange(0.05, 0.2)
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
		{
			Name: ProphetsTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Lawful:          0.5,
				Authoritaristic: 0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ReligiousLawTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Lawful: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.Metadata.IsLawful() {
					return true, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: ReincarnationTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Individualistic: 0.75,
				Complicated:     0.75,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.HasReincarnation() {
					return true, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SanctionedFalseConversionsTheologyTrait, // If your life is threatened, it is acceptable to confess a false faith, as long as you keep the truth in your heart.
			ReligionMetadata: &religionMetadata{
				Pacifistic:      0.25,
				Liberal:         1,
				Individualistic: 0.75,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: TreeConnectionTheologyTrait, // Trees are the essence of life, and we must be near them.
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Chthonic:     0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.Metadata.IsNaturalistic() {
					return true, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: AnimalConnectionTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Chthonic:     0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.Metadata.IsNaturalistic() {
					return true, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: BlindsightTheologyTrait, // Only the blind can perceive true reality
			ReligionMetadata: &religionMetadata{
				Chthonic: 1,
				Ascetic:  0.5,
			},
			BaseCoef: t.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: AstrologyTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
				Educational:  0.5,
				Complicated:  1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: RepentanceTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Lawful:  0.25,
				Ascetic: 0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: MartyrdomTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Ascetic:        0.25,
				Collectivistic: 0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				var addCoef float64
				for _, trait := range selectedTraits {
					switch trait.Name {
					case MessiahTheologyTrait:
						coef, err := pm.RandFloat64InRange(0.01, 0.125)
						if err != nil {
							return false, err
						}
						addCoef += coef
					case ProphetsTheologyTrait:
						coef, err := pm.RandFloat64InRange(0.1, 0.2)
						if err != nil {
							return false, err
						}
						addCoef += coef
					}
				}
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: IndulgencesTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Plutocratic: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Type.IsAtheism() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolinessTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Ascetic:         0.75,
				Individualistic: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				for _, goal := range r.Doctrine.HighGoal.Goals {
					switch goal.Name {
					case BringHolinessDownToTheWorldHighGoal:
						coef, err := pm.RandFloat64InRange(0.03, 0.1)
						if err != nil {
							return false, err
						}
						addCoef += coef
					case BecomePerfectAndSaintsHighGoal:
						coef, err := pm.RandFloat64InRange(0.05, 0.15)
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
			Name: CelibacyTheologyTrait,
			ReligionMetadata: &religionMetadata{
				SexualStrictness: 1,
				Ascetic:          1,
				Individualistic:  0.5,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.Attributes.Clerics.HasClerics {
					if r.Attributes.Clerics.Limitations.Marriage.IsDisallowed() {
						coef, err := pm.RandFloat64InRange(0.05, 0.2)
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
			Name: FlagellantismTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Ascetic: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				var addCoef float64
				for _, trait := range selectedTraits {
					switch trait.Name {
					case RepentanceTheologyTrait:
						fallthrough
					case MartyrdomTheologyTrait:
						coef, err := pm.RandFloat64InRange(0.05, 0.15)
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
			Name: FeedTheWorldTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Altruistic: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HolyArmyTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Chthonic:       0.25,
				Aggressive:     1,
				Collectivistic: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				var addCoef float64
				if r.Metadata.IsAggressive() {
					coef, err := pm.RandFloat64InRange(0.05, 0.125)
					if err != nil {
						return false, err
					}
					addCoef += coef
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DefendersOfFaithTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Lawful:     0.25,
				Aggressive: 0.75,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: NonViolentResistanceTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Pacifistic: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Metadata.IsAggressive() {
					p, err := pm.RandFloat64InRange(0, 0.05)
					if err != nil {
						return false, err
					}

					return pm.GetRandomBool(p)
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: NoMoreKillingTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Altruistic: 1,
				Pacifistic: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				if r.Metadata.IsAggressive() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: AnimalMessengersTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 1,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: SanctuaryOfMindTheologyTrait,
			ReligionMetadata: &religionMetadata{
				Individualistic: 1,
				Complicated:     0.25,
			},
			BaseCoef: t.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: PhantomsTheologyTrait,
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

func (t *Theology) HasReincarnation() bool {
	if len(t.Traits) == 0 {
		return false
	}

	for _, trait := range t.Traits {
		if trait == nil {
			continue
		}
		if trait.Name == ReincarnationTheologyTrait {
			return true
		}
	}

	return false
}

func GetTheologySimilarityCoef(t1, t2 *Theology) float64 {
	if t1.IsZero() && t2.IsZero() {
		return 1
	}
	if t1.IsZero() || t2.IsZero() {
		return 0
	}

	similarityTraits := []struct {
		value float64
		coef  float64
	}{
		{
			value: GetTraitsSimilarityCoef(t1.Traits, t2.Traits),
			coef:  0.05,
		},
		{
			value: GetTraitsSimilarityCoef(t1.Cults, t2.Cults),
			coef:  0.2,
		},
		{
			value: GetRulesSimilarityCoef(t1.Rules, t2.Rules),
			coef:  0.1,
		},
		{
			value: GetTaboosSimilarityCoef(t1.Taboos, t2.Taboos),
			coef:  0.3,
		},
		{
			value: GetRitualsSimilarityCoef(t1.Rituals, t2.Rituals),
			coef:  0.10,
		},
		{
			value: GetHolydaysSimilarityCoef(t1.Holydays, t2.Holydays),
			coef:  0.05,
		},
		{
			value: GetConversionSimilarityCoef(t1.Conversion, t2.Conversion),
			coef:  0.05,
		},
		{
			value: GetMarriageTraditionSimilarityCoef(t1.MarriageTradition, t2.MarriageTradition),
			coef:  0.15,
		},
	}

	var out float64
	for _, t := range similarityTraits {
		out += t.value * t.coef
	}

	return out
}

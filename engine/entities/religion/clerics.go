package religion

import (
	"fmt"
	"strings"

	g "persons_generator/engine/entities/gender"
	pm "persons_generator/engine/probability_machine"
)

type Clerics struct {
	religion *Religion   `json:"-" bson:"-"`
	attrs    *Attributes `json:"-" bson:"-"`

	HasClerics  bool                `json:"has_clerics" bson:"has_clerics"`
	Appointment *ClericsAppointment `json:"appointment" bson:"appointment"`
	Limitations *ClericsLimitations `json:"limitations" bson:"limitations"`
	Traits      []*trait            `json:"traits" bson:"traits"`
	Functions   []*trait            `json:"functions" bson:"functions"`
}

func (as *Attributes) generateClerics() (*Clerics, error) {
	cs := &Clerics{religion: as.religion, attrs: as}

	hasClerics, err := cs.generateHasClerics()
	if err != nil {
		return nil, err
	}
	cs.HasClerics = hasClerics
	if cs.HasClerics {
		a, err := cs.generateAppointment()
		if err != nil {
			return nil, err
		}
		cs.Appointment = a
		ls, err := cs.generateLimitations()
		if err != nil {
			return nil, err
		}
		cs.Limitations = ls
		ts, err := generateTraits(as.religion, cs.getAllClericsTraits(), generateTraitsOpts{
			Label: "Clerics.generateTraits",
			Min:   1,
			Max:   3,
		})
		if err != nil {
			return nil, err
		}
		cs.Traits = ts
		fs, err := generateTraits(as.religion, cs.getAllClericsFunctions(), generateTraitsOpts{
			Label: "Clerics.generateFunctions",
			Min:   1,
			Max:   2,
		})
		if err != nil {
			return nil, err
		}
		cs.Functions = fs
	}

	return cs, nil
}

func (cs *Clerics) Print() {
	if !cs.HasClerics {
		fmt.Printf("No clerics (religion_name=%s):\n", cs.religion.Name)
		return
	}
	fmt.Printf("Clerics (religion_name=%s):\n", cs.religion.Name)
	cs.Appointment.Print()
	cs.Limitations.Print()
	fmt.Printf("Clerics Traits (religion_name=%s):\n", cs.religion.Name)
	for _, trait := range cs.Traits {
		fmt.Printf(" - %s\n", trait.Name)
	}
	fmt.Printf("Clerics Functions (religion_name=%s):\n", cs.religion.Name)
	for _, function := range cs.Functions {
		fmt.Printf(" - %s\n", function.Name)
	}
}

func (cs *Clerics) SerializeAppointment() string {
	if cs == nil || !cs.HasClerics || cs.Appointment == nil {
		return ""
	}

	out := make([]string, 0, 2)
	if cs.Appointment.IsCivil {
		out = append(out, "clerics are civil.")
	} else {
		out = append(out, "clerics are not civil.")
	}
	if cs.Appointment.IsRevocable {
		out = append(out, "cleric position is revocable.")
	} else {
		out = append(out, "cleric position is not revocable.")
	}

	return strings.Join(out, " ")
}

func (cs *Clerics) SerializeLimitations() string {
	if cs == nil || !cs.HasClerics || cs.Limitations == nil {
		return ""
	}

	out := make([]string, 0, 2)
	switch cs.Limitations.AcceptableGender {
	case g.OnlyMen:
		out = append(out, "only men can be clerics.")
	case g.MenAndWomen:
		out = append(out, "men amd women can be clerics.")
	case g.OnlyWomen:
		out = append(out, "only women can be clerics.")
	}

	switch cs.Limitations.Marriage {
	case AlwaysAllowed:
		out = append(out, "clerics can get married.")
	case MustBeApproved:
		out = append(out, "clerics can get married just after approval.")
	case Disallowed:
		out = append(out, "clerics can not get married.")
	}

	return strings.Join(out, " ")
}

func (cs *Clerics) SerializeTraits() []string {
	if cs == nil || !cs.HasClerics {
		return []string{}
	}

	return extractTraitNames(cs.Traits)
}

func (cs *Clerics) SerializeFunctions() []string {
	if cs == nil || !cs.HasClerics {
		return []string{}
	}

	return extractTraitNames(cs.Functions)
}

func (cs *Clerics) generateHasClerics() (bool, error) {
	primaryProbability, err := pm.RandFloat64InRange(0.6, 0.85)
	if err != nil {
		return false, err
	}
	if cs.religion.Metadata.IsLawful() {
		p, err := pm.RandFloat64InRange(0.05, 0.15)
		if err != nil {
			return false, err
		}
		primaryProbability += p
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProbability))
}

type ClericsAppointment struct {
	religion *Religion `json:"-" bson:"-"`

	IsCivil     bool `json:"is_civil" bson:"is_civil"`
	IsRevocable bool `json:"is_revocable" bson:"is_revocable"`
}

func (cs *Clerics) generateAppointment() (*ClericsAppointment, error) {
	out := &ClericsAppointment{
		religion: cs.religion,
	}
	isCivil, err := CalculateProbabilityFromReligionMetadata(cs.religion.M.BaseCoef, cs.religion, &religionMetadata{Liberal: 1}, CalcProbOpts{})
	if err != nil {
		return nil, err
	}
	out.IsCivil = isCivil
	isRevocable, err := CalculateProbabilityFromReligionMetadata(cs.religion.M.BaseCoef, cs.religion, &religionMetadata{Authoritaristic: 1}, CalcProbOpts{})
	if err != nil {
		return nil, err
	}
	out.IsRevocable = isRevocable

	return out, nil
}

func (cas *ClericsAppointment) Print() {
	if cas == nil {
		return
	}
	isCivilStr := "can not be"
	if cas.IsCivil {
		isCivilStr = "can be"
	}
	IsRevocableStr := "is not"
	if cas.IsRevocable {
		IsRevocableStr = "is"
	}
	fmt.Printf("Cleric position %s civil and it %s revocable\n", isCivilStr, IsRevocableStr)
}

type ClericsLimitations struct {
	religion *Religion `json:"-" bson:"-"`

	AcceptableGender g.Acceptance `json:"acceptable" bson:"acceptable"`
	Marriage         Permission   `json:"marriage" bson:"marriage"`
}

func (cs *Clerics) generateLimitations() (*ClericsLimitations, error) {
	cls := &ClericsLimitations{
		religion: cs.religion,
	}
	ag, err := g.GetClericalCustom(cs.religion.M.BaseCoef, cs.religion.GenderDominance)
	if err != nil {
		return nil, err
	}
	cls.AcceptableGender = ag

	alwaysAllowed, err := pm.RandFloat64InRange(0.25, 0.75)
	if err != nil {
		return nil, err
	}
	mustBeApproved, err := pm.RandFloat64InRange(0.25, 0.75)
	if err != nil {
		return nil, err
	}
	disallowed, err := pm.RandFloat64InRange(0.25, 0.75)
	if err != nil {
		return nil, err
	}
	m, err := getPermissionByProbability(alwaysAllowed, mustBeApproved, disallowed)
	if err != nil {
		return nil, err
	}
	cls.Marriage = m

	return cls, nil
}

func (cls *ClericsLimitations) Print() {
	if cls == nil {
		return
	}
	fmt.Printf("%s can be clerics and marriage is %s\n", cls.AcceptableGender, cls.Marriage)
}

func (cs *Clerics) getAllClericsTraits() []*trait {
	return []*trait{
		{
			Name: NakedPriestsClericTrait,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Chthonic:     0.25,
			},
			BaseCoef: cs.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				addCoef, err := pm.RandFloat64InRange(0.05, 0.15)
				if err != nil {
					return false, err
				}
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef-addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HeadOfFaithClericTrait,
			ReligionMetadata: &religionMetadata{
				Plutocratic:     0.5,
				Lawful:          0.5,
				Authoritaristic: 1,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: InternalCirclesOfInitiationClericTrait,
			ReligionMetadata: &religionMetadata{
				Authoritaristic: 0.5,
				Collectivistic:  0.5,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				var addCoef float64
				for _, trait := range selectedTraits {
					if trait.Name == HeadOfFaithClericTrait {
						c, err := pm.RandFloat64InRange(0.01, 0.1)
						if err != nil {
							return false, err
						}
						addCoef += c
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef+addCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: MendicantPreachersClericTrait,
			ReligionMetadata: &religionMetadata{
				Ascetic: 0.75,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: MonasticismClericTrait,
			ReligionMetadata: &religionMetadata{
				SexualStrictness: 0.5,
				Ascetic:          0.75,
				Authoritaristic:  0.75,
				Collectivistic:   0.1,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: VowOfPovertyClericTrait,
			ReligionMetadata: &religionMetadata{
				Ascetic:         1,
				Authoritaristic: 0.25,
				Individualistic: 1,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: PatronshipClericTrait,
			ReligionMetadata: &religionMetadata{
				Plutocratic:     0.25,
				Authoritaristic: 0.5,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: WarriorMonksClericTrait,
			ReligionMetadata: &religionMetadata{
				Aggressive:      1,
				Ascetic:         0.25,
				Authoritaristic: 0.25,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				var hasMonasticism bool
				for _, trait := range selectedTraits {
					if trait.Name == MonasticismClericTrait {
						hasMonasticism = true
					}
				}
				if !hasMonasticism {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: IsHereditaryClericTrait,
			ReligionMetadata: &religionMetadata{
				Plutocratic: 1,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				if cs.Limitations.Marriage.IsDisallowed() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (cs *Clerics) getAllClericsFunctions() []*trait {
	return []*trait{
		{
			Name: ControlClericFunction,
			ReligionMetadata: &religionMetadata{
				Authoritaristic: 1,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: AlmsAndPacificationClericFunction,
			ReligionMetadata: &religionMetadata{
				Pacifistic: 0.75,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: HealClericFuction,
			ReligionMetadata: &religionMetadata{
				Altruistic: 1,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: RecruitmentClericFunction,
			ReligionMetadata: &religionMetadata{
				Collectivistic: 1,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: TeachClericFunction,
			ReligionMetadata: &religionMetadata{
				Lawful:      0.5,
				Educational: 1,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: OracleClericFunction,
			ReligionMetadata: &religionMetadata{
				Chthonic: 0.25,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DivinerClericFunction,
			ReligionMetadata: &religionMetadata{
				Chthonic: 0.25,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: DruidClericFunction,
			ReligionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Chthonic:     0.25,
			},
			BaseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.BaseCoef, r, self.ReligionMetadata, CalcProbOpts{})
			},
		},
	}
}

func GetClericsSimilarityCoef(c1, c2 *Clerics) float64 {
	if c1 == nil && c2 == nil {
		return 1
	}
	if c1 == nil || c2 == nil {
		return 0
	}

	similarityTraits := []struct {
		enable bool
		value  float64
		coef   float64
	}{
		{
			enable: c1.HasClerics == c2.HasClerics,
			value:  1,
			coef:   0.3,
		},
		{
			enable: true,
			value:  GetClericsAppointmentSimilarityCoef(c1.Appointment, c2.Appointment),
			coef:   0.15,
		},
		{
			enable: true,
			value:  GetClericsLimitationsSimilarityCoef(c1.Limitations, c2.Limitations),
			coef:   0.15,
		},
		{
			enable: true,
			value:  GetTraitsSimilarityCoef(c1.Traits, c2.Traits),
			coef:   0.2,
		},
		{
			enable: true,
			value:  GetTraitsSimilarityCoef(c1.Functions, c2.Functions),
			coef:   0.2,
		},
	}

	var out float64
	for _, t := range similarityTraits {
		if t.enable {
			out += t.value * t.coef
		}
	}

	return out
}

func GetClericsAppointmentSimilarityCoef(c1, c2 *ClericsAppointment) float64 {
	if c1 == nil && c2 == nil {
		return 1
	}
	if c1 == nil || c2 == nil {
		return 0
	}

	var out float64
	if c1.IsCivil == c2.IsCivil {
		out += 0.5
	}
	if c1.IsRevocable == c2.IsRevocable {
		out += 0.5
	}

	return out
}

func GetClericsLimitationsSimilarityCoef(c1, c2 *ClericsLimitations) float64 {
	if c1 == nil && c2 == nil {
		return 1
	}
	if c1 == nil || c2 == nil {
		return 0
	}

	var out float64
	if c1.AcceptableGender == c2.AcceptableGender {
		out += 0.5
	}
	if c1.Marriage == c2.Marriage {
		out += 0.5
	}

	return out
}

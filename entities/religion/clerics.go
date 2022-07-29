package religion

import (
	"fmt"

	g "persons_generator/entities/gender"
	pm "persons_generator/probability_machine"
)

type Clerics struct {
	religion *Religion
	attrs    *Attributes

	HasClerics  bool
	Appointment *ClericsAppointment
	Limitations *ClericsLimitations
	Traits      []*trait
	Functions   []*trait
}

func (as *Attributes) generateClerics() *Clerics {
	cs := &Clerics{religion: as.religion, attrs: as}
	if !cs.generateHasClerics() {
		return cs
	}

	cs.HasClerics = cs.generateHasClerics()
	if cs.HasClerics {
		cs.Appointment = cs.generateAppointment()
		cs.Limitations = cs.generateLimitations()
		cs.Traits = generateTraits(as.religion, cs.getAllClericsTraits(), generateTraitsOpts{
			Label: "Clerics.generateTraits",
			Min:   1,
			Max:   3,
		})
		cs.Functions = generateTraits(as.religion, cs.getAllClericsFunctions(), generateTraitsOpts{
			Label: "Clerics.generateFunctions",
			Min:   1,
			Max:   2,
		})
	}

	return cs
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

func (cs *Clerics) generateHasClerics() bool {
	primaryProbability := pm.RandFloat64InRange(0.6, 0.85)
	if cs.religion.metadata.IsLawful() {
		primaryProbability += pm.RandFloat64InRange(0.05, 0.15)
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProbability))
}

type ClericsAppointment struct {
	religion *Religion

	IsCivil     bool
	IsRevocable bool
}

func (cs *Clerics) generateAppointment() *ClericsAppointment {
	return &ClericsAppointment{
		religion:    cs.religion,
		IsCivil:     CalculateProbabilityFromReligionMetadata(cs.religion.M.BaseCoef, cs.religion, &religionMetadata{Liberal: 1}, CalcProbOpts{}),
		IsRevocable: CalculateProbabilityFromReligionMetadata(cs.religion.M.BaseCoef, cs.religion, &religionMetadata{Authoritaristic: 1}, CalcProbOpts{}),
	}
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
	religion *Religion

	AcceptableGender g.Acceptance
	Marriage         Permission
}

func (cs *Clerics) generateLimitations() *ClericsLimitations {
	cls := &ClericsLimitations{
		religion: cs.religion,

		AcceptableGender: g.GetClericalCustom(cs.religion.M.BaseCoef, cs.religion.GenderDominance),
	}

	var (
		alwaysAllowed  = pm.RandFloat64InRange(0.25, 0.75)
		mustBeApproved = pm.RandFloat64InRange(0.25, 0.75)
		disallowed     = pm.RandFloat64InRange(0.25, 0.75)
	)
	cls.Marriage = getPermissionByProbability(alwaysAllowed, mustBeApproved, disallowed)

	return cls
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
			Name: "NakedPriests",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Chthonic:     0.25,
			},
			baseCoef: cs.religion.M.LowBaseCoef - pm.RandFloat64InRange(0.05, 0.15),
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "HeadOfFaith",
			_religionMetadata: &religionMetadata{
				Plutocratic:     0.5,
				Lawful:          0.5,
				Authoritaristic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "InternalCirclesOfInitiation",
			_religionMetadata: &religionMetadata{
				Authoritaristic: 0.5,
				Collectivistic:  0.5,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) bool {
				var addCoef float64
				for _, trait := range selectedTraits {
					if trait.Name == "HeadOfFaith" {
						addCoef += pm.RandFloat64InRange(0.01, 0.1)
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "MendicantPreachers",
			_religionMetadata: &religionMetadata{
				Ascetic: 0.75,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Monasticism",
			_religionMetadata: &religionMetadata{
				SexualStrictness: 0.5,
				Ascetic:          0.75,
				Authoritaristic:  0.75,
				Collectivistic:   0.1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "VowOfPoverty",
			_religionMetadata: &religionMetadata{
				Ascetic:         1,
				Authoritaristic: 0.25,
				Individualistic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Patronship",
			_religionMetadata: &religionMetadata{
				Plutocratic:     0.25,
				Authoritaristic: 0.5,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "WarriorMonks",
			_religionMetadata: &religionMetadata{
				Aggressive:      1,
				Ascetic:         0.25,
				Authoritaristic: 0.25,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) bool {
				var hasMonasticism bool
				for _, trait := range selectedTraits {
					if trait.Name == "Monasticism" {
						hasMonasticism = true
					}
				}
				if !hasMonasticism {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "IsHereditary",
			_religionMetadata: &religionMetadata{
				Plutocratic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) bool {
				if cs.Limitations.Marriage.IsDisallowed() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (cs *Clerics) getAllClericsFunctions() []*trait {
	return []*trait{
		{
			Name: "Control",
			_religionMetadata: &religionMetadata{
				Authoritaristic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "AlmsAndPacification",
			_religionMetadata: &religionMetadata{
				Pacifistic: 0.75,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Heal",
			_religionMetadata: &religionMetadata{
				Altruistic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Recruitment",
			_religionMetadata: &religionMetadata{
				Collectivistic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Teach",
			_religionMetadata: &religionMetadata{
				Lawful:      0.5,
				Educational: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Oracle",
			_religionMetadata: &religionMetadata{
				Chthonic: 0.25,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Diviner",
			_religionMetadata: &religionMetadata{
				Chthonic: 0.25,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Druid",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Chthonic:     0.25,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

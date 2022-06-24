package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
)

type Clerics struct {
	religion *Religion
	attrs    *Attributes

	HasClerics  bool
	Appointment *ClericsAppointment
	Limitations *ClericsLimitations
	Traits      []*clericsTrait
	Functions   []*clericsFunction
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
		cs.Traits = cs.generateTraits(1, 3)
		cs.Functions = cs.generateFunctions(1, 2)
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
	if cs.religion.IsLawful() {
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
		isCivilStr = "is"
	}
	fmt.Printf("Cleric position %s civil and it %s revocable\n", isCivilStr, IsRevocableStr)
}

type ClericsLimitations struct {
	religion *Religion

	AcceptableGender GenderAcceptance
	Marriage         Permission
}

func (cs *Clerics) generateLimitations() *ClericsLimitations {
	cls := &ClericsLimitations{religion: cs.religion}

	var (
		baseCoef    = cs.religion.M.BaseCoef
		men         = pm.RandFloat64InRange(0.05, 0.15)
		menAndWomen = pm.RandFloat64InRange(0.05, 0.15)
		women       = pm.RandFloat64InRange(0.05, 0.15)
	)
	switch {
	case cs.religion.GenderDominance.IsMaleDominate():
		men += baseCoef * pm.RandFloat64InRange(0.15, 25) * cs.religion.GenderDominance.GetCoef()
	case cs.religion.GenderDominance.IsEquality():
		menAndWomen += baseCoef * pm.RandFloat64InRange(0.15, 25) * cs.religion.GenderDominance.GetCoef()
	case cs.religion.GenderDominance.IsFemaleDominate():
		women += baseCoef * pm.RandFloat64InRange(0.15, 25) * cs.religion.GenderDominance.GetCoef()
	}
	cls.AcceptableGender = getGenderAcceptanceByProbability(baseCoef*men, baseCoef*menAndWomen, baseCoef*women)

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

type clericsTrait struct {
	_religionMetadata *religionMetadata
	baseCoef          float64
	Name              string
	Calc              func(r *Religion, self *clericsTrait, selectedTraits []*clericsTrait) bool
}

func (cs *Clerics) generateTraits(min, max int) []*clericsTrait {
	if min < 0 {
		panic("[Clerics.generateTraits] min can not be less than 0")
	}
	allTraits := cs.getAllClericsTraits()
	if max > len(allTraits) {
		panic("[Clerics.generateTraits] max can not be greater than allTraits length")
	}

	traits := make([]*clericsTrait, 0, len(allTraits))
	for count := 0; count < 20; count++ {
		for _, trait := range allTraits {
			if trait.Calc(cs.religion, trait, traits) {
				traits = append(traits, &clericsTrait{
					_religionMetadata: trait._religionMetadata,
					baseCoef:          trait.baseCoef,
					Name:              trait.Name,
					Calc:              trait.Calc,
				})
			}
			if len(traits) == max {
				break
			}
		}
		if len(traits) == max {
			break
		}
		if len(traits) >= min {
			break
		}
	}

	for _, trait := range traits {
		cs.religion.UpdateMetadata(UpdateReligionMetadata(*cs.religion.metadata, *trait._religionMetadata))
	}

	return traits
}

func (cs *Clerics) getAllClericsTraits() []*clericsTrait {
	return []*clericsTrait{
		{
			Name: "NakedPriests",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Chthonic:     0.25,
			},
			baseCoef: cs.religion.M.LowBaseCoef - pm.RandFloat64InRange(0.05, 0.15),
			Calc: func(r *Religion, self *clericsTrait, _ []*clericsTrait) bool {
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
			Calc: func(r *Religion, self *clericsTrait, _ []*clericsTrait) bool {
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
			Calc: func(r *Religion, self *clericsTrait, selectedTraits []*clericsTrait) bool {
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
			Calc: func(r *Religion, self *clericsTrait, _ []*clericsTrait) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Monasticism",
			_religionMetadata: &religionMetadata{
				Ascetic:         0.75,
				Authoritaristic: 0.75,
				Collectivistic:  0.25,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *clericsTrait, _ []*clericsTrait) bool {
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
			Calc: func(r *Religion, self *clericsTrait, _ []*clericsTrait) bool {
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
			Calc: func(r *Religion, self *clericsTrait, _ []*clericsTrait) bool {
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
			Calc: func(r *Religion, self *clericsTrait, selectedTraits []*clericsTrait) bool {
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
			Calc: func(r *Religion, self *clericsTrait, selectedTraits []*clericsTrait) bool {
				if cs.Limitations.Marriage.IsDisallowed() {
					return false
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

type clericsFunction struct {
	_religionMetadata *religionMetadata
	baseCoef          float64
	Name              string
	Calc              func(r *Religion, self *clericsFunction, selectedFunctions []*clericsFunction) bool
}

func (cs *Clerics) generateFunctions(min, max int) []*clericsFunction {
	if min < 0 {
		panic("[Clerics.generateFunctions] min can not be less than 0")
	}
	allFunctions := cs.getAllClericsFunctions()
	if max > len(allFunctions) {
		panic("[Clerics.generateFunctions] max can not be greater than allFunctions length")
	}

	functions := make([]*clericsFunction, 0, len(allFunctions))
	for count := 0; count < 20; count++ {
		for _, function := range allFunctions {
			if function.Calc(cs.religion, function, functions) {
				functions = append(functions, &clericsFunction{
					_religionMetadata: function._religionMetadata,
					baseCoef:          function.baseCoef,
					Name:              function.Name,
					Calc:              function.Calc,
				})
			}
			if len(functions) == max {
				break
			}
		}
		if len(functions) == max {
			break
		}
		if len(functions) >= min {
			break
		}
	}

	for _, function := range functions {
		cs.religion.UpdateMetadata(UpdateReligionMetadata(*cs.religion.metadata, *function._religionMetadata))
	}

	return functions
}

func (cs *Clerics) getAllClericsFunctions() []*clericsFunction {
	return []*clericsFunction{
		{
			Name: "Control",
			_religionMetadata: &religionMetadata{
				Authoritaristic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *clericsFunction, _ []*clericsFunction) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "AlmsAndPacification",
			_religionMetadata: &religionMetadata{
				Pacifistic: 0.75,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *clericsFunction, _ []*clericsFunction) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Heal",
			_religionMetadata: &religionMetadata{
				Altruistic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *clericsFunction, _ []*clericsFunction) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Recruitment",
			_religionMetadata: &religionMetadata{
				Collectivistic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *clericsFunction, _ []*clericsFunction) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Teach",
			_religionMetadata: &religionMetadata{
				Lawful: 0.5,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *clericsFunction, _ []*clericsFunction) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Oracle",
			_religionMetadata: &religionMetadata{
				Chthonic: 0.25,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *clericsFunction, _ []*clericsFunction) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "Diviner",
			_religionMetadata: &religionMetadata{
				Chthonic: 0.25,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *clericsFunction, _ []*clericsFunction) bool {
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
			Calc: func(r *Religion, self *clericsFunction, _ []*clericsFunction) bool {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (cs *Clerics) GetAggressiveCriterias() float64 {
	if len(cs.Traits) == 0 {
		return 0
	}

	var criterias float64
	for _, trait := range cs.Traits {
		switch trait.Name {
		case "WarriorMonks":
			criterias += 1
		}
	}

	return criterias
}

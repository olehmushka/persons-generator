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
	cs.Appointment = cs.generateAppointment()
	cs.Limitations = cs.generateLimitations()
	cs.Traits = cs.generateTraits(1, 3)
	cs.Functions = cs.generateFunctions(1, 2)

	return cs
}

func (cs *Clerics) Print() {
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

type ClericsAppointment struct {
	religion *Religion

	IsHereditary bool // make it traits
	IsCivil      bool
	IsRevocable  bool
}

func (cs *Clerics) generateAppointment() *ClericsAppointment {
	cas := &ClericsAppointment{religion: cs.religion}

	return cas
}

func (cas *ClericsAppointment) Print() {
	fmt.Printf("ClericsAppointment (religion_name=%s):\n", cas.religion.Name)
}

type ClericsLimitations struct {
	religion *Religion

	AcceptableGender GenderAcceptance
	Marriage         Permission
}

func (cs *Clerics) generateLimitations() *ClericsLimitations {
	cls := &ClericsLimitations{religion: cs.religion}

	return cls
}

func (cls *ClericsLimitations) Print() {
	fmt.Printf("ClericsLimitations (religion_name=%s):\n", cls.religion.Name)
	fmt.Printf("Clerics AcceptableGender=%s, MarriageLimitations=%s\n", cls.AcceptableGender, cls.Marriage)
}

type clericsTrait struct {
	_religionMetadata *updateReligionMetadata
	Index             int
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
		for i, trait := range allTraits {
			if trait.Calc(cs.religion, trait, traits) {
				traits = append(traits, &clericsTrait{
					_religionMetadata: trait._religionMetadata,
					Index:             i,
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
			Name:              "NakedPriests",
			_religionMetadata: &updateReligionMetadata{},
			Calc: func(r *Religion, self *clericsTrait, _ []*clericsTrait) bool {
				baseCoef := pm.RandFloat64InRange(0.5, 0.9)

				return CalculateProbabilityFromReligionMetadata(baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "HeadOfFaith",
			_religionMetadata: &updateReligionMetadata{},
			Calc: func(r *Religion, self *clericsTrait, _ []*clericsTrait) bool {
				baseCoef := pm.RandFloat64InRange(0.5, 0.9)

				return CalculateProbabilityFromReligionMetadata(baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "SecretCirclesOfInitiation",
			_religionMetadata: &updateReligionMetadata{},
			Calc: func(r *Religion, self *clericsTrait, _ []*clericsTrait) bool {
				baseCoef := pm.RandFloat64InRange(0.5, 0.9)

				return CalculateProbabilityFromReligionMetadata(baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "MendicantPreachers",
			_religionMetadata: &updateReligionMetadata{},
			Calc: func(r *Religion, self *clericsTrait, _ []*clericsTrait) bool {
				baseCoef := pm.RandFloat64InRange(0.5, 0.9)

				return CalculateProbabilityFromReligionMetadata(baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Monasticism",
			_religionMetadata: &updateReligionMetadata{},
			Calc: func(r *Religion, self *clericsTrait, _ []*clericsTrait) bool {
				baseCoef := pm.RandFloat64InRange(0.5, 0.9)

				return CalculateProbabilityFromReligionMetadata(baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "VowOfPoverty",
			_religionMetadata: &updateReligionMetadata{},
			Calc: func(r *Religion, self *clericsTrait, _ []*clericsTrait) bool {
				baseCoef := pm.RandFloat64InRange(0.5, 0.9)

				return CalculateProbabilityFromReligionMetadata(baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

type clericsFunction struct {
	_religionMetadata *updateReligionMetadata
	Index             int
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
		for i, function := range allFunctions {
			if function.Calc(cs.religion, function, functions) {
				functions = append(functions, &clericsFunction{
					_religionMetadata: function._religionMetadata,
					Index:             i,
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
			Name:              "Control",
			_religionMetadata: &updateReligionMetadata{},
			Calc: func(r *Religion, self *clericsFunction, _ []*clericsFunction) bool {
				baseCoef := pm.RandFloat64InRange(0.5, 0.9)

				return CalculateProbabilityFromReligionMetadata(baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "AlmsAndPacification",
			_religionMetadata: &updateReligionMetadata{},
			Calc: func(r *Religion, self *clericsFunction, _ []*clericsFunction) bool {
				baseCoef := pm.RandFloat64InRange(0.5, 0.9)

				return CalculateProbabilityFromReligionMetadata(baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name:              "Recruitment",
			_religionMetadata: &updateReligionMetadata{},
			Calc: func(r *Religion, self *clericsFunction, _ []*clericsFunction) bool {
				baseCoef := pm.RandFloat64InRange(0.5, 0.9)

				return CalculateProbabilityFromReligionMetadata(baseCoef, r, *self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

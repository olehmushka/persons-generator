package religion

import (
	"fmt"

	g "persons_generator/engine/entities/gender"
	pm "persons_generator/engine/probability_machine"
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

func (cs *Clerics) generateHasClerics() (bool, error) {
	primaryProbability, err := pm.RandFloat64InRange(0.6, 0.85)
	if err != nil {
		return false, err
	}
	if cs.religion.metadata.IsLawful() {
		p, err := pm.RandFloat64InRange(0.05, 0.15)
		if err != nil {
			return false, err
		}
		primaryProbability += p
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProbability))
}

type ClericsAppointment struct {
	religion *Religion

	IsCivil     bool
	IsRevocable bool
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
	religion *Religion

	AcceptableGender g.Acceptance
	Marriage         Permission
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
	cls.Marriage = getPermissionByProbability(alwaysAllowed, mustBeApproved, disallowed)

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
			Name: "naked_priests",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.25,
				Chthonic:     0.25,
			},
			baseCoef: cs.religion.M.LowBaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				addCoef, err := pm.RandFloat64InRange(0.05, 0.15)
				if err != nil {
					return false, err
				}
				return CalculateProbabilityFromReligionMetadata(self.baseCoef-addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "head_of_faith",
			_religionMetadata: &religionMetadata{
				Plutocratic:     0.5,
				Lawful:          0.5,
				Authoritaristic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "internal_circles_of_initiation",
			_religionMetadata: &religionMetadata{
				Authoritaristic: 0.5,
				Collectivistic:  0.5,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				var addCoef float64
				for _, trait := range selectedTraits {
					if trait.Name == "head_of_faith" {
						c, err := pm.RandFloat64InRange(0.01, 0.1)
						if err != nil {
							return false, err
						}
						addCoef += c
					}
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef+addCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "mendicant_preachers",
			_religionMetadata: &religionMetadata{
				Ascetic: 0.75,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "monasticism",
			_religionMetadata: &religionMetadata{
				SexualStrictness: 0.5,
				Ascetic:          0.75,
				Authoritaristic:  0.75,
				Collectivistic:   0.1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "vow_of_poverty",
			_religionMetadata: &religionMetadata{
				Ascetic:         1,
				Authoritaristic: 0.25,
				Individualistic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "patronship",
			_religionMetadata: &religionMetadata{
				Plutocratic:     0.25,
				Authoritaristic: 0.5,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "warrior_monks",
			_religionMetadata: &religionMetadata{
				Aggressive:      1,
				Ascetic:         0.25,
				Authoritaristic: 0.25,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				var hasMonasticism bool
				for _, trait := range selectedTraits {
					if trait.Name == "monasticism" {
						hasMonasticism = true
					}
				}
				if !hasMonasticism {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "is_hereditary",
			_religionMetadata: &religionMetadata{
				Plutocratic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, selectedTraits []*trait) (bool, error) {
				if cs.Limitations.Marriage.IsDisallowed() {
					return false, nil
				}

				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

func (cs *Clerics) getAllClericsFunctions() []*trait {
	return []*trait{
		{
			Name: "control",
			_religionMetadata: &religionMetadata{
				Authoritaristic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "alms_and_pacification",
			_religionMetadata: &religionMetadata{
				Pacifistic: 0.75,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "heal",
			_religionMetadata: &religionMetadata{
				Altruistic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "recruitment",
			_religionMetadata: &religionMetadata{
				Collectivistic: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "teach",
			_religionMetadata: &religionMetadata{
				Lawful:      0.5,
				Educational: 1,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "oracle",
			_religionMetadata: &religionMetadata{
				Chthonic: 0.25,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "diviner",
			_religionMetadata: &religionMetadata{
				Chthonic: 0.25,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
		{
			Name: "druid",
			_religionMetadata: &religionMetadata{
				Naturalistic: 0.5,
				Chthonic:     0.25,
			},
			baseCoef: cs.religion.M.BaseCoef,
			Calc: func(r *Religion, self *trait, _ []*trait) (bool, error) {
				return CalculateProbabilityFromReligionMetadata(self.baseCoef, r, self._religionMetadata, CalcProbOpts{})
			},
		},
	}
}

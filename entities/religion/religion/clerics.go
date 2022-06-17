package religion

import "fmt"

type Clerics struct {
	religion *Religion
	attrs    *Attributes
}

func (as *Attributes) generateClerics() *Clerics {
	cs := &Clerics{religion: as.religion, attrs: as}

	return cs
}

func (cs *Clerics) Print() {
	fmt.Printf("Clerics (religion_name=%s):\n", cs.religion.Name)
}

/*

type Clerics struct {
	Function             ClericalFunction
	WhoCanBeCleric       GenderAcceptance
	Appointment          *ClericAppointment
	ClericMarriage       Permission
	NakedPriests         bool
	PrimacyOfHeadOfFaith bool
}

type ClericAppointment struct {
	IsCivil     bool
	IsRevocable bool
}

type ClericalFunction string

const (
	Control             ClericalFunction = "control"
	AlmsAndPacification ClericalFunction = "alms_and_pacification"
	Recruitment         ClericalFunction = "recruitment"
)

*/

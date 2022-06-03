package religion

type Clerics struct {
	Function       ClericalFunction
	WhoCanBeCleric GenderAcceptance
	Appointment    *ClericAppointment
	ClericMarriage Permission
	NakedPriests   bool
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

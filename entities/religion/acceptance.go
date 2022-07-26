package religion

import pm "persons_generator/probability_machine"

type Acceptance string

const (
	Accepted Acceptance = "accepted"
	Shunned  Acceptance = "shunned"
	Criminal Acceptance = "criminal"
)

func (a Acceptance) IsAccepted() bool {
	return a == Accepted
}

func (a Acceptance) IsShunned() bool {
	return a == Shunned
}

func (a Acceptance) IsCriminal() bool {
	return a == Criminal
}

func getAcceptanceByProbability(accepted, shunned, criminal float64) Acceptance {
	return Acceptance(pm.GetRandomFromSeveral(map[string]float64{
		string(Accepted): pm.PrepareProbability(accepted),
		string(Shunned):  pm.PrepareProbability(shunned),
		string(Criminal): pm.PrepareProbability(criminal),
	}))
}

package religion

import (
	"persons_generator/core/wrapped_error"
	pm "persons_generator/engine/probability_machine"
)

type Acceptance string

const (
	Accepted Acceptance = "accepted"
	Shunned  Acceptance = "shunned"
	Criminal Acceptance = "criminal"
)

func (a Acceptance) String() string {
	return string(a)
}

func (a Acceptance) IsAccepted() bool {
	return a == Accepted
}

func (a Acceptance) IsShunned() bool {
	return a == Shunned
}

func (a Acceptance) IsCriminal() bool {
	return a == Criminal
}

func (a Acceptance) Value() float64 {
	if a.IsAccepted() {
		return 1
	}
	if a.IsShunned() {
		return 0.5
	}
	if a.IsCriminal() {
		return 0
	}

	return 0
}

func getAcceptanceByProbability(accepted, shunned, criminal float64) (Acceptance, error) {
	acceptance, err := pm.GetRandomFromSeveral(map[string]float64{
		string(Accepted): pm.PrepareProbability(accepted),
		string(Shunned):  pm.PrepareProbability(shunned),
		string(Criminal): pm.PrepareProbability(criminal),
	})
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generace acceptance for religion")
	}

	return Acceptance(acceptance), nil
}

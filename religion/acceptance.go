package religion

import (
	"persons_generator/entities/religion"
	pm "persons_generator/probability_machine"
)

func geAcceptanceByProbability(accepted, shunned, criminal float64) religion.Acceptance {
	m := map[string]float64{
		string(religion.Accepted): pm.PrepareProbability(accepted),
		string(religion.Shunned):  pm.PrepareProbability(shunned),
		string(religion.Criminal): pm.PrepareProbability(criminal),
	}
	return religion.Acceptance(pm.GetRandomFromSeveral(m))
}

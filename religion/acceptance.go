package religion

import (
	"persons_generator/entities/religion"
	pm "persons_generator/probability-machine"
)

func geAcceptanceByProbability(accepted, shunned, criminal int) religion.Acceptance {
	m := map[string]int{
		string(religion.Accepted): pm.PrepareProbability(accepted),
		string(religion.Shunned):  pm.PrepareProbability(shunned),
		string(religion.Criminal): pm.PrepareProbability(criminal),
	}
	return religion.Acceptance(pm.GetRandomFromSeveral(m))
}

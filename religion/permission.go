package religion

import (
	"persons_generator/entities/religion"
	pm "persons_generator/probability_machine"
)

func getPermissionByProbability(alwaysAllowed, mustBeApproved, disallowed float64) religion.Permission {
	m := map[string]float64{
		string(religion.AlwaysAllowed):  pm.PrepareProbability(alwaysAllowed),
		string(religion.MustBeApproved): pm.PrepareProbability(mustBeApproved),
		string(religion.Disallowed):     pm.PrepareProbability(disallowed),
	}
	return religion.Permission(pm.GetRandomFromSeveral(m))
}

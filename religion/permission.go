package religion

import (
	"persons_generator/entities/religion"
	pm "persons_generator/probability-machine"
)

func getPermissionByProbability(alwaysAllowed, mustBeApproved, disallowed int) religion.Permission {
	m := map[string]int{
		string(religion.AlwaysAllowed):  pm.PrepareProbability(alwaysAllowed),
		string(religion.MustBeApproved): pm.PrepareProbability(mustBeApproved),
		string(religion.Disallowed):     pm.PrepareProbability(disallowed),
	}
	return religion.Permission(pm.GetRandomFromSeveral(m))
}

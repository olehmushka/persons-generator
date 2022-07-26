package religion

import pm "persons_generator/probability_machine"

type Permission string

const (
	AlwaysAllowed  Permission = "always_allowed"
	MustBeApproved Permission = "must_be_approved"
	Disallowed     Permission = "disallowed"
)

func getPermissionByProbability(alwaysAllowed, mustBeApproved, disallowed float64) Permission {
	return Permission(pm.GetRandomFromSeveral(map[string]float64{
		string(AlwaysAllowed):  pm.PrepareProbability(alwaysAllowed),
		string(MustBeApproved): pm.PrepareProbability(mustBeApproved),
		string(Disallowed):     pm.PrepareProbability(disallowed),
	}))
}

func (p Permission) IsAlwaysAllowed() bool {
	return p == AlwaysAllowed
}

func (p Permission) IsMustBeApproved() bool {
	return p == MustBeApproved
}

func (p Permission) IsDisallowed() bool {
	return p == Disallowed
}

package religion

import (
	"persons_generator/core/wrapped_error"
	pm "persons_generator/engine/probability_machine"
)

type Permission string

const (
	AlwaysAllowed  Permission = "always_allowed"
	MustBeApproved Permission = "must_be_approved"
	Disallowed     Permission = "disallowed"
)

func (p Permission) String() string {
	return string(p)
}

func getPermissionByProbability(alwaysAllowed, mustBeApproved, disallowed float64) (Permission, error) {
	p, err := pm.GetRandomFromSeveral(map[string]float64{
		string(AlwaysAllowed):  pm.PrepareProbability(alwaysAllowed),
		string(MustBeApproved): pm.PrepareProbability(mustBeApproved),
		string(Disallowed):     pm.PrepareProbability(disallowed),
	})
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generate religion permission")
	}

	return Permission(p), nil
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

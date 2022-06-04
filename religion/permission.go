package religion

import "persons_generator/entities/religion"

func getPermission(alwaysAllowed, mustBeApproved, disallowed int) religion.Permission {
	if disallowed >= alwaysAllowed && disallowed >= mustBeApproved {
		return religion.Disallowed
	}

	if mustBeApproved >= alwaysAllowed {
		return religion.MustBeApproved
	}

	return religion.AlwaysAllowed
}

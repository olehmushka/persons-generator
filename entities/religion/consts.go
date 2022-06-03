package religion

type Permission string

const (
	AlwaysAllowed  Permission = "always_allowed"
	MustBeApproved Permission = "must_be_approved"
	Disallowed     Permission = "disallowed"
)

type Acceptance string

const (
	Accepted Acceptance = "accepted"
	Shunned  Acceptance = "shunned"
	Criminal Acceptance = "criminal"
)

type GenderAcceptance string

const (
	OnlyMen     GenderAcceptance = "only_men"
	MenAndWomen GenderAcceptance = "men_and_women"
	OnlyWomen   GenderAcceptance = "only_women"
)

type Attitude string

const (
	Virtue          Attitude = "virtue"
	NeutralAttitude Attitude = "neutral"
	Sin             Attitude = "sin"
)

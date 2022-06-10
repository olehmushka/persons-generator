package religion

type Permission string

const (
	AlwaysAllowed  Permission = "always_allowed"
	MustBeApproved Permission = "must_be_approved"
	Disallowed     Permission = "disallowed"
)

type Acceptance string

func (a Acceptance) IsAccepted() bool {
	return a == Accepted
}

func (a Acceptance) IsShunned() bool {
	return a == Shunned
}

func (a Acceptance) IsCriminal() bool {
	return a == Criminal
}

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

func (a Attitude) IsVirtue() bool {
	return a == Virtue
}

func (a Attitude) IsNeutral() bool {
	return a == NeutralAttitude
}

func (a Attitude) IsSin() bool {
	return a == Sin
}

const (
	Virtue          Attitude = "virtue"
	NeutralAttitude Attitude = "neutral"
	Sin             Attitude = "sin"
)

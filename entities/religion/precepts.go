package religion

type Precepts struct {
	Marriage       *Marriage
	SinsAndVirtues *SinsAndVirtues
	Rituals        *Rituals
}

type Marriage struct {
	Kind          MarriageKind
	Divorce       Permission
	Bastardry     Bastardry
	Consanguinity Consanguinity
}

type MarriageKind string

const (
	Monogamy              MarriageKind = "monogamy"
	Polygamy              MarriageKind = "polygamy"
	ConsortsAndConcubines MarriageKind = "consorts_and_concubines"
)

type Bastardry string

const (
	NoBastards       Bastardry = "no_bastards"
	Legitimization   Bastardry = "legitimization"
	NoLegitimization Bastardry = "no_legitimization"
)

type Consanguinity string

const (
	CloseKinTaboo        Consanguinity = "close_kin_taboo"
	CousinMarriage       Consanguinity = "cousin_marriage"
	AvunculateMarriage   Consanguinity = "avunculate_marriage"
	UnrestrictedMarriage Consanguinity = "unrestricted_marriage"
)

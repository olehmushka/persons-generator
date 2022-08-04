package religion

import (
	"fmt"

	"persons_generator/entities/culture"
	pm "persons_generator/probability_machine"
)

type MarriageTradition struct {
	religion *Religion

	Kind          MarriageKind
	Bastardry     Bastardry
	Consanguinity Consanguinity
	Divorce       Permission
}

func (t *Theology) generateMarriageTradition(c *culture.Culture) *MarriageTradition {
	mt := &MarriageTradition{religion: t.religion}
	mt.Kind = mt.generateKind(c)
	mt.Bastardry = mt.generateBastardry()
	mt.Consanguinity = mt.generateConsanguinity()
	mt.Divorce = mt.generateDivorce()

	return mt
}

func (mt *MarriageTradition) Print() {
	fmt.Printf("MarriageTradition (religion_name=%s):\n", mt.religion.Name)
	fmt.Printf("MarriageKind=%s, Bastardry=%s, Consanguinity=%s, Divorce=%s\n", mt.Kind, mt.Bastardry, mt.Consanguinity, mt.Divorce)
}

type MarriageKind string

const (
	Monogamy              MarriageKind = "monogamy"
	Polygamy              MarriageKind = "polygamy"
	ConsortsAndConcubines MarriageKind = "consorts_and_concubines"
)

func (mt *MarriageTradition) generateKind(c *culture.Culture) MarriageKind {
	if mk := getMarriageKindByCulture(c); mk != "" {
		return mk
	}

	var (
		monogamy              = pm.RandFloat64InRange(0.25, 0.35)
		consortsAndConcubines = pm.RandFloat64InRange(0.15, 0.25)
		polygamy              = pm.RandFloat64InRange(0.05, 0.15)
	)

	switch {
	case mt.religion.GenderDominance.IsMaleDominate():
		consortsAndConcubines += pm.RandFloat64InRange(0.01, 0.05)
		polygamy += pm.RandFloat64InRange(0.01, 0.1)
	case mt.religion.GenderDominance.IsEquality():
		monogamy += pm.RandFloat64InRange(0.01, 0.1)
	case mt.religion.GenderDominance.IsFemaleDominate():
		monogamy += pm.RandFloat64InRange(0.1, 0.2)
	}

	if mt.religion.metadata.IsSexualActive() {
		consortsAndConcubines += pm.RandFloat64InRange(0.01, 0.1)
		polygamy += pm.RandFloat64InRange(0.01, 0.1)
	}
	if mt.religion.metadata.IsSexualStrictness() {
		monogamy += pm.RandFloat64InRange(0.1, 0.2)
	}
	if mt.religion.metadata.IsHedonistic() {
		consortsAndConcubines += pm.RandFloat64InRange(0.01, 0.05)
		polygamy += pm.RandFloat64InRange(0.01, 0.05)
	}

	return MarriageKind(pm.GetRandomFromSeveral(map[string]float64{
		string(Monogamy):              pm.PrepareProbability(monogamy),
		string(ConsortsAndConcubines): pm.PrepareProbability(consortsAndConcubines),
		string(Polygamy):              pm.PrepareProbability(polygamy),
	}))
}

func getMarriageKindByCulture(c *culture.Culture) MarriageKind {
	if c == nil {
		return ""
	}

	for _, t := range c.Traditions {
		if t == nil {
			continue
		}

		switch t.Name {
		case culture.MonogamousTradition.Name:
			return Polygamy
		case culture.ConcubinesTradition.Name:
			return ConsortsAndConcubines
		case culture.PolygamousTradition.Name:
			return Polygamy
		}
	}

	return ""
}

type Bastardry string

const (
	NoBastards       Bastardry = "no_bastards"
	Legitimization   Bastardry = "legitimization"
	NoLegitimization Bastardry = "no_legitimization"
)

func (mt *MarriageTradition) generateBastardry() Bastardry {
	var (
		noBastards       = pm.RandFloat64InRange(0.055, 0.155)
		legitimization   = pm.RandFloat64InRange(0.15, 0.25)
		noLegitimization = pm.RandFloat64InRange(0.06, 0.16)
	)

	switch {
	case mt.religion.GenderDominance.IsMaleDominate():
		noBastards += pm.RandFloat64InRange(0.03, 0.075)
		legitimization += pm.RandFloat64InRange(0.07, 0.13)
		noLegitimization += pm.RandFloat64InRange(0.03, 0.075)
	case mt.religion.GenderDominance.IsEquality():
		noBastards += pm.RandFloat64InRange(0.03, 0.075)
		legitimization += pm.RandFloat64InRange(0.07, 0.13)
	case mt.religion.GenderDominance.IsFemaleDominate():
		noBastards += pm.RandFloat64InRange(0.1, 0.2)
		legitimization += pm.RandFloat64InRange(0.07, 0.13)
	}

	if mt.religion.metadata.IsLawful() {
		legitimization += pm.RandFloat64InRange(0.05, 0.1)
	}
	if mt.religion.metadata.IsAltruistic() {
		noBastards += pm.RandFloat64InRange(0.04, 0.09)
		legitimization += pm.RandFloat64InRange(0.04, 0.09)
	}

	return Bastardry(pm.GetRandomFromSeveral(map[string]float64{
		string(NoBastards):       pm.PrepareProbability(noBastards),
		string(Legitimization):   pm.PrepareProbability(legitimization),
		string(NoLegitimization): pm.PrepareProbability(noLegitimization),
	}))
}

type Consanguinity string

const (
	CloseKinTaboo        Consanguinity = "close_kin_taboo"       // people can not marry family members
	CousinMarriage       Consanguinity = "cousin_marriage"       // people can marry cousins, but not other family members
	AvunculateMarriage   Consanguinity = "avunculate_marriage"   // people can not marry close family members
	UnrestrictedMarriage Consanguinity = "unrestricted_marriage" // people can marry all family members
)

func (mt *MarriageTradition) generateConsanguinity() Consanguinity {
	var (
		closeKinTaboo        = pm.RandFloat64InRange(0.15, 0.25)
		cousinMarriage       = pm.RandFloat64InRange(0.14, 0.26)
		avunculateMarriage   = pm.RandFloat64InRange(0.25, 0.35)
		unrestrictedMarriage = pm.RandFloat64InRange(0.05, 0.15)
	)

	switch {
	case mt.religion.GenderDominance.IsMaleDominate():
		closeKinTaboo += pm.RandFloat64InRange(0.025, 0.075)
		cousinMarriage += pm.RandFloat64InRange(0.05, 0.15)
		avunculateMarriage += pm.RandFloat64InRange(0.05, 0.1)
		unrestrictedMarriage += pm.RandFloat64InRange(0.015, 0.035)
	case mt.religion.GenderDominance.IsEquality():
		closeKinTaboo += pm.RandFloat64InRange(0.015, 0.035)
		cousinMarriage += pm.RandFloat64InRange(0.05, 0.1)
		avunculateMarriage += pm.RandFloat64InRange(0.05, 0.15)
		unrestrictedMarriage += pm.RandFloat64InRange(0.025, 0.075)
	case mt.religion.GenderDominance.IsFemaleDominate():
		closeKinTaboo += pm.RandFloat64InRange(0.05, 0.1)
		cousinMarriage += pm.RandFloat64InRange(0.05, 0.15)
		avunculateMarriage += pm.RandFloat64InRange(0.025, 0.075)
		unrestrictedMarriage += pm.RandFloat64InRange(0.015, 0.035)
	}

	if mt.religion.metadata.IsSexualActive() {
		unrestrictedMarriage += pm.RandFloat64InRange(0.01, 0.03)
	}
	if mt.religion.metadata.IsPacifistic() {
		avunculateMarriage += pm.RandFloat64InRange(0.01, 0.03)
	}
	if mt.religion.metadata.IsLawful() {
		closeKinTaboo += pm.RandFloat64InRange(0.01, 0.05)
		cousinMarriage += pm.RandFloat64InRange(0.01, 0.03)
	}
	if mt.religion.metadata.IsLiberal() {
		unrestrictedMarriage += pm.RandFloat64InRange(0.01, 0.03)
	}

	return Consanguinity(pm.GetRandomFromSeveral(map[string]float64{
		string(CloseKinTaboo):        pm.PrepareProbability(closeKinTaboo),
		string(CousinMarriage):       pm.PrepareProbability(cousinMarriage),
		string(AvunculateMarriage):   pm.PrepareProbability(avunculateMarriage),
		string(UnrestrictedMarriage): pm.PrepareProbability(unrestrictedMarriage),
	}))
}

func (mt *MarriageTradition) generateDivorce() Permission {
	var (
		alwaysAllowed  = pm.RandFloat64InRange(0.1, 0.2)
		mustBeApproved = pm.RandFloat64InRange(0.2, 0.3)
		disallowed     = pm.RandFloat64InRange(0.1, 0.2)
	)

	switch {
	case mt.religion.GenderDominance.IsMaleDominate():
		alwaysAllowed += pm.RandFloat64InRange(0.015, 0.035)
		mustBeApproved += pm.RandFloat64InRange(0.05, 0.15)
		disallowed += pm.RandFloat64InRange(0.1, 0.2)
	case mt.religion.GenderDominance.IsEquality():
		alwaysAllowed += pm.RandFloat64InRange(0.1, 0.2)
		mustBeApproved += pm.RandFloat64InRange(0.05, 0.15)
		disallowed += pm.RandFloat64InRange(0.015, 0.035)
	case mt.religion.GenderDominance.IsFemaleDominate():
		alwaysAllowed += pm.RandFloat64InRange(0.015, 0.035)
		mustBeApproved += pm.RandFloat64InRange(0.1, 0.2)
		disallowed += pm.RandFloat64InRange(0.05, 0.15)
	}

	if mt.religion.metadata.IsSexualActive() {
		alwaysAllowed += pm.RandFloat64InRange(0.01, 0.03)
	}
	if mt.religion.metadata.IsSexualStrictness() {
		disallowed += pm.RandFloat64InRange(0.01, 0.03)
	}
	if mt.religion.metadata.IsLawful() {
		mustBeApproved += pm.RandFloat64InRange(0.01, 0.03)
	}
	if mt.religion.metadata.IsAuthoritaristic() {
		disallowed += pm.RandFloat64InRange(0.01, 0.03)
	}
	if mt.religion.metadata.IsLiberal() {
		alwaysAllowed += pm.RandFloat64InRange(0.01, 0.03)
	}

	return getPermissionByProbability(alwaysAllowed, mustBeApproved, disallowed)
}

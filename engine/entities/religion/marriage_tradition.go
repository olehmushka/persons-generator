package religion

import (
	"fmt"

	"persons_generator/engine/entities/culture"
	pm "persons_generator/engine/probability_machine"
)

type MarriageTradition struct {
	religion *Religion

	Kind          MarriageKind
	Bastardry     Bastardry
	Consanguinity Consanguinity
	Divorce       Permission
}

func (t *Theology) generateMarriageTradition(c *culture.Culture) (*MarriageTradition, error) {
	mt := &MarriageTradition{religion: t.religion}
	k, err := mt.generateKind(c)
	if err != nil {
		return nil, err
	}
	mt.Kind = k
	b, err := mt.generateBastardry()
	if err != nil {
		return nil, err
	}
	mt.Bastardry = b
	con, err := mt.generateConsanguinity()
	if err != nil {
		return nil, err
	}
	mt.Consanguinity = con
	d, err := mt.generateDivorce()
	if err != nil {
		return nil, err
	}
	mt.Divorce = d

	return mt, nil
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

func (mt *MarriageTradition) generateKind(c *culture.Culture) (MarriageKind, error) {
	if mk := getMarriageKindByCulture(c); mk != "" {
		return mk, nil
	}

	monogamy, err := pm.RandFloat64InRange(0.25, 0.35)
	if err != nil {
		return "", err
	}
	consortsAndConcubines, err := pm.RandFloat64InRange(0.15, 0.25)
	if err != nil {
		return "", err
	}
	polygamy, err := pm.RandFloat64InRange(0.05, 0.15)
	if err != nil {
		return "", err
	}

	switch {
	case mt.religion.GenderDominance.IsMaleDominate():
		consortsAndConcubinesP, err := pm.RandFloat64InRange(0.01, 0.05)
		if err != nil {
			return "", err
		}
		consortsAndConcubines += consortsAndConcubinesP
		polygamyP, err := pm.RandFloat64InRange(0.01, 0.1)
		if err != nil {
			return "", err
		}
		polygamy += polygamyP
	case mt.religion.GenderDominance.IsEquality():
		monogamyP, err := pm.RandFloat64InRange(0.01, 0.1)
		if err != nil {
			return "", err
		}
		monogamy += monogamyP
	case mt.religion.GenderDominance.IsFemaleDominate():
		monogamyP, err := pm.RandFloat64InRange(0.1, 0.2)
		if err != nil {
			return "", err
		}
		monogamy += monogamyP
	}

	if mt.religion.metadata.IsSexualActive() {
		consortsAndConcubinesP, err := pm.RandFloat64InRange(0.01, 0.1)
		if err != nil {
			return "", err
		}
		consortsAndConcubines += consortsAndConcubinesP
		polygamyP, err := pm.RandFloat64InRange(0.01, 0.1)
		if err != nil {
			return "", err
		}
		polygamy += polygamyP
	}
	if mt.religion.metadata.IsSexualStrictness() {
		monogamyP, err := pm.RandFloat64InRange(0.1, 0.2)
		if err != nil {
			return "", err
		}
		monogamy += monogamyP
	}
	if mt.religion.metadata.IsHedonistic() {
		consortsAndConcubinesP, err := pm.RandFloat64InRange(0.01, 0.05)
		if err != nil {
			return "", err
		}
		consortsAndConcubines += consortsAndConcubinesP
		polygamyP, err := pm.RandFloat64InRange(0.01, 0.05)
		if err != nil {
			return "", err
		}
		polygamy += polygamyP
	}

	return MarriageKind(pm.GetRandomFromSeveral(map[string]float64{
		string(Monogamy):              pm.PrepareProbability(monogamy),
		string(ConsortsAndConcubines): pm.PrepareProbability(consortsAndConcubines),
		string(Polygamy):              pm.PrepareProbability(polygamy),
	})), nil
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

func (mt *MarriageTradition) generateBastardry() (Bastardry, error) {
	noBastards, err := pm.RandFloat64InRange(0.055, 0.155)
	if err != nil {
		return "", err
	}
	legitimization, err := pm.RandFloat64InRange(0.15, 0.25)
	if err != nil {
		return "", err
	}
	noLegitimization, err := pm.RandFloat64InRange(0.06, 0.16)
	if err != nil {
		return "", err
	}

	switch {
	case mt.religion.GenderDominance.IsMaleDominate():
		noBastardsP, err := pm.RandFloat64InRange(0.03, 0.075)
		if err != nil {
			return "", err
		}
		noBastards += noBastardsP
		legitimizationP, err := pm.RandFloat64InRange(0.07, 0.13)
		if err != nil {
			return "", err
		}
		legitimization += legitimizationP
		noLegitimizationP, err := pm.RandFloat64InRange(0.03, 0.075)
		if err != nil {
			return "", err
		}
		noLegitimization += noLegitimizationP
	case mt.religion.GenderDominance.IsEquality():
		noBastardsP, err := pm.RandFloat64InRange(0.03, 0.075)
		if err != nil {
			return "", err
		}
		noBastards += noBastardsP
		legitimizationP, err := pm.RandFloat64InRange(0.07, 0.13)
		if err != nil {
			return "", err
		}
		legitimization += legitimizationP
	case mt.religion.GenderDominance.IsFemaleDominate():
		noBastardsP, err := pm.RandFloat64InRange(0.1, 0.2)
		if err != nil {
			return "", err
		}
		noBastards += noBastardsP
		legitimizationP, err := pm.RandFloat64InRange(0.07, 0.13)
		if err != nil {
			return "", err
		}
		legitimization += legitimizationP
	}

	if mt.religion.metadata.IsLawful() {
		legitimizationP, err := pm.RandFloat64InRange(0.05, 0.1)
		if err != nil {
			return "", err
		}
		legitimization += legitimizationP
	}
	if mt.religion.metadata.IsAltruistic() {
		noBastardsP, err := pm.RandFloat64InRange(0.04, 0.09)
		if err != nil {
			return "", err
		}
		noBastards += noBastardsP
		legitimizationP, err := pm.RandFloat64InRange(0.04, 0.09)
		if err != nil {
			return "", err
		}
		legitimization += legitimizationP
	}

	return Bastardry(pm.GetRandomFromSeveral(map[string]float64{
		string(NoBastards):       pm.PrepareProbability(noBastards),
		string(Legitimization):   pm.PrepareProbability(legitimization),
		string(NoLegitimization): pm.PrepareProbability(noLegitimization),
	})), nil
}

type Consanguinity string

const (
	CloseKinTaboo        Consanguinity = "close_kin_taboo"       // people can not marry family members
	CousinMarriage       Consanguinity = "cousin_marriage"       // people can marry cousins, but not other family members
	AvunculateMarriage   Consanguinity = "avunculate_marriage"   // people can not marry close family members
	UnrestrictedMarriage Consanguinity = "unrestricted_marriage" // people can marry all family members
)

func (mt *MarriageTradition) generateConsanguinity() (Consanguinity, error) {
	closeKinTaboo, err := pm.RandFloat64InRange(0.15, 0.25)
	if err != nil {
		return "", err
	}
	cousinMarriage, err := pm.RandFloat64InRange(0.14, 0.26)
	if err != nil {
		return "", err
	}
	avunculateMarriage, err := pm.RandFloat64InRange(0.25, 0.35)
	if err != nil {
		return "", err
	}
	unrestrictedMarriage, err := pm.RandFloat64InRange(0.05, 0.15)
	if err != nil {
		return "", err
	}

	switch {
	case mt.religion.GenderDominance.IsMaleDominate():
		closeKinTabooP, err := pm.RandFloat64InRange(0.025, 0.075)
		if err != nil {
			return "", err
		}
		closeKinTaboo += closeKinTabooP
		cousinMarriageP, err := pm.RandFloat64InRange(0.05, 0.15)
		if err != nil {
			return "", err
		}
		cousinMarriage += cousinMarriageP
		avunculateMarriageP, err := pm.RandFloat64InRange(0.05, 0.1)
		if err != nil {
			return "", err
		}
		avunculateMarriage += avunculateMarriageP
		unrestrictedMarriageP, err := pm.RandFloat64InRange(0.015, 0.035)
		if err != nil {
			return "", err
		}
		unrestrictedMarriage += unrestrictedMarriageP
	case mt.religion.GenderDominance.IsEquality():
		closeKinTabooP, err := pm.RandFloat64InRange(0.015, 0.035)
		if err != nil {
			return "", err
		}
		closeKinTaboo += closeKinTabooP
		cousinMarriageP, err := pm.RandFloat64InRange(0.05, 0.1)
		if err != nil {
			return "", err
		}
		cousinMarriage += cousinMarriageP
		avunculateMarriageP, err := pm.RandFloat64InRange(0.05, 0.15)
		if err != nil {
			return "", err
		}
		avunculateMarriage += avunculateMarriageP
		unrestrictedMarriageP, err := pm.RandFloat64InRange(0.025, 0.075)
		if err != nil {
			return "", err
		}
		unrestrictedMarriage += unrestrictedMarriageP
	case mt.religion.GenderDominance.IsFemaleDominate():
		closeKinTabooP, err := pm.RandFloat64InRange(0.05, 0.1)
		if err != nil {
			return "", err
		}
		closeKinTaboo += closeKinTabooP
		cousinMarriageP, err := pm.RandFloat64InRange(0.05, 0.15)
		if err != nil {
			return "", err
		}
		cousinMarriage += cousinMarriageP
		avunculateMarriageP, err := pm.RandFloat64InRange(0.025, 0.075)
		if err != nil {
			return "", err
		}
		avunculateMarriage += avunculateMarriageP
		unrestrictedMarriageP, err := pm.RandFloat64InRange(0.015, 0.035)
		if err != nil {
			return "", err
		}
		unrestrictedMarriage += unrestrictedMarriageP
	}

	if mt.religion.metadata.IsSexualActive() {
		unrestrictedMarriageP, err := pm.RandFloat64InRange(0.01, 0.03)
		if err != nil {
			return "", err
		}
		unrestrictedMarriage += unrestrictedMarriageP
	}
	if mt.religion.metadata.IsPacifistic() {
		avunculateMarriageP, err := pm.RandFloat64InRange(0.01, 0.03)
		if err != nil {
			return "", err
		}
		avunculateMarriage += avunculateMarriageP
	}
	if mt.religion.metadata.IsLawful() {
		closeKinTabooP, err := pm.RandFloat64InRange(0.01, 0.05)
		if err != nil {
			return "", err
		}
		closeKinTaboo += closeKinTabooP
		cousinMarriageP, err := pm.RandFloat64InRange(0.01, 0.03)
		if err != nil {
			return "", err
		}
		cousinMarriage += cousinMarriageP
	}
	if mt.religion.metadata.IsLiberal() {
		unrestrictedMarriageP, err := pm.RandFloat64InRange(0.01, 0.03)
		if err != nil {
			return "", err
		}
		unrestrictedMarriage += unrestrictedMarriageP
	}

	return Consanguinity(pm.GetRandomFromSeveral(map[string]float64{
		string(CloseKinTaboo):        pm.PrepareProbability(closeKinTaboo),
		string(CousinMarriage):       pm.PrepareProbability(cousinMarriage),
		string(AvunculateMarriage):   pm.PrepareProbability(avunculateMarriage),
		string(UnrestrictedMarriage): pm.PrepareProbability(unrestrictedMarriage),
	})), nil
}

func (mt *MarriageTradition) generateDivorce() (Permission, error) {
	alwaysAllowed, err := pm.RandFloat64InRange(0.1, 0.2)
	if err != nil {
		return "", err
	}
	mustBeApproved, err := pm.RandFloat64InRange(0.2, 0.3)
	if err != nil {
		return "", err
	}
	disallowed, err := pm.RandFloat64InRange(0.1, 0.2)
	if err != nil {
		return "", err
	}

	switch {
	case mt.religion.GenderDominance.IsMaleDominate():
		alwaysAllowedP, err := pm.RandFloat64InRange(0.015, 0.035)
		if err != nil {
			return "", err
		}
		alwaysAllowed += alwaysAllowedP
		mustBeApprovedP, err := pm.RandFloat64InRange(0.05, 0.15)
		if err != nil {
			return "", err
		}
		mustBeApproved += mustBeApprovedP
		disallowedP, err := pm.RandFloat64InRange(0.1, 0.2)
		if err != nil {
			return "", err
		}
		disallowed += disallowedP
	case mt.religion.GenderDominance.IsEquality():
		alwaysAllowedP, err := pm.RandFloat64InRange(0.1, 0.2)
		if err != nil {
			return "", err
		}
		alwaysAllowed += alwaysAllowedP
		mustBeApprovedP, err := pm.RandFloat64InRange(0.05, 0.15)
		if err != nil {
			return "", err
		}
		mustBeApproved += mustBeApprovedP
		disallowedP, err := pm.RandFloat64InRange(0.015, 0.035)
		if err != nil {
			return "", err
		}
		disallowed += disallowedP
	case mt.religion.GenderDominance.IsFemaleDominate():
		alwaysAllowedP, err := pm.RandFloat64InRange(0.015, 0.035)
		if err != nil {
			return "", err
		}
		alwaysAllowed += alwaysAllowedP
		mustBeApprovedP, err := pm.RandFloat64InRange(0.1, 0.2)
		if err != nil {
			return "", err
		}
		mustBeApproved += mustBeApprovedP
		disallowedP, err := pm.RandFloat64InRange(0.05, 0.15)
		if err != nil {
			return "", err
		}
		disallowed += disallowedP
	}

	if mt.religion.metadata.IsSexualActive() {
		alwaysAllowedP, err := pm.RandFloat64InRange(0.01, 0.03)
		if err != nil {
			return "", err
		}
		alwaysAllowed += alwaysAllowedP
	}
	if mt.religion.metadata.IsSexualStrictness() {
		disallowedP, err := pm.RandFloat64InRange(0.01, 0.03)
		if err != nil {
			return "", err
		}
		disallowed += disallowedP
	}
	if mt.religion.metadata.IsLawful() {
		mustBeApprovedP, err := pm.RandFloat64InRange(0.01, 0.03)
		if err != nil {
			return "", err
		}
		mustBeApproved += mustBeApprovedP
	}
	if mt.religion.metadata.IsAuthoritaristic() {
		disallowedP, err := pm.RandFloat64InRange(0.01, 0.03)
		if err != nil {
			return "", err
		}
		disallowed += disallowedP
	}
	if mt.religion.metadata.IsLiberal() {
		alwaysAllowedP, err := pm.RandFloat64InRange(0.01, 0.03)
		if err != nil {
			return "", err
		}
		alwaysAllowed += alwaysAllowedP
	}

	return getPermissionByProbability(alwaysAllowed, mustBeApproved, disallowed), nil
}

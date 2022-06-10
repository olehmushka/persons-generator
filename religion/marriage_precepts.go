package religion

import (
	"fmt"

	"persons_generator/entities"
	"persons_generator/entities/religion"
	pm "persons_generator/probability-machine"
)

func generateMarriagePercepts(r *entities.Religion) {
	fmt.Println("[generateMarriagePercepts] started")
	defer fmt.Println("[generateMarriagePercepts] finished")

	r.Theology.Precepts.Marriage = &religion.Marriage{}
	r.Theology.Precepts.Marriage.Kind = getMarriageKind(r)
	r.Theology.Precepts.Marriage.Divorce = getDivorce(r)
	r.Theology.Precepts.Marriage.Bastardry = getBastardry(r)
	r.Theology.Precepts.Marriage.Consanguinity = getConsanguinity(r)
}

func getMarriageKind(r *entities.Religion) religion.MarriageKind {
	var (
		monogamy              = 0.3
		consortsAndConcubines = 0.2
		polygamy              = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		monogamy += 0.1
		consortsAndConcubines += 0.05
		polygamy += 0.1
	case r.Doctrines.Base.Polytheism:
		monogamy += 0.1
		consortsAndConcubines += 0.05
		polygamy += 0.1
	case r.Doctrines.Base.DeityDualism:
		monogamy += 0.1
		consortsAndConcubines += 0.1
	case r.Doctrines.Base.Deism:
		monogamy += 0.1
		consortsAndConcubines += 0.1
		polygamy += 0.1
	case r.Doctrines.Base.Henothism:
		monogamy += 0.1
		polygamy += 0.1
	case r.Doctrines.Base.Monolatry:
		monogamy += 0.1
		polygamy += 0.1
	case r.Doctrines.Base.Omnism:
		polygamy += 0.1
		consortsAndConcubines += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		consortsAndConcubines += 0.03
		polygamy += 0.05
	case r.Doctrines.Gender.Equality:
		monogamy += 0.05
	case r.Doctrines.Gender.FemaleDominance:
		monogamy += 0.15
	}

	if r.Doctrines.Asceticism {
		monogamy += 0.085
	}
	if r.Doctrines.Monasticism {
		monogamy += 0.01
	}
	if r.Doctrines.Polyamory {
		polygamy += 0.5
		consortsAndConcubines += 0.1
	}
	if r.Doctrines.ReligiousLaw {
		monogamy += 0.1
		polygamy += 0.1
		consortsAndConcubines += 0.1
	}
	if r.Doctrines.Raider {
		polygamy += 0.5
	}
	if r.Doctrines.Hedonism {
		polygamy += 0.5
		consortsAndConcubines += 0.5
	}

	m := map[string]float64{
		string(religion.Monogamy):              pm.PrepareProbability(monogamy),
		string(religion.ConsortsAndConcubines): pm.PrepareProbability(consortsAndConcubines),
		string(religion.Polygamy):              pm.PrepareProbability(polygamy),
	}
	return religion.MarriageKind(pm.GetRandomFromSeveral(m))
}

func getDivorce(r *entities.Religion) religion.Permission {
	var (
		alwaysAllowed  = 0.1
		mustBeApproved = 0.25
		disallowed     = 0.1
	)

	switch {
	case r.Doctrines.Base.Monotheism:
		alwaysAllowed += 0.1
		mustBeApproved += 0.18
		disallowed += 0.15
	case r.Doctrines.Base.Polytheism:
		alwaysAllowed += 0.1
		disallowed += 0.1
	case r.Doctrines.Base.DeityDualism:
		mustBeApproved += 0.1
		disallowed += 0.1
	case r.Doctrines.Base.Deism:
		alwaysAllowed += 0.01
	case r.Doctrines.Base.Henothism:
		alwaysAllowed += 0.1
		disallowed += 0.1
	case r.Doctrines.Base.Monolatry:
		alwaysAllowed += 0.1
		disallowed += 0.1
	case r.Doctrines.Base.Omnism:
		alwaysAllowed += 0.1
		mustBeApproved += 0.1
		disallowed += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		mustBeApproved += 0.1
		disallowed += 0.15
	case r.Doctrines.Gender.Equality:
		alwaysAllowed += 0.05
		mustBeApproved += 0.1
	case r.Doctrines.Gender.FemaleDominance:
		alwaysAllowed += 0.15
		mustBeApproved += 0.05
	}

	if r.Doctrines.FullTolerance {
		alwaysAllowed += 0.1
	}
	if r.Doctrines.Asceticism {
		mustBeApproved += 0.2
		disallowed += 0.3
	}
	if r.Doctrines.Legalism {
		mustBeApproved += 0.2
	}
	if r.Doctrines.Polyamory {
		alwaysAllowed += 0.5
		mustBeApproved += 0.1
	}
	if r.Doctrines.ReligiousLaw {
		mustBeApproved += 0.2
	}
	if r.Doctrines.Pacifism {
		alwaysAllowed += 0.1
		mustBeApproved += 0.1
	}
	if r.Doctrines.Pacifism {
		alwaysAllowed += 0.1
		mustBeApproved += 0.1
	}
	if r.Doctrines.Raider {
		alwaysAllowed += 0.01
		mustBeApproved += 0.1
		disallowed += 0.1
	}
	if r.Doctrines.Proselytizer {
		disallowed += 0.1
	}
	if r.Doctrines.Hedonism {
		alwaysAllowed += 0.3
	}

	return getPermissionByProbability(alwaysAllowed, mustBeApproved, disallowed)
}

func getBastardry(r *entities.Religion) religion.Bastardry {
	var (
		noBastards       = 0.105
		legitimization   = 0.2
		noLegitimization = 0.11
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		noBastards += 0.12
		legitimization += 0.18
		noLegitimization += 0.15
	case r.Doctrines.Base.Polytheism:
		noBastards += 0.13
		legitimization += 0.18
	case r.Doctrines.Base.DeityDualism:
		legitimization += 0.15
	case r.Doctrines.Base.Deism:
		noBastards += 0.17
		legitimization += 0.19
	case r.Doctrines.Base.Henothism:
		noBastards += 0.15
		legitimization += 0.16
	case r.Doctrines.Base.Monolatry:
		noBastards += 0.12
		legitimization += 0.17
	case r.Doctrines.Base.Omnism:
		noBastards += 0.19
		legitimization += 0.17
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		legitimization += 0.11
		noBastards += 0.051
		noLegitimization += 0.055
	case r.Doctrines.Gender.Equality:
		noBastards += 0.05
		legitimization += 0.1
	case r.Doctrines.Gender.FemaleDominance:
		noBastards += 0.159
		legitimization += 0.109
	}

	if r.Doctrines.FullTolerance {
		noBastards += 0.3
		legitimization += 0.05
	}
	if r.Doctrines.Asceticism {
		noLegitimization += 0.3
		legitimization += 0.05
	}
	if r.Doctrines.Legalism {
		legitimization += 0.2
	}
	if r.Doctrines.Polyamory {
		noBastards += 0.6
		legitimization += 0.01
	}
	if r.Doctrines.ReligiousLaw {
		legitimization += 0.21
		noLegitimization += 0.29
	}
	if r.Doctrines.Pacifism {
		legitimization += 0.1
		noBastards += 0.07
	}
	if r.Doctrines.SacredChildbirth {
		noBastards += 0.5
		legitimization += 0.1
	}
	if r.Doctrines.Raider {
		noBastards += 0.15
		legitimization += 0.19
		noLegitimization += 0.13
	}

	m := map[string]float64{
		string(religion.NoBastards):       pm.PrepareProbability(noBastards),
		string(religion.Legitimization):   pm.PrepareProbability(legitimization),
		string(religion.NoLegitimization): pm.PrepareProbability(noLegitimization),
	}
	return religion.Bastardry(pm.GetRandomFromSeveral(m))
}

func getConsanguinity(r *entities.Religion) religion.Consanguinity {
	var (
		closeKinTaboo        = 0.2
		cousinMarriage       = 0.2
		avunculateMarriage   = 0.3
		unrestrictedMarriage = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		avunculateMarriage += 0.15
		cousinMarriage += 0.1
		closeKinTaboo += 0.04
		unrestrictedMarriage += 0.01
	case r.Doctrines.Base.Polytheism:
		avunculateMarriage += 0.17
		cousinMarriage += 0.12
		closeKinTaboo += 0.02
		unrestrictedMarriage += 0.01
	case r.Doctrines.Base.DeityDualism:
		avunculateMarriage += 0.17
		cousinMarriage += 0.12
		closeKinTaboo += 0.05
		unrestrictedMarriage += 0.015
	case r.Doctrines.Base.Deism:
		avunculateMarriage += 0.13
		cousinMarriage += 0.12
		closeKinTaboo += 0.07
		unrestrictedMarriage += 0.02
	case r.Doctrines.Base.Henothism:
		avunculateMarriage += 0.12
		cousinMarriage += 0.12
		closeKinTaboo += 0.05
		unrestrictedMarriage += 0.01
	case r.Doctrines.Base.Monolatry:
		avunculateMarriage += 0.15
		cousinMarriage += 0.12
		closeKinTaboo += 0.045
		unrestrictedMarriage += 0.015
	case r.Doctrines.Base.Omnism:
		avunculateMarriage += 0.15
		cousinMarriage += 0.14
		closeKinTaboo += 0.13
		unrestrictedMarriage += 0.03
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		avunculateMarriage += 0.053
		cousinMarriage += 0.044
		closeKinTaboo += 0.045
		unrestrictedMarriage += 0.06
	case r.Doctrines.Gender.Equality:
		avunculateMarriage += 0.15
		cousinMarriage += 0.094
		closeKinTaboo += 0.073
		unrestrictedMarriage += 0.039
	case r.Doctrines.Gender.FemaleDominance:
		avunculateMarriage += 0.085
		cousinMarriage += 0.0999
		closeKinTaboo += 0.045
		unrestrictedMarriage += 0.055
	}

	if r.Doctrines.FullTolerance {
		unrestrictedMarriage += 0.3
	}
	if r.Doctrines.Asceticism {
		avunculateMarriage += 0.2
		cousinMarriage += 0.02
	}
	if r.Doctrines.Polyamory {
		cousinMarriage += 1
		closeKinTaboo += 2
		unrestrictedMarriage += 0.4
	}
	if r.Doctrines.ReligiousLaw {
		avunculateMarriage += 0.35
		cousinMarriage += 0.15
	}
	if r.Doctrines.SacredChildbirth {
		avunculateMarriage += 0.2
	}
	if r.Doctrines.Hedonism {
		closeKinTaboo += 0.35
		unrestrictedMarriage += 0.4
	}

	m := map[string]float64{
		string(religion.CloseKinTaboo):        pm.PrepareProbability(closeKinTaboo),
		string(religion.CousinMarriage):       pm.PrepareProbability(cousinMarriage),
		string(religion.AvunculateMarriage):   pm.PrepareProbability(avunculateMarriage),
		string(religion.UnrestrictedMarriage): pm.PrepareProbability(unrestrictedMarriage),
	}
	return religion.Consanguinity(pm.GetRandomFromSeveral(m))
}

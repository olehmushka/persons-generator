package religion

import (
	"persons_generator/entities"
	"persons_generator/entities/religion"
	pm "persons_generator/probability-machine"
)

func generateMarriagePercepts(r *entities.Religion) {
	r.Theology.Precepts.Marriage = &religion.Marriage{}
	r.Theology.Precepts.Marriage.Kind = getMarriageKind(r)
	r.Theology.Precepts.Marriage.Divorce = getDivorce(r)
	r.Theology.Precepts.Marriage.Bastardry = getBastardry(r)
	r.Theology.Precepts.Marriage.Consanguinity = getConsanguinity(r)
}

func getMarriageKind(r *entities.Religion) religion.MarriageKind {
	var (
		monogamy              = 30
		consortsAndConcubines = 20
		polygamy              = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		monogamy += 10
		polygamy += 10
	case r.Doctrines.Base.Polytheism:
		monogamy += 10
		polygamy += 10
	case r.Doctrines.Base.DeityDualism:
		monogamy += 10
		consortsAndConcubines += 10
	case r.Doctrines.Base.Deism:
		monogamy += 10
		consortsAndConcubines += 10
		polygamy += 10
	case r.Doctrines.Base.Henothism:
		monogamy += 10
		polygamy += 10
	case r.Doctrines.Base.Monolatry:
		monogamy += 10
		polygamy += 10
	case r.Doctrines.Base.Omnism:
		polygamy += 10
		consortsAndConcubines += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		polygamy += 5
	case r.Doctrines.Gender.Equality:
		monogamy += 5
	case r.Doctrines.Gender.FemaleDominance:
		monogamy += 10
	}

	if r.Doctrines.Asceticism {
		monogamy += 10
	}
	if r.Doctrines.Monasticism {
		monogamy += 10
	}
	if r.Doctrines.Polyamory {
		polygamy += 50
	}
	if r.Doctrines.ReligiousLaw {
		monogamy += 10
		polygamy += 10
		consortsAndConcubines += 10
	}
	if r.Doctrines.Raider {
		polygamy += 50
	}
	if r.Doctrines.Hedonism {
		polygamy += 50
		consortsAndConcubines += 50
	}

	m := map[string]int{
		string(religion.Monogamy):              pm.PrepareProbability(monogamy),
		string(religion.ConsortsAndConcubines): pm.PrepareProbability(consortsAndConcubines),
		string(religion.Polygamy):              pm.PrepareProbability(polygamy),
	}
	return religion.MarriageKind(pm.GetRandomFromSeveral(m))
}

func getDivorce(r *entities.Religion) religion.Permission {
	var (
		alwaysAllowed  = 10
		mustBeApproved = 20
		disallowed     = 10
	)

	switch {
	case r.Doctrines.Base.Monotheism:
		alwaysAllowed += 10
		mustBeApproved += 10
		disallowed += 10
	case r.Doctrines.Base.Polytheism:
		alwaysAllowed += 10
		disallowed += 10
	case r.Doctrines.Base.DeityDualism:
		mustBeApproved += 10
		disallowed += 10
	case r.Doctrines.Base.Deism:
		alwaysAllowed++
	case r.Doctrines.Base.Henothism:
		alwaysAllowed += 10
		disallowed += 10
	case r.Doctrines.Base.Monolatry:
		alwaysAllowed += 10
		disallowed += 10
	case r.Doctrines.Base.Omnism:
		alwaysAllowed += 10
		mustBeApproved += 10
		disallowed += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		mustBeApproved += 10
		disallowed += 15
	case r.Doctrines.Gender.Equality:
		alwaysAllowed += 5
		mustBeApproved += 10
	case r.Doctrines.Gender.FemaleDominance:
		alwaysAllowed += 15
		mustBeApproved += 5
	}

	if r.Doctrines.FullTolerance {
		alwaysAllowed += 10
	}
	if r.Doctrines.Asceticism {
		mustBeApproved += 20
		disallowed += 30
	}
	if r.Doctrines.Legalism {
		mustBeApproved += 20
	}
	if r.Doctrines.Polyamory {
		alwaysAllowed += 50
	}
	if r.Doctrines.ReligiousLaw {
		mustBeApproved += 20
	}
	if r.Doctrines.Pacifism {
		alwaysAllowed += 10
		mustBeApproved += 10
	}
	if r.Doctrines.Pacifism {
		alwaysAllowed += 10
		mustBeApproved += 10
	}
	if r.Doctrines.Raider {
		mustBeApproved += 10
		disallowed += 10
	}
	if r.Doctrines.Proselytizer {
		disallowed += 10
	}
	if r.Doctrines.Hedonism {
		alwaysAllowed += 30
	}

	return getPermissionByProbability(alwaysAllowed, mustBeApproved, disallowed)
}

func getBastardry(r *entities.Religion) religion.Bastardry {
	var (
		noBastards       = 10
		legitimization   = 20
		noLegitimization = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		noBastards += 10
		legitimization += 10
		noLegitimization += 10
	case r.Doctrines.Base.Polytheism:
		noBastards += 10
		legitimization += 10
	case r.Doctrines.Base.DeityDualism:
		legitimization += 10
	case r.Doctrines.Base.Deism:
		noBastards += 10
		legitimization += 10
	case r.Doctrines.Base.Henothism:
		noBastards += 10
		legitimization += 10
	case r.Doctrines.Base.Monolatry:
		noBastards += 10
		legitimization += 10
	case r.Doctrines.Base.Omnism:
		noBastards += 10
		legitimization += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		legitimization += 10
		noBastards += 5
	case r.Doctrines.Gender.Equality:
		noBastards += 5
		legitimization += 10
	case r.Doctrines.Gender.FemaleDominance:
		noBastards += 15
		legitimization += 10
	}

	if r.Doctrines.FullTolerance {
		noBastards += 30
	}
	if r.Doctrines.Asceticism {
		noLegitimization += 30
	}
	if r.Doctrines.Legalism {
		legitimization += 20
	}
	if r.Doctrines.Polyamory {
		noBastards += 60
	}
	if r.Doctrines.ReligiousLaw {
		legitimization += 20
		noLegitimization += 20
	}
	if r.Doctrines.Pacifism {
		legitimization += 10
	}
	if r.Doctrines.SacredChildbirth {
		noBastards += 50
	}
	if r.Doctrines.Raider {
		noBastards += 10
		legitimization += 10
		noLegitimization += 10
	}

	m := map[string]int{
		string(religion.NoBastards):       pm.PrepareProbability(noBastards),
		string(religion.Legitimization):   pm.PrepareProbability(legitimization),
		string(religion.NoLegitimization): pm.PrepareProbability(noLegitimization),
	}
	return religion.Bastardry(pm.GetRandomFromSeveral(m))
}

func getConsanguinity(r *entities.Religion) religion.Consanguinity {
	var (
		closeKinTaboo        = 20
		cousinMarriage       = 20
		avunculateMarriage   = 30
		unrestrictedMarriage = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		avunculateMarriage += 15
		cousinMarriage += 10
	case r.Doctrines.Base.Polytheism:
		avunculateMarriage += 10
		cousinMarriage += 10
	case r.Doctrines.Base.DeityDualism:
		avunculateMarriage += 10
		cousinMarriage += 10
	case r.Doctrines.Base.Deism:
		avunculateMarriage += 10
		cousinMarriage += 10
	case r.Doctrines.Base.Henothism:
		avunculateMarriage += 10
		cousinMarriage += 10
	case r.Doctrines.Base.Monolatry:
		avunculateMarriage += 10
		cousinMarriage += 10
	case r.Doctrines.Base.Omnism:
		closeKinTaboo += 10
		avunculateMarriage += 10
		cousinMarriage += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		unrestrictedMarriage += 6
		closeKinTaboo += 5
		cousinMarriage += 4
		avunculateMarriage += 3
	case r.Doctrines.Gender.Equality:
		avunculateMarriage += 5
	case r.Doctrines.Gender.FemaleDominance:
		unrestrictedMarriage += 4
		closeKinTaboo += 3
		cousinMarriage += 6
		avunculateMarriage += 5
	}

	if r.Doctrines.FullTolerance {
		unrestrictedMarriage += 30
	}
	if r.Doctrines.Asceticism {
		avunculateMarriage += 20
	}
	if r.Doctrines.Polyamory {
		unrestrictedMarriage += 40
	}
	if r.Doctrines.ReligiousLaw {
		avunculateMarriage += 20
		cousinMarriage += 10
	}
	if r.Doctrines.SacredChildbirth {
		avunculateMarriage += 20
	}
	if r.Doctrines.Hedonism {
		unrestrictedMarriage += 40
	}

	m := map[string]int{
		string(religion.CloseKinTaboo):        pm.PrepareProbability(closeKinTaboo),
		string(religion.CousinMarriage):       pm.PrepareProbability(cousinMarriage),
		string(religion.AvunculateMarriage):   pm.PrepareProbability(avunculateMarriage),
		string(religion.UnrestrictedMarriage): pm.PrepareProbability(unrestrictedMarriage),
	}
	return religion.Consanguinity(pm.GetRandomFromSeveral(m))
}

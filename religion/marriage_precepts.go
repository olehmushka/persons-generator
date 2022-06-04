package religion

import (
	"persons_generator/entities"
	"persons_generator/entities/religion"
)

func generateMarriagePercepts(r *entities.Religion) {
	r.Theology.Precepts.Marriage = &religion.Marriage{}
	r.Theology.Precepts.Marriage.Kind = getMarriageKind(r)
	r.Theology.Precepts.Marriage.Divorce = getDivorce(r)
	r.Theology.Precepts.Marriage.Bastardry = getBastardry(r)
	r.Theology.Precepts.Marriage.Consanguinity = getConsanguinity(r)
}

func getMarriageKind(r *entities.Religion) religion.MarriageKind {
	var monogamy, polygamy, consortsAndConcubines int
	switch {
	case r.Doctrines.Main.Monotheism:
	case r.Doctrines.Main.Polytheism:
	case r.Doctrines.Main.DeityDualism:
	case r.Doctrines.Main.Deism:
	case r.Doctrines.Main.Henothism:
	case r.Doctrines.Main.Monolatry:
	case r.Doctrines.Main.Omnism:
		polygamy++
		consortsAndConcubines++
	}

	if r.Doctrines.Asceticism {
		monogamy++
	}
	if r.Doctrines.Monasticism {
		monogamy++
	}
	if r.Doctrines.Polyamory {
		polygamy++
	}
	if r.Doctrines.ReligiousLaw {
		monogamy++
		polygamy++
		consortsAndConcubines++
	}
	if r.Doctrines.Raider {
		polygamy++
	}
	if r.Doctrines.Hedonism {
		polygamy++
		consortsAndConcubines++
	}

	if monogamy >= polygamy && monogamy > consortsAndConcubines {
		return religion.Monogamy
	}
	if polygamy >= consortsAndConcubines {
		return religion.Polygamy
	}

	return religion.ConsortsAndConcubines
}

func getDivorce(r *entities.Religion) religion.Permission {
	var alwaysAllowed, mustBeApproved, disallowed int
	mustBeApproved++

	switch {
	case r.Doctrines.Main.Monotheism:
		alwaysAllowed++
		mustBeApproved++
		disallowed++
	case r.Doctrines.Main.Polytheism:
		alwaysAllowed++
		disallowed++
	case r.Doctrines.Main.DeityDualism:
		mustBeApproved++
		disallowed++
	case r.Doctrines.Main.Deism:
		alwaysAllowed++
	case r.Doctrines.Main.Henothism:
		alwaysAllowed++
		disallowed++
	case r.Doctrines.Main.Monolatry:
		alwaysAllowed++
		disallowed++
	case r.Doctrines.Main.Omnism:
		alwaysAllowed++
		mustBeApproved++
		disallowed++
	}

	if r.Doctrines.FullTolerance {
		alwaysAllowed++
	}
	if r.Doctrines.Asceticism {
		mustBeApproved++
		disallowed++
	}
	if r.Doctrines.Legalism {
		mustBeApproved++
	}
	if r.Doctrines.Polyamory {
		alwaysAllowed += 2
	}
	if r.Doctrines.ReligiousLaw {
		mustBeApproved++
	}
	if r.Doctrines.Pacifism {
		alwaysAllowed++
		mustBeApproved++
	}
	if r.Doctrines.Pacifism {
		alwaysAllowed++
		mustBeApproved++
	}
	if r.Doctrines.Raider {
		mustBeApproved++
		disallowed++
	}
	if r.Doctrines.Proselytizer {
		disallowed++
	}
	if r.Doctrines.Hedonism {
		alwaysAllowed++
	}

	return getPermission(alwaysAllowed, mustBeApproved, disallowed)
}

func getBastardry(r *entities.Religion) religion.Bastardry {
	var noBastards, legitimization, noLegitimization int
	switch {
	case r.Doctrines.Main.Monotheism:
		noBastards++
		legitimization++
		noLegitimization++
	case r.Doctrines.Main.Polytheism:
		noBastards++
		legitimization++
	case r.Doctrines.Main.DeityDualism:
		legitimization++
	case r.Doctrines.Main.Deism:
		noBastards++
		legitimization++
	case r.Doctrines.Main.Henothism:
		noBastards++
		legitimization++
	case r.Doctrines.Main.Monolatry:
		noBastards++
		legitimization++
	case r.Doctrines.Main.Omnism:
		noBastards++
		legitimization++
	}

	if r.Doctrines.FullTolerance {
		noBastards++
	}
	if r.Doctrines.Asceticism {
		noLegitimization++
	}
	if r.Doctrines.Legalism {
		legitimization++
	}
	if r.Doctrines.Polyamory {
		noBastards += 5
	}
	if r.Doctrines.ReligiousLaw {
		legitimization++
		noLegitimization++
	}
	if r.Doctrines.Pacifism {
		legitimization++
	}
	if r.Doctrines.SacredChildbirth {
		noBastards++
	}
	if r.Doctrines.Raider {
		noBastards++
		legitimization++
		noLegitimization++
	}

	if legitimization >= noBastards && legitimization > noLegitimization {
		return religion.Legitimization
	}
	if noLegitimization > noBastards {
		return religion.NoLegitimization
	}

	return religion.NoBastards
}

func getConsanguinity(r *entities.Religion) religion.Consanguinity {
	var closeKinTaboo, cousinMarriage, avunculateMarriage, unrestrictedMarriage int
	avunculateMarriage++
	switch {
	case r.Doctrines.Main.Monotheism:
		avunculateMarriage++
		cousinMarriage++
	case r.Doctrines.Main.Polytheism:
		avunculateMarriage++
		cousinMarriage++
	case r.Doctrines.Main.DeityDualism:
		avunculateMarriage++
		cousinMarriage++
	case r.Doctrines.Main.Deism:
		avunculateMarriage++
		cousinMarriage++
	case r.Doctrines.Main.Henothism:
		avunculateMarriage++
		cousinMarriage++
	case r.Doctrines.Main.Monolatry:
		avunculateMarriage++
		cousinMarriage++
	case r.Doctrines.Main.Omnism:
		closeKinTaboo++
		avunculateMarriage++
		cousinMarriage++
	}

	if r.Doctrines.FullTolerance {
		unrestrictedMarriage++
	}
	if r.Doctrines.Asceticism {
		avunculateMarriage++
	}
	if r.Doctrines.Polyamory {
		unrestrictedMarriage++
	}
	if r.Doctrines.ReligiousLaw {
		avunculateMarriage++
		cousinMarriage++
	}
	if r.Doctrines.SacredChildbirth {
		avunculateMarriage++
	}
	if r.Doctrines.Hedonism {
		unrestrictedMarriage++
	}

	if avunculateMarriage >= cousinMarriage && avunculateMarriage >= closeKinTaboo && avunculateMarriage >= unrestrictedMarriage {
		return religion.AvunculateMarriage
	}
	if cousinMarriage > closeKinTaboo && cousinMarriage >= unrestrictedMarriage {
		return religion.CousinMarriage
	}
	if closeKinTaboo >= unrestrictedMarriage {
		return religion.CloseKinTaboo
	}

	return religion.UnrestrictedMarriage
}

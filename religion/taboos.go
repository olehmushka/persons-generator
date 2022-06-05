package religion

import (
	"persons_generator/entities"
	rel "persons_generator/entities/religion"
)

func generateTaboos(r *entities.Religion) {
	t := &rel.Taboos{}
	t.RaisingTaboos = getRaisingTaboos(r)
	t.SexualTaboos = getSexualTaboos(r)
	t.EatingTaboos = getEatingTaboos(r)
	t.KillingTaboos = getKillingTaboos(r)
	t.BodyModificationTaboos = getBodyModificationTaboos(r)
	t.Witchcraft = getWitchcraft(r)
	t.Nudism = getNudism(r)
	r.Taboos = t
}

func getRaisingTaboos(r *entities.Religion) *rel.RaisingTaboos {
	rt := &rel.RaisingTaboos{}
	rt.RaisingPlants = getRaisingPlantsAcceptance(r)
	rt.RaisingAnimals = getRaisingAnimalsAcceptance(r)
	rt.RaisingFungus = getRaisingFungusAcceptance(r)
	return rt
}

func getRaisingPlantsAcceptance(r *entities.Religion) rel.Acceptance {
	var (
		accepted = 15
		shunned  = 10
		criminal = 5
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 15
	case r.Doctrines.Base.Polytheism:
		accepted += 15
	case r.Doctrines.Base.DeityDualism:
		accepted += 15
	case r.Doctrines.Base.Deism:
		accepted += 12
	case r.Doctrines.Base.Henothism:
		accepted += 15
	case r.Doctrines.Base.Monolatry:
		accepted += 14
	case r.Doctrines.Base.Omnism:
		accepted += 15
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 15
	case r.Doctrines.Gender.Equality:
		accepted += 15
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 15
	}

	if r.Doctrines.FullTolerance {
		accepted += 20
	}
	if r.Doctrines.Astrology {
		accepted += 30
	}
	if r.Doctrines.Legalism {
		accepted += 30
	}
	if r.Doctrines.SanctityOfNature {
		shunned += 10
	}
	if r.Doctrines.Pacifism {
		accepted += 25
	}
	if r.Doctrines.Reincarnation {
		accepted += 5
	}
	if r.Doctrines.SacredChildbirth {
		accepted += 5
	}
	if r.Doctrines.SunWorship {
		accepted += 10
	}
	if r.Doctrines.MoonWorship {
		accepted += 5
	}
	if r.Doctrines.Darkness {
		shunned += 15
	}
	if r.Doctrines.LiveUnderGround {
		criminal += 20
	}
	if r.Doctrines.TreeConnection {
		accepted += 10
	}
	if r.Doctrines.Raider {
		shunned += 10
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getRaisingAnimalsAcceptance(r *entities.Religion) rel.Acceptance {
	var (
		accepted = 15
		shunned  = 10
		criminal = 5
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 20
	case r.Doctrines.Base.Polytheism:
		accepted += 15
	case r.Doctrines.Base.DeityDualism:
		accepted += 17
	case r.Doctrines.Base.Deism:
		accepted += 15
	case r.Doctrines.Base.Henothism:
		accepted += 16
	case r.Doctrines.Base.Monolatry:
		accepted += 15
	case r.Doctrines.Base.Omnism:
		accepted += 20
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 20
	case r.Doctrines.Gender.Equality:
		accepted += 15
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 10
	}

	if r.Doctrines.FullTolerance {
		accepted += 20
	}
	if r.Doctrines.Legalism {
		accepted += 30
	}
	if r.Doctrines.SanctityOfNature {
		shunned += 5
	}
	if r.Doctrines.Pacifism {
		accepted += 7
		shunned += 5
	}
	if r.Doctrines.Reincarnation {
		shunned += 5
	}
	if r.Doctrines.Darkness {
		criminal += 15
	}
	if r.Doctrines.LiveUnderGround {
		criminal += 20
	}
	if r.Doctrines.AnimalConnection {
		accepted += 20
	}
	if r.Doctrines.Raider {
		accepted += 10
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getRaisingFungusAcceptance(r *entities.Religion) rel.Acceptance {
	var (
		accepted = 20
		shunned  = 15
		criminal = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 20
	case r.Doctrines.Base.Polytheism:
		accepted += 25
	case r.Doctrines.Base.DeityDualism:
		accepted += 22
	case r.Doctrines.Base.Deism:
		accepted += 20
	case r.Doctrines.Base.Henothism:
		accepted += 18
	case r.Doctrines.Base.Monolatry:
		accepted += 17
	case r.Doctrines.Base.Omnism:
		accepted += 25
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 20
	case r.Doctrines.Gender.Equality:
		accepted += 20
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 30
	}

	if r.Doctrines.FullTolerance {
		accepted += 20
	}
	if r.Doctrines.Asceticism {
		accepted += 5
	}
	if r.Doctrines.Astrology {
		accepted += 5
	}
	if r.Doctrines.Esotericism {
		accepted += 10
	}
	if r.Doctrines.SanctityOfNature {
		accepted += 5
	}
	if r.Doctrines.Pacifism {
		accepted += 10
	}
	if r.Doctrines.MoonWorship {
		accepted += 30
	}
	if r.Doctrines.Darkness {
		accepted += 30
	}
	if r.Doctrines.LiveUnderGround {
		accepted += 40
	}
	if r.Doctrines.Raider {
		shunned += 15
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getSexualTaboos(r *entities.Religion) *rel.SexualTaboos {
	st := &rel.SexualTaboos{}
	st.SameSexRelations = getSameSexRelations(r)
	st.MaleAdultery = getMaleAdultery(r)
	st.FemaleAdultery = getFemaleAdultery(r)
	st.Deviancy = getDeviancy(r)
	return st
}

func getSameSexRelations(r *entities.Religion) rel.Acceptance {
	var (
		accepted = 5
		shunned  = 15
		criminal = 30
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		shunned += 10
		criminal += 30
	case r.Doctrines.Base.Polytheism:
		accepted += 20
		shunned += 25
		criminal += 10
	case r.Doctrines.Base.DeityDualism:
		accepted += 5
		shunned += 15
		criminal += 18
	case r.Doctrines.Base.Deism:
		accepted += 20
		shunned += 20
		criminal += 10
	case r.Doctrines.Base.Henothism:
		accepted += 12
		shunned += 25
		criminal += 15
	case r.Doctrines.Base.Monolatry:
		accepted += 15
		shunned += 25
		criminal += 12
	case r.Doctrines.Base.Omnism:
		accepted += 20
		shunned += 25
		criminal += 5
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		shunned += 10
		criminal += 25
	case r.Doctrines.Gender.Equality:
		accepted += 15
		shunned += 20
		criminal += 15
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 17
		shunned += 25
		criminal += 10
	}

	if r.Doctrines.FullTolerance {
		accepted += 20
	}
	if r.Doctrines.Prophets {
		criminal += 10
	}
	if r.Doctrines.Astrology {
		shunned += 5
		accepted += 5
	}
	if r.Doctrines.Esotericism {
		shunned += 5
		accepted += 3
	}
	if r.Doctrines.Legalism {
		shunned += 5
		criminal += 10
	}
	if r.Doctrines.Polyamory {
		accepted += 40
	}
	if r.Doctrines.ReligiousLaw {
		shunned += 20
		criminal += 40
	}
	if r.Doctrines.RitualHospitality {
		accepted += 3
	}
	if r.Doctrines.SacredChildbirth {
		shunned += 7
		criminal += 10
	}
	if r.Doctrines.Raider {
		shunned += 30
		criminal += 15
	}
	if r.Doctrines.Hedonism {
		accepted += 15
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getMaleAdultery(r *entities.Religion) rel.Acceptance {
	var (
		accepted = 15
		shunned  = 20
		criminal = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 2
		shunned += 20
		criminal += 5
	case r.Doctrines.Base.Polytheism:
		accepted += 5
		shunned += 18
		criminal += 2
	case r.Doctrines.Base.DeityDualism:
		accepted += 3
		shunned += 20
		criminal += 3
	case r.Doctrines.Base.Deism:
		accepted += 2
		shunned += 20
		criminal += 2
	case r.Doctrines.Base.Henothism:
		accepted += 3
		shunned += 19
		criminal += 3
	case r.Doctrines.Base.Monolatry:
		accepted += 4
		shunned += 18
		criminal += 2
	case r.Doctrines.Base.Omnism:
		accepted += 10
		shunned += 15
		criminal += 1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 10
		shunned += 20
	case r.Doctrines.Gender.Equality:
		shunned += 30
	case r.Doctrines.Gender.FemaleDominance:
		criminal += 30
		shunned += 40
	}

	if r.Doctrines.FullTolerance {
		accepted += 10
	}
	if r.Doctrines.Asceticism {
		shunned += 15
		criminal += 5
	}
	if r.Doctrines.Legalism {
		shunned += 10
		criminal += 3
	}
	if r.Doctrines.Polyamory {
		accepted += 40
	}
	if r.Doctrines.ReligiousLaw {
		shunned += 20
	}
	if r.Doctrines.Raider {
		shunned += 10
		accepted += 20
	}
	if r.Doctrines.Hedonism {
		accepted += 40
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getFemaleAdultery(r *entities.Religion) rel.Acceptance {
	var (
		accepted = 10
		shunned  = 15
		criminal = 20
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 1
		shunned += 35
		criminal += 15
	case r.Doctrines.Base.Polytheism:
		accepted += 3
		shunned += 30
		criminal += 5
	case r.Doctrines.Base.DeityDualism:
		accepted += 2
		shunned += 30
		criminal += 5
	case r.Doctrines.Base.Deism:
		accepted += 5
		shunned += 30
		criminal += 5
	case r.Doctrines.Base.Henothism:
		accepted += 1
		shunned += 30
		criminal += 5
	case r.Doctrines.Base.Monolatry:
		accepted += 1
		shunned += 31
		criminal += 5
	case r.Doctrines.Base.Omnism:
		accepted += 4
		shunned += 25
		criminal += 2
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		criminal += 60
		shunned += 40
	case r.Doctrines.Gender.Equality:
		shunned += 30
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 20
		shunned += 45
	}

	if r.Doctrines.FullTolerance {
		accepted += 10
	}
	if r.Doctrines.Prophets {
		shunned += 20
	}
	if r.Doctrines.Asceticism {
		shunned += 20
		criminal += 8
	}
	if r.Doctrines.Legalism {
		shunned += 20
		criminal += 10
	}
	if r.Doctrines.Polyamory {
		accepted += 40
	}
	if r.Doctrines.Legalism {
		shunned += 30
		criminal += 20
	}
	if r.Doctrines.AncestorWorship {
		shunned += 25
	}
	if r.Doctrines.Raider {
		shunned += 25
		criminal += 10
	}
	if r.Doctrines.Hedonism {
		shunned += 25
		accepted += 20
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getDeviancy(r *entities.Religion) rel.Acceptance {
	var (
		accepted = 20
		shunned  = 15
		criminal = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		shunned += 20
		criminal += 27
	case r.Doctrines.Base.Polytheism:
		accepted += 20
		shunned += 30
		criminal += 8
	case r.Doctrines.Base.DeityDualism:
		accepted += 1
		shunned += 22
		criminal += 17
	case r.Doctrines.Base.Deism:
		accepted += 25
		shunned += 28
		criminal += 9
	case r.Doctrines.Base.Henothism:
		accepted += 11
		shunned += 30
		criminal += 10
	case r.Doctrines.Base.Monolatry:
		accepted += 13
		shunned += 28
		criminal += 10
	case r.Doctrines.Base.Omnism:
		accepted += 20
		shunned += 25
		criminal += 5
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		shunned += 20
		criminal += 25
	case r.Doctrines.Gender.Equality:
		accepted += 5
		shunned += 25
		criminal += 15
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 3
		shunned += 25
		criminal += 20
	}

	if r.Doctrines.FullTolerance {
		accepted += 20
	}
	if r.Doctrines.Prophets {
		shunned += 10
		criminal += 8
	}
	if r.Doctrines.Astrology {
		shunned += 15
		accepted += 2
	}
	if r.Doctrines.Esotericism {
		shunned += 15
		accepted += 1
	}
	if r.Doctrines.Legalism {
		shunned += 15
		criminal += 20
	}
	if r.Doctrines.Polyamory {
		accepted += 30
	}
	if r.Doctrines.ReligiousLaw {
		shunned += 25
		criminal += 40
	}
	if r.Doctrines.RitualHospitality {
		accepted += 1
	}
	if r.Doctrines.SacredChildbirth {
		shunned += 17
		criminal += 15
	}
	if r.Doctrines.Raider {
		shunned += 35
		criminal += 18
	}
	if r.Doctrines.Hedonism {
		accepted += 15
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getEatingTaboos(r *entities.Religion) *rel.EatingTaboos {
	return &rel.EatingTaboos{}
}

func getKillingTaboos(r *entities.Religion) *rel.KillingTaboos {
	return &rel.KillingTaboos{}
}

func getBodyModificationTaboos(r *entities.Religion) *rel.BodyModificationTaboos {
	return &rel.BodyModificationTaboos{}
}

func getWitchcraft(r *entities.Religion) rel.Acceptance {
	return ""
}

func getNudism(r *entities.Religion) rel.Acceptance {
	return ""
}

package religion

import (
	"fmt"

	"persons_generator/entities"
	rel "persons_generator/entities/religion"
)

func generateTaboos(r *entities.Religion) {
	fmt.Println("[generateTaboos] started")
	defer fmt.Println("[generateTaboos] finished")

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
		accepted = 0.15
		shunned  = 0.1
		criminal = 0.05
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.15
	case r.Doctrines.Base.Polytheism:
		accepted += 0.15
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.15
	case r.Doctrines.Base.Deism:
		accepted += 0.12
	case r.Doctrines.Base.Henothism:
		accepted += 0.15
	case r.Doctrines.Base.Monolatry:
		accepted += 0.14
	case r.Doctrines.Base.Omnism:
		accepted += 0.15
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.05
	case r.Doctrines.Gender.Equality:
		accepted += 0.05
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.05
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.01
	}
	if r.Doctrines.Astrology {
		accepted += 0.05
	}
	if r.Doctrines.Legalism {
		accepted += 0.05
	}
	if r.Doctrines.SanctityOfNature {
		shunned += 0.05
	}
	if r.Doctrines.Pacifism {
		accepted += 0.1
	}
	if r.Doctrines.Reincarnation {
		accepted += 0.01
	}
	if r.Doctrines.SacredChildbirth {
		accepted += 0.01
	}
	if r.Doctrines.SunWorship {
		accepted += 0.1
	}
	if r.Doctrines.MoonWorship {
		accepted += 0.005
	}
	if r.Doctrines.Darkness {
		shunned += 0.05
	}
	if r.Doctrines.LiveUnderGround {
		criminal += 0.07
	}
	if r.Doctrines.TreeConnection {
		accepted += 0.08
	}
	if r.Doctrines.Raider {
		shunned += 0.1
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getRaisingAnimalsAcceptance(r *entities.Religion) rel.Acceptance {
	var (
		accepted = 0.15
		shunned  = 0.1
		criminal = 0.05
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.2
	case r.Doctrines.Base.Polytheism:
		accepted += 0.15
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.17
	case r.Doctrines.Base.Deism:
		accepted += 0.15
	case r.Doctrines.Base.Henothism:
		accepted += 0.16
	case r.Doctrines.Base.Monolatry:
		accepted += 0.15
	case r.Doctrines.Base.Omnism:
		accepted += 0.2
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.05
	case r.Doctrines.Gender.Equality:
		accepted += 0.04
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.037
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.05
	}
	if r.Doctrines.Legalism {
		accepted += 0.08
	}
	if r.Doctrines.SanctityOfNature {
		shunned += 0.01
	}
	if r.Doctrines.Pacifism {
		accepted += 0.03
		shunned += 0.01
	}
	if r.Doctrines.Reincarnation {
		shunned += 0.02
	}
	if r.Doctrines.Darkness {
		criminal += 0.09
	}
	if r.Doctrines.LiveUnderGround {
		criminal += 0.05
	}
	if r.Doctrines.AnimalConnection {
		accepted += 0.05
	}
	if r.Doctrines.Raider {
		accepted += 0.1
		shunned += 0.02
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getRaisingFungusAcceptance(r *entities.Religion) rel.Acceptance {
	var (
		accepted = 0.211
		shunned  = 0.15
		criminal = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.2
	case r.Doctrines.Base.Polytheism:
		accepted += 0.25
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.22
	case r.Doctrines.Base.Deism:
		accepted += 0.2
	case r.Doctrines.Base.Henothism:
		accepted += 0.18
	case r.Doctrines.Base.Monolatry:
		accepted += 0.17
	case r.Doctrines.Base.Omnism:
		accepted += 0.25
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.05
	case r.Doctrines.Gender.Equality:
		accepted += 0.05
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.07
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.05
	}
	if r.Doctrines.Asceticism {
		accepted += 0.02
	}
	if r.Doctrines.Astrology {
		accepted += 0.05
	}
	if r.Doctrines.Esotericism {
		accepted += 0.08
	}
	if r.Doctrines.SanctityOfNature {
		accepted += 0.01
	}
	if r.Doctrines.Pacifism {
		accepted += 0.02
	}
	if r.Doctrines.MoonWorship {
		accepted += 0.1
	}
	if r.Doctrines.Darkness {
		accepted += 0.11
	}
	if r.Doctrines.LiveUnderGround {
		accepted += 0.25
	}
	if r.Doctrines.Raider {
		shunned += 0.1
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
		accepted = 0.05
		shunned  = 0.15
		criminal = 0.3
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		shunned += 0.1
		criminal += 0.3
	case r.Doctrines.Base.Polytheism:
		accepted += 0.2
		shunned += 0.25
		criminal += 0.19
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.05
		shunned += 0.15
		criminal += 0.18
	case r.Doctrines.Base.Deism:
		accepted += 0.2
		shunned += 0.2
		criminal += 0.1
	case r.Doctrines.Base.Henothism:
		accepted += 0.12
		shunned += 0.25
		criminal += 0.15
	case r.Doctrines.Base.Monolatry:
		accepted += 0.15
		shunned += 0.25
		criminal += 0.12
	case r.Doctrines.Base.Omnism:
		accepted += 0.2
		shunned += 0.25
		criminal += 0.05
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		shunned += 0.1
		criminal += 0.25
	case r.Doctrines.Gender.Equality:
		accepted += 0.15
		shunned += 0.2
		criminal += 0.159
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.17
		shunned += 0.25
		criminal += 0.175
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.05
	}
	if r.Doctrines.Prophets {
		criminal += 0.03
	}
	if r.Doctrines.Astrology {
		shunned += 0.01
		accepted += 0.01
	}
	if r.Doctrines.Esotericism {
		shunned += 0.02
		accepted += 0.01
	}
	if r.Doctrines.Legalism {
		shunned += 0.02
		criminal += 0.06
	}
	if r.Doctrines.Polyamory {
		accepted += 0.15
		shunned += 0.07
	}
	if r.Doctrines.ReligiousLaw {
		shunned += 0.05
		criminal += 0.2
	}
	if r.Doctrines.RitualHospitality {
		accepted += 0.03
		shunned += 0.02
	}
	if r.Doctrines.SacredChildbirth {
		shunned += 0.03
		criminal += 0.07
	}
	if r.Doctrines.Raider {
		shunned += 0.15
		criminal += 0.1
	}
	if r.Doctrines.Hedonism {
		accepted += 0.05
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getMaleAdultery(r *entities.Religion) rel.Acceptance {
	var (
		accepted = 0.15
		shunned  = 0.2
		criminal = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.02
		shunned += 0.2
		criminal += 0.05
	case r.Doctrines.Base.Polytheism:
		accepted += 0.07
		shunned += 0.18
		criminal += 0.02
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.049
		shunned += 0.2
		criminal += 0.03
	case r.Doctrines.Base.Deism:
		accepted += 0.03
		shunned += 0.2
		criminal += 0.02
	case r.Doctrines.Base.Henothism:
		accepted += 0.04
		shunned += 0.19
		criminal += 0.03
	case r.Doctrines.Base.Monolatry:
		accepted += 0.04
		shunned += 0.18
		criminal += 0.02
	case r.Doctrines.Base.Omnism:
		accepted += 0.1
		shunned += 0.15
		criminal += 0.01
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.1
		shunned += 0.2
	case r.Doctrines.Gender.Equality:
		shunned += 0.3
	case r.Doctrines.Gender.FemaleDominance:
		criminal += 0.3
		shunned += 0.4
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.1
	}
	if r.Doctrines.Asceticism {
		shunned += 0.15
		criminal += 0.05
	}
	if r.Doctrines.Legalism {
		shunned += 0.1
		criminal += 0.03
	}
	if r.Doctrines.Polyamory {
		accepted += 0.4
	}
	if r.Doctrines.ReligiousLaw {
		shunned += 0.2
	}
	if r.Doctrines.Raider {
		shunned += 0.1
		accepted += 0.2
	}
	if r.Doctrines.Hedonism {
		accepted += 0.4
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getFemaleAdultery(r *entities.Religion) rel.Acceptance {
	var (
		accepted = 0.1
		shunned  = 0.15
		criminal = 0.2
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.01
		shunned += 0.35
		criminal += 0.15
	case r.Doctrines.Base.Polytheism:
		accepted += 0.05
		shunned += 0.315
		criminal += 0.15
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.02
		shunned += 0.31
		criminal += 0.05
	case r.Doctrines.Base.Deism:
		accepted += 0.041
		shunned += 0.305
		criminal += 0.05
	case r.Doctrines.Base.Henothism:
		accepted += 0.01
		shunned += 0.3
		criminal += 0.05
	case r.Doctrines.Base.Monolatry:
		accepted += 0.01
		shunned += 0.31
		criminal += 0.05
	case r.Doctrines.Base.Omnism:
		accepted += 0.04
		shunned += 0.025
		criminal += 0.02
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		shunned += 0.4
		criminal += 0.6
	case r.Doctrines.Gender.Equality:
		shunned += 0.3
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.35
		shunned += 0.45
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.15
	}
	if r.Doctrines.Prophets {
		shunned += 0.2
	}
	if r.Doctrines.Asceticism {
		shunned += 0.2
		criminal += 0.08
	}
	if r.Doctrines.Legalism {
		shunned += 0.2
		criminal += 0.19
	}
	if r.Doctrines.Polyamory {
		accepted += 0.4
	}
	if r.Doctrines.Legalism {
		shunned += 0.3
		criminal += 0.2
	}
	if r.Doctrines.AncestorWorship {
		shunned += 0.25
	}
	if r.Doctrines.Raider {
		shunned += 0.25
		criminal += 0.1
	}
	if r.Doctrines.Hedonism {
		shunned += 0.25
		accepted += 0.2
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getDeviancy(r *entities.Religion) rel.Acceptance {
	var (
		accepted = 0.2
		shunned  = 0.15
		criminal = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		shunned += 0.2
		criminal += 0.279
	case r.Doctrines.Base.Polytheism:
		accepted += 0.2
		shunned += 0.39
		criminal += 0.08
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.019
		shunned += 0.225
		criminal += 0.17
	case r.Doctrines.Base.Deism:
		accepted += 0.25
		shunned += 0.28
		criminal += 0.09
	case r.Doctrines.Base.Henothism:
		accepted += 0.11
		shunned += 0.3
		criminal += 0.1
	case r.Doctrines.Base.Monolatry:
		accepted += 0.13
		shunned += 0.28
		criminal += 0.19
	case r.Doctrines.Base.Omnism:
		accepted += 0.2
		shunned += 0.25
		criminal += 0.05
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		shunned += 0.2
		criminal += 0.25
	case r.Doctrines.Gender.Equality:
		accepted += 0.05
		shunned += 0.25
		criminal += 0.15
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.03
		shunned += 0.25
		criminal += 0.2
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.2
	}
	if r.Doctrines.Prophets {
		shunned += 0.1
		criminal += 0.08
	}
	if r.Doctrines.Astrology {
		shunned += 0.15
		accepted += 0.02
	}
	if r.Doctrines.Esotericism {
		shunned += 0.15
		accepted += 0.01
	}
	if r.Doctrines.Legalism {
		shunned += 0.15
		criminal += 0.2
	}
	if r.Doctrines.Polyamory {
		accepted += 0.3
	}
	if r.Doctrines.ReligiousLaw {
		shunned += 0.25
		criminal += 0.4
	}
	if r.Doctrines.RitualHospitality {
		accepted += 0.01
		shunned += 0.005
	}
	if r.Doctrines.SacredChildbirth {
		shunned += 0.17
		criminal += 0.15
	}
	if r.Doctrines.Raider {
		shunned += 0.35
		criminal += 0.18
	}
	if r.Doctrines.Hedonism {
		accepted += 0.15
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

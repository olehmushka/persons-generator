package religion

import (
	"fmt"

	"persons_generator/entities"
	"persons_generator/entities/religion"
	rel "persons_generator/entities/religion"
)

func generateTaboos(r *entities.Religion) {
	fmt.Println("[generateTaboos] started")
	defer fmt.Println("[generateTaboos] finished")

	r.Taboos = &rel.Taboos{}
	getRaisingTaboos(r)
	getSexualTaboos(r)
	getEatingTaboos(r)
	getKillingTaboos(r)
	getBodyModificationTaboos(r)
	r.Taboos.Witchcraft = getWitchcraft(r)
	r.Taboos.Nudism = getNudism(r)
}

func getRaisingTaboos(r *entities.Religion) {
	r.Taboos.RaisingTaboos = &rel.RaisingTaboos{}
	r.Taboos.RaisingTaboos.RaisingPlants = getRaisingPlantsAcceptance(r)
	r.Taboos.RaisingTaboos.RaisingAnimals = getRaisingAnimalsAcceptance(r)
	r.Taboos.RaisingTaboos.RaisingFungus = getRaisingFungusAcceptance(r)
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

func getSexualTaboos(r *entities.Religion) {
	r.Taboos.SexualTaboos = &rel.SexualTaboos{}
	r.Taboos.SexualTaboos.SameSexRelations = getSameSexRelations(r)
	r.Taboos.SexualTaboos.MaleAdultery = getMaleAdultery(r)
	r.Taboos.SexualTaboos.FemaleAdultery = getFemaleAdultery(r)
	r.Taboos.SexualTaboos.Deviancy = getDeviancy(r)
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

func getEatingTaboos(r *entities.Religion) {
	r.Taboos.EatingTaboos = &rel.EatingTaboos{}
	r.Taboos.EatingTaboos.EatAnimals = getEatAnimals(r)
	r.Taboos.EatingTaboos.EatInsects = getEatInsects(r)
	r.Taboos.EatingTaboos.EatFungus = getEatFungus(r)
	r.Taboos.EatingTaboos.Cannibalism = getCannibalism(r)
	r.Taboos.EatingTaboos.DrinkingBlood = getDrinkingBlood(r)
	r.Taboos.EatingTaboos.DrinkingStrongAlcohol = getDrinkingStrongAlcohol(r)
	r.Taboos.EatingTaboos.DrinkingNotStrongAlcohol = getDrinkingNotStrongAlcohol(r)
	r.Taboos.EatingTaboos.UseNicotine = getUseNicotine(r)
	r.Taboos.EatingTaboos.UseCannabis = getUseCannabis(r)
	r.Taboos.EatingTaboos.UseHallucinogens = getUseHallucinogens(r)
	r.Taboos.EatingTaboos.UseCNSStimulants = getUseCNSStimulants(r)
	r.Taboos.EatingTaboos.UseOpium = getUseOpium(r)
}

func getEatAnimals(r *entities.Religion) religion.Acceptance {
	var (
		accepted = 0.3
		shunned  = 0.2
		criminal = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.05
		criminal += 0.01
	case r.Doctrines.Base.Polytheism:
		accepted += 0.05
		shunned += 0.01
		criminal += 0.009
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.05
		shunned += 0.02
		criminal += 0.01
	case r.Doctrines.Base.Deism:
		accepted += 0.05
		shunned += 0.03
		criminal += 0.02
	case r.Doctrines.Base.Henothism:
		accepted += 0.05
		shunned += 0.01
		criminal += 0.01
	case r.Doctrines.Base.Monolatry:
		accepted += 0.05
		shunned += 0.02
		criminal += 0.015
	case r.Doctrines.Base.Omnism:
		accepted += 0.05
		shunned += 0.01
		criminal += 0.005
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.1
	case r.Doctrines.Gender.Equality:
		shunned += 0.01
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.03
		shunned += 0.01
		criminal += 0.005
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.05
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
		shunned += 0.02
		criminal += 0.01
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
		accepted += 0.01
		shunned += 0.009
		criminal += 0.008
	}
	if r.Doctrines.Legalism {
		accepted += 0.01
		criminal += 0.005
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
		accepted += 0.01
		criminal += 0.01
	}
	if r.Doctrines.SanctityOfNature {
		shunned += 0.01
		criminal += 0.007
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
		shunned += 0.02
		criminal += 0.01
	}
	if r.Doctrines.RitualHospitality {
		accepted += 0.01
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
		accepted += 0.01
	}
	if r.Doctrines.SunWorship {
		accepted += 0.01
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
		accepted += 0.02
	}
	if r.Doctrines.Darkness {
		accepted += 0.02
	}
	if r.Doctrines.LiveUnderGround {
		accepted += 0.01
	}
	if r.Doctrines.TreeConnection {
		shunned += 0.01
	}
	if r.Doctrines.AnimalConnection {
		accepted += 0.01
		shunned += 0.02
		criminal += 0.01
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
		accepted += 0.03
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
		accepted += 0.05
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
		accepted += 0.01
		shunned += 0.01
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		accepted += 0.05
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
		shunned += 0.01
		criminal += 0.005
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
		shunned += 0.005
		criminal += 0.001
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getEatInsects(r *entities.Religion) religion.Acceptance {
	var (
		accepted = 0.1
		shunned  = 0.15
		criminal = 0.05
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.05
		shunned += 0.05
		criminal += 0.01
	case r.Doctrines.Base.Polytheism:
		accepted += 0.05
		shunned += 0.04
		criminal += 0.01
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.05
		shunned += 0.05
		criminal += 0.05
	case r.Doctrines.Base.Deism:
		accepted += 0.1
		shunned += 0.05
		criminal += 0.02
	case r.Doctrines.Base.Henothism:
		accepted += 0.05
		shunned += 0.045
		criminal += 0.02
	case r.Doctrines.Base.Monolatry:
		accepted += 0.05
		shunned += 0.03
		criminal += 0.015
	case r.Doctrines.Base.Omnism:
		accepted += 0.07
		shunned += 0.03
		criminal += 0.01
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.03
		shunned += 0.015
	case r.Doctrines.Gender.Equality:
		accepted += 0.01
		shunned += 0.2
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.005
		criminal += 0.01
		shunned += 0.03
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
		accepted += 0.02
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
		accepted += 0.01
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
		accepted += 0.01
		shunned += 0.01
		criminal += 0.01
	}
	if r.Doctrines.SanctityOfNature {
		shunned += 0.01
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
		shunned += 0.02
		criminal += 0.01
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
	}
	if r.Doctrines.MoonWorship {
		accepted += 0.02
		shunned += 0.01
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
		accepted += 0.05
	}
	if r.Doctrines.LiveUnderGround {
		accepted += 0.1
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
		accepted += 0.01
		shunned += 0.01
	}
	if r.Doctrines.Raider {
		accepted += 0.005
		shunned += 0.01
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
		shunned += 0.03
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
		shunned += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getEatFungus(r *entities.Religion) religion.Acceptance {
	var (
		accepted = 0.3
		shunned  = 0.15
		criminal = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.05
		shunned += 0.01
		criminal += 0.005
	case r.Doctrines.Base.Polytheism:
		accepted += 0.09
		shunned += 0.01
		criminal += 0.01
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.05
		shunned += 0.02
		criminal += 0.005
	case r.Doctrines.Base.Deism:
		accepted += 0.1
		shunned += 0.02
		criminal += 0.03
	case r.Doctrines.Base.Henothism:
		accepted += 0.05
		shunned += 0.03
		criminal += 0.02
	case r.Doctrines.Base.Monolatry:
		accepted += 0.09
		shunned += 0.03
		criminal += 0.02
	case r.Doctrines.Base.Omnism:
		accepted += 0.1
		shunned += 0.04
		criminal += 0.02
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.05
		shunned += 0.01
	case r.Doctrines.Gender.Equality:
		accepted += 0.03
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.01
		shunned += 0.01
		criminal += 0.0001
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
		accepted += 0.01
	}
	if r.Doctrines.Astrology {
		accepted += 0.01
	}
	if r.Doctrines.Esotericism {
		accepted += 0.05
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
		accepted += 0.01
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
		accepted += 0.001
		shunned += 0.001
		criminal += 0.001
	}
	if r.Doctrines.SanctityOfNature {
		accepted += 0.01
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
		accepted += 0.01
	}
	if r.Doctrines.Pacifism {
		accepted += 0.02
	}
	if r.Doctrines.Reincarnation {
		accepted += 0.005
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		shunned += 0.01
		criminal += 0.005
	}
	if r.Doctrines.MoonWorship {
		accepted += 0.05
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
		accepted += 0.09
	}
	if r.Doctrines.LiveUnderGround {
		accepted += 0.09
	}
	if r.Doctrines.TreeConnection {
		accepted += 0.02
	}
	if r.Doctrines.AnimalConnection {
		shunned += 0.01
	}
	if r.Doctrines.Blindsight {
		accepted += 0.01
	}
	if r.Doctrines.Raider {
		accepted += 0.01
		shunned += 0.015
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
		accepted += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		accepted += 0.02
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getCannibalism(r *entities.Religion) religion.Acceptance {
	var (
		accepted = 0.05
		shunned  = 0.1
		criminal = 0.3
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.01
		shunned += 0.07
		criminal += 0.09
	case r.Doctrines.Base.Polytheism:
		accepted += 0.03
		shunned += 0.07
		criminal += 0.08
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.02
		shunned += 0.07
		criminal += 0.077
	case r.Doctrines.Base.Deism:
		accepted += 0.03
		shunned += 0.08
		criminal += 0.05
	case r.Doctrines.Base.Henothism:
		accepted += 0.025
		shunned += 0.069
		criminal += 0.065
	case r.Doctrines.Base.Monolatry:
		accepted += 0.03
		shunned += 0.07
		criminal += 0.075
	case r.Doctrines.Base.Omnism:
		accepted += 0.03
		shunned += 0.07
		criminal += 0.08
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.01
		shunned += 0.05
		criminal += 0.02
	case r.Doctrines.Gender.Equality:
		shunned += 0.03
	case r.Doctrines.Gender.FemaleDominance:
		shunned += 0.1
		criminal += 0.05
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
		shunned += 0.02
		criminal += 0.005
	}
	if r.Doctrines.Asceticism {
		shunned += 0.01
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
		accepted += 0.01
	}
	if r.Doctrines.Legalism {
		criminal += 0.03
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
		criminal += 0.02
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
		accepted += 0.01
	}
	if r.Doctrines.Pacifism {
		criminal += 0.01
	}
	if r.Doctrines.Reincarnation {
		criminal += 0.01
	}
	if r.Doctrines.RitualHospitality {
		shunned += 0.01
	}
	if r.Doctrines.SacredChildbirth {
		shunned += 0.02
		criminal += 0.01
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		shunned += 0.01
		criminal += 0.02
	}
	if r.Doctrines.MoonWorship {
		accepted += 0.01
		shunned += 0.02
	}
	if r.Doctrines.PainIsVirtue {
		accepted += 0.01
		shunned += 0.01
	}
	if r.Doctrines.Darkness {
		accepted += 0.05
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
		criminal += 0.01
	}
	if r.Doctrines.AnimalConnection {
		shunned += 0.01
		criminal += 0.005
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
		accepted += 0.01
		shunned += 0.005
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
		accepted += 0.01
		shunned += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
		shunned += 0.01
		criminal += 0.005
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
		shunned += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
		accepted += 0.005
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
		shunned += 0.005
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
		criminal += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
		criminal += 0.05
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getDrinkingBlood(r *entities.Religion) religion.Acceptance {
	var (
		accepted = 0.1
		shunned  = 0.15
		criminal = 0.12
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.05
		shunned += 0.07
		criminal += 0.04
	case r.Doctrines.Base.Polytheism:
		accepted += 0.06
		shunned += 0.06
		criminal += 0.06
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.05
		shunned += 0.08
		criminal += 0.05
	case r.Doctrines.Base.Deism:
		accepted += 0.05
		shunned += 0.05
		criminal += 0.05
	case r.Doctrines.Base.Henothism:
		accepted += 0.05
		shunned += 0.07
		criminal += 0.06
	case r.Doctrines.Base.Monolatry:
		accepted += 0.06
		shunned += 0.06
		criminal += 0.06
	case r.Doctrines.Base.Omnism:
		accepted += 0.06
		shunned += 0.05
		criminal += 0.05
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.01
		shunned += 0.03
		criminal += 0.01
	case r.Doctrines.Gender.Equality:
		shunned += 0.03
		criminal += 0.05
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.02
		shunned += 0.05
		criminal += 0.05
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
		criminal += 0.01
	}
	if r.Doctrines.Asceticism {
		criminal += 0.01
	}
	if r.Doctrines.Astrology {
		accepted += 0.01
	}
	if r.Doctrines.Esotericism {
		accepted += 0.01
		criminal += 0.01
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
		accepted += 0.02
		criminal += 0.02
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
		criminal += 0.05
	}
	if r.Doctrines.Pacifism {
		shunned += 0.03
	}
	if r.Doctrines.Reincarnation {
		shunned += 0.01
		criminal += 0.01
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
		accepted += 0.01
		criminal += 0.01
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		shunned += 0.02
		criminal += 0.01
	}
	if r.Doctrines.MoonWorship {
		accepted += 0.01
	}
	if r.Doctrines.PainIsVirtue {
		accepted += 0.01
	}
	if r.Doctrines.Darkness {
		accepted += 0.05
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
		accepted += 0.05
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
		shunned += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
		criminal += 0.05
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
		accepted += 0.03
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
		criminal += 0.03
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
		shunned += 0.01
		criminal += 0.02
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getDrinkingStrongAlcohol(r *entities.Religion) religion.Acceptance {
	var (
		accepted = 0.15
		shunned  = 0.15
		criminal = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.1
		shunned += 0.07
		criminal += 0.03
	case r.Doctrines.Base.Polytheism:
		accepted += 0.11
		shunned += 0.03
		criminal += 0.02
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.1
		shunned += 0.09
		criminal += 0.03
	case r.Doctrines.Base.Deism:
		accepted += 0.1
		shunned += 0.1
		criminal += 0.04
	case r.Doctrines.Base.Henothism:
		accepted += 0.1
		shunned += 0.04
		criminal += 0.025
	case r.Doctrines.Base.Monolatry:
		accepted += 0.12
		shunned += 0.035
		criminal += 0.02
	case r.Doctrines.Base.Omnism:
		accepted += 0.1
		shunned += 0.02
		criminal += 0.01
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.1
		shunned += 0.05
	case r.Doctrines.Gender.Equality:
		accepted += 0.06
		shunned += 0.05
		criminal += 0.01
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.01
		shunned += 0.02
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
		shunned += 0.01
		criminal += 0.005
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
		accepted += 0.01
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
		accepted += 0.01
		shunned += 0.005
	}
	if r.Doctrines.ReligiousLaw {
		criminal += 0.005
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		criminal += 0.01
	}
	if r.Doctrines.AncestorWorship {
		accepted += 0.01
		shunned += 0.005
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
		shunned += 0.01
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
		shunned += 0.005
		criminal += 0.01
	}
	if r.Doctrines.SanctionedFalseConversions {
		accepted += 0.01
	}
	if r.Doctrines.SunWorship {
		shunned += 0.01
	}
	if r.Doctrines.MoonWorship {
		accepted += 0.01
	}
	if r.Doctrines.PainIsVirtue {
		accepted += 0.01
	}
	if r.Doctrines.Darkness {
		accepted += 0.05
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
		accepted += 0.01
	}
	if r.Doctrines.Raider {
		accepted += 0.01
	}
	if r.Doctrines.Proselytizer {
		shunned += 0.01
		criminal += 0.02
	}
	if r.Doctrines.Hedonism {
		accepted += 0.02
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		accepted += 0.02
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
		shunned += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
		accepted += 0.005
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
		shunned += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
		criminal += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
		shunned += 0.005
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
		shunned += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
		shunned += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
		criminal += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
		accepted += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
		shunned += 0.01
		criminal += 0.02
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
		accepted += 0.01
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getDrinkingNotStrongAlcohol(r *entities.Religion) religion.Acceptance {
	var (
		accepted = 0.2
		shunned  = 0.1
		criminal = 0.02
	)
	switch {
	case r.Taboos.EatingTaboos.DrinkingStrongAlcohol.IsAccepted():
		accepted += 0.1
	case r.Taboos.EatingTaboos.DrinkingStrongAlcohol.IsShunned():
		shunned += 0.1
	case r.Taboos.EatingTaboos.DrinkingStrongAlcohol.IsCriminal():
		shunned += 0.07
		criminal += 0.1
	}

	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.1
		shunned += 0.05
		criminal += 0.01
	case r.Doctrines.Base.Polytheism:
		accepted += 0.15
		shunned += 0.03
		criminal += 0.01
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.1
		shunned += 0.04
		criminal += 0.01
	case r.Doctrines.Base.Deism:
		accepted += 0.1
		shunned += 0.03
		criminal += 0.01
	case r.Doctrines.Base.Henothism:
		accepted += 0.15
		shunned += 0.035
		criminal += 0.01
	case r.Doctrines.Base.Monolatry:
		accepted += 0.155
		shunned += 0.03
		criminal += 0.01
	case r.Doctrines.Base.Omnism:
		accepted += 0.159
		shunned += 0.031
		criminal += 0.01
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.05
		shunned += 0.01
	case r.Doctrines.Gender.Equality:
		accepted += 0.02
		shunned += 0.01
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.01
		shunned += 0.01
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
		shunned += 0.01
		criminal += 0.02
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
		accepted += 0.01
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
		accepted += 0.01
	}
	if r.Doctrines.ReligiousLaw {
		accepted += 0.01
		criminal += 0.01
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		shunned += 0.01
		criminal += 0.02
	}
	if r.Doctrines.AncestorWorship {
		shunned += 0.01
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
		shunned += 0.01
		criminal += 0.001
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
		shunned += 0.01
		criminal += 0.005
	}
	if r.Doctrines.SanctionedFalseConversions {
		accepted += 0.01
	}
	if r.Doctrines.SunWorship {
		accepted += 0.02
		shunned += 0.01
	}
	if r.Doctrines.MoonWorship {
		accepted += 0.01
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
		accepted += 0.05
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
		accepted += 0.01
	}
	if r.Doctrines.Raider {
		accepted += 0.02
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
		shunned += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
		shunned += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
		shunned += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getUseNicotine(r *entities.Religion) religion.Acceptance {
	var (
		accepted = 0.2
		shunned  = 0.2
		criminal = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.1
		shunned += 0.05
		criminal += 0.03
	case r.Doctrines.Base.Polytheism:
		accepted += 0.1
		shunned += 0.04
		criminal += 0.02
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.1
		shunned += 0.04
		criminal += 0.028
	case r.Doctrines.Base.Deism:
		accepted += 0.1
		shunned += 0.01
		criminal += 0.01
	case r.Doctrines.Base.Henothism:
		accepted += 0.1
		shunned += 0.045
		criminal += 0.025
	case r.Doctrines.Base.Monolatry:
		accepted += 0.11
		shunned += 0.035
		criminal += 0.015
	case r.Doctrines.Base.Omnism:
		accepted += 0.1
		shunned += 0.035
		criminal += 0.029
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.05
		shunned += 0.01
	case r.Doctrines.Gender.Equality:
		accepted += 0.03
		shunned += 0.01
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.01
		shunned += 0.03
		criminal += 0.02
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
		shunned += 0.01
		criminal += 0.02
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
		accepted += 0.01
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
		accepted += 0.02
		criminal += 0.02
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
		shunned += 0.01
		criminal += 0.02
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		accepted += 0.005
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
		accepted += 0.01
	}
	if r.Doctrines.Darkness {
		accepted += 0.05
	}
	if r.Doctrines.LiveUnderGround {
		shunned += 0.05
		criminal += 0.03
	}
	if r.Doctrines.TreeConnection {
		shunned += 0.01
	}
	if r.Doctrines.AnimalConnection {
		shunned += 0.01
	}
	if r.Doctrines.Blindsight {
		accepted += 0.01
	}
	if r.Doctrines.Raider {
		accepted += 0.003
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
		accepted += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
		shunned += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
		accepted += 0.001
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getUseCannabis(r *entities.Religion) religion.Acceptance {
	var (
		accepted = 0.2
		shunned  = 0.1
		criminal = 0.15
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.06
		shunned += 0.05
		criminal += 0.05
	case r.Doctrines.Base.Polytheism:
		accepted += 0.12
		shunned += 0.04
		criminal += 0.02
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.08
		shunned += 0.043
		criminal += 0.03
	case r.Doctrines.Base.Deism:
		accepted += 0.12
		shunned += 0.02
		criminal += 0.02
	case r.Doctrines.Base.Henothism:
		accepted += 0.11
		shunned += 0.041
		criminal += 0.021
	case r.Doctrines.Base.Monolatry:
		accepted += 0.121
		shunned += 0.04
		criminal += 0.019
	case r.Doctrines.Base.Omnism:
		accepted += 0.13
		shunned += 0.035
		criminal += 0.018
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.05
		shunned += 0.01
	case r.Doctrines.Gender.Equality:
		accepted += 0.01
		shunned += 0.005
		criminal += 0.003
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.03
		shunned += 0.02
		criminal += 0.01
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
		shunned += 0.01
		criminal += 0.02
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
		accepted += 0.03
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
		accepted += 0.02
		criminal += 0.02
	}
	if r.Doctrines.SanctityOfNature {
		accepted += 0.005
		shunned += 0.003
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		criminal += 0.02
	}
	if r.Doctrines.AncestorWorship {
		accepted += 0.02
		shunned += 0.01
	}
	if r.Doctrines.Pacifism {
		accepted += 0.01
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
		shunned += 0.01
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		accepted += 0.015
		shunned += 0.005
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
	}
	if r.Doctrines.LiveUnderGround {
		shunned += 0.1
		criminal += 0.15
	}
	if r.Doctrines.TreeConnection {
		accepted += 0.01
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
		accepted += 0.01
	}
	if r.Doctrines.Raider {
		accepted += 0.001
	}
	if r.Doctrines.Proselytizer {
		accepted += 0.01
		shunned += 0.015
		criminal += 0.01
	}
	if r.Doctrines.Hedonism {
		accepted += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
		accepted += 0.02
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
		shunned += 0.005
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
		accepted += 0.002
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
		shunned += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
		shunned += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
		accepted += 0.002
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
		accepted += 0.001
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
		shunned += 0.01
		criminal += 0.005
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getUseHallucinogens(r *entities.Religion) religion.Acceptance {
	var (
		accepted = 0.1
		shunned  = 0.15
		criminal = 0.16
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.05
		shunned += 0.06
		criminal += 0.065
	case r.Doctrines.Base.Polytheism:
		accepted += 0.07
		shunned += 0.04
		criminal += 0.04
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.055
		shunned += 0.055
		criminal += 0.045
	case r.Doctrines.Base.Deism:
		accepted += 0.08
		shunned += 0.035
		criminal += 0.035
	case r.Doctrines.Base.Henothism:
		accepted += 0.07
		shunned += 0.045
		criminal += 0.043
	case r.Doctrines.Base.Monolatry:
		accepted += 0.075
		shunned += 0.04
		criminal += 0.038
	case r.Doctrines.Base.Omnism:
		accepted += 0.077
		shunned += 0.039
		criminal += 0.041
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.05
		shunned += 0.06
		criminal += 0.01
	case r.Doctrines.Gender.Equality:
		accepted += 0.01
		shunned += 0.03
		criminal += 0.02
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.025
		shunned += 0.04
		criminal += 0.01
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
		accepted += 0.05
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
		accepted += 0.01
		criminal += 0.02
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		accepted += 0.001
	}
	if r.Doctrines.AncestorWorship {
		accepted += 0.001
	}
	if r.Doctrines.Pacifism {
		accepted += 0.001
	}
	if r.Doctrines.Reincarnation {
		accepted += 0.001
	}
	if r.Doctrines.RitualHospitality {
		accepted += 0.001
	}
	if r.Doctrines.SacredChildbirth {
		shunned += 0.01
		criminal += 0.011
	}
	if r.Doctrines.SanctionedFalseConversions {
		accepted += 0.015
	}
	if r.Doctrines.SunWorship {
	}
	if r.Doctrines.MoonWorship {
		accepted += 0.01
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
		accepted += 0.01
	}
	if r.Doctrines.LiveUnderGround {
		accepted += 0.005
	}
	if r.Doctrines.TreeConnection {
		accepted += 0.005
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
		accepted += 0.01
	}
	if r.Doctrines.Raider {
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
		accepted += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
		shunned += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
		shunned += 0.005
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
		shunned += 0.005
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
		accepted += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

// Amphetamines, Cocaine, etc
func getUseCNSStimulants(r *entities.Religion) religion.Acceptance {
	var (
		accepted = 0.1
		shunned  = 0.15
		criminal = 0.16
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.04
		shunned += 0.05
		criminal += 0.06
	case r.Doctrines.Base.Polytheism:
		accepted += 0.05
		shunned += 0.04
		criminal += 0.0459
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.048
		shunned += 0.02
		criminal += 0.02
	case r.Doctrines.Base.Deism:
		accepted += 0.06
		shunned += 0.05
		criminal += 0.07
	case r.Doctrines.Base.Henothism:
		accepted += 0.055
		shunned += 0.041
		criminal += 0.0455
	case r.Doctrines.Base.Monolatry:
		accepted += 0.05
		shunned += 0.049
		criminal += 0.0469
	case r.Doctrines.Base.Omnism:
		accepted += 0.04
		shunned += 0.049
		criminal += 0.0479
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.02
		shunned += 0.03
		criminal += 0.04
	case r.Doctrines.Gender.Equality:
		accepted += 0.005
		shunned += 0.03
		criminal += 0.03
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.001
		shunned += 0.04
		criminal += 0.05
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
		shunned += 0.01
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
		accepted += 0.01
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
		accepted += 0.01
	}
	if r.Doctrines.ReligiousLaw {
		accepted += 0.01
		criminal += 0.03
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		accepted += 0.01
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
		shunned += 0.01
		criminal += 0.007
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
		shunned += 0.01
		criminal += 0.02
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
		accepted += 0.01
	}
	if r.Doctrines.LiveUnderGround {
		accepted += 0.01
	}
	if r.Doctrines.TreeConnection {
		accepted += 0.005
		shunned += 0.01
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
		accepted += 0.005
	}
	if r.Doctrines.Raider {
		accepted += 0.05
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
		accepted += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
		criminal += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
		shunned += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
		shunned += 0.005
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
		accepted += 0.005
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
		shunned += 0.005
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
		accepted += 0.02
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
		shunned += 0.01
		criminal += 0.02
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
		criminal += 0.001
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
		accepted += 0.001
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getUseOpium(r *entities.Religion) religion.Acceptance {
	var (
		accepted = 0.05
		shunned  = 0.1
		criminal = 0.2
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.01
		shunned += 0.04
		criminal += 0.09
	case r.Doctrines.Base.Polytheism:
		accepted += 0.06
		shunned += 0.07
		criminal += 0.08
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.05
		shunned += 0.04
		criminal += 0.1
	case r.Doctrines.Base.Deism:
		accepted += 0.1
		shunned += 0.09
		criminal += 0.1
	case r.Doctrines.Base.Henothism:
		accepted += 0.055
		shunned += 0.079
		criminal += 0.089
	case r.Doctrines.Base.Monolatry:
		accepted += 0.06
		shunned += 0.071
		criminal += 0.081
	case r.Doctrines.Base.Omnism:
		accepted += 0.067
		shunned += 0.07
		criminal += 0.085
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.05
		shunned += 0.06
		criminal += 0.07
	case r.Doctrines.Gender.Equality:
		accepted += 0.03
		shunned += 0.08
		criminal += 0.08
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.03
		shunned += 0.09
		criminal += 0.1
	}

	if r.Doctrines.FullTolerance {
		accepted += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
		criminal += 0.01
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
		accepted += 0.005
	}
	if r.Doctrines.Legalism {
		criminal += 0.005
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		criminal += 0.01
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
		shunned += 0.01
		criminal += 0.03
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
	}
	if r.Doctrines.MoonWorship {
		accepted += 0.001
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
		accepted += 0.01
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
	}
	if r.Doctrines.Proselytizer {
		criminal += 0.05
	}
	if r.Doctrines.Hedonism {
		accepted += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
		shunned += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
		shunned += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
		criminal += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
		shunned += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
		accepted += 0.001
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getKillingTaboos(r *entities.Religion) {
	r.Taboos.KillingTaboos = &rel.KillingTaboos{}
	r.Taboos.KillingTaboos.Killing = getKilling(r)
	r.Taboos.KillingTaboos.KillAnimals = getKillAnimals(r)
	r.Taboos.KillingTaboos.KillHollyAnimals = getKillHollyAnimals(r)
	r.Taboos.KillingTaboos.KillHumans = getKillHumans(r)
	r.Taboos.KillingTaboos.Kinslaying = getKinslaying(r)
	r.Taboos.KillingTaboos.Suicide = getSuicide(r)
}

func getKilling(r *entities.Religion) religion.Acceptance {
	var (
		accepted = 0.2
		shunned  = 0.16
		criminal = 0.09
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		accepted += 0.1
		shunned += 0.1
		criminal += 0.05
	case r.Doctrines.Base.Polytheism:
		accepted += 0.12
		shunned += 0.18
		criminal += 0.02
	case r.Doctrines.Base.DeityDualism:
		accepted += 0.105
		shunned += 0.11
		criminal += 0.05
	case r.Doctrines.Base.Deism:
		accepted += 0.1
		shunned += 0.1
		criminal += 0.09
	case r.Doctrines.Base.Henothism:
		accepted += 0.11
		shunned += 0.12
		criminal += 0.05
	case r.Doctrines.Base.Monolatry:
		accepted += 0.115
		shunned += 0.1
		criminal += 0.04
	case r.Doctrines.Base.Omnism:
		accepted += 0.1
		shunned += 0.1
		criminal += 0.02
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		accepted += 0.1
		shunned += 0.05
		criminal += 0.05
	case r.Doctrines.Gender.Equality:
		accepted += 0.05
		shunned += 0.05
		criminal += 0.055
	case r.Doctrines.Gender.FemaleDominance:
		accepted += 0.01
		shunned += 0.05
		criminal += 0.1
	}

	if r.Doctrines.FullTolerance {
		criminal += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
		accepted += 0.01
		shunned += 0.005
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
		accepted += 0.01
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
		accepted += 0.02
		criminal += 0.02
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		accepted += 0.05
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
		shunned += 0.05
		criminal += 0.1
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
		shunned += 0.01
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
		accepted += 0.01
	}
	if r.Doctrines.Darkness {
		accepted += 0.05
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
		accepted += 0.1
	}
	if r.Doctrines.Proselytizer {
		accepted += 0.01
	}
	if r.Doctrines.Hedonism {
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
		shunned += 0.005
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
		accepted += 0.03
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
		shunned += 0.01
		criminal += 0.005
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
		accepted += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
		shunned += 0.005
		criminal += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
		accepted += 0.05
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
		shunned += 0.02
		criminal += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
		criminal += 0.05
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
		accepted += 0.05
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

// @todo continue
func getKillAnimals(r *entities.Religion) religion.Acceptance {
	var (
		accepted = 0.2
		shunned  = 0.1
		criminal = 0.03
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
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getKillHollyAnimals(r *entities.Religion) religion.Acceptance {
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
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getKillHumans(r *entities.Religion) religion.Acceptance {
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
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getKinslaying(r *entities.Religion) religion.Acceptance {
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
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getSuicide(r *entities.Religion) religion.Acceptance {
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
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getBodyModificationTaboos(r *entities.Religion) {
	r.Taboos.BodyModificationTaboos = &rel.BodyModificationTaboos{}
	r.Taboos.BodyModificationTaboos.Tattoo = getTattoo(r)
}

func getTattoo(r *entities.Religion) religion.Acceptance {
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
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getWitchcraft(r *entities.Religion) rel.Acceptance {
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
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

func getNudism(r *entities.Religion) rel.Acceptance {
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
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lie.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return geAcceptanceByProbability(accepted, shunned, criminal)
}

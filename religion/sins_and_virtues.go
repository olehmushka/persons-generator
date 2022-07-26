package religion

import (
	"fmt"

	"persons_generator/entities"
	"persons_generator/entities/religion"
	pm "persons_generator/probability_machine"
)

func generateSinsAndVirtues(r *entities.Religion) {
	fmt.Println("[generateSinsAndVirtues] started")
	defer fmt.Println("[rgenerateSinsAndVirtues] finished")

	r.Theology.Precepts.SinsAndVirtues = &religion.SinsAndVirtues{}
	r.Theology.Precepts.SinsAndVirtues.Profanity = getProfanity(r)
	r.Theology.Precepts.SinsAndVirtues.HaveOtherGods = getHaveOtherGods(r)
	r.Theology.Precepts.SinsAndVirtues.Sloth = getSloth(r)
	r.Theology.Precepts.SinsAndVirtues.Greed = getGreed(r)
	r.Theology.Precepts.SinsAndVirtues.Charity = getCharity(r)
	r.Theology.Precepts.SinsAndVirtues.Gluttony = getGluttony(r)
	r.Theology.Precepts.SinsAndVirtues.Temperance = getTemperance(r)
	r.Theology.Precepts.SinsAndVirtues.Pride = getPride(r)
	r.Theology.Precepts.SinsAndVirtues.Humility = getHumility(r)
	r.Theology.Precepts.SinsAndVirtues.Lust = getLust(r)
	r.Theology.Precepts.SinsAndVirtues.Chaste = getChaste(r)
	r.Theology.Precepts.SinsAndVirtues.Wrath = getWrath(r)
	r.Theology.Precepts.SinsAndVirtues.Patience = getPatience(r)
	r.Theology.Precepts.SinsAndVirtues.Pain = getPain(r)
	r.Theology.Precepts.SinsAndVirtues.Sadism = getSadism(r)
	r.Theology.Precepts.SinsAndVirtues.Empathy = getEmpathy(r)
	r.Theology.Precepts.SinsAndVirtues.Stealing = getStealing(r)
	r.Theology.Precepts.SinsAndVirtues.Lie = getLie(r)
	r.Theology.Precepts.SinsAndVirtues.Honest = getHonest(r)
	r.Theology.Precepts.SinsAndVirtues.HonorParents = getHonorParents(r)
}

func getAttitudeByProbability(virtue, neutral, sin float64) religion.Attitude {
	m := map[string]float64{
		string(religion.Virtue):          pm.PrepareProbability(virtue),
		string(religion.NeutralAttitude): pm.PrepareProbability(neutral),
		string(religion.Sin):             pm.PrepareProbability(sin),
	}
	return religion.Attitude(pm.GetRandomFromSeveral(m))
}

func getProfanity(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 0.01
		neutral = 0.05
		sin     = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 0.4
	case r.Doctrines.Base.Polytheism:
		sin += 0.38
	case r.Doctrines.Base.DeityDualism:
		sin += 0.39
	case r.Doctrines.Base.Deism:
		sin += 0.059
		neutral += 0.1
	case r.Doctrines.Base.Henothism:
		sin += 0.19
	case r.Doctrines.Base.Monolatry:
		sin += 0.15
	case r.Doctrines.Base.Omnism:
		neutral += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 0.17
	case r.Doctrines.Gender.Equality:
		sin += 0.152
	case r.Doctrines.Gender.FemaleDominance:
		sin += 0.11
	}

	if r.Doctrines.FullTolerance {
		neutral += 0.12
		virtue += 0.071
	}
	if r.Doctrines.Messiah {
		sin += 0.15
	}
	if r.Doctrines.Prophets {
		sin += 0.169
	}
	if r.Doctrines.Asceticism {
		sin += 0.17
	}
	if r.Doctrines.Legalism {
		sin += 0.12
	}
	if r.Doctrines.ReligiousLaw {
		sin += 0.22
	}
	if r.Doctrines.TaxNonbelievers {
		sin += 0.1999
	}
	if r.Doctrines.UnrelentingFaith {
		sin += 0.1999
	}
	if r.Doctrines.Darkness {
		virtue += 0.091
	}
	if r.Doctrines.Proselytizer {
		sin += 0.3
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getHaveOtherGods(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 0.07
		neutral = 0.1
		sin     = 0.15
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 0.5
	case r.Doctrines.Base.Polytheism:
		neutral += 0.15
		virtue += 0.09
	case r.Doctrines.Base.DeityDualism:
		neutral += 0.1
		sin += 0.05
	case r.Doctrines.Base.Deism:
		virtue += 0.09
		neutral += 0.11
		sin += 0.099
	case r.Doctrines.Base.Henothism:
		neutral += 0.1
		sin += 0.2
	case r.Doctrines.Base.Monolatry:
		virtue += 0.11
		neutral += 0.15
		sin += 0.19
	case r.Doctrines.Base.Omnism:
		virtue += 0.12
		neutral += 0.15
		sin += 0.17
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 0.06
	case r.Doctrines.Gender.Equality:
		sin += 0.02
	case r.Doctrines.Gender.FemaleDominance:
		sin += 0.01
	}

	if r.Doctrines.Messiah {
		sin += 0.1
	}
	if r.Doctrines.Prophets {
		sin += 0.2
	}
	if r.Doctrines.MendicantPreachers {
		sin += 0.1
	}
	if r.Doctrines.Polyamory {
		virtue += 0.11
		neutral += 0.12
	}
	if r.Doctrines.ReligiousLaw {
		sin += 0.15
	}
	if r.Doctrines.TaxNonbelievers {
		sin += 0.11
	}
	if r.Doctrines.UnrelentingFaith {
		sin += 0.3
	}
	if r.Doctrines.AncestorWorship {
		sin += 0.1
	}
	if r.Doctrines.Proselytizer {
		sin += 0.3
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getSloth(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 0.1
		neutral = 0.15
		sin     = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		neutral += 0.035
		sin += 0.1
	case r.Doctrines.Base.Polytheism:
		neutral += 0.1
		sin += 0.03
	case r.Doctrines.Base.DeityDualism:
		sin += 0.1
	case r.Doctrines.Base.Deism:
		neutral += 0.1
	case r.Doctrines.Base.Henothism:
		sin += 0.1
	case r.Doctrines.Base.Monolatry:
		neutral += 0.1
	case r.Doctrines.Base.Omnism:
		neutral += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 0.07
	case r.Doctrines.Gender.Equality:
		sin += 0.04
	case r.Doctrines.Gender.FemaleDominance:
		sin += 0.03
	}

	if r.Doctrines.FullTolerance {
		virtue += 0.1
		neutral += 0.15
	}
	if r.Doctrines.Asceticism {
		sin += 0.1
	}
	if r.Doctrines.Legalism {
		sin += 0.1
	}
	if r.Doctrines.Monasticism {
		sin += 0.1
	}
	if r.Doctrines.Polyamory {
		virtue += 0.1
		neutral += 0.1
	}
	if r.Doctrines.ReligiousLaw {
		neutral += 0.03
		sin += 0.1
	}
	if r.Doctrines.RitualHospitality {
		sin += 0.1
	}
	if r.Doctrines.Hedonism {
		virtue += 0.1
		neutral += 0.07
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getGreed(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 0.100009
		neutral = 0.105
		sin     = 0.18
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 0.19
	case r.Doctrines.Base.Polytheism:
		neutral += 0.07
		sin += 0.12
	case r.Doctrines.Base.DeityDualism:
		sin += 0.19
	case r.Doctrines.Base.Deism:
		neutral += 0.1
		sin += 0.15
	case r.Doctrines.Base.Henothism:
		neutral += 0.1
		sin += 0.16
	case r.Doctrines.Base.Monolatry:
		neutral += 0.1
		sin += 0.145
	case r.Doctrines.Base.Omnism:
		neutral += 0.1
		sin += 0.12
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 0.15
	case r.Doctrines.Gender.Equality:
		sin += 0.2
	case r.Doctrines.Gender.FemaleDominance:
		sin += 0.3
	}

	if r.Doctrines.ReligiousLaw {
		sin += 0.1
	}
	if r.Doctrines.Monasticism {
		sin += 0.15
	}
	if r.Doctrines.RitualHospitality {
		sin += 0.1
	}
	if r.Doctrines.Raider {
		virtue += 0.2
	}
	if r.Doctrines.Hedonism {
		virtue += 0.2
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getCharity(r *entities.Religion) religion.Attitude {
	if r.Theology.Precepts.SinsAndVirtues.Greed == religion.Sin {
		return religion.Virtue
	}
	if r.Theology.Precepts.SinsAndVirtues.Greed == religion.Virtue {
		return religion.Sin
	}

	var (
		virtue  = 0.19
		neutral = 0.12
		sin     = 0.05
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 0.2
	case r.Doctrines.Base.Polytheism:
		virtue += 0.2
	case r.Doctrines.Base.DeityDualism:
		virtue += 0.2
	case r.Doctrines.Base.Deism:
		virtue += 0.2
	case r.Doctrines.Base.Henothism:
		virtue += 0.2
	case r.Doctrines.Base.Monolatry:
		virtue += 0.2
	case r.Doctrines.Base.Omnism:
		virtue += 0.2
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		virtue += 0.1
	case r.Doctrines.Gender.Equality:
		virtue += 0.3
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 0.35
	}

	if r.Doctrines.FullTolerance {
		virtue += 0.2
	}
	if r.Doctrines.Asceticism {
		virtue += 0.1
	}
	if r.Doctrines.Darkness {
		neutral += 0.1
	}
	if r.Doctrines.Raider {
		neutral += 0.25
		sin += 0.3
	}
	if r.Doctrines.Hedonism {
		neutral += 0.15
		sin += 0.1
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getGluttony(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 0.1
		neutral = 0.1
		sin     = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 0.1
	case r.Doctrines.Base.Polytheism:
		neutral += 0.11
		sin += 0.09
	case r.Doctrines.Base.DeityDualism:
		sin += 0.1
	case r.Doctrines.Base.Deism:
		neutral += 0.1
		sin += 0.05
	case r.Doctrines.Base.Henothism:
		neutral += 0.1
		sin += 0.06
	case r.Doctrines.Base.Monolatry:
		neutral += 0.1
		sin += 0.04
	case r.Doctrines.Base.Omnism:
		virtue += 0.01
		neutral += 0.15
		sin += 0.06
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 0.1
	case r.Doctrines.Gender.Equality:
		sin += 0.14
	case r.Doctrines.Gender.FemaleDominance:
		sin += 0.18
	}

	if r.Doctrines.Asceticism {
		sin += 0.3
	}
	if r.Doctrines.Legalism {
		sin += 0.1
	}
	if r.Doctrines.MendicantPreachers {
		sin += 0.1
	}
	if r.Doctrines.Monasticism {
		sin += 0.3
	}
	if r.Doctrines.Polyamory {
		virtue += 0.1
	}
	if r.Doctrines.RitualHospitality {
		neutral += 0.1
		sin += 0.1
	}
	if r.Doctrines.PainIsVirtue {
		sin += 0.1
	}
	if r.Doctrines.Raider {
		virtue += 0.3
		neutral += 0.1
	}
	if r.Doctrines.Hedonism {
		virtue += 0.4
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getTemperance(r *entities.Religion) religion.Attitude {
	if r.Theology.Precepts.SinsAndVirtues.Gluttony == religion.Sin {
		return religion.Virtue
	}
	if r.Theology.Precepts.SinsAndVirtues.Gluttony == religion.Virtue {
		return religion.Sin
	}

	var (
		virtue  = 0.18
		neutral = 0.12
		sin     = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 0.11
	case r.Doctrines.Base.Polytheism:
		neutral += 0.1
	case r.Doctrines.Base.DeityDualism:
		virtue += 0.105
	case r.Doctrines.Base.Deism:
		neutral += 0.101
	case r.Doctrines.Base.Henothism:
		neutral += 0.1
	case r.Doctrines.Base.Monolatry:
		neutral += 0.1
	case r.Doctrines.Base.Omnism:
		neutral += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		virtue += 0.125
	case r.Doctrines.Gender.Equality:
		virtue += 0.12
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 0.11
	}

	if r.Doctrines.Asceticism {
		virtue += 0.1
	}
	if r.Doctrines.Legalism {
		virtue += 0.1
	}
	if r.Doctrines.MendicantPreachers {
		virtue += 0.1
	}
	if r.Doctrines.Monasticism {
		virtue += 0.1
	}
	if r.Doctrines.Polyamory {
		neutral += 0.1
	}
	if r.Doctrines.RitualHospitality {
		virtue += 0.12
		neutral += 0.1
	}
	if r.Doctrines.PainIsVirtue {
		virtue += 0.1
	}
	if r.Doctrines.Raider {
		neutral += 0.045
		sin += 0.1
	}
	if r.Doctrines.Hedonism {
		neutral += 0.04
		sin += 0.1
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getPride(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 0.1
		neutral = 0.11
		sin     = 0.16
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 0.1
		neutral += 0.15
		sin += 0.3
	case r.Doctrines.Base.Polytheism:
		neutral += 0.1
	case r.Doctrines.Base.DeityDualism:
		neutral += 0.1
	case r.Doctrines.Base.Deism:
		neutral += 0.1
	case r.Doctrines.Base.Henothism:
		neutral += 0.1
	case r.Doctrines.Base.Monolatry:
		neutral += 0.1
	case r.Doctrines.Base.Omnism:
		neutral += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 0.09
	case r.Doctrines.Gender.Equality:
		sin += 0.08
	case r.Doctrines.Gender.FemaleDominance:
		neutral += 0.1
		sin += 0.07
	}

	if r.Doctrines.Messiah {
		sin += 0.2
	}
	if r.Doctrines.Prophets {
		sin += 0.1
	}
	if r.Doctrines.Legalism {
		sin += 0.1
	}
	if r.Doctrines.Polyamory {
		neutral += 0.1
	}
	if r.Doctrines.AncestorWorship {
		virtue += 0.1
	}
	if r.Doctrines.Pacifism {
		sin += 0.1
	}
	if r.Doctrines.RitualHospitality {
		sin += 0.15
	}
	if r.Doctrines.SanctionedFalseConversions {
		sin += 0.2
	}
	if r.Doctrines.Proselytizer {
		virtue += 0.1
		neutral += 0.115
		sin += 0.13
	}
	if r.Doctrines.Hedonism {
		virtue += 0.3
		neutral += 0.21
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getHumility(r *entities.Religion) religion.Attitude {
	if r.Theology.Precepts.SinsAndVirtues.Pride == religion.Sin {
		return religion.Virtue
	}
	if r.Theology.Precepts.SinsAndVirtues.Pride == religion.Virtue {
		return religion.Sin
	}

	var (
		virtue  = 0.19
		neutral = 0.11
		sin     = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 0.1
	case r.Doctrines.Base.Polytheism:
		virtue += 0.05
		neutral += 0.12
	case r.Doctrines.Base.DeityDualism:
		virtue += 0.1
	case r.Doctrines.Base.Deism:
		virtue += 0.05
		neutral += 0.14
	case r.Doctrines.Base.Henothism:
		virtue += 0.05
		neutral += 0.12
	case r.Doctrines.Base.Monolatry:
		virtue += 0.05
		neutral += 0.13
	case r.Doctrines.Base.Omnism:
		virtue += 0.05
		neutral += 0.12
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		neutral += 0.05
	case r.Doctrines.Gender.Equality:
		neutral += 0.05
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 0.05
	}

	if r.Doctrines.Asceticism {
		virtue += 0.1
	}
	if r.Doctrines.Legalism {
		virtue += 0.1
	}
	if r.Doctrines.MendicantPreachers {
		virtue += 0.1
	}
	if r.Doctrines.Monasticism {
		virtue += 0.1
	}
	if r.Doctrines.RitualHospitality {
		virtue += 0.1
	}
	if r.Doctrines.SanctionedFalseConversions {
		virtue += 0.1
	}
	if r.Doctrines.Darkness {
		neutral += 0.12
		sin += 0.1
	}
	if r.Doctrines.Raider {
		neutral += 0.09
		sin += 0.1
	}
	if r.Doctrines.Hedonism {
		neutral += 0.11
		sin += 0.1
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getLust(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 0.1
		neutral = 0.12
		sin     = 0.13
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 0.3
	case r.Doctrines.Base.Polytheism:
		virtue += 0.1
		neutral += 0.13
		sin += 0.11
	case r.Doctrines.Base.DeityDualism:
		neutral += 0.1
		sin += 0.13
	case r.Doctrines.Base.Deism:
		neutral += 0.1
	case r.Doctrines.Base.Henothism:
		neutral += 0.1
	case r.Doctrines.Base.Monolatry:
		neutral += 0.1
	case r.Doctrines.Base.Omnism:
		neutral += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 0.15
	case r.Doctrines.Gender.Equality:
		virtue += 0.005
		neutral += 0.01
		sin += 0.011
	case r.Doctrines.Gender.FemaleDominance:
		neutral += 0.15
	}

	if r.Doctrines.FullTolerance {
		virtue += 0.1
		neutral += 0.03
	}
	if r.Doctrines.Prophets {
		sin += 0.1
	}
	if r.Doctrines.Asceticism {
		sin += 0.4
	}
	if r.Doctrines.Legalism {
		sin += 0.1
	}
	if r.Doctrines.Monasticism {
		sin += 0.3
	}
	if r.Doctrines.Polyamory {
		virtue += 0.21
	}
	if r.Doctrines.SanctityOfNature {
		virtue += 0.1
	}
	if r.Doctrines.Pacifism {
		neutral += 0.1
	}
	if r.Doctrines.SacredChildbirth {
		virtue += 0.105
		neutral += 0.12
	}
	if r.Doctrines.Raider {
		virtue += 0.2
		neutral += 0.15
	}
	if r.Doctrines.Hedonism {
		virtue += 0.4
		neutral += 0.205
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getChaste(r *entities.Religion) religion.Attitude {
	if r.Theology.Precepts.SinsAndVirtues.Lust == religion.Sin {
		return religion.Virtue
	}
	if r.Theology.Precepts.SinsAndVirtues.Lust == religion.Virtue {
		return religion.Sin
	}

	var (
		virtue  = 0.15
		neutral = 0.11
		sin     = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 0.1
	case r.Doctrines.Base.Polytheism:
		virtue += 0.1
	case r.Doctrines.Base.DeityDualism:
		virtue += 0.1
	case r.Doctrines.Base.Deism:
		neutral += 0.1
	case r.Doctrines.Base.Henothism:
		neutral += 0.1
	case r.Doctrines.Base.Monolatry:
		neutral += 0.1
	case r.Doctrines.Base.Omnism:
		neutral += 0.1
		virtue += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		virtue += 0.2
	case r.Doctrines.Gender.Equality:
		neutral += 0.1
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 0.1
	}

	if r.Doctrines.FullTolerance {
		neutral += 0.1
	}
	if r.Doctrines.Messiah {
		virtue += 0.1
	}
	if r.Doctrines.Asceticism {
		virtue += 0.4
	}
	if r.Doctrines.Legalism {
		virtue += 0.1
	}
	if r.Doctrines.Monasticism {
		virtue += 0.2
	}
	if r.Doctrines.Polyamory {
		neutral += 0.23
		sin += 0.3
	}
	if r.Doctrines.ReligiousLaw {
		virtue += 0.1
	}
	if r.Doctrines.SacredChildbirth {
		neutral += 0.3
	}
	if r.Doctrines.Raider {
		neutral += 0.1
	}
	if r.Doctrines.Hedonism {
		neutral += 0.15
		sin += 0.3
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getWrath(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 0.1
		neutral = 0.13
		sin     = 0.16
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		neutral += 0.1
		sin += 0.12
	case r.Doctrines.Base.Polytheism:
		neutral += 0.1
	case r.Doctrines.Base.DeityDualism:
		neutral += 0.1
		sin += 0.12
	case r.Doctrines.Base.Deism:
		neutral += 0.1
	case r.Doctrines.Base.Henothism:
		neutral += 0.1
	case r.Doctrines.Base.Monolatry:
		neutral += 0.1
	case r.Doctrines.Base.Omnism:
		neutral += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		neutral += 0.1
		sin += 0.03
	case r.Doctrines.Gender.Equality:
		sin += 0.07
	case r.Doctrines.Gender.FemaleDominance:
		sin += 0.1
	}

	if r.Doctrines.Asceticism {
		sin += 0.1
	}
	if r.Doctrines.Esotericism {
		sin += 0.1
	}
	if r.Doctrines.Polyamory {
		sin += 0.1
	}
	if r.Doctrines.UnrelentingFaith {
		neutral += 0.1
	}
	if r.Doctrines.Pacifism {
		sin += 0.4
	}
	if r.Doctrines.SacredChildbirth {
		sin += 0.1
	}
	if r.Doctrines.Darkness {
		virtue += 0.1
	}
	if r.Doctrines.PainIsVirtue {
		virtue += 0.1
		neutral += 0.05
	}
	if r.Doctrines.Raider {
		virtue += 0.5
	}
	if r.Doctrines.Proselytizer {
		neutral += 0.1
	}
	if r.Doctrines.Hedonism {
		neutral += 0.1
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getPatience(r *entities.Religion) religion.Attitude {
	if r.Theology.Precepts.SinsAndVirtues.Wrath == religion.Sin {
		return religion.Virtue
	}
	if r.Theology.Precepts.SinsAndVirtues.Wrath == religion.Virtue {
		return religion.Sin
	}

	var (
		virtue  = 0.15
		neutral = 0.12
		sin     = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 0.12
	case r.Doctrines.Base.Polytheism:
		virtue += 0.105
	case r.Doctrines.Base.DeityDualism:
		virtue += 0.11
	case r.Doctrines.Base.Deism:
		virtue += 0.103
	case r.Doctrines.Base.Henothism:
		virtue += 0.106
	case r.Doctrines.Base.Monolatry:
		virtue += 0.1055
	case r.Doctrines.Base.Omnism:
		virtue += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		virtue += 0.05
	case r.Doctrines.Gender.Equality:
		virtue += 0.06
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 0.08
	}

	if r.Doctrines.Asceticism {
		virtue += 0.3
	}
	if r.Doctrines.Astrology {
		virtue += 0.1
	}
	if r.Doctrines.Pacifism {
		virtue += 0.4
	}
	if r.Doctrines.PainIsVirtue {
		virtue += 0.2
	}
	if r.Doctrines.Raider {
		sin += 0.3
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getPain(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 0.1
		neutral = 0.15
		sin     = 0.11
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		neutral += 0.1
	case r.Doctrines.Base.Polytheism:
		neutral += 0.1
	case r.Doctrines.Base.DeityDualism:
		neutral += 0.1
	case r.Doctrines.Base.Deism:
		neutral += 0.1
	case r.Doctrines.Base.Henothism:
		neutral += 0.1
	case r.Doctrines.Base.Monolatry:
		neutral += 0.1
	case r.Doctrines.Base.Omnism:
		neutral += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		neutral += 0.05
	case r.Doctrines.Gender.Equality:
		neutral += 0.05
	case r.Doctrines.Gender.FemaleDominance:
		sin += 0.05
	}

	if r.Doctrines.PainIsVirtue {
		virtue += 0.6
	}
	if r.Doctrines.Hedonism {
		sin += 0.2
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getSadism(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 0.1
		neutral = 0.12
		sin     = 0.199
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 0.2
	case r.Doctrines.Base.Polytheism:
		sin += 0.2
	case r.Doctrines.Base.DeityDualism:
		sin += 0.2
	case r.Doctrines.Base.Deism:
		neutral += 0.09
		sin += 0.19
	case r.Doctrines.Base.Henothism:
		sin += 0.2
	case r.Doctrines.Base.Monolatry:
		sin += 0.2
	case r.Doctrines.Base.Omnism:
		sin += 0.2
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 0.1
	case r.Doctrines.Gender.Equality:
		sin += 0.15
	case r.Doctrines.Gender.FemaleDominance:
		sin += 0.18
	}

	if r.Doctrines.FullTolerance {
		sin += 0.2
	}
	if r.Doctrines.Legalism {
		sin += 0.1
	}
	if r.Doctrines.Polyamory {
		sin += 0.205
	}
	if r.Doctrines.ReligiousLaw {
		sin += 0.1
	}
	if r.Doctrines.SanctityOfNature {
		sin += 0.1
	}
	if r.Doctrines.Pacifism {
		sin += 0.4
	}
	if r.Doctrines.Reincarnation {
		sin += 0.1
	}
	if r.Doctrines.RitualHospitality {
		sin += 0.2
	}
	if r.Doctrines.PainIsVirtue {
		virtue += 0.3
		neutral += 0.11
	}
	if r.Doctrines.Darkness {
		virtue += 0.3
		neutral += 0.1
	}
	if r.Doctrines.TreeConnection {
		sin += 0.1
	}
	if r.Doctrines.AnimalConnection {
		sin += 0.15
	}
	if r.Doctrines.Raider {
		virtue += 0.25
		neutral += 0.15
	}
	if r.Doctrines.Hedonism {
		neutral += 0.1
		sin += 0.01
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getEmpathy(r *entities.Religion) religion.Attitude {
	if r.Theology.Precepts.SinsAndVirtues.Sadism == religion.Sin {
		return religion.Virtue
	}
	if r.Theology.Precepts.SinsAndVirtues.Sadism == religion.Virtue {
		return religion.Sin
	}

	var (
		virtue  = 0.19
		neutral = 0.12
		sin     = 0.1
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 0.1
	case r.Doctrines.Base.Polytheism:
		virtue += 0.1
	case r.Doctrines.Base.DeityDualism:
		virtue += 0.1
	case r.Doctrines.Base.Deism:
		virtue += 0.1
	case r.Doctrines.Base.Henothism:
		virtue += 0.1
	case r.Doctrines.Base.Monolatry:
		virtue += 0.1
	case r.Doctrines.Base.Omnism:
		virtue += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		virtue += 0.1
	case r.Doctrines.Gender.Equality:
		virtue += 0.15
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 0.2
	}

	if r.Doctrines.FullTolerance {
		virtue += 0.3
	}
	if r.Doctrines.Legalism {
		virtue += 0.1
	}
	if r.Doctrines.Polyamory {
		virtue += 0.2
	}
	if r.Doctrines.ReligiousLaw {
		virtue += 0.1
	}
	if r.Doctrines.SanctityOfNature {
		virtue += 0.3
	}
	if r.Doctrines.Pacifism {
		virtue += 0.309
	}
	if r.Doctrines.Reincarnation {
		virtue += 0.1
	}
	if r.Doctrines.RitualHospitality {
		virtue += 0.309
	}
	if r.Doctrines.PainIsVirtue {
		neutral += 0.1
	}
	if r.Doctrines.Darkness {
		neutral += 0.1
	}
	if r.Doctrines.TreeConnection {
		virtue += 0.1
	}
	if r.Doctrines.AnimalConnection {
		virtue += 0.19
	}
	if r.Doctrines.Raider {
		neutral += 0.1
		sin += 0.2
	}
	if r.Doctrines.Hedonism {
		neutral += 0.1
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getStealing(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 0.1
		neutral = 0.13
		sin     = 0.18
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 0.1
	case r.Doctrines.Base.Polytheism:
		sin += 0.1
	case r.Doctrines.Base.DeityDualism:
		sin += 0.1
	case r.Doctrines.Base.Deism:
		neutral += 0.1
	case r.Doctrines.Base.Henothism:
		sin += 0.1
	case r.Doctrines.Base.Monolatry:
		sin += 0.1
	case r.Doctrines.Base.Omnism:
		sin += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 0.2
	case r.Doctrines.Gender.Equality:
		sin += 0.15
	case r.Doctrines.Gender.FemaleDominance:
		sin += 0.109
	}

	if r.Doctrines.Prophets {
		sin += 0.1
	}
	if r.Doctrines.Legalism {
		sin += 0.2
	}
	if r.Doctrines.ReligiousLaw {
		sin += 0.29
	}
	if r.Doctrines.TaxNonbelievers {
		sin += 0.1
	}
	if r.Doctrines.UnrelentingFaith {
		sin += 0.1
	}
	if r.Doctrines.Darkness {
		virtue += 0.3
		neutral += 0.12
	}
	if r.Doctrines.LiveUnderGround {
		virtue += 0.1
		neutral += 0.05
	}
	if r.Doctrines.Hedonism {
		neutral += 0.1
		sin += 0.05
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getLie(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 0.11
		neutral = 0.12
		sin     = 0.16
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 0.1
	case r.Doctrines.Base.Polytheism:
		neutral += 0.1
	case r.Doctrines.Base.DeityDualism:
		sin += 0.1
	case r.Doctrines.Base.Deism:
		neutral += 0.1
	case r.Doctrines.Base.Henothism:
		neutral += 0.1
	case r.Doctrines.Base.Monolatry:
		neutral += 0.1
	case r.Doctrines.Base.Omnism:
		neutral += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 0.2
	case r.Doctrines.Gender.Equality:
		sin += 0.15
	case r.Doctrines.Gender.FemaleDominance:
		sin += 0.1109
	}

	if r.Doctrines.Prophets {
		sin += 0.1
	}
	if r.Doctrines.Monasticism {
		sin += 0.1
	}
	if r.Doctrines.Legalism {
		sin += 0.2
	}
	if r.Doctrines.Polyamory {
		virtue += 0.01
		neutral += 0.1
	}
	if r.Doctrines.ReligiousLaw {
		sin += 0.3
	}
	if r.Doctrines.Pacifism {
		neutral += 0.1
	}
	if r.Doctrines.RitualHospitality {
		neutral += 0.1
	}
	if r.Doctrines.Darkness {
		virtue += 0.3
	}
	if r.Doctrines.Raider {
		sin += 0.1
	}
	if r.Doctrines.Hedonism {
		neutral += 0.1
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getHonest(r *entities.Religion) religion.Attitude {
	if r.Theology.Precepts.SinsAndVirtues.Lie == religion.Sin {
		return religion.Virtue
	}
	if r.Theology.Precepts.SinsAndVirtues.Lie == religion.Virtue {
		return religion.Sin
	}

	var (
		virtue  = 0.19
		neutral = 0.14
		sin     = 0.101
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 0.2
	case r.Doctrines.Base.Polytheism:
		virtue += 0.15
	case r.Doctrines.Base.DeityDualism:
		virtue += 0.2
	case r.Doctrines.Base.Deism:
		virtue += 0.105
	case r.Doctrines.Base.Henothism:
		virtue += 0.16
	case r.Doctrines.Base.Monolatry:
		virtue += 0.14
	case r.Doctrines.Base.Omnism:
		virtue += 0.13
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		virtue += 0.2
	case r.Doctrines.Gender.Equality:
		virtue += 0.2
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 0.18
	}

	if r.Doctrines.Prophets {
		virtue += 0.1
	}
	if r.Doctrines.Legalism {
		virtue += 0.3
	}
	if r.Doctrines.Polyamory {
		virtue += 0.1
		neutral += 0.085
	}
	if r.Doctrines.ReligiousLaw {
		virtue += 0.3
	}
	if r.Doctrines.Pacifism {
		virtue += 0.1
	}
	if r.Doctrines.RitualHospitality {
		virtue += 0.1
	}
	if r.Doctrines.Darkness {
		neutral += 0.1
	}
	if r.Doctrines.Raider {
		virtue += 0.1
	}
	if r.Doctrines.Hedonism {
		virtue += 0.1
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getHonorParents(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 0.18
		neutral = 0.14
		sin     = 0.101
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 0.19
	case r.Doctrines.Base.Polytheism:
		neutral += 0.1
	case r.Doctrines.Base.DeityDualism:
		virtue += 0.1
	case r.Doctrines.Base.Deism:
		neutral += 0.1
	case r.Doctrines.Base.Henothism:
		virtue += 0.1
	case r.Doctrines.Base.Monolatry:
		virtue += 0.1
	case r.Doctrines.Base.Omnism:
		neutral += 0.1
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		virtue += 0.25
	case r.Doctrines.Gender.Equality:
		virtue += 0.2
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 0.15
	}

	if r.Doctrines.Prophets {
		virtue += 0.105
	}
	if r.Doctrines.Legalism {
		virtue += 0.105
	}
	if r.Doctrines.Polyamory {
		virtue += 0.1
	}
	if r.Doctrines.ReligiousLaw {
		virtue += 0.3
	}
	if r.Doctrines.AncestorWorship {
		virtue += 0.6
	}
	if r.Doctrines.Reincarnation {
		virtue += 0.15
	}
	if r.Doctrines.RitualHospitality {
		virtue += 0.1
	}
	if r.Doctrines.Darkness {
		neutral += 0.1
	}
	if r.Doctrines.Hedonism {
		neutral += 0.1
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

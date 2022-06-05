package religion

import (
	"persons_generator/entities"
	"persons_generator/entities/religion"
	pm "persons_generator/probability-machine"
)

func generateSinsAndVirtues(r *entities.Religion) {
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

func getAttitudeByProbability(virtue, neutral, sin int) religion.Attitude {
	m := map[string]int{
		string(religion.Virtue):          pm.PrepareProbability(virtue),
		string(religion.NeutralAttitude): pm.PrepareProbability(neutral),
		string(religion.Sin):             pm.PrepareProbability(sin),
	}
	return religion.Attitude(pm.GetRandomFromSeveral(m))
}

func getProfanity(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 40
	case r.Doctrines.Base.Polytheism:
		sin += 40
	case r.Doctrines.Base.DeityDualism:
		sin += 40
	case r.Doctrines.Base.Deism:
		neutral += 10
	case r.Doctrines.Base.Henothism:
		sin += 10
	case r.Doctrines.Base.Monolatry:
		sin += 10
	case r.Doctrines.Base.Omnism:
		neutral += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 7
	case r.Doctrines.Gender.Equality:
		sin += 2
	case r.Doctrines.Gender.FemaleDominance:
		sin += 1
	}

	if r.Doctrines.FullTolerance {
		neutral += 10
		virtue += 10
	}
	if r.Doctrines.Messiah {
		sin += 10
	}
	if r.Doctrines.Prophets {
		sin += 10
	}
	if r.Doctrines.Asceticism {
		sin += 10
	}
	if r.Doctrines.Legalism {
		sin += 10
	}
	if r.Doctrines.ReligiousLaw {
		sin += 20
	}
	if r.Doctrines.TaxNonbelievers {
		sin += 10
	}
	if r.Doctrines.UnrelentingFaith {
		sin += 10
	}
	if r.Doctrines.Darkness {
		virtue += 10
	}
	if r.Doctrines.Proselytizer {
		sin += 30
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getHaveOtherGods(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 50
	case r.Doctrines.Base.Polytheism:
		neutral += 10
		virtue += 10
	case r.Doctrines.Base.DeityDualism:
		neutral += 10
	case r.Doctrines.Base.Deism:
		neutral += 10
		virtue += 10
	case r.Doctrines.Base.Henothism:
		sin += 20
		neutral += 10
	case r.Doctrines.Base.Monolatry:
		neutral += 10
		virtue += 10
	case r.Doctrines.Base.Omnism:
		virtue += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 6
	case r.Doctrines.Gender.Equality:
		sin += 2
	case r.Doctrines.Gender.FemaleDominance:
		sin += 1
	}

	if r.Doctrines.Messiah {
		sin += 10
	}
	if r.Doctrines.Prophets {
		sin += 20
	}
	if r.Doctrines.MendicantPreachers {
		sin += 10
	}
	if r.Doctrines.Polyamory {
		virtue += 10
		neutral += 10
	}
	if r.Doctrines.ReligiousLaw {
		sin += 10
	}
	if r.Doctrines.TaxNonbelievers {
		sin += 10
	}
	if r.Doctrines.UnrelentingFaith {
		sin += 30
	}
	if r.Doctrines.AncestorWorship {
		sin += 10
	}
	if r.Doctrines.Proselytizer {
		sin += 30
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getSloth(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 10
	case r.Doctrines.Base.Polytheism:
		neutral += 10
	case r.Doctrines.Base.DeityDualism:
		sin += 10
	case r.Doctrines.Base.Deism:
		neutral += 10
	case r.Doctrines.Base.Henothism:
		sin += 10
	case r.Doctrines.Base.Monolatry:
		neutral += 10
	case r.Doctrines.Base.Omnism:
		neutral += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 7
	case r.Doctrines.Gender.Equality:
		sin += 4
	case r.Doctrines.Gender.FemaleDominance:
		sin += 3
	}

	if r.Doctrines.FullTolerance {
		neutral += 10
		virtue += 10
	}
	if r.Doctrines.Asceticism {
		sin += 10
	}
	if r.Doctrines.Legalism {
		sin += 10
	}
	if r.Doctrines.Monasticism {
		sin += 10
	}
	if r.Doctrines.Polyamory {
		neutral += 10
		virtue += 10
	}
	if r.Doctrines.ReligiousLaw {
		sin += 10
	}
	if r.Doctrines.RitualHospitality {
		sin += 10
	}
	if r.Doctrines.Hedonism {
		virtue += 10
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getGreed(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 10
	case r.Doctrines.Base.Polytheism:
		neutral += 10
	case r.Doctrines.Base.DeityDualism:
		sin += 10
	case r.Doctrines.Base.Deism:
		neutral += 10
	case r.Doctrines.Base.Henothism:
		neutral += 10
	case r.Doctrines.Base.Monolatry:
		neutral += 10
	case r.Doctrines.Base.Omnism:
		neutral += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 15
	case r.Doctrines.Gender.Equality:
		sin += 20
	case r.Doctrines.Gender.FemaleDominance:
		sin += 30
	}

	if r.Doctrines.ReligiousLaw {
		sin += 10
	}
	if r.Doctrines.Monasticism {
		sin += 10
	}
	if r.Doctrines.RitualHospitality {
		sin += 10
	}
	if r.Doctrines.Raider {
		virtue += 20
	}
	if r.Doctrines.Hedonism {
		virtue += 20
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
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 20
	case r.Doctrines.Base.Polytheism:
		virtue += 20
	case r.Doctrines.Base.DeityDualism:
		virtue += 20
	case r.Doctrines.Base.Deism:
		virtue += 20
	case r.Doctrines.Base.Henothism:
		virtue += 20
	case r.Doctrines.Base.Monolatry:
		virtue += 20
	case r.Doctrines.Base.Omnism:
		virtue += 20
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		virtue += 10
	case r.Doctrines.Gender.Equality:
		virtue += 30
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 35
	}

	if r.Doctrines.FullTolerance {
		virtue += 20
	}
	if r.Doctrines.Asceticism {
		virtue += 10
	}
	if r.Doctrines.Darkness {
		neutral += 10
	}
	if r.Doctrines.Raider {
		sin += 30
	}
	if r.Doctrines.Hedonism {
		sin += 10
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getGluttony(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 10
	case r.Doctrines.Base.Polytheism:
		neutral += 10
	case r.Doctrines.Base.DeityDualism:
		sin += 10
	case r.Doctrines.Base.Deism:
		neutral += 10
	case r.Doctrines.Base.Henothism:
		neutral += 10
	case r.Doctrines.Base.Monolatry:
		neutral += 10
	case r.Doctrines.Base.Omnism:
		neutral += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 10
	case r.Doctrines.Gender.Equality:
		sin += 14
	case r.Doctrines.Gender.FemaleDominance:
		sin += 18
	}

	if r.Doctrines.Asceticism {
		sin += 30
	}
	if r.Doctrines.Legalism {
		sin += 10
	}
	if r.Doctrines.MendicantPreachers {
		sin += 10
	}
	if r.Doctrines.Monasticism {
		sin += 30
	}
	if r.Doctrines.Polyamory {
		virtue += 10
	}
	if r.Doctrines.RitualHospitality {
		neutral += 10
		sin += 10
	}
	if r.Doctrines.PainIsVirtue {
		sin += 10
	}
	if r.Doctrines.Raider {
		virtue += 30
	}
	if r.Doctrines.Hedonism {
		virtue += 40
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
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 10
	case r.Doctrines.Base.Polytheism:
		neutral += 10
	case r.Doctrines.Base.DeityDualism:
		virtue += 10
	case r.Doctrines.Base.Deism:
		neutral += 10
	case r.Doctrines.Base.Henothism:
		neutral += 10
	case r.Doctrines.Base.Monolatry:
		neutral += 10
	case r.Doctrines.Base.Omnism:
		neutral += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		virtue += 12
	case r.Doctrines.Gender.Equality:
		virtue += 12
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 11
	}

	if r.Doctrines.Asceticism {
		virtue += 10
	}
	if r.Doctrines.Legalism {
		virtue += 10
	}
	if r.Doctrines.MendicantPreachers {
		virtue += 10
	}
	if r.Doctrines.Monasticism {
		virtue += 10
	}
	if r.Doctrines.Polyamory {
		neutral += 10
	}
	if r.Doctrines.RitualHospitality {
		neutral += 10
		virtue += 10
	}
	if r.Doctrines.PainIsVirtue {
		virtue += 10
	}
	if r.Doctrines.Raider {
		sin += 10
	}
	if r.Doctrines.Hedonism {
		sin += 10
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getPride(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 10
		neutral += 10
		sin += 30
	case r.Doctrines.Base.Polytheism:
		neutral += 10
	case r.Doctrines.Base.DeityDualism:
		neutral += 10
	case r.Doctrines.Base.Deism:
		neutral += 10
	case r.Doctrines.Base.Henothism:
		neutral += 10
	case r.Doctrines.Base.Monolatry:
		neutral += 10
	case r.Doctrines.Base.Omnism:
		neutral += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 9
	case r.Doctrines.Gender.Equality:
		sin += 8
	case r.Doctrines.Gender.FemaleDominance:
		neutral += 10
	}

	if r.Doctrines.Messiah {
		sin += 20
	}
	if r.Doctrines.Prophets {
		sin += 10
	}
	if r.Doctrines.Legalism {
		sin += 10
	}
	if r.Doctrines.Polyamory {
		neutral += 10
	}
	if r.Doctrines.AncestorWorship {
		virtue += 10
	}
	if r.Doctrines.Pacifism {
		sin += 10
	}
	if r.Doctrines.RitualHospitality {
		sin += 15
	}
	if r.Doctrines.SanctionedFalseConversions {
		sin += 20
	}
	if r.Doctrines.Proselytizer {
		virtue += 10
		sin += 10
		neutral += 10
	}
	if r.Doctrines.Hedonism {
		virtue += 30
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
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 10
	case r.Doctrines.Base.Polytheism:
		neutral += 10
	case r.Doctrines.Base.DeityDualism:
		virtue += 10
	case r.Doctrines.Base.Deism:
		neutral += 10
	case r.Doctrines.Base.Henothism:
		neutral += 10
	case r.Doctrines.Base.Monolatry:
		neutral += 10
	case r.Doctrines.Base.Omnism:
		neutral += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		neutral += 5
	case r.Doctrines.Gender.Equality:
		neutral += 5
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 5
	}

	if r.Doctrines.Asceticism {
		virtue += 10
	}
	if r.Doctrines.Legalism {
		virtue += 10
	}
	if r.Doctrines.MendicantPreachers {
		virtue += 10
	}
	if r.Doctrines.Monasticism {
		virtue += 10
	}
	if r.Doctrines.RitualHospitality {
		virtue += 10
	}
	if r.Doctrines.SanctionedFalseConversions {
		virtue += 10
	}
	if r.Doctrines.Darkness {
		sin += 10
	}
	if r.Doctrines.Raider {
		sin += 10
	}
	if r.Doctrines.Hedonism {
		sin += 10
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getLust(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 30
	case r.Doctrines.Base.Polytheism:
		neutral += 10
		virtue += 10
	case r.Doctrines.Base.DeityDualism:
		sin += 10
	case r.Doctrines.Base.Deism:
		neutral += 10
	case r.Doctrines.Base.Henothism:
		neutral += 10
	case r.Doctrines.Base.Monolatry:
		neutral += 10
	case r.Doctrines.Base.Omnism:
		neutral += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 15
	case r.Doctrines.Gender.Equality:
	case r.Doctrines.Gender.FemaleDominance:
		neutral += 15
	}

	if r.Doctrines.FullTolerance {
		virtue += 10
	}
	if r.Doctrines.Prophets {
		sin += 10
	}
	if r.Doctrines.Asceticism {
		sin += 40
	}
	if r.Doctrines.Legalism {
		sin += 10
	}
	if r.Doctrines.Monasticism {
		sin += 30
	}
	if r.Doctrines.Polyamory {
		virtue += 10
	}
	if r.Doctrines.SanctityOfNature {
		virtue += 10
	}
	if r.Doctrines.Pacifism {
		neutral += 10
	}
	if r.Doctrines.SacredChildbirth {
		virtue += 10
	}
	if r.Doctrines.Raider {
		virtue += 20
	}
	if r.Doctrines.Hedonism {
		virtue += 40
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
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 10
	case r.Doctrines.Base.Polytheism:
		virtue += 10
	case r.Doctrines.Base.DeityDualism:
		virtue += 10
	case r.Doctrines.Base.Deism:
		neutral += 10
	case r.Doctrines.Base.Henothism:
		neutral += 10
	case r.Doctrines.Base.Monolatry:
		neutral += 10
	case r.Doctrines.Base.Omnism:
		neutral += 10
		virtue += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		virtue += 20
	case r.Doctrines.Gender.Equality:
		neutral += 10
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 10
	}

	if r.Doctrines.FullTolerance {
		neutral += 10
	}
	if r.Doctrines.Messiah {
		virtue += 10
	}
	if r.Doctrines.Asceticism {
		virtue += 40
	}
	if r.Doctrines.Legalism {
		virtue += 10
	}
	if r.Doctrines.Monasticism {
		virtue += 20
	}
	if r.Doctrines.Polyamory {
		sin += 30
	}
	if r.Doctrines.ReligiousLaw {
		virtue += 10
	}
	if r.Doctrines.SacredChildbirth {
		neutral += 30
	}
	if r.Doctrines.Raider {
		neutral += 10
	}
	if r.Doctrines.Hedonism {
		sin += 30
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getWrath(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		neutral += 10
		sin += 10
	case r.Doctrines.Base.Polytheism:
		neutral += 10
	case r.Doctrines.Base.DeityDualism:
		neutral += 10
		sin += 10
	case r.Doctrines.Base.Deism:
		neutral += 10
	case r.Doctrines.Base.Henothism:
		neutral += 10
	case r.Doctrines.Base.Monolatry:
		neutral += 10
	case r.Doctrines.Base.Omnism:
		neutral += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		neutral += 10
	case r.Doctrines.Gender.Equality:
		sin += 7
	case r.Doctrines.Gender.FemaleDominance:
		sin += 10
	}

	if r.Doctrines.Asceticism {
		sin += 10
	}
	if r.Doctrines.Esotericism {
		sin += 10
	}
	if r.Doctrines.Polyamory {
		sin += 10
	}
	if r.Doctrines.UnrelentingFaith {
		neutral += 10
	}
	if r.Doctrines.Pacifism {
		sin += 40
	}
	if r.Doctrines.SacredChildbirth {
		sin += 10
	}
	if r.Doctrines.Darkness {
		virtue += 10
	}
	if r.Doctrines.PainIsVirtue {
		virtue += 10
	}
	if r.Doctrines.Raider {
		virtue += 50
	}
	if r.Doctrines.Proselytizer {
		neutral += 10
	}
	if r.Doctrines.Hedonism {
		neutral += 10
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
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 10
	case r.Doctrines.Base.Polytheism:
		virtue += 10
	case r.Doctrines.Base.DeityDualism:
		virtue += 10
	case r.Doctrines.Base.Deism:
		virtue += 10
	case r.Doctrines.Base.Henothism:
		virtue += 10
	case r.Doctrines.Base.Monolatry:
		virtue += 10
	case r.Doctrines.Base.Omnism:
		virtue += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		virtue += 5
	case r.Doctrines.Gender.Equality:
		virtue += 6
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 8
	}

	if r.Doctrines.Asceticism {
		virtue += 30
	}
	if r.Doctrines.Astrology {
		virtue += 10
	}
	if r.Doctrines.Pacifism {
		virtue += 40
	}
	if r.Doctrines.PainIsVirtue {
		virtue += 20
	}
	if r.Doctrines.Raider {
		sin += 30
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getPain(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		neutral += 10
	case r.Doctrines.Base.Polytheism:
		neutral += 10
	case r.Doctrines.Base.DeityDualism:
		neutral += 10
	case r.Doctrines.Base.Deism:
		neutral += 10
	case r.Doctrines.Base.Henothism:
		neutral += 10
	case r.Doctrines.Base.Monolatry:
		neutral += 10
	case r.Doctrines.Base.Omnism:
		neutral += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		neutral += 5
	case r.Doctrines.Gender.Equality:
		neutral += 5
	case r.Doctrines.Gender.FemaleDominance:
		sin += 5
	}

	if r.Doctrines.PainIsVirtue {
		virtue += 60
	}
	if r.Doctrines.Hedonism {
		sin += 20
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getSadism(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 20
	case r.Doctrines.Base.Polytheism:
		sin += 20
	case r.Doctrines.Base.DeityDualism:
		sin += 20
	case r.Doctrines.Base.Deism:
		sin += 20
	case r.Doctrines.Base.Henothism:
		sin += 20
	case r.Doctrines.Base.Monolatry:
		sin += 20
	case r.Doctrines.Base.Omnism:
		sin += 20
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 10
	case r.Doctrines.Gender.Equality:
		sin += 15
	case r.Doctrines.Gender.FemaleDominance:
		sin += 18
	}

	if r.Doctrines.FullTolerance {
		sin += 20
	}
	if r.Doctrines.Legalism {
		sin += 10
	}
	if r.Doctrines.Polyamory {
		sin += 20
	}
	if r.Doctrines.ReligiousLaw {
		sin += 10
	}
	if r.Doctrines.SanctityOfNature {
		sin += 10
	}
	if r.Doctrines.Pacifism {
		sin += 40
	}
	if r.Doctrines.Reincarnation {
		sin += 10
	}
	if r.Doctrines.RitualHospitality {
		sin += 20
	}
	if r.Doctrines.PainIsVirtue {
		virtue += 30
	}
	if r.Doctrines.Darkness {
		virtue += 30
	}
	if r.Doctrines.TreeConnection {
		sin += 10
	}
	if r.Doctrines.AnimalConnection {
		sin += 10
	}
	if r.Doctrines.Raider {
		virtue += 30
	}
	if r.Doctrines.Hedonism {
		neutral += 10
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
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 10
	case r.Doctrines.Base.Polytheism:
		virtue += 10
	case r.Doctrines.Base.DeityDualism:
		virtue += 10
	case r.Doctrines.Base.Deism:
		virtue += 10
	case r.Doctrines.Base.Henothism:
		virtue += 10
	case r.Doctrines.Base.Monolatry:
		virtue += 10
	case r.Doctrines.Base.Omnism:
		virtue += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		virtue += 10
	case r.Doctrines.Gender.Equality:
		virtue += 15
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 20
	}

	if r.Doctrines.FullTolerance {
		virtue += 30
	}
	if r.Doctrines.Legalism {
		virtue += 10
	}
	if r.Doctrines.Polyamory {
		virtue += 20
	}
	if r.Doctrines.ReligiousLaw {
		virtue += 10
	}
	if r.Doctrines.SanctityOfNature {
		virtue += 30
	}
	if r.Doctrines.Pacifism {
		virtue += 30
	}
	if r.Doctrines.Reincarnation {
		virtue += 10
	}
	if r.Doctrines.RitualHospitality {
		virtue += 30
	}
	if r.Doctrines.PainIsVirtue {
		neutral += 10
	}
	if r.Doctrines.Darkness {
		neutral += 10
	}
	if r.Doctrines.TreeConnection {
		virtue += 10
	}
	if r.Doctrines.AnimalConnection {
		virtue += 10
	}
	if r.Doctrines.Raider {
		sin += 20
	}
	if r.Doctrines.Hedonism {
		neutral += 10
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getStealing(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 10
	case r.Doctrines.Base.Polytheism:
		sin += 10
	case r.Doctrines.Base.DeityDualism:
		sin += 10
	case r.Doctrines.Base.Deism:
		neutral += 10
	case r.Doctrines.Base.Henothism:
		sin += 10
	case r.Doctrines.Base.Monolatry:
		sin += 10
	case r.Doctrines.Base.Omnism:
		sin += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 20
	case r.Doctrines.Gender.Equality:
		sin += 15
	case r.Doctrines.Gender.FemaleDominance:
		sin += 10
	}

	if r.Doctrines.Prophets {
		sin += 10
	}
	if r.Doctrines.Legalism {
		sin += 20
	}
	if r.Doctrines.ReligiousLaw {
		sin += 20
	}
	if r.Doctrines.TaxNonbelievers {
		sin += 10
	}
	if r.Doctrines.UnrelentingFaith {
		sin += 10
	}
	if r.Doctrines.Darkness {
		virtue += 30
	}
	if r.Doctrines.LiveUnderGround {
		virtue += 10
	}
	if r.Doctrines.Hedonism {
		neutral += 10
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getLie(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		sin += 10
	case r.Doctrines.Base.Polytheism:
		neutral += 10
	case r.Doctrines.Base.DeityDualism:
		sin += 10
	case r.Doctrines.Base.Deism:
		neutral += 10
	case r.Doctrines.Base.Henothism:
		neutral += 10
	case r.Doctrines.Base.Monolatry:
		neutral += 10
	case r.Doctrines.Base.Omnism:
		neutral += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		sin += 20
	case r.Doctrines.Gender.Equality:
		sin += 15
	case r.Doctrines.Gender.FemaleDominance:
		sin += 10
	}

	if r.Doctrines.Prophets {
		sin += 10
	}
	if r.Doctrines.Monasticism {
		sin += 10
	}
	if r.Doctrines.Legalism {
		sin += 20
	}
	if r.Doctrines.Polyamory {
		neutral += 10
	}
	if r.Doctrines.ReligiousLaw {
		sin += 30
	}
	if r.Doctrines.Pacifism {
		neutral += 10
	}
	if r.Doctrines.RitualHospitality {
		neutral += 10
	}
	if r.Doctrines.Darkness {
		virtue += 30
	}
	if r.Doctrines.Raider {
		sin += 10
	}
	if r.Doctrines.Hedonism {
		neutral += 10
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
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 20
	case r.Doctrines.Base.Polytheism:
		virtue += 10
	case r.Doctrines.Base.DeityDualism:
		virtue += 20
	case r.Doctrines.Base.Deism:
		virtue += 10
	case r.Doctrines.Base.Henothism:
		virtue += 10
	case r.Doctrines.Base.Monolatry:
		virtue += 10
	case r.Doctrines.Base.Omnism:
		virtue += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		virtue += 20
	case r.Doctrines.Gender.Equality:
		virtue += 20
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 15
	}

	if r.Doctrines.Prophets {
		virtue += 10
	}
	if r.Doctrines.Legalism {
		virtue += 30
	}
	if r.Doctrines.Polyamory {
		virtue += 10
	}
	if r.Doctrines.ReligiousLaw {
		virtue += 30
	}
	if r.Doctrines.Pacifism {
		virtue += 10
	}
	if r.Doctrines.RitualHospitality {
		virtue += 10
	}
	if r.Doctrines.Darkness {
		neutral += 10
	}
	if r.Doctrines.Raider {
		virtue += 10
	}
	if r.Doctrines.Hedonism {
		virtue += 10
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

func getHonorParents(r *entities.Religion) religion.Attitude {
	var (
		virtue  = 10
		neutral = 10
		sin     = 10
	)
	switch {
	case r.Doctrines.Base.Monotheism:
		virtue += 10
	case r.Doctrines.Base.Polytheism:
		neutral += 10
	case r.Doctrines.Base.DeityDualism:
		virtue += 10
	case r.Doctrines.Base.Deism:
		neutral += 10
	case r.Doctrines.Base.Henothism:
		virtue += 10
	case r.Doctrines.Base.Monolatry:
		virtue += 10
	case r.Doctrines.Base.Omnism:
		neutral += 10
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		virtue += 25
	case r.Doctrines.Gender.Equality:
		virtue += 20
	case r.Doctrines.Gender.FemaleDominance:
		virtue += 15
	}

	if r.Doctrines.Prophets {
		virtue += 10
	}
	if r.Doctrines.Legalism {
		virtue += 10
	}
	if r.Doctrines.Polyamory {
		virtue += 10
	}
	if r.Doctrines.ReligiousLaw {
		virtue += 30
	}
	if r.Doctrines.AncestorWorship {
		virtue += 60
	}
	if r.Doctrines.Reincarnation {
		virtue += 10
	}
	if r.Doctrines.RitualHospitality {
		virtue += 10
	}
	if r.Doctrines.Darkness {
		neutral += 10
	}
	if r.Doctrines.Hedonism {
		neutral += 10
	}

	return getAttitudeByProbability(virtue, neutral, sin)
}

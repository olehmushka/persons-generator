package religion

import (
	"persons_generator/entities"
	"persons_generator/entities/religion"
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

func getAttitude(virtue, neutral, sin int) religion.Attitude {
	if virtue >= neutral && virtue > sin {
		return religion.Virtue
	}
	if sin >= neutral && sin > virtue {
		return religion.Sin
	}

	return religion.NeutralAttitude
}

func getProfanity(r *entities.Religion) religion.Attitude {
	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		sin += 2
	case r.Doctrines.Main.Polytheism:
		sin += 2
	case r.Doctrines.Main.DeityDualism:
		sin += 2
	case r.Doctrines.Main.Deism:
		neutral++
	case r.Doctrines.Main.Henothism:
		sin++
	case r.Doctrines.Main.Monolatry:
		sin++
	case r.Doctrines.Main.Omnism:
		neutral++
	}

	if r.Doctrines.FullTolerance {
		neutral++
		virtue++
	}
	if r.Doctrines.Messiah {
		sin++
	}
	if r.Doctrines.Prophets {
		sin++
	}
	if r.Doctrines.Asceticism {
		sin++
	}
	if r.Doctrines.Legalism {
		sin++
	}
	if r.Doctrines.ReligiousLaw {
		sin += 2
	}
	if r.Doctrines.TaxNonbelievers {
		sin++
	}
	if r.Doctrines.UnrelentingFaith {
		sin++
	}
	if r.Doctrines.Darkness {
		virtue++
	}
	if r.Doctrines.Proselytizer {
		sin++
	}

	return getAttitude(virtue, neutral, sin)
}

func getHaveOtherGods(r *entities.Religion) religion.Attitude {
	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		sin += 2
	case r.Doctrines.Main.Polytheism:
		neutral++
		virtue++
	case r.Doctrines.Main.DeityDualism:
		neutral++
	case r.Doctrines.Main.Deism:
		neutral++
		virtue++
	case r.Doctrines.Main.Henothism:
		sin++
		neutral++
	case r.Doctrines.Main.Monolatry:
		neutral++
		virtue++
	case r.Doctrines.Main.Omnism:
		virtue++
	}

	if r.Doctrines.Messiah {
		sin++
	}
	if r.Doctrines.Prophets {
		sin += 2
	}
	if r.Doctrines.MendicantPreachers {
		sin++
	}
	if r.Doctrines.Polyamory {
		virtue++
		neutral++
	}
	if r.Doctrines.ReligiousLaw {
		sin++
	}
	if r.Doctrines.TaxNonbelievers {
		sin++
	}
	if r.Doctrines.UnrelentingFaith {
		sin += 2
	}
	if r.Doctrines.AncestorWorship {
		sin++
	}
	if r.Doctrines.Proselytizer {
		sin += 2
	}

	return getAttitude(virtue, neutral, sin)
}

func getSloth(r *entities.Religion) religion.Attitude {
	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		sin++
	case r.Doctrines.Main.Polytheism:
		neutral++
	case r.Doctrines.Main.DeityDualism:
		sin++
	case r.Doctrines.Main.Deism:
		neutral++
	case r.Doctrines.Main.Henothism:
		sin++
	case r.Doctrines.Main.Monolatry:
		neutral++
	case r.Doctrines.Main.Omnism:
		neutral++
	}

	if r.Doctrines.FullTolerance {
		neutral++
		virtue++
	}
	if r.Doctrines.Asceticism {
		sin++
	}
	if r.Doctrines.Legalism {
		sin++
	}
	if r.Doctrines.Monasticism {
		sin++
	}
	if r.Doctrines.Polyamory {
		neutral++
		virtue++
	}
	if r.Doctrines.ReligiousLaw {
		sin++
	}
	if r.Doctrines.RitualHospitality {
		sin++
	}
	if r.Doctrines.Hedonism {
		virtue++
	}

	return getAttitude(virtue, neutral, sin)
}

func getGreed(r *entities.Religion) religion.Attitude {
	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		sin++
	case r.Doctrines.Main.Polytheism:
		neutral++
	case r.Doctrines.Main.DeityDualism:
		sin++
	case r.Doctrines.Main.Deism:
		neutral++
	case r.Doctrines.Main.Henothism:
		neutral++
	case r.Doctrines.Main.Monolatry:
		neutral++
	case r.Doctrines.Main.Omnism:
		neutral++
	}

	if r.Doctrines.ReligiousLaw {
		sin++
	}
	if r.Doctrines.Monasticism {
		sin++
	}
	if r.Doctrines.RitualHospitality {
		sin++
	}
	if r.Doctrines.Raider {
		virtue++
	}
	if r.Doctrines.Hedonism {
		virtue++
	}

	return getAttitude(virtue, neutral, sin)
}

func getCharity(r *entities.Religion) religion.Attitude {
	if r.Theology.Precepts.SinsAndVirtues.Greed == religion.Sin {
		return religion.Virtue
	}
	if r.Theology.Precepts.SinsAndVirtues.Greed == religion.Virtue {
		return religion.Sin
	}

	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		virtue++
	case r.Doctrines.Main.Polytheism:
		virtue++
	case r.Doctrines.Main.DeityDualism:
		virtue++
	case r.Doctrines.Main.Deism:
		virtue++
	case r.Doctrines.Main.Henothism:
		virtue++
	case r.Doctrines.Main.Monolatry:
		virtue++
	case r.Doctrines.Main.Omnism:
		virtue++
	}

	if r.Doctrines.FullTolerance {
		virtue++
	}
	if r.Doctrines.Asceticism {
		virtue++
	}
	if r.Doctrines.Darkness {
		neutral++
	}
	if r.Doctrines.Raider {
		sin++
	}
	if r.Doctrines.Hedonism {
		sin++
	}

	return getAttitude(virtue, neutral, sin)
}

func getGluttony(r *entities.Religion) religion.Attitude {
	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		sin++
	case r.Doctrines.Main.Polytheism:
		neutral++
	case r.Doctrines.Main.DeityDualism:
		sin++
	case r.Doctrines.Main.Deism:
		neutral++
	case r.Doctrines.Main.Henothism:
		neutral++
	case r.Doctrines.Main.Monolatry:
		neutral++
	case r.Doctrines.Main.Omnism:
		neutral++
	}

	if r.Doctrines.Asceticism {
		sin++
	}
	if r.Doctrines.Legalism {
		sin++
	}
	if r.Doctrines.MendicantPreachers {
		sin++
	}
	if r.Doctrines.Monasticism {
		sin++
	}
	if r.Doctrines.Polyamory {
		virtue++
	}
	if r.Doctrines.RitualHospitality {
		neutral++
		sin++
	}
	if r.Doctrines.PainIsVirtue {
		sin++
	}
	if r.Doctrines.Raider {
		virtue++
	}
	if r.Doctrines.Hedonism {
		virtue++
	}

	return getAttitude(virtue, neutral, sin)
}

func getTemperance(r *entities.Religion) religion.Attitude {
	if r.Theology.Precepts.SinsAndVirtues.Gluttony == religion.Sin {
		return religion.Virtue
	}
	if r.Theology.Precepts.SinsAndVirtues.Gluttony == religion.Virtue {
		return religion.Sin
	}

	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		virtue++
	case r.Doctrines.Main.Polytheism:
		neutral++
	case r.Doctrines.Main.DeityDualism:
		virtue++
	case r.Doctrines.Main.Deism:
		neutral++
	case r.Doctrines.Main.Henothism:
		neutral++
	case r.Doctrines.Main.Monolatry:
		neutral++
	case r.Doctrines.Main.Omnism:
		neutral++
	}

	if r.Doctrines.Asceticism {
		virtue++
	}
	if r.Doctrines.Legalism {
		virtue++
	}
	if r.Doctrines.MendicantPreachers {
		virtue++
	}
	if r.Doctrines.Monasticism {
		virtue++
	}
	if r.Doctrines.Polyamory {
		neutral++
	}
	if r.Doctrines.RitualHospitality {
		neutral++
		virtue++
	}
	if r.Doctrines.PainIsVirtue {
		virtue++
	}
	if r.Doctrines.Raider {
		sin++
	}
	if r.Doctrines.Hedonism {
		sin++
	}

	return getAttitude(virtue, neutral, sin)
}

func getPride(r *entities.Religion) religion.Attitude {
	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		virtue++
		neutral++
		sin++
	case r.Doctrines.Main.Polytheism:
		neutral++
	case r.Doctrines.Main.DeityDualism:
		neutral++
	case r.Doctrines.Main.Deism:
		neutral++
	case r.Doctrines.Main.Henothism:
		neutral++
	case r.Doctrines.Main.Monolatry:
		neutral++
	case r.Doctrines.Main.Omnism:
		neutral++
	}

	if r.Doctrines.Messiah {
		sin++
	}
	if r.Doctrines.Prophets {
		sin++
	}
	if r.Doctrines.Legalism {
		sin++
	}
	if r.Doctrines.Polyamory {
		neutral++
	}
	if r.Doctrines.AncestorWorship {
		virtue++
	}
	if r.Doctrines.Pacifism {
		sin++
	}
	if r.Doctrines.RitualHospitality {
		sin++
	}
	if r.Doctrines.SanctionedFalseConversions {
		sin++
	}
	if r.Doctrines.Proselytizer {
		virtue++
		sin++
		neutral++
	}
	if r.Doctrines.Hedonism {
		virtue++
	}

	return getAttitude(virtue, neutral, sin)
}

func getHumility(r *entities.Religion) religion.Attitude {
	if r.Theology.Precepts.SinsAndVirtues.Pride == religion.Sin {
		return religion.Virtue
	}
	if r.Theology.Precepts.SinsAndVirtues.Pride == religion.Virtue {
		return religion.Sin
	}

	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		virtue++
	case r.Doctrines.Main.Polytheism:
		neutral++
	case r.Doctrines.Main.DeityDualism:
		virtue++
	case r.Doctrines.Main.Deism:
		neutral++
	case r.Doctrines.Main.Henothism:
		neutral++
	case r.Doctrines.Main.Monolatry:
		neutral++
	case r.Doctrines.Main.Omnism:
		neutral++
	}

	if r.Doctrines.Asceticism {
		virtue++
	}
	if r.Doctrines.Legalism {
		virtue++
	}
	if r.Doctrines.MendicantPreachers {
		virtue++
	}
	if r.Doctrines.Monasticism {
		virtue++
	}
	if r.Doctrines.RitualHospitality {
		virtue++
	}
	if r.Doctrines.SanctionedFalseConversions {
		virtue++
	}
	if r.Doctrines.Darkness {
		sin++
	}
	if r.Doctrines.Raider {
		sin++
	}
	if r.Doctrines.Hedonism {
		sin++
	}

	return getAttitude(virtue, neutral, sin)
}

func getLust(r *entities.Religion) religion.Attitude {
	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		sin++
	case r.Doctrines.Main.Polytheism:
		neutral++
		virtue++
	case r.Doctrines.Main.DeityDualism:
		sin++
	case r.Doctrines.Main.Deism:
		neutral++
	case r.Doctrines.Main.Henothism:
		neutral++
	case r.Doctrines.Main.Monolatry:
		neutral++
	case r.Doctrines.Main.Omnism:
		neutral++
	}

	if r.Doctrines.FullTolerance {
		virtue++
	}
	if r.Doctrines.Prophets {
		sin++
	}
	if r.Doctrines.Asceticism {
		sin++
	}
	if r.Doctrines.Legalism {
		sin++
	}
	if r.Doctrines.Monasticism {
		sin++
	}
	if r.Doctrines.Polyamory {
		virtue++
	}
	if r.Doctrines.SanctityOfNature {
		virtue++
	}
	if r.Doctrines.Pacifism {
		neutral++
	}
	if r.Doctrines.SacredChildbirth {
		virtue++
	}
	if r.Doctrines.Raider {
		virtue++
	}
	if r.Doctrines.Hedonism {
		virtue++
	}

	return getAttitude(virtue, neutral, sin)
}

func getChaste(r *entities.Religion) religion.Attitude {
	if r.Theology.Precepts.SinsAndVirtues.Lust == religion.Sin {
		return religion.Virtue
	}
	if r.Theology.Precepts.SinsAndVirtues.Lust == religion.Virtue {
		return religion.Sin
	}

	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		virtue++
	case r.Doctrines.Main.Polytheism:
		virtue++
	case r.Doctrines.Main.DeityDualism:
		virtue++
	case r.Doctrines.Main.Deism:
		neutral++
	case r.Doctrines.Main.Henothism:
		neutral++
	case r.Doctrines.Main.Monolatry:
		neutral++
	case r.Doctrines.Main.Omnism:
		neutral++
		virtue++
	}

	if r.Doctrines.FullTolerance {
		neutral++
	}
	if r.Doctrines.Messiah {
		virtue++
	}
	if r.Doctrines.Asceticism {
		virtue++
	}
	if r.Doctrines.Legalism {
		virtue++
	}
	if r.Doctrines.Monasticism {
		virtue++
	}
	if r.Doctrines.Polyamory {
		sin++
	}
	if r.Doctrines.ReligiousLaw {
		virtue++
	}
	if r.Doctrines.SacredChildbirth {
		neutral++
	}
	if r.Doctrines.Raider {
		neutral++
	}
	if r.Doctrines.Hedonism {
		sin++
	}

	return getAttitude(virtue, neutral, sin)
}

func getWrath(r *entities.Religion) religion.Attitude {
	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		neutral++
		sin++
	case r.Doctrines.Main.Polytheism:
		neutral++
	case r.Doctrines.Main.DeityDualism:
		neutral++
		sin++
	case r.Doctrines.Main.Deism:
		neutral++
	case r.Doctrines.Main.Henothism:
		neutral++
	case r.Doctrines.Main.Monolatry:
		neutral++
	case r.Doctrines.Main.Omnism:
		neutral++
	}

	if r.Doctrines.Asceticism {
		sin++
	}
	if r.Doctrines.Esotericism {
		sin++
	}
	if r.Doctrines.Polyamory {
		sin++
	}
	if r.Doctrines.UnrelentingFaith {
		neutral++
	}
	if r.Doctrines.Pacifism {
		sin++
	}
	if r.Doctrines.SacredChildbirth {
		sin++
	}
	if r.Doctrines.Darkness {
		virtue++
	}
	if r.Doctrines.PainIsVirtue {
		virtue++
	}
	if r.Doctrines.Raider {
		virtue++
	}
	if r.Doctrines.Proselytizer {
		neutral++
	}
	if r.Doctrines.Hedonism {
		neutral++
	}

	return getAttitude(virtue, neutral, sin)
}

func getPatience(r *entities.Religion) religion.Attitude {
	if r.Theology.Precepts.SinsAndVirtues.Wrath == religion.Sin {
		return religion.Virtue
	}
	if r.Theology.Precepts.SinsAndVirtues.Wrath == religion.Virtue {
		return religion.Sin
	}

	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		virtue++
	case r.Doctrines.Main.Polytheism:
		virtue++
	case r.Doctrines.Main.DeityDualism:
		virtue++
	case r.Doctrines.Main.Deism:
		virtue++
	case r.Doctrines.Main.Henothism:
		virtue++
	case r.Doctrines.Main.Monolatry:
		virtue++
	case r.Doctrines.Main.Omnism:
		virtue++
	}

	if r.Doctrines.Asceticism {
		virtue++
	}
	if r.Doctrines.Astrology {
		virtue++
	}
	if r.Doctrines.Pacifism {
		virtue++
	}
	if r.Doctrines.PainIsVirtue {
		virtue++
	}
	if r.Doctrines.Raider {
		sin++
	}

	return getAttitude(virtue, neutral, sin)
}

func getPain(r *entities.Religion) religion.Attitude {
	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		neutral++
	case r.Doctrines.Main.Polytheism:
		neutral++
	case r.Doctrines.Main.DeityDualism:
		neutral++
	case r.Doctrines.Main.Deism:
		neutral++
	case r.Doctrines.Main.Henothism:
		neutral++
	case r.Doctrines.Main.Monolatry:
		neutral++
	case r.Doctrines.Main.Omnism:
		neutral++
	}

	if r.Doctrines.PainIsVirtue {
		virtue += 3
	}
	if r.Doctrines.Hedonism {
		sin++
	}

	return getAttitude(virtue, neutral, sin)
}

func getSadism(r *entities.Religion) religion.Attitude {
	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		sin++
	case r.Doctrines.Main.Polytheism:
		sin++
	case r.Doctrines.Main.DeityDualism:
		sin++
	case r.Doctrines.Main.Deism:
		sin++
	case r.Doctrines.Main.Henothism:
		sin++
	case r.Doctrines.Main.Monolatry:
		sin++
	case r.Doctrines.Main.Omnism:
		sin++
	}

	if r.Doctrines.FullTolerance {
		sin++
	}
	if r.Doctrines.Legalism {
		sin++
	}
	if r.Doctrines.Polyamory {
		sin++
	}
	if r.Doctrines.ReligiousLaw {
		sin++
	}
	if r.Doctrines.SanctityOfNature {
		sin++
	}
	if r.Doctrines.Pacifism {
		sin++
	}
	if r.Doctrines.Reincarnation {
		sin++
	}
	if r.Doctrines.RitualHospitality {
		sin++
	}
	if r.Doctrines.PainIsVirtue {
		virtue++
	}
	if r.Doctrines.Darkness {
		virtue++
	}
	if r.Doctrines.TreeConnection {
		sin++
	}
	if r.Doctrines.AnimalConnection {
		sin++
	}
	if r.Doctrines.Raider {
		virtue++
	}
	if r.Doctrines.Hedonism {
		neutral++
	}

	return getAttitude(virtue, neutral, sin)
}

func getEmpathy(r *entities.Religion) religion.Attitude {
	if r.Theology.Precepts.SinsAndVirtues.Sadism == religion.Sin {
		return religion.Virtue
	}
	if r.Theology.Precepts.SinsAndVirtues.Sadism == religion.Virtue {
		return religion.Sin
	}

	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		virtue++
	case r.Doctrines.Main.Polytheism:
		virtue++
	case r.Doctrines.Main.DeityDualism:
		virtue++
	case r.Doctrines.Main.Deism:
		virtue++
	case r.Doctrines.Main.Henothism:
		virtue++
	case r.Doctrines.Main.Monolatry:
		virtue++
	case r.Doctrines.Main.Omnism:
		virtue++
	}

	if r.Doctrines.FullTolerance {
		virtue++
	}
	if r.Doctrines.Legalism {
		virtue++
	}
	if r.Doctrines.Polyamory {
		virtue++
	}
	if r.Doctrines.ReligiousLaw {
		virtue++
	}
	if r.Doctrines.SanctityOfNature {
		virtue++
	}
	if r.Doctrines.Pacifism {
		virtue++
	}
	if r.Doctrines.Reincarnation {
		virtue++
	}
	if r.Doctrines.RitualHospitality {
		virtue++
	}
	if r.Doctrines.PainIsVirtue {
		neutral++
	}
	if r.Doctrines.Darkness {
		neutral++
	}
	if r.Doctrines.TreeConnection {
		virtue++
	}
	if r.Doctrines.AnimalConnection {
		virtue++
	}
	if r.Doctrines.Raider {
		sin++
	}
	if r.Doctrines.Hedonism {
		neutral++
	}

	return getAttitude(virtue, neutral, sin)
}

func getStealing(r *entities.Religion) religion.Attitude {
	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		sin++
	case r.Doctrines.Main.Polytheism:
		sin++
	case r.Doctrines.Main.DeityDualism:
		sin++
	case r.Doctrines.Main.Deism:
		neutral++
	case r.Doctrines.Main.Henothism:
		sin++
	case r.Doctrines.Main.Monolatry:
		sin++
	case r.Doctrines.Main.Omnism:
		sin++
	}

	if r.Doctrines.Prophets {
		sin++
	}
	if r.Doctrines.Legalism {
		sin++
	}
	if r.Doctrines.ReligiousLaw {
		sin++
	}
	if r.Doctrines.TaxNonbelievers {
		sin++
	}
	if r.Doctrines.UnrelentingFaith {
		sin++
	}
	if r.Doctrines.Darkness {
		virtue++
	}
	if r.Doctrines.LiveUnderGround {
		virtue++
	}
	if r.Doctrines.Hedonism {
		neutral++
	}

	return getAttitude(virtue, neutral, sin)
}

func getLie(r *entities.Religion) religion.Attitude {
	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		sin++
	case r.Doctrines.Main.Polytheism:
		neutral++
	case r.Doctrines.Main.DeityDualism:
		sin++
	case r.Doctrines.Main.Deism:
		neutral++
	case r.Doctrines.Main.Henothism:
		neutral++
	case r.Doctrines.Main.Monolatry:
		neutral++
	case r.Doctrines.Main.Omnism:
		neutral++
	}

	if r.Doctrines.Prophets {
		sin++
	}
	if r.Doctrines.Monasticism {
		sin++
	}
	if r.Doctrines.Legalism {
		sin++
	}
	if r.Doctrines.Polyamory {
		neutral++
	}
	if r.Doctrines.ReligiousLaw {
		sin++
	}
	if r.Doctrines.Pacifism {
		neutral++
	}
	if r.Doctrines.RitualHospitality {
		neutral++
	}
	if r.Doctrines.Darkness {
		virtue++
	}
	if r.Doctrines.Raider {
		sin++
	}
	if r.Doctrines.Hedonism {
		neutral++
	}

	return getAttitude(virtue, neutral, sin)
}

func getHonest(r *entities.Religion) religion.Attitude {
	if r.Theology.Precepts.SinsAndVirtues.Lie == religion.Sin {
		return religion.Virtue
	}
	if r.Theology.Precepts.SinsAndVirtues.Lie == religion.Virtue {
		return religion.Sin
	}

	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		virtue++
	case r.Doctrines.Main.Polytheism:
		virtue++
	case r.Doctrines.Main.DeityDualism:
		virtue++
	case r.Doctrines.Main.Deism:
		virtue++
	case r.Doctrines.Main.Henothism:
		virtue++
	case r.Doctrines.Main.Monolatry:
		virtue++
	case r.Doctrines.Main.Omnism:
		virtue++
	}

	if r.Doctrines.Prophets {
		virtue++
	}
	if r.Doctrines.Legalism {
		virtue++
	}
	if r.Doctrines.Polyamory {
		virtue++
	}
	if r.Doctrines.ReligiousLaw {
		virtue++
	}
	if r.Doctrines.Pacifism {
		virtue++
	}
	if r.Doctrines.RitualHospitality {
		virtue++
	}
	if r.Doctrines.Darkness {
		neutral++
	}
	if r.Doctrines.Raider {
		virtue++
	}
	if r.Doctrines.Hedonism {
		virtue++
	}

	return getAttitude(virtue, neutral, sin)
}

func getHonorParents(r *entities.Religion) religion.Attitude {
	var virtue, neutral, sin int
	switch {
	case r.Doctrines.Main.Monotheism:
		virtue++
	case r.Doctrines.Main.Polytheism:
		neutral++
	case r.Doctrines.Main.DeityDualism:
		virtue++
	case r.Doctrines.Main.Deism:
		neutral++
	case r.Doctrines.Main.Henothism:
		virtue++
	case r.Doctrines.Main.Monolatry:
		virtue++
	case r.Doctrines.Main.Omnism:
		neutral++
	}

	if r.Doctrines.Prophets {
		virtue++
	}
	if r.Doctrines.Legalism {
		virtue++
	}
	if r.Doctrines.Polyamory {
		virtue++
	}
	if r.Doctrines.ReligiousLaw {
		virtue++
	}
	if r.Doctrines.AncestorWorship {
		virtue += 2
	}
	if r.Doctrines.Reincarnation {
		virtue++
	}
	if r.Doctrines.RitualHospitality {
		virtue++
	}
	if r.Doctrines.Darkness {
		neutral++
	}
	if r.Doctrines.Hedonism {
		neutral++
	}

	return getAttitude(virtue, neutral, sin)
}

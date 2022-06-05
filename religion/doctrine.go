package religion

import (
	"persons_generator/entities"
	rel "persons_generator/entities/religion"
	pm "persons_generator/probability-machine"
)

func generateDoctrine(r *entities.Religion) {
	doc := &rel.Doctrines{
		Base: getMainDoctrine(),
	}

	doc.Gender = getGenderDoctrine(doc)
	doc.FullTolerance = getFullTolerance(doc)
	doc.Messiah = getMessiah(doc)
	doc.Prophets = getProphets(doc)
	doc.Asceticism = getAsceticism(doc)
	doc.Astrology = getAstrology(doc)
	doc.Esotericism = getEsotericism(doc)
	doc.MendicantPreachers = getMendicantPreachers(doc)
	doc.Monasticism = getMonasticism(doc)
	doc.Polyamory = getPolyamory(doc)
	doc.ReligiousLaw = getReligiousLaw(doc)
	doc.Legalism = getLegalism(doc)
	doc.SanctityOfNature = getSanctityOfNature(doc)
	doc.TaxNonbelievers = getTaxNonbelievers(doc)
	doc.UnrelentingFaith = getUnrelentingFaith(doc)
	doc.AncestorWorship = getAncestorWorship(doc)
	doc.Pacifism = getPacifism(doc)
	doc.Reincarnation = getReincarnation(doc)
	doc.RitualHospitality = getRitualHospitality(doc)
	doc.SacredChildbirth = getSacredChildbirth(doc)
	doc.SanctionedFalseConversions = getSanctionedFalseConversions(doc)
	doc.SunWorship = getSunWorship(doc)
	doc.MoonWorship = getMoonWorship(doc)
	doc.PainIsVirtue = getPainIsVirtue(doc)
	doc.Darkness = getDarkness(doc)
	doc.LiveUnderGround = getLiveUnderGround(doc)
	doc.TreeConnection = getTreeConnection(doc)
	doc.AnimalConnection = getAnimalConnection(doc)
	doc.Blindsight = getBlindsight(doc)
	doc.Raider = getRaider(doc)
	doc.Proselytizer = getProselytizer(doc)
	doc.Hedonism = getHedonism(doc)

	r.Doctrines = doc
}

func getMainDoctrine() *rel.BaseDoctrine {
	md := &rel.BaseDoctrine{}
	count := 0
	for {
		if pm.GetRandomBool(60) {
			md.Monotheism = true
			break
		}
		if pm.GetRandomBool(60) {
			md.Polytheism = true
			break
		}
		if pm.GetRandomBool(60) {
			md.DeityDualism = true
			break
		}
		if pm.GetRandomBool(20) {
			md.Deism = true
			break
		}
		if pm.GetRandomBool(40) {
			md.Henothism = true
			break
		}
		if pm.GetRandomBool(55) {
			md.Monolatry = true
			break
		}
		if pm.GetRandomBool(20) {
			md.Omnism = true
			break
		}

		if count == 10 {
			md.Omnism = true
			break
		}
		count++
	}
	return md
}

func getGenderDoctrine(doc *rel.Doctrines) *rel.GenderDoctrine {
	var (
		male     = 30
		equality = 25
		female   = 10
	)
	switch {
	case doc.Base.Monotheism:
		male += 15
		equality += 8
	case doc.Base.Polytheism:
		male += 15
		equality += 10
		female += 5
	case doc.Base.DeityDualism:
		male += 12
		equality += 10
		female += 7
	case doc.Base.Deism:
		male += 10
		equality += 20
		female += 5
	case doc.Base.Henothism:
		male += 15
		equality += 10
		female += 5
	case doc.Base.Monolatry:
		male += 15
		equality += 10
		female += 10
	case doc.Base.Omnism:
		male += 10
		equality += 15
		female += 5
	}

	gd := &rel.GenderDoctrine{}
	count := 0
	for {
		if pm.GetRandomBool(male) {
			gd.MaleDominance = true
			break
		}
		if pm.GetRandomBool(equality) {
			gd.Equality = true
			break
		}
		if pm.GetRandomBool(female) {
			gd.FemaleDominance = true
			break
		}
		if count == 10 {
			gd.Equality = true
			break
		}
		count++
	}
	return gd
}

func getFullTolerance(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 10
	case doc.Base.Polytheism:
		primaryProb = 15
	case doc.Base.DeityDualism:
		primaryProb = 12
	case doc.Base.Deism:
		primaryProb = 80
	case doc.Base.Henothism:
		primaryProb = 13
	case doc.Base.Monolatry:
		primaryProb = 20
	case doc.Base.Omnism:
		primaryProb = 85
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb -= 30
	case doc.Gender.Equality:
		primaryProb += 30
	case doc.Gender.FemaleDominance:
		primaryProb -= 20
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getMessiah(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 70
	case doc.Base.Polytheism:
		primaryProb = 15
	case doc.Base.DeityDualism:
		primaryProb = 60
	case doc.Base.Deism:
		primaryProb = 2
	case doc.Base.Henothism:
		primaryProb = 20
	case doc.Base.Monolatry:
		primaryProb = 12
	case doc.Base.Omnism:
		primaryProb = 30
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 10
	case doc.Gender.Equality:
		primaryProb += 5
	case doc.Gender.FemaleDominance:
		primaryProb -= 5
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getProphets(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 80
	case doc.Base.Polytheism:
		primaryProb = 50
	case doc.Base.DeityDualism:
		primaryProb = 70
	case doc.Base.Deism:
		primaryProb = 10
	case doc.Base.Henothism:
		primaryProb = 50
	case doc.Base.Monolatry:
		primaryProb = 50
	case doc.Base.Omnism:
		primaryProb = 50
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 10
	case doc.Gender.Equality:
	case doc.Gender.FemaleDominance:
		primaryProb -= 10
	}

	if doc.Messiah {
		primaryProb += 30
	}
	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getAsceticism(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 45
	case doc.Base.Polytheism:
		primaryProb = 15
	case doc.Base.DeityDualism:
		primaryProb = 45
	case doc.Base.Deism:
		primaryProb = 5
	case doc.Base.Henothism:
		primaryProb = 20
	case doc.Base.Monolatry:
		primaryProb = 20
	case doc.Base.Omnism:
		primaryProb = 7
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 10
	case doc.Gender.Equality:
	case doc.Gender.FemaleDominance:
		primaryProb -= 10
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getAstrology(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 40
	case doc.Base.Polytheism:
		primaryProb = 70
	case doc.Base.DeityDualism:
		primaryProb = 55
	case doc.Base.Deism:
		primaryProb = 70
	case doc.Base.Henothism:
		primaryProb = 65
	case doc.Base.Monolatry:
		primaryProb = 65
	case doc.Base.Omnism:
		primaryProb = 55
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 10
	case doc.Gender.Equality:
	case doc.Gender.FemaleDominance:
		primaryProb += 5
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getEsotericism(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 40
	case doc.Base.Polytheism:
		primaryProb = 70
	case doc.Base.DeityDualism:
		primaryProb = 55
	case doc.Base.Deism:
		primaryProb = 70
	case doc.Base.Henothism:
		primaryProb = 65
	case doc.Base.Monolatry:
		primaryProb = 65
	case doc.Base.Omnism:
		primaryProb = 55
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 5
	case doc.Gender.Equality:
		primaryProb += 7
	case doc.Gender.FemaleDominance:
		primaryProb += 5
	}

	if doc.Astrology {
		primaryProb += 20
	}
	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getMendicantPreachers(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 55
	case doc.Base.Polytheism:
		primaryProb = 45
	case doc.Base.DeityDualism:
		primaryProb = 45
	case doc.Base.Deism:
		primaryProb = 0
	case doc.Base.Henothism:
		primaryProb = 10
	case doc.Base.Monolatry:
		primaryProb = 10
	case doc.Base.Omnism:
		primaryProb = 5
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 5
	case doc.Gender.Equality:
		primaryProb += 2
	case doc.Gender.FemaleDominance:
		primaryProb -= 5
	}

	if doc.Asceticism {
		primaryProb += 40
	}
	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getMonasticism(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 45
	case doc.Base.Polytheism:
		primaryProb = 25
	case doc.Base.DeityDualism:
		primaryProb = 30
	case doc.Base.Deism:
		primaryProb = 5
	case doc.Base.Henothism:
		primaryProb = 30
	case doc.Base.Monolatry:
		primaryProb = 30
	case doc.Base.Omnism:
		primaryProb = 40
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 7
	case doc.Gender.Equality:
		primaryProb += 2
	case doc.Gender.FemaleDominance:
		primaryProb -= 3
	}

	if doc.Asceticism {
		primaryProb += 45
	}
	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getPolyamory(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 20
	case doc.Base.Polytheism:
		primaryProb = 30
	case doc.Base.DeityDualism:
		primaryProb = 25
	case doc.Base.Deism:
		primaryProb = 40
	case doc.Base.Henothism:
		primaryProb = 28
	case doc.Base.Monolatry:
		primaryProb = 29
	case doc.Base.Omnism:
		primaryProb = 38
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 5
	case doc.Gender.Equality:
		primaryProb += 20
	case doc.Gender.FemaleDominance:
		primaryProb -= 2
	}

	if doc.FullTolerance {
		primaryProb += 40
	}
	if doc.Asceticism {
		primaryProb -= 45
	}
	if doc.Messiah {
		primaryProb -= 5
	}
	if doc.Prophets {
		primaryProb -= 10
	}
	if doc.Monasticism {
		primaryProb -= 10
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getReligiousLaw(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 50
	case doc.Base.Polytheism:
		primaryProb = 18
	case doc.Base.DeityDualism:
		primaryProb = 45
	case doc.Base.Deism:
		primaryProb = 3
	case doc.Base.Henothism:
		primaryProb = 40
	case doc.Base.Monolatry:
		primaryProb = 25
	case doc.Base.Omnism:
		primaryProb = 18
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 10
	case doc.Gender.Equality:
		primaryProb -= 5
	case doc.Gender.FemaleDominance:
		primaryProb -= 15
	}

	if doc.FullTolerance {
		primaryProb -= 35
	}
	if doc.Prophets {
		primaryProb += 35
	}
	if doc.Asceticism {
		primaryProb += 15
	}
	if doc.Polyamory {
		primaryProb -= 15
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getLegalism(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 30
	case doc.Base.Polytheism:
		primaryProb = 10
	case doc.Base.DeityDualism:
		primaryProb = 20
	case doc.Base.Deism:
		primaryProb = 30
	case doc.Base.Henothism:
		primaryProb = 20
	case doc.Base.Monolatry:
		primaryProb = 18
	case doc.Base.Omnism:
		primaryProb = 30
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 7
	case doc.Gender.Equality:
		primaryProb += 2
	case doc.Gender.FemaleDominance:
		primaryProb -= 5
	}

	if doc.FullTolerance {
		primaryProb -= 45
	}
	if doc.Prophets {
		primaryProb += 30
	}
	if doc.Asceticism {
		primaryProb += 25
	}
	if doc.Polyamory {
		primaryProb -= 40
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getSanctityOfNature(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 40
	case doc.Base.Polytheism:
		primaryProb = 60
	case doc.Base.DeityDualism:
		primaryProb = 50
	case doc.Base.Deism:
		primaryProb = 60
	case doc.Base.Henothism:
		primaryProb = 50
	case doc.Base.Monolatry:
		primaryProb = 55
	case doc.Base.Omnism:
		primaryProb = 50
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 2
	case doc.Gender.Equality:
		primaryProb += 5
	case doc.Gender.FemaleDominance:
		primaryProb += 15
	}

	if doc.Prophets {
		primaryProb += 5
	}
	if doc.Asceticism {
		primaryProb += 5
	}
	if doc.Astrology {
		primaryProb += 5
	}
	if doc.Esotericism {
		primaryProb += 5
	}
	if doc.Polyamory {
		primaryProb += 5
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getTaxNonbelievers(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 50
	case doc.Base.Polytheism:
		primaryProb = 5
	case doc.Base.DeityDualism:
		primaryProb = 20
	case doc.Base.Deism:
		primaryProb = 0
	case doc.Base.Henothism:
		primaryProb = 15
	case doc.Base.Monolatry:
		primaryProb = 10
	case doc.Base.Omnism:
		primaryProb = 0
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 7
	case doc.Gender.Equality:
		primaryProb += 5
	case doc.Gender.FemaleDominance:
	}

	if doc.FullTolerance {
		primaryProb -= 10
	}
	if doc.Messiah {
		primaryProb += 5
	}
	if doc.Prophets {
		primaryProb += 3
	}
	if doc.Monasticism {
		primaryProb += 5
	}
	if doc.ReligiousLaw {
		primaryProb += 10
	}
	if doc.Legalism {
		primaryProb += 5
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getUnrelentingFaith(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 50
	case doc.Base.Polytheism:
		primaryProb = 22
	case doc.Base.DeityDualism:
		primaryProb = 25
	case doc.Base.Deism:
		primaryProb = 0
	case doc.Base.Henothism:
		primaryProb = 24
	case doc.Base.Monolatry:
		primaryProb = 23
	case doc.Base.Omnism:
		primaryProb = 0
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 10
	case doc.Gender.Equality:
		primaryProb -= 15
	case doc.Gender.FemaleDominance:
		primaryProb -= 15
	}

	if doc.FullTolerance {
		primaryProb -= 50
	}
	if doc.Messiah {
		primaryProb += 5
	}
	if doc.Prophets {
		primaryProb += 5
	}
	if doc.Asceticism {
		primaryProb += 10
	}
	if doc.Esotericism {
		primaryProb -= 10
	}
	if doc.Monasticism {
		primaryProb += 10
	}
	if doc.Polyamory {
		primaryProb -= 40
	}
	if doc.TaxNonbelievers {
		primaryProb += 3
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getAncestorWorship(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 10
	case doc.Base.Polytheism:
		primaryProb = 80
	case doc.Base.DeityDualism:
		primaryProb = 40
	case doc.Base.Deism:
		primaryProb = 50
	case doc.Base.Henothism:
		primaryProb = 85
	case doc.Base.Monolatry:
		primaryProb = 86
	case doc.Base.Omnism:
		primaryProb = 70
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 10
	case doc.Gender.Equality:
		primaryProb -= 10
	case doc.Gender.FemaleDominance:
		primaryProb -= 2
	}

	if doc.Astrology {
		primaryProb += 10
	}
	if doc.Esotericism {
		primaryProb += 5
	}
	if doc.SanctityOfNature {
		primaryProb += 5
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getPacifism(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 15
	case doc.Base.Polytheism:
		primaryProb = 15
	case doc.Base.DeityDualism:
		primaryProb = 15
	case doc.Base.Deism:
		primaryProb = 25
	case doc.Base.Henothism:
		primaryProb = 15
	case doc.Base.Monolatry:
		primaryProb = 15
	case doc.Base.Omnism:
		primaryProb = 35
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb -= 10
	case doc.Gender.Equality:
	case doc.Gender.FemaleDominance:
		primaryProb += 20
	}

	if doc.FullTolerance {
		primaryProb += 30
	}
	if doc.Polyamory {
		primaryProb += 10
	}
	if doc.SanctityOfNature {
		primaryProb += 20
	}
	if doc.UnrelentingFaith {
		primaryProb -= 60
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getReincarnation(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 40
	case doc.Base.Polytheism:
		primaryProb = 50
	case doc.Base.DeityDualism:
		primaryProb = 45
	case doc.Base.Deism:
		primaryProb = 55
	case doc.Base.Henothism:
		primaryProb = 45
	case doc.Base.Monolatry:
		primaryProb = 50
	case doc.Base.Omnism:
		primaryProb = 55
	}

	switch {
	case doc.Gender.MaleDominance:
	case doc.Gender.Equality:
		primaryProb += 1
	case doc.Gender.FemaleDominance:
		primaryProb += 1
	}

	if doc.FullTolerance {
		primaryProb += 5
	}
	if doc.Asceticism {
		primaryProb += 3
	}
	if doc.Esotericism {
		primaryProb += 10
	}
	if doc.Monasticism {
		primaryProb += 3
	}
	if doc.Polyamory {
		primaryProb += 5
	}
	if doc.UnrelentingFaith {
		primaryProb -= 15
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getRitualHospitality(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 40
	case doc.Base.Polytheism:
		primaryProb = 40
	case doc.Base.DeityDualism:
		primaryProb = 40
	case doc.Base.Deism:
		primaryProb = 20
	case doc.Base.Henothism:
		primaryProb = 40
	case doc.Base.Monolatry:
		primaryProb = 40
	case doc.Base.Omnism:
		primaryProb = 50
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 5
	case doc.Gender.Equality:
		primaryProb += 7
	case doc.Gender.FemaleDominance:
		primaryProb += 10
	}

	if doc.FullTolerance {
		primaryProb += 5
	}
	if doc.Astrology {
		primaryProb += 2
	}
	if doc.Polyamory {
		primaryProb += 7
	}
	if doc.ReligiousLaw {
		primaryProb += 10
	}
	if doc.UnrelentingFaith {
		primaryProb -= 7
	}
	if doc.Pacifism {
		primaryProb += 5
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getSacredChildbirth(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 20
	case doc.Base.Polytheism:
		primaryProb = 60
	case doc.Base.DeityDualism:
		primaryProb = 50
	case doc.Base.Deism:
		primaryProb = 20
	case doc.Base.Henothism:
		primaryProb = 40
	case doc.Base.Monolatry:
		primaryProb = 40
	case doc.Base.Omnism:
		primaryProb = 50
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb -= 10
	case doc.Gender.Equality:
		primaryProb += 5
	case doc.Gender.FemaleDominance:
		primaryProb += 40
	}

	if doc.Messiah {
		primaryProb += 10
	}
	if doc.Astrology {
		primaryProb += 15
	}
	if doc.Esotericism {
		primaryProb += 5
	}
	if doc.Polyamory {
		primaryProb += 30
	}
	if doc.SanctityOfNature {
		primaryProb += 5
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getSanctionedFalseConversions(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 30
	case doc.Base.Polytheism:
		primaryProb = 30
	case doc.Base.DeityDualism:
		primaryProb = 30
	case doc.Base.Deism:
		primaryProb = 100
	case doc.Base.Henothism:
		primaryProb = 40
	case doc.Base.Monolatry:
		primaryProb = 40
	case doc.Base.Omnism:
		primaryProb = 40
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb -= 10
	case doc.Gender.Equality:
		primaryProb += 15
	case doc.Gender.FemaleDominance:
		primaryProb += 20
	}

	if doc.Asceticism {
		primaryProb -= 5
	}
	if doc.UnrelentingFaith {
		primaryProb -= 20
	}
	if doc.Pacifism {
		primaryProb += 20
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getSunWorship(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 50
	case doc.Base.Polytheism:
		primaryProb = 90
	case doc.Base.DeityDualism:
		primaryProb = 60
	case doc.Base.Deism:
		primaryProb = 40
	case doc.Base.Henothism:
		primaryProb = 80
	case doc.Base.Monolatry:
		primaryProb = 70
	case doc.Base.Omnism:
		primaryProb = 60
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 5
	case doc.Gender.Equality:
	case doc.Gender.FemaleDominance:
	}

	if doc.Astrology {
		primaryProb += 10
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getMoonWorship(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 10
	case doc.Base.Polytheism:
		primaryProb = 90
	case doc.Base.DeityDualism:
		primaryProb = 40
	case doc.Base.Deism:
		primaryProb = 40
	case doc.Base.Henothism:
		primaryProb = 50
	case doc.Base.Monolatry:
		primaryProb = 50
	case doc.Base.Omnism:
		primaryProb = 50
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 3
	case doc.Gender.Equality:
	case doc.Gender.FemaleDominance:
		primaryProb += 7
	}

	if doc.Astrology {
		primaryProb += 20
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getPainIsVirtue(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 20
	case doc.Base.Polytheism:
		primaryProb = 20
	case doc.Base.DeityDualism:
		primaryProb = 20
	case doc.Base.Deism:
		primaryProb = 21
	case doc.Base.Henothism:
		primaryProb = 20
	case doc.Base.Monolatry:
		primaryProb = 20
	case doc.Base.Omnism:
		primaryProb = 19
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb -= 5
	case doc.Gender.Equality:
		primaryProb -= 6
	case doc.Gender.FemaleDominance:
		primaryProb -= 7
	}

	if doc.Asceticism {
		primaryProb += 15
	}
	if doc.Esotericism {
		primaryProb += 5
	}
	if doc.Polyamory {
		primaryProb -= 5
	}
	if doc.SacredChildbirth {
		primaryProb -= 5
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getDarkness(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 5
	case doc.Base.Polytheism:
		primaryProb = 10
	case doc.Base.DeityDualism:
		primaryProb = 7
	case doc.Base.Deism:
		primaryProb = 5
	case doc.Base.Henothism:
		primaryProb = 10
	case doc.Base.Monolatry:
		primaryProb = 10
	case doc.Base.Omnism:
		primaryProb = 10
	}

	switch {
	case doc.Gender.MaleDominance:
	case doc.Gender.Equality:
	case doc.Gender.FemaleDominance:
	}

	if doc.Astrology {
		primaryProb -= 5
	}
	if doc.Esotericism {
		primaryProb += 5
	}
	if doc.SunWorship {
		primaryProb -= 100
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getLiveUnderGround(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 5
	case doc.Base.Polytheism:
		primaryProb = 5
	case doc.Base.DeityDualism:
		primaryProb = 7
	case doc.Base.Deism:
		primaryProb = 5
	case doc.Base.Henothism:
		primaryProb = 7
	case doc.Base.Monolatry:
		primaryProb = 7
	case doc.Base.Omnism:
		primaryProb = 7
	}

	switch {
	case doc.Gender.MaleDominance:
	case doc.Gender.Equality:
	case doc.Gender.FemaleDominance:
	}

	if doc.Astrology {
		primaryProb -= 10
	}
	if doc.Esotericism {
		primaryProb += 10
	}
	if doc.SunWorship {
		primaryProb -= 100
	}
	if doc.MoonWorship {
		primaryProb -= 100
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getTreeConnection(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 30
	case doc.Base.Polytheism:
		primaryProb = 55
	case doc.Base.DeityDualism:
		primaryProb = 40
	case doc.Base.Deism:
		primaryProb = 50
	case doc.Base.Henothism:
		primaryProb = 40
	case doc.Base.Monolatry:
		primaryProb = 50
	case doc.Base.Omnism:
		primaryProb = 45
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb -= 5
	case doc.Gender.Equality:
		primaryProb += 5
	case doc.Gender.FemaleDominance:
		primaryProb += 3
	}

	if doc.Esotericism {
		primaryProb += 7
	}
	if doc.SanctityOfNature {
		primaryProb += 7
	}
	if doc.AncestorWorship {
		primaryProb += 10
	}
	if doc.Reincarnation {
		primaryProb += 10
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getAnimalConnection(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 30
	case doc.Base.Polytheism:
		primaryProb = 60
	case doc.Base.DeityDualism:
		primaryProb = 40
	case doc.Base.Deism:
		primaryProb = 60
	case doc.Base.Henothism:
		primaryProb = 60
	case doc.Base.Monolatry:
		primaryProb = 60
	case doc.Base.Omnism:
		primaryProb = 50
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 3
	case doc.Gender.Equality:
		primaryProb += 5
	case doc.Gender.FemaleDominance:
		primaryProb += 7
	}

	if doc.Esotericism {
		primaryProb += 7
	}
	if doc.SanctityOfNature {
		primaryProb += 7
	}
	if doc.Pacifism {
		primaryProb += 5
	}
	if doc.Reincarnation {
		primaryProb += 15
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getBlindsight(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 20
	case doc.Base.Polytheism:
		primaryProb = 23
	case doc.Base.DeityDualism:
		primaryProb = 21
	case doc.Base.Deism:
		primaryProb = 25
	case doc.Base.Henothism:
		primaryProb = 23
	case doc.Base.Monolatry:
		primaryProb = 23
	case doc.Base.Omnism:
		primaryProb = 26
	}

	switch {
	case doc.Gender.MaleDominance:
	case doc.Gender.Equality:
		primaryProb -= 3
	case doc.Gender.FemaleDominance:
		primaryProb -= 7
	}

	if doc.Astrology {
		primaryProb -= 10
	}
	if doc.SunWorship {
		primaryProb -= 100
	}
	if doc.MoonWorship {
		primaryProb -= 100
	}
	if doc.Darkness {
		primaryProb += 60
	}
	if doc.PainIsVirtue {
		primaryProb += 20
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getRaider(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 28
	case doc.Base.Polytheism:
		primaryProb = 50
	case doc.Base.DeityDualism:
		primaryProb = 35
	case doc.Base.Deism:
		primaryProb = 38
	case doc.Base.Henothism:
		primaryProb = 52
	case doc.Base.Monolatry:
		primaryProb = 51
	case doc.Base.Omnism:
		primaryProb = 45
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 7
	case doc.Gender.Equality:
	case doc.Gender.FemaleDominance:
		primaryProb -= 15
	}

	if doc.ReligiousLaw {
		primaryProb -= 10
	}
	if doc.Legalism {
		primaryProb -= 10
	}
	if doc.TaxNonbelievers {
		primaryProb += 5
	}
	if doc.UnrelentingFaith {
		primaryProb += 5
	}
	if doc.AncestorWorship {
		primaryProb += 5
	}
	if doc.Pacifism {
		primaryProb -= 100
	}
	if doc.Reincarnation {
		primaryProb -= 100
	}
	if doc.PainIsVirtue {
		primaryProb += 50
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getProselytizer(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 55
	case doc.Base.Polytheism:
		primaryProb = 20
	case doc.Base.DeityDualism:
		primaryProb = 30
	case doc.Base.Deism:
		primaryProb = 10
	case doc.Base.Henothism:
		primaryProb = 30
	case doc.Base.Monolatry:
		primaryProb = 22
	case doc.Base.Omnism:
		primaryProb = 10
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 7
	case doc.Gender.Equality:
	case doc.Gender.FemaleDominance:
		primaryProb -= 3
	}

	if doc.FullTolerance {
		primaryProb -= 30
	}
	if doc.Messiah {
		primaryProb += 20
	}
	if doc.Prophets {
		primaryProb += 10
	}
	if doc.MendicantPreachers {
		primaryProb += 15
	}
	if doc.Polyamory {
		primaryProb -= 20
	}
	if doc.UnrelentingFaith {
		primaryProb += 30
	}
	if doc.PainIsVirtue {
		primaryProb += 5
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getHedonism(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Base.Monotheism:
		primaryProb = 15
	case doc.Base.Polytheism:
		primaryProb = 60
	case doc.Base.DeityDualism:
		primaryProb = 15
	case doc.Base.Deism:
		primaryProb = 60
	case doc.Base.Henothism:
		primaryProb = 40
	case doc.Base.Monolatry:
		primaryProb = 45
	case doc.Base.Omnism:
		primaryProb = 50
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 10
	case doc.Gender.Equality:
		primaryProb += 5
	case doc.Gender.FemaleDominance:
		primaryProb += 10
	}

	if doc.FullTolerance {
		primaryProb += 5
	}
	if doc.Asceticism {
		primaryProb -= 100
	}
	if doc.MendicantPreachers {
		primaryProb -= 50
	}
	if doc.Monasticism {
		primaryProb -= 50
	}
	if doc.Polyamory {
		primaryProb += 100
	}
	if doc.PainIsVirtue {
		primaryProb -= 100
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

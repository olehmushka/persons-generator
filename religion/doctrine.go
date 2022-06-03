package religion

import (
	"persons_generator/entities"
	rel "persons_generator/entities/religion"
	pm "persons_generator/probability-machine"
)

func generateDoctrine(r *entities.Religion) {
	doc := &rel.Doctrines{
		Main: getMainDoctrine(),
	}

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

func getMainDoctrine() *rel.MainDoctrine {
	md := &rel.MainDoctrine{}
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

func getFullTolerance(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Main.Monotheism:
		primaryProb = 5
	case doc.Main.Polytheism:
		primaryProb = 15
	case doc.Main.DeityDualism:
		primaryProb = 7
	case doc.Main.Deism:
		primaryProb = 80
	case doc.Main.Henothism:
		primaryProb = 12
	case doc.Main.Monolatry:
		primaryProb = 20
	case doc.Main.Omnism:
		primaryProb = 85
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getMessiah(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Main.Monotheism:
		primaryProb = 80
	case doc.Main.Polytheism:
		primaryProb = 15
	case doc.Main.DeityDualism:
		primaryProb = 60
	case doc.Main.Deism:
		primaryProb = 2
	case doc.Main.Henothism:
		primaryProb = 20
	case doc.Main.Monolatry:
		primaryProb = 12
	case doc.Main.Omnism:
		primaryProb = 30
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getProphets(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Main.Monotheism:
		primaryProb = 80
	case doc.Main.Polytheism:
		primaryProb = 50
	case doc.Main.DeityDualism:
		primaryProb = 70
	case doc.Main.Deism:
		primaryProb = 10
	case doc.Main.Henothism:
		primaryProb = 50
	case doc.Main.Monolatry:
		primaryProb = 50
	case doc.Main.Omnism:
		primaryProb = 50
	}

	if doc.Messiah {
		primaryProb += 30
	}
	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getAsceticism(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Main.Monotheism:
		primaryProb = 45
	case doc.Main.Polytheism:
		primaryProb = 15
	case doc.Main.DeityDualism:
		primaryProb = 45
	case doc.Main.Deism:
		primaryProb = 5
	case doc.Main.Henothism:
		primaryProb = 20
	case doc.Main.Monolatry:
		primaryProb = 20
	case doc.Main.Omnism:
		primaryProb = 7
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getAstrology(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Main.Monotheism:
		primaryProb = 40
	case doc.Main.Polytheism:
		primaryProb = 70
	case doc.Main.DeityDualism:
		primaryProb = 55
	case doc.Main.Deism:
		primaryProb = 70
	case doc.Main.Henothism:
		primaryProb = 65
	case doc.Main.Monolatry:
		primaryProb = 65
	case doc.Main.Omnism:
		primaryProb = 55
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getEsotericism(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Main.Monotheism:
		primaryProb = 40
	case doc.Main.Polytheism:
		primaryProb = 70
	case doc.Main.DeityDualism:
		primaryProb = 55
	case doc.Main.Deism:
		primaryProb = 70
	case doc.Main.Henothism:
		primaryProb = 65
	case doc.Main.Monolatry:
		primaryProb = 65
	case doc.Main.Omnism:
		primaryProb = 55
	}

	if doc.Astrology {
		primaryProb += 20
	}
	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getMendicantPreachers(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Main.Monotheism:
		primaryProb = 55
	case doc.Main.Polytheism:
		primaryProb = 45
	case doc.Main.DeityDualism:
		primaryProb = 45
	case doc.Main.Deism:
		primaryProb = 0
	case doc.Main.Henothism:
		primaryProb = 10
	case doc.Main.Monolatry:
		primaryProb = 10
	case doc.Main.Omnism:
		primaryProb = 5
	}

	if doc.Asceticism {
		primaryProb += 40
	}
	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getMonasticism(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Main.Monotheism:
		primaryProb = 45
	case doc.Main.Polytheism:
		primaryProb = 25
	case doc.Main.DeityDualism:
		primaryProb = 30
	case doc.Main.Deism:
		primaryProb = 5
	case doc.Main.Henothism:
		primaryProb = 30
	case doc.Main.Monolatry:
		primaryProb = 30
	case doc.Main.Omnism:
		primaryProb = 40
	}
	if doc.Asceticism {
		primaryProb += 45
	}
	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getPolyamory(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Main.Monotheism:
		primaryProb = 20
	case doc.Main.Polytheism:
		primaryProb = 30
	case doc.Main.DeityDualism:
		primaryProb = 25
	case doc.Main.Deism:
		primaryProb = 40
	case doc.Main.Henothism:
		primaryProb = 28
	case doc.Main.Monolatry:
		primaryProb = 29
	case doc.Main.Omnism:
		primaryProb = 38
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
	case doc.Main.Monotheism:
		primaryProb = 50
	case doc.Main.Polytheism:
		primaryProb = 18
	case doc.Main.DeityDualism:
		primaryProb = 45
	case doc.Main.Deism:
		primaryProb = 3
	case doc.Main.Henothism:
		primaryProb = 40
	case doc.Main.Monolatry:
		primaryProb = 25
	case doc.Main.Omnism:
		primaryProb = 18
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
	case doc.Main.Monotheism:
		primaryProb = 30
	case doc.Main.Polytheism:
		primaryProb = 10
	case doc.Main.DeityDualism:
		primaryProb = 20
	case doc.Main.Deism:
		primaryProb = 30
	case doc.Main.Henothism:
		primaryProb = 20
	case doc.Main.Monolatry:
		primaryProb = 18
	case doc.Main.Omnism:
		primaryProb = 30
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
	case doc.Main.Monotheism:
		primaryProb = 40
	case doc.Main.Polytheism:
		primaryProb = 60
	case doc.Main.DeityDualism:
		primaryProb = 50
	case doc.Main.Deism:
		primaryProb = 60
	case doc.Main.Henothism:
		primaryProb = 50
	case doc.Main.Monolatry:
		primaryProb = 55
	case doc.Main.Omnism:
		primaryProb = 50
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
	case doc.Main.Monotheism:
		primaryProb = 50
	case doc.Main.Polytheism:
		primaryProb = 5
	case doc.Main.DeityDualism:
		primaryProb = 20
	case doc.Main.Deism:
		primaryProb = 0
	case doc.Main.Henothism:
		primaryProb = 15
	case doc.Main.Monolatry:
		primaryProb = 10
	case doc.Main.Omnism:
		primaryProb = 0
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
	case doc.Main.Monotheism:
		primaryProb = 50
	case doc.Main.Polytheism:
		primaryProb = 22
	case doc.Main.DeityDualism:
		primaryProb = 25
	case doc.Main.Deism:
		primaryProb = 0
	case doc.Main.Henothism:
		primaryProb = 24
	case doc.Main.Monolatry:
		primaryProb = 23
	case doc.Main.Omnism:
		primaryProb = 0
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
	case doc.Main.Monotheism:
		primaryProb = 10
	case doc.Main.Polytheism:
		primaryProb = 80
	case doc.Main.DeityDualism:
		primaryProb = 40
	case doc.Main.Deism:
		primaryProb = 50
	case doc.Main.Henothism:
		primaryProb = 85
	case doc.Main.Monolatry:
		primaryProb = 86
	case doc.Main.Omnism:
		primaryProb = 70
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
	case doc.Main.Monotheism:
		primaryProb = 15
	case doc.Main.Polytheism:
		primaryProb = 15
	case doc.Main.DeityDualism:
		primaryProb = 15
	case doc.Main.Deism:
		primaryProb = 25
	case doc.Main.Henothism:
		primaryProb = 15
	case doc.Main.Monolatry:
		primaryProb = 15
	case doc.Main.Omnism:
		primaryProb = 35
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
	case doc.Main.Monotheism:
		primaryProb = 40
	case doc.Main.Polytheism:
		primaryProb = 50
	case doc.Main.DeityDualism:
		primaryProb = 45
	case doc.Main.Deism:
		primaryProb = 55
	case doc.Main.Henothism:
		primaryProb = 45
	case doc.Main.Monolatry:
		primaryProb = 50
	case doc.Main.Omnism:
		primaryProb = 55
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
	case doc.Main.Monotheism:
		primaryProb = 40
	case doc.Main.Polytheism:
		primaryProb = 40
	case doc.Main.DeityDualism:
		primaryProb = 40
	case doc.Main.Deism:
		primaryProb = 20
	case doc.Main.Henothism:
		primaryProb = 40
	case doc.Main.Monolatry:
		primaryProb = 40
	case doc.Main.Omnism:
		primaryProb = 50
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
	case doc.Main.Monotheism:
		primaryProb = 20
	case doc.Main.Polytheism:
		primaryProb = 60
	case doc.Main.DeityDualism:
		primaryProb = 50
	case doc.Main.Deism:
		primaryProb = 20
	case doc.Main.Henothism:
		primaryProb = 40
	case doc.Main.Monolatry:
		primaryProb = 40
	case doc.Main.Omnism:
		primaryProb = 50
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
	case doc.Main.Monotheism:
		primaryProb = 30
	case doc.Main.Polytheism:
		primaryProb = 30
	case doc.Main.DeityDualism:
		primaryProb = 30
	case doc.Main.Deism:
		primaryProb = 100
	case doc.Main.Henothism:
		primaryProb = 40
	case doc.Main.Monolatry:
		primaryProb = 40
	case doc.Main.Omnism:
		primaryProb = 40
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
	case doc.Main.Monotheism:
		primaryProb = 50
	case doc.Main.Polytheism:
		primaryProb = 90
	case doc.Main.DeityDualism:
		primaryProb = 60
	case doc.Main.Deism:
		primaryProb = 40
	case doc.Main.Henothism:
		primaryProb = 80
	case doc.Main.Monolatry:
		primaryProb = 70
	case doc.Main.Omnism:
		primaryProb = 60
	}

	if doc.Astrology {
		primaryProb += 10
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getMoonWorship(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Main.Monotheism:
		primaryProb = 10
	case doc.Main.Polytheism:
		primaryProb = 90
	case doc.Main.DeityDualism:
		primaryProb = 40
	case doc.Main.Deism:
		primaryProb = 40
	case doc.Main.Henothism:
		primaryProb = 50
	case doc.Main.Monolatry:
		primaryProb = 50
	case doc.Main.Omnism:
		primaryProb = 50
	}

	if doc.Astrology {
		primaryProb += 20
	}

	return pm.GetRandomBool(preparePrimaryProbability(primaryProb))
}

func getPainIsVirtue(doc *rel.Doctrines) bool {
	var primaryProb int
	switch {
	case doc.Main.Monotheism:
		primaryProb = 20
	case doc.Main.Polytheism:
		primaryProb = 20
	case doc.Main.DeityDualism:
		primaryProb = 20
	case doc.Main.Deism:
		primaryProb = 21
	case doc.Main.Henothism:
		primaryProb = 20
	case doc.Main.Monolatry:
		primaryProb = 20
	case doc.Main.Omnism:
		primaryProb = 19
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
	case doc.Main.Monotheism:
		primaryProb = 5
	case doc.Main.Polytheism:
		primaryProb = 10
	case doc.Main.DeityDualism:
		primaryProb = 7
	case doc.Main.Deism:
		primaryProb = 5
	case doc.Main.Henothism:
		primaryProb = 10
	case doc.Main.Monolatry:
		primaryProb = 10
	case doc.Main.Omnism:
		primaryProb = 10
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
	case doc.Main.Monotheism:
		primaryProb = 5
	case doc.Main.Polytheism:
		primaryProb = 5
	case doc.Main.DeityDualism:
		primaryProb = 7
	case doc.Main.Deism:
		primaryProb = 5
	case doc.Main.Henothism:
		primaryProb = 7
	case doc.Main.Monolatry:
		primaryProb = 7
	case doc.Main.Omnism:
		primaryProb = 7
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
	case doc.Main.Monotheism:
		primaryProb = 30
	case doc.Main.Polytheism:
		primaryProb = 55
	case doc.Main.DeityDualism:
		primaryProb = 40
	case doc.Main.Deism:
		primaryProb = 50
	case doc.Main.Henothism:
		primaryProb = 40
	case doc.Main.Monolatry:
		primaryProb = 50
	case doc.Main.Omnism:
		primaryProb = 45
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
	case doc.Main.Monotheism:
		primaryProb = 30
	case doc.Main.Polytheism:
		primaryProb = 60
	case doc.Main.DeityDualism:
		primaryProb = 40
	case doc.Main.Deism:
		primaryProb = 60
	case doc.Main.Henothism:
		primaryProb = 60
	case doc.Main.Monolatry:
		primaryProb = 60
	case doc.Main.Omnism:
		primaryProb = 50
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
	case doc.Main.Monotheism:
		primaryProb = 20
	case doc.Main.Polytheism:
		primaryProb = 23
	case doc.Main.DeityDualism:
		primaryProb = 21
	case doc.Main.Deism:
		primaryProb = 25
	case doc.Main.Henothism:
		primaryProb = 23
	case doc.Main.Monolatry:
		primaryProb = 23
	case doc.Main.Omnism:
		primaryProb = 26
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
	case doc.Main.Monotheism:
		primaryProb = 28
	case doc.Main.Polytheism:
		primaryProb = 50
	case doc.Main.DeityDualism:
		primaryProb = 35
	case doc.Main.Deism:
		primaryProb = 38
	case doc.Main.Henothism:
		primaryProb = 52
	case doc.Main.Monolatry:
		primaryProb = 51
	case doc.Main.Omnism:
		primaryProb = 45
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
	case doc.Main.Monotheism:
		primaryProb = 55
	case doc.Main.Polytheism:
		primaryProb = 20
	case doc.Main.DeityDualism:
		primaryProb = 30
	case doc.Main.Deism:
		primaryProb = 10
	case doc.Main.Henothism:
		primaryProb = 30
	case doc.Main.Monolatry:
		primaryProb = 22
	case doc.Main.Omnism:
		primaryProb = 10
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
	case doc.Main.Monotheism:
		primaryProb = 15
	case doc.Main.Polytheism:
		primaryProb = 60
	case doc.Main.DeityDualism:
		primaryProb = 15
	case doc.Main.Deism:
		primaryProb = 60
	case doc.Main.Henothism:
		primaryProb = 40
	case doc.Main.Monolatry:
		primaryProb = 45
	case doc.Main.Omnism:
		primaryProb = 50
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

package religion

import (
	"fmt"

	"persons_generator/entities"
	rel "persons_generator/entities/religion"
	pm "persons_generator/probability-machine"
)

func generateDoctrine(r *entities.Religion) {
	fmt.Println("[generateDoctrine] started")
	defer fmt.Println("[generateDoctrine] finished")

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
	fmt.Println("[getMainDoctrine] started")
	defer fmt.Println("[getMainDoctrine] finished")

	md := &rel.BaseDoctrine{}
	count := 0
	for {
		if pm.GetRandomBool(0.6) {
			md.Monotheism = true
			break
		}
		if pm.GetRandomBool(0.6) {
			md.Polytheism = true
			break
		}
		if pm.GetRandomBool(0.6) {
			md.DeityDualism = true
			break
		}
		if pm.GetRandomBool(0.2) {
			md.Deism = true
			break
		}
		if pm.GetRandomBool(0.4) {
			md.Henothism = true
			break
		}
		if pm.GetRandomBool(0.55) {
			md.Monolatry = true
			break
		}
		if pm.GetRandomBool(0.2) {
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
		male     = 0.3
		equality = 0.25
		female   = 0.1
	)
	switch {
	case doc.Base.Monotheism:
		male += 0.15
		equality += 0.08
	case doc.Base.Polytheism:
		male += 0.15
		equality += 0.1
		female += 0.05
	case doc.Base.DeityDualism:
		male += 0.12
		equality += 0.1
		female += 0.07
	case doc.Base.Deism:
		male += 0.1
		equality += 0.2
		female += 0.05
	case doc.Base.Henothism:
		male += 0.15
		equality += 0.1
		female += 0.05
	case doc.Base.Monolatry:
		male += 0.15
		equality += 0.1
		female += 0.1
	case doc.Base.Omnism:
		male += 0.1
		equality += 0.15
		female += 0.05
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
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.1
	case doc.Base.Polytheism:
		primaryProb = 0.15
	case doc.Base.DeityDualism:
		primaryProb = 0.12
	case doc.Base.Deism:
		primaryProb = 0.8
	case doc.Base.Henothism:
		primaryProb = 0.13
	case doc.Base.Monolatry:
		primaryProb = 0.2
	case doc.Base.Omnism:
		primaryProb = 0.85
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb -= 0.3
	case doc.Gender.Equality:
		primaryProb += 0.3
	case doc.Gender.FemaleDominance:
		primaryProb -= 0.2
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getMessiah(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.7
	case doc.Base.Polytheism:
		primaryProb = 0.15
	case doc.Base.DeityDualism:
		primaryProb = 0.6
	case doc.Base.Deism:
		primaryProb = 0.02
	case doc.Base.Henothism:
		primaryProb = 0.2
	case doc.Base.Monolatry:
		primaryProb = 0.12
	case doc.Base.Omnism:
		primaryProb = 0.3
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.1
	case doc.Gender.Equality:
		primaryProb += 0.05
	case doc.Gender.FemaleDominance:
		primaryProb -= 0.05
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getProphets(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.8
	case doc.Base.Polytheism:
		primaryProb = 0.5
	case doc.Base.DeityDualism:
		primaryProb = 0.7
	case doc.Base.Deism:
		primaryProb = 0.1
	case doc.Base.Henothism:
		primaryProb = 0.5
	case doc.Base.Monolatry:
		primaryProb = 0.5
	case doc.Base.Omnism:
		primaryProb = 0.5
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.1
	case doc.Gender.Equality:
	case doc.Gender.FemaleDominance:
		primaryProb -= 0.1
	}

	if doc.Messiah {
		primaryProb += 0.3
	}
	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getAsceticism(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.35
	case doc.Base.Polytheism:
		primaryProb = 0.15
	case doc.Base.DeityDualism:
		primaryProb = 0.45
	case doc.Base.Deism:
		primaryProb = 0.05
	case doc.Base.Henothism:
		primaryProb = 0.2
	case doc.Base.Monolatry:
		primaryProb = 0.2
	case doc.Base.Omnism:
		primaryProb = 0.7
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.1
	case doc.Gender.Equality:
	case doc.Gender.FemaleDominance:
		primaryProb -= 0.1
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getAstrology(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.4
	case doc.Base.Polytheism:
		primaryProb = 0.7
	case doc.Base.DeityDualism:
		primaryProb = 0.55
	case doc.Base.Deism:
		primaryProb = 0.7
	case doc.Base.Henothism:
		primaryProb = 0.65
	case doc.Base.Monolatry:
		primaryProb = 0.65
	case doc.Base.Omnism:
		primaryProb = 0.55
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.1
	case doc.Gender.Equality:
	case doc.Gender.FemaleDominance:
		primaryProb += 0.05
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getEsotericism(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.4
	case doc.Base.Polytheism:
		primaryProb = 0.7
	case doc.Base.DeityDualism:
		primaryProb = 0.55
	case doc.Base.Deism:
		primaryProb = 0.7
	case doc.Base.Henothism:
		primaryProb = 0.65
	case doc.Base.Monolatry:
		primaryProb = 0.65
	case doc.Base.Omnism:
		primaryProb = 0.55
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.05
	case doc.Gender.Equality:
		primaryProb += 0.07
	case doc.Gender.FemaleDominance:
		primaryProb += 0.05
	}

	if doc.Astrology {
		primaryProb += 0.2
	}
	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getMendicantPreachers(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.55
	case doc.Base.Polytheism:
		primaryProb = 0.45
	case doc.Base.DeityDualism:
		primaryProb = 0.45
	case doc.Base.Deism:
	case doc.Base.Henothism:
		primaryProb = 0.1
	case doc.Base.Monolatry:
		primaryProb = 0.1
	case doc.Base.Omnism:
		primaryProb = 0.05
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.05
	case doc.Gender.Equality:
		primaryProb += 0.02
	case doc.Gender.FemaleDominance:
		primaryProb -= 0.05
	}

	if doc.Asceticism {
		primaryProb += 0.4
	}
	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getMonasticism(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.45
	case doc.Base.Polytheism:
		primaryProb = 0.25
	case doc.Base.DeityDualism:
		primaryProb = 0.3
	case doc.Base.Deism:
		primaryProb = 0.05
	case doc.Base.Henothism:
		primaryProb = 0.3
	case doc.Base.Monolatry:
		primaryProb = 0.3
	case doc.Base.Omnism:
		primaryProb = 0.4
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.07
	case doc.Gender.Equality:
		primaryProb += 0.02
	case doc.Gender.FemaleDominance:
		primaryProb -= 0.03
	}

	if doc.Asceticism {
		primaryProb += 0.45
	}
	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getPolyamory(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.2
	case doc.Base.Polytheism:
		primaryProb = 0.3
	case doc.Base.DeityDualism:
		primaryProb = 0.25
	case doc.Base.Deism:
		primaryProb = 0.4
	case doc.Base.Henothism:
		primaryProb = 0.28
	case doc.Base.Monolatry:
		primaryProb = 0.29
	case doc.Base.Omnism:
		primaryProb = 0.38
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.05
	case doc.Gender.Equality:
		primaryProb += 0.02
	case doc.Gender.FemaleDominance:
		primaryProb -= 0.02
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

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getReligiousLaw(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.5
	case doc.Base.Polytheism:
		primaryProb = 0.18
	case doc.Base.DeityDualism:
		primaryProb = 0.45
	case doc.Base.Deism:
		primaryProb = 0.03
	case doc.Base.Henothism:
		primaryProb = 0.4
	case doc.Base.Monolatry:
		primaryProb = 0.25
	case doc.Base.Omnism:
		primaryProb = 0.18
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.1
	case doc.Gender.Equality:
		primaryProb -= 0.05
	case doc.Gender.FemaleDominance:
		primaryProb -= 0.15
	}

	if doc.FullTolerance {
		primaryProb -= 0.35
	}
	if doc.Prophets {
		primaryProb += 0.35
	}
	if doc.Asceticism {
		primaryProb += 0.15
	}
	if doc.Polyamory {
		primaryProb -= 0.15
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getLegalism(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.3
	case doc.Base.Polytheism:
		primaryProb = 0.1
	case doc.Base.DeityDualism:
		primaryProb = 0.2
	case doc.Base.Deism:
		primaryProb = 0.3
	case doc.Base.Henothism:
		primaryProb = 0.2
	case doc.Base.Monolatry:
		primaryProb = 0.18
	case doc.Base.Omnism:
		primaryProb = 0.3
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.07
	case doc.Gender.Equality:
		primaryProb += 0.02
	case doc.Gender.FemaleDominance:
		primaryProb -= 0.05
	}

	if doc.FullTolerance {
		primaryProb -= 0.45
	}
	if doc.Prophets {
		primaryProb += 0.3
	}
	if doc.Asceticism {
		primaryProb += 0.25
	}
	if doc.Polyamory {
		primaryProb -= 0.4
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getSanctityOfNature(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.4
	case doc.Base.Polytheism:
		primaryProb = 0.6
	case doc.Base.DeityDualism:
		primaryProb = 0.5
	case doc.Base.Deism:
		primaryProb = 0.6
	case doc.Base.Henothism:
		primaryProb = 0.5
	case doc.Base.Monolatry:
		primaryProb = 0.55
	case doc.Base.Omnism:
		primaryProb = 0.5
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.02
	case doc.Gender.Equality:
		primaryProb += 0.05
	case doc.Gender.FemaleDominance:
		primaryProb += 0.15
	}

	if doc.Prophets {
		primaryProb += 0.05
	}
	if doc.Asceticism {
		primaryProb += 0.05
	}
	if doc.Astrology {
		primaryProb += 0.05
	}
	if doc.Esotericism {
		primaryProb += 0.05
	}
	if doc.Polyamory {
		primaryProb += 0.05
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getTaxNonbelievers(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.5
	case doc.Base.Polytheism:
		primaryProb = 0.05
	case doc.Base.DeityDualism:
		primaryProb = 0.2
	case doc.Base.Deism:
		primaryProb = 0.03
	case doc.Base.Henothism:
		primaryProb = 0.15
	case doc.Base.Monolatry:
		primaryProb = 0.1
	case doc.Base.Omnism:
		primaryProb = 0.02
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.07
	case doc.Gender.Equality:
		primaryProb += 0.05
	case doc.Gender.FemaleDominance:
		primaryProb += 0.001
	}

	if doc.FullTolerance {
		primaryProb -= 0.1
	}
	if doc.Messiah {
		primaryProb += 0.05
	}
	if doc.Prophets {
		primaryProb += 0.03
	}
	if doc.Monasticism {
		primaryProb += 0.05
	}
	if doc.ReligiousLaw {
		primaryProb += 0.1
	}
	if doc.Legalism {
		primaryProb += 0.05
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getUnrelentingFaith(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.39
	case doc.Base.Polytheism:
		primaryProb = 0.22
	case doc.Base.DeityDualism:
		primaryProb = 0.25
	case doc.Base.Deism:
		primaryProb = 0.0003
	case doc.Base.Henothism:
		primaryProb = 0.24
	case doc.Base.Monolatry:
		primaryProb = 0.23
	case doc.Base.Omnism:
		primaryProb = 0.0002
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.1
	case doc.Gender.Equality:
		primaryProb += 0.001
	case doc.Gender.FemaleDominance:
		primaryProb += 0.01
	}

	if doc.FullTolerance {
		primaryProb -= 0.25
	}
	if doc.Messiah {
		primaryProb += 0.03
	}
	if doc.Prophets {
		primaryProb += 0.03
	}
	if doc.Asceticism {
		primaryProb += 0.07
	}
	if doc.Esotericism {
		primaryProb -= 0.05
	}
	if doc.Monasticism {
		primaryProb += 0.05
	}
	if doc.Polyamory {
		primaryProb -= 0.02
	}
	if doc.TaxNonbelievers {
		primaryProb += 0.015
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getAncestorWorship(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.1
	case doc.Base.Polytheism:
		primaryProb = 0.8
	case doc.Base.DeityDualism:
		primaryProb = 0.4
	case doc.Base.Deism:
		primaryProb = 0.5
	case doc.Base.Henothism:
		primaryProb = 0.85
	case doc.Base.Monolatry:
		primaryProb = 0.86
	case doc.Base.Omnism:
		primaryProb = 0.7
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.1
	case doc.Gender.Equality:
		primaryProb -= 0.002
	case doc.Gender.FemaleDominance:
		primaryProb -= 0.002
	}

	if doc.Astrology {
		primaryProb += 0.1
	}
	if doc.Esotericism {
		primaryProb += 0.05
	}
	if doc.SanctityOfNature {
		primaryProb += 0.05
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getPacifism(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.15
	case doc.Base.Polytheism:
		primaryProb = 0.165
	case doc.Base.DeityDualism:
		primaryProb = 0.151
	case doc.Base.Deism:
		primaryProb = 0.25
	case doc.Base.Henothism:
		primaryProb = 0.15
	case doc.Base.Monolatry:
		primaryProb = 0.15
	case doc.Base.Omnism:
		primaryProb = 0.359
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb -= 0.1
	case doc.Gender.Equality:
		primaryProb += 0.05
	case doc.Gender.FemaleDominance:
		primaryProb += 0.2
	}

	if doc.FullTolerance {
		primaryProb += 0.3
	}
	if doc.Polyamory {
		primaryProb += 0.1
	}
	if doc.SanctityOfNature {
		primaryProb += 0.2
	}
	if doc.UnrelentingFaith {
		primaryProb -= 0.6
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getReincarnation(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.3
	case doc.Base.Polytheism:
		primaryProb = 0.45
	case doc.Base.DeityDualism:
		primaryProb = 0.425
	case doc.Base.Deism:
		primaryProb = 0.5
	case doc.Base.Henothism:
		primaryProb = 0.415
	case doc.Base.Monolatry:
		primaryProb = 0.5
	case doc.Base.Omnism:
		primaryProb = 0.55
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.005
	case doc.Gender.Equality:
		primaryProb += 0.01
	case doc.Gender.FemaleDominance:
		primaryProb += 0.001
	}

	if doc.FullTolerance {
		primaryProb += 0.05
	}
	if doc.Asceticism {
		primaryProb += 0.03
	}
	if doc.Esotericism {
		primaryProb += 0.1
	}
	if doc.Monasticism {
		primaryProb += 0.03
	}
	if doc.Polyamory {
		primaryProb += 0.06
	}
	if doc.UnrelentingFaith {
		primaryProb -= 0.15
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getRitualHospitality(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.35
	case doc.Base.Polytheism:
		primaryProb = 0.3
	case doc.Base.DeityDualism:
		primaryProb = 0.33
	case doc.Base.Deism:
		primaryProb = 0.2
	case doc.Base.Henothism:
		primaryProb = 0.33
	case doc.Base.Monolatry:
		primaryProb = 0.33
	case doc.Base.Omnism:
		primaryProb = 0.37
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.05
	case doc.Gender.Equality:
		primaryProb += 0.07
	case doc.Gender.FemaleDominance:
		primaryProb += 0.1
	}

	if doc.FullTolerance {
		primaryProb += 0.05
	}
	if doc.Astrology {
		primaryProb += 0.02
	}
	if doc.Polyamory {
		primaryProb += 0.07
	}
	if doc.ReligiousLaw {
		primaryProb += 0.1
	}
	if doc.UnrelentingFaith {
		primaryProb -= 0.07
	}
	if doc.Pacifism {
		primaryProb += 0.05
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getSacredChildbirth(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.2
	case doc.Base.Polytheism:
		primaryProb = 0.473
	case doc.Base.DeityDualism:
		primaryProb = 0.485
	case doc.Base.Deism:
		primaryProb = 0.2
	case doc.Base.Henothism:
		primaryProb = 0.4
	case doc.Base.Monolatry:
		primaryProb = 0.4
	case doc.Base.Omnism:
		primaryProb = 0.41
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb -= 0.05
	case doc.Gender.Equality:
		primaryProb += 0.05
	case doc.Gender.FemaleDominance:
		primaryProb += 0.33
	}

	if doc.Messiah {
		primaryProb += 0.1
	}
	if doc.Astrology {
		primaryProb += 0.15
	}
	if doc.Esotericism {
		primaryProb += 0.05
	}
	if doc.Polyamory {
		primaryProb += 0.3
	}
	if doc.SanctityOfNature {
		primaryProb += 0.23
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getSanctionedFalseConversions(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.3
	case doc.Base.Polytheism:
		primaryProb = 0.3
	case doc.Base.DeityDualism:
		primaryProb = 0.3
	case doc.Base.Deism:
		primaryProb = 0.95
	case doc.Base.Henothism:
		primaryProb = 0.4
	case doc.Base.Monolatry:
		primaryProb = 0.4
	case doc.Base.Omnism:
		primaryProb = 0.4
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb -= 0.05
	case doc.Gender.Equality:
		primaryProb += 0.15
	case doc.Gender.FemaleDominance:
		primaryProb += 0.17
	}

	if doc.Asceticism {
		primaryProb -= 0.05
	}
	if doc.UnrelentingFaith {
		primaryProb -= 0.2
	}
	if doc.Pacifism {
		primaryProb += 0.2
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getSunWorship(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.35
	case doc.Base.Polytheism:
		primaryProb = 0.9
	case doc.Base.DeityDualism:
		primaryProb = 0.5
	case doc.Base.Deism:
		primaryProb = 0.4
	case doc.Base.Henothism:
		primaryProb = 0.73
	case doc.Base.Monolatry:
		primaryProb = 0.7
	case doc.Base.Omnism:
		primaryProb = 0.65
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.05
	case doc.Gender.Equality:
		primaryProb += 0.03
	case doc.Gender.FemaleDominance:
		primaryProb += 0.01
	}

	if doc.Astrology {
		primaryProb += 0.1
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getMoonWorship(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.1
	case doc.Base.Polytheism:
		primaryProb = 0.88
	case doc.Base.DeityDualism:
		primaryProb = 0.4
	case doc.Base.Deism:
		primaryProb = 0.4
	case doc.Base.Henothism:
		primaryProb = 0.5
	case doc.Base.Monolatry:
		primaryProb = 0.5
	case doc.Base.Omnism:
		primaryProb = 0.5
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.03
	case doc.Gender.Equality:
		primaryProb += 0.01
	case doc.Gender.FemaleDominance:
		primaryProb += 0.07
	}

	if doc.Astrology {
		primaryProb += 0.33
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getPainIsVirtue(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.1
	case doc.Base.Polytheism:
		primaryProb = 0.12
	case doc.Base.DeityDualism:
		primaryProb = 0.11
	case doc.Base.Deism:
		primaryProb = 0.09
	case doc.Base.Henothism:
		primaryProb = 0.12
	case doc.Base.Monolatry:
		primaryProb = 0.125
	case doc.Base.Omnism:
		primaryProb = 0.09
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.007
	case doc.Gender.Equality:
		primaryProb += 0.006
	case doc.Gender.FemaleDominance:
		primaryProb += 0.005
	}

	if doc.Asceticism {
		primaryProb += 0.125
	}
	if doc.Esotericism {
		primaryProb += 0.0445
	}
	if doc.Polyamory {
		primaryProb -= 0.0005
	}
	if doc.SacredChildbirth {
		primaryProb += 0.005
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getDarkness(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.05
	case doc.Base.Polytheism:
		primaryProb = 0.1
	case doc.Base.DeityDualism:
		primaryProb = 0.07
	case doc.Base.Deism:
		primaryProb = 0.05
	case doc.Base.Henothism:
		primaryProb = 0.1
	case doc.Base.Monolatry:
		primaryProb = 0.1
	case doc.Base.Omnism:
		primaryProb = 0.1
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.005
	case doc.Gender.Equality:
		primaryProb += 0.004
	case doc.Gender.FemaleDominance:
		primaryProb += 0.006
	}

	if doc.Astrology {
		primaryProb -= 0.05
	}
	if doc.Esotericism {
		primaryProb += 0.05
	}
	if doc.SunWorship {
		primaryProb -= 0.9
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getLiveUnderGround(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.029
	case doc.Base.Polytheism:
		primaryProb = 0.031
	case doc.Base.DeityDualism:
		primaryProb = 0.0349
	case doc.Base.Deism:
		primaryProb = 0.0291
	case doc.Base.Henothism:
		primaryProb = 0.027
	case doc.Base.Monolatry:
		primaryProb = 0.0297
	case doc.Base.Omnism:
		primaryProb = 0.0257
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.0005
	case doc.Gender.Equality:
		primaryProb += 0.0004
	case doc.Gender.FemaleDominance:
		primaryProb += 0.0005
	}

	if doc.Astrology {
		primaryProb -= 0.1
	}
	if doc.Esotericism {
		primaryProb += 0.1
	}
	if doc.SunWorship {
		primaryProb -= 0.99
	}
	if doc.MoonWorship {
		primaryProb -= 0.98
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getTreeConnection(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.25
	case doc.Base.Polytheism:
		primaryProb = 0.5
	case doc.Base.DeityDualism:
		primaryProb = 0.31
	case doc.Base.Deism:
		primaryProb = 0.35
	case doc.Base.Henothism:
		primaryProb = 0.4
	case doc.Base.Monolatry:
		primaryProb = 0.45
	case doc.Base.Omnism:
		primaryProb = 0.47
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.005
	case doc.Gender.Equality:
		primaryProb += 0.05
	case doc.Gender.FemaleDominance:
		primaryProb += 0.03
	}

	if doc.Esotericism {
		primaryProb += 0.07
	}
	if doc.SanctityOfNature {
		primaryProb += 0.07
	}
	if doc.AncestorWorship {
		primaryProb += 0.1
	}
	if doc.Reincarnation {
		primaryProb += 0.15
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getAnimalConnection(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.15
	case doc.Base.Polytheism:
		primaryProb = 0.5
	case doc.Base.DeityDualism:
		primaryProb = 0.35
	case doc.Base.Deism:
		primaryProb = 0.56
	case doc.Base.Henothism:
		primaryProb = 0.486
	case doc.Base.Monolatry:
		primaryProb = 0.516
	case doc.Base.Omnism:
		primaryProb = 0.5
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.023
	case doc.Gender.Equality:
		primaryProb += 0.0485
	case doc.Gender.FemaleDominance:
		primaryProb += 0.067
	}

	if doc.Esotericism {
		primaryProb += 0.067
	}
	if doc.SanctityOfNature {
		primaryProb += 0.097
	}
	if doc.Pacifism {
		primaryProb += 0.095
	}
	if doc.Reincarnation {
		primaryProb += 0.15
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getBlindsight(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.092
	case doc.Base.Polytheism:
		primaryProb = 0.0923
	case doc.Base.DeityDualism:
		primaryProb = 0.1121
	case doc.Base.Deism:
		primaryProb = 0.125
	case doc.Base.Henothism:
		primaryProb = 0.0927
	case doc.Base.Monolatry:
		primaryProb = 0.0933
	case doc.Base.Omnism:
		primaryProb = 0.0946
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.0001
	case doc.Gender.Equality:
		primaryProb -= 0.0003
	case doc.Gender.FemaleDominance:
		primaryProb -= 0.007
	}

	if doc.Astrology {
		primaryProb -= 0.1
	}
	if doc.SunWorship {
		primaryProb -= 0.95
	}
	if doc.MoonWorship {
		primaryProb -= 0.95
	}
	if doc.Darkness {
		primaryProb += 0.6
	}
	if doc.PainIsVirtue {
		primaryProb += 0.2
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getRaider(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.18
	case doc.Base.Polytheism:
		primaryProb = 0.35
	case doc.Base.DeityDualism:
		primaryProb = 0.25
	case doc.Base.Deism:
		primaryProb = 0.198
	case doc.Base.Henothism:
		primaryProb = 0.3652
	case doc.Base.Monolatry:
		primaryProb = 0.3351
	case doc.Base.Omnism:
		primaryProb = 0.2045
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.07
	case doc.Gender.Equality:
		primaryProb += 0.0001
	case doc.Gender.FemaleDominance:
		primaryProb -= 0.015
	}

	if doc.ReligiousLaw {
		primaryProb -= 0.1
	}
	if doc.Legalism {
		primaryProb -= 0.1
	}
	if doc.TaxNonbelievers {
		primaryProb += 0.035
	}
	if doc.UnrelentingFaith {
		primaryProb += 0.035
	}
	if doc.AncestorWorship {
		primaryProb += 0.035
	}
	if doc.Pacifism {
		primaryProb -= 0.9
	}
	if doc.Reincarnation {
		primaryProb -= 0.8
	}
	if doc.PainIsVirtue {
		primaryProb += 0.35
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getProselytizer(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.5
	case doc.Base.Polytheism:
		primaryProb = 0.2
	case doc.Base.DeityDualism:
		primaryProb = 0.3
	case doc.Base.Deism:
		primaryProb = 0.1
	case doc.Base.Henothism:
		primaryProb = 0.3
	case doc.Base.Monolatry:
		primaryProb = 0.22
	case doc.Base.Omnism:
		primaryProb = 0.1
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.07
	case doc.Gender.Equality:
		primaryProb += 0.01
	case doc.Gender.FemaleDominance:
		primaryProb += 0.02
	}

	if doc.FullTolerance {
		primaryProb -= 0.23
	}
	if doc.Messiah {
		primaryProb += 0.192
	}
	if doc.Prophets {
		primaryProb += 0.851
	}
	if doc.MendicantPreachers {
		primaryProb += 0.15
	}
	if doc.Polyamory {
		primaryProb -= 0.152
	}
	if doc.UnrelentingFaith {
		primaryProb += 0.3
	}
	if doc.PainIsVirtue {
		primaryProb += 0.05
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getHedonism(doc *rel.Doctrines) bool {
	var primaryProb float64
	switch {
	case doc.Base.Monotheism:
		primaryProb = 0.12
	case doc.Base.Polytheism:
		primaryProb = 0.45
	case doc.Base.DeityDualism:
		primaryProb = 0.155
	case doc.Base.Deism:
		primaryProb = 0.496
	case doc.Base.Henothism:
		primaryProb = 0.41
	case doc.Base.Monolatry:
		primaryProb = 0.43
	case doc.Base.Omnism:
		primaryProb = 0.475
	}

	switch {
	case doc.Gender.MaleDominance:
		primaryProb += 0.05
	case doc.Gender.Equality:
		primaryProb += 0.055
	case doc.Gender.FemaleDominance:
		primaryProb += 0.057
	}

	if doc.FullTolerance {
		primaryProb += 0.05
	}
	if doc.Asceticism {
		primaryProb -= 0.9
	}
	if doc.MendicantPreachers {
		primaryProb -= 0.5
	}
	if doc.Monasticism {
		primaryProb -= 0.5
	}
	if doc.Polyamory {
		primaryProb += 0.9
	}
	if doc.PainIsVirtue {
		primaryProb -= 0.9
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

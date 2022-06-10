package religion

import (
	"fmt"

	"persons_generator/entities"
	"persons_generator/entities/religion"
	pm "persons_generator/probability-machine"
)

func generateRituals(r *entities.Religion) {
	fmt.Println("[generateRituals] started")
	defer fmt.Println("[generateRituals] finished")
	r.Theology.Precepts.Rituals = &religion.Rituals{}
	r.Theology.Precepts.Rituals.RitualCelebrations = getRitualCelebrations(r)
	r.Theology.Precepts.Rituals.Funerals = getFunerals(r)
	r.Theology.Precepts.Rituals.Sacrifices = getSacrifices(r)
	r.Theology.Precepts.Rituals.Adorcism = getAdorcism(r)
	r.Theology.Precepts.Rituals.Exorcism = getExorcism(r)
	r.Theology.Precepts.Rituals.Orgy = getOrgy(r)
	r.Theology.Precepts.Rituals.SmokeCircle = getSmokeCircle(r)
	r.Theology.Precepts.Rituals.GladiatorDuel = getGladiatorDuel(r)
	r.Theology.Precepts.Rituals.Blinding = getBlinding(r)
	r.Theology.Precepts.Rituals.RitualCannibalism = getRitualCannibalism(r)
}

func getRitualCelebrations(r *entities.Religion) *religion.Celebrations {
	c := &religion.Celebrations{}
	c.DanceParty = getDanceParty(r)
	c.DrumParty = getDrumParty(r)
	c.SocialFestivals = getSocialFestivals(r)
	c.TreeCelebration = getTreeCelebration(r)
	return c
}

func getDanceParty(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.25
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.4
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.35
	case r.Doctrines.Base.Deism:
		primaryProb = 0.33
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.32
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.34
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.41
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.11
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.15
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.16
	}

	if r.Doctrines.FullTolerance {
		primaryProb += 0.03
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
		primaryProb += 0.01
	}
	if r.Doctrines.Asceticism {
		primaryProb -= 0.01
	}
	if r.Doctrines.Astrology {
		primaryProb += 0.01
	}
	if r.Doctrines.Esotericism {
		primaryProb += 0.03
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
		primaryProb += 0.005
	}
	if r.Doctrines.Monasticism {
		primaryProb += 0.005
	}
	if r.Doctrines.Polyamory {
		primaryProb += 0.05
	}
	if r.Doctrines.ReligiousLaw {
	}
	if r.Doctrines.SanctityOfNature {
		primaryProb += 0.02
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		primaryProb += 0.02
	}
	if r.Doctrines.AncestorWorship {
		primaryProb += 0.05
	}
	if r.Doctrines.Pacifism {
		primaryProb += 0.01
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
		primaryProb += 0.02
	}
	if r.Doctrines.SacredChildbirth {
		primaryProb += 0.01
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		primaryProb += 0.1
	}
	if r.Doctrines.MoonWorship {
		primaryProb += 0.1
	}
	if r.Doctrines.PainIsVirtue {
		primaryProb += 0.03
	}
	if r.Doctrines.Darkness {
		primaryProb += 0.01
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
		primaryProb -= 0.03
	}
	if r.Doctrines.Raider {
		primaryProb += 0.01
	}
	if r.Doctrines.Proselytizer {
		primaryProb += 0.03
	}
	if r.Doctrines.Hedonism {
		primaryProb += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		primaryProb += 0.05
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
		primaryProb += 0.05
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
		primaryProb += 0.005
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
		primaryProb += 0.001
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
		primaryProb += 0.005
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
		primaryProb += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
		primaryProb += 0.009
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
		primaryProb += 0.05
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
		primaryProb += 0.009
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
		primaryProb += 0.02
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
		primaryProb += 0.04
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
		primaryProb += 0.06
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
		primaryProb += 0.02
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
		primaryProb += 0.04
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
		primaryProb += 0.09
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
		primaryProb += 0.05
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
		primaryProb += 0.02
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
		primaryProb += 0.02
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
		primaryProb += 0.02
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
		primaryProb += 0.005
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
		primaryProb += 0.005
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Stealing.IsVirtue():
		primaryProb += 0.05
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

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getDrumParty(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.24
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.43
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.355
	case r.Doctrines.Base.Deism:
		primaryProb = 0.345
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.385
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.44
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.46
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.15
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.1
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.12
	}

	if r.Doctrines.FullTolerance {
		primaryProb += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
		primaryProb += 0.01
	}
	if r.Doctrines.Asceticism {
		primaryProb += 0.01
	}
	if r.Doctrines.Astrology {
		primaryProb += 0.01
	}
	if r.Doctrines.Esotericism {
		primaryProb += 0.03
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
		primaryProb += 0.02
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
		primaryProb += 0.01
	}
	if r.Doctrines.ReligiousLaw {
	}
	if r.Doctrines.SanctityOfNature {
		primaryProb += 0.01
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		primaryProb += 0.02
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
		primaryProb += 0.05
	}
	if r.Doctrines.MoonWorship {
		primaryProb += 0.06
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
		primaryProb += 0.01
	}
	if r.Doctrines.LiveUnderGround {
		primaryProb += 0.01
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
		primaryProb += 0.05
	}
	if r.Doctrines.Raider {
		primaryProb += 0.1
	}
	if r.Doctrines.Proselytizer {
		primaryProb += 0.03
	}
	if r.Doctrines.Hedonism {
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		primaryProb += 0.02
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
		primaryProb += 0.02
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
		primaryProb -= 0.01
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
		primaryProb -= 0.01
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
		primaryProb -= 0.01
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
		primaryProb += 0.02
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
		primaryProb -= 0.01
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

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getSocialFestivals(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.3
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.445
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.39
	case r.Doctrines.Base.Deism:
		primaryProb = 0.395
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.45
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.46
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.46
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.1
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.19
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.12
	}

	if r.Doctrines.FullTolerance {
		primaryProb += 0.05
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
		primaryProb += 0.05
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
		primaryProb += 0.1
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
		primaryProb += 0.05
	}
	if r.Doctrines.ReligiousLaw {
		primaryProb += 0.05
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
		primaryProb += 0.02
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
		primaryProb += 0.1
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		primaryProb += 0.05
	}
	if r.Doctrines.MoonWorship {
		primaryProb += 0.05
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
		primaryProb -= 0.01
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
		primaryProb += 0.01
	}
	if r.Doctrines.Proselytizer {
		primaryProb += 0.15
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
		primaryProb += 0.01
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
		primaryProb -= 0.01
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
		primaryProb += 0.02
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
		primaryProb += 0.03
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		primaryProb += 0.01
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
		primaryProb += 0.03
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
		primaryProb += 0.01
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
		primaryProb += 0.05
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

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getTreeCelebration(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.15
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.4
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.25
	case r.Doctrines.Base.Deism:
		primaryProb = 0.39
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.38
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.39
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.4
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.08
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.1
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.11
	}

	if r.Doctrines.FullTolerance {
		primaryProb += 0.01
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
		primaryProb += 0.01
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
		primaryProb += 0.2
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
		primaryProb += 0.05
	}
	if r.Doctrines.Pacifism {
		primaryProb += 0.01
	}
	if r.Doctrines.Reincarnation {
		primaryProb += 0.1
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		primaryProb += 0.01
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
	}
	if r.Doctrines.LiveUnderGround {
		primaryProb -= 0.01
	}
	if r.Doctrines.TreeConnection {
		primaryProb += 0.3
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
		primaryProb += 0.01
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
		primaryProb += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
		primaryProb -= 0.02
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
		primaryProb += 0.01
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

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getFunerals(r *entities.Religion) *religion.Funerals {
	f := &religion.Funerals{}
	f.SkyBurials = getSkyBurials(r)
	f.CaveBurials = getCaveBurials(r)
	f.UndergroundBurials = getUndergroundBurials(r)
	return f
}

func getSkyBurials(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.5
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.35
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.41
	case r.Doctrines.Base.Deism:
		primaryProb = 0.5
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.38
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.365
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.4
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.05
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.1
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.04
	}

	if r.Doctrines.FullTolerance {
		primaryProb += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
		primaryProb += 0.01
	}
	if r.Doctrines.Asceticism {
		primaryProb += 0.1
	}
	if r.Doctrines.Astrology {
		primaryProb += 0.05
	}
	if r.Doctrines.Esotericism {
		primaryProb += 0.01
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
		primaryProb += 0.05
	}
	if r.Doctrines.SanctityOfNature {
		primaryProb += 0.01
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		primaryProb += 0.01
	}
	if r.Doctrines.AncestorWorship {
		primaryProb -= 0.02
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
		primaryProb -= 0.05
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		primaryProb += 0.05
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
		primaryProb += 0.01
	}
	if r.Doctrines.Darkness {
		primaryProb += 0.01
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
		primaryProb += 0.01
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
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
		primaryProb -= 0.01
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
		primaryProb += 0.01
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

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getCaveBurials(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.48
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.55
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.53
	case r.Doctrines.Base.Deism:
		primaryProb = 0.5
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.555
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.555
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.5
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.07
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.5
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.06
	}

	if r.Doctrines.FullTolerance {
		primaryProb += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
		primaryProb += 0.01
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
		primaryProb += 0.05
	}
	if r.Doctrines.SanctityOfNature {
		primaryProb += 0.01
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
		primaryProb += 0.3
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
		primaryProb += 0.1
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
		primaryProb += 0.08
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
		primaryProb += 0.2
	}
	if r.Doctrines.LiveUnderGround {
		primaryProb -= 0.01
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
		primaryProb += 0.01
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
		primaryProb -= 0.01
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
		primaryProb += 0.02
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		primaryProb += 0.01
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
		primaryProb += 0.05
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
		primaryProb += 0.1
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getUndergroundBurials(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.55
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.5
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.53
	case r.Doctrines.Base.Deism:
		primaryProb = 0.5
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.5
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.5
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.5
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.08
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.8
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.08
	}

	if r.Doctrines.FullTolerance {
		primaryProb += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
		primaryProb += 0.09
	}
	if r.Doctrines.Astrology {
		primaryProb += 0.01
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
		primaryProb += 0.05
	}
	if r.Doctrines.SanctityOfNature {
		primaryProb += 0.01
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
		primaryProb += 0.2
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
		primaryProb += 0.1
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		primaryProb += 0.01
	}
	if r.Doctrines.MoonWorship {
		primaryProb += 0.01
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
		primaryProb += 0.1
	}
	if r.Doctrines.LiveUnderGround {
		primaryProb += 0.2
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
		primaryProb += 0.02
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
		primaryProb += 0.01
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
		primaryProb += 0.01
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
		primaryProb += 0.05
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getSacrifices(r *entities.Religion) *religion.Sacrifices {
	s := &religion.Sacrifices{}
	s.RitualSuicide = getRitualSuicide(r)
	s.AnimalSacrifice = getAnimalSacrifice(r)
	s.HumanSacrifice = getHumanSacrifice(r)
	s.PlantsSacrifice = getPlantsSacrifice(r)
	return s
}

func getRitualSuicide(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.02
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.03
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.021
	case r.Doctrines.Base.Deism:
		primaryProb = 0.025
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.031
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.03
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.0209
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.009
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.011
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.008
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
		primaryProb -= 0.01
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		primaryProb += 0.01
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
		primaryProb += 0.01
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
		primaryProb -= 0.01
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
		primaryProb += 0.01
	}
	if r.Doctrines.Darkness {
		primaryProb += 0.02
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
		primaryProb += 0.001
	}
	if r.Doctrines.Raider {
		primaryProb += 0.01
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		primaryProb -= 0.005
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
		primaryProb -= 0.005
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
		primaryProb -= 0.005
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
		primaryProb -= 0.005
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
		primaryProb += 0.005
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		primaryProb -= 0.005
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
		primaryProb -= 0.005
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
		primaryProb += 0.01
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
		primaryProb += 0.05
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
		primaryProb -= 0.01
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

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getAnimalSacrifice(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.3
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.55
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.3
	case r.Doctrines.Base.Deism:
		primaryProb = 0.05
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.5
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.55
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.59
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.05
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.07
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.04
	}

	if r.Doctrines.FullTolerance {
		primaryProb += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
		primaryProb += 0.05
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
		primaryProb += 0.1
	}
	if r.Doctrines.Esotericism {
		primaryProb += 0.05
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
		primaryProb += 0.1
	}
	if r.Doctrines.SanctityOfNature {
		primaryProb -= 0.01
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		primaryProb += 0.05
	}
	if r.Doctrines.AncestorWorship {
		primaryProb += 0.1
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
		primaryProb -= 0.01
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		primaryProb += 0.2
	}
	if r.Doctrines.MoonWorship {
		primaryProb += 0.2
	}
	if r.Doctrines.PainIsVirtue {
		primaryProb += 0.3
	}
	if r.Doctrines.Darkness {
		primaryProb += 0.3
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
		primaryProb -= 0.2
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
		primaryProb += 0.05
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
		primaryProb -= 0.005
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
		primaryProb += 0.02
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsVirtue():
		primaryProb -= 0.005
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
		primaryProb += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
		primaryProb += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
		primaryProb += 0.005
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
		primaryProb += 0.02
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
		primaryProb += 0.1
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
		primaryProb += 0.05
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
		primaryProb -= 0.05
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
		primaryProb += 0.2
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
		primaryProb += 0.1
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
		primaryProb += 0.01
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

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getHumanSacrifice(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.01
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.02
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.015
	case r.Doctrines.Base.Deism:
		primaryProb = 0.009
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.02
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.025
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.02
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.009
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.01
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.008
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
		primaryProb += 0.01
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
		primaryProb += 0.01
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
	}
	if r.Doctrines.Pacifism {
		primaryProb -= 0.05
	}
	if r.Doctrines.Reincarnation {
		primaryProb += 0.01
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
		primaryProb -= 0.1
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		primaryProb += 0.05
	}
	if r.Doctrines.MoonWorship {
	}
	if r.Doctrines.PainIsVirtue {
		primaryProb += 0.1
	}
	if r.Doctrines.Darkness {
		primaryProb += 0.2
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
		primaryProb += 0.02
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
		primaryProb += 0.01
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
		primaryProb += 0.05
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
		primaryProb -= 0.01
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
		primaryProb += 0.03
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
		primaryProb += 0.1
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
		primaryProb -= 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
		primaryProb -= 0.01
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

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getPlantsSacrifice(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.4
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.6
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.5
	case r.Doctrines.Base.Deism:
		primaryProb = 0.6
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.6
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.6
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.609
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.15
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.22
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.3
	}

	if r.Doctrines.FullTolerance {
		primaryProb += 0.1
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
		primaryProb += 0.1
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
		primaryProb += 0.05
	}
	if r.Doctrines.Esotericism {
		primaryProb += 0.1
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
		primaryProb += 0.08
	}
	if r.Doctrines.SanctityOfNature {
		primaryProb += 0.01
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
		primaryProb += 0.15
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
		primaryProb += 0.1
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		primaryProb += 0.2
	}
	if r.Doctrines.MoonWorship {
		primaryProb += 0.2
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
		primaryProb -= 0.05
	}
	if r.Doctrines.AnimalConnection {
		primaryProb += 0.01
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
		primaryProb += 0.03
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsVirtue():
		primaryProb += 0.1
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Charity.IsSin():
		primaryProb -= 0.1
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

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getAdorcism(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.01
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.05
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.01
	case r.Doctrines.Base.Deism:
		primaryProb = 0.051
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.02
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.03
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.055
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.02
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.01
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.03
	}

	if r.Doctrines.FullTolerance {
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
		primaryProb -= 0.01
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
		primaryProb += 0.01
	}
	if r.Doctrines.Esotericism {
		primaryProb += 0.01
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
		primaryProb += 0.01
	}
	if r.Doctrines.ReligiousLaw {
		primaryProb -= 0.05
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		primaryProb += 0.01
	}
	if r.Doctrines.AncestorWorship {
		primaryProb += 0.01
	}
	if r.Doctrines.Pacifism {
	}
	if r.Doctrines.Reincarnation {
		primaryProb += 0.01
	}
	if r.Doctrines.RitualHospitality {
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		primaryProb += 0.001
	}
	if r.Doctrines.MoonWorship {
		primaryProb += 0.001
	}
	if r.Doctrines.PainIsVirtue {
		primaryProb += 0.01
	}
	if r.Doctrines.Darkness {
		primaryProb += 0.2
	}
	if r.Doctrines.LiveUnderGround {
		primaryProb += 0.01
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
		primaryProb += 0.01
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
		primaryProb += 0.01
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
		primaryProb += 0.02
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
		primaryProb -= 0.01
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsSin():
		primaryProb += 0.03
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsVirtue():
		primaryProb += 0.005
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
		primaryProb += 0.05
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
		primaryProb -= 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
		primaryProb -= 0.01
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
		primaryProb += 0.05
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
		primaryProb -= 0.03
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
		primaryProb += 0.1
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
		primaryProb -= 0.03
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
		primaryProb -= 0.1
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
		primaryProb += 0.1
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

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getExorcism(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.3
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.15
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.29
	case r.Doctrines.Base.Deism:
		primaryProb = 0.28
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.19
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.14
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.15
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.2
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.18
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.175
	}

	if r.Doctrines.FullTolerance {
	}
	if r.Doctrines.Messiah {
		primaryProb += 0.05
	}
	if r.Doctrines.Prophets {
		primaryProb += 0.02
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
	}
	if r.Doctrines.Esotericism {
		primaryProb += 0.01
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
		primaryProb += 0.05
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		primaryProb += 0.15
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
		primaryProb += 0.015
	}
	if r.Doctrines.MoonWorship {
		primaryProb += 0.01
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
		primaryProb -= 0.01
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
		primaryProb += 0.01
	}
	if r.Doctrines.Raider {
	}
	if r.Doctrines.Proselytizer {
		primaryProb += 0.21
	}
	if r.Doctrines.Hedonism {
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		primaryProb -= 0.04
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
		primaryProb += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
		primaryProb -= 0.05
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
		primaryProb += 0.01
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
		primaryProb += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
		primaryProb -= 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
		primaryProb -= 0.005
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
		primaryProb -= 0.01
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
		primaryProb += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsVirtue():
		primaryProb -= 0.005
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
		primaryProb += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
		primaryProb += 0.005
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
		primaryProb -= 0.005
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
		primaryProb += 0.03
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
		primaryProb += 0.15
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
		primaryProb -= 0.09
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

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getOrgy(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.05
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.25
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.06
	case r.Doctrines.Base.Deism:
		primaryProb = 0.15
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.24
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.25
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.2
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.1
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.09
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.09
	}

	if r.Doctrines.FullTolerance {
		primaryProb += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
		primaryProb -= 0.02
	}
	if r.Doctrines.Astrology {
		primaryProb += 0.01
	}
	if r.Doctrines.Esotericism {
		primaryProb += 0.01
	}
	if r.Doctrines.Legalism {
		primaryProb -= 0.01
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
		primaryProb -= 0.04
	}
	if r.Doctrines.Polyamory {
		primaryProb += 0.3
	}
	if r.Doctrines.ReligiousLaw {
	}
	if r.Doctrines.SanctityOfNature {
		primaryProb += 0.01
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
		primaryProb += 0.005
	}
	if r.Doctrines.SacredChildbirth {
		primaryProb += 0.01
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		primaryProb += 0.01
	}
	if r.Doctrines.MoonWorship {
		primaryProb += 0.02
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
		primaryProb += 0.09
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
		primaryProb += 0.05
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
		primaryProb += 0.25
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		primaryProb += 0.05
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
		primaryProb += 0.05
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
		primaryProb -= 0.01
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
		primaryProb += 0.03
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
		primaryProb -= 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
		primaryProb -= 0.01
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
		primaryProb += 0.01
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
		primaryProb += 0.35
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Lust.IsSin():
		primaryProb -= 0.3
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsVirtue():
		primaryProb -= 0.2
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Chaste.IsSin():
		primaryProb += 0.3
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
		primaryProb -= 0.05
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
		primaryProb += 0.05
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getSmokeCircle(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.15
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.25
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.16
	case r.Doctrines.Base.Deism:
		primaryProb = 0.26
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.25
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.26
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.265
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.1
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.15
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.09
	}

	if r.Doctrines.FullTolerance {
		primaryProb += 0.01
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
	}
	if r.Doctrines.Astrology {
		primaryProb += 0.02
	}
	if r.Doctrines.Esotericism {
		primaryProb += 0.2
	}
	if r.Doctrines.Legalism {
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
		primaryProb += 0.01
	}
	if r.Doctrines.ReligiousLaw {
	}
	if r.Doctrines.SanctityOfNature {
		primaryProb += 0.15
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
	}
	if r.Doctrines.AncestorWorship {
		primaryProb += 0.2
	}
	if r.Doctrines.Pacifism {
		primaryProb += 0.01
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
		primaryProb += 0.05
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
		primaryProb += 0.01
	}
	if r.Doctrines.SunWorship {
		primaryProb += 0.01
	}
	if r.Doctrines.MoonWorship {
		primaryProb += 0.01
	}
	if r.Doctrines.PainIsVirtue {
		primaryProb += 0.005
	}
	if r.Doctrines.Darkness {
		primaryProb += 0.01
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
		primaryProb += 0.1
	}
	if r.Doctrines.Raider {
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
		primaryProb += 0.01
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
		primaryProb += 0.01
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
		primaryProb += 0.01
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
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
		primaryProb += 0.001
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
		primaryProb += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
		primaryProb += 0.01
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
		primaryProb += 0.01
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
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getGladiatorDuel(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.15
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.2
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.16
	case r.Doctrines.Base.Deism:
		primaryProb = 0.1
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.19
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.2
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.18
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.2
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.16
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.05
	}

	if r.Doctrines.FullTolerance {
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
		primaryProb += 0.01
	}
	if r.Doctrines.Astrology {
		primaryProb += 0.05
	}
	if r.Doctrines.Esotericism {
	}
	if r.Doctrines.Legalism {
		primaryProb += 0.1
	}
	if r.Doctrines.MendicantPreachers {
	}
	if r.Doctrines.Monasticism {
	}
	if r.Doctrines.Polyamory {
	}
	if r.Doctrines.ReligiousLaw {
		primaryProb += 0.05
	}
	if r.Doctrines.SanctityOfNature {
	}
	if r.Doctrines.TaxNonbelievers {
	}
	if r.Doctrines.UnrelentingFaith {
		primaryProb += 0.15
	}
	if r.Doctrines.AncestorWorship {
		primaryProb += 0.1
	}
	if r.Doctrines.Pacifism {
		primaryProb -= 0.05
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
		primaryProb += 0.01
	}
	if r.Doctrines.MoonWorship {
		primaryProb += 0.01
	}
	if r.Doctrines.PainIsVirtue {
	}
	if r.Doctrines.Darkness {
		primaryProb += 0.05
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
		primaryProb += 0.2
	}
	if r.Doctrines.Proselytizer {
		primaryProb += 0.02
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
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
		primaryProb -= 0.01
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
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Gluttony.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsVirtue():
		primaryProb -= 0.01
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Temperance.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsVirtue():
		primaryProb += 0.08
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pride.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Humility.IsVirtue():
		primaryProb -= 0.02
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
		primaryProb += 0.25
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Wrath.IsSin():
		primaryProb -= 0.1
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsVirtue():
		primaryProb -= 0.01
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Patience.IsSin():
		primaryProb += 0.09
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsVirtue():
		primaryProb += 0.04
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
		primaryProb -= 0.1
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
		primaryProb += 0.2
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
		primaryProb -= 0.05
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
		primaryProb -= 0.05
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
		primaryProb += 0.06
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
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Honest.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsVirtue():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HonorParents.IsSin():
	}

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getBlinding(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.01
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.02
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.01
	case r.Doctrines.Base.Deism:
		primaryProb = 0.01
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.02
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.02
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.01
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.01
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.019
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.013
	}

	if r.Doctrines.FullTolerance {
	}
	if r.Doctrines.Messiah {
	}
	if r.Doctrines.Prophets {
	}
	if r.Doctrines.Asceticism {
		primaryProb += 0.01
	}
	if r.Doctrines.Astrology {
		primaryProb -= 0.2
	}
	if r.Doctrines.Esotericism {
		primaryProb += 0.01
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
		primaryProb += 0.01
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
		primaryProb -= 0.1
	}
	if r.Doctrines.MoonWorship {
		primaryProb -= 0.1
	}
	if r.Doctrines.PainIsVirtue {
		primaryProb += 0.09
	}
	if r.Doctrines.Darkness {
		primaryProb += 0.1
	}
	if r.Doctrines.LiveUnderGround {
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
		primaryProb += 0.4
	}
	if r.Doctrines.Raider {
	}
	if r.Doctrines.Proselytizer {
	}
	if r.Doctrines.Hedonism {
		primaryProb -= 0.1
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsVirtue():
		primaryProb -= 0.01
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Profanity.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsVirtue():
		primaryProb -= 0.01
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.HaveOtherGods.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsVirtue():
		primaryProb -= 0.01
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sloth.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Greed.IsVirtue():
		primaryProb -= 0.01
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
		primaryProb -= 0.01
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
		primaryProb -= 0.02
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
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
		primaryProb -= 0.02
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

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

func getRitualCannibalism(r *entities.Religion) bool {
	var primaryProb float64
	switch {
	case r.Doctrines.Base.Monotheism:
		primaryProb = 0.005
	case r.Doctrines.Base.Polytheism:
		primaryProb = 0.01
	case r.Doctrines.Base.DeityDualism:
		primaryProb = 0.006
	case r.Doctrines.Base.Deism:
		primaryProb = 0.007
	case r.Doctrines.Base.Henothism:
		primaryProb = 0.009
	case r.Doctrines.Base.Monolatry:
		primaryProb = 0.01
	case r.Doctrines.Base.Omnism:
		primaryProb = 0.009
	}

	switch {
	case r.Doctrines.Gender.MaleDominance:
		primaryProb += 0.003
	case r.Doctrines.Gender.Equality:
		primaryProb += 0.004
	case r.Doctrines.Gender.FemaleDominance:
		primaryProb += 0.002
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
		primaryProb += 0.01
	}
	if r.Doctrines.Esotericism {
		primaryProb += 0.01
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
		primaryProb += 0.01
	}
	if r.Doctrines.AncestorWorship {
		primaryProb += 0.01
	}
	if r.Doctrines.Pacifism {
		primaryProb -= 0.05
	}
	if r.Doctrines.Reincarnation {
	}
	if r.Doctrines.RitualHospitality {
		primaryProb -= 0.05
	}
	if r.Doctrines.SacredChildbirth {
	}
	if r.Doctrines.SanctionedFalseConversions {
	}
	if r.Doctrines.SunWorship {
		primaryProb += 0.003
	}
	if r.Doctrines.MoonWorship {
		primaryProb += 0.005
	}
	if r.Doctrines.PainIsVirtue {
		primaryProb += 0.01
	}
	if r.Doctrines.Darkness {
		primaryProb += 0.1
	}
	if r.Doctrines.LiveUnderGround {
		primaryProb += 0.01
	}
	if r.Doctrines.TreeConnection {
	}
	if r.Doctrines.AnimalConnection {
	}
	if r.Doctrines.Blindsight {
	}
	if r.Doctrines.Raider {
		primaryProb += 0.01
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
		primaryProb += 0.02
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
		primaryProb += 0.03
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
		primaryProb += 0.02
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
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Pain.IsSin():
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsVirtue():
		primaryProb += 0.01
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Sadism.IsSin():
		primaryProb -= 0.03
	}

	switch {
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsVirtue():
		primaryProb -= 0.05
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsNeutral():
	case r.Theology.Precepts.SinsAndVirtues.Empathy.IsSin():
		primaryProb += 0.06
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

	return pm.GetRandomBool(pm.PrepareProbability(primaryProb))
}

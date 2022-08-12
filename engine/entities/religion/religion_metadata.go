package religion

import (
	"fmt"

	pm "persons_generator/engine/probability_machine"
)

type Metadata struct {
	BaseCoef     float64
	LowBaseCoef  float64
	HighBaseCoef float64

	MaxMetadataValue float64
}

type religionMetadata struct {
	Naturalistic float64

	SexualActive     float64
	SexualStrictness float64

	Chthonic    float64
	Plutocratic float64
	Altruistic  float64
	Lawful      float64
	Educational float64

	Aggressive float64
	Pacifistic float64

	Hedonistic float64
	Ascetic    float64

	Authoritaristic float64
	Liberal         float64

	Individualistic float64
	Collectivistic  float64

	Simple      float64
	Complicated float64
}

func (rm *religionMetadata) IsNaturalistic() bool {
	if rm == nil {
		return false
	}

	return rm.Naturalistic > 2.5
}

func (rm *religionMetadata) IsSexualActive() bool {
	if rm == nil {
		return false
	}

	return rm.SexualActive >= 2
}

func (rm *religionMetadata) IsSexualStrictness() bool {
	if rm == nil {
		return false
	}

	return rm.SexualActive >= 2
}

func (rm *religionMetadata) IsChthonic() bool {
	if rm == nil {
		return false
	}

	return rm.Chthonic >= 2
}

func (rm *religionMetadata) IsPlutocratic() bool {
	if rm == nil {
		return false
	}

	return rm.Plutocratic >= 2
}

func (rm *religionMetadata) IsAltruistic() bool {
	if rm == nil {
		return false
	}

	return rm.Altruistic >= 2
}

func (rm *religionMetadata) IsLawful() bool {
	if rm == nil {
		return false
	}

	return rm.Lawful >= 2
}

func (rm *religionMetadata) IsEducational() bool {
	if rm == nil {
		return false
	}

	return rm.Educational >= 2
}

func (rm *religionMetadata) IsAggressive() bool {
	if rm == nil {
		return false
	}

	return rm.Aggressive >= 2
}

func (rm *religionMetadata) IsPacifistic() bool {
	if rm == nil {
		return false
	}

	return rm.Pacifistic >= 2
}

func (rm *religionMetadata) IsHedonistic() bool {
	if rm == nil {
		return false
	}

	return rm.Hedonistic >= 2
}

func (rm *religionMetadata) IsAscetic() bool {
	if rm == nil {
		return false
	}

	return rm.Ascetic >= 2
}

func (rm *religionMetadata) IsAuthoritaristic() bool {
	if rm == nil {
		return false
	}

	return rm.Authoritaristic >= 2
}

func (rm *religionMetadata) IsLiberal() bool {
	if rm == nil {
		return false
	}

	return rm.Liberal >= 2
}

func (rm *religionMetadata) IsIndividualistic() bool {
	if rm == nil {
		return false
	}

	return rm.Individualistic >= 2
}

func (rm *religionMetadata) IsCollectivistic() bool {
	if rm == nil {
		return false
	}

	return rm.Collectivistic >= 2
}

func (rm *religionMetadata) IsSimple() bool {
	if rm == nil {
		return false
	}

	return rm.Simple >= 2
}

func (rm *religionMetadata) IsComplicated() bool {
	if rm == nil {
		return false
	}

	return rm.Complicated >= 2
}

func (r *Religion) generateMetadata() *religionMetadata {
	rm := &religionMetadata{}
	switch {
	case r.Type.IsMonotheism():
		rm.Lawful += pm.RandFloat64InRange(0.01, 0.1)
		rm.Authoritaristic += pm.RandFloat64InRange(0.01, 0.05)
	case r.Type.IsPolytheism():
		rm.Naturalistic += pm.RandFloat64InRange(0.01, 0.25)
		rm.Liberal += pm.RandFloat64InRange(0.01, 0.05)
		rm.Collectivistic += pm.RandFloat64InRange(0.01, 0.05)
	case r.Type.IsDeityDualism():
	case r.Type.IsDeism():
		rm.Naturalistic += pm.RandFloat64InRange(0.01, 0.25)
		rm.Liberal += pm.RandFloat64InRange(0.01, 0.075)
	case r.Type.IsAtheism():
		rm.Liberal += pm.RandFloat64InRange(0.01, 0.1)
	}

	return rm
}

func updateReligionMetadata(r *Religion, rm, u religionMetadata, coef float64) *religionMetadata {
	if u.Naturalistic > 0 {
		rm.Naturalistic = prepareMetadataValueForUpdate(r, rm.Naturalistic+u.Naturalistic*coef)
	}

	if u.SexualActive > 0 {
		rm.SexualActive = prepareMetadataValueForUpdate(r, rm.SexualActive+u.SexualActive*coef)
	}
	if u.SexualStrictness > 0 {
		rm.SexualStrictness = prepareMetadataValueForUpdate(r, rm.SexualStrictness+u.SexualStrictness*coef)
	}

	if u.Chthonic > 0 {
		rm.Chthonic = prepareMetadataValueForUpdate(r, rm.Chthonic+u.Chthonic*coef)
	}
	if u.Plutocratic > 0 {
		rm.Plutocratic = prepareMetadataValueForUpdate(r, rm.Plutocratic+u.Plutocratic*coef)
	}
	if u.Altruistic > 0 {
		rm.Altruistic = prepareMetadataValueForUpdate(r, rm.Altruistic+u.Altruistic*coef)
	}
	if u.Lawful > 0 {
		rm.Lawful = prepareMetadataValueForUpdate(r, rm.Lawful+u.Lawful*coef)
	}
	if u.Educational > 0 {
		rm.Educational = prepareMetadataValueForUpdate(r, rm.Educational+u.Educational*coef)
	}

	if u.Aggressive > 0 {
		rm.Aggressive = prepareMetadataValueForUpdate(r, rm.Aggressive+u.Aggressive*coef)
	}
	if u.Pacifistic > 0 {
		rm.Pacifistic = prepareMetadataValueForUpdate(r, rm.Pacifistic+u.Pacifistic*coef)
	}

	if u.Hedonistic > 0 {
		rm.Hedonistic = prepareMetadataValueForUpdate(r, rm.Hedonistic+u.Hedonistic*coef)
	}
	if u.Ascetic > 0 {
		rm.Ascetic = prepareMetadataValueForUpdate(r, rm.Ascetic+u.Ascetic*coef)
	}

	if u.Authoritaristic > 0 {
		rm.Authoritaristic = prepareMetadataValueForUpdate(r, rm.Authoritaristic+u.Authoritaristic*coef)
	}
	if u.Liberal > 0 {
		rm.Liberal = prepareMetadataValueForUpdate(r, rm.Liberal+u.Liberal*coef)
	}

	if u.Individualistic > 0 {
		rm.Individualistic = prepareMetadataValueForUpdate(r, rm.Individualistic+u.Individualistic*coef)
	}
	if u.Collectivistic > 0 {
		rm.Collectivistic = prepareMetadataValueForUpdate(r, rm.Collectivistic+u.Collectivistic*coef)
	}

	if u.Simple > 0 {
		rm.Simple = prepareMetadataValueForUpdate(r, rm.Simple+u.Simple*coef)
	}
	if u.Complicated > 0 {
		rm.Complicated = prepareMetadataValueForUpdate(r, rm.Complicated+u.Complicated*coef)
	}

	return &rm
}

func UpdateReligionMetadata(r *Religion, rm, u religionMetadata) *religionMetadata {
	return updateReligionMetadata(r, rm, u, 1)
}

func UpdateReligionMetadataWithAcceptance(r *Religion, rm, u religionMetadata, a Acceptance) *religionMetadata {
	var coef float64
	switch a {
	case Accepted:
		coef = pm.RandFloat64InRange(0.3, 0.5)
	case Shunned:
		coef = -pm.RandFloat64InRange(0.05, 0.25)
	case Criminal:
		coef = -pm.RandFloat64InRange(0.3, 0.5)
	}
	return updateReligionMetadata(r, rm, u, coef)
}

type CalcProbOpts struct {
	Log   bool
	Label string
}

func getRMProbability(coef float64, isMatchSame, isMatchContrary bool) float64 {
	var (
		sameCoef     = pm.RandFloat64InRange(0.4, 0.8)
		contraryCoef = pm.RandFloat64InRange(0.3, 0.7)
		newCoef      = pm.RandFloat64InRange(0.5, 1)
	)
	if isMatchSame && isMatchContrary {
		return coef * pm.PrepareProbability(sameCoef-contraryCoef)
	}
	if isMatchSame && !isMatchContrary {
		return coef * sameCoef
	}
	if !isMatchSame && !isMatchContrary {
		return coef * sameCoef * newCoef
	}

	return 0
}

func CalculateProbabilityFromReligionMetadata(baseCoef float64, r *Religion, u *religionMetadata, opts CalcProbOpts) bool {
	baseCoef = pm.PrepareCoef(baseCoef)
	var (
		primaryProbability float64
		randCoef           = pm.RandFloat64InRange(0.9, 1.1)
		ideasCount         int
	)

	if u.Naturalistic > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsNaturalistic(), false)
		ideasCount++
	}

	if u.SexualActive > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsSexualActive(), r.metadata.IsSexualStrictness())
		ideasCount++
	}
	if u.SexualStrictness > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsSexualStrictness(), r.metadata.IsSexualActive())
		ideasCount++
	}

	if u.Chthonic > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsChthonic(), false)
		ideasCount++
	}
	if u.Plutocratic > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsPlutocratic(), r.metadata.IsAltruistic())
		ideasCount++
	}
	if u.Altruistic > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsAltruistic(), r.metadata.IsPlutocratic())
		ideasCount++
	}
	if u.Lawful > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsLawful(), false)
		ideasCount++
	}
	if u.Educational > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsEducational(), false)
		ideasCount++
	}

	if u.Aggressive > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsAggressive(), r.metadata.IsPacifistic())
		ideasCount++
	}
	if u.Pacifistic > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsPacifistic(), r.metadata.IsAggressive())
		ideasCount++
	}

	if u.Hedonistic > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsHedonistic(), r.metadata.IsAscetic())
		ideasCount++
	}
	if u.Ascetic > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsAscetic(), r.metadata.IsHedonistic())
		ideasCount++
	}

	if u.Authoritaristic > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsAuthoritaristic(), r.metadata.IsLiberal())
		ideasCount++
	}
	if u.Ascetic > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsLiberal(), r.metadata.IsAuthoritaristic())
		ideasCount++
	}

	if u.Individualistic > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsIndividualistic(), r.metadata.IsCollectivistic())
		ideasCount++
	}
	if u.Ascetic > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsCollectivistic(), r.metadata.IsIndividualistic())
		ideasCount++
	}

	if u.Simple > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsSimple(), r.metadata.IsComplicated())
		ideasCount++
	}
	if u.Ascetic > 0 {
		primaryProbability += getRMProbability(randCoef, r.metadata.IsComplicated(), r.metadata.IsSimple())
		ideasCount++
	}

	probability := baseCoef * primaryProbability
	if ideasCount > 0 {
		probability = probability / float64(ideasCount)
	}

	if opts.Log {
		fmt.Printf("\n>>>>>>>>>>\nLabel: %s\nprobability: %f\n<<<<<<<<<<<<\n", opts.Label, probability)
	}

	return pm.GetRandomBool(pm.PrepareProbability(probability))
}

func getRMAcceptanceProbability(coef float64, isMatchSame, isMatchContrary bool) (float64, float64, float64) {
	var (
		sameCoef     = pm.RandFloat64InRange(0.4, 0.8)
		contraryCoef = pm.RandFloat64InRange(0.3, 0.7)
		newCoef      = pm.RandFloat64InRange(0.5, 1)

		newProbability = coef * sameCoef * newCoef
	)
	if isMatchSame && isMatchContrary {
		return coef * sameCoef, coef * pm.PrepareProbability(sameCoef-contraryCoef), coef * contraryCoef
	}
	if isMatchSame && !isMatchContrary {
		return coef * sameCoef, newProbability, 0
	}
	if !isMatchSame && isMatchContrary {
		return 0, newProbability, coef * contraryCoef
	}
	if !isMatchSame && !isMatchContrary {
		return newProbability, newProbability, newProbability
	}

	return 0, 0, 0
}

func CalculateAcceptanceFromReligionMetadata(acceptedBaseCoef, stunnedBaseCoef, criminalBaseCoef float64, r *Religion, u *religionMetadata, opts CalcProbOpts) Acceptance {
	acceptedBaseCoef = pm.PrepareCoef(acceptedBaseCoef)
	stunnedBaseCoef = pm.PrepareCoef(stunnedBaseCoef)
	criminalBaseCoef = pm.PrepareCoef(criminalBaseCoef)
	var (
		accepted, shunned, criminal float64
		randCoef                    = pm.RandFloat64InRange(0.9, 1.1)
		ideasCount                  int
	)

	if u.Naturalistic > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsNaturalistic(), false)
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.SexualActive > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsSexualActive(), false)
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Chthonic > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsChthonic(), false)
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Plutocratic > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsPlutocratic(), r.metadata.IsAltruistic())
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Altruistic > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsAltruistic(), r.metadata.IsPlutocratic())
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Lawful > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsLawful(), false)
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Educational > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsEducational(), false)
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if u.Aggressive > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsAggressive(), r.metadata.IsPacifistic())
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Pacifistic > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsPacifistic(), r.metadata.IsAggressive())
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if u.Hedonistic > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsHedonistic(), r.metadata.IsAscetic())
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Ascetic > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsAscetic(), r.metadata.IsHedonistic())
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if u.Authoritaristic > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsAuthoritaristic(), r.metadata.IsLiberal())
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Ascetic > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsLiberal(), r.metadata.IsAuthoritaristic())
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if u.Individualistic > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsIndividualistic(), r.metadata.IsCollectivistic())
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Ascetic > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsCollectivistic(), r.metadata.IsIndividualistic())
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if u.Simple > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsSimple(), r.metadata.IsComplicated())
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Ascetic > 0 {
		acc, shun, crim := getRMAcceptanceProbability(randCoef, r.metadata.IsComplicated(), r.metadata.IsSimple())
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	accepted = accepted / float64(ideasCount)
	shunned = shunned / float64(ideasCount)
	criminal = criminal / float64(ideasCount)

	accepted = pm.PrepareProbability(acceptedBaseCoef * accepted)
	shunned = pm.PrepareProbability(stunnedBaseCoef * shunned)
	criminal = pm.PrepareProbability(criminalBaseCoef * criminal)
	if opts.Log {
		fmt.Printf("\n>>>>>>>>>>\nLabel: %s\naccepted: %f, shunned: %f, criminal: %f\n<<<<<<<<<<<<\n", opts.Label, accepted, shunned, criminal)
	}

	return getAcceptanceByProbability(accepted, shunned, criminal)
}

func prepareMetadataValueForUpdate(r *Religion, v float64) float64 {
	var out float64
	if v > 0 {
		out = v
	}

	switch rel := out / r.M.MaxMetadataValue; {
	case rel < 0.25:
		return out * pm.RandFloat64InRange(1, 1.01)
	case rel < 0.5:
		return out * pm.RandFloat64InRange(1, 1.001)
	case rel < 0.75:
		return out
	case rel < 0.9:
		return out * pm.RandFloat64InRange(0.999, 0.9999)
	case rel < 1:
		return out * pm.RandFloat64InRange(0.98, 0.99)
	default:
		return r.M.MaxMetadataValue
	}
}

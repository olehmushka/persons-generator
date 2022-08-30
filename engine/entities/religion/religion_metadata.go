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

func (r *Religion) generateMetadata() (*religionMetadata, error) {
	rm := &religionMetadata{}
	switch {
	case r.Type.IsMonotheism():
		pLawful, err := pm.RandFloat64InRange(0.01, 0.1)
		if err != nil {
			return nil, err
		}
		rm.Lawful += pLawful
		pAuthoritaristic, err := pm.RandFloat64InRange(0.01, 0.05)
		if err != nil {
			return nil, err
		}
		rm.Authoritaristic += pAuthoritaristic
	case r.Type.IsPolytheism():
		pNaturalistic, err := pm.RandFloat64InRange(0.01, 0.25)
		if err != nil {
			return nil, err
		}
		rm.Naturalistic += pNaturalistic
		pLiberal, err := pm.RandFloat64InRange(0.01, 0.05)
		if err != nil {
			return nil, err
		}
		rm.Liberal += pLiberal
		pCollectivistic, err := pm.RandFloat64InRange(0.01, 0.05)
		if err != nil {
			return nil, err
		}
		rm.Collectivistic += pCollectivistic
	case r.Type.IsDeityDualism():
	case r.Type.IsDeism():
		pNaturalistic, err := pm.RandFloat64InRange(0.01, 0.25)
		if err != nil {
			return nil, err
		}
		rm.Naturalistic += pNaturalistic
		pLiberal, err := pm.RandFloat64InRange(0.01, 0.075)
		if err != nil {
			return nil, err
		}
		rm.Liberal += pLiberal
	case r.Type.IsAtheism():
		pLiberal, err := pm.RandFloat64InRange(0.01, 0.1)
		if err != nil {
			return nil, err
		}
		rm.Liberal += pLiberal
	}

	return rm, nil
}

func updateReligionMetadata(r *Religion, rm, u religionMetadata, coef float64) (*religionMetadata, error) {
	if u.Naturalistic > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Naturalistic+u.Naturalistic*coef)
		if err != nil {
			return nil, err
		}
		rm.Naturalistic = p
	}

	if u.SexualActive > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.SexualActive+u.SexualActive*coef)
		if err != nil {
			return nil, err
		}
		rm.SexualActive = p
	}
	if u.SexualStrictness > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.SexualStrictness+u.SexualStrictness*coef)
		if err != nil {
			return nil, err
		}
		rm.SexualStrictness = p
	}

	if u.Chthonic > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Chthonic+u.Chthonic*coef)
		if err != nil {
			return nil, err
		}
		rm.Chthonic = p
	}
	if u.Plutocratic > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Plutocratic+u.Plutocratic*coef)
		if err != nil {
			return nil, err
		}
		rm.Plutocratic = p
	}
	if u.Altruistic > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Altruistic+u.Altruistic*coef)
		if err != nil {
			return nil, err
		}
		rm.Altruistic = p
	}
	if u.Lawful > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Lawful+u.Lawful*coef)
		if err != nil {
			return nil, err
		}
		rm.Lawful = p
	}
	if u.Educational > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Educational+u.Educational*coef)
		if err != nil {
			return nil, err
		}
		rm.Educational = p
	}

	if u.Aggressive > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Aggressive+u.Aggressive*coef)
		if err != nil {
			return nil, err
		}
		rm.Aggressive = p
	}
	if u.Pacifistic > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Pacifistic+u.Pacifistic*coef)
		if err != nil {
			return nil, err
		}
		rm.Pacifistic = p
	}

	if u.Hedonistic > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Hedonistic+u.Hedonistic*coef)
		if err != nil {
			return nil, err
		}
		rm.Hedonistic = p
	}
	if u.Ascetic > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Ascetic+u.Ascetic*coef)
		if err != nil {
			return nil, err
		}
		rm.Ascetic = p
	}

	if u.Authoritaristic > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Authoritaristic+u.Authoritaristic*coef)
		if err != nil {
			return nil, err
		}
		rm.Authoritaristic = p
	}
	if u.Liberal > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Liberal+u.Liberal*coef)
		if err != nil {
			return nil, err
		}
		rm.Liberal = p
	}

	if u.Individualistic > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Individualistic+u.Individualistic*coef)
		if err != nil {
			return nil, err
		}
		rm.Individualistic = p
	}
	if u.Collectivistic > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Collectivistic+u.Collectivistic*coef)
		if err != nil {
			return nil, err
		}
		rm.Collectivistic = p
	}

	if u.Simple > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Simple+u.Simple*coef)
		if err != nil {
			return nil, err
		}
		rm.Simple = p
	}
	if u.Complicated > 0 {
		p, err := prepareMetadataValueForUpdate(r, rm.Complicated+u.Complicated*coef)
		if err != nil {
			return nil, err
		}
		rm.Complicated = p
	}

	return &rm, nil
}

func UpdateReligionMetadata(r *Religion, rm, u religionMetadata) (*religionMetadata, error) {
	return updateReligionMetadata(r, rm, u, 1)
}

func UpdateReligionMetadataWithAcceptance(r *Religion, rm, u religionMetadata, a Acceptance) (*religionMetadata, error) {
	var coef float64
	switch a {
	case Accepted:
		c, err := pm.RandFloat64InRange(0.3, 0.5)
		if err != nil {
			return nil, err
		}
		coef = c
	case Shunned:
		c, err := pm.RandFloat64InRange(0.05, 0.25)
		if err != nil {
			return nil, err
		}
		coef = -c
	case Criminal:
		c, err := pm.RandFloat64InRange(0.3, 0.5)
		if err != nil {
			return nil, err
		}
		coef = -c
	}
	return updateReligionMetadata(r, rm, u, coef)
}

type CalcProbOpts struct {
	Log   bool
	Label string
}

func getRMProbability(coef float64, isMatchSame, isMatchContrary bool) (float64, error) {
	sameCoef, err := pm.RandFloat64InRange(0.4, 0.8)
	if err != nil {
		return 0, err
	}
	contraryCoef, err := pm.RandFloat64InRange(0.3, 0.7)
	if err != nil {
		return 0, err
	}
	newCoef, err := pm.RandFloat64InRange(0.5, 1)
	if err != nil {
		return 0, err
	}

	if isMatchSame && isMatchContrary {
		return coef * pm.PrepareProbability(sameCoef-contraryCoef), nil
	}
	if isMatchSame && !isMatchContrary {
		return coef * sameCoef, nil
	}
	if !isMatchSame && !isMatchContrary {
		return coef * sameCoef * newCoef, nil
	}

	return 0, nil
}

func CalculateProbabilityFromReligionMetadata(baseCoef float64, r *Religion, u *religionMetadata, opts CalcProbOpts) (bool, error) {
	baseCoef = pm.PrepareCoef(baseCoef)
	var (
		primaryProbability float64
		ideasCount         int
	)
	randCoef, err := pm.RandFloat64InRange(0.9, 1.1)
	if err != nil {
		return false, err
	}

	if u.Naturalistic > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsNaturalistic(), false)
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if u.SexualActive > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsSexualActive(), r.metadata.IsSexualStrictness())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if u.SexualStrictness > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsSexualStrictness(), r.metadata.IsSexualActive())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if u.Chthonic > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsChthonic(), false)
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if u.Plutocratic > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsPlutocratic(), r.metadata.IsAltruistic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if u.Altruistic > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsAltruistic(), r.metadata.IsPlutocratic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if u.Lawful > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsLawful(), false)
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if u.Educational > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsEducational(), false)
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if u.Aggressive > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsAggressive(), r.metadata.IsPacifistic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if u.Pacifistic > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsPacifistic(), r.metadata.IsAggressive())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if u.Hedonistic > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsHedonistic(), r.metadata.IsAscetic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if u.Ascetic > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsAscetic(), r.metadata.IsHedonistic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if u.Authoritaristic > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsAuthoritaristic(), r.metadata.IsLiberal())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if u.Ascetic > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsLiberal(), r.metadata.IsAuthoritaristic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if u.Individualistic > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsIndividualistic(), r.metadata.IsCollectivistic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if u.Ascetic > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsCollectivistic(), r.metadata.IsIndividualistic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if u.Simple > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsSimple(), r.metadata.IsComplicated())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if u.Ascetic > 0 {
		p, err := getRMProbability(randCoef, r.metadata.IsComplicated(), r.metadata.IsSimple())
		if err != nil {
			return false, err
		}
		primaryProbability += p
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

func getRMAcceptanceProbability(coef float64, isMatchSame, isMatchContrary bool) (float64, float64, float64, error) {
	sameCoef, err := pm.RandFloat64InRange(0.4, 0.8)
	if err != nil {
		return 0, 0, 0, err
	}
	contraryCoef, err := pm.RandFloat64InRange(0.3, 0.7)
	if err != nil {
		return 0, 0, 0, err
	}
	newCoef, err := pm.RandFloat64InRange(0.5, 1)
	if err != nil {
		return 0, 0, 0, err
	}

	newProbability := coef * sameCoef * newCoef
	if isMatchSame && isMatchContrary {
		return coef * sameCoef, coef * pm.PrepareProbability(sameCoef-contraryCoef), coef * contraryCoef, nil
	}
	if isMatchSame && !isMatchContrary {
		return coef * sameCoef, newProbability, 0, nil
	}
	if !isMatchSame && isMatchContrary {
		return 0, newProbability, coef * contraryCoef, nil
	}
	if !isMatchSame && !isMatchContrary {
		return newProbability, newProbability, newProbability, nil
	}

	return 0, 0, 0, nil
}

func CalculateAcceptanceFromReligionMetadata(acceptedBaseCoef, stunnedBaseCoef, criminalBaseCoef float64, r *Religion, u *religionMetadata, opts CalcProbOpts) (Acceptance, error) {
	acceptedBaseCoef = pm.PrepareCoef(acceptedBaseCoef)
	stunnedBaseCoef = pm.PrepareCoef(stunnedBaseCoef)
	criminalBaseCoef = pm.PrepareCoef(criminalBaseCoef)
	var (
		accepted, shunned, criminal float64
		ideasCount                  int
	)
	randCoef, err := pm.RandFloat64InRange(0.9, 1.1)
	if err != nil {
		return "", err
	}

	if u.Naturalistic > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsNaturalistic(), false)
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.SexualActive > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsSexualActive(), false)
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Chthonic > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsChthonic(), false)
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Plutocratic > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsPlutocratic(), r.metadata.IsAltruistic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Altruistic > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsAltruistic(), r.metadata.IsPlutocratic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Lawful > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsLawful(), false)
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Educational > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsEducational(), false)
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if u.Aggressive > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsAggressive(), r.metadata.IsPacifistic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Pacifistic > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsPacifistic(), r.metadata.IsAggressive())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if u.Hedonistic > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsHedonistic(), r.metadata.IsAscetic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Ascetic > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsAscetic(), r.metadata.IsHedonistic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if u.Authoritaristic > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsAuthoritaristic(), r.metadata.IsLiberal())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Ascetic > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsLiberal(), r.metadata.IsAuthoritaristic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if u.Individualistic > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsIndividualistic(), r.metadata.IsCollectivistic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Ascetic > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsCollectivistic(), r.metadata.IsIndividualistic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if u.Simple > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsSimple(), r.metadata.IsComplicated())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if u.Ascetic > 0 {
		acc, shun, crim, err := getRMAcceptanceProbability(randCoef, r.metadata.IsComplicated(), r.metadata.IsSimple())
		if err != nil {
			return "", err
		}
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

	return getAcceptanceByProbability(accepted, shunned, criminal), nil
}

func prepareMetadataValueForUpdate(r *Religion, v float64) (float64, error) {
	var out float64
	if v > 0 {
		out = v
	}

	switch rel := out / r.M.MaxMetadataValue; {
	case rel < 0.25:
		p, err := pm.RandFloat64InRange(1, 1.01)
		if err != nil {
			return 0, err
		}
		return out * p, nil
	case rel < 0.5:
		p, err := pm.RandFloat64InRange(1, 1.001)
		if err != nil {
			return 0, err
		}
		return out * p, nil
	case rel < 0.75:
		return out, nil
	case rel < 0.9:
		p, err := pm.RandFloat64InRange(0.999, 0.9999)
		if err != nil {
			return 0, err
		}
		return out * p, nil
	case rel < 1:
		p, err := pm.RandFloat64InRange(0.98, 0.99)
		if err != nil {
			return 0, err
		}
		return out * p, nil
	default:
		return r.M.MaxMetadataValue, nil
	}
}

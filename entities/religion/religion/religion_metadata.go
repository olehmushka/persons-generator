package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
)

type Metadata struct {
	BaseCoef     float64
	LowBaseCoef  float64
	HighBaseCoef float64
}

type religionMetadata struct {
	Naturalistic float64
	SexualActive float64
	Chthonic     float64
	Plutocratic  float64
	Altruistic   float64
	Lawful       float64
	Educational  float64

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

	return rm.Naturalistic > 2
}

func (rm *religionMetadata) IsSexualActive() bool {
	if rm == nil {
		return false
	}

	return rm.SexualActive > 2
}

func (rm *religionMetadata) IsChthonic() bool {
	if rm == nil {
		return false
	}

	return rm.Chthonic > 2
}

func (rm *religionMetadata) IsPlutocratic() bool {
	if rm == nil {
		return false
	}

	return rm.Plutocratic > 2
}

func (rm *religionMetadata) IsAltruistic() bool {
	if rm == nil {
		return false
	}

	return rm.Altruistic > 2
}

func (rm *religionMetadata) IsLawful() bool {
	if rm == nil {
		return false
	}

	return rm.Lawful > 2
}

func (rm *religionMetadata) IsEducational() bool {
	if rm == nil {
		return false
	}

	return rm.Educational > 2
}

func (rm *religionMetadata) IsAggressive() bool {
	if rm == nil {
		return false
	}

	return rm.Aggressive > 2
}

func (rm *religionMetadata) IsPacifistic() bool {
	if rm == nil {
		return false
	}

	return rm.Pacifistic > 2
}

func (rm *religionMetadata) IsHedonistic() bool {
	if rm == nil {
		return false
	}

	return rm.Hedonistic > 2
}

func (rm *religionMetadata) IsAscetic() bool {
	if rm == nil {
		return false
	}

	return rm.Ascetic > 2
}

func (rm *religionMetadata) IsAuthoritaristic() bool {
	if rm == nil {
		return false
	}

	return rm.Authoritaristic > 2
}

func (rm *religionMetadata) IsLiberal() bool {
	if rm == nil {
		return false
	}

	return rm.Liberal > 2
}

func (rm *religionMetadata) IsIndividualistic() bool {
	if rm == nil {
		return false
	}

	return rm.Individualistic > 2
}

func (rm *religionMetadata) IsCollectivistic() bool {
	if rm == nil {
		return false
	}

	return rm.Collectivistic > 2
}

func (rm *religionMetadata) IsSimple() bool {
	if rm == nil {
		return false
	}

	return rm.Simple > 2
}

func (rm *religionMetadata) IsComplicated() bool {
	if rm == nil {
		return false
	}

	return rm.Complicated > 2
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

func UpdateReligionMetadata(rm, u religionMetadata) *religionMetadata {
	if u.Naturalistic > 0 {
		rm.Naturalistic += u.Naturalistic
	}
	if u.SexualActive > 0 {
		rm.SexualActive += u.SexualActive
	}
	if u.Chthonic > 0 {
		rm.Chthonic += u.Chthonic
	}
	if u.Plutocratic > 0 {
		rm.Plutocratic += u.Plutocratic
	}
	if u.Altruistic > 0 {
		rm.Altruistic += u.Altruistic
	}
	if u.Lawful > 0 {
		rm.Lawful += u.Lawful
	}
	if u.Educational > 0 {
		rm.Educational += u.Educational
	}

	if u.Aggressive > 0 {
		rm.Aggressive += u.Aggressive
	}
	if u.Pacifistic > 0 {
		rm.Pacifistic += u.Pacifistic
	}

	if u.Hedonistic > 0 {
		rm.Hedonistic += u.Hedonistic
	}
	if u.Ascetic > 0 {
		rm.Ascetic += u.Ascetic
	}

	if u.Authoritaristic > 0 {
		rm.Authoritaristic += u.Authoritaristic
	}
	if u.Liberal > 0 {
		rm.Liberal += u.Liberal
	}

	if u.Individualistic > 0 {
		rm.Individualistic += u.Individualistic
	}
	if u.Collectivistic > 0 {
		rm.Collectivistic += u.Collectivistic
	}

	if u.Simple > 0 {
		rm.Simple += u.Simple
	}
	if u.Complicated > 0 {
		rm.Complicated += u.Complicated
	}

	return &rm
}

type CalcProbOpts struct {
	Log bool
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
	baseCoef = pm.PrepareProbability(baseCoef)
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
		primaryProbability += getRMProbability(randCoef, r.metadata.IsSexualActive(), false)
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

	probability := pm.PrepareProbability(baseCoef * primaryProbability)
	if ideasCount > 0 {
		probability = probability / float64(ideasCount)
	}

	if opts.Log {
		fmt.Printf(">>>>>>>>>>\nprobability: %f\n<<<<<<<<<<<<", probability)
	}

	return pm.GetRandomBool(probability)
}

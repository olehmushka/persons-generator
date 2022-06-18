package religion

import (
	"fmt"

	pm "persons_generator/probability-machine"
)

type religionMetadata struct {
	Centralized   float64
	Decentralized float64

	RealLifeOriented  float64
	AfterlifeOriented float64

	InsideDirected  float64
	OutsideDirected float64

	Fanaticism float64
	Strictness float64

	// Specials traits
	Hedonism  float64
	Ascetic   float64
	Chthonic  float64
	Elitaric  float64
	Primitive float64
	Organized float64
}

func (r *Religion) generateMetadata() *religionMetadata {
	rm := &religionMetadata{}
	switch {
	case r.Type.IsMonotheism():
		rm.Centralized += pm.RandFloat64InRange(0.05, 0.15)
		rm.Decentralized += pm.RandFloat64InRange(0.01, 0.05)
		rm.RealLifeOriented += pm.RandFloat64InRange(0.02, 0.065)
		rm.AfterlifeOriented += pm.RandFloat64InRange(0.04, 0.1)
		rm.InsideDirected += pm.RandFloat64InRange(0.02, 0.08)
		rm.OutsideDirected += pm.RandFloat64InRange(0.02, 0.095)
		rm.Fanaticism += pm.RandFloat64InRange(0.03, 0.1)
		rm.Strictness += pm.RandFloat64InRange(0.02, 0.1)
	case r.Type.IsPolytheism():
		rm.Centralized += pm.RandFloat64InRange(0.01, 0.1)
		rm.Decentralized += pm.RandFloat64InRange(0.05, 0.12)
		rm.RealLifeOriented += pm.RandFloat64InRange(0.019, 0.05)
		rm.AfterlifeOriented += pm.RandFloat64InRange(0.02, 0.06)
		rm.InsideDirected += pm.RandFloat64InRange(0.015, 0.07)
		rm.OutsideDirected += pm.RandFloat64InRange(0.015, 0.085)
		rm.Fanaticism += pm.RandFloat64InRange(0.03, 0.1)
		rm.Strictness += pm.RandFloat64InRange(0.02, 0.08)
	case r.Type.IsDeityDualism():
		rm.Centralized += pm.RandFloat64InRange(0.02, 0.11)
		rm.Decentralized += pm.RandFloat64InRange(0.02, 0.08)
		rm.RealLifeOriented += pm.RandFloat64InRange(0.0195, 0.06)
		rm.AfterlifeOriented += pm.RandFloat64InRange(0.035, 0.065)
		rm.InsideDirected += pm.RandFloat64InRange(0.017, 0.075)
		rm.OutsideDirected += pm.RandFloat64InRange(0.017, 0.087)
		rm.Fanaticism += pm.RandFloat64InRange(0.029, 0.095)
		rm.Strictness += pm.RandFloat64InRange(0.02, 0.092)
	case r.Type.IsDeism():
		rm.Centralized += pm.RandFloat64InRange(0.01, 0.04)
		rm.Decentralized += pm.RandFloat64InRange(0.04, 0.13)
		rm.RealLifeOriented += pm.RandFloat64InRange(0.03, 0.08)
		rm.AfterlifeOriented += pm.RandFloat64InRange(0.005, 0.03)
		rm.InsideDirected += pm.RandFloat64InRange(0.03, 0.075)
		rm.OutsideDirected += pm.RandFloat64InRange(0.02, 0.08)
		rm.Fanaticism += pm.RandFloat64InRange(0.005, 0.03)
		rm.Strictness += pm.RandFloat64InRange(0.02, 0.06)
	case r.Type.IsAtheism():
		rm.Centralized += pm.RandFloat64InRange(0.02, 0.1)
		rm.Decentralized += pm.RandFloat64InRange(0.02, 0.1)
		rm.RealLifeOriented += pm.RandFloat64InRange(0.1, 0.2)
		rm.AfterlifeOriented += 0
		rm.InsideDirected += pm.RandFloat64InRange(0.01, 0.05)
		rm.OutsideDirected += pm.RandFloat64InRange(0.04, 0.097)
		rm.Fanaticism += pm.RandFloat64InRange(0.025, 0.07)
		rm.Strictness += pm.RandFloat64InRange(0.01, 0.06)
	}

	genderCoef := r.GenderDominance.GetCoef()
	switch {
	case r.GenderDominance.IsMaleDominate():
		rm.RealLifeOriented += genderCoef * pm.RandFloat64InRange(0.01, 0.03)
		rm.AfterlifeOriented += genderCoef * pm.RandFloat64InRange(0.005, 0.015)
		rm.InsideDirected += genderCoef * pm.RandFloat64InRange(0.005, 0.015)
		rm.OutsideDirected += genderCoef * pm.RandFloat64InRange(0.01, 0.03)
		rm.Fanaticism += genderCoef * pm.RandFloat64InRange(0.02, 0.04)
		rm.Strictness += genderCoef * pm.RandFloat64InRange(0.01, 0.03)
	case r.GenderDominance.IsEquality():
		rm.RealLifeOriented += genderCoef * pm.RandFloat64InRange(0.01, 0.02)
		rm.AfterlifeOriented += genderCoef * pm.RandFloat64InRange(0.01, 0.02)
		rm.InsideDirected += genderCoef * pm.RandFloat64InRange(0.01, 0.02)
		rm.OutsideDirected += genderCoef * pm.RandFloat64InRange(0.01, 0.02)
		rm.Fanaticism += genderCoef * pm.RandFloat64InRange(0.005, 0.015)
		rm.Strictness += genderCoef * pm.RandFloat64InRange(0.01, 0.02)
	case r.GenderDominance.IsFemaleDominate():
		rm.RealLifeOriented += genderCoef * pm.RandFloat64InRange(0.005, 0.015)
		rm.AfterlifeOriented += genderCoef * pm.RandFloat64InRange(0.01, 0.03)
		rm.InsideDirected += genderCoef * pm.RandFloat64InRange(0.01, 0.02)
		rm.OutsideDirected += genderCoef * pm.RandFloat64InRange(0.005, 0.015)
		rm.Fanaticism += genderCoef * pm.RandFloat64InRange(0.03, 0.04)
		rm.Strictness += genderCoef * pm.RandFloat64InRange(0.01, 0.02)
	}

	return rm
}

type updateReligionMetadata struct {
	Centralized   *float64
	Decentralized *float64

	RealLifeOriented  *float64
	AfterlifeOriented *float64

	InsideDirected  *float64
	OutsideDirected *float64

	Fanaticism *float64
	Strictness *float64

	// Specials traits
	Hedonism  *float64
	Ascetic   *float64
	Chthonic  *float64
	Elitaric  *float64
	Primitive *float64
	Organized *float64
}

func UpdateReligionMetadata(rm religionMetadata, u updateReligionMetadata) *religionMetadata {
	if u.Centralized != nil {
		rm.Centralized = CalculateReligionMetadataScoreIncrease(rm.Centralized, *u.Centralized)
	}
	if u.Decentralized != nil {
		rm.Decentralized = CalculateReligionMetadataScoreIncrease(rm.Decentralized, *u.Decentralized)
	}

	if u.RealLifeOriented != nil {
		rm.RealLifeOriented = CalculateReligionMetadataScoreIncrease(rm.RealLifeOriented, *u.RealLifeOriented)
	}
	if u.AfterlifeOriented != nil {
		rm.AfterlifeOriented = CalculateReligionMetadataScoreIncrease(rm.AfterlifeOriented, *u.AfterlifeOriented)
	}

	if u.InsideDirected != nil {
		rm.InsideDirected = CalculateReligionMetadataScoreIncrease(rm.InsideDirected, *u.InsideDirected)
	}
	if u.OutsideDirected != nil {
		rm.OutsideDirected = CalculateReligionMetadataScoreIncrease(rm.OutsideDirected, *u.OutsideDirected)
	}

	if u.Fanaticism != nil {
		rm.Fanaticism = CalculateReligionMetadataScoreIncrease(rm.Fanaticism, *u.Fanaticism)
	}
	if u.Strictness != nil {
		rm.Strictness = CalculateReligionMetadataScoreIncrease(rm.Strictness, *u.Strictness)
	}

	if u.Hedonism != nil {
		rm.Hedonism = CalculateReligionMetadataScoreIncrease(rm.Hedonism, *u.Hedonism)
	}
	if u.Ascetic != nil {
		rm.Ascetic = CalculateReligionMetadataScoreIncrease(rm.Ascetic, *u.Ascetic)
	}
	if u.Chthonic != nil {
		rm.Chthonic = CalculateReligionMetadataScoreIncrease(rm.Chthonic, *u.Chthonic)
	}
	if u.Elitaric != nil {
		rm.Elitaric = CalculateReligionMetadataScoreIncrease(rm.Elitaric, *u.Elitaric)
	}
	if u.Primitive != nil {
		rm.Primitive = CalculateReligionMetadataScoreIncrease(rm.Primitive, *u.Primitive)
	}
	if u.Organized != nil {
		rm.Organized = CalculateReligionMetadataScoreIncrease(rm.Organized, *u.Organized)
	}

	return &rm
}

type CalcProbOpts struct {
	Log bool
}

func CalculateProbabilityFromReligionMetadata(baseCoef float64, r *Religion, u updateReligionMetadata, opts CalcProbOpts) bool {
	calcFunc := func(base, inc float64) float64 {
		if base == 0 && inc == 0 {
			return 0
		}
		if base == 0 {
			return pm.RandFloat64InRange(0, 0.25) * baseCoef
		}

		var (
			probability   float64
			signCoef      = 0.5 // for inc > base
			percentOfDiff = inc / base
		)
		if base > inc {
			signCoef = 1
		}

		switch {
		case percentOfDiff > 1:
			probability = pm.RandFloat64InRange(1, 2)
		case percentOfDiff > 0.75 && percentOfDiff <= 1:
			probability = pm.RandFloat64InRange(0.5, 1)
		case percentOfDiff > 0.5 && percentOfDiff <= 0.75:
			probability = pm.RandFloat64InRange(0.25, 0.75)
		case percentOfDiff > 0.25 && percentOfDiff <= 0.5:
			probability = pm.RandFloat64InRange(0.1, 0.5)
		case percentOfDiff >= 0 && percentOfDiff <= 0.25:
			probability = pm.RandFloat64InRange(0, 0.25)
		}

		return signCoef * baseCoef * probability
	}

	var primaryProbability float64
	var treatsCount int
	if u.RealLifeOriented != nil {
		primaryProbability += calcFunc(r.metadata.RealLifeOriented, *u.RealLifeOriented)
		treatsCount++
	}
	if u.AfterlifeOriented != nil {
		primaryProbability += calcFunc(r.metadata.AfterlifeOriented, *u.AfterlifeOriented)
		treatsCount++
	}

	if u.InsideDirected != nil {
		primaryProbability += calcFunc(r.metadata.InsideDirected, *u.InsideDirected)
		treatsCount++
	}
	if u.OutsideDirected != nil {
		primaryProbability += calcFunc(r.metadata.OutsideDirected, *u.OutsideDirected)
		treatsCount++
	}

	if u.Fanaticism != nil {
		primaryProbability += calcFunc(r.metadata.Fanaticism, *u.Fanaticism)
		treatsCount++
	}
	if u.Strictness != nil {
		primaryProbability += calcFunc(r.metadata.Strictness, *u.Strictness)
		treatsCount++
	}

	if u.Hedonism != nil {
		primaryProbability += calcFunc(r.metadata.Hedonism, *u.Hedonism)
		treatsCount++
	}
	if u.Ascetic != nil {
		primaryProbability += calcFunc(r.metadata.Ascetic, *u.Ascetic)
		treatsCount++
	}
	if u.Chthonic != nil {
		primaryProbability += calcFunc(r.metadata.Chthonic, *u.Chthonic)
		treatsCount++
	}
	if u.Elitaric != nil {
		primaryProbability += calcFunc(r.metadata.Elitaric, *u.Elitaric)
		treatsCount++
	}
	if u.Primitive != nil {
		primaryProbability += calcFunc(r.metadata.Primitive, *u.Primitive)
		treatsCount++
	}
	if u.Organized != nil {
		primaryProbability += calcFunc(r.metadata.Organized, *u.Organized)
		treatsCount++
	}
	primaryProbability = primaryProbability / float64(treatsCount)
	if opts.Log {
		fmt.Printf("<<<<<\nPrimProb: %f\n>>>>>>>>", primaryProbability)
	}

	return pm.GetRandomBool(primaryProbability)
}

func CalculateReligionMetadataScoreIncrease(src, inc float64) float64 {
	incCoef := pm.RandFloat64InRange(0.05, 0.15)

	return src + inc*incCoef
}

func Float64(input float64) *float64 {
	return &input
}

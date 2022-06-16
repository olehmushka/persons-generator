package religion

import (
	"fmt"
	"math"

	pm "persons_generator/probability-machine"
)

type religionMetadata struct {
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
		rm.RealLifeOriented += 0.037
		rm.AfterlifeOriented += 0.05
		rm.InsideDirected += 0.02
		rm.OutsideDirected += 0.04
		rm.Fanaticism += 0.05
		rm.Strictness += 0.04
	case r.Type.IsPolytheism():
		rm.RealLifeOriented += 0.035
		rm.AfterlifeOriented += 0.02
		rm.InsideDirected += 0.015
		rm.OutsideDirected += 0.03
		rm.Fanaticism += 0.05
		rm.Strictness += 0.03
	case r.Type.IsDeityDualism():
		rm.RealLifeOriented += 0.036
		rm.AfterlifeOriented += 0.025
		rm.InsideDirected += 0.017
		rm.OutsideDirected += 0.035
		rm.Fanaticism += 0.049
		rm.Strictness += 0.035
	case r.Type.IsDeism():
		rm.RealLifeOriented += 0.045
		rm.AfterlifeOriented += 0.005
		rm.InsideDirected += 0.03
		rm.OutsideDirected += 0.025
		rm.Fanaticism += 0.01
		rm.Strictness += 0.036
	case r.Type.IsAtheism():
		rm.RealLifeOriented += 0.1
		rm.AfterlifeOriented += 0
		rm.InsideDirected += 0.01
		rm.OutsideDirected += 0.045
		rm.Fanaticism += 0.04
		rm.Strictness += 0.02
	}

	genderCoef := r.GenderDominance.GetCoef()
	switch {
	case r.GenderDominance.IsMaleDominate():
		rm.RealLifeOriented += 0.02 * genderCoef
		rm.AfterlifeOriented += 0.01 * genderCoef
		rm.InsideDirected += 0.01 * genderCoef
		rm.OutsideDirected += 0.02 * genderCoef
		rm.Fanaticism += 0.03 * genderCoef
		rm.Strictness += 0.02 * genderCoef
	case r.GenderDominance.IsEquality():
		rm.RealLifeOriented += 0.015 * genderCoef
		rm.AfterlifeOriented += 0.015 * genderCoef
		rm.InsideDirected += 0.015 * genderCoef
		rm.OutsideDirected += 0.015 * genderCoef
		rm.Fanaticism += 0.01 * genderCoef
		rm.Strictness += 0.015 * genderCoef
	case r.GenderDominance.IsFemaleDominate():
		rm.RealLifeOriented += 0.01 * genderCoef
		rm.AfterlifeOriented += 0.02 * genderCoef
		rm.InsideDirected += 0.015 * genderCoef
		rm.OutsideDirected += 0.01 * genderCoef
		rm.Fanaticism += 0.035 * genderCoef
		rm.Strictness += 0.015 * genderCoef
	}

	return rm
}

type updateReligionMetadata struct {
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
		if base == inc {
			return 1
		}
		result := (0.1 / math.Abs(base-inc))
		if result >= 2 {
			return 0.05 * result * baseCoef
		}

		return 0.1 * result * baseCoef
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
	return src + inc*0.1
}

func Float64(input float64) *float64 {
	return &input
}

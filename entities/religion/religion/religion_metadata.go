package religion

type religionMetadata struct {
	RealLifeOriented  float64
	AfterlifeOriented float64

	InsideDirected  float64
	OutsideDirected float64

	Fanaticism float64
	Strictness float64
	Hedonism   float64
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
		rm.Hedonism += 0.01
	case r.Type.IsPolytheism():
		rm.RealLifeOriented += 0.035
		rm.AfterlifeOriented += 0.02
		rm.InsideDirected += 0.015
		rm.OutsideDirected += 0.03
		rm.Fanaticism += 0.05
		rm.Strictness += 0.03
		rm.Hedonism += 0.05
	case r.Type.IsDeityDualism():
		rm.RealLifeOriented += 0.036
		rm.AfterlifeOriented += 0.025
		rm.InsideDirected += 0.017
		rm.OutsideDirected += 0.035
		rm.Fanaticism += 0.049
		rm.Strictness += 0.035
		rm.Hedonism += 0.02
	case r.Type.IsDeism():
		rm.RealLifeOriented += 0.045
		rm.AfterlifeOriented += 0.005
		rm.InsideDirected += 0.03
		rm.OutsideDirected += 0.025
		rm.Fanaticism += 0.01
		rm.Strictness += 0.036
		rm.Hedonism += 0.035
	case r.Type.IsAtheism():
		rm.RealLifeOriented += 0.1
		rm.AfterlifeOriented += 0
		rm.InsideDirected += 0.01
		rm.OutsideDirected += 0.045
		rm.Fanaticism += 0.04
		rm.Strictness += 0.02
		rm.Hedonism += 0.07
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
		rm.Hedonism += 0.02 * genderCoef
	case r.GenderDominance.IsEquality():
		rm.RealLifeOriented += 0.015 * genderCoef
		rm.AfterlifeOriented += 0.015 * genderCoef
		rm.InsideDirected += 0.015 * genderCoef
		rm.OutsideDirected += 0.015 * genderCoef
		rm.Fanaticism += 0.01 * genderCoef
		rm.Strictness += 0.015 * genderCoef
		rm.Hedonism += 0.01 * genderCoef
	case r.GenderDominance.IsFemaleDominate():
		rm.RealLifeOriented += 0.01 * genderCoef
		rm.AfterlifeOriented += 0.02 * genderCoef
		rm.InsideDirected += 0.015 * genderCoef
		rm.OutsideDirected += 0.01 * genderCoef
		rm.Fanaticism += 0.035 * genderCoef
		rm.Strictness += 0.015 * genderCoef
		rm.Hedonism += 0.01 * genderCoef
	}

	return rm
}

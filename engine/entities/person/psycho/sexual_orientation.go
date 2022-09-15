package psycho

import pm "persons_generator/engine/probability_machine"

type SexualOrientation string

func (so SexualOrientation) IsHeterosexual() bool {
	return so == Heterosexual
}

func (so SexualOrientation) IsHomosexual() bool {
	return so == Homosexual
}

func (so SexualOrientation) IsBisexual() bool {
	return so == Bisexual
}

func (so SexualOrientation) IsAsexual() bool {
	return so == Asexual
}

func ProduceSexualOrientation() SexualOrientation {
	var (
		heterosexual = 0.95
		homosexual   = 0.01
		bisexual     = 0.01
		asexual      = 0.01
	)

	return SexualOrientation(pm.GetRandomFromSeveral(map[string]float64{
		string(Heterosexual): pm.PrepareProbability(heterosexual),
		string(Homosexual):   pm.PrepareProbability(homosexual),
		string(Bisexual):     pm.PrepareProbability(bisexual),
		string(Asexual):      pm.PrepareProbability(asexual),
	}))
}

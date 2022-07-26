package religion

import pm "persons_generator/probability_machine"

type Tolerance string

const (
	FullTolerance     Tolerance = "FullTolerance"
	HighTolerance     Tolerance = "HighTolerance"
	ModerateTolerance Tolerance = "ModerateTolerance"
	LowTolerance      Tolerance = "LowTolerance"
	NoTolerance       Tolerance = "NoTolerance"
)

func getToleranceByProbability(fullTolerance, highTolerance, moderateTolerance, lowTolerance, noTolerance float64) Tolerance {
	return Tolerance(pm.GetRandomFromSeveral(map[string]float64{
		string(FullTolerance):     pm.PrepareProbability(fullTolerance),
		string(HighTolerance):     pm.PrepareProbability(highTolerance),
		string(ModerateTolerance): pm.PrepareProbability(moderateTolerance),
		string(LowTolerance):      pm.PrepareProbability(lowTolerance),
		string(NoTolerance):       pm.PrepareProbability(noTolerance),
	}))
}

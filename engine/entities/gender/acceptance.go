package gender

import pm "persons_generator/engine/probability_machine"

type Acceptance string

const (
	OnlyMen     Acceptance = "only_men"
	MenAndWomen Acceptance = "men_and_women"
	OnlyWomen   Acceptance = "only_women"
)

func GetAcceptanceByProbability(onlyMen, menAndWomen, onlyWomen float64) Acceptance {
	return Acceptance(pm.GetRandomFromSeveral(map[string]float64{
		string(OnlyMen):     pm.PrepareProbability(onlyMen),
		string(MenAndWomen): pm.PrepareProbability(menAndWomen),
		string(OnlyWomen):   pm.PrepareProbability(onlyWomen),
	}))
}

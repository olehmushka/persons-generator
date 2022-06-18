package religion

import pm "persons_generator/probability-machine"

type GenderAcceptance string

const (
	OnlyMen     GenderAcceptance = "only_men"
	MenAndWomen GenderAcceptance = "men_and_women"
	OnlyWomen   GenderAcceptance = "only_women"
)

func getGenderAcceptanceByProbability(onlyMen, menAndWomen, onlyWomen float64) GenderAcceptance {
	m := map[string]float64{
		string(OnlyMen):     pm.PrepareProbability(onlyMen),
		string(MenAndWomen): pm.PrepareProbability(menAndWomen),
		string(OnlyWomen):   pm.PrepareProbability(onlyWomen),
	}
	return GenderAcceptance(pm.GetRandomFromSeveral(m))
}

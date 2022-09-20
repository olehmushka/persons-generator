package gender

import (
	"fmt"
	"persons_generator/core/wrapped_error"
	pm "persons_generator/engine/probability_machine"
)

type Acceptance string

func (a Acceptance) String() string {
	return string(a)
}

const (
	OnlyMen     Acceptance = "only_men"
	MenAndWomen Acceptance = "men_and_women"
	OnlyWomen   Acceptance = "only_women"
)

func GetAcceptanceByProbability(onlyMen, menAndWomen, onlyWomen float64) (Acceptance, error) {
	out, err := pm.GetRandomFromSeveral(map[string]float64{
		string(OnlyMen):     pm.PrepareProbability(onlyMen),
		string(MenAndWomen): pm.PrepareProbability(menAndWomen),
		string(OnlyWomen):   pm.PrepareProbability(onlyWomen),
	})
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not generate acceptance (only_men=%f, men_and_woman=%f, only_women=%f)", onlyMen, menAndWomen, onlyWomen))
	}

	return Acceptance(out), nil
}

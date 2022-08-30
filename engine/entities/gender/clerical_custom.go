package gender

import pm "persons_generator/engine/probability_machine"

func GetClericalCustom(baseCoef float64, dom *Dominance) (Acceptance, error) {
	men, err := pm.RandFloat64InRange(0.05, 0.15)
	if err != nil {
		return "", err
	}
	menAndWomen, err := pm.RandFloat64InRange(0.05, 0.15)
	if err != nil {
		return "", err
	}
	women, err := pm.RandFloat64InRange(0.05, 0.15)
	if err != nil {
		return "", err
	}
	randCoef, err := pm.RandFloat64InRange(0.15, 25)
	if err != nil {
		return "", err
	}

	switch {
	case dom.IsMaleDominate():
		men += baseCoef * randCoef * dom.GetCoef()
	case dom.IsEquality():
		menAndWomen += baseCoef * randCoef * dom.GetCoef()
	case dom.IsFemaleDominate():
		women += baseCoef * randCoef * dom.GetCoef()
	}

	return GetAcceptanceByProbability(baseCoef*men, baseCoef*menAndWomen, baseCoef*women), nil
}

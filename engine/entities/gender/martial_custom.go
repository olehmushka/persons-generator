package gender

import pm "persons_generator/engine/probability_machine"

func GetMartialCustom(baseCoef float64, dom *Dominance) (Acceptance, error) {
	men, err := pm.RandFloat64InRange(0.25, 0.5)
	if err != nil {
		return "", err
	}
	menAndWomen, err := pm.RandFloat64InRange(0.03, 0.1)
	if err != nil {
		return "", err
	}
	women, err := pm.RandFloat64InRange(0.01, 0.075)
	if err != nil {
		return "", err
	}

	switch {
	case dom.IsMaleDominate():
		randCoef, err := pm.RandFloat64InRange(0.15, 25)
		if err != nil {
			return "", err
		}
		men += baseCoef * randCoef * dom.GetCoef()
	case dom.IsEquality():
		randCoef, err := pm.RandFloat64InRange(0.12, 22)
		if err != nil {
			return "", err
		}
		menAndWomen += baseCoef * randCoef * dom.GetCoef()
	case dom.IsFemaleDominate():
		randCoef, err := pm.RandFloat64InRange(0.1, 2)
		if err != nil {
			return "", err
		}
		women += baseCoef * randCoef * dom.GetCoef()
	}

	return GetAcceptanceByProbability(baseCoef*men, baseCoef*menAndWomen, baseCoef*women), nil
}

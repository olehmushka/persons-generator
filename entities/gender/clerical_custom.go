package gender

import pm "persons_generator/probability_machine"

func GetClericalCustom(baseCoef float64, dom *Dominance) Acceptance {
	var (
		men         = pm.RandFloat64InRange(0.05, 0.15)
		menAndWomen = pm.RandFloat64InRange(0.05, 0.15)
		women       = pm.RandFloat64InRange(0.05, 0.15)
	)
	switch {
	case dom.IsMaleDominate():
		men += baseCoef * pm.RandFloat64InRange(0.15, 25) * dom.GetCoef()
	case dom.IsEquality():
		menAndWomen += baseCoef * pm.RandFloat64InRange(0.15, 25) * dom.GetCoef()
	case dom.IsFemaleDominate():
		women += baseCoef * pm.RandFloat64InRange(0.15, 25) * dom.GetCoef()
	}

	return GetAcceptanceByProbability(baseCoef*men, baseCoef*menAndWomen, baseCoef*women)
}

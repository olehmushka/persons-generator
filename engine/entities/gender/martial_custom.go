package gender

import pm "persons_generator/engine/probability_machine"

func GetMartialCustom(baseCoef float64, dom *Dominance) Acceptance {
	var (
		men         = pm.RandFloat64InRange(0.25, 0.5)
		menAndWomen = pm.RandFloat64InRange(0.03, 0.1)
		women       = pm.RandFloat64InRange(0.01, 0.075)
	)
	switch {
	case dom.IsMaleDominate():
		men += baseCoef * pm.RandFloat64InRange(0.15, 25) * dom.GetCoef()
	case dom.IsEquality():
		menAndWomen += baseCoef * pm.RandFloat64InRange(0.12, 22) * dom.GetCoef()
	case dom.IsFemaleDominate():
		women += baseCoef * pm.RandFloat64InRange(0.1, 2) * dom.GetCoef()
	}

	return GetAcceptanceByProbability(baseCoef*men, baseCoef*menAndWomen, baseCoef*women)
}

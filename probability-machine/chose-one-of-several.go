package probabilitymachine

import "fmt"

func GetRandomFromSeveral[T string](values map[T]float64) T {
	if len(values) == 0 {
		fmt.Printf("[GetRandomFromSeveral] Values count is zero!!!\n\n\n")
	}

	var (
		valuesWithZeroCount int
		valuesWith1Count    int
	)
	for _, prob := range values {
		if prob == 0 {
			valuesWithZeroCount++
		}
		if prob == 1 {
			valuesWith1Count++
		}
	}
	if valuesWithZeroCount == len(values) {
		fmt.Printf("[GetRandomFromSeveral] All values are zero!!!\n\n\n")
	}
	if valuesWith1Count > 1 {
		fmt.Printf("[GetRandomFromSeveral] Several values are 1!!!\n\n\n")
	}

	preparedValues := cloneStringProbabilityMap(values)
	for value, prob := range preparedValues {
		preparedValues[value] = PrepareProbability(prob - 0.01)
	}

	tempValues := make(map[T]float64)

	for {
		var iterValues map[T]float64
		if len(tempValues) == 0 {
			iterValues = cloneStringProbabilityMap(preparedValues)
		} else {
			iterValues = cloneStringProbabilityMap(tempValues)
		}
		tempValues = make(map[T]float64)

		for value, prob := range iterValues {
			if GetRandomBool(prob) {
				tempValues[value] = prob
			}
		}

		if len(tempValues) == 1 {
			for value := range tempValues {
				return value
			}
		}
	}
}

func cloneStringProbabilityMap[T string](m map[T]float64) map[T]float64 {
	clonedMap := make(map[T]float64)
	for key, value := range m {
		clonedMap[key] = value
	}

	return clonedMap
}

package probability_machine

import (
	"persons_generator/core/wrapped_error"
)

func GetRandomFromSeveral[T string](values map[T]float64) (T, error) {
	var zero T
	if len(values) == 0 {
		return zero, wrapped_error.NewInternalServerError(nil, "[GetRandomFromSeveral] Values count is zero")
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
		return zero, wrapped_error.NewInternalServerError(nil, "[GetRandomFromSeveral] All values are zero")
	}
	if valuesWith1Count > 1 {
		values = PrepareMapToPickRandomValue(values)
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
			if r, _ := GetRandomBool(prob); r {
				tempValues[value] = prob
			}
		}

		if len(tempValues) == 1 {
			for value := range tempValues {
				return value, nil
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

func PrepareMapToPickRandomValue[T string](values map[T]float64) map[T]float64 {
	var totalScore float64
	for _, score := range values {
		totalScore += score
	}

	out := make(map[T]float64, len(values))
	for key, value := range values {
		out[key] = value / totalScore
	}

	return out
}

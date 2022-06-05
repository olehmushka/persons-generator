package probabilitymachine

func GetRandomFromSeveral[T string](values map[T]int) T {
	tempValues := make(map[T]int)
	for {
		var iterValues map[T]int
		if len(tempValues) == 0 {
			iterValues = cloneStringProbabilityMap(values)
		} else {
			iterValues = cloneStringProbabilityMap(tempValues)
		}
		tempValues = make(map[T]int)

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

func cloneStringProbabilityMap[T string](m map[T]int) map[T]int {
	clonedMap := make(map[T]int)
	for key, value := range m {
		clonedMap[key] = value
	}

	return clonedMap
}

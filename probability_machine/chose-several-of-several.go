package probability_machine

func GetRandomSeveralFromSeveral[T string](values map[T]float64, l int) []T {
	result := make([]T, 0, l)
	for {
		val := GetRandomFromSeveral(values)
		if includes(result, val) {
			continue
		}
		result = append(result, val)
		if len(result) == l {
			break
		}
	}

	return result
}

func includes[T string](values []T, v T) bool {
	for _, val := range values {
		if val == v {
			return true
		}
	}

	return false
}

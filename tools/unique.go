package tools

func Unique[T int | float64 | string](sl []T) []T {
	if len(sl) <= 1 {
		return sl
	}

	var (
		keys = make(map[T]bool)
		list = make([]T, 0)
	)
	for _, entry := range sl {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

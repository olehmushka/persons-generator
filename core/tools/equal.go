package tools

func IsMapEqual(maps ...map[string]float64) bool {
	if len(maps) == 0 {
		return true
	}

	for _, mm := range maps {
		for key, value := range mm {
			for _, m := range maps {
				v, ok := m[key]
				if !ok {
					return false
				}
				if v != value {
					return false
				}
			}
		}
	}

	return true
}

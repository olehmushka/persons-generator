package tools

func Chunk[T interface{}](total int, preferred []T) [][]T {
	if len(preferred) == 0 {
		return [][]T{}
	}

	if total == 0 {
		return [][]T{}
	}

	divided := make([][]T, (len(preferred)+total-1)/total)
	prev := 0
	i := 0
	till := len(preferred) - total
	for prev < till {
		next := prev + total
		divided[i] = preferred[prev:next]
		prev = next
		i++
	}
	divided[i] = preferred[prev:]

	return divided
}

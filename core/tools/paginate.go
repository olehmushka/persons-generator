package tools

func Paginate[T any](elements []T, offset int, limit int) []T {
	if offset > len(elements) {
		offset = len(elements)
	}

	end := offset + limit
	if end > len(elements) {
		end = len(elements)
	}

	return elements[offset:end]
}

package name_generator

func GetLastElement[T interface{}](sl []T) T {
	return sl[len(sl)-1]
}

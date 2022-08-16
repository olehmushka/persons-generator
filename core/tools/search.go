package tools

import "fmt"

func Search[T interface{}, Q string | int](elements []T, getFieldFn func(T) Q, q Q) T {
	var value T
	for _, el := range elements {
		field := getFieldFn(el)
		if IsString(field) && IsString(q) && ContainString(fmt.Sprint(field), fmt.Sprint(q)) {
			value = el
		}
		if IsInt(field) && IsInt(q) && field == q {
			value = el
		}
	}

	return value
}

func SearchMany[T interface{}, Q string | int](elements []T, getFieldFn func(T) Q, q Q) []T {
	values := make([]T, 0)
	for _, el := range elements {
		field := getFieldFn(el)
		if IsString(field) && IsString(q) && ContainString(fmt.Sprint(field), fmt.Sprint(q)) {
			values = append(values, el)
		}
		if IsInt(field) && IsInt(q) && field == q {
			values = append(values, el)
		}
	}

	return values
}

package tools

func MaxNumeric[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](args ...T) T {
	var out T
	for _, arg := range args {
		if arg > out {
			out = arg
		}
	}

	return out
}

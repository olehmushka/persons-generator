package tools

func GetStrLen(str string, def int) int {
	if l := len(str); l > 0 {
		return l
	}

	return def
}

func GetCharByIndex(str string, index int, def string) string {
	if len(str) <= index {
		return def
	}

	for i, char := range str {
		if i == index {
			return string(char)
		}
	}

	return def
}

func GetLastStrChar(str string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str
	}

	return string(str[len(str)-1])
}

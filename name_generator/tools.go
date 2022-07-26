package name_generator

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func getStrLenOrMin(str string, min int) int {
	l := len(str)
	if l == 0 {
		return min
	}
	return l
}

func getCharByIndex(str string, index int, def string) string {
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

func mergeChains(c1, c2 Chain) Chain {
	out := make(Chain, (len(c1)+len(c2))/2)
	for key, value := range c1 {
		out[key] = value
	}
	for key, value := range c2 {
		if v, ok := out[key]; ok {
			out[key] = mergeStrSlicesWithoutDuplicates(v, value)
		}
	}

	return out
}

func mergeStrSlicesWithoutDuplicates(ss1, ss2 []string) []string {
	out := make([]string, 0, (len(ss1)+len(ss2))/2)
	out = append(out, ss1...)
	out = append(out, ss2...)

	return removeDuplicateStr(out)
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func getLastStrChar(str string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str
	}

	return string(str[len(str)-1])
}

func hasInsideStrWordsLessThan(str string, min int) bool {
	for _, word := range strings.Split(str, " ") {
		if len(word) < min {
			return true
		}
	}
	return false
}

func makeInsideWordsCapitalized(str string) string {
	var out string
	for _, word := range strings.Split(str, " ") {
		out += capitalize(word)
	}
	return out
}

func capitalize(str string) string {
	if str == "" {
		return ""
	}

	return cases.Title(language.English).String(str)
}

package tools

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Capitalize(str string) string {
	if str == "" {
		return ""
	}

	return cases.Title(language.English).String(str)
}

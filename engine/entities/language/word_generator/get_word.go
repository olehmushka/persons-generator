package word_generator

import (
	"fmt"
	"strings"

	"persons_generator/core/tools"
	pm "persons_generator/engine/probability_machine"
)

func GetWord(base string, nameBases map[string][]string, min, max int, dupl string) string {
	if base == "" {
		panic("Please define a base")
	}

	var data Chain
	for i := 0; i < 20; i++ {
		var ok bool
		if data, ok = Chains[base]; !ok {
			UpdateChains(base, nameBases)
			continue
		}
		break
	}

	if len(data) == 0 {
		panic(fmt.Sprintf("name_base %s is incorrect! [1]", base))
	}

	val, ok := data[""]
	if !ok {
		panic(fmt.Sprintf("name_base %s is incorrect! [2]", base))
	}

	v := val
	var (
		cur = tools.RandomValueOfSlice(pm.RandFloat64, v)
		w   = ""
	)

	for i := 0; i < 20; i++ {
		if cur == "" {
			// end of word
			if len(w) < min {
				cur = ""
				w = ""
				if val, ok := data[""]; ok {
					v = val
				}
			} else {
				break
			}
		} else {
			if len(w)+len(cur) > max {
				// word too long
				if len(w) < min {
					w += cur
				}
				break
			} else {
				if val, ok := data[tools.GetLastStrChar(cur)]; ok {
					v = val
				}
			}
		}

		w += cur
		cur = tools.RandomValueOfSlice(pm.RandFloat64, v)
	}

	// parse word to get a final name
	l := tools.GetLastStrChar(w) // last letter
	if l == "'" || l == " " || l == "-" {
		w = w[0 : len(w)-1] // not allow some characters at the end
	}
	var name string
	for i, c := range w {
		var nextC, afterNextC string
		if len(w) > i+1 {
			nextC = string(w[i+1])
		}
		if len(w) > i+2 {
			afterNextC = string(w[i+2])
		}
		// duplication is not allowed
		if (nextC != "" && string(c) == nextC) && strings.Contains(dupl, string(c)) {
			continue
		}
		if len(name) == 0 {
			name += strings.ToUpper(string(c))
			continue
		}
		// remove space after hyphen
		if tools.GetLastStrChar(name) == "-" && string(c) == " " {
			continue
		}
		// capitalize letter after space
		if tools.GetLastStrChar(name) == " " {
			name += strings.ToUpper(string(c))
		}
		// capitalize letter after hyphen
		if tools.GetLastStrChar(name) == "-" {
			name += strings.ToUpper(string(c))
		}
		// "ae" => "e"
		if string(c) == "a" && (nextC != "" && nextC == "e") {
			continue
		}
		// remove three same letters in a row
		if i+2 < len(name) && (nextC != "" && string(c) == nextC) && (afterNextC != "" && string(c) == afterNextC) {
			continue
		}

		name += string(c)
	}

	// join the word if any part has only 1 letter
	if hasInsideStrWordsLessThan(name, 2) {
		name = makeInsideWordsCapitalized(name)
	}

	if len(name) < 2 {
		fmt.Printf("name is too short! Random name will be selected\n")
		name = tools.RandomValueOfSlice(pm.RandFloat64, nameBases[base])
	}

	return name
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
		out += tools.Capitalize(word)
	}
	return out
}

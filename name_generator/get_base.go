package name_generator

import (
	"fmt"
	"strings"
)

func GetBase(base int, nameBases [][]string, min, max int, dupl string) string {
	if base >= len(Chains) || Chains[base] == nil {
		UpdateChains(base, nameBases)
	}
	data := Chains[base]
	var v []string
	if val, ok := data[""]; ok {
		v = val
	}
	var (
		cur = RandomValOfSlice(v)
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
				if val, ok := data[getLastStrChar(cur)]; ok {
					v = val
				}
			}
		}

		w += cur
		cur = RandomValOfSlice(v)
	}

	// parse word to get a final name
	l := getLastStrChar(w) // last letter
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
		}
		// remove space after hyphen
		if getLastStrChar(name) == "-" && string(c) == " " {
			continue
		}
		// capitalize letter after space
		if getLastStrChar(name) == " " {
			name += strings.ToUpper(string(c))
		}
		// capitalize letter after hyphen
		if getLastStrChar(name) == "-" {
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
		fmt.Printf("name is too short! Random name will be selected")
		name = RandomValOfSlice(nameBases[base])
	}

	return name
}

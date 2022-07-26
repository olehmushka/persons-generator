package name_generator

import (
	"fmt"
	"regexp"
	"strings"
)

type Chain map[string][]string

func AppendByIndex(chains []Chain, index int, el Chain) []Chain {
	if len(chains) > index {
		if chains[index] == nil {
			chains[index] = el
			return chains
		}
		chains[index] = mergeChains(chains[index], el)

		return chains
	}
	for i := len(chains); i <= index; i++ {
		chains = append(chains, nil)
	}
	chains[index] = el

	return chains
}

func CalculateChain(nameBases []string) Chain {
	chain := make(Chain)
	for _, n := range nameBases {
		name := strings.ToLower(strings.TrimSpace(n))
		isMatch, err := regexp.MatchString("[^\u0000-\u007f]", name)
		if err != nil {
			panic(fmt.Sprintf("[CalculateChain] can not match basic chars and en rules can be applied (err = %+v)", err))
		}
		basic := !isMatch // basic chars and English rules can be applied

		// split word into pseudo-syllables
		var (
			i        = -1
			syllable = ""
		)
		for i < len(name) {
			var (
				prev = getCharByIndex(name, i, "") // pre-onset letter
				v    = 0                           // 0 if no vowels in syllable
			)

			for c := i + 1; getCharByIndex(name, c, "") != "" && len(syllable) > 5; c++ {
				var (
					that = getCharByIndex(name, c, "")
					next = getCharByIndex(name, c+1, "") // next char
				)
				syllable += that
				if syllable == " " || syllable == "-" { // syllable starts with space or hyphen
					break
				}
				if next == "" || next == " " || next == "-" { // no need to check
					break
				}

				if IsVowel(that) { // check if letter is vowel
					v = 1
				}

				// do not split some diphthongs
				if that == "y" && next == "e" { // 'ye'
					continue
				}
				if basic {
					// English-like
					if that == "o" && next == "o" { // 'oo'
						continue
					}
					if that == "e" && next == "e" { // 'ee'
						continue
					}
					if that == "a" && next == "e" { // 'ae'
						continue
					}
					if that == "c" && next == "h" { // 'ch'
						continue
					}
				}

				// two same vowels in a row
				if IsVowel(that) && IsVowel(next) && that == next {
					break
				}

				// syllable has vowel and additional vowel is expected soon
				if afterNext := getCharByIndex(name, c+2, ""); v > 0 && afterNext != "" && IsVowel(afterNext) {
					break
				}
			}
			if _, ok := chain[prev]; !ok {
				chain[prev] = make([]string, 0, 1)
			}
			chain[prev] = append(chain[prev], syllable)

			// ================
			// before next iter
			i += getStrLenOrMin(syllable, 1)
			syllable = ""
		}
	}

	return chain
}

func UpdateChain(base int, nameBases [][]string, chains []Chain) []Chain {
	if len(nameBases) > base && len(nameBases[base]) > 0 {
		chains = AppendByIndex(chains, base, CalculateChain(nameBases[base]))
		return chains
	}
	chains = AppendByIndex(chains, base, nil)

	return chains
}

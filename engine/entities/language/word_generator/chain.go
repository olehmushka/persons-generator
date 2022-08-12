package word_generator

import (
	"fmt"
	"regexp"
	"strings"

	"persons_generator/core/tools"
)

type Chain map[string][]string

func AppendByBase(chains map[string]Chain, base string, el Chain) map[string]Chain {
	if len(chains) == 0 {
		chains = make(map[string]Chain)
	}

	chain, ok := chains[base]
	if !ok {
		chains[base] = el
		return chains
	}
	chains[base] = mergeChains(chain, el)
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
				prev = tools.GetCharByIndex(name, i, "") // pre-onset letter
				v    = 0                                 // 0 if no vowels in syllable
			)

			for c := i + 1; tools.GetCharByIndex(name, c, "") != "" && len(syllable) < 5; c++ {
				var (
					that = tools.GetCharByIndex(name, c, "")
					next = tools.GetCharByIndex(name, c+1, "") // next char
				)
				syllable += that
				if syllable == " " || syllable == "-" { // syllable starts with space or hyphen
					break
				}
				if next == "" || next == " " || next == "-" { // no need to check
					break
				}

				if tools.IsVowel(that) { // check if letter is vowel
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
				if tools.IsVowel(that) && tools.IsVowel(next) && that == next {
					break
				}

				// syllable has vowel and additional vowel is expected soon
				if afterNext := tools.GetCharByIndex(name, c+2, ""); v > 0 && afterNext != "" && tools.IsVowel(afterNext) {
					break
				}
			}
			if _, ok := chain[prev]; !ok {
				chain[prev] = make([]string, 0, 1)
			}
			chain[prev] = append(chain[prev], syllable)

			// ================
			// before next iter
			i += tools.GetStrLen(syllable, 1)
			syllable = ""
		}
	}

	return chain
}

func UpdateChain(base string, nameBases map[string][]string, chains map[string]Chain) map[string]Chain {
	if chain, ok := nameBases[base]; ok {
		if len(chain) == 0 {
			panic(fmt.Sprintf("word base is empty (base=%s)", base))
		}
		chains = AppendByBase(chains, base, CalculateChain(chain))
		return chains
	}

	chains = AppendByBase(chains, base, nil)

	return chains
}

func mergeChains(c1, c2 Chain) Chain {
	out := make(Chain, (len(c1)+len(c2))/2)
	for key, value := range c1 {
		out[key] = value
	}
	for key, value := range c2 {
		if v, ok := out[key]; ok {
			out[key] = tools.MergeUnique(v, value)
		}
	}

	return out
}

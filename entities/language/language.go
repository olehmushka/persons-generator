package language

import (
	"fmt"

	wg "persons_generator/entities/language/word_generator"
	pm "persons_generator/probability_machine"
	"persons_generator/tools"
)

type Language struct {
	Name      string
	Subfamily *Subfamily
	WordBase  *WordBase
	IsLiving  bool
}

func New(preferred []string) *Language {
	if len(preferred) == 0 {
		return tools.RandomValueOfSlice(pm.RandFloat64, AllLanguages)
	}
	langs := make([]*Language, 0, len(preferred))
	for _, pref := range preferred {
		if lang := GetLanguageByName(pref); lang != nil {
			langs = append(langs, lang)
		}
	}
	switch len(langs) {
	case 0:
		return tools.RandomValueOfSlice(pm.RandFloat64, AllLanguages)
	case 1:
		return langs[0]
	default:
		return tools.RandomValueOfSlice(pm.RandFloat64, langs)
	}
}

func (l *Language) GetWord() string {
	if l.WordBase == nil {
		panic(fmt.Sprintf("can not get word base (language=%s)", l.Name))
	}
	return wg.GetWord(l.WordBase.Name, ExtractWords(AllWordBases), l.WordBase.Min, l.WordBase.Max, l.WordBase.Dupl)
}

func GetLanguageByName(name string) *Language {
	return tools.Search(AllLanguages, func(el *Language) string { return el.Name }, name)
}

func (l *Language) GetCultureName() string {
	return wg.GetCultureName(l.GetWord())
}

func (l *Language) GetReligionName() string {
	return wg.GetReligionName(l.GetWord())
}

func IsLanguagesEqual(l1, l2 *Language) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil {
		return false
	}
	if l2 == nil {
		return false
	}

	return l1.Name == l2.Name
}

func GetLanguagesBySubfamilyName(name string) []*Language {
	if name == "" {
		return []*Language{}
	}

	out := make([]*Language, 0)
	sfs := GetSubfamiliesWithChildrenByName(name)
	if len(sfs) == 0 {
		return []*Language{}
	}

	for _, lang := range AllLanguages {
		for _, sf := range sfs {
			if IsSubfamiliesEqual(sf, lang.Subfamily) {
				out = append(out, lang)
			}
		}
	}

	return out
}

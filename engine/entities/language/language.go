package language

import (
	"fmt"

	"persons_generator/core/tools"
	"persons_generator/core/wrapped_error"
	g "persons_generator/engine/entities/gender"
	wg "persons_generator/engine/entities/language/word_generator"
	pm "persons_generator/engine/probability_machine"
)

type Language struct {
	ID          string       `json:"id" bson:"id"`
	Name        string       `json:"name" bson:"name"`
	Subfamily   *Subfamily   `json:"subfamily" bson:"subfamily,omitempty"`
	WordBaseRef *WordBaseRef `json:"word_base_ref" bson:"word_base_ref,omitempty"`
	WordBase    *WordBase    `json:"word_base" bson:"word_base,omitempty"`
	IsLiving    bool         `json:"is_living" bson:"is_living"`
	Origin      Origin       `json:"origin" bson:"origin"`
}

func New(preferred []string) (*Language, error) {
	if len(preferred) == 0 {
		return tools.RandomValueOfSlice(pm.RandFloat64, AllLanguages)
	}
	langs := make([]*Language, 0, len(preferred))
	for _, pref := range preferred {
		if lang := GetLanguageByName(pref); lang != nil {
			langs = append(langs, lang...)
		}
	}

	switch len(langs) {
	case 0:
		return tools.RandomValueOfSlice(pm.RandFloat64, AllLanguages)
	case 1:
		return langs[0], nil
	default:
		return tools.RandomValueOfSlice(pm.RandFloat64, langs)
	}
}

func (l *Language) getWordBase() (*WordBase, error) {
	if l.WordBase != nil {
		return l.WordBase, nil
	}
	wb, err := l.WordBaseRef.LoadWordBase()
	if err != nil {
		return nil, err
	}
	l.WordBase = wb

	return wb, nil
}

func (l *Language) GetWord() (string, error) {
	wb, err := l.getWordBase()
	if err != nil {
		return "", err
	}
	if wb == nil {
		return "", wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not get word base (language=%s)", l.Name))
	}

	return wg.GetWord(wb.Name, ExtractWords(AllWordBases), wb.Min, wb.Max, wb.Dupl)
}

func (l *Language) IsZero() bool {
	return l == nil
}

func (l *Language) Print() {
	if l == nil {
		return
	}
	fmt.Printf("language: %s (is_living: %t)\n", l.Name, l.IsLiving)
	l.Subfamily.Print()
}

func GetLanguageByName(name string) []*Language {
	if name == "" {
		return nil
	}

	return tools.SearchMany(AllLanguages, func(l *Language) string { return l.Name }, name)
}

func FindLanguage(id, name string) *Language {
	for _, lang := range AllLanguages {
		isIDMatch := true
		if id != "" && lang.ID != id {
			isIDMatch = false
		}

		isNameMatch := true
		if name != "" && lang.Name != name {
			isNameMatch = false
		}
		if isIDMatch && isNameMatch {
			return lang
		}
	}

	return nil
}

func (l *Language) GetCultureName() (string, error) {
	w, err := l.GetWord()
	if err != nil {
		return "", err
	}

	return wg.GetCultureName(w), nil
}

func (l *Language) GetReligionName() (string, error) {
	w, err := l.GetWord()
	if err != nil {
		return "", err
	}

	return wg.GetReligionName(w), nil
}

func (l *Language) GetOwnName(sex g.Sex) (string, error) {
	if sex == g.MaleSex {
		return tools.RandomValueOfSlice(pm.RandFloat64, l.WordBase.MaleOwnNames)
	}
	if sex == g.FemaleSex {
		return tools.RandomValueOfSlice(pm.RandFloat64, l.WordBase.FemaleOwnNames)
	}

	return "", nil
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

func GetLanguageSubfamilyChains(accum []string, sf *Subfamily) []string {
	if sf == nil {
		return accum
	}
	if sf.ExtendedSubfamily == nil {
		return accum
	}

	return GetLanguageSubfamilyChains(
		append(accum, sf.Name),
		sf.ExtendedSubfamily,
	)
}

func GetFullLanguageChains(lang *Language) []string {
	if lang == nil {
		return []string{}
	}
	out := []string{lang.Name}
	out = GetLanguageSubfamilyChains(out, lang.Subfamily)

	return append(out, string(lang.Subfamily.Family))
}

func GetLanguageKinship(l1, l2 *Language) int {
	var (
		l1Chains = GetFullLanguageChains(l1)
		l2Chains = GetFullLanguageChains(l2)
	)
	if len(l1Chains) == 0 || len(l2Chains) == 0 {
		return -1
	}
	if l1Chains[0] == l2Chains[0] {
		return 0
	}
	if l1Chains[len(l1Chains)-1] != l2Chains[len(l2Chains)-1] {
		return -1
	}

	maxIter := len(l1Chains)
	if maxIter < len(l2Chains) {
		maxIter = len(l2Chains)
	}
	for i := 0; i < maxIter; i++ {
		var (
			l1El = l1Chains[len(l1Chains)-1]
			l2El = l2Chains[len(l2Chains)-1]
		)
		if len(l1Chains) > i {
			l1El = l1Chains[i]
		}
		if len(l2Chains) > i {
			l2El = l2Chains[i]
		}
		if l1El == l2El {
			return i
		}
	}

	return -1
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

func GetLanguageSimilarityCoef(l1, l2 *Language) (float64, error) {
	if l1.IsZero() && l2.IsZero() {
		return 1, nil
	}
	if l1.IsZero() || l2.IsZero() {
		return 0, wrapped_error.NewInternalServerError(nil, "can not compare languages if one of it is <nil>")
	}

	switch kinship := GetLanguageKinship(l1, l2); kinship {
	case -1:
		return 0, nil
	case 0:
		return 1, nil
	case 1:
		return 0.75, nil
	default:
		return 1 / float64(kinship), nil
	}
}

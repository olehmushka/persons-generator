package culture

import (
	"errors"
	"fmt"

	"persons_generator/core/tools"
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/language"
	pm "persons_generator/engine/probability_machine"
)

type Culture struct {
	Proto []*Culture
	Name  string

	Abstuct         *AbstructCulture
	Root            *Root
	Language        *language.Language
	Ethos           *Ethos
	Traditions      []*Tradition
	GenderDominance *g.Dominance
	MartialCustom   g.Acceptance
}

func New(preferred *Preference) (*Culture, error) {
	proto, err := getProtoCultures(preferred)
	if err != nil {
		return nil, err
	}
	if len(proto) == 0 {
		return nil, errors.New("proto cultures can not be zero")
	}

	return NewWithProto(proto)
}

func getProtoCultures(preferred *Preference) ([]*Culture, error) {
	if preferred == nil {
		return GetRandomProtoCultures(1, 3)
	}

	switch preferred.Kind {
	case StrictPrefKind:
		return getProtoCulturesForStrictKind(preferred)
	case HybridPrefKind:
		return getProtoCulturesForHybridKind(preferred)
	default:
		return nil, fmt.Errorf("unexpected preference kind (kind_name=%s)", preferred.Kind)
	}
}

func getProtoCulturesForStrictKind(preferred *Preference) ([]*Culture, error) {
	if preferred == nil {
		return nil, errors.New("preferred can not be nil")
	}
	if preferred.Kind == StrictPrefKind && len(preferred.Names) != preferred.Amount {
		return nil, fmt.Errorf("for strict kind of preference number preferred names (%d) can not be not equal to amount (%d)", len(preferred.Names), preferred.Amount)
	}

	out, err := GetProtoCulturesByPreferred(preferred.Names)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func getProtoCulturesForHybridKind(preferred *Preference) ([]*Culture, error) {
	if preferred == nil {
		return nil, errors.New("preferred can not be nil")
	}

	out, err := GetProtoCulturesByPreferred(preferred.Names)
	if err != nil {
		return nil, err
	}
	if len(out) == preferred.Amount {
		return out, nil
	}
	for i := 0; i < 20; i++ {
		out = append(out, tools.RandomValuesOfSlice(pm.RandFloat64, AllCultures, preferred.Amount-len(out))...)
		out = UniqueCultures(out)
		if len(out) == preferred.Amount {
			break
		}
	}

	return out, nil
}

func NewWithProto(proto []*Culture) (*Culture, error) {
	c := &Culture{Proto: proto}
	c.GenderDominance = getGenderDominance(c.Proto)
	c.MartialCustom = getMartialCustom(c.Proto)
	c.Language = language.New(getLanguageNamesFromProto(c.Proto))
	name, err := c.Language.GetCultureName()
	if err != nil {
		return nil, err
	}
	c.Name = name
	c.Ethos = getEthos(c.Proto)
	c.Traditions = getTraditions(c.Proto)
	c.Root = getRoot(c.Proto)

	return c, nil
}

func NewCultures(amount int, preferred []*Preference) ([]*Culture, error) {
	cultures := make([]*Culture, amount)
	if amount < len(preferred) {
		return nil, fmt.Errorf("amount (%d) can not be less than length of preferred (length=%d)", amount, len(preferred))
	}
	for i := range cultures {
		var p *Preference
		if len(preferred) > i {
			p = preferred[i]
		}
		c, err := New(p)
		if err != nil {
			return nil, err
		}
		cultures[i] = c
	}

	return cultures, nil
}

func (c *Culture) Print() {
	if c == nil {
		return
	}

	fmt.Printf("Culture (name=%s)\n", c.Name)
	c.Root.Print()
	if len(c.Proto) > 0 {
		fmt.Printf("Proto cultures (culture_name=%s)\n", c.Name)
		for _, p := range c.Proto {
			if p == nil {
				continue
			}
			fmt.Printf(" - %s\n", p.Name)
		}
	}
	c.Language.Print()
	c.Ethos.Print()
	c.PrintTraditions()
	c.GenderDominance.Print()
	c.PrintMartialCustom()
}

func getLanguageNamesFromProto(proto []*Culture) []string {
	names := make([]string, 0, len(proto))
	for _, p := range proto {
		if p == nil {
			continue
		}
		if p.Language == nil {
			continue
		}
		names = append(names, p.Language.Name)
	}

	return tools.Unique(names)
}

func GetProtoCulturesByPreferred(names []string) ([]*Culture, error) {
	out := make([]*Culture, len(names))
	for i := range out {
		found := GetProtoCultureByPreferredName(names[i])
		if found == nil {
			return nil, fmt.Errorf("can not found proto culture by name (name=%s)", names[i])
		}
		out[i] = found
	}

	return out, nil
}

func GetProtoCultureByPreferredName(prefName string) *Culture {
	for _, c := range AllCultures {
		if tools.ContainString(c.Name, prefName) {
			return c
		}
	}

	return nil
}

func UniqueCultures(cultures []*Culture) []*Culture {
	if len(cultures) <= 1 {
		return cultures
	}

	preOut := make(map[string]*Culture)
	for _, c := range cultures {
		preOut[c.Name] = c
	}

	out := make([]*Culture, 0, len(preOut))
	for _, c := range preOut {
		out = append(out, c)
	}

	return out
}

func GetRandomProtoCultures(min, max int) ([]*Culture, error) {
	if max > len(AllCultures) {
		return nil, fmt.Errorf("number of all cultures can not be less than max for random generation (%d)", max)
	}
	var (
		amount = pm.RandIntInRange(min, max)
		out    = make([]*Culture, 0, amount)
	)
	for i := 0; i < 20; i++ {
		out = append(out, tools.RandomValuesOfSlice(pm.RandFloat64, AllCultures, amount-len(out))...)
		out = UniqueCultures(out)
		if len(out) == amount {
			break
		}
	}

	return out, nil
}

func GetCultureByName(name string, list []*Culture) *Culture {
	if name == "" || len(list) == 0 {
		return nil
	}

	return tools.Search(list, func(c *Culture) string { return c.Name }, name)
}

func GetCulturesByName(name string, list []*Culture) []*Culture {
	if name == "" || len(list) == 0 {
		return nil
	}

	return tools.SearchMany(list, func(c *Culture) string { return c.Name }, name)
}

func MapCultureNames(cultures []*Culture) []string {
	if len(cultures) == 0 {
		return []string{}
	}

	out := make([]string, len(cultures))
	for i := range out {
		out[i] = cultures[i].Name
	}

	return out
}

package culture

import (
	"fmt"

	g "persons_generator/entities/gender"
	"persons_generator/entities/language"
	pm "persons_generator/probability_machine"
	"persons_generator/tools"
)

type Culture struct {
	Proto []*Culture
	Name  string

	Root            *Root
	Language        *language.Language
	Ethos           *Ethos
	Traditions      []*Tradition
	GenderDominance *g.Dominance
	MartialCustom   g.Acceptance
}

func New(preferred []string) *Culture {
	c := &Culture{}
	c.Proto = getProtoCultures(preferred)
	c.GenderDominance = getGenderDominance(c.Proto)
	c.MartialCustom = getMartialCustom(c.Proto)
	c.Language = language.New(getLanguageNamesFromProto(c.Proto))
	c.Name = c.Language.GetCultureName()
	c.Ethos = getEthos(c.Proto)
	c.Traditions = getTraditions(c.Proto)

	return c
}

func NewCultures(amount int, preferred []string) []*Culture {
	cultures := make([]*Culture, amount)
	chunk := chunkPreferred(amount, preferred)
	for i := range cultures {
		cultures[i] = New(chunk[i])
	}

	return cultures
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

func chunkPreferred(amount int, preferred []string) [][]string {
	if amount == 0 {
		return [][]string{}
	}

	if len(preferred) == 0 {
		out := make([][]string, amount)
		for i := range out {
			out[i] = []string{}
		}
		return out
	}

	out := make([][]string, amount)
	if amount >= len(preferred) {
		for i := range out {
			if len(preferred) > i {
				out[i] = []string{preferred[i]}
				continue
			}
			out[i] = []string{}
		}
		return out
	}

	return tools.Chunk(amount, preferred)
}

func getProtoCultures(preferred []string) []*Culture {
	expectedAmount := pm.RandIntInRange(1, 2)
	if found := GetProtoCulturesByPreferred(preferred); len(found) > 0 {
		return tools.RandomValuesOfSlice(pm.RandFloat64, found, expectedAmount)
	}

	return tools.RandomValuesOfSlice(pm.RandFloat64, AllCultures, expectedAmount)
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

func GetProtoCulturesByPreferred(preferred []string) []*Culture {
	if len(preferred) == 0 {
		return []*Culture{}
	}

	out := make([]*Culture, 0, len(preferred))
	for _, pref := range preferred {
		if found := GetProtoCultureByPreferred(pref); found != nil {
			out = append(out, found)
		}
	}

	return out
}

func GetProtoCultureByPreferred(pref string) *Culture {
	for _, c := range AllCultures {
		if tools.ContainString(c.Name, pref) {
			return c
		}
	}

	return nil
}
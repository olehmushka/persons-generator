package culture

import (
	g "persons_generator/entities/gender"
	"persons_generator/entities/language"
	"persons_generator/tools"
)

// bc "persons_generator/entities/base_culture"
// pc "persons_generator/entities/proto_culture"
// pm "persons_generator/probability_machine"
// "persons_generator/tools"

type Culture struct {
	Proto []*Culture
	Name  string

	Root            *Root
	Language        *language.Language
	Ethos           *Ethos
	Traditions      []*Tradition
	GenderDominance *g.Dominance
	MartialCustom   g.Acceptance
	// Proto                 *pc.ProtoCulture
	// InheritedBaseCultures []*bc.BaseCulture
}

func New(preferred []string) *Culture {
	c := &Culture{}
	c.Proto = GetProtoCultures(preferred)
	c.GenderDominance = g.NewDominance()
	c.MartialCustom = g.GetMartialCustom(1, c.GenderDominance)
	c.Language = language.New(GetLanguageNamesFromProto(c.Proto))
	c.Name = c.Language.GetCultureName()

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

func chunkPreferred(amount int, preferred []string) [][]string {
	out := make([][]string, amount)
	if amount <= len(preferred) {
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

func GetProtoCultures(preferred []string) []*Culture {
	return nil
}

func GetLanguageNamesFromProto(proto []*Culture) []string {
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

// func NewMany(amount int, preferred []string) []*Culture {
// 	out := make([]*Culture, amount)
// 	chunks := tools.Chunk(amount, preferred)
// 	for i := range out {
// 		p := Params{
// 			MinInheritedBaseCultures: 1,
// 			MaxInheritedBaseCultures: 3,
// 		}
// 		if len(chunks) > i {
// 			p.PreferredCultures = chunks[i]
// 		}
// 		out[i] = New(p)
// 	}

// 	return out
// }

// func New(params Params) *Culture {
// 	if len(params.PreferredCultures) == 0 {
// 		proto := pc.GetRandomProtoCulture()
// 		return &Culture{
// 			Proto:                 proto,
// 			InheritedBaseCultures: bc.GetRandomBaseCultures(bc.GetBaseCulturesByProto(proto), params.MinInheritedBaseCultures, params.MaxInheritedBaseCultures),
// 		}
// 	}

// 	if protoCultures := pc.GetProtoCulturesByPref(params.PreferredCultures); len(protoCultures) > 0 {
// 		proto := tools.RandomValueOfSlice(pm.RandFloat64, protoCultures)
// 		return &Culture{
// 			Proto:                 proto,
// 			InheritedBaseCultures: bc.GetRandomBaseCultures(bc.GetBaseCulturesByProto(proto), params.MinInheritedBaseCultures, params.MaxInheritedBaseCultures),
// 		}
// 	}

// 	if baseCultures := bc.GetBaseCulturesByPref(params.PreferredCultures); len(baseCultures) > 0 {
// 		baseCulture := tools.RandomValueOfSlice(pm.RandFloat64, baseCultures)
// 		return &Culture{
// 			Proto:                 baseCulture.Proto,
// 			InheritedBaseCultures: []*bc.BaseCulture{baseCulture},
// 		}
// 	}

// 	proto := pc.GetRandomProtoCulture()

// 	return &Culture{
// 		Proto:                 proto,
// 		InheritedBaseCultures: bc.GetRandomBaseCultures(bc.GetBaseCulturesByProto(proto), params.MinInheritedBaseCultures, params.MaxInheritedBaseCultures),
// 	}
// }

// func (c *Culture) GetCultureName() string {
// 	if c == nil {
// 		return ""
// 	}

// 	return c.Name
// }

// func (c *Culture) HasLastName() bool {
// 	if c == nil {
// 		return false
// 	}
// 	for _, cl := range bc.CulturesWithoutLastName {
// 		if bc.IsEqualCulture(cl, c) {
// 			return false
// 		}
// 	}

// 	return true
// }

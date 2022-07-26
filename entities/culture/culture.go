package culture

import (
	bc "persons_generator/entities/base_culture"
	pc "persons_generator/entities/proto_culture"
	pm "persons_generator/probability_machine"
	"persons_generator/tools"
)

type Culture struct {
	Name string

	Proto                 *pc.ProtoCulture
	InheritedBaseCultures []*bc.BaseCulture
}

func NewMany(amount int, preferred []string) []*Culture {
	out := make([]*Culture, amount)
	chunks := tools.Chunk(amount, preferred)
	for i := range out {
		p := Params{
			MinInheritedBaseCultures: 1,
			MaxInheritedBaseCultures: 3,
		}
		if len(chunks) > i {
			p.PreferredCultures = chunks[i]
		}
		out[i] = New(p)
	}

	return out
}

func New(params Params) *Culture {
	if len(params.PreferredCultures) == 0 {
		proto := pc.GetRandomProtoCulture()
		return &Culture{
			Proto:                 proto,
			InheritedBaseCultures: bc.GetRandomBaseCultures(bc.GetBaseCulturesByProto(proto), params.MinInheritedBaseCultures, params.MaxInheritedBaseCultures),
		}
	}

	if protoCultures := pc.GetProtoCulturesByPref(params.PreferredCultures); len(protoCultures) > 0 {
		proto := tools.RandomValueOfSlice(pm.RandFloat64, protoCultures)
		return &Culture{
			Proto:                 proto,
			InheritedBaseCultures: bc.GetRandomBaseCultures(bc.GetBaseCulturesByProto(proto), params.MinInheritedBaseCultures, params.MaxInheritedBaseCultures),
		}
	}

	if baseCultures := bc.GetBaseCulturesByPref(params.PreferredCultures); len(baseCultures) > 0 {
		baseCulture := tools.RandomValueOfSlice(pm.RandFloat64, baseCultures)
		return &Culture{
			Proto:                 baseCulture.Proto,
			InheritedBaseCultures: []*bc.BaseCulture{baseCulture},
		}
	}

	proto := pc.GetRandomProtoCulture()

	return &Culture{
		Proto:                 proto,
		InheritedBaseCultures: bc.GetRandomBaseCultures(bc.GetBaseCulturesByProto(proto), params.MinInheritedBaseCultures, params.MaxInheritedBaseCultures),
	}
}

func (c *Culture) GetCultureName() string {
	if c == nil {
		return ""
	}

	return c.Name
}

func (c *Culture) HasLastName() bool {
	if c == nil {
		return false
	}
	for _, cl := range bc.CulturesWithoutLastName {
		if bc.IsEqualCulture(cl, c) {
			return false
		}
	}

	return true
}

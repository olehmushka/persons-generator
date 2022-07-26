package culture

import (
	bc "persons_generator/entities/base_culture"
	pc "persons_generator/entities/proto_culture"
)

type CultureInterface interface {
	GetCultureName() string
	HasLastName() bool
}

type Culture struct {
	Name string

	Proto                 *pc.ProtoCulture
	InheritedBaseCultures []*bc.BaseCulture
}

func New(params Params) *Culture {
	proto := pc.GetRandomProtoCulture()
	ibcs := bc.GetRandomBaseCultures(bc.GetBaseCulturesByProto(proto), params.MinInheritedBaseCultures, params.MaxInheritedBaseCultures)
	return &Culture{
		Proto:                 proto,
		InheritedBaseCultures: ibcs,
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
		if IsEqualCulture(cl, c) {
			return false
		}
	}

	return true
}

func IsEqualCulture(c1, c2 CultureInterface) bool {
	if c1 == nil && c2 == nil {
		return true
	}
	if c1 == nil {
		return false
	}
	if c2 == nil {
		return false
	}

	return c1.GetCultureName() == c2.GetCultureName()
}

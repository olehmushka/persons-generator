package language

import (
	"fmt"

	"persons_generator/core/tools"
)

type Subfamily struct {
	Name              string     `json:"name"`
	Family            Family     `json:"family"`
	ExtendedSubfamily *Subfamily `json:"extended_subfamily"`
	IsLiving          bool       `json:"is_living"`
}

func (sf *Subfamily) Print() {
	if sf == nil {
		return
	}
	fmt.Printf("subfamily: %s (family: %s, is_living: %t)\n", sf.Name, sf.Family, sf.IsLiving)
	if sf.ExtendedSubfamily != nil {
		sf.ExtendedSubfamily.Print()
	}
}

func GetSubfamiliesByName(name string) []*Subfamily {
	if name == "" {
		return []*Subfamily{}
	}

	out := make([]*Subfamily, 0)
	for _, sf := range AllSubfamilies {
		if tools.ContainString(sf.Name, name) {
			out = append(out, sf)
		}
	}

	return out
}

func GetSubfamilyChildren(subfamily *Subfamily) []*Subfamily {
	if subfamily == nil {
		return []*Subfamily{}
	}

	out := make([]*Subfamily, 0)
	for _, sf := range AllSubfamilies {
		if IsSubfamiliesEqual(sf.ExtendedSubfamily, subfamily) {
			out = append(out, sf)
			if children := GetSubfamilyChildren(sf); len(children) > 0 {
				out = append(out, children...)
			}
		}
	}

	return out
}

func GetSubfamiliesWithChildrenByName(name string) []*Subfamily {
	if name == "" {
		return []*Subfamily{}
	}

	out := make([]*Subfamily, 0)
	for _, sf := range AllSubfamilies {
		if tools.ContainString(sf.Name, name) {
			out = append(out, sf)
		}
	}
	if len(out) == 0 {
		return out
	}

	for _, sf := range out {
		if children := GetSubfamilyChildren(sf); len(children) > 0 {
			out = append(out, children...)
		}
	}

	return out
}

func IsSubfamiliesEqual(l1, l2 *Subfamily) bool {
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

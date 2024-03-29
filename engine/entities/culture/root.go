package culture

import (
	"fmt"

	"persons_generator/core/tools"
	pm "persons_generator/engine/probability_machine"
)

type Root struct {
	Name string `json:"name" bson:"name"`
}

func (r *Root) SerializeName() string {
	if r == nil {
		return ""
	}

	return r.Name
}

func (r *Root) Print() {
	if r == nil {
		return
	}
	fmt.Printf("Root: %s\n", r.Name)
}

func getRoot(proto []*Culture) (*Root, error) {
	rs := make([]*Root, 0, len(proto))
	for _, p := range proto {
		rs = append(rs, p.Root)
	}

	unique := UniqueRoot(rs)
	if len(unique) == 1 {
		return unique[0], nil
	}

	return tools.RandomValueOfSlice(pm.RandFloat64, unique)
}

func UniqueRoot(rs []*Root) []*Root {
	if len(rs) <= 1 {
		return rs
	}

	preOut := make(map[string]*Root)
	for _, t := range rs {
		preOut[t.Name] = t
	}

	out := make([]*Root, 0, len(preOut))
	for _, t := range preOut {
		out = append(out, t)
	}

	return out
}

func IsRootEqual(r1, r2 *Root) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}

	return r1.Name == r2.Name
}

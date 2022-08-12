package culture

import (
	"fmt"

	pm "persons_generator/engine/probability_machine"
)

type Ethos struct {
	Name string

	IsDiplomatic     bool
	IsWarlike        bool
	IsAdministrative bool
	IsIntrigue       bool
	IsScholarly      bool
}

func (e *Ethos) Print() {
	if e == nil {
		return
	}
	fmt.Printf("Ethos: %s\n", e.Name)
}

func getEthos(proto []*Culture) *Ethos {
	ethoses := UniqueEthoses(ExtractEthoses(proto))
	m := make(map[string]int)
	for _, c := range proto {
		if _, ok := m[c.Ethos.Name]; ok {
			m[c.Ethos.Name]++
			continue
		}
		m[c.Ethos.Name] = 1
	}

	mWithProb := make(map[string]float64, len(m))
	for name, amount := range m {
		mWithProb[name] = float64(amount) / float64(len(proto))
	}

	chosenEthos := pm.GetRandomFromSeveral(mWithProb)
	for _, e := range ethoses {
		if e.Name == chosenEthos {
			return e
		}
	}

	return nil
}

func ExtractEthoses(cultures []*Culture) []*Ethos {
	if len(cultures) == 0 {
		return []*Ethos{}
	}

	out := make([]*Ethos, 0, len(cultures))
	for _, c := range cultures {
		if c == nil {
			continue
		}
		out = append(out, c.Ethos)
	}

	return out
}

func UniqueEthoses(ethoses []*Ethos) []*Ethos {
	if len(ethoses) <= 1 {
		return ethoses
	}

	preOut := make(map[string]*Ethos)
	for _, e := range ethoses {
		preOut[e.Name] = e
	}

	out := make([]*Ethos, 0, len(preOut))
	for _, e := range preOut {
		out = append(out, e)
	}

	return out
}

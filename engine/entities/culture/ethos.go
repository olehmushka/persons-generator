package culture

import (
	"fmt"

	"persons_generator/core/wrapped_error"
	pm "persons_generator/engine/probability_machine"
)

type Ethos struct {
	Name string `json:"name" bson:"name"`

	IsDiplomatic     bool `json:"is_diplomatic" bson:"is_diplomatic"`
	IsWarlike        bool `json:"is_warlike" bson:"is_warlike"`
	IsAdministrative bool `json:"is_administrative" bson:"is_administrative"`
	IsIntrigue       bool `json:"is_intrigue" bson:"is_intrigue"`
	IsScholarly      bool `json:"is_scholarly" bson:"is_scholarly"`
}

func (e *Ethos) IsZero() bool {
	return e == nil
}

func (e *Ethos) Serialize() string {
	if e == nil {
		return ""
	}

	return e.Name
}

func (e *Ethos) Print() {
	if e == nil {
		return
	}
	fmt.Printf("Ethos: %s\n", e.Name)
}

func getRandomEthosFromCultures(proto []*Culture) (*Ethos, error) {
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

	chosenEthos, err := pm.GetRandomFromSeveral(mWithProb)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate ethos")
	}
	for _, e := range ethoses {
		if e.Name == chosenEthos {
			return e, nil
		}
	}

	return nil, nil
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

func GetEthosSimilarityCoef(e1, e2 *Ethos) float64 {
	if e1.IsZero() && e2.IsZero() {
		return 1
	}
	if e1.IsZero() || e2.IsZero() {
		return 0
	}

	if e1.Name == e2.Name {
		return 1
	}

	similarityTraits := []struct {
		enable bool
		coef   float64
	}{
		{
			enable: e1.IsDiplomatic == e2.IsDiplomatic,
			coef:   0.2,
		},
		{
			enable: e1.IsWarlike == e2.IsWarlike,
			coef:   0.2,
		},
		{
			enable: e1.IsAdministrative == e2.IsAdministrative,
			coef:   0.2,
		},
		{
			enable: e1.IsIntrigue == e2.IsIntrigue,
			coef:   0.2,
		},
		{
			enable: e1.IsScholarly == e2.IsScholarly,
			coef:   0.2,
		},
	}

	var out float64
	for _, t := range similarityTraits {
		if t.enable {
			out += t.coef
		}
	}

	return out
}

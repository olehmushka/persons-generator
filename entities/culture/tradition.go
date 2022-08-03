package culture

import (
	"fmt"

	pm "persons_generator/probability_machine"
	"persons_generator/tools"
)

type Tradition struct {
	Name              string
	PreferredEthoses  []*Ethos
	Type              TraditionType
	CompatibilityFunc func(c *Culture) bool
}

type TraditionType string

const (
	RealmTraditionType    TraditionType = "realm_tradition_type"
	WarfareTraditionType  TraditionType = "warfare_tradition_type"
	SocialTraditionType   TraditionType = "social_tradition_type"
	RitualTraditionType   TraditionType = "ritual_tradition_type"
	RegionalTraditionType TraditionType = "regional_tradition_type"
)

func (c *Culture) PrintTraditions() {
	if c == nil {
		return
	}
	fmt.Printf("Traditions: (culture_name=%s)\n", c.Name)
	for _, t := range c.Traditions {
		fmt.Printf(" - %s (type: %s)\n", t.Name, t.Type)
	}
}

func getTraditions(proto []*Culture) []*Tradition {
	ts := make([]*Tradition, 0, len(proto))
	var traditionsAmount int
	for _, p := range proto {
		ts = append(ts, p.Traditions...)
		traditionsAmount += len(p.Traditions)
	}

	return tools.RandomValuesOfSlice(
		pm.RandFloat64,
		UniqueTraditions(ts),
		traditionsAmount/len(proto),
	)
}

func UniqueTraditions(ts []*Tradition) []*Tradition {
	if len(ts) <= 1 {
		return ts
	}

	preOut := make(map[string]*Tradition)
	for _, t := range ts {
		preOut[t.Name] = t
	}

	out := make([]*Tradition, 0, len(preOut))
	for _, t := range preOut {
		out = append(out, t)
	}

	return out
}

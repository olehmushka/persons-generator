package culture

import (
	"fmt"
	"persons_generator/core/tools"
	pm "persons_generator/engine/probability_machine"
)

type Tradition struct {
	Name              string                `json:"name" bson:"name"`
	PreferredEthoses  []*Ethos              `json:"preferred_ethoses" bson:"preferred_ethoses"`
	Type              TraditionType         `json:"type" bson:"type"`
	CompatibilityFunc func(c *Culture) bool `json:"-" bson:"-"`
}

type TraditionType string

const (
	RealmTraditionType    TraditionType = "realm_tradition_type"
	WarfareTraditionType  TraditionType = "warfare_tradition_type"
	SocialTraditionType   TraditionType = "social_tradition_type"
	RitualTraditionType   TraditionType = "ritual_tradition_type"
	RegionalTraditionType TraditionType = "regional_tradition_type"
)

func (t *Tradition) IsZero() bool {
	return t == nil
}

func (t *Tradition) String() string {
	if t == nil {
		return ""
	}

	return t.Name
}

func (c *Culture) PrintTraditions() {
	if c == nil {
		return
	}
	fmt.Printf("Traditions: (culture_name=%s)\n", c.Name)
	for _, t := range c.Traditions {
		fmt.Printf(" - %s (type: %s)\n", t.Name, t.Type)
	}
}

func getTraditions(proto []*Culture, ethos *Ethos) ([]*Tradition, error) {
	out := make([]*Tradition, 0, len(proto))
	allTraditions := make([]*Tradition, 0, len(proto))
	for _, p := range proto {
		for _, t := range p.Traditions {
			for _, e := range t.PreferredEthoses {
				if e.Name == ethos.Name {
					out = append(out, t)
					break
				}
			}
			allTraditions = append(allTraditions, t)
		}
	}
	randomTraditions, err := tools.RandomValuesOfSlice(
		pm.RandFloat64,
		UniqueTraditions(allTraditions),
		len(allTraditions)/len(proto),
	)
	if err != nil {
		return nil, err
	}

	return UniqueTraditions(append(out, randomTraditions...)), nil
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

func GetTraditionsSimilarityCoef(t1, t2 []*Tradition) float64 {
	// @TODO finish it

	return 0
}

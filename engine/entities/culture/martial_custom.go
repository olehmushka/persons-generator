package culture

import (
	"fmt"

	g "persons_generator/engine/entities/gender"
)

func (c *Culture) PrintMartialCustom() {
	if c == nil {
		return
	}
	fmt.Printf("MartialCustom: %s (culture_name: %s)\n", c.MartialCustom, c.Name)
}

func getMartialCustom(proto []*Culture) g.Acceptance {
	var (
		onlyMen     float64
		menAndWomen float64
		onlyWomen   float64
	)
	for _, c := range proto {
		switch c.MartialCustom {
		case g.OnlyMen:
			onlyMen++
		case g.MenAndWomen:
			menAndWomen++
		case g.OnlyWomen:
			onlyWomen++
		}
	}
	var (
		totalScore      = onlyMen + menAndWomen + onlyWomen
		onlyMenProb     = onlyMen / totalScore
		menAndWomenProb = menAndWomen / totalScore
		onlyWomenProb   = onlyWomen / totalScore
	)

	return g.GetAcceptanceByProbability(onlyMenProb, menAndWomenProb, onlyWomenProb)
}

package culture

import (
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/influence"
)

func getGenderDominance(proto []*Culture) *g.Dominance {
	var isMatriarchal bool
	for _, c := range proto {
		if c == nil || len(c.Traditions) == 0 {
			continue
		}

		for _, t := range c.Traditions {
			if t == nil {
				continue
			}
			if t.Name == MatriarchalTradition.Name {
				isMatriarchal = true
			}
		}
	}

	var (
		maleDominance   float64
		femaleDominance float64
		equalDominance  float64

		strongInfluence   float64
		moderateInfluence float64
		weakInfluence     float64
	)
	for _, c := range proto {
		switch c.GenderDominance.Dominance {
		case g.MaleDominance:
			maleDominance++
		case g.FemaleDominance:
			femaleDominance++
		case g.EqualityDominance:
			equalDominance++
		}
		switch c.GenderDominance.Influence {
		case influence.StrongInfluence:
			strongInfluence++
		case influence.ModerateInfluence:
			moderateInfluence++
		case influence.WeakInfluence:
			weakInfluence++
		}
	}
	var (
		totalDominancesScore = maleDominance + femaleDominance + equalDominance
		maleDominanceProb    = maleDominance / totalDominancesScore
		femaleDominanceProb  = femaleDominance / totalDominancesScore
		equalDominanceProb   = equalDominance / totalDominancesScore

		totalInfluencesScore  = strongInfluence + moderateInfluence + weakInfluence
		strongInfluenceProb   = strongInfluence / totalInfluencesScore
		moderateInfluenceProb = moderateInfluence / totalInfluencesScore
		weakInfluenceProb     = weakInfluence / totalInfluencesScore
	)

	if isMatriarchal {
		return g.NewDominanceWithParams(
			g.FemaleDominance,
			influence.GetInfluenceByProbability(strongInfluenceProb, moderateInfluenceProb, weakInfluenceProb),
		)
	}

	return g.NewDominanceWithParams(
		g.GetGenderDominanceByProbability(maleDominanceProb, equalDominanceProb, femaleDominanceProb),
		influence.GetInfluenceByProbability(strongInfluenceProb, moderateInfluenceProb, weakInfluenceProb),
	)
}

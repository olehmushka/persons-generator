package culture

import (
	g "persons_generator/engine/entities/gender"
	"persons_generator/engine/entities/influence"
)

func getGenderDominance(proto []*Culture) (*g.Dominance, error) {
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
		i, err := influence.GetInfluenceByProbability(strongInfluenceProb, moderateInfluenceProb, weakInfluenceProb)
		if err != nil {
			return nil, err
		}
		return g.NewDominanceWithParams(g.FemaleDominance, i), nil
	}
	dominance, err := g.GetGenderDominanceByProbability(maleDominanceProb, equalDominanceProb, femaleDominanceProb)
	if err != nil {
		return nil, err
	}
	i, err := influence.GetInfluenceByProbability(strongInfluenceProb, moderateInfluenceProb, weakInfluenceProb)
	if err != nil {
		return nil, err
	}

	return g.NewDominanceWithParams(dominance, i), nil
}

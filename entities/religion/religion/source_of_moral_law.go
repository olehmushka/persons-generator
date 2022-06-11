package religion

import pm "persons_generator/probability-machine"

type SourceOfMoralLaw string

const (
	DeitySourceOfMoralLaw  SourceOfMoralLaw = "Deity"
	NoneSourceOfMoralLaw   SourceOfMoralLaw = "None"
	HumanSourceOfMoralLaw  SourceOfMoralLaw = "Human"
	NatureSourceOfMoralLaw SourceOfMoralLaw = "Nature"
)

func (d *Doctrine) generateSourceOfMoralLaw() SourceOfMoralLaw {
	return ""
}

func geSourceOfMoralLawByProbability(deity, none, human, nature float64) SourceOfMoralLaw {
	m := map[string]float64{
		string(DeitySourceOfMoralLaw):  pm.PrepareProbability(deity),
		string(NoneSourceOfMoralLaw):   pm.PrepareProbability(none),
		string(HumanSourceOfMoralLaw):  pm.PrepareProbability(human),
		string(NatureSourceOfMoralLaw): pm.PrepareProbability(nature),
	}
	return SourceOfMoralLaw(pm.GetRandomFromSeveral(m))
}

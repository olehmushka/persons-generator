package religion

type Doctrine struct {
	religion *Religion
}

func NewDoctrine(r *Religion) *Doctrine {
	d := &Doctrine{religion: r}

	return d
}

type DeityNature string

const (
	MajoryGoodDeityNature DeityNature = "MajoryGood"
	MinoryGoodDeityNature DeityNature = "MinoryGood"
	CombinedDeityNature   DeityNature = "Combined"
	MinoryEvilDeityNature DeityNature = "MinoryEvil"
	MajoryEvilDeityNature DeityNature = "MajoryEvil"
)

type HumanNature string

const (
	GoodHumanNature    HumanNature = "Good"
	NeutralHumanNature HumanNature = "Neutral"
	EvilHumanNature    HumanNature = "Evil"
)

type SourceOfMoralLaw string

const (
	DeitySourceOfMoralLaw  SourceOfMoralLaw = "Deity"
	NoneSourceOfMoralLaw   SourceOfMoralLaw = "None"
	HumanSourceOfMoralLaw  SourceOfMoralLaw = "Human"
	NatureSourceOfMoralLaw SourceOfMoralLaw = "Nature"
)

package religion

type Doctrine struct {
	religion *Religion

	Deity            *DeityDoctrine
	Human            *HumanDoctrine
	Social           *SocialDoctrine
	SourceOfMoralLaw SourceOfMoralLaw
	Afterlife        *Afterlife
}

func NewDoctrine(r *Religion) *Doctrine {
	d := &Doctrine{religion: r}
	d.Deity = d.generateDeityDoctrine()
	d.Human = d.generateHumanDoctrine()
	d.Social = d.generateSocialDoctrine()
	d.SourceOfMoralLaw = d.generateSourceOfMoralLaw()
	d.Afterlife = d.generateAfterlife()

	return d
}

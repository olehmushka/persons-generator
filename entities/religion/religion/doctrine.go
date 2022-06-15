package religion

import "fmt"

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

func (d *Doctrine) Print() {
	fmt.Printf("Doctrine (religion=%s)\n", d.religion.Name)
	d.Deity.Print()
	d.Human.Print()
	d.Social.Print()
	fmt.Printf("Source of moral law is %s\n", d.SourceOfMoralLaw)
}

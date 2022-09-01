package religion

import "fmt"

type Doctrine struct {
	religion *Religion `json:"-"`

	HighGoal         *HighGoal        `json:"high_goal"`
	Deity            *DeityDoctrine   `json:"deitu"`
	Human            *HumanDoctrine   `json:"human"`
	Social           *SocialDoctrine  `json:"social"`
	SourceOfMoralLaw SourceOfMoralLaw `json:"source_of_moral_law"`
	Afterlife        *Afterlife       `json:"afterlife"`
}

func NewDoctrine(r *Religion) (*Doctrine, error) {
	d := &Doctrine{religion: r}
	hg, err := d.generateHighGoal()
	if err != nil {
		return nil, err
	}
	d.HighGoal = hg
	deity, err := d.generateDeityDoctrine()
	if err != nil {
		return nil, err
	}
	d.Deity = deity
	h, err := d.generateHumanDoctrine()
	if err != nil {
		return nil, err
	}
	d.Human = h
	social, err := d.generateSocialDoctrine()
	if err != nil {
		return nil, err
	}
	d.Social = social
	sourceOfMoralLaw, err := d.generateSourceOfMoralLaw()
	if err != nil {
		return nil, err
	}
	d.SourceOfMoralLaw = sourceOfMoralLaw
	al, err := d.generateAfterlife()
	if err != nil {
		return nil, err
	}
	d.Afterlife = al

	return d, nil
}

func (d *Doctrine) Print() {
	fmt.Printf("Doctrine (religion=%s)\n", d.religion.Name)
	d.HighGoal.Print()
	d.Deity.Print()
	d.Human.Print()
	d.Social.Print()
	d.SourceOfMoralLaw.Print()
	d.Afterlife.Print()
}

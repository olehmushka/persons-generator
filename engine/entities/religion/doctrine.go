package religion

import "fmt"

type Doctrine struct {
	religion *Religion `json:"-" bson:"-"`

	HighGoal         *HighGoal        `json:"high_goal" bson:"high_goal"`
	Deity            *DeityDoctrine   `json:"deitu" bson:"deitu"`
	Human            *HumanDoctrine   `json:"human" bson:"human"`
	Social           *SocialDoctrine  `json:"social" bson:"social"`
	SourceOfMoralLaw SourceOfMoralLaw `json:"source_of_moral_law" bson:"source_of_moral_law"`
	Afterlife        *Afterlife       `json:"afterlife" bson:"afterlife"`
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

func GetDoctrineSimilarityCoef(d1, d2 *Doctrine) float64 {
	if d1 == nil && d2 == nil {
		return 1
	}
	if d1 == nil || d2 == nil {
		return 0
	}

	similarityTraits := []struct {
		enable bool
		value  float64
		coef   float64
	}{
		{
			enable: true,
			value:  GetHighGoalSimilarityCoef(d1.HighGoal, d2.HighGoal),
			coef:   0.1,
		},
		{
			enable: true,
			value:  GetDeityDoctrineSimilarityCoef(d1.Deity, d2.Deity),
			coef:   0.2,
		},
		{
			enable: true,
			value:  GetHumanDoctrineSimilarityCoef(d1.Human, d2.Human),
			coef:   0.2,
		},
		{
			enable: true,
			value:  GetSocialDoctrineSimilarityCoef(d1.Social, d2.Social),
			coef:   0.15,
		},
		{
			enable: d1.SourceOfMoralLaw == d2.SourceOfMoralLaw,
			value:  1,
			coef:   0.2,
		},
		{
			enable: true,
			value:  GetAfterlifeSimilarityCoef(d1.Afterlife, d2.Afterlife),
			coef:   0.25,
		},
	}

	var out float64
	for _, t := range similarityTraits {
		if t.enable {
			out += t.value * t.coef
		}
	}

	return out
}

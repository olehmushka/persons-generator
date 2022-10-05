package entities

type Language struct {
	ID        string     `json:"id" bson:"id"`
	Name      string     `json:"name" bson:"name"`
	Subfamily *Subfamily `json:"subfamily" bson:"subfamily,omitempty"`
	WordBase  *WordBase  `json:"word_base" bson:"word_base,omitempty"`
	IsLiving  bool       `json:"is_living" bson:"is_living"`
}

type Subfamily struct {
	Name              string     `json:"name" bson:"name"`
	FamilyName        string     `json:"family" bson:"family"`
	ExtendedSubfamily *Subfamily `json:"extended_subfamily" bson:"extended_subfamily"`
	IsLiving          bool       `json:"is_living" bson:"is_living"`
}

type WordBase struct {
	FemaleOwnNames []string `json:"female_own_names" bson:"female_own_names"`
	MaleOwnNames   []string `json:"male_own_names" bson:"male_own_names"`
	Words          []string `json:"words" bson:"words"`
	Name           string   `json:"name" bson:"name"`
	Min            int      `json:"min" bson:"min"`
	Max            int      `json:"max" bson:"max"`
	Dupl           string   `json:"dupl" bson:"dupl"`
	M              float64  `json:"m" bson:"m"`
}

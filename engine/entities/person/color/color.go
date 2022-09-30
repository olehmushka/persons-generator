package color

type Color struct {
	Name    string `json:"name" bson:"name"`
	Hex     string `json:"hex" bson:"hex"`
	Palette string `json:"palette" bson:"palette"`
}

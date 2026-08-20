package entities

type Person struct {
	ID         string     `json:"id"`
	OwnName    string     `json:"own_name"`
	CultureID  string     `json:"culture_id,omitempty"`
	ReligionID string     `json:"religion_id,omitempty"`
	Human      *Human     `json:"human"`
	Traits     []string   `json:"traits"`
	Spouces    []string   `json:"spouces"`
	DeathYear  int        `json:"death_year"`
	FatherID   string     `json:"father_id"`
	MotherID   string     `json:"mother_id"`
	Coordinate Coordinate `json:"coordinate"`
	WorldID    string     `json:"world_id"`
}

type Human struct {
	Sex string `json:"sex"`
	Age int    `json:"age"`

	FaceShape         string  `json:"face_shape"`
	EyesColor         Color   `json:"eyes_color"`
	EyesType          string  `json:"eyes_type"`
	EarsType          string  `json:"ears_type"`
	NoseType          string  `json:"nose_type"`
	LipsType          string  `json:"lips_type"`
	HairColor         Color   `json:"hair_color"`
	ScalpHairTexture  string  `json:"scalp_hair_texture"`
	ScalpHairDensity  string  `json:"scalp_hair_density"`
	FaceHairDensity   string  `json:"face_hair_density"`
	HeightInCm        float64 `json:"height_in_cm"`
	ShoeEUSize        int     `json:"shoe_eu_size"`
	SkinColor         Color   `json:"skin_color"`
	SexualOrientation string  `json:"sexual_orientation"`
	Temperament       string  `json:"temperament"`
}

type Color struct {
	Name    string `json:"name"`
	Hex     string `json:"hex"`
	Palette string `json:"palette"`
}

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

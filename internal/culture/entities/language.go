package entities

type Language struct {
	Name      string             `json:"name"`
	Subfamily *LanguageSubfamily `json:"subfamily"`
}

package entities

type LanguageSubfamily struct {
	Name              string
	FamilyName        string
	ExtendedSubfamily *LanguageSubfamily
}

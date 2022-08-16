package entities

type Culture struct {
	Name            string
	Proto           []*Culture
	CultureGroup    *CultureGroup
	RootCultureName string
	Language        *Language
	EthosName       string
	Traditions      []string
	DominantSex     string
	MartialCustom   string
}

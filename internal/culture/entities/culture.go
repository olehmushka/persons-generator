package entities

type Culture struct {
	Name            string
	Proto           []*Culture
	WideCulture     *WideCulture
	RootCultureName string
	Language        *Language
	EthosName       string
	Traditions      []string
	DominantSex     string
	MartialCustom   string
}

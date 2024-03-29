package entities

type Culture struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	Proto           []*Culture    `json:"proto"`
	CultureGroup    *CultureGroup `json:"culture_group"`
	RootCultureName string        `json:"root_culture_name"`
	Language        *Language     `json:"language"`
	EthosName       string        `json:"ethos_name"`
	Traditions      []string      `jon:"traditions"`
	DominantSex     string        `json:"dominant_sex"`
	MartialCustom   string        `json:"martial_custom"`
}

package http

import "github.com/google/uuid"

type CulturePreferred struct {
	Names  []string `json:"names"`
	Amount int      `json:"amount"`
	Kind   string   `json:"kind"`
}

type PostCreatCulturesRequest struct {
	Preferred []*CulturePreferred `json:"preferred"`
	Amount    int                 `json:"amount"`
}

type PostCreateCulturesResponse struct {
	Data []*Culture `json:"data"`
}

type GetProtoCulturesResponse struct {
	Data  []*Culture `json:"data"`
	Total int        `json:"total"`
}

type GetCultureByIDResponse struct {
	Data *Culture `json:"data"`
}

type Culture struct {
	ID              uuid.UUID     `json:"id"`
	Name            string        `json:"name"`
	Proto           []*Culture    `json:"proto"`
	CultureGroup    *CultureGroup `json:"culture_group"`
	RootCultureName string        `json:"root_culture_name"`
	Language        *Language     `json:"language"`
	EthosName       string        `json:"ethos_name"`
	Traditions      []string      `json:"traditions"`
	DominantSex     string        `json:"dominant_sex"`
	MartialCustom   string        `json:"martial_custom"`
}

type LanguageSubfamily struct {
	Name              string             `json:"name"`
	FamilyName        string             `json:"family_name"`
	ExtendedSubfamily *LanguageSubfamily `json:"extended_subfamily"`
}

type Language struct {
	Name      string             `json:"name"`
	Subfamily *LanguageSubfamily `json:"subfamily"`
}

type CultureGroup struct {
	Name            string `json:"name"`
	RootCultureName string `json:"root_culture_name"`
}

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

type ReligionPreferred struct {
	CultureIDs []string `json:"culture_ids"`
	Amount     int      `json:"amount"`
}

type PostCreateReligionsRequest struct {
	Preferred []*ReligionPreferred `json:"preferred"`
	Amount    int                  `json:"amount"`
}

type PostCreateReligionsResponse struct {
	Data  []*Religion `json:"data"`
	Total int         `json:"total"`
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

type Religion struct {
	ID                  uuid.UUID         `json:"id"`
	Name                string            `json:"name"`
	Type                ReligionType      `json:"type"`
	DominantSex         string            `json:"dominant_sex"`
	HighGoals           []string          `json:"high_goals"`
	DeityGoodness       string            `json:"deity_goodness"`
	DeityTraits         []string          `json:"deity_traits"`
	HumanGoodness       string            `json:"human_goodness"`
	HumanTraits         []string          `json:"human_traits"`
	SocialTraits        []string          `json:"social_traits"`
	SourceOfMoralLaw    string            `json:"source_of_moral_law"`
	Afterlife           *Afterlife        `json:"afterlife"`
	Attributes          []string          `json:"attributes"`
	Clerics             *Clerics          `json:"clerics"`
	HolyScriptureTraits []string          `json:"holy_scripture_traits"`
	TempleTraits        []string          `json:"temple_traits"`
	TheologyTraits      []string          `json:"theology_traits"`
	Cults               []string          `json:"cults"`
	Rules               []string          `json:"rules"`
	Taboos              []Taboo           `json:"taboos"`
	Rituals             []string          `json:"rituals"`
	Holydays            []string          `json:"holydays"`
	ConversionTraits    []string          `json:"conversion_traits"`
	MarriageTradition   MarriageTradition `json:"marriage_tradition"`
}

type Afterlife struct {
	Participants AfterlifeParticipants `json:"participants"`
	Traits       []string              `json:"traits"`
}

type AfterlifeParticipants struct {
	ForTopBelievers    string `json:"for_top_believers"`
	ForBelievers       string `json:"for_believers"`
	ForUntrueBelievers string `json:"for_untrue_believers"`
	ForAtheists        string `json:"for_atheists"`
}

type Clerics struct {
	IsCivil                  bool     `json:"is_civil"`
	IsRevocable              bool     `json:"is_revocable"`
	AcceptableClericSex      string   `json:"acceptable_cleric_sex"`
	AcceptableClericMarriage string   `json:"acceptable_cleric_marriage"`
	Traits                   []string `json:"traits"`
	Functions                []string `json:"functions"`
}

type MarriageTradition struct {
	Kind          string `json:"kind"`
	Bastardry     string `json:"bastardry"`
	Consanguinity string `json:"consanguinity"`
	Divorce       string `json:"divorce"`
}

type Taboo struct {
	Action     string `json:"action"`
	Acceptance string `json:"acceptance"`
}

type ReligionType struct {
	Name        string `json:"name"`
	SubtypeName string `json:"subtype_name"`
}

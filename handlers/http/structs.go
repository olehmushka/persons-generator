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
	Data []*SerializedCulture `json:"data"`
}

type GetProtoCulturesResponse struct {
	Data  []*SerializedCulture `json:"data"`
	Total int                  `json:"total"`
}

type GetCultureByIDResponse struct {
	Data *SerializedCulture `json:"data"`
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
	Data  []*SerializedReligion `json:"data"`
	Total int                   `json:"total"`
}

type SerializedCulture struct {
	ID              uuid.UUID            `json:"id"`
	Name            string               `json:"name"`
	ProtoCultures   []*SerializedCulture `json:"proto_cultures"`
	AbstructCulture string               `json:"abstruct_culture"`
	Root            string               `json:"root"`
	Language        *SerializedLanguage  `json:"language"`
	Ethos           string               `json:"ethos"`
	Traditions      []string             `json:"traditions"`
	GenderDominance string               `json:"gender_dominance"`
	MartialCustom   string               `json:"martial_custom"`
}

type SerializedLanguage struct {
	Name      string                       `json:"name"`
	Subfamily *SerailizedLanguageSubfamily `json:"subfamily"`
}

type SerailizedLanguageSubfamily struct {
	Name              string                       `json:"name"`
	Family            string                       `json:"family"`
	ExtendedSubfamily *SerailizedLanguageSubfamily `json:"extended_subfamily"`
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

type SerializedReligion struct {
	ID                  uuid.UUID           `json:"id"`
	Name                string              `json:"name"`
	Type                string              `json:"type"`
	GenderDominance     string              `json:"gender_dominance"`
	HighGoals           []string            `json:"high_goals"`
	DeityNature         string              `json:"deity_nature"`
	DeityTraits         []string            `json:"deity_traits"`
	HumanNature         string              `json:"human_nature"`
	HumanTraits         []string            `json:"human_traits"`
	SocialTraits        []string            `json:"social_traits"`
	SourceOfMoralLaw    string              `json:"source_of_moral_law"`
	Afterlife           map[string]string   `json:"afterlife"`
	AfterlifeTraits     []string            `json:"afterlife_traits"`
	Traits              []string            `json:"traits"`
	ClericsAppointment  string              `json:"clerics_appointment"`
	ClericsLimitations  string              `json:"clerics_limitations"`
	ClericsTraits       []string            `json:"clerics_traits"`
	ClericsFunctions    []string            `json:"clerics_functions"`
	HolyScriptureTraits []string            `json:"holy_scripture_traits"`
	TempleTraits        []string            `json:"temple_traits"`
	TheologyTraits      []string            `json:"theology_traits"`
	Cults               []string            `json:"cults"`
	Rules               []string            `json:"rules"`
	Taboos              []map[string]string `json:"taboos"`
	Rituals             []string            `json:"rituals"`
	Holydays            []string            `json:"holydays"`
	Conversion          []string            `json:"conversion"`
	MarriageTradition   map[string]string   `json:"marriage_tradition"`
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

type CreateWorldPersonsRequest struct {
	PersonsAmount           *int     `json:"persons_amount"`
	MalePersonsAmount       *int     `json:"male_persons_amount"`
	FemalePersonsAmount     *int     `json:"female_persons_amount"`
	CultureReligionIDsPairs []string `json:"culture_religion_ids_pairs"`
}

type CreateWorldPersonsResponse struct {
	WorldID uuid.UUID `json:"world_id"`
}

type GetWorldPersonsResponse struct {
	Data  []*Person `json:"data"`
	Total int       `json:"total"`
}

type Person struct{}

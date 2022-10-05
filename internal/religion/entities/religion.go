package entities

type Religion struct {
	ID                  string            `json:"id"`
	Name                string            `json:"name"`
	Type                Type              `json:"type"`
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

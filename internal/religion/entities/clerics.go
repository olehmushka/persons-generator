package entities

type Clerics struct {
	IsCivil                  bool     `json:"is_civil"`
	IsRevocable              bool     `json:"is_revocable"`
	AcceptableClericSex      string   `json:"acceptable_cleric_sex"`
	AcceptableClericMarriage string   `json:"acceptable_cleric_marriage"`
	Traits                   []string `json:"traits"`
	Functions                []string `json:"functions"`
}

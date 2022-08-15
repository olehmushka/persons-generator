package entities

type CulturePreference struct {
	Names  []string `json:"names"`
	Amount int      `json:"amount"`
	Kind   string   `json:"kind"`
}

package culture

type Preference struct {
	Names  []string       `json:"names"`
	Amount int            `json:"amount"`
	Kind   PreferenceKind `json:"kind"`
}

type PreferenceKind string

const (
	StrictPrefKind PreferenceKind = "strict"
	HybridPrefKind PreferenceKind = "hybrid"
)

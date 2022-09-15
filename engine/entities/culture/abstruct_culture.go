package culture

type AbstructCulture struct {
	Name string `json:"name"`
	Root *Root  `json:"root"`
}

func (c *AbstructCulture) IsZero() bool {
	return c == nil
}

func GetAbstructCultureSimilarityCoef(c1, c2 *AbstructCulture) float64 {
	if c1.IsZero() && c2.IsZero() {
		return 1
	}
	if c1.IsZero() || c2.IsZero() {
		return 0
	}

	if c1.Name == c2.Name {
		return 1
	}
	if IsRootEqual(c1.Root, c2.Root) {
		return 0.5
	}

	return 0
}

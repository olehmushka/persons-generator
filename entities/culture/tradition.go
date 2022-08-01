package culture

type Tradition struct {
	Name              string
	PreferredEthoses  []*Ethos
	Type              TraditionType
	CompatibilityFunc func(c *Culture) bool
}

type TraditionType string

const (
	RealmTraditionType    TraditionType = "realm_tradition_type"
	WarfareTraditionType  TraditionType = "warfare_tradition_type"
	SocialTraditionType   TraditionType = "social_tradition_type"
	RitualTraditionType   TraditionType = "ritual_tradition_type"
	RegionalTraditionType TraditionType = "regional_tradition_type"
)

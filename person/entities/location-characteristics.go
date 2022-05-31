package entities

type LocationType string

const (
	LocationTypeCity         LocationType = "city"
	LocationTypeTown         LocationType = "town"
	LocationTypeSmallTown    LocationType = "small_town"
	LocationTypeBigVillage   LocationType = "big_village"
	LocationTypeVillage      LocationType = "village"
	LocationTypeSmallVillage LocationType = "small_village"
	LocationTypeForest       LocationType = "forest"
	LocationTypeIsland       LocationType = "island"
	LocationTypeDesert       LocationType = "desert"
)

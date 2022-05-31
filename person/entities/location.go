package entities

import "persons_generator/vector"

type Country struct {
	Name         string
	MainCultures vector.Vector[*Culture]
	CountryType  CountryType
	Districts    vector.Vector[*District]
}

type District struct {
	Name         string
	MainCultures vector.Vector[*Culture]
	Locations    vector.Vector[*Location]
	Country      *Country
}

type Location struct {
	Name         string
	MainCultures vector.Vector[*Culture]
	LocationType LocationType
	District     *District
	Country      *Country
}

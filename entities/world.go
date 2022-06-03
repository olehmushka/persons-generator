package entities

type World struct {
	Size      int
	Locations [][]*Location
}

type Location struct {
	Coordinates     *Coordinate
	PrimaryCulture  *Culture
	PrimaryReligion *Religion
	Population      []*Human
}

type Coordinate struct {
	X, Y int
}

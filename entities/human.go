package entities

type Human struct {
	Name    *HumanName
	Body    *Body
	Psycho  *Psycho
	Parents *Parents
	Home    *Location
}

type Parents struct {
	Father *Human
	Mother *Human
}

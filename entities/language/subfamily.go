package language

type Subfamily struct {
	Name              string
	Family            Family
	ExtendedSubfamily *Subfamily
	IsLiving          bool
}

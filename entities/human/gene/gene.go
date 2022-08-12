package gene

type Gene interface {
	Byteble
	Type() string
	Produce() (Byteble, error)
	Children() []Gene
}

func Merge(g1, g2 Gene) Gene {
	return nil
}

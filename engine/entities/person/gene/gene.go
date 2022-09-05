package gene

import "persons_generator/engine/entities/gender"

type Gene interface {
	Byteble
	Type() string
	Produce(sex gender.Sex) (Byteble, error)
	Children() []Gene
}

func Merge(g1, g2 Gene) Gene {
	return nil
}

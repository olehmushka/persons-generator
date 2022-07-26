package culture

import "persons_generator/entities/religion/religion"

type Params struct {
	Religion *religion.Religion

	MinInheritedBaseCultures int
	MaxInheritedBaseCultures int
}

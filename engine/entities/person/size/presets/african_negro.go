package presets

import "persons_generator/engine/entities/person/size"

var AfricanNegroSizePreset = size.NewSizeGene(
	size.NewHeightGene(155, 175, 205),
	size.NewShoeSizeGene(22, 25, 28),
)

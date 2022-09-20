package presets

import "persons_generator/engine/entities/person/size"

var NegritoSizePreset = size.NewSizeGene(
	size.NewHeightGene(155, 170, 190),
	size.NewShoeSizeGene(22, 24, 27),
)

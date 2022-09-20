package presets

import "persons_generator/engine/entities/person/size"

var AlpineSizePreset = size.NewSizeGene(
	size.NewHeightGene(165, 177, 205),
	size.NewShoeSizeGene(23, 25, 28),
)

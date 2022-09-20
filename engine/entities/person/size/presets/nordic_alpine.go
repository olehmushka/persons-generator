package presets

import "persons_generator/engine/entities/person/size"

var NordicAlpineSizePreset = size.NewSizeGene(
	size.NewHeightGene(157, 175.5, 205),
	size.NewShoeSizeGene(22, 25, 28),
)

package presets

import "persons_generator/engine/entities/person/size"

var EastBalticSizePreset = size.NewSizeGene(
	size.NewHeightGene(157, 177, 205),
	size.NewShoeSizeGene(22, 25, 28),
)

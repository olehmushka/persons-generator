package presets

import "persons_generator/engine/entities/person/size"

var DinaricSizePreset = size.NewSizeGene(
	size.NewHeightGene(156, 175, 205),
	size.NewShoeSizeGene(22, 25, 28),
)

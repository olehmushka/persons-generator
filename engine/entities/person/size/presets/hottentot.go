package presets

import "persons_generator/engine/entities/person/size"

var HottentotSizePreset = size.NewSizeGene(
	size.NewHeightGene(155, 179, 205),
	size.NewShoeSizeGene(22, 25.5, 28),
)

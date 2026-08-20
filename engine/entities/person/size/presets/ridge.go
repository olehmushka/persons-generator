package presets

import "persons_generator/engine/entities/person/size"

var RidgeSizePreset = size.NewSizeGene(
	size.NewHeightGene(155, 168, 195),
	size.NewShoeSizeGene(22, 24, 28),
)

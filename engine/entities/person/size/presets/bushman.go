package presets

import "persons_generator/engine/entities/person/size"

var BushmanSizePreset = size.NewSizeGene(
	size.NewHeightGene(160, 180, 210),
	size.NewShoeSizeGene(23, 25.5, 28),
)

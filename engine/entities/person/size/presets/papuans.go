package presets

import "persons_generator/engine/entities/person/size"

var PapuansSizePreset = size.NewSizeGene(
	size.NewHeightGene(155, 170, 190),
	size.NewShoeSizeGene(22, 24, 28),
)

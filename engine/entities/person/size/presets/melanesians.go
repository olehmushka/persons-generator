package presets

import "persons_generator/engine/entities/person/size"

var MelanesiansSizePreset = size.NewSizeGene(
	size.NewHeightGene(155, 167, 195),
	size.NewShoeSizeGene(22, 24.75, 28),
)

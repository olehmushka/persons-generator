package presets

import "persons_generator/engine/entities/person/size"

var ClassicMongoloidSizePreset = size.NewSizeGene(
	size.NewHeightGene(155, 166.5, 190),
	size.NewShoeSizeGene(22, 23.75, 27),
)

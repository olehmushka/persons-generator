package presets

import "persons_generator/engine/entities/person/size"

var ArcticMongoloidSizePreset = size.NewSizeGene(
	size.NewHeightGene(155, 165, 190),
	size.NewShoeSizeGene(22, 24, 27),
)

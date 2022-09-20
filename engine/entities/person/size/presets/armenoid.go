package presets

import "persons_generator/engine/entities/person/size"

var ArmenoidSizePreset = size.NewSizeGene(
	size.NewHeightGene(155, 172, 200),
	size.NewShoeSizeGene(22, 24, 27.5),
)

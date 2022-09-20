package presets

import "persons_generator/engine/entities/person/size"

var NordicSizePreset = size.NewSizeGene(
	size.NewHeightGene(160, 179, 207),
	size.NewShoeSizeGene(22.5, 25.5, 28),
)

package presets

import "persons_generator/engine/entities/person/size"

var NiloticNegroSizePreset = size.NewSizeGene(
	size.NewHeightGene(155, 180, 210),
	size.NewShoeSizeGene(23, 25.5, 28),
)

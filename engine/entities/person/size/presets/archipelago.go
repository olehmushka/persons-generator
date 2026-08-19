package presets

import "persons_generator/engine/entities/person/size"

var ArchipelagoSizePreset = size.NewSizeGene(
	size.NewHeightGene(169, 180, 210),
	size.NewShoeSizeGene(24, 27, 28),
)

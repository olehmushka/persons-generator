package presets

import "persons_generator/engine/entities/person/size"

var PlateauSizePreset = size.NewSizeGene(
	size.NewHeightGene(153, 170, 195),
	size.NewShoeSizeGene(22, 24, 27),
)

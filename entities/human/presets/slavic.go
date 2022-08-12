package presets

import (
	"persons_generator/entities/human/body/presets"
	"persons_generator/entities/human/human"
)

var SlavicHumanPreset = human.NewHumanGene(
	presets.SlavicBodyPreset,
)

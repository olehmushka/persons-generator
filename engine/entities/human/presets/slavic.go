package presets

import (
	"persons_generator/engine/entities/human/body/presets"
	"persons_generator/engine/entities/human/human"
)

var SlavicHumanPreset = human.NewHumanGene(
	presets.SlavicBodyPreset,
)

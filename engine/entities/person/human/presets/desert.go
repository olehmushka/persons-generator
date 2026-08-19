package presets

import (
	bodyPresets "persons_generator/engine/entities/person/body/presets"
	"persons_generator/engine/entities/person/human"
	psychoPresets "persons_generator/engine/entities/person/psycho/presets"
)

var DesertHumanPreset = human.NewGene(
	bodyPresets.DesertBodyPreset,
	psychoPresets.DesertPsychoPreset,
)

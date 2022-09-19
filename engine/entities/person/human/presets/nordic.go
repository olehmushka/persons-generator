package presets

import (
	bodyPresets "persons_generator/engine/entities/person/body/presets"
	"persons_generator/engine/entities/person/human"
	psychoPresets "persons_generator/engine/entities/person/psycho/presets"
)

var NordicHumanPreset = human.NewGene(
	bodyPresets.NordicBodyPreset,
	psychoPresets.NordicPsychoPreset,
)

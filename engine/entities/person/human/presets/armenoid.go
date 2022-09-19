package presets

import (
	bodyPresets "persons_generator/engine/entities/person/body/presets"
	"persons_generator/engine/entities/person/human"
	psychoPresets "persons_generator/engine/entities/person/psycho/presets"
)

var ArmenoidHumanPreset = human.NewGene(
	bodyPresets.ArmenoidBodyPreset,
	psychoPresets.ArmenoidPsychoPreset,
)

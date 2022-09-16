package presets

import (
	"persons_generator/engine/entities/person/psycho"
	temperamentPresets "persons_generator/engine/entities/person/temperament/presets"
)

var DinaricPsychoPreset = psycho.NewPsychoGene(
	temperamentPresets.DinaricTemperamentPreset,
)

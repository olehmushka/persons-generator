package presets

import (
	"persons_generator/engine/entities/person/body"
	facePresets "persons_generator/engine/entities/person/face/presets"
	hairPresets "persons_generator/engine/entities/person/hair/presets"
)

var SlavicBodyPreset = body.NewBodyGene(
	facePresets.SlavicFacePreset,
	hairPresets.SlavicHairPreset,
)
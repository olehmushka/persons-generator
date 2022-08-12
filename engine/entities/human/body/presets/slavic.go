package presets

import (
	"persons_generator/engine/entities/human/body"
	"persons_generator/engine/entities/human/face/presets"
)

var SlavicBodyPreset = body.NewBodyGene(
	presets.SlavicFacePreset,
)

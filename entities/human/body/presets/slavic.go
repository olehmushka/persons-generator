package presets

import (
	"persons_generator/entities/human/body"
	"persons_generator/entities/human/face/presets"
)

var SlavicBodyPreset = body.NewBodyGene(
	presets.SlavicFacePreset,
)

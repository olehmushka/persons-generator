package presets

import (
	"persons_generator/engine/entities/person/body"
	facePresets "persons_generator/engine/entities/person/face/presets"
	hairPresets "persons_generator/engine/entities/person/hair/presets"
	sizePresets "persons_generator/engine/entities/person/size/presets"
	skinPresets "persons_generator/engine/entities/person/skin/presets"
)

var MalanesiansBodyPreset = body.NewBodyGene(
	facePresets.MelanesiansFacePreset,
	hairPresets.MalanesiansHairPreset,
	sizePresets.MalanesiansSizePreset,
	skinPresets.MalanesiansSkinPreset,
)

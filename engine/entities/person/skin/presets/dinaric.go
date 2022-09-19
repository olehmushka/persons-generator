package presets

import (
	"persons_generator/engine/entities/person/skin"
)

var DinaricSkinPreset = skin.NewSkinGene(
	skin.NewSkinColorGene(map[string]float64{}, 1),
)

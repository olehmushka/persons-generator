package presets

import (
	"persons_generator/engine/entities/person/skin"
)

var AinuSkinPreset = skin.NewSkinGene(
	skin.NewSkinColorGene(map[string]float64{}, 1),
)

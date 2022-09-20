package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/skin"
)

var MelanesiansSkinPreset = skin.NewSkinGene(
	skin.NewSkinColorGene(map[string]float64{
		color.BrownSkinColorPalette: 1,
	}, 3),
)

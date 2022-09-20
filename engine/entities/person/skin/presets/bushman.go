package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/skin"
)

var BushmanSkinPreset = skin.NewSkinGene(
	skin.NewSkinColorGene(map[string]float64{
		color.BrownSkinColorPalette: 1,
	}, 1),
)

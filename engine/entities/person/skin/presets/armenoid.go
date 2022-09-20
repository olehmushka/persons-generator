package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/skin"
)

var ArmenoidSkinPreset = skin.NewSkinGene(
	skin.NewSkinColorGene(map[string]float64{
		color.OliveSkinColorPalette: 1,
	}, 3),
)

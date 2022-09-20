package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/skin"
)

var CelticSkinPreset = skin.NewSkinGene(
	skin.NewSkinColorGene(map[string]float64{
		color.CaucasianSkinColorPalette: 1,
	}, 0),
)

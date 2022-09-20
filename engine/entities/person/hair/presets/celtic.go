package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/hair"
)

var CelticHairPreset = hair.NewHairGene(
	hair.NewScalpHairGene(
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type):    0,
			string(hair.CurlyHairTexture.Type):    0.3,
			string(hair.FineHairTexture.Type):     0.2,
			string(hair.StraightHairTexture.Type): 0.15,
			string(hair.ThickHairTexture.Type):    0.15,
			string(hair.ThinHairTexture.Type):     0.15,
			string(hair.WavyHairTexture.Type):     0.3,
		}),
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.25,
				hair.MediumHairDensity.Type: 0.35,
				hair.HighHairDensity.Type:   0.8,
			},
		),
	),
	hair.NewFaceHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.3,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.3,
			},
		),
	),
	hair.NewAxillaryHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.3,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.3,
			},
		),
	),
	hair.NewChestAndAbdomenHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.3,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.3,
			},
		),
	),
	hair.NewArmsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.3,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.3,
			},
		),
	),
	hair.NewPubicHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.3,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.3,
			},
		),
	),
	hair.NewLegsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.3,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.3,
			},
		),
	),
	hair.NewFeetHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.3,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.3,
			},
		),
	),
	hair.NewHairColorGene(map[string]float64{
		color.BlackHairColorPalette:            0.01,
		color.MediumBrownHairColorPalette:      0.35,
		color.LightBrownHairColorPalette:       0.35,
		color.BlondeHairColorPalette:           0.35,
		color.DarkRedHairColorPalette:          0.4,
		color.StrawberryBlondeHairColorPalette: 0.35,
	}),
)

package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/hair"
)

var NordicAlpineHairPreset = hair.NewHairGene(
	hair.NewScalpHairGene(
		hair.NewHairTextureGene(map[string]float64{
			string(hair.CurlyHairTexture.Type):    0.1,
			string(hair.FineHairTexture.Type):     0.2,
			string(hair.StraightHairTexture.Type): 0.6,
			string(hair.ThickHairTexture.Type):    0.2,
			string(hair.ThinHairTexture.Type):     0.2,
			string(hair.WavyHairTexture.Type):     0.15,
		}),
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.9,
			},
		),
	),
	hair.NewFaceHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.4,
			},
		),
	),
	hair.NewAxillaryHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.4,
			},
		),
	),
	hair.NewChestAndAbdomenHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.4,
			},
		),
	),
	hair.NewArmsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.4,
			},
		),
	),
	hair.NewPubicHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.4,
			},
		),
	),
	hair.NewLegsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.4,
			},
		),
	),
	hair.NewFeetHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.4,
			},
		),
	),
	hair.NewHairColorGene(map[string]float64{
		color.BlackHairColorPalette:            0.05,
		color.MediumBrownHairColorPalette:      0.4,
		color.LightBrownHairColorPalette:       0.5,
		color.BlondeHairColorPalette:           0.7,
		color.DarkRedHairColorPalette:          0.1,
		color.StrawberryBlondeHairColorPalette: 0.6,
	}),
)

package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/hair"
)

var AmericanIndiansHairPreset = hair.NewHairGene(
	hair.NewScalpHairGene(
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type):    0,
			string(hair.CurlyHairTexture.Type):    0,
			string(hair.FineHairTexture.Type):     0.1,
			string(hair.StraightHairTexture.Type): 0.5,
			string(hair.ThickHairTexture.Type):    0.2,
			string(hair.ThinHairTexture.Type):     0.2,
			string(hair.WavyHairTexture.Type):     0.05,
		}),
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.25,
				hair.HighHairDensity.Type:   0.7,
			},
		),
	),
	hair.NewFaceHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.8,
				hair.MediumHairDensity.Type: 0.15,
				hair.HighHairDensity.Type:   0.1,
			},
		),
	),
	hair.NewAxillaryHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.8,
				hair.MediumHairDensity.Type: 0.15,
				hair.HighHairDensity.Type:   0.1,
			},
		),
	),
	hair.NewChestAndAbdomenHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.8,
				hair.MediumHairDensity.Type: 0.15,
				hair.HighHairDensity.Type:   0.1,
			},
		),
	),
	hair.NewArmsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.8,
				hair.MediumHairDensity.Type: 0.15,
				hair.HighHairDensity.Type:   0.1,
			},
		),
	),
	hair.NewPubicHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.8,
				hair.MediumHairDensity.Type: 0.15,
				hair.HighHairDensity.Type:   0.1,
			},
		),
	),
	hair.NewLegsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.8,
				hair.MediumHairDensity.Type: 0.15,
				hair.HighHairDensity.Type:   0.1,
			},
		),
	),
	hair.NewFeetHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.8,
				hair.MediumHairDensity.Type: 0.15,
				hair.HighHairDensity.Type:   0.1,
			},
		),
	),
	hair.NewHairColorGene(map[string]float64{
		color.BlackHairColorPalette:            1,
		color.MediumBrownHairColorPalette:      0,
		color.LightBrownHairColorPalette:       0,
		color.BlondeHairColorPalette:           0,
		color.DarkRedHairColorPalette:          0,
		color.StrawberryBlondeHairColorPalette: 0,
	}),
)

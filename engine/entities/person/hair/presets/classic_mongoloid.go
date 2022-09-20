package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/hair"
)

var ClassicMongoloidHairPreset = hair.NewHairGene(
	hair.NewScalpHairGene(
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type):    0,
			string(hair.CurlyHairTexture.Type):    0.25,
			string(hair.FineHairTexture.Type):     0.05,
			string(hair.StraightHairTexture.Type): 0.25,
			string(hair.ThickHairTexture.Type):    0.2,
			string(hair.ThinHairTexture.Type):     0.2,
			string(hair.WavyHairTexture.Type):     0.35,
		}),
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.15,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.8,
			},
		),
	),
	hair.NewFaceHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.65,
				hair.MediumHairDensity.Type: 0.25,
				hair.HighHairDensity.Type:   0.1,
			},
		),
	),
	hair.NewAxillaryHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.65,
				hair.MediumHairDensity.Type: 0.25,
				hair.HighHairDensity.Type:   0.1,
			},
		),
	),
	hair.NewChestAndAbdomenHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.65,
				hair.MediumHairDensity.Type: 0.25,
				hair.HighHairDensity.Type:   0.1,
			},
		),
	),
	hair.NewArmsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.65,
				hair.MediumHairDensity.Type: 0.25,
				hair.HighHairDensity.Type:   0.1,
			},
		),
	),
	hair.NewPubicHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.65,
				hair.MediumHairDensity.Type: 0.25,
				hair.HighHairDensity.Type:   0.1,
			},
		),
	),
	hair.NewLegsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.65,
				hair.MediumHairDensity.Type: 0.25,
				hair.HighHairDensity.Type:   0.1,
			},
		),
	),
	hair.NewFeetHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.65,
				hair.MediumHairDensity.Type: 0.25,
				hair.HighHairDensity.Type:   0.1,
			},
		),
	),
	hair.NewHairColorGene(map[string]float64{
		color.BlackHairColorPalette:            0.8,
		color.MediumBrownHairColorPalette:      0.01,
		color.LightBrownHairColorPalette:       0.0001,
		color.BlondeHairColorPalette:           0.00001,
		color.DarkRedHairColorPalette:          0,
		color.StrawberryBlondeHairColorPalette: 0.005,
	}),
)

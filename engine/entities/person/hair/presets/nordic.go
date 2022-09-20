package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/hair"
)

var NordicHairPreset = hair.NewHairGene(
	hair.NewScalpHairGene(
		hair.NewHairTextureGene(map[string]float64{
			string(hair.CurlyHairTexture.Type):    0.05,
			string(hair.FineHairTexture.Type):     0.1,
			string(hair.StraightHairTexture.Type): 0.5,
			string(hair.ThickHairTexture.Type):    0.4,
			string(hair.ThinHairTexture.Type):     0.5,
			string(hair.WavyHairTexture.Type):     0.3,
		}),
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.35,
				hair.HighHairDensity.Type:   0.75,
			},
		),
	),
	hair.NewFaceHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.35,
				hair.HighHairDensity.Type:   0.65,
			},
		),
	),
	hair.NewAxillaryHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.2,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.5,
			},
		),
	),
	hair.NewChestAndAbdomenHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.2,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.5,
			},
		),
	),
	hair.NewArmsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.2,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.5,
			},
		),
	),
	hair.NewPubicHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.2,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.5,
			},
		),
	),
	hair.NewLegsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.2,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.5,
			},
		),
	),
	hair.NewFeetHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.2,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.5,
			},
		),
	),
	hair.NewHairColorGene(map[string]float64{
		color.BlackHairColorPalette:            0.05,
		color.MediumBrownHairColorPalette:      0.3,
		color.LightBrownHairColorPalette:       0.6,
		color.BlondeHairColorPalette:           0.5,
		color.DarkRedHairColorPalette:          0.25,
		color.StrawberryBlondeHairColorPalette: 0.5,
	}),
)

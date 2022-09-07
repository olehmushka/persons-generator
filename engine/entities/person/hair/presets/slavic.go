package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/hair"
)

var SlavicHairPreset = hair.NewHairGene(
	hair.NewScalpHairGene(
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type):    0.001,
			string(hair.CurlyHairTexture.Type):    0.15,
			string(hair.FineHairTexture.Type):     0.25,
			string(hair.StraightHairTexture.Type): 0.25,
			string(hair.ThickHairTexture.Type):    0.15,
			string(hair.ThinHairTexture.Type):     0.15,
			string(hair.WavyHairTexture.Type):     0.25,
		}),
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.2,
				hair.MediumHairDensity.Type: 0.35,
				hair.HighHairDensity.Type:   0.2,
			},
		),
	),
	hair.NewFaceHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.2,
				hair.MediumHairDensity.Type: 0.35,
				hair.HighHairDensity.Type:   0.2,
			},
		),
	),
	hair.NewAxillaryHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.2,
				hair.MediumHairDensity.Type: 0.35,
				hair.HighHairDensity.Type:   0.2,
			},
		),
	),
	hair.NewChestAndAbdomenHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.2,
				hair.MediumHairDensity.Type: 0.35,
				hair.HighHairDensity.Type:   0.2,
			},
		),
	),
	hair.NewArmsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.2,
				hair.MediumHairDensity.Type: 0.35,
				hair.HighHairDensity.Type:   0.2,
			},
		),
	),
	hair.NewPubicHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.2,
				hair.MediumHairDensity.Type: 0.35,
				hair.HighHairDensity.Type:   0.2,
			},
		),
	),
	hair.NewLegsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.2,
				hair.MediumHairDensity.Type: 0.35,
				hair.HighHairDensity.Type:   0.2,
			},
		),
	),
	hair.NewFeetHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.2,
				hair.MediumHairDensity.Type: 0.35,
				hair.HighHairDensity.Type:   0.2,
			},
		),
	),
	hair.NewHairColorGene(map[string]float64{
		color.BlackHairColorPalette:            0.1,
		color.MediumBrownHairColorPalette:      0.2,
		color.LightBrownHairColorPalette:       0.2,
		color.BlondeHairColorPalette:           0.15,
		color.DarkRedHairColorPalette:          0.1,
		color.StrawberryBlondeHairColorPalette: 0.2,
	}),
)

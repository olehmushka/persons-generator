package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/hair"
)

var PapuansHairPreset = hair.NewHairGene(
	hair.NewScalpHairGene(
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type):    0,
			string(hair.CurlyHairTexture.Type):    0,
			string(hair.FineHairTexture.Type):     0,
			string(hair.StraightHairTexture.Type): 0,
			string(hair.ThickHairTexture.Type):    0,
			string(hair.ThinHairTexture.Type):     0,
			string(hair.WavyHairTexture.Type):     0,
		}),
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0,
				hair.MediumHairDensity.Type: 0,
				hair.HighHairDensity.Type:   0,
			},
		),
	),
	hair.NewFaceHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0,
				hair.MediumHairDensity.Type: 0,
				hair.HighHairDensity.Type:   0,
			},
		),
	),
	hair.NewAxillaryHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0,
				hair.MediumHairDensity.Type: 0,
				hair.HighHairDensity.Type:   0,
			},
		),
	),
	hair.NewChestAndAbdomenHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0,
				hair.MediumHairDensity.Type: 0,
				hair.HighHairDensity.Type:   0,
			},
		),
	),
	hair.NewArmsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0,
				hair.MediumHairDensity.Type: 0,
				hair.HighHairDensity.Type:   0,
			},
		),
	),
	hair.NewPubicHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0,
				hair.MediumHairDensity.Type: 0,
				hair.HighHairDensity.Type:   0,
			},
		),
	),
	hair.NewLegsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0,
				hair.MediumHairDensity.Type: 0,
				hair.HighHairDensity.Type:   0,
			},
		),
	),
	hair.NewFeetHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0,
				hair.MediumHairDensity.Type: 0,
				hair.HighHairDensity.Type:   0,
			},
		),
	),
	hair.NewHairColorGene(map[string]float64{
		color.BlackHairColorPalette:            0,
		color.MediumBrownHairColorPalette:      0,
		color.LightBrownHairColorPalette:       0,
		color.BlondeHairColorPalette:           0,
		color.DarkRedHairColorPalette:          0,
		color.StrawberryBlondeHairColorPalette: 0,
	}),
)

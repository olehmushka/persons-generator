package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/hair"
)

var NordicMedititerraneanHairPreset = hair.NewHairGene(
	hair.NewScalpHairGene(
		hair.NewHairTextureGene(map[string]float64{
			string(hair.CurlyHairTexture.Type):    0.3,
			string(hair.FineHairTexture.Type):     0.1,
			string(hair.StraightHairTexture.Type): 0.25,
			string(hair.ThickHairTexture.Type):    0.2,
			string(hair.ThinHairTexture.Type):     0.2,
			string(hair.WavyHairTexture.Type):     0.25,
		}),
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.8,
			},
		),
	),
	hair.NewFaceHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.35,
				hair.HighHairDensity.Type:   0.7,
			},
		),
	),
	hair.NewAxillaryHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.15,
				hair.MediumHairDensity.Type: 0.45,
				hair.HighHairDensity.Type:   0.75,
			},
		),
	),
	hair.NewChestAndAbdomenHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.15,
				hair.MediumHairDensity.Type: 0.45,
				hair.HighHairDensity.Type:   0.75,
			},
		),
	),
	hair.NewArmsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.15,
				hair.MediumHairDensity.Type: 0.45,
				hair.HighHairDensity.Type:   0.75,
			},
		),
	),
	hair.NewPubicHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.15,
				hair.MediumHairDensity.Type: 0.45,
				hair.HighHairDensity.Type:   0.75,
			},
		),
	),
	hair.NewLegsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.15,
				hair.MediumHairDensity.Type: 0.45,
				hair.HighHairDensity.Type:   0.75,
			},
		),
	),
	hair.NewFeetHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.15,
				hair.MediumHairDensity.Type: 0.45,
				hair.HighHairDensity.Type:   0.75,
			},
		),
	),
	hair.NewHairColorGene(map[string]float64{
		color.BlackHairColorPalette:            0.05,
		color.MediumBrownHairColorPalette:      0.1,
		color.LightBrownHairColorPalette:       0.25,
		color.BlondeHairColorPalette:           0.25,
		color.DarkRedHairColorPalette:          0.05,
		color.StrawberryBlondeHairColorPalette: 0.3,
	}),
)

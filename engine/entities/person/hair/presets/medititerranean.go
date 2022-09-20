package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/hair"
)

var MedititerraneanHairPreset = hair.NewHairGene(
	hair.NewScalpHairGene(
		hair.NewHairTextureGene(map[string]float64{
			string(hair.CurlyHairTexture.Type):    0.35,
			string(hair.FineHairTexture.Type):     0.15,
			string(hair.StraightHairTexture.Type): 0.275,
			string(hair.ThickHairTexture.Type):    0.275,
			string(hair.ThinHairTexture.Type):     0.275,
			string(hair.WavyHairTexture.Type):     0.35,
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
				hair.LowHairDensity.Type:    0.4,
				hair.MediumHairDensity.Type: 0.5,
				hair.HighHairDensity.Type:   0.6,
			},
		),
	),
	hair.NewAxillaryHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.7,
			},
		),
	),
	hair.NewChestAndAbdomenHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.7,
			},
		),
	),
	hair.NewArmsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.7,
			},
		),
	),
	hair.NewPubicHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.7,
			},
		),
	),
	hair.NewLegsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.7,
			},
		),
	),
	hair.NewFeetHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.6,
				hair.HighHairDensity.Type:   0.7,
			},
		),
	),
	hair.NewHairColorGene(map[string]float64{
		color.BlackHairColorPalette:            0.6,
		color.MediumBrownHairColorPalette:      0.1,
		color.LightBrownHairColorPalette:       0.05,
		color.BlondeHairColorPalette:           0.05,
		color.DarkRedHairColorPalette:          0.025,
		color.StrawberryBlondeHairColorPalette: 0.05,
	}),
)

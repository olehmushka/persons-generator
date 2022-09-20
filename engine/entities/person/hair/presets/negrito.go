package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/hair"
)

var NegritoHairPreset = hair.NewHairGene(
	hair.NewScalpHairGene(
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type):    0.9,
			string(hair.CurlyHairTexture.Type):    0,
			string(hair.FineHairTexture.Type):     0,
			string(hair.StraightHairTexture.Type): 0.1,
			string(hair.ThickHairTexture.Type):    0,
			string(hair.ThinHairTexture.Type):     0,
			string(hair.WavyHairTexture.Type):     0,
		}),
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.7,
			},
		),
	),
	hair.NewFaceHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.9,
				hair.MediumHairDensity.Type: 0.1,
				hair.HighHairDensity.Type:   0.01,
			},
		),
	),
	hair.NewAxillaryHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.95,
				hair.MediumHairDensity.Type: 0.05,
				hair.HighHairDensity.Type:   0.01,
			},
		),
	),
	hair.NewChestAndAbdomenHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.95,
				hair.MediumHairDensity.Type: 0.05,
				hair.HighHairDensity.Type:   0.01,
			},
		),
	),
	hair.NewArmsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.95,
				hair.MediumHairDensity.Type: 0.05,
				hair.HighHairDensity.Type:   0.01,
			},
		),
	),
	hair.NewPubicHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.95,
				hair.MediumHairDensity.Type: 0.05,
				hair.HighHairDensity.Type:   0.01,
			},
		),
	),
	hair.NewLegsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.95,
				hair.MediumHairDensity.Type: 0.05,
				hair.HighHairDensity.Type:   0.01,
			},
		),
	),
	hair.NewFeetHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.95,
				hair.MediumHairDensity.Type: 0.05,
				hair.HighHairDensity.Type:   0.01,
			},
		),
	),
	hair.NewHairColorGene(map[string]float64{
		color.BlackHairColorPalette: 1,
	}),
)

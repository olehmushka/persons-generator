package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/hair"
)

var PapuansHairPreset = hair.NewHairGene(
	hair.NewScalpHairGene(
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type): 1,
		}),
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.05,
				hair.MediumHairDensity.Type: 0.1,
				hair.HighHairDensity.Type:   0.9,
			},
		),
	),
	hair.NewFaceHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.1,
				hair.MediumHairDensity.Type: 0.2,
				hair.HighHairDensity.Type:   0.7,
			},
		),
	),
	hair.NewAxillaryHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.6,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.2,
			},
		),
	),
	hair.NewChestAndAbdomenHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.6,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.2,
			},
		),
	),
	hair.NewArmsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.6,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.2,
			},
		),
	),
	hair.NewPubicHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.6,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.2,
			},
		),
	),
	hair.NewLegsHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.6,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.2,
			},
		),
	),
	hair.NewFeetHairGene(
		hair.NewHairDensityGene(
			map[string]float64{
				hair.LowHairDensity.Type:    0.6,
				hair.MediumHairDensity.Type: 0.3,
				hair.HighHairDensity.Type:   0.2,
			},
		),
	),
	hair.NewHairColorGene(map[string]float64{
		color.BlackHairColorPalette: 1,
	}),
)

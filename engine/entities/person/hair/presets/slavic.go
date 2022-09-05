package presets

import "persons_generator/engine/entities/person/hair"

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
		hair.LightAshBlondeHairColor.Name:    0.25,
		hair.LightBlondeHairColor.Name:       0.25,
		hair.LightGoldenBlondeHairColor.Name: 0.25,
		hair.BeelineHoneyHairColor.Name:      0.25,
		hair.MediumChampagneHairColor.Name:   0.25,
		hair.ButterscotchHairColor.Name:      0.25,
		hair.LightCoolBrownHairColor.Name:    0.25,
		hair.LightBrownHairColor.Name:        0.25,
		hair.LightGoldenBrownHairColor.Name:  0.25,
		hair.ChocolateBrownHairColor.Name:    0.25,
		hair.DarkGoldenBrownHairColor.Name:   0.25,
		hair.MediumAshBrownHairColor.Name:    0.25,
		hair.ReddishBlondeHairColor.Name:     0.1,
		hair.LightAuburnHairColor.Name:       0.1,
		hair.MediumAuburnHairColor.Name:      0.1,
		hair.RedHotCinnamonHairColor.Name:    0.1,
		hair.ExpressoHairColor.Name:          0,
		hair.JetBlackHairColor.Name:          0,
	}),
)

package presets

import "persons_generator/engine/entities/person/hair"

var SlavicHairPreset = hair.NewHairGene(
	hair.NewScalpHairGene(
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
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type):    0.001,
			string(hair.CurlyHairTexture.Type):    0.15,
			string(hair.FineHairTexture.Type):     0.25,
			string(hair.StraightHairTexture.Type): 0.25,
			string(hair.ThickHairTexture.Type):    0.15,
			string(hair.ThinHairTexture.Type):     0.15,
			string(hair.WavyHairTexture.Type):     0.25,
		}),
	),
	hair.NewFaceHairGene(
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
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type):    0.001,
			string(hair.CurlyHairTexture.Type):    0.15,
			string(hair.FineHairTexture.Type):     0.25,
			string(hair.StraightHairTexture.Type): 0.25,
			string(hair.ThickHairTexture.Type):    0.15,
			string(hair.ThinHairTexture.Type):     0.15,
			string(hair.WavyHairTexture.Type):     0.25,
		}),
	),
	hair.NewAxillaryHairGene(
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
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type):    0.001,
			string(hair.CurlyHairTexture.Type):    0.15,
			string(hair.FineHairTexture.Type):     0.25,
			string(hair.StraightHairTexture.Type): 0.25,
			string(hair.ThickHairTexture.Type):    0.15,
			string(hair.ThinHairTexture.Type):     0.15,
			string(hair.WavyHairTexture.Type):     0.25,
		}),
	),
	hair.NewChestAndAbdomenHairGene(
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
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type):    0.001,
			string(hair.CurlyHairTexture.Type):    0.15,
			string(hair.FineHairTexture.Type):     0.25,
			string(hair.StraightHairTexture.Type): 0.25,
			string(hair.ThickHairTexture.Type):    0.15,
			string(hair.ThinHairTexture.Type):     0.15,
			string(hair.WavyHairTexture.Type):     0.25,
		}),
	),
	hair.NewArmsHairGene(
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
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type):    0.001,
			string(hair.CurlyHairTexture.Type):    0.15,
			string(hair.FineHairTexture.Type):     0.25,
			string(hair.StraightHairTexture.Type): 0.25,
			string(hair.ThickHairTexture.Type):    0.15,
			string(hair.ThinHairTexture.Type):     0.15,
			string(hair.WavyHairTexture.Type):     0.25,
		}),
	),
	hair.NewPubicHairGene(
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
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type):    0.001,
			string(hair.CurlyHairTexture.Type):    0.15,
			string(hair.FineHairTexture.Type):     0.25,
			string(hair.StraightHairTexture.Type): 0.25,
			string(hair.ThickHairTexture.Type):    0.15,
			string(hair.ThinHairTexture.Type):     0.15,
			string(hair.WavyHairTexture.Type):     0.25,
		}),
	),
	hair.NewLegsHairGene(
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
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type):    0.001,
			string(hair.CurlyHairTexture.Type):    0.15,
			string(hair.FineHairTexture.Type):     0.25,
			string(hair.StraightHairTexture.Type): 0.25,
			string(hair.ThickHairTexture.Type):    0.15,
			string(hair.ThinHairTexture.Type):     0.15,
			string(hair.WavyHairTexture.Type):     0.25,
		}),
	),
	hair.NewFeetHairGene(
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
		hair.NewHairTextureGene(map[string]float64{
			string(hair.BlackHairTexture.Type):    0.001,
			string(hair.CurlyHairTexture.Type):    0.15,
			string(hair.FineHairTexture.Type):     0.25,
			string(hair.StraightHairTexture.Type): 0.25,
			string(hair.ThickHairTexture.Type):    0.15,
			string(hair.ThinHairTexture.Type):     0.15,
			string(hair.WavyHairTexture.Type):     0.25,
		}),
	),
)

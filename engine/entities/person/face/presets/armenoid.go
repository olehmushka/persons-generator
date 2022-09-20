package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var ArmenoidFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):      0.2,
		string(face.Oblong):    0.15,
		string(face.VTriangle): 0.25,
		string(face.Square):    0.1,
		string(face.Diamond):   0.15,
		string(face.Rectangle): 0.05,
		string(face.Heart):     0.15,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.RoundEyesType):          0.01,
			string(face.RoundishAlmondEyesType): 0.05,
			string(face.AlmondEyesType):         0.1,
			string(face.DroopyEyesType):         0.25,
			string(face.DroopyHoodedEyesType):   0.15,
			string(face.HoodedEyesType):         0.1,
		}),
		face.NewEyesColorGene(map[string]float64{
			color.BrownEyesColorPalette:    0.75,
			color.HazelEyesColorPalette:    0.15,
			color.AmberEyesColorPalette:    0.15,
			color.GreenEyesColorPalette:    0.05,
			color.GreyEyesColorPalette:     0.075,
			color.TealBlueEyesColorPalette: 0.05,
			color.CyanEyesColorPalette:     0.025,
			color.IndigoEyesColorPalette:   0.01,
		}),
	),
	face.NewEarsGene(map[string]float64{
		string(face.SquareEarType):        0.25,
		string(face.PointedEarType):       0.25,
		string(face.NarrowEarType):        0.25,
		string(face.StickingEarType):      0.25,
		string(face.RoundEarFreeLobeType): 0.25,
		string(face.AttachedLobeType):     0.25,
		string(face.BroadLobeType):        0.25,
	}),
	face.NewNoseGene(map[string]float64{
		string(face.FleshyNoseType):   0.2,
		string(face.TurnedUpNoseType): 0.01,
		string(face.HawkNoseType):     0.35,
		string(face.GreekNoseType):    0.35,
		string(face.RomanNoseType):    0.15,
		string(face.BumpyNoseType):    0.15,
		string(face.NixonNoseType):    0.225,
		string(face.BulbousNoseType):  0.223,
		string(face.ComboNoseType):    0.2,
		string(face.SnubNoseType):     0.05,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.FullLipsType):           0.05,
		string(face.HeavyUpperLipsType):     0.05,
		string(face.WideLipsType):           0.05,
		string(face.HeavyLowerLipsType):     0.05,
		string(face.ThinLipsType):           0.2,
		string(face.HeardShapedLipsType):    0.35,
		string(face.DownwardTurnedLipsType): 0.2,
	}),
)

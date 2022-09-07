package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var SlavicFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):      0.3,
		string(face.Circle):    0.4,
		string(face.Oblong):    0.1,
		string(face.VTriangle): 0.08,
		string(face.Square):    0.09,
		string(face.Diamond):   0.01,
		string(face.Rectangle): 0.03,
		string(face.Heart):     0.03,
		string(face.ATriangle): 0.03,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.RoundEyesType):              0.5,
			string(face.RoundishAlmondEyesType):     0.4,
			string(face.AlmondEyesType):             0.3,
			string(face.DroopyEyesType):             0.05,
			string(face.DroopyHoodedEyesType):       0.05,
			string(face.HoodedEyesType):             0.05,
			string(face.AsianEyesType):              0.01,
			string(face.ChildishRoundAsianEyesType): 0.01,
		}),
		face.NewEyesColorGene(map[string]float64{
			color.BrownEyesColorPalette:    0.1,
			color.HazelEyesColorPalette:    0.25,
			color.AmberEyesColorPalette:    0.25,
			color.GreenEyesColorPalette:    0.15,
			color.GreyEyesColorPalette:     0.15,
			color.TealBlueEyesColorPalette: 0.15,
			color.CyanEyesColorPalette:     0.1,
			color.IndigoEyesColorPalette:   0.1,
		}),
	),
	face.NewEarsGene(map[string]float64{
		string(face.SquareEarType):        0.25,
		string(face.PointedEarType):       0.1,
		string(face.NarrowEarType):        0.1,
		string(face.StickingEarType):      0.25,
		string(face.RoundEarFreeLobeType): 0.25,
		string(face.AttachedLobeType):     0.25,
		string(face.BroadLobeType):        0.25,
	}),
	face.NewNoseGene(map[string]float64{
		string(face.FleshyNoseType):    0.25,
		string(face.TurnedUpNoseType):  0.2,
		string(face.HawkNoseType):      0.1,
		string(face.GreekNoseType):     0.03,
		string(face.RomanNoseType):     0.3,
		string(face.BumpyNoseType):     0.25,
		string(face.NixonNoseType):     0.1,
		string(face.BulbousNoseType):   0.15,
		string(face.ComboNoseType):     0.01,
		string(face.SnubNoseType):      0.05,
		string(face.NubianNoseType):    0.01,
		string(face.EastAsianNoseType): 0.01,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.FullLipsType):           0.2,
		string(face.HeavyUpperLipsType):     0.2,
		string(face.WideLipsType):           0.2,
		string(face.HeavyLowerLipsType):     0.25,
		string(face.ThinLipsType):           0.2,
		string(face.HeardShapedLipsType):    0.2,
		string(face.DownwardTurnedLipsType): 0.25,
	}),
)

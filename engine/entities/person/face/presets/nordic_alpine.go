package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var NordicAlpineFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):      0,
		string(face.Circle):    0,
		string(face.Oblong):    0,
		string(face.VTriangle): 0,
		string(face.Square):    0,
		string(face.Diamond):   0,
		string(face.Rectangle): 0,
		string(face.Heart):     0,
		string(face.ATriangle): 0,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.RoundEyesType):              0,
			string(face.RoundishAlmondEyesType):     0,
			string(face.AlmondEyesType):             0,
			string(face.DroopyEyesType):             0,
			string(face.DroopyHoodedEyesType):       0,
			string(face.HoodedEyesType):             0,
			string(face.AsianEyesType):              0,
			string(face.ChildishRoundAsianEyesType): 0,
		}),
		face.NewEyesColorGene(map[string]float64{
			color.BrownEyesColorPalette:    0,
			color.HazelEyesColorPalette:    0,
			color.AmberEyesColorPalette:    0,
			color.GreenEyesColorPalette:    0,
			color.GreyEyesColorPalette:     0,
			color.TealBlueEyesColorPalette: 0,
			color.CyanEyesColorPalette:     0,
			color.IndigoEyesColorPalette:   0,
		}),
	),
	face.NewEarsGene(map[string]float64{
		string(face.SquareEarType):        0,
		string(face.PointedEarType):       0,
		string(face.NarrowEarType):        0,
		string(face.StickingEarType):      0,
		string(face.RoundEarFreeLobeType): 0,
		string(face.AttachedLobeType):     0,
		string(face.BroadLobeType):        0,
	}),
	face.NewNoseGene(map[string]float64{
		string(face.FleshyNoseType):    0,
		string(face.TurnedUpNoseType):  0,
		string(face.HawkNoseType):      0,
		string(face.GreekNoseType):     0,
		string(face.RomanNoseType):     0,
		string(face.BumpyNoseType):     0,
		string(face.NixonNoseType):     0,
		string(face.BulbousNoseType):   0,
		string(face.ComboNoseType):     0,
		string(face.SnubNoseType):      0,
		string(face.NubianNoseType):    0,
		string(face.EastAsianNoseType): 0,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.FullLipsType):           0,
		string(face.HeavyUpperLipsType):     0,
		string(face.WideLipsType):           0,
		string(face.HeavyLowerLipsType):     0,
		string(face.ThinLipsType):           0,
		string(face.HeardShapedLipsType):    0,
		string(face.DownwardTurnedLipsType): 0,
	}),
)

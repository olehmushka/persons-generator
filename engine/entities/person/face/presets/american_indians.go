package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var AmericanIndiansFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):      0.2,
		string(face.Circle):    0.15,
		string(face.Oblong):    0.1,
		string(face.VTriangle): 0.2,
		string(face.Diamond):   0.3,
		string(face.Rectangle): 0.2,
		string(face.Heart):     0.25,
		string(face.ATriangle): 0.2,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.DroopyEyesType):             0.1,
			string(face.DroopyHoodedEyesType):       0.15,
			string(face.HoodedEyesType):             0.25,
			string(face.AsianEyesType):              0.4,
			string(face.ChildishRoundAsianEyesType): 0.4,
		}),
		face.NewEyesColorGene(map[string]float64{
			color.BrownEyesColorPalette:  0.9,
			color.HazelEyesColorPalette:  0.05,
			color.AmberEyesColorPalette:  0.01,
			color.IndigoEyesColorPalette: 0.05,
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
		string(face.FleshyNoseType):    0.2,
		string(face.TurnedUpNoseType):  0.01,
		string(face.HawkNoseType):      0.35,
		string(face.GreekNoseType):     0.01,
		string(face.RomanNoseType):     0.4,
		string(face.BumpyNoseType):     0.2,
		string(face.NixonNoseType):     0.15,
		string(face.BulbousNoseType):   0.1,
		string(face.ComboNoseType):     0.2,
		string(face.SnubNoseType):      0.05,
		string(face.EastAsianNoseType): 0.2,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.FullLipsType):           0.1,
		string(face.HeavyUpperLipsType):     0.12,
		string(face.WideLipsType):           0.12,
		string(face.HeavyLowerLipsType):     0.15,
		string(face.ThinLipsType):           0.2,
		string(face.HeardShapedLipsType):    0.3,
		string(face.DownwardTurnedLipsType): 0.1,
	}),
)

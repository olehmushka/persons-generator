package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var IndonesianMongoloidFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):      0.1,
		string(face.Oblong):    0.05,
		string(face.VTriangle): 0.35,
		string(face.Diamond):   0.1,
		string(face.Rectangle): 0.05,
		string(face.Heart):     0.25,
		string(face.ATriangle): 0.01,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.HoodedEyesType):             0.15,
			string(face.AsianEyesType):              0.45,
			string(face.ChildishRoundAsianEyesType): 0.35,
		}),
		face.NewEyesColorGene(map[string]float64{
			color.BrownEyesColorPalette: 1,
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
		string(face.FleshyNoseType):    0.25,
		string(face.TurnedUpNoseType):  0.05,
		string(face.BumpyNoseType):     0.1,
		string(face.NixonNoseType):     0.05,
		string(face.BulbousNoseType):   0.1,
		string(face.ComboNoseType):     0.15,
		string(face.SnubNoseType):      0.05,
		string(face.EastAsianNoseType): 0.2,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.FullLipsType):           0.01,
		string(face.HeavyUpperLipsType):     0.05,
		string(face.WideLipsType):           0.1,
		string(face.HeavyLowerLipsType):     0.05,
		string(face.ThinLipsType):           0.1,
		string(face.HeardShapedLipsType):    0.3,
		string(face.DownwardTurnedLipsType): 0.25,
	}),
)

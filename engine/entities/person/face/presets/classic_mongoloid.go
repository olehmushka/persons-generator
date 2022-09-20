package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var ClassicMongoloidFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):   0.25,
		string(face.Circle): 0.35,
		string(face.Oblong): 0.2,
		string(face.Square): 0.2,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.HoodedEyesType):             0.1,
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
		string(face.FleshyNoseType):    0.2,
		string(face.TurnedUpNoseType):  0.01,
		string(face.HawkNoseType):      0.03,
		string(face.RomanNoseType):     0.1,
		string(face.BumpyNoseType):     0.2,
		string(face.NixonNoseType):     0.05,
		string(face.BulbousNoseType):   0.1,
		string(face.ComboNoseType):     0.1,
		string(face.SnubNoseType):      0.05,
		string(face.EastAsianNoseType): 0.35,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.FullLipsType):           0.1,
		string(face.WideLipsType):           0.2,
		string(face.ThinLipsType):           0.2,
		string(face.HeardShapedLipsType):    0,
		string(face.DownwardTurnedLipsType): 0.15,
	}),
)

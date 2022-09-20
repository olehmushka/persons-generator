package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var ArcticMongoloidFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):      0.25,
		string(face.Circle):    0.3,
		string(face.Oblong):    0.25,
		string(face.Square):    0.3,
		string(face.Rectangle): 0.1,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.HoodedEyesType):             0.3,
			string(face.AsianEyesType):              0.4,
			string(face.ChildishRoundAsianEyesType): 0.3,
		}),
		face.NewEyesColorGene(map[string]float64{
			color.BrownEyesColorPalette: 0.98,
			color.HazelEyesColorPalette: 0.01,
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
		string(face.FleshyNoseType):    0.1,
		string(face.HawkNoseType):      0.25,
		string(face.RomanNoseType):     0.15,
		string(face.BumpyNoseType):     0.1,
		string(face.NixonNoseType):     0.3,
		string(face.BulbousNoseType):   0.25,
		string(face.ComboNoseType):     0.2,
		string(face.SnubNoseType):      0.05,
		string(face.EastAsianNoseType): 0.1,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.HeavyUpperLipsType):  0.15,
		string(face.WideLipsType):        0.15,
		string(face.HeavyLowerLipsType):  0.15,
		string(face.ThinLipsType):        0.15,
		string(face.HeardShapedLipsType): 0.2,
	}),
)

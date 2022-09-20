package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var AfricanNegroFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):      0.3,
		string(face.Circle):    0.3,
		string(face.Oblong):    0.25,
		string(face.VTriangle): 0.01,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.RoundEyesType):          0.4,
			string(face.RoundishAlmondEyesType): 0.4,
			string(face.AlmondEyesType):         0.4,
			string(face.DroopyEyesType):         0.4,
			string(face.DroopyHoodedEyesType):   0.2,
			string(face.HoodedEyesType):         0.2,
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
		string(face.BumpyNoseType):  0.1,
		string(face.ComboNoseType):  0.1,
		string(face.SnubNoseType):   0.05,
		string(face.NubianNoseType): 0.5,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.FullLipsType):        0.8,
		string(face.HeavyUpperLipsType):  0.2,
		string(face.WideLipsType):        0.2,
		string(face.HeavyLowerLipsType):  0.1,
		string(face.HeardShapedLipsType): 0.2,
	}),
)

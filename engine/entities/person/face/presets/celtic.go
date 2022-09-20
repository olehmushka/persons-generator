package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var CelticFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):      0.1,
		string(face.Circle):    0.15,
		string(face.Oblong):    0.22,
		string(face.VTriangle): 0.15,
		string(face.Square):    0.2,
		string(face.Diamond):   0.2,
		string(face.Rectangle): 0.2,
		string(face.Heart):     0.17,
		string(face.ATriangle): 0.15,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.RoundEyesType):          0.1,
			string(face.RoundishAlmondEyesType): 0.15,
			string(face.AlmondEyesType):         0.2,
			string(face.DroopyEyesType):         0.25,
			string(face.DroopyHoodedEyesType):   0.15,
		}),
		face.NewEyesColorGene(map[string]float64{
			color.BrownEyesColorPalette:    0.2,
			color.HazelEyesColorPalette:    0.2,
			color.AmberEyesColorPalette:    0.2,
			color.GreenEyesColorPalette:    0.3,
			color.GreyEyesColorPalette:     0.3,
			color.TealBlueEyesColorPalette: 0.2,
			color.CyanEyesColorPalette:     0.2,
			color.IndigoEyesColorPalette:   0.2,
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
		string(face.FleshyNoseType):   0.1,
		string(face.TurnedUpNoseType): 0.2,
		string(face.HawkNoseType):     0.1,
		string(face.RomanNoseType):    0.35,
		string(face.BumpyNoseType):    0.1,
		string(face.NixonNoseType):    0.2,
		string(face.BulbousNoseType):  0.05,
		string(face.ComboNoseType):    0.1,
		string(face.SnubNoseType):     0.05,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.HeavyUpperLipsType):     0.05,
		string(face.WideLipsType):           0.05,
		string(face.HeavyLowerLipsType):     0.05,
		string(face.ThinLipsType):           0.15,
		string(face.HeardShapedLipsType):    0.2,
		string(face.DownwardTurnedLipsType): 0.2,
	}),
)

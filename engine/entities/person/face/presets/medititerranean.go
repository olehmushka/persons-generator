package presets

import (
	"persons_generator/engine/entities/person/color"
	"persons_generator/engine/entities/person/face"
)

var MedititerraneanFacePreset = face.NewFaceGene(
	face.NewFaceShapeGene(map[string]float64{
		string(face.Oval):      0.25,
		string(face.Circle):    0.05,
		string(face.Oblong):    0.2,
		string(face.VTriangle): 0.05,
		string(face.Rectangle): 0.2,
		string(face.Heart):     0.2,
	}),
	face.NewEyesGene(
		face.NewEyesTypeGene(map[string]float64{
			string(face.RoundEyesType):          0.1,
			string(face.RoundishAlmondEyesType): 0.15,
			string(face.AlmondEyesType):         0.35,
			string(face.DroopyEyesType):         0.2,
		}),
		face.NewEyesColorGene(map[string]float64{
			color.BrownEyesColorPalette:    0.6,
			color.HazelEyesColorPalette:    0.25,
			color.AmberEyesColorPalette:    0.25,
			color.GreenEyesColorPalette:    0.05,
			color.GreyEyesColorPalette:     0.1,
			color.TealBlueEyesColorPalette: 0.05,
			color.CyanEyesColorPalette:     0.05,
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
		string(face.FleshyNoseType):   0.1,
		string(face.TurnedUpNoseType): 0.1,
		string(face.HawkNoseType):     0.1,
		string(face.GreekNoseType):    0.15,
		string(face.RomanNoseType):    0.35,
		string(face.BumpyNoseType):    0.1,
		string(face.NixonNoseType):    0.12,
		string(face.BulbousNoseType):  0.1,
		string(face.ComboNoseType):    0.15,
		string(face.SnubNoseType):     0.05,
		string(face.NubianNoseType):   0.01,
	}),
	face.NewMouthGene(map[string]float64{
		string(face.FullLipsType):           0.05,
		string(face.HeavyUpperLipsType):     0.1,
		string(face.WideLipsType):           0.01,
		string(face.HeavyLowerLipsType):     0.1,
		string(face.ThinLipsType):           0.25,
		string(face.HeardShapedLipsType):    0.25,
		string(face.DownwardTurnedLipsType): 0.25,
	}),
)

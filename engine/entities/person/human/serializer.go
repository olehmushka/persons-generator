package human

import "persons_generator/engine/entities/person/color"

type SerializedHuman struct {
	Sex string `json:"sex" bson:"sex"`
	Age int    `json:"age" bson:"age"`

	FaceShape         string      `json:"face_shape" bson:"face_shape"`
	EyesColor         color.Color `json:"eyes_color" bson:"eyes_color"`
	EyesType          string      `json:"eyes_type" bson:"eyes_type"`
	EarsType          string      `json:"ears_type" bson:"ears_type"`
	NoseType          string      `json:"mose_type" bson:"mose_type"`
	LipsType          string      `json:"lips_type" bson:"lips_type"`
	HairColor         color.Color `json:"hair_color" bson:"hair_color"`
	ScalpHairTexture  string      `json:"scalp_hair_texture" bson:"scalp_hair_texture"`
	ScalpHairDensity  string      `json:"scalp_hair_density" bson:"scalp_hair_density"`
	FaceHairDensity   string      `json:"face_hair_density" bson:"face_hair_density"`
	HeightInCm        float64     `json:"height_in_cm" bson:"height_in_cm"`
	ShoeEUSize        int         `json:"shoe_eu_size" bson:"shoe_eu_size"`
	SkinColor         color.Color `json:"skin_color" bson:"skin_color"`
	SexualOrientation string      `json:"sexual_orientation" bson:"sexual_orientation"`
	Temperament       string      `json:"temperament" bson:"temperament"`
}

func (h *Human) Serialize() *SerializedHuman {
	if h == nil {
		return nil
	}

	return &SerializedHuman{
		Sex:               h.Sex.String(),
		Age:               h.Age,
		FaceShape:         h.Body.Face.Shape.String(),
		EyesColor:         h.Body.Face.Eyes.Color.Value(),
		EyesType:          h.Body.Face.Eyes.Type.String(),
		EarsType:          h.Body.Face.Ears.Type.String(),
		NoseType:          h.Body.Face.Nose.Type.String(),
		LipsType:          h.Body.Face.Mouth.LipsType.String(),
		HairColor:         h.Body.Hair.Color.Value(),
		ScalpHairTexture:  h.Body.Hair.Scalp.Texture.Serialize(),
		ScalpHairDensity:  h.Body.Hair.Scalp.Density.Serialize(),
		FaceHairDensity:   h.Body.Hair.Face.Density.Serialize(),
		HeightInCm:        h.Body.Size.Height.HeightInCm,
		ShoeEUSize:        h.Body.Size.ShoeSize.EUSize,
		SkinColor:         h.Body.Skin.Color.Value(),
		SexualOrientation: h.Psycho.SexualOrientation.String(),
		Temperament:       h.Psycho.Temperament.Serialize(),
	}
}

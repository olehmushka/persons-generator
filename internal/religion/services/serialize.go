package services

import (
	g "persons_generator/engine/entities/gender"
	engineReligion "persons_generator/engine/entities/religion"
	"persons_generator/internal/religion/entities"
)

func serializeReligions(in []*engineReligion.Religion) []*entities.Religion {
	out := make([]*entities.Religion, len(in))
	for i := range out {
		out[i] = serializeReligion(in[i])
	}

	return out
}

func serializeReligion(in *engineReligion.Religion) *entities.Religion {
	if in == nil {
		return nil
	}

	return &entities.Religion{
		ID:                  in.ID,
		Name:                in.Name,
		Type:                serializeType(in.Type),
		DominantSex:         serializeDominantSex(in.GenderDominance),
		HighGoals:           serializeHighGoals(in.Doctrine.HighGoal),
		DeityGoodness:       serializeDeityGoodness(in.Doctrine.Deity),
		DeityTraits:         serializeDeityTraits(in.Doctrine.Deity),
		HumanGoodness:       serializeHumanGoodness(in.Doctrine.Human),
		HumanTraits:         serializeHumanTraits(in.Doctrine.Human),
		SocialTraits:        serializeSocialTraits(in.Doctrine.Social),
		SourceOfMoralLaw:    in.Doctrine.SourceOfMoralLaw.String(),
		Afterlife:           serializeAfterlife(in.Doctrine.Afterlife),
		Attributes:          serializeAttributes(in.Attributes),
		Clerics:             serializeClerics(in.Attributes),
		HolyScriptureTraits: serializeHolyScriptureTraits(in.Attributes),
		TempleTraits:        serializeTempleTraits(in.Attributes),
		TheologyTraits:      serializeTheologyTraits(in.Theology),
		Cults:               serializeCults(in.Theology),
		Rules:               serializeRules(in.Theology),
		Taboos:              serializeTaboos(in.Theology),
		Rituals:             serializeRituals(in.Theology),
		Holydays:            serializeHolydays(in.Theology),
		ConversionTraits:    serializeConversionTraits(in.Theology),
		MarriageTradition:   serializeMarriageTradition(in.Theology),
	}
}

func serializeType(in *engineReligion.Type) entities.Type {
	if in == nil {
		return entities.Type{}
	}

	return entities.Type{
		Name:        in.Type.String(),
		SubtypeName: in.Subtype.String(),
	}
}

func serializeDominantSex(in *g.Dominance) string {
	if in == nil {
		return ""
	}

	return in.Dominance.String()
}

func serializeHighGoals(in *engineReligion.HighGoal) []string {
	if in == nil {
		return []string{}
	}
	out := make([]string, len(in.Goals))
	for i := range out {
		out[i] = in.Goals[i].Name
	}

	return out
}

func serializeDeityGoodness(in *engineReligion.DeityDoctrine) string {
	if in == nil || in.Nature == nil {
		return ""
	}

	return in.Nature.Goodness.Goodness.String()
}

func serializeDeityTraits(in *engineReligion.DeityDoctrine) []string {
	if in == nil || in.Nature == nil {
		return []string{}
	}
	out := make([]string, len(in.Nature.Traits))
	for i := range out {
		out[i] = in.Nature.Traits[i].Name
	}

	return out
}

func serializeHumanGoodness(in *engineReligion.HumanDoctrine) string {
	if in == nil || in.Nature == nil {
		return ""
	}

	return in.Nature.Goodness.Goodness.String()
}

func serializeHumanTraits(in *engineReligion.HumanDoctrine) []string {
	if in == nil || in.Nature == nil {
		return []string{}
	}
	out := make([]string, len(in.Nature.Traits))
	for i := range out {
		out[i] = in.Nature.Traits[i].Name
	}

	return out
}

func serializeSocialTraits(in *engineReligion.SocialDoctrine) []string {
	if in == nil || len(in.Traits) == 0 {
		return []string{}
	}
	out := make([]string, len(in.Traits))
	for i := range out {
		out[i] = in.Traits[i].Name
	}

	return out
}

func serializeAfterlife(in *engineReligion.Afterlife) *entities.Afterlife {
	if in == nil || !in.IsExists {
		return nil
	}
	out := &entities.Afterlife{
		Participants: entities.AfterlifeParticipants{
			ForTopBelievers:    in.Participants.ForTopBelievers.String(),
			ForBelievers:       in.Participants.ForBelievers.String(),
			ForUntrueBelievers: in.Participants.ForUntrueBelievers.String(),
			ForAtheists:        in.Participants.ForAtheists.String(),
		},
	}
	traits := make([]string, len(in.Traits))
	for i := range traits {
		traits[i] = in.Traits[i].Name
	}
	out.Traits = traits

	return out
}

func serializeAttributes(in *engineReligion.Attributes) []string {
	if in == nil {
		return []string{}
	}
	out := make([]string, len(in.Traits))
	for i := range out {
		out[i] = in.Traits[i].Name
	}

	return out
}

func serializeClerics(in *engineReligion.Attributes) *entities.Clerics {
	if in == nil || !in.Clerics.HasClerics {
		return nil
	}
	out := &entities.Clerics{
		IsCivil:                  in.Clerics.Appointment.IsCivil,
		IsRevocable:              in.Clerics.Appointment.IsRevocable,
		AcceptableClericSex:      in.Clerics.Limitations.AcceptableGender.String(),
		AcceptableClericMarriage: in.Clerics.Limitations.Marriage.String(),
	}
	traits := make([]string, len(in.Clerics.Traits))
	for i := range traits {
		traits[i] = in.Clerics.Traits[i].Name
	}
	out.Traits = traits
	functions := make([]string, len(in.Clerics.Functions))
	for i := range functions {
		functions[i] = in.Clerics.Functions[i].Name
	}
	out.Functions = functions

	return out
}

func serializeHolyScriptureTraits(in *engineReligion.Attributes) []string {
	if in == nil || !in.HolyScripture.HasHolyScripture {
		return []string{}
	}
	out := make([]string, len(in.HolyScripture.Traits))
	for i := range out {
		out[i] = in.HolyScripture.Traits[i].Name
	}

	return out
}

func serializeTempleTraits(in *engineReligion.Attributes) []string {
	if in == nil || len(in.Temples.Traits) == 0 {
		return []string{}
	}
	out := make([]string, len(in.Temples.Traits))
	for i := range out {
		out[i] = in.Temples.Traits[i].Name
	}

	return out
}

func serializeTheologyTraits(in *engineReligion.Theology) []string {
	if in == nil || len(in.Traits) == 0 {
		return []string{}
	}
	out := make([]string, len(in.Traits))
	for i := range out {
		out[i] = in.Traits[i].Name
	}

	return out
}

func serializeCults(in *engineReligion.Theology) []string {
	if in == nil || len(in.Cults) == 0 {
		return []string{}
	}
	out := make([]string, len(in.Cults))
	for i := range out {
		out[i] = in.Cults[i].Name
	}

	return out
}

func serializeRules(in *engineReligion.Theology) []string {
	if in == nil || len(in.Rules.Rules) == 0 {
		return []string{}
	}
	out := make([]string, len(in.Rules.Rules))
	for i := range out {
		out[i] = in.Rules.Rules[i].Name
	}

	return out
}

func serializeTaboos(in *engineReligion.Theology) []entities.Taboo {
	if in == nil || len(in.Taboos.Taboos) == 0 {
		return []entities.Taboo{}
	}
	out := make([]entities.Taboo, len(in.Taboos.Taboos))
	for i := range out {
		out[i] = entities.Taboo{
			Action:     in.Taboos.Taboos[i].Name,
			Acceptance: in.Taboos.Taboos[i].Acceptance.String(),
		}
	}

	return out
}

func serializeRituals(in *engineReligion.Theology) []string {
	if in == nil {
		return []string{}
	}
	out := make([]string, 0, len(in.Rituals.Funeral)+len(in.Rituals.Holyday)+len(in.Rituals.Initiation)+len(in.Rituals.Sacrifice))
	for _, r := range in.Rituals.Funeral {
		out = append(out, r.Name)
	}
	for _, r := range in.Rituals.Holyday {
		out = append(out, r.Name)
	}
	for _, r := range in.Rituals.Initiation {
		out = append(out, r.Name)
	}
	for _, r := range in.Rituals.Sacrifice {
		out = append(out, r.Name)
	}

	return out
}

func serializeHolydays(in *engineReligion.Theology) []string {
	if in == nil || in.Holydays == nil {
		return []string{}
	}
	out := make([]string, len(in.Holydays.Holydays))
	for i := range out {
		out[i] = in.Holydays.Holydays[i].Name
	}

	return out
}

func serializeConversionTraits(in *engineReligion.Theology) []string {
	if in == nil || in.Conversion == nil {
		return []string{}
	}
	out := make([]string, len(in.Conversion.Traits))
	for i := range out {
		out[i] = in.Conversion.Traits[i].Name
	}

	return out
}

func serializeMarriageTradition(in *engineReligion.Theology) entities.MarriageTradition {
	if in == nil || in.MarriageTradition == nil {
		return entities.MarriageTradition{}
	}

	return entities.MarriageTradition{
		Kind:          in.MarriageTradition.Kind.String(),
		Bastardry:     in.MarriageTradition.Bastardry.String(),
		Consanguinity: in.MarriageTradition.Consanguinity.String(),
		Divorce:       in.MarriageTradition.Divorce.String(),
	}
}

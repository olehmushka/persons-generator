package http

import (
	"persons_generator/internal/culture/entities"
	personsEntities "persons_generator/internal/persons/entities"
	religionEntities "persons_generator/internal/religion/entities"
)

func serializeCultures(cultures []*entities.Culture) []*Culture {
	out := make([]*Culture, len(cultures))
	for i := range out {
		out[i] = serializeCulture(cultures[i])
	}

	return out
}

func serializeCulture(in *entities.Culture) *Culture {
	return &Culture{
		ID:              in.ID,
		Name:            in.Name,
		Proto:           serializeCultures(in.Proto),
		CultureGroup:    serializeCultureGroup(in.CultureGroup),
		RootCultureName: in.RootCultureName,
		Language: &Language{
			Name:      in.Language.Name,
			Subfamily: serializeLanguageSubfamily(in.Language.Subfamily),
		},
		EthosName:     in.EthosName,
		Traditions:    in.Traditions,
		DominantSex:   in.DominantSex,
		MartialCustom: in.MartialCustom,
	}
}

func serializeCultureGroup(in *entities.CultureGroup) *CultureGroup {
	if in == nil {
		return nil
	}

	return &CultureGroup{
		Name:            in.Name,
		RootCultureName: in.RootCultureName,
	}
}

func serializeLanguageSubfamily(in *entities.LanguageSubfamily) *LanguageSubfamily {
	if in == nil {
		return nil
	}
	return &LanguageSubfamily{
		Name:              in.Name,
		FamilyName:        in.FamilyName,
		ExtendedSubfamily: serializeLanguageSubfamily(in.ExtendedSubfamily),
	}
}

func serializeReligions(religions []*religionEntities.Religion) []*Religion {
	out := make([]*Religion, len(religions))
	for i := range out {
		out[i] = serializeReligion(religions[i])
	}

	return out
}

func serializeReligion(in *religionEntities.Religion) *Religion {
	return &Religion{
		ID:                  in.ID,
		Name:                in.Name,
		Type:                serializeReligionType(in.Type),
		DominantSex:         in.DominantSex,
		HighGoals:           in.HighGoals,
		DeityGoodness:       in.DeityGoodness,
		DeityTraits:         in.DeityTraits,
		HumanGoodness:       in.HumanGoodness,
		HumanTraits:         in.HumanTraits,
		SocialTraits:        in.SocialTraits,
		SourceOfMoralLaw:    in.SourceOfMoralLaw,
		Afterlife:           serializeAfterlife(in.Afterlife),
		Attributes:          in.Attributes,
		Clerics:             serializeClerics(in.Clerics),
		HolyScriptureTraits: in.HolyScriptureTraits,
		TempleTraits:        in.TempleTraits,
		TheologyTraits:      in.TheologyTraits,
		Cults:               in.Cults,
		Rules:               in.Rules,
		Taboos:              serializeTaboos(in.Taboos),
		Rituals:             in.Rituals,
		Holydays:            in.Holydays,
		ConversionTraits:    in.ConversionTraits,
		MarriageTradition:   serializeMarriageTradition(in.MarriageTradition),
	}
}

func serializeReligionType(in religionEntities.Type) ReligionType {
	return ReligionType{
		Name:        in.Name,
		SubtypeName: in.SubtypeName,
	}
}

func serializeAfterlife(in *religionEntities.Afterlife) *Afterlife {
	if in == nil {
		return nil
	}

	return &Afterlife{
		Participants: AfterlifeParticipants{
			ForTopBelievers:    in.Participants.ForTopBelievers,
			ForBelievers:       in.Participants.ForBelievers,
			ForUntrueBelievers: in.Participants.ForUntrueBelievers,
			ForAtheists:        in.Participants.ForAtheists,
		},
		Traits: in.Traits,
	}
}

func serializeClerics(in *religionEntities.Clerics) *Clerics {
	if in == nil {
		return nil
	}

	return &Clerics{
		IsCivil:                  in.IsCivil,
		IsRevocable:              in.IsRevocable,
		AcceptableClericSex:      in.AcceptableClericSex,
		AcceptableClericMarriage: in.AcceptableClericMarriage,
		Traits:                   in.Traits,
		Functions:                in.Functions,
	}
}

func serializeTaboos(in []religionEntities.Taboo) []Taboo {
	if in == nil {
		return []Taboo{}
	}
	out := make([]Taboo, len(in))
	for i := range out {
		out[i] = Taboo{
			Action:     in[i].Action,
			Acceptance: in[i].Acceptance,
		}
	}

	return out
}

func serializeMarriageTradition(in religionEntities.MarriageTradition) MarriageTradition {
	return MarriageTradition{
		Kind:          in.Kind,
		Bastardry:     in.Bastardry,
		Consanguinity: in.Consanguinity,
		Divorce:       in.Divorce,
	}
}

func serializePersons(in []*personsEntities.Person) []*Person {
	out := make([]*Person, len(in))
	for i := range out {
		out[i] = serializePerson(in[i])
	}

	return out
}

func serializePerson(in *personsEntities.Person) *Person {
	return &Person{}
}

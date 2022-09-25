package http

import (
	"encoding/json"
	"persons_generator/core/wrapped_error"
	"persons_generator/engine/entities/culture"
	"persons_generator/engine/entities/religion"
	"persons_generator/internal/culture/entities"
	personsEntities "persons_generator/internal/persons/entities"
	religionEntities "persons_generator/internal/religion/entities"
)

func serializeCultures(cultures []*culture.SerializedCulture) ([]*SerializedCulture, error) {
	out := make([]*SerializedCulture, len(cultures))
	for i := range out {
		var err error
		out[i], err = serializeCulture(cultures[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can no serialize cultures")
		}
	}

	return out, nil
}

func serializeCulture(in *culture.SerializedCulture) (*SerializedCulture, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not marshal in serialize culture")
	}

	var out SerializedCulture
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not unmarshal in serialize culture")
	}

	return &out, nil
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

func serializeReligions(religions []*religion.SerializedReligion) ([]*SerializedReligion, error) {
	out := make([]*SerializedReligion, len(religions))
	for i := range out {
		var err error
		out[i], err = serializeReligion(religions[i])
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can no serialize religions")
		}
	}

	return out, nil
}

func serializeReligion(in *religion.SerializedReligion) (*SerializedReligion, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not marshal in serialize religion")
	}

	var out SerializedReligion
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not unmarshal in serialize religion")
	}

	return &out, nil
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

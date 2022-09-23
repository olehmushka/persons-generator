package religion

import "github.com/google/uuid"

type SerializedReligion struct {
	ID                  uuid.UUID           `json:"id"`
	Name                string              `json:"name"`
	Type                string              `json:"type"`
	GenderDominance     string              `json:"gender_dominance"`
	HighGoals           []string            `json:"high_goals"`
	DeityNature         string              `json:"deity_nature"`
	DeityTraits         []string            `json:"deity_traits"`
	HumanNature         string              `json:"human_nature"`
	HumanTraits         []string            `json:"human_traits"`
	SocialTraits        []string            `json:"social_traits"`
	SourceOfMoralLaw    string              `json:"source_of_moral_law"`
	Afterlife           map[string]string   `json:"afterlife"`
	AfterlifeTraits     []string            `json:"afterlife_traits"`
	Traits              []string            `json:"traits"`
	ClericsAppointment  string              `json:"clerics_appointment"`
	ClericsLimitations  string              `json:"clerics_limitations"`
	ClericsTraits       []string            `json:"clerics_traits"`
	ClericsFunctions    []string            `json:"clerics_functions"`
	HolyScriptureTraits []string            `json:"holy_scripture_traits"`
	TempleTraits        []string            `json:"temple_traits"`
	TheologyTraits      []string            `json:"theology_traits"`
	Cults               []string            `json:"cults"`
	Rules               []string            `json:"rules"`
	Taboos              []map[string]string `json:"taboos"`
	Rituals             []string            `json:"rituals"`
	Holydays            []string            `json:"holydays"`
	Conversion          []string            `json:"conversion"`
	MarriageTradition   map[string]string   `json:"marriage_tradition"`
}

func (r *Religion) Serialize() *SerializedReligion {
	if r == nil {
		return nil
	}

	return &SerializedReligion{
		ID:                  r.ID,
		Name:                r.Name,
		Type:                r.Type.String(),
		GenderDominance:     r.GenderDominance.String(),
		HighGoals:           r.Doctrine.HighGoal.Serialize(),
		DeityNature:         r.Doctrine.Deity.Nature.Goodness.String(),
		DeityTraits:         r.Doctrine.Deity.SerializeNatureTraits(),
		HumanNature:         r.Doctrine.Human.Nature.Goodness.String(),
		HumanTraits:         r.Doctrine.Human.SerializeNatureTraits(),
		SocialTraits:        r.Doctrine.Social.SerializeTraits(),
		SourceOfMoralLaw:    r.Doctrine.SourceOfMoralLaw.String(),
		Afterlife:           r.Doctrine.Afterlife.SerializeParticipant(),
		AfterlifeTraits:     r.Doctrine.Afterlife.SerializeTraits(),
		Traits:              r.Attributes.SerializeTraits(),
		ClericsAppointment:  r.Attributes.Clerics.SerializeAppointment(),
		ClericsLimitations:  r.Attributes.Clerics.SerializeLimitations(),
		ClericsTraits:       r.Attributes.Clerics.SerializeTraits(),
		ClericsFunctions:    r.Attributes.Clerics.SerializeFunctions(),
		HolyScriptureTraits: r.Attributes.HolyScripture.SerializeTraits(),
		TempleTraits:        r.Attributes.Temples.SerializeTraits(),
		TheologyTraits:      r.Theology.SerializeTraits(),
		Cults:               r.Theology.SerializeCults(),
		Rules:               r.Theology.SerializeRules(),
		Taboos:              r.Theology.SerializeTaboos(),
		Rituals:             r.Theology.SerializeRituals(),
		Holydays:            r.Theology.Holydays.Serialize(),
		Conversion:          r.Theology.Conversion.Serialize(),
		MarriageTradition:   r.Theology.MarriageTradition.Serialize(),
	}
}

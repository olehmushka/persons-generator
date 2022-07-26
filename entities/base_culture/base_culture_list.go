package base_culture

import pc "persons_generator/entities/proto_culture"

var (
	Italic   = &BaseCulture{Proto: pc.EuropeanProtoCulture, Name: "italic"}
	Etrurian = &BaseCulture{Proto: pc.EuropeanProtoCulture, Name: "etrurian"}
	Roman    = &BaseCulture{
		Proto:             pc.EuropeanProtoCulture,
		Name:              "roman",
		InheritedCultures: []*BaseCulture{Italic, Etrurian},
	}
	Italian = &BaseCulture{
		Proto:             pc.EuropeanProtoCulture,
		Name:              "italian",
		InheritedCultures: []*BaseCulture{Roman},
	}
	Celtic = &BaseCulture{Proto: pc.EuropeanProtoCulture, Name: "celtic"}
	French = &BaseCulture{
		Proto:             pc.EuropeanProtoCulture,
		Name:              "french",
		InheritedCultures: []*BaseCulture{Roman, Celtic},
	}
	British = &BaseCulture{
		Proto:             pc.EuropeanProtoCulture,
		Name:              "british",
		InheritedCultures: []*BaseCulture{Celtic},
	}
	Welsh = &BaseCulture{
		Proto:             pc.EuropeanProtoCulture,
		Name:              "welsh",
		InheritedCultures: []*BaseCulture{British},
	}
	Irish = &BaseCulture{
		Proto:             pc.EuropeanProtoCulture,
		Name:              "irish",
		InheritedCultures: []*BaseCulture{Celtic},
	}
	Iberian = &BaseCulture{
		Proto:             pc.EuropeanProtoCulture,
		Name:              "iberian",
		InheritedCultures: []*BaseCulture{Celtic},
	}
	Castillian = &BaseCulture{
		Proto:             pc.EuropeanProtoCulture,
		Name:              "castillian",
		InheritedCultures: []*BaseCulture{Roman, Iberian},
	}
	Spanish = &BaseCulture{
		Proto:             pc.EuropeanProtoCulture,
		Name:              "spanish",
		InheritedCultures: []*BaseCulture{Castillian},
	}
	Greek        = &BaseCulture{Proto: pc.EuropeanProtoCulture, Name: "greek"}
	Mesopotamian = &BaseCulture{
		Proto: pc.MesopotamianProtoCulture,
		Name:  "mesopotamian",
	}
	Babylonian = &BaseCulture{
		Proto:             pc.MesopotamianProtoCulture,
		Name:              "babylonian",
		InheritedCultures: []*BaseCulture{Mesopotamian},
	}
	Semitic = &BaseCulture{
		Proto:             pc.MesopotamianProtoCulture,
		Name:              "semitic",
		InheritedCultures: []*BaseCulture{Mesopotamian},
	}
	Jewish = &BaseCulture{
		Proto:             pc.MesopotamianProtoCulture,
		Name:              "jewish",
		InheritedCultures: []*BaseCulture{Semitic},
	}
	Egyptian = &BaseCulture{Proto: pc.EgyptianProtoCulture, Name: "egyptian"}
	Germanic = &BaseCulture{Proto: pc.EuropeanProtoCulture, Name: "germanic"}
	Anglic   = &BaseCulture{
		Proto:             pc.EuropeanProtoCulture,
		Name:              "Anglic",
		InheritedCultures: []*BaseCulture{Germanic},
	}
	Scandinavian = &BaseCulture{
		Proto:             pc.EuropeanProtoCulture,
		Name:              "Scandinavian",
		InheritedCultures: []*BaseCulture{Germanic},
	}
	Slavic    = &BaseCulture{Proto: pc.EuropeanProtoCulture, Name: "slavic"}
	Ruthenian = &BaseCulture{
		Proto:             pc.EuropeanProtoCulture,
		Name:              "ruthenian",
		InheritedCultures: []*BaseCulture{Slavic},
	}
	Baltic = &BaseCulture{Proto: pc.EuropeanProtoCulture, Name: "baltic"}
	Uralic = &BaseCulture{Proto: pc.UralicProtoCulture, Name: "uralic"}
	Finnic = &BaseCulture{
		Proto:             pc.UralicProtoCulture,
		Name:              "finnic",
		InheritedCultures: []*BaseCulture{Uralic},
	}
	Korean     = &BaseCulture{Proto: pc.EastAsianProtoCulture, Name: "korean"}
	Chinese    = &BaseCulture{Proto: pc.EastAsianProtoCulture, Name: "chinese"}
	Japanese   = &BaseCulture{Proto: pc.EastAsianProtoCulture, Name: "japanese"}
	Portuguese = &BaseCulture{
		Proto: pc.EuropeanProtoCulture, Name: "portuguese",
		InheritedCultures: []*BaseCulture{Roman, Iberian},
	}
	Nahuatl   = &BaseCulture{Proto: pc.MesoamericanProtoCulture, Name: "nahuatl"}
	Hungarian = &BaseCulture{
		Proto: pc.UralicProtoCulture, Name: "hungarian",
		InheritedCultures: []*BaseCulture{Uralic},
	}
	Turkish = &BaseCulture{Proto: pc.TurkProtoCulture, Name: "turkish"}
	Berber  = &BaseCulture{Proto: pc.AfricanProtoCulture, Name: "berber"}
	Arabic  = &BaseCulture{
		Proto: pc.MesopotamianProtoCulture, Name: "arabic",
		InheritedCultures: []*BaseCulture{Semitic},
	}
	Inuit      = &BaseCulture{Proto: pc.InuitProtoCulture, Name: "inuit"}
	Basque     = &BaseCulture{Proto: pc.EuropeanProtoCulture, Name: "basque"}
	Nigerian   = &BaseCulture{Proto: pc.AfricanProtoCulture, Name: "nigerian"}
	Iranian    = &BaseCulture{Proto: pc.MesopotamianProtoCulture, Name: "iranian"}
	Polynesian = &BaseCulture{Proto: pc.PolynesianProtoCulture, Name: "polynesian"}
	Hawaiian   = &BaseCulture{
		Proto: pc.EastAsianProtoCulture, Name: "hawaiian",
		InheritedCultures: []*BaseCulture{Polynesian},
	}
	Arian  = &BaseCulture{Proto: pc.IndianProtoCulture, Name: "arian"}
	Indian = &BaseCulture{
		Proto: pc.IndianProtoCulture, Name: "indian",
		InheritedCultures: []*BaseCulture{Arian},
	}
	Karnataka = &BaseCulture{
		Proto: pc.IndianProtoCulture, Name: "karnataka",
		InheritedCultures: []*BaseCulture{Indian},
	}
	Quechua    = &BaseCulture{Proto: pc.SouthAmericanProtoCulture, Name: "quechua"}
	Swahili    = &BaseCulture{Proto: pc.AfricanProtoCulture, Name: "swahili"}
	Vietnamese = &BaseCulture{Proto: pc.EastAsianProtoCulture, Name: "vietnamese"}
	Cantonese  = &BaseCulture{
		Proto: pc.EastAsianProtoCulture, Name: "cantonese",
		InheritedCultures: []*BaseCulture{Chinese},
	}
	Mongolian = &BaseCulture{Proto: pc.MongolianProtoCulture, Name: "mongolian"}
	Elven     = &BaseCulture{Proto: pc.ElvenProtoCulture, Name: "elven"}
	DarkElven = &BaseCulture{Proto: pc.DarkElvenProtoCulture, Name: "dark_elven"}
	Dwarven   = &BaseCulture{Proto: pc.DwarvenProtoCulture, Name: "dwarven"}
	Goblin    = &BaseCulture{Proto: pc.GoblinProtoCulture, Name: "goblin"}
	Orc       = &BaseCulture{Proto: pc.OrcProtoCulture, Name: "orc"}
	Giant     = &BaseCulture{Proto: pc.GiantProtoCulture, Name: "giant"}
	Draconic  = &BaseCulture{Proto: pc.DraconicProtoCulture, Name: "draconic"}
	Arachnid  = &BaseCulture{Proto: pc.ArachnidProtoCulture, Name: "arachnid"}
	Serpents  = &BaseCulture{Proto: pc.SerpentsProtoCulture, Name: "serpents"}
)

var AllBaseCultures = []*BaseCulture{
	Italic,
	Etrurian,
	Roman,
	Italian,
	Celtic,
	French,
	British,
	Welsh,
	Irish,
	Iberian,
	Castillian,
	Spanish,
	Greek,
	Mesopotamian,
	Babylonian,
	Semitic,
	Jewish,
	Egyptian,
	Germanic,
	Anglic,
	Scandinavian,
	Slavic,
	Ruthenian,
	Baltic,
	Uralic,
	Finnic,
	Korean,
	Chinese,
	Japanese,
	Portuguese,
	Nahuatl,
	Hungarian,
	Turkish,
	Berber,
	Arabic,
	Inuit,
	Basque,
	Nigerian,
	Iranian,
	Polynesian,
	Hawaiian,
	Arian,
	Indian,
	Karnataka,
	Quechua,
	Swahili,
	Vietnamese,
	Cantonese,
	Mongolian,
	Elven,
	DarkElven,
	Dwarven,
	Goblin,
	Orc,
	Giant,
	Draconic,
	Arachnid,
	Serpents,
}

var CulturesWithoutLastName = []*BaseCulture{
	Mesopotamian,
	Babylonian,
	Egyptian,
}

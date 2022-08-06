package language

import "persons_generator/tools"

var AllFemaleOwnNames = tools.MergeUnique(
	AllFantasyFemaleOwnNames,
	AllIndoEuropeanFemaleOwnNames,
	AllEastAsianFemaleOwnNames,
	AllIsolatedFemaleOwnNames,
	AllUralicFemaleOwnNames,
	AllOtherFemaleOwnNames,
)

// Fantasy
var (
	FemaleOwnNamesSindarin = []string{}
	FemaleOwnNamesDunmeris = []string{}
	FemaleOwnNamesDwemeris = []string{}
	FemaleOwnNamesGoblins  = []string{}
	FemaleOwnNamesOrc      = []string{}
	FemaleOwnNamesGiant    = []string{}
	FemaleOwnNamesDraconic = []string{}
	FemaleOwnNamesArachnid = []string{}
	FemaleOwnNamesSerpents = []string{}
)

var AllFantasyFemaleOwnNames = tools.MergeUnique(
	FemaleOwnNamesSindarin,
	FemaleOwnNamesDunmeris,
	FemaleOwnNamesDwemeris,
	FemaleOwnNamesGoblins,
	FemaleOwnNamesOrc,
	FemaleOwnNamesGiant,
	FemaleOwnNamesDraconic,
	FemaleOwnNamesArachnid,
	FemaleOwnNamesSerpents,
)

var (
	FemaleOwnNamesAlbanian   = []string{}
	FemaleOwnNamesArmenian   = []string{}
	FemaleOwnNamesRuthenian  = []string{}
	FemaleOwnNamesLithuanian = []string{}
	FemaleOwnNamesCeltic     = []string{}
	FemaleOwnNamesDacian     = []string{}
	FemaleOwnNamesNordic     = []string{}
	FemaleOwnNamesEnglish    = []string{}
	FemaleOwnNamesGerman     = []string{}
	FemaleOwnNamesGreek      = []string{}
	FemaleOwnNamesIllyrian   = []string{}
	FemaleOwnNamesHindi      = []string{}
	FemaleOwnNamesIranian    = []string{}
	FemaleOwnNamesLatin      = []string{}
	FemaleOwnNamesPortuguese = []string{}
	FemaleOwnNamesSpanish    = []string{}
	FemaleOwnNamesFrench     = []string{}
	FemaleOwnNamesItalian    = []string{}
	FemaleOwnNamesEtruscan   = []string{}
)

var AllIndoEuropeanFemaleOwnNames = tools.MergeUnique(
	FemaleOwnNamesAlbanian,
	FemaleOwnNamesArmenian,
	FemaleOwnNamesRuthenian,
	FemaleOwnNamesLithuanian,
	FemaleOwnNamesCeltic,
	FemaleOwnNamesDacian,
	FemaleOwnNamesNordic,
	FemaleOwnNamesEnglish,
	FemaleOwnNamesGerman,
	FemaleOwnNamesGreek,
	FemaleOwnNamesIllyrian,
	FemaleOwnNamesHindi,
	FemaleOwnNamesIranian,
	FemaleOwnNamesLatin,
	FemaleOwnNamesPortuguese,
	FemaleOwnNamesSpanish,
	FemaleOwnNamesFrench,
	FemaleOwnNamesItalian,
	FemaleOwnNamesEtruscan,
)

// EastAsian
var (
	FemaleOwnNamesChinese    = []string{}
	FemaleOwnNamesCantonese  = []string{}
	FemaleOwnNamesJapanese   = []string{}
	FemaleOwnNamesKorean     = []string{}
	FemaleOwnNamesVietnamese = []string{}
	FemaleOwnNamesKannada    = []string{}
)

var AllEastAsianFemaleOwnNames = tools.MergeUnique(
	FemaleOwnNamesChinese,
	FemaleOwnNamesCantonese,
	FemaleOwnNamesJapanese,
	FemaleOwnNamesKorean,
	FemaleOwnNamesVietnamese,
	FemaleOwnNamesKannada,
)

// Isolated
var (
	FemaleOwnNamesBasque   = []string{}
	FemaleOwnNamesGeorgian = []string{}
)
var AllIsolatedFemaleOwnNames = tools.MergeUnique(FemaleOwnNamesBasque, FemaleOwnNamesGeorgian)

// Uralic
var (
	FemaleOwnNamesFinnic    = []string{}
	FemaleOwnNamesHungarian = []string{}
)
var AllUralicFemaleOwnNames = tools.MergeUnique(FemaleOwnNamesFinnic, FemaleOwnNamesHungarian)

// Other
var (
	FemaleOwnNamesTurkish      = []string{}
	FemaleOwnNamesMongolian    = []string{}
	FemaleOwnNamesSwahili      = []string{}
	FemaleOwnNamesNigerian     = []string{}
	FemaleOwnNamesNahuatl      = []string{}
	FemaleOwnNamesInuit        = []string{}
	FemaleOwnNamesHawaiian     = []string{}
	FemaleOwnNamesQuechua      = []string{}
	FemaleOwnNamesBerber       = []string{}
	FemaleOwnNamesMesopotamian = []string{}
	FemaleOwnNamesArabic       = []string{}
)

var AllOtherFemaleOwnNames = tools.MergeUnique(
	FemaleOwnNamesTurkish,
	FemaleOwnNamesMongolian,
	FemaleOwnNamesSwahili,
	FemaleOwnNamesNigerian,
	FemaleOwnNamesNahuatl,
	FemaleOwnNamesInuit,
	FemaleOwnNamesHawaiian,
	FemaleOwnNamesQuechua,
	FemaleOwnNamesBerber,
	FemaleOwnNamesMesopotamian,
	FemaleOwnNamesArabic,
)

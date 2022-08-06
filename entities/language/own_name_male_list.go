package language

import "persons_generator/tools"

var AllMaleOwnNames = tools.MergeUnique(
	AllFantasyMaleOwnNames,
	AllIndoEuropeanMaleOwnNames,
	AllEastAsianMaleOwnNames,
	AllIsolatedMaleOwnNames,
	AllUralicMaleOwnNames,
	AllOtherMaleOwnNames,
)

// Fantasy
var (
	MaleOwnNamesSindarin = []string{}
	MaleOwnNamesDunmeris = []string{}
	MaleOwnNamesDwemeris = []string{}
	MaleOwnNamesGoblins  = []string{}
	MaleOwnNamesOrc      = []string{}
	MaleOwnNamesGiant    = []string{}
	MaleOwnNamesDraconic = []string{}
	MaleOwnNamesArachnid = []string{}
	MaleOwnNamesSerpents = []string{}
)

var AllFantasyMaleOwnNames = tools.MergeUnique(
	MaleOwnNamesSindarin,
	MaleOwnNamesDunmeris,
	MaleOwnNamesDwemeris,
	MaleOwnNamesGoblins,
	MaleOwnNamesOrc,
	MaleOwnNamesGiant,
	MaleOwnNamesDraconic,
	MaleOwnNamesArachnid,
	MaleOwnNamesSerpents,
)

var (
	MaleOwnNamesAlbanian   = []string{}
	MaleOwnNamesArmenian   = []string{}
	MaleOwnNamesRuthenian  = []string{}
	MaleOwnNamesLithuanian = []string{}
	MaleOwnNamesCeltic     = []string{}
	MaleOwnNamesDacian     = []string{}
	MaleOwnNamesNordic     = []string{}
	MaleOwnNamesEnglish    = []string{}
	MaleOwnNamesGerman     = []string{}
	MaleOwnNamesGreek      = []string{}
	MaleOwnNamesIllyrian   = []string{}
	MaleOwnNamesHindi      = []string{}
	MaleOwnNamesIranian    = []string{}
	MaleOwnNamesLatin      = []string{}
	MaleOwnNamesPortuguese = []string{}
	MaleOwnNamesSpanish    = []string{}
	MaleOwnNamesFrench     = []string{}
	MaleOwnNamesItalian    = []string{}
	MaleOwnNamesEtruscan   = []string{}
)

var AllIndoEuropeanMaleOwnNames = tools.MergeUnique(
	MaleOwnNamesAlbanian,
	MaleOwnNamesArmenian,
	MaleOwnNamesRuthenian,
	MaleOwnNamesLithuanian,
	MaleOwnNamesCeltic,
	MaleOwnNamesDacian,
	MaleOwnNamesNordic,
	MaleOwnNamesEnglish,
	MaleOwnNamesGerman,
	MaleOwnNamesGreek,
	MaleOwnNamesIllyrian,
	MaleOwnNamesHindi,
	MaleOwnNamesIranian,
	MaleOwnNamesLatin,
	MaleOwnNamesPortuguese,
	MaleOwnNamesSpanish,
	MaleOwnNamesFrench,
	MaleOwnNamesItalian,
	MaleOwnNamesEtruscan,
)

// EastAsian
var (
	MaleOwnNamesChinese    = []string{}
	MaleOwnNamesCantonese  = []string{}
	MaleOwnNamesJapanese   = []string{}
	MaleOwnNamesKorean     = []string{}
	MaleOwnNamesVietnamese = []string{}
	MaleOwnNamesKannada    = []string{}
)

var AllEastAsianMaleOwnNames = tools.MergeUnique(
	MaleOwnNamesChinese,
	MaleOwnNamesCantonese,
	MaleOwnNamesJapanese,
	MaleOwnNamesKorean,
	MaleOwnNamesVietnamese,
	MaleOwnNamesKannada,
)

// Isolated
var (
	MaleOwnNamesBasque   = []string{}
	MaleOwnNamesGeorgian = []string{}
)
var AllIsolatedMaleOwnNames = tools.MergeUnique(MaleOwnNamesBasque, MaleOwnNamesGeorgian)

// Uralic
var (
	MaleOwnNamesFinnic    = []string{}
	MaleOwnNamesHungarian = []string{}
)
var AllUralicMaleOwnNames = tools.MergeUnique(MaleOwnNamesFinnic, MaleOwnNamesHungarian)

// Other
var (
	MaleOwnNamesTurkish      = []string{}
	MaleOwnNamesMongolian    = []string{}
	MaleOwnNamesSwahili      = []string{}
	MaleOwnNamesNigerian     = []string{}
	MaleOwnNamesNahuatl      = []string{}
	MaleOwnNamesInuit        = []string{}
	MaleOwnNamesHawaiian     = []string{}
	MaleOwnNamesQuechua      = []string{}
	MaleOwnNamesBerber       = []string{}
	MaleOwnNamesMesopotamian = []string{}
	MaleOwnNamesArabic       = []string{}
)

var AllOtherMaleOwnNames = tools.MergeUnique(
	MaleOwnNamesTurkish,
	MaleOwnNamesMongolian,
	MaleOwnNamesSwahili,
	MaleOwnNamesNigerian,
	MaleOwnNamesNahuatl,
	MaleOwnNamesInuit,
	MaleOwnNamesHawaiian,
	MaleOwnNamesQuechua,
	MaleOwnNamesBerber,
	MaleOwnNamesMesopotamian,
	MaleOwnNamesArabic,
)

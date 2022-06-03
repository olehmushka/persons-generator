package religion

type Theology struct {
	Precepts      *Precepts
	HolyScripture *HolyScripture
}

type HolyScripture struct {
	Authors       HolyScriptureAuthors
	Understanding HolyScriptureUnderstanding
	IsWorshiped   bool
}

type HolyScriptureAuthors string

const (
	Unknown       HolyScriptureAuthors = "unknown"
	OneHuman      HolyScriptureAuthors = "one_human"
	SeveralHumans HolyScriptureAuthors = "several_humans"
	NoOne         HolyScriptureAuthors = "no_one"
)

type HolyScriptureUnderstanding string

const (
	Literal HolyScriptureUnderstanding = "literal"
)

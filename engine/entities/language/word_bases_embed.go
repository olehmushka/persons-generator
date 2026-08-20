package language

import "embed"

//go:embed word_bases/*.json
var wordBasesFS embed.FS

package language

type WordBase struct {
	Words []string
	Name  string
	Min   int
	Max   int
	Dupl  string
	M     float64
}

func ExtractWords(wbs []*WordBase) map[string][]string {
	out := make(map[string][]string, len(wbs))
	for _, wb := range wbs {
		out[wb.Name] = wb.Words
	}

	return out
}

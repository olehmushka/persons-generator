package language

type WordBase struct {
	FemaleOwnNames []string `json:"female_own_names"`
	MaleOwnNames   []string `json:"male_own_names"`
	Words          []string `json:"words"`
	Name           string   `json:"name"`
	Min            int      `json:"min"`
	Max            int      `json:"max"`
	Dupl           string   `json:"dupl"`
	M              float64  `json:"m"`
}

func ExtractWords(wbs []*WordBase) map[string][]string {
	out := make(map[string][]string, len(wbs))
	for _, wb := range wbs {
		out[wb.Name] = wb.Words
	}

	return out
}

func SetWordBases() (err error) {
	AllWordBases = make([]*WordBase, len(RefAllWordBases))
	for i := range AllWordBases {
		AllWordBases[i], err = RefAllWordBases[i].LoadWordBase()
		if err != nil {
			return err
		}
	}

	return err
}

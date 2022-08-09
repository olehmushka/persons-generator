package language

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

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

func getWordBaseFromJSON(filename string) (*WordBase, error) {
	jsonFile, err := os.Open("entities/language/word_bases/" + filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var wb WordBase
	if err := json.Unmarshal(byteValue, &wb); err != nil {
		return nil, err
	}

	return &wb, nil
}

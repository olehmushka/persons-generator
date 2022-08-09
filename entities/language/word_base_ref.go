package language

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const PrefixWordBasesPath = "entities/language/word_bases/"

type WordBaseRef struct {
	filename string
}

func (wbr *WordBaseRef) LoadWordBase() (*WordBase, error) {
	jsonFile, err := os.Open(PrefixWordBasesPath + wbr.filename)
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

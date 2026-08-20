package language

import (
	"encoding/json"
)

type WordBaseRef struct {
	Filename string `json:"filename" bson:"filename"`
}

func (wbr *WordBaseRef) LoadWordBase() (*WordBase, error) {
	byteValue, err := wordBasesFS.ReadFile("word_bases/" + wbr.Filename)
	if err != nil {
		return nil, err
	}

	var wb WordBase
	if err := json.Unmarshal(byteValue, &wb); err != nil {
		return nil, err
	}

	return &wb, nil
}

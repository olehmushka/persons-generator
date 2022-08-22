package entities

type Afterlife struct {
	Participants AfterlifeParticipants `json:"participants"`
	Traits       []string              `json:"traits"`
}

type AfterlifeParticipants struct {
	ForTopBelievers    string `json:"for_top_believers"`
	ForBelievers       string `json:"for_believers"`
	ForUntrueBelievers string `json:"for_untrue_believers"`
	ForAtheists        string `json:"for_atheists"`
}

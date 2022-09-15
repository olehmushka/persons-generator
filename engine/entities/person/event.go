package person

type Event struct {
	Year  int         `json:"year"`
	Name  EventName   `json:"name"`
	Value interface{} `json:"value"`
}

type EventName string

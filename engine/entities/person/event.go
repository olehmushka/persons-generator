package person

type Event struct {
	Year int       `json:"year"`
	Name EventName `json:"name"`
}

type EventName string

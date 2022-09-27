package tools

import "time"

var TimeFormat = time.RFC3339

func SerializeTime(in time.Time) string {
	return in.Format(TimeFormat)
}

func DeserializeTime(in string) (time.Time, error) {
	return time.Parse(TimeFormat, in)
}

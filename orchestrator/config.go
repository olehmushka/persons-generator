package orchestrator

type Config struct {
	WorldSize int `default:"30"`
	Religion  ReligionConfig
}

type ReligionConfig struct {
	Amount int `default:"3"`
}

type CultureConfig struct {
	Preferred []string
	Amount    int `default:"3"`
}

package orchestrator

type Config struct {
	WorldSize int `default:"30"`
	Religion  ReligionConfig
	Culture   CultureConfig
}

type ReligionConfig struct {
	Amount int `default:"3"`
}

type CultureConfig struct {
	Preferred []string
	Amount    int `default:"3"`
}

package config

type HTTPServer struct {
	Address string `env:"HTTP_SERVER_ADDRESS"`
}

package main

import (
	"fmt"
	"persons_generator/config"
	"persons_generator/person"
)

func main() {
	cfg := config.New()
	persons := person.GenerateMany(&person.GenerateManyConfig{
		Mode: cfg.Mode,
	})
	fmt.Println(persons)
}

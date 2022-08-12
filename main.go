package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"persons_generator/cli"
	l "persons_generator/engine/entities/language"

	"github.com/sirupsen/logrus"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	if err := l.SetWordBases(); err != nil {
		fmt.Printf("[language.SetWordBases] error = %+v\n", err)
		os.Exit(1)
		return
	}

	if err := cli.Execute(os.Args[1:]); err != nil {
		logrus.WithField("msg", "cli.Execute error").Error(err)
		os.Exit(1)
	}
}

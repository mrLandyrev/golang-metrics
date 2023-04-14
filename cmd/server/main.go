package main

import (
	"flag"
	"os"

	"github.com/mrLandyrev/golang-metrics/internal/server/app"
)

func main() {
	a := flag.String("a", "localhost:8080", "address")
	flag.Parse()

	if envAddress := os.Getenv("ADDRESS"); envAddress != "" {
		a = &envAddress
	}

	app.NewApp(*a).Run()
}

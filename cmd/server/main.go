package main

import (
	"flag"
	"os"

	"github.com/mrLandyrev/golang-metrics/internal/server/app"
)

var config app.ServerConfig

func buildConfig() {
	flag.StringVar(&config.Address, "a", "localhost:8080", "metrics server address")
	flag.Parse()

	if envA := os.Getenv("ADDRESS"); envA != "" {
		config.Address = envA
	}
}

func main() {
	buildConfig()
	app.NewServerApp(config).Run()
}

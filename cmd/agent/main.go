package main

import (
	"flag"
	"os"
	"time"

	"github.com/mrLandyrev/golang-metrics/internal/agent/app"
)

func main() {
	serverAddress := flag.String("a", "localhost:8080", "metrics server address")
	syncInteval := flag.Duration("r", time.Second*10, "time between sync metrics with server")
	collectInterval := flag.Duration("p", time.Second*2, "time between run collect metrics")
	flag.Parse()

	if envA := os.Getenv("ADDRESS"); envA != "" {
		serverAddress = &envA
	}

	if envR := os.Getenv("REPORT_INTERVAL"); envR != "" {
		parsed, err := time.ParseDuration(envR)
		if err != nil {
			panic(err)
		}
		*syncInteval = parsed
	}

	if envP := os.Getenv("POLL_INTERVAL"); envP != "" {
		parsed, err := time.ParseDuration(envP)
		if err != nil {
			panic(err)
		}
		*collectInterval = parsed
	}

	app.NewApp(*serverAddress, *syncInteval, *collectInterval).Run()
}

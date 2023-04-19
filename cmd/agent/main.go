package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/mrLandyrev/golang-metrics/internal/agent/app"
)

func main() {
	serverAddress := flag.String("a", "localhost:8080", "metrics server address")
	syncInteval := flag.Int("r", 10, "time between sync metrics with server in seconds")
	collectInterval := flag.Int("p", 2, "time between run collect metrics in seconds")
	flag.Parse()

	if envA := os.Getenv("ADDRESS"); envA != "" {
		serverAddress = &envA
	}

	if envR := os.Getenv("REPORT_INTERVAL"); envR != "" {
		parsed, err := strconv.Atoi(envR)
		if err != nil {
			panic(err)
		}
		*syncInteval = parsed
	}

	if envP := os.Getenv("POLL_INTERVAL"); envP != "" {
		parsed, err := strconv.Atoi(envP)
		if err != nil {
			panic(err)
		}
		*collectInterval = parsed
	}

	app.NewApp(*serverAddress, *syncInteval, *collectInterval).Run()
}

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/mrLandyrev/golang-metrics/internal/agent/app"
)

var config app.Config

func buildConfig() {
	flag.StringVar(&config.ServerAddress, "a", "localhost:8080", "metrics server address")
	flag.IntVar(&config.SyncInteval, "r", 10, "time between sync metrics with server in seconds")
	flag.IntVar(&config.CollectInterval, "p", 2, "time between run collect metrics in seconds")
	flag.StringVar(&config.SignKey, "k", "", "sign data key")
	flag.Parse()

	if envA := os.Getenv("ADDRESS"); envA != "" {
		config.ServerAddress = envA
	}

	if envR := os.Getenv("REPORT_INTERVAL"); envR != "" {
		parsed, err := strconv.Atoi(envR)
		if err != nil {
			fmt.Println("Cannot parse REPORT_INTERVAL variable")
		} else {
			config.SyncInteval = parsed
		}
	}

	if envP := os.Getenv("POLL_INTERVAL"); envP != "" {
		parsed, err := strconv.Atoi(envP)
		if err != nil {
			fmt.Println("Cannot parse POLL_INTERVAL variable")
		} else {
			config.CollectInterval = parsed
		}
	}

	if envK := os.Getenv("KEY"); envK != "" {
		config.SignKey = envK
	}
}

func main() {
	buildConfig()
	agent := app.NewAgentApp(config)
	agent.Run()
	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	<-gracefulStop
	agent.Stop()
}

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/mrLandyrev/golang-metrics/internal/server/app"
)

var config app.ServerConfig

func buildConfig() {
	var storeInterval int
	flag.StringVar(&config.Address, "a", "localhost:8080", "metrics server address")
	flag.IntVar(&storeInterval, "i", 300, "time between store metrics to file")
	flag.StringVar(&config.FileStoragePath, "f", "/tmp/metrics-db.json", "path to file where storage metrics")
	flag.BoolVar(&config.NeedRestore, "r", true, "need restore data on startup")
	flag.StringVar(&config.DatabaseConnection, "d", "", "database connection string")
	flag.StringVar(&config.SignKey, "k", "", "sign data key")
	flag.Parse()

	if envA := os.Getenv("ADDRESS"); envA != "" {
		config.Address = envA
	}

	if envI := os.Getenv("STORE_INTERVAL"); envI != "" {
		parsed, err := strconv.Atoi(envI)
		if err != nil {
			fmt.Println("Cannot parse STORE_INTERVAL variable")
		} else {
			storeInterval = parsed
		}
	}

	config.StoreInterval = time.Second * time.Duration(storeInterval)

	if envF := os.Getenv("FILE_STORAGE_PATH"); envF != "" {
		config.FileStoragePath = envF
	}

	if envR := os.Getenv("RESTORE"); envR != "" {
		parsed, err := strconv.ParseBool(envR)
		if err != nil {
			fmt.Println("Cannot parse RESTORE variable")
		} else {
			config.NeedRestore = parsed
		}
	}

	if envD := os.Getenv("DATABASE_DSN"); envD != "" {
		config.DatabaseConnection = envD
	}

	if envK := os.Getenv("KEY"); envK != "" {
		config.SignKey = envK
	}
}

func main() {
	buildConfig()
	app := app.NewServerApp(config)
	app.Run()
	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	<-gracefulStop
	app.Stop()
}

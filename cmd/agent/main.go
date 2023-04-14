package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/mrLandyrev/golang-metrics/internal/agent/app"
)

func main() {
	a := flag.String("a", "localhost:8080", "endpoint")
	r := flag.Int64("r", 10, "report interval")
	p := flag.Int64("p", 2, "poll interval")
	flag.Parse()

	if envA := os.Getenv("ADDRESS"); envA != "" {
		a = &envA
	}

	if envR := os.Getenv("REPORT_INTERVAL"); envR != "" {
		parsed, err := strconv.ParseInt(envR, 10, 64)
		if err != nil {
			panic(err)
		}
		*r = parsed
	}

	if envP := os.Getenv("POLL_INTERVAL"); envP != "" {
		parsed, err := strconv.ParseInt(envP, 10, 64)
		if err != nil {
			panic(err)
		}
		*p = parsed
	}

	app.NewApp(*a, *r, *p).Run()
}

package main

import (
	"flag"

	"github.com/mrLandyrev/golang-metrics/internal/agent/app"
)

func main() {
	a := flag.String("a", "localhost:8080", "endpoint")
	r := flag.Int64("r", 10, "report interval")
	p := flag.Int64("p", 2, "poll interval")
	flag.Parse()

	app.NewApp(*a, *r, *p).Run()
}

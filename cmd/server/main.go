package main

import (
	"flag"

	"github.com/mrLandyrev/golang-metrics/internal/server/app"
)

func main() {
	a := flag.String("a", "localhost:8080", "address")
	flag.Parse()

	app.NewApp(*a).Run()
}

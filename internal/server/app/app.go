package app

import (
	"net/http"

	"github.com/mrLandyrev/golang-metrics/internal/server/app/transport/rest"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/factory"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/repository"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/service"
	"github.com/mrLandyrev/golang-metrics/pkg/router"
)

type App struct {
	router *router.Router
}

func (app *App) Run() {
	http.ListenAndServe(":8080", app.router)
}

func NewApp() *App {
	metricsRepository := repository.NewMemoryMetricsRepository()
	metricsFactory := factory.NewMetricsFactory()
	metricsService := service.NewMetricsService(metricsRepository, metricsFactory)

	router := router.NewRouter()

	router.Use("POST", "/update/:kind/:name/:value", rest.BuildUpdateMetricHandler(metricsService))
	router.Use("GET", "/get/:kind/:name", rest.BuildGetMetricHandler(metricsService))

	return &App{router: router}
}

package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrLandyrev/golang-metrics/internal/server/app/transport/rest"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/factory"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/repository"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/service"
)

type App struct {
	router *gin.Engine
	a      string
}

func (app *App) Run() {
	http.ListenAndServe(app.a, app.router)
}

func NewApp(a string) *App {
	metricsRepository := repository.NewMemoryMetricsRepository()
	metricsFactory := factory.NewMetricsFactory()
	metricsService := service.NewMetricsService(metricsRepository, metricsFactory)

	router := gin.New()

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", rest.BuildGetAllMetricHandler(metricsService))
	router.POST("/update/:kind/:name/:value", rest.BuildUpdateMetricHandler(metricsService))
	router.GET("/value/:kind/:name", rest.BuildGetMetricHandler(metricsService))

	return &App{router: router, a: a}
}

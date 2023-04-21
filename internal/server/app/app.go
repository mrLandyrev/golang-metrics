package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrLandyrev/golang-metrics/internal/metrics"
	"github.com/mrLandyrev/golang-metrics/internal/server/app/transport/rest"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/service"
)

type App struct {
	router  *gin.Engine
	address string
}

func (app *App) Run() {
	http.ListenAndServe(app.address, app.router)
}

func NewApp(address string) *App {
	// build dependencies
	metricsRepository := metrics.NewMemoryMetricsRepository()
	metricsFactory := metrics.NewMetricsFactory()
	metricsService := service.NewMetricsService(metricsRepository, metricsFactory)

	router := gin.New()

	// register handlers
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", rest.BuildGetAllMetricHandler(metricsService))
	router.POST("/update/:kind/:name/:value", rest.BuildUpdateMetricHandler(metricsService))
	router.GET("/value/:kind/:name", rest.BuildGetMetricHandler(metricsService))

	return &App{router: router, address: address}
}

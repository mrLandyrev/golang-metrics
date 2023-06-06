package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrLandyrev/golang-metrics/internal/metrics"
	"github.com/mrLandyrev/golang-metrics/internal/metrics/storage"
	"github.com/mrLandyrev/golang-metrics/internal/server/app/transport/rest"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/service"
	"go.uber.org/zap"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type ServerApp struct {
	server *http.Server
	db     *sql.DB
	flush  func() error
}

func (app *ServerApp) Run() {
	go func() {
		app.server.ListenAndServe()
	}()
}

func (app *ServerApp) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	app.server.Shutdown(ctx)
	if app.flush != nil {
		app.flush()
	}
	if app.db != nil {
		app.db.Close()
	}
}

func NewServerApp(config ServerConfig) *ServerApp {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln("Cannot create logger")
	}

	logger.Debug(fmt.Sprint(config))

	db, err := sql.Open("pgx", config.DatabaseConnection)
	if err != nil {
		logger.Fatal("Cannot connect database")
	}

	var metricsRepository service.MetricsRepository
	var flushCallback func() error
	// build dependencies
	if config.DatabaseConnection != "" {
		metricsRepository, err = storage.NewDatabaseMetricsRepository(db)
		if err != nil {
			logger.Fatal(err.Error())
		}
	} else if config.FileStoragePath != "" {
		fileMetricsRepository, err := storage.NewFileMetricsRepository(config.FileStoragePath, config.StoreInterval, config.NeedRestore)
		if err != nil {
			logger.Fatal(err.Error())
		}
		metricsRepository = fileMetricsRepository
		flushCallback = fileMetricsRepository.Flush
	} else {
		metricsRepository = storage.NewMemoryMetricsRepository()
	}
	metricsFactory := metrics.NewMetricsFactory()
	metricsService := service.NewMetricsService(metricsRepository, metricsFactory)

	router := gin.New()

	router.Use(rest.LoggingMiddleware(logger))
	router.Use(rest.GzipMiddleware())

	// register handlers
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", rest.BuildGetAllMetricHandler(metricsService))
	router.POST("/update/:kind/:name/:value", rest.BuildUpdateMetricHandler(metricsService))
	router.POST("/update", rest.BuildJSONUpdateMetricHandler(metricsService))
	router.POST("/value", rest.BuildJSONGetMetricHandler(metricsService))
	router.POST("/updates", rest.BuildJSONBatchUpdateMetricsHandler(metricsService))
	router.GET("/value/:kind/:name", rest.BuildGetMetricHandler(metricsService))
	router.GET("/ping", rest.BuildPingHandler(db))

	server := http.Server{
		Addr:    config.Address,
		Handler: router,
	}

	return &ServerApp{server: &server, db: db, flush: flushCallback}
}

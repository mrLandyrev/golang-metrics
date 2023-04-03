package metrics

import (
	"fmt"
	"net/http"

	"github.com/mrLandyrev/golang-metrics/internal/router"
)

type Server struct {
	metricsService *Service
}

func (server *Server) Listen() {
	router := router.NewRouter()
	router.Use("POST", "/update/:kind/:name/:value", server.updateHandler)
	router.Use("GET", "/get/:kind/:name", server.getHandler)

	http.ListenAndServe(":8080", router)
}

func (server *Server) updateHandler(w http.ResponseWriter, r *http.Request, c *router.Context) {
	err := server.metricsService.AddRecord(c.GetPathParam("kind"), c.GetPathParam("name"), c.GetPathParam("value"))

	switch err {
	case nil:
		break
	case ErrUnknownMetricKind:
		w.WriteHeader(http.StatusNotImplemented)
		return
	case ErrIncorrectMetricValue:
		w.WriteHeader(http.StatusBadRequest)
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (server *Server) getHandler(w http.ResponseWriter, r *http.Request, c *router.Context) {
	item, err := server.metricsService.GetRecord(c.GetPathParam("kind"), c.GetPathParam("name"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if item == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, item.GetStrValue())
}

func NewServer(metricsService *Service) *Server {
	return &Server{metricsService: metricsService}
}

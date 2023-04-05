package factory

import (
	"errors"

	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/types"
)

var ErrUnknownMetricKind = errors.New("unknown metric kind detected")

type MetricsFactory interface {
	GetInstance(kind string, name string) (types.Metric, error)
}

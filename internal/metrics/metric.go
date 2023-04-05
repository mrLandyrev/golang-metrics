package metrics

import "errors"

var (
	ErrIncorrectMetricValue = errors.New("incorrect metric value")
	ErrUnknownMetricKind    = errors.New("unknown metric kind detected")
)

type Metric interface {
	GetName() string
	GetKind() string
	AddStrValue(value string) error
	GetStrValue() string
}

package metrics

import "errors"

var (
	ErrIncorrectMetricValue = errors.New("Incorrect metric value")
)

type Metric interface {
	GetName() string
	GetKind() string
	AddStrValue(value string) error
	GetStrValue() string
}

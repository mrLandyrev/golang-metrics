package types

import "errors"

var (
	ErrIncorrectMetricValue = errors.New("incorrect metric value")
)

type Metric interface {
	GetName() string
	GetKind() string
	AddValue(value string) error
	GetValue() string
}

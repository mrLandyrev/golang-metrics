package metrics

import "errors"

var (
	ErrUnknownMetricKind    = errors.New("unknow metric kind")
	ErrIncorrectMetricValue = errors.New("incorrect metric value")
)

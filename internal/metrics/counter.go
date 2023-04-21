package metrics

import "strconv"

type CounterMetric struct {
	name  string
	value int64
}

func (counter *CounterMetric) Name() string {
	return counter.name
}

func (counter *CounterMetric) Kind() string {
	return "counter"
}

func (counter *CounterMetric) AddValue(value string) error {
	convertedValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return ErrIncorrectMetricValue
	}

	counter.value += convertedValue

	return nil
}

func (counter *CounterMetric) Value() string {
	return strconv.FormatInt(counter.value, 10)
}

func NewCounterMetric(name string) *CounterMetric {
	return &CounterMetric{name: name, value: 0}
}

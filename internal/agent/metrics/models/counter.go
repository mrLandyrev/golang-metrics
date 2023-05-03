package models

import "strconv"

type CounterMetric struct {
	name  string
	value int64
}

func (counter *CounterMetric) GetName() string {
	return counter.name
}

func (counter *CounterMetric) GetKind() string {
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

func (counter *CounterMetric) GetValue() string {
	return strconv.FormatInt(counter.value, 10)
}

func NewCounterMetric(name string) *CounterMetric {
	return &CounterMetric{name: name, value: 0}
}

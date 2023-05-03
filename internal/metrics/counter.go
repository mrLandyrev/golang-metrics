package metrics

import "strconv"

type CounterMetric struct {
	NameData  string
	ValueData int64
}

func (counter *CounterMetric) Name() string {
	return counter.NameData
}

func (counter *CounterMetric) Kind() string {
	return "counter"
}

func (counter *CounterMetric) AddValue(value string) error {
	convertedValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return ErrIncorrectMetricValue
	}

	counter.ValueData += convertedValue

	return nil
}

func (counter *CounterMetric) Value() string {
	return strconv.FormatInt(counter.ValueData, 10)
}

func NewCounterMetric(name string) *CounterMetric {
	return &CounterMetric{NameData: name, ValueData: 0}
}

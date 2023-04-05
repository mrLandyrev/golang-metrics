package types

import "strconv"

type GaugeMetric struct {
	name  string
	value float64
}

func (gauge *GaugeMetric) GetName() string {
	return gauge.name
}

func (gauge *GaugeMetric) GetKind() string {
	return "gauge"
}

func (gauge *GaugeMetric) AddStrValue(value string) error {
	convertedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return ErrIncorrectMetricValue
	}

	gauge.value = convertedValue

	return nil
}

func (gauge *GaugeMetric) GetStrValue() string {
	return strconv.FormatFloat(gauge.value, 'E', -1, 64)
}

func NewGaugeMetric(name string) *GaugeMetric {
	return &GaugeMetric{name: name, value: 0}
}

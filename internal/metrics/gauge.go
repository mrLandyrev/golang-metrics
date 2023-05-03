package metrics

import "strconv"

type GaugeMetric struct {
	name  string
	value float64
}

func (gauge *GaugeMetric) Name() string {
	return gauge.name
}

func (gauge *GaugeMetric) Kind() string {
	return "gauge"
}

func (gauge *GaugeMetric) AddValue(value string) error {
	convertedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return ErrIncorrectMetricValue
	}

	gauge.value = convertedValue

	return nil
}

func (gauge *GaugeMetric) Value() string {
	return strconv.FormatFloat(gauge.value, 'f', -1, 64)
}

func NewGaugeMetric(name string) *GaugeMetric {
	return &GaugeMetric{name: name, value: 0}
}

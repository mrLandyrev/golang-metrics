package metrics

import "strconv"

type GaugeMetric struct {
	NameData  string
	ValueData float64
}

func (gauge *GaugeMetric) Name() string {
	return gauge.NameData
}

func (gauge *GaugeMetric) Kind() string {
	return "gauge"
}

func (gauge *GaugeMetric) AddValue(value string) error {
	convertedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return ErrIncorrectMetricValue
	}

	gauge.ValueData = convertedValue

	return nil
}

func (gauge *GaugeMetric) Value() string {
	return strconv.FormatFloat(gauge.ValueData, 'f', -1, 64)
}

func NewGaugeMetric(name string) *GaugeMetric {
	return &GaugeMetric{NameData: name, ValueData: 0}
}

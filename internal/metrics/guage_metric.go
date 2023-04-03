package metrics

import "strconv"

type GuageMetric struct {
	name  string
	value float64
}

func (guage *GuageMetric) GetName() string {
	return guage.name
}

func (guage *GuageMetric) GetKind() string {
	return "guage"
}

func (guage *GuageMetric) AddStrValue(value string) error {
	convertedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}

	guage.value = convertedValue

	return nil
}

func (guage *GuageMetric) GetStrValue() string {
	return strconv.FormatFloat(guage.value, 'E', -1, 64)
}

func NewGuageMetric(name string) *GuageMetric {
	return &GuageMetric{name: name, value: 0}
}

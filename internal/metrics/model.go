package metrics

type Metric interface {
	Name() string
	Kind() string
	AddValue(value string) error
	Value() string
}

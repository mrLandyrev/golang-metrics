package metrics

type Metric interface {
	GetName() string
	GetKind() string
	AddStrValue(value string) error
	GetStrValue() string
}

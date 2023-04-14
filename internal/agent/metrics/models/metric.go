package models

type Metric interface {
	GetName() string
	GetKind() string
	AddValue(value string) error
	GetValue() string
}

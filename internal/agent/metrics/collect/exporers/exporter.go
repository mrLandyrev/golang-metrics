package exporers

type Exporter interface {
	Collect() error
	Clear()
}

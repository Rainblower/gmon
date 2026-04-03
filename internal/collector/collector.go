package collector

type Metric struct {
	Name  string
	Value string
}

type Collector interface {
	Name() string
	Collect() ([]Metric, error)
}

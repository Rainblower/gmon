package monitor

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gmon/internal/collector"
	"gmon/internal/display"
)

type Monitor struct {
	collectors []collector.Collector
	writer     *display.Writer
}

func New(collectors ...collector.Collector) *Monitor {
	return &Monitor{
		collectors: collectors,
		writer:     &display.Writer{},
	}
}

func (m *Monitor) RunOnce() {
	m.collect()
	m.writer.Print()
}

func (m *Monitor) RunRealTime(interval time.Duration) {
	fmt.Print("\033[?1049h\033[?25l")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Print("\033[?25h\033[?1049l")
		os.Exit(0)
	}()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		m.writer.Reset()
		m.collect()
		m.writer.Flush()
		<-ticker.C
	}
}

type result struct {
	name    string
	metrics []collector.Metric
	err     error
}

func (m *Monitor) collect() {
	ch := make(chan result, len(m.collectors))

	for _, c := range m.collectors {
		go func(c collector.Collector) {
			metrics, err := c.Collect()
			ch <- result{name: c.Name(), metrics: metrics, err: err}
		}(c)
	}

	results := make([]result, 0, len(m.collectors))
	for range m.collectors {
		results = append(results, <-ch)
	}

	order := make(map[string]int)
	for i, c := range m.collectors {
		order[c.Name()] = i
	}

	sorted := make([]result, len(results))
	for _, r := range results {
		sorted[order[r.name]] = r
	}

	for _, r := range sorted {
		if r.err != nil {
			m.writer.Render(r.name, []collector.Metric{{Name: "Error", Value: r.err.Error()}})
			continue
		}
		m.writer.Render(r.name, r.metrics)
	}
}

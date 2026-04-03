package collector

import (
	"fmt"
	"gmon/internal/utils"

	"github.com/shirou/gopsutil/v4/mem"
)

type MemoryCollector struct{}

func (m *MemoryCollector) Name() string {
	return "Memory"
}

func (m *MemoryCollector) Collect() ([]Metric, error) {
	v, err := mem.VirtualMemory()

	if err != nil {
		return nil, fmt.Errorf("failed to get memory info: %w", err)
	}

	return []Metric{
		{Name: "Used", Value: utils.FormatBytes(v.Used)},
		{Name: "Total", Value: utils.FormatBytes(v.Total)},
		{Name: "Usage", Value: fmt.Sprintf("%.1f%%", v.UsedPercent)},
	}, nil
}

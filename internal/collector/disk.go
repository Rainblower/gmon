package collector

import (
	"fmt"
	"gmon/internal/utils"

	"github.com/shirou/gopsutil/v4/disk"
)

type DiskCollector struct{}

func (m *DiskCollector) Name() string {
	return "Disk"
}

func (m *DiskCollector) Collect() ([]Metric, error) {

	partitions, err := disk.Partitions(false)

	if err != nil {
		return nil, fmt.Errorf("failed to get disk partitions: %w", err)
	}

	metrics := []Metric{
		{Name: "Filesystem", Value: fmt.Sprintf("%-11s %-11s %-11s %5s  Mounted on", "Size", "Used", "Avail", "Use%")},
	}

	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue // пропускаем недоступные разделы
		}

		metrics = append(metrics, Metric{
			Name: p.Device,
			Value: fmt.Sprintf("%-11s %-11s %-11s %4.1f%%  %s",
				utils.FormatBytes(usage.Total),
				utils.FormatBytes(usage.Used),
				utils.FormatBytes(usage.Free),
				usage.UsedPercent,
				p.Mountpoint,
			),
		})
	}

	return metrics, nil
}

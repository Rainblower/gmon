package collector

import (
	"fmt"
	"gmon/internal/utils"
	"strings"

	"github.com/shirou/gopsutil/v4/host"
)

type InfoCollector struct{}

func (m *InfoCollector) Name() string {
	return "Info"
}

func (m *InfoCollector) Collect() ([]Metric, error) {
	info, err := host.Info()

	if err != nil {
		return nil, err
	}

	platform := utils.Capitalize(info.Platform)
	OS := utils.Capitalize(info.OS)

	fullOSName := fmt.Sprintf("%s %s %s", platform, OS, info.KernelArch)

	uptime := m.getUpTime(info.Uptime)

	return []Metric{
		{Name: "OS", Value: fullOSName},
		{Name: "Host", Value: info.Hostname},
		{Name: "Kernel", Value: info.KernelVersion},
		{Name: "Uptime", Value: uptime},
	}, nil
}

func (m *InfoCollector) getUpTime(seconds uint64) string {
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24

	remainingHours := hours % 24
	remainingMinutes := minutes % 60

	var parts []string
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%d days", days))
	}
	if remainingHours > 0 {
		parts = append(parts, fmt.Sprintf("%d hours", remainingHours))
	}
	parts = append(parts, fmt.Sprintf("%d mins", remainingMinutes))

	return strings.Join(parts, " ")
}

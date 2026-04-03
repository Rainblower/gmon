package collector

import (
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/sensors"
)

type CPUCollector struct{}

func (m *CPUCollector) Name() string {
	return "CPU"
}

var cpuMainKeys = []string{"coretemp_packageid0", "k10temp_tctl", "k10temp_tdie", "zenpower_tdie"}

func (m *CPUCollector) Collect() ([]Metric, error) {
	info, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	var metrics []Metric

	// CPU Name
	if len(info) > 0 {
		metrics = append(metrics, Metric{Name: "CPU", Value: info[0].ModelName})
	}

	// CPU Temperature
	temps, _ := sensors.SensorsTemperatures()
	for _, t := range temps {
		for _, key := range cpuMainKeys {
			if strings.ToLower(t.SensorKey) == key {
				metrics = append(metrics, Metric{Name: "CPU Temp", Value: fmt.Sprintf("%.1f°C", t.Temperature)})
			}
		}
	}

	// CPU per core usage
	percents, _ := cpu.Percent(0, true)
	for i, p := range percents {
		metrics = append(metrics, Metric{Name: fmt.Sprintf("Core %d", i), Value: fmt.Sprintf("%.1f%%", p)})
	}

	return metrics, nil
}

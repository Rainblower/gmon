package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"gmon/internal/collector"
	"gmon/internal/monitor"
)

var (
	flagInfo     bool
	flagCPU      bool
	flagMemory   bool
	flagDisks    bool
	flagRealTime bool
)

var rootCmd = &cobra.Command{
	Use:     "gmon",
	Version: "1.0.0",
	Short:   "A lightweight system resource monitor written in Go",
	Long: `gmon is a simple and fast command-line tool for real-time monitoring 
of system resources: CPU, memory, disk.`,
	// Run вызывается когда пользователь запускает программу
	Run: func(cmd *cobra.Command, args []string) {
		all := []collector.Collector{
			&collector.InfoCollector{},
			&collector.CPUCollector{},
			&collector.MemoryCollector{},
			&collector.DiskCollector{},
		}

		var collectors []collector.Collector
		noFlags := !flagCPU && !flagMemory && !flagInfo && !flagDisks
		if noFlags {
			collectors = all
		} else {
			if flagInfo {
				collectors = append(collectors, &collector.InfoCollector{})
			}
			if flagCPU {
				collectors = append(collectors, &collector.CPUCollector{})
			}
			if flagMemory {
				collectors = append(collectors, &collector.MemoryCollector{})
			}
			if flagDisks {
				collectors = append(collectors, &collector.DiskCollector{})
			}
		}

		m := monitor.New(collectors...)

		if flagRealTime || noFlags {
			m.RunRealTime(1 * time.Second)
		} else {
			m.RunOnce()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&flagInfo, "info", "i", false, "show system info")
	rootCmd.Flags().BoolVarP(&flagCPU, "cpu", "c", false, "show cpu info")
	rootCmd.Flags().BoolVarP(&flagDisks, "disk", "d", false, "show disk info")
	rootCmd.Flags().BoolVarP(&flagMemory, "memory", "m", false, "show memory info")
	rootCmd.Flags().BoolVarP(&flagRealTime, "realtime", "r", false, "real-time monitoring")
}

package utils

import "fmt"

func FormatBytes(b uint64) string {
	const gb = 1 << 30
	const mb = 1 << 20

	if b >= gb {
		return fmt.Sprintf("%.2f GB", float64(b)/gb)
	}
	return fmt.Sprintf("%.2f MB", float64(b)/mb)
}

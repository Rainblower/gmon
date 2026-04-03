package display

import (
	"fmt"
	"os"
	"strings"

	"gmon/internal/collector"

	"golang.org/x/term"
)

type Writer struct {
	buf strings.Builder
}

// Reset сбрасывает буфер перед новым кадром.
func (w *Writer) Reset() {
	w.buf.Reset()
}

func (w *Writer) Flush() {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || width <= 0 {
		width = 80
	}

	lines := strings.Split(w.buf.String(), "\n")
	var sb strings.Builder
	sb.WriteString("\033[H")
	for _, line := range lines {
		if len(line) > width {
			line = line[:width]
		}
		sb.WriteString("\033[2K\r")
		sb.WriteString(line)
		sb.WriteString("\r\n")
	}
	sb.WriteString("\033[J")
	fmt.Print(sb.String())
}

func (w *Writer) Print() {
	fmt.Print(w.buf.String())
}

func (w *Writer) PrintHeader(header string) {
	fmt.Fprintln(&w.buf, header)
}

func (w *Writer) Render(groupName string, metrics []collector.Metric) {
	fmt.Fprintf(&w.buf, "=== %s ===\n", groupName)

	for _, m := range metrics {
		fmt.Fprintf(&w.buf, "  %-20s %s\n", m.Name+":", m.Value)
	}
}

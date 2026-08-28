package report

import (
	"strings"
	"waveboard/internal/domain"
)

func Table(progress []domain.Progress) string {
	lines := []string{"wave,pending,picked,review,shortage,packed"}
	for _, p := range progress {
		lines = append(lines, strings.Join([]string{p.WaveID, itoa(p.Pending), itoa(p.Picked), itoa(p.Review), itoa(p.Shortage), itoa(p.Packed)}, ","))
	}
	return strings.Join(lines, "\n")
}
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	buf := ""
	for n > 0 {
		buf = string(rune('0'+n%10)) + buf
		n /= 10
	}
	return sign + buf
}

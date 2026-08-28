package report

import (
	"fmt"
	"waveboard/internal/domain"
)

func Summary(p domain.Progress) string {
	return fmt.Sprintf("wave %s pending=%d picked=%d review=%d shortage=%d packed=%d", p.WaveID, p.Pending, p.Picked, p.Review, p.Shortage, p.Packed)
}
func Completion(p domain.Progress) float64 {
	total := p.Pending + p.Picked + p.Review + p.Shortage + p.Packed
	if total == 0 {
		return 0
	}
	return float64(p.Packed) / float64(total)
}
func NeedsAttention(p domain.Progress) bool { return p.Shortage > 0 || p.Review > p.Picked }

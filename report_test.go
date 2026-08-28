package waveboard

import (
	"testing"
	"waveboard/internal/domain"
	"waveboard/internal/report"
)

func TestReport(t *testing.T) {
	p := domain.Progress{WaveID: "W", Packed: 2, Pending: 2}
	if report.Completion(p) != .5 {
		t.Fatal("completion")
	}
	if !report.NeedsAttention(domain.Progress{Shortage: 1}) {
		t.Fatal("attention")
	}
}

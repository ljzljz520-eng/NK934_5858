package waveboard

import (
	"testing"
	"waveboard/internal/domain"
)

func TestDomainRules(t *testing.T) {
	if !domain.CanTransition(domain.StatusPending, domain.StatusPicked) {
		t.Fatal("transition")
	}
	if domain.CanTransition(domain.StatusPacked, domain.StatusPending) {
		t.Fatal("terminal")
	}
	if domain.ValidStatus("bad") {
		t.Fatal("status")
	}
}

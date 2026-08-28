package domain

import "fmt"

func ValidateWave(w Wave) error {
	if w.ID == "" {
		return fmt.Errorf("wave id required")
	}
	if w.Zone == "" {
		return fmt.Errorf("zone required")
	}
	if w.Priority < 0 {
		return fmt.Errorf("priority invalid")
	}
	return nil
}
func ValidateTask(t PickTask) error {
	if t.ID == "" || t.WaveID == "" {
		return fmt.Errorf("task identity required")
	}
	if t.Quantity <= 0 {
		return fmt.Errorf("quantity invalid")
	}
	if t.Location == "" || t.SKU == "" {
		return fmt.Errorf("location and sku required")
	}
	return nil
}
func ValidStatus(s string) bool {
	switch s {
	case StatusPending, StatusPicked, StatusReview, StatusShortage, StatusPacked:
		return true
	default:
		return false
	}
}

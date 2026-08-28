package fixture

import (
	"waveboard/internal/domain"
	"waveboard/internal/service"
)

func Seed(m *service.Manager) error {
	for _, w := range []domain.Wave{{ID: "W-100", Zone: "A", Status: "open", Priority: 1}, {ID: "W-200", Zone: "B", Status: "open", Priority: 2}} {
		if e := m.CreateWave(w); e != nil {
			return e
		}
	}
	tasks := []domain.PickTask{{ID: "T-1", WaveID: "W-100", Location: "A-01", SKU: "SKU-RED", Quantity: 2, Status: domain.StatusPending}, {ID: "T-2", WaveID: "W-100", Location: "A-02", SKU: "SKU-BLU", Quantity: 1, Status: domain.StatusPicked}, {ID: "T-3", WaveID: "W-200", Location: "B-01", SKU: "SKU-GRN", Quantity: 4, Status: domain.StatusReview}, {ID: "T-4", WaveID: "W-200", Location: "B-02", SKU: "SKU-YEL", Quantity: 3, Status: domain.StatusShortage}}
	for _, t := range tasks {
		if e := m.AddTask(t); e != nil {
			return e
		}
	}
	return nil
}

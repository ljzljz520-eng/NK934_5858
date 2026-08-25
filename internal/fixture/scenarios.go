package fixture

import "waveboard/internal/domain"

func ScenarioTasks() []domain.PickTask {
	return []domain.PickTask{{ID: "S-1", WaveID: "W-100", Location: "A-03", SKU: "SKU-001", Quantity: 1, Status: domain.StatusPending}, {ID: "S-2", WaveID: "W-100", Location: "A-04", SKU: "SKU-002", Quantity: 2, Status: domain.StatusPacked}}
}
func ScenarioWaves() []domain.Wave {
	return []domain.Wave{{ID: "W-300", Zone: "C", Status: "open", Priority: 3}, {ID: "W-400", Zone: "D", Status: "open", Priority: 4}}
}

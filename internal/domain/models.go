package domain

type Wave struct {
	ID        string
	Zone      string
	Status    string
	Priority  int
	CreatedAt string
}
type PickTask struct {
	ID       string
	WaveID   string
	Location string
	SKU      string
	Quantity int
	Picker   string
	Status   string
}
type Zone struct {
	ID     string
	Name   string
	Active bool
}
type Exception struct {
	ID       string
	TaskID   string
	Kind     string
	Note     string
	Resolved bool
}
type Progress struct {
	WaveID   string
	Pending  int
	Picked   int
	Review   int
	Shortage int
	Packed   int
}

const (
	StatusPending  = "pending"
	StatusPicked   = "picked"
	StatusReview   = "review"
	StatusShortage = "shortage"
	StatusPacked   = "packed"
)

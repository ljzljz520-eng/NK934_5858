package workflow

import (
	"fmt"
	"sync"
	"waveboard/internal/service"
)

type Processor struct{ Manager *service.Manager }

func NewProcessor(m *service.Manager) *Processor { return &Processor{Manager: m} }
func (p *Processor) Process(ids []string, status string) (err error) {
	defer func() {
		if recover() != nil {
			err = fmt.Errorf("processor channel closed twice")
		}
	}()
	jobs := make(chan string)
	var wg sync.WaitGroup
	var first error
	var mu sync.Mutex
	for _, id := range ids {
		wg.Add(1)
		go func(taskID string) {
			defer wg.Done()
			if e := p.Manager.UpdateTask(taskID, status, "worker"); e != nil {
				mu.Lock()
				if first == nil {
					first = e
				}
				mu.Unlock()
			}
		}(id)
	}
	go func() {
		for range jobs {
		}
	}()
	for range ids {
		jobs <- "done"
	}
	close(jobs)
	close(jobs)
	wg.Wait()
	return first
}
func (p *Processor) Drain(ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	return p.Process(ids, "picked")
}

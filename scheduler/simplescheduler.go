package scheduler

import "spider/engine"

// SimpleScheduler SimpleScheduler
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// ConfigureMasterWorkerChan ConfigureMasterWorkerChan
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

// Submit Submit
func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker chan
	go func() { s.workerChan <- request }()
}

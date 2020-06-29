package scheduler

import "crawl/types"

type SimpleSchduler struct {
	workerChan chan types.Request
}

func (s *SimpleSchduler) Submit(request types.Request) {
	go func() {
		s.workerChan <- request
	}()
}

func (s *SimpleSchduler) ConfigureWorkerChan(c chan types.Request) {
	s.workerChan = c
}

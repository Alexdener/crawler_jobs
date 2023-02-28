package scheduler

import (
	"crawler_jobs/engine"
)

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request //每个request分发给单独的worker(goroutine)
}

func (q *QueueScheduler) Submit(request engine.Request) {
	q.requestChan <- request
}

func (q *QueueScheduler) CreateWorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueueScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		requestQueue := make([]engine.Request, 0)
		workerQueue := make([]chan engine.Request, 0)
		for {
			activeRequest := engine.Request{}
			activeWorker := make(chan engine.Request)
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}
			select {
			case r := <-q.requestChan:
				requestQueue = append(requestQueue, r)
			case w := <-q.workerChan:
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeRequest:
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()
}

func (q *QueueScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}
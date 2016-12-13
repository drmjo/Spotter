package main

import "fmt"

// Note: just handling request type for the time being.
type Dispatcher struct {
	WorkerPool chan chan string
	work       WorkRequest
}

var ResultChannel chan WorkResponse

func NewDispatcher(work *WorkRequest) *Dispatcher {
	pool := make(chan chan string, (*work).NumberOfWorkers)
	return &Dispatcher{WorkerPool: pool, work: *work}
}

func (d *Dispatcher) Run() {

	ResultChannel := make(chan WorkResponse, d.work.NumberOfRequests)
	// start goroutines equal to worker number then go through request count and handout requests.
	for i := 0; i < d.work.NumberOfWorkers; i++ {
		fmt.Printf("Starting Worker: %d \n", i+1)
		worker := NewWorker(i, d.WorkerPool)
		worker.Start(ResultChannel)
	}

	go d.dispatch()
	for i := 0; i < d.work.NumberOfRequests; i++ {
		result := <-ResultChannel
		fmt.Printf("%d %s\n", result.id, result.message)
	}
}

func (d *Dispatcher) dispatch() {
	for i := 0; i < d.work.NumberOfRequests; i++ {
		go func(work WorkRequest) {
			jobChannel := <-d.WorkerPool

			jobChannel <- d.work.RequestVerb
		}(d.work)
	}
}

package main

import (
	"fmt"
	"net/http"
)

type Dispatcher struct {
	WorkerPool chan chan WorkRequest
	work       WorkRequest
}

var ResultChannel chan WorkResponse
var client *http.Client

func NewDispatcher(work *WorkRequest) *Dispatcher {
	pool := make(chan chan WorkRequest, (*work).NumberOfWorkers)
	return &Dispatcher{WorkerPool: pool, work: *work}
}

func (d *Dispatcher) Run() {

	ResultChannel := make(chan WorkResponse, d.work.NumberOfRequests)
	// NOTE: Can set redirect policy here and other http client settings.
	client := &http.Client{}
	// NOTE: start goroutines equal to worker number then go through request count and handout requests.
	for i := 0; i < d.work.NumberOfWorkers; i++ {
		fmt.Printf("Starting Worker: %d \n", i+1)
		worker := NewWorker(i, d.WorkerPool)
		worker.Start(ResultChannel)
	}

	go d.dispatch()
	for i := 0; i < d.work.NumberOfRequests; i++ {
		result := <-ResultChannel
		// NOTE: should collect results in some structure here then do aggregations after the fact.
		fmt.Printf("%d %s\n", result.id, result.message)
	}
}

func (d *Dispatcher) dispatch() {
	for i := 0; i < d.work.NumberOfRequests; i++ {
		go func(work WorkRequest) {
			jobChannel := <-d.WorkerPool

			jobChannel <- d.work
		}(d.work)
	}
}
